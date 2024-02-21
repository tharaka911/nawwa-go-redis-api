#### **pre requested
```
install golang on your system
install git
```
##### install go in ubuntu server
```sh
cd /tmp && wget https://go.dev/dl/go1.21.5.linux-amd64.tar.gz && tar -C /usr/local -xzf go1.21.5.linux-amd64.tar.gz
```

```
export PATH=$PATH:/usr/local/go/bin && go version
or
cd
nano .bashrc
add below line
export PATH=$PATH:/usr/local/go/bin
source .bashrc
```
#### 1 clone the code
```sh
git clone https://github.com/tharaka911/nawwa-go-redis-api.git
```
#### 2 run the application
```sh
go run main.go
```

#### 3 build the binaries
```sh
go build .
```
