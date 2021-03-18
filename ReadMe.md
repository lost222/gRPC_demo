
[toc]



## 使用流程

1. 使用proto定义数据和服务形式
2. 使用protoc和go相关插件，.proto编译为pb.go
3. import pb ./pb.go, 实现server接口，定义服务的行为





## proto

Protobuf是Protocol Buffers的简称，它是Google公司开发的一种数据描述语言，用于描述一种轻便高效的结构化数据存储格式



### 定义服务
~~~protobuf
service StringService{
  rpc Concat(StringRequest) returns (StringResponse) {};
  rpc Diff(StringRequest) returns (StringResponse) {};
  rpc ServerStream(StringRequest) return (Stream StringResponse) {};
  rpc ClientStream(Stream StringRequest) return (StringResponse) {};
  rpc C_SStream(Stream StringRequest) return (Stream StringResponse) {};
}
~~~
service定义了哪些服务通过gRPC进行，实际上在编译为.go之后会变成各种方法。

传输的`StringRequest` 和`StringResponse`是message形式

 流式编程。 依据Stream关键字的位置分为

* 一元
* 服务端
* 客户端
* 双向

区别是一元必须是同步的RPC， 停等协议。 而拿到Stream的可以通过Steam结构体多次发送数据，类似TCP流水线。 gRPC的stream好用，是因为它基于HTTP2。

#### HTTP2 

* stream并发：

  考虑一个情况， client端顺序发送3个请求A,B,C。A请求的数据量最大。 在HTTP中，会分别请求和断开， HTTP头部要发三次，TCP连接也要连接三次。 所以在HTTP1.1里，允许在一个TCP连接上发送多个请求。 节省一部分空间。 但是问题就是B和C要等着A传输完了才能开始传输。 而A又很大，B和C又可能是更急切的需求。 

  HTTP2提供新的抽象层Stream，这个时候允许为A，B,  C分别开一个Stream。 并发请求。 但是TCP底层还是只有一个Socket，所以实际上操作是轮询式地从三个Stream里发送包。 这里有一个新问题，就是如果轮询的时候Astream事实上接收了过大的数据，导致TCP流量控制了，那B和C虽然原理上并发，但是事实上又阻塞了。 所以Stream也是有流量控制的。 和TCP流量控制的目的一样： 各个Stream尽量公平。 



###  定义消息

~~~protobuf
message SearchRequest {
  string query = 1;
  int32 page_number = 2;
  int32 result_per_page = 3;
}
~~~

message定义的是传递的数据形式， 编译为go之后会成为struct。

* 类型： string int32
* 编号：在message定义中每个字段都有一个唯一的编号，这些编号被用来在二进制消息体中识别你定义的这些字段
* 枚举类型
* 使用message作为message

~~~protobuf
message SearchResponse {
  repeated Result results = 1;
}

message Result {
  string url = 1;
  string title = 2;
  repeated string snippets = 3;
}
~~~


### 编译


~~~shell
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative helloworld/helloworld.proto
~~~





## client代码

1. 调用grpc.Dial 建立网络连接
2. 创建gRPC client(protoBuff生成)
3. 调用client.Concat进行RPC调用



## server代码

1. 置顶grpc的监听IP和端口
2. 调用grpc.Newserver()建立RPC server
3. 调用生成的pb.go中的Register---server来注册服务
4. grpc.Server监听



## service struct ： 对pb.go中的server接口的实现

实际上就是要实现.proto文件里的方法。即`Concat`,`diff`

定义服务的行为

