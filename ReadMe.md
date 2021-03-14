@todo 命令行用处

~~~shell
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative helloworld/helloworld.proto
~~~


## proto
### service
~~~protobuf
service StringService{
  rpc Concat(StringRequest) returns (StringResponse) {};
  rpc Diff(StringRequest) returns (StringResponse) {};
}
~~~
service定义了哪些服务通过gRPC进行




## client代码

1. 调用grpc.Dial 建立网络连接
2. 创建gRPC client(protoBuff生成)
3. 调用client.Concat进行RPC调用