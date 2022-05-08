package causal

import (
	"context"
	"encoding/json"
	"fmt"
	config "hckvstore/config"
	Per "hckvstore/persister"
	RPC "hckvstore/rpc/causalrpc"
	"hckvstore/util"
	"log"
	"math/rand"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type CausalEntity struct {
	vectorClock sync.Map
	members     []string
	delay       int
	address     string
	persister   *Per.Persister
	// client      RPC.CAUSALClient
	// me          int32
	rwmu *sync.RWMutex

	// waitCh  chan config.OpCausalVector
	applyCh chan int
	log     []Log
}

type Log struct {
	Command config.Op
}

func (ce *CausalEntity) MergeVC(vc sync.Map) {
	vc.Range(func(k, v interface{}) bool {
		val, ok := ce.vectorClock.Load(k)
		if !ok {
			ce.vectorClock.Store(k, v)
		} else {
			if v.(int32) > val.(int32) {
				ce.vectorClock.Store(k, v)
			}
		}
		return true
	})
	// for member := range vc {
	// 	if vc[member] > ce.vectorClock[member] {
	// 		ce.vectorClock[member] = vc[member]
	// 	}
	// }
}

func (ce *CausalEntity) IsUpper(vc sync.Map) bool {
	util.DPrintf("IsUpper(): ce.vc: %v, arg_vc: %v", util.BecomeMap(ce.vectorClock), util.BecomeMap(vc))
	if util.Len(vc) == 0 {
		return true
	}
	if util.Len(ce.vectorClock) != util.Len(vc) {
		return false
	} else {
		// vc_temp := ce.vectorClock
		// for member := range vc_temp {
		// 	// key在vc中存在
		// 	if _, ok := vc[member]; ok {
		// 		if vc_temp[member] >= vc[member] {
		// 			continue
		// 		} else {
		// 			return false
		// 		}
		// 	}
		// 	return false
		// }
		// return true
		res := true
		ce.vectorClock.Range(func(key, value interface{}) bool {
			v, ok := vc.Load(key)
			if ok {
				if value.(int32) >= v.(int32) {
					return true
				} else {
					res = false
					return false
				}
			}
			res = false
			return false
		})
		return res
	}
}

// this method must be called by KV Server instead of CausalEntity
func (ce *CausalEntity) Start(command interface{}, vcFromClientArg map[string]int32) (map[string]int32, bool) {
	// ce.mu.Lock()
	// defer ce.mu.Unlock()
	vcFromClient := util.BecomeSyncMap(vcFromClientArg)
	newLog := Log{
		command.(config.Op),
	}
	util.DPrintf("Log in Start(): %v ", newLog)
	util.DPrintf("vcFromClient in Start(): %v", vcFromClient)
	if newLog.Command.Option == "Put" {
		isUpper := ce.IsUpper(vcFromClient)
		if isUpper {
			// ce.vectorClock[ce.address+"1"] += 1
			val, _ := ce.vectorClock.Load(ce.address + "1")
			ce.vectorClock.Store(ce.address+"1", val.(int32)+1)
			data, _ := json.Marshal(newLog)
			args := &RPC.AppendEntriesInCausalArgs{
				Log:         data,
				VectorClock: util.BecomeMap(ce.vectorClock),
			}
			for i := 0; i < len(ce.members); i++ {
				if ce.members[i] != ce.address {
					go ce.sendAppendEntriesInCausal(ce.members[i], args)
				}
			}
			ce.log = append(ce.log, newLog)
			ce.persister.Put(newLog.Command.Key, newLog.Command.Value)
			ce.applyCh <- 1
			util.DPrintf("applyCh unread buffer: %v", len(ce.applyCh))
			return util.BecomeMap(ce.vectorClock), true
		} else {
			return nil, false
		}
	} else if newLog.Command.Option == "Get" {
		if util.Len(vcFromClient) == 0 {
			return nil, true
		} else {
			if ce.IsUpper(vcFromClient) {
				return nil, true
			}
			return nil, false
		}
	}
	util.DPrintf("here is Start() in Causal: log command option is false")
	return nil, false
}

// s0 --> other servers
func (ce *CausalEntity) sendAppendEntriesInCausal(address string, args *RPC.AppendEntriesInCausalArgs) (*RPC.AppendEntriesInCausalReply, bool) {
	fmt.Println("here is sendAppendEntriesInCausal() ---------> ", address)
	// 随机等待，模拟延迟
	time.Sleep(time.Millisecond * time.Duration(ce.delay+rand.Intn(25)))
	// conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := RPC.NewCAUSALClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	reply, err := client.AppendEntriesInCausal(ctx, args)
	if err != nil {
		fmt.Println(" sendAppendEntriesInCausal could not greet: ", err, address)
		return reply, false
	}
	return reply, true
}

func (ce *CausalEntity) AppendEntriesInCausal(ctx context.Context, args *RPC.AppendEntriesInCausalArgs) (reply *RPC.AppendEntriesInCausalReply, err error) {
	// ce.mu.Lock()
	// defer ce.mu.Unlock()
	fmt.Println("==============AppendEntriesInCausal RPC Call From Others==============")
	util.DPrintf("AppendEntriesInCausalArgs: %v", args)
	reply = &RPC.AppendEntriesInCausalReply{}
	remoteVC := util.BecomeSyncMap(args.VectorClock)
	ok := ce.IsUpper(remoteVC)
	if ok {
		// 本地vc更大，忽略请求
		reply.Success = true
		return reply, nil
	} else {
		// 本地vc更小或者mix
		// merge remote vc and apply cmd
		ce.MergeVC(remoteVC)
		// 解析args
		var log Log
		json.Unmarshal(args.Log, &log)
		ce.log = append(ce.log, log)
		if log.Command.Option == "Put" {
			ce.persister.Put(log.Command.Key, log.Command.Value)
		}
		return reply, nil
	}
}

func (ce *CausalEntity) RegisterServer(address string) {
	// Register Server
	for {
		lis, err := net.Listen("tcp", address)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		RPC.RegisterCAUSALServer(s, ce)
		// Register reflection service on gRPC server.
		reflection.Register(s)
		if err := s.Serve(lis); err != nil {
			fmt.Printf("failed to serve: %v \n", err)
		}
	}

}

func MakeCausalEntity(add string, mem []string, persist *Per.Persister, mu *sync.RWMutex, applyCh chan int, delay int) *CausalEntity {
	ce := &CausalEntity{}
	// if len(mem) <= 1 {
	// 	panic("#######Address is less 1, you should set follower's address!######")
	// }
	ce.address = add
	ce.persister = persist
	ce.applyCh = applyCh
	ce.delay = delay
	ce.rwmu = mu
	// for concurrent map write and read
	ce.members = make([]string, len(mem))
	// ce.vectorClock = make(sync.Map)
	// ce.waitCh = make(chan config.Op, 100)
	for i := 0; i < len(mem); i++ {
		ce.members[i] = mem[i]
		// 因为client记录的是30011端口，而非3001端口，统一记录
		add_temp := mem[i] + "1"
		ce.vectorClock.Store(add_temp, int32(0))
	}
	fmt.Println("members: ", ce.members)
	fmt.Println("vectorClock: ", util.BecomeMap(ce.vectorClock))
	go ce.RegisterServer(ce.address)
	// 省略了后台的心跳检测
	// 通过Start()
	return ce
}

func MakeTestVC(vectorClock sync.Map) *CausalEntity {
	ce := &CausalEntity{}
	ce.vectorClock = vectorClock
	return ce
}

func (ce *CausalEntity) PrintVC() {
	fmt.Println(util.BecomeMap(ce.vectorClock))
}
