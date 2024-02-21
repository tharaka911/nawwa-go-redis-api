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
### how to run go binaries as a service in linux

```sh
nano /lib/systemd/system/nawwa-go-redis-api.service
```
#### adding the below lines to golanghttp.service
```sh
[Unit]
Description=nawwa-go-redis-api
[Service]
Type=simple
Restart=always
RestartSec=5s
ExecStart=<executable location>/go-redis-api
WorkingDirectory=<executable location>
[Install]
WantedBy=multi-user.target
```
#### View the logs
```sh
last logs -> journalctl -xeu nawwa-go-redis-api
```
```sh
live logs -> journalctl -xeu nawwa-go-redis-api -f
```
#### service manipulation 
```sh
service nawwa-go-redis-api start
```
```sh
service nawwa-go-redis-api status
```
```sh
service nawwa-go-redis-api stop
```
#### more infor
```sh
https://medium.com/@mnabilarta/deploying-go-app-on-ubuntu-server-e5d1e45162ca
```

