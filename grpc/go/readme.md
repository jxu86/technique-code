`环境:`     
`系统: mac`     
`Protocol Buffers v3.11.4`



* protocal buffer安装

从`https://github.com/google/protobuf/releases`下周解压包，解压然后执行
```
./configure --prefix=/usr/local/protobuf
sudo make
sudo make install
```
最后把环境变量写到/etc/profile      
export PROTOBUF=/usr/local/protobuf     
export PATH=$PROTOBUF/bin:$PATH

```
source /etc/profile
```


* golang protobuf安装

```
go get -u github.com/golang/protobuf/proto // golang protobuf 库
go get -u github.com/golang/protobuf/protoc-gen-go //protoc --go_out 工具
```

* 安装 gRPC-go

```
go get google.golang.org/grpc
```


* 生成客户端和服务器端代码  
进入routeguide目录
```
protoc --go_out=plugins=grpc:. route_guide.proto
```

参考:   
[gRPC 官方文档中文版](http://doc.oschina.net/grpc?t=60133)