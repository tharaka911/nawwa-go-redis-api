##### install go in ubuntu server
```sh
cd /tmp && wget https://go.dev/dl/go1.22.2.linux-amd64.tar.gz && tar -C /usr/local -xzf go1.22.2.linux-amd64.tar.gz
```

##### open the bashrc under root user
```sh
cd && nano .bashrc
```

##### add below line at the below of .bashrc
```sh
export PATH=$PATH:/usr/local/go/bin
```
##### save and exit

##### reload the bashrc
```sh
source .bashrc
```
#### 1. clone the code
```sh
git clone https://github.com/tharaka911/nawwa-go-redis-api.git
```
#### 2. run the application
```sh
go run main.go
```
#### 3. build the binaries
```sh
go build main.go
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
ExecStart=<executable location>/main
WorkingDirectory=<executable location>
[Install]
WantedBy=multi-user.target
```
#### View the logs
##### last logs -> 
```sh
journalctl -xeu nawwa-go-redis-api
```
##### live logs -> 
```sh
journalctl -xeu nawwa-go-redis-api -f
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


