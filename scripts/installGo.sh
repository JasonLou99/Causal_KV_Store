wget https://studygolang.com/dl/golang/go1.15.4.linux-amd64.tar.gz
sudo tar -zxvf go1.15.4.linux-amd64.tar.gz -C ~/
mkdir ~/Go
echo "export GOPATH=~/Go" >>  ~/.bashrc 
echo "export GOROOT=~/go" >> ~/.bashrc 
echo "export GOTOOLS=\$GOROOT/pkg/tool" >> ~/.bashrc
echo "export PATH=\$PATH:\$GOROOT/bin:\$GOPATH/bin" >> ~/.bashrc
source ~/.bashrc
