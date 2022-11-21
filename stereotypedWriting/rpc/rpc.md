### GO RPC
#### grpc 框架
##### 四类服务方法
  - 单项 RPC  
  即客户端发送一个请求给服务端，从服务端获取一个应答，就像一次普通的函数调用。
  - 流式传输
    - HTTP/2中有两个概念，流（stream）与帧（frame），其中帧作为HTTP/2中通信的最小传输单位，
    通常一个请求或响应会被分为一个或多个帧传输，流则表示已建立连接的虚拟通道，可以传输多次请求或响应。
    每个帧中包含Stream Identifier，标志所属流。HTTP/2通过流与帧实现多路复用，对于相同域名的请求，
    通过Stream Identifier标识可在同一个流中进行，从而减少连接开销。 而gRPC基于HTTP/2协议传输，
    自然而然也实现了流式传输，其中gRPC中共有以下三种类型的流
    - 方式
      - 服务端流式 RPC
      即客户端发送一个请求给服务端，可获取一个数据流用来读取一系列消息。客户端从返回的数据流里一直读取直到没有更多消息为止。
      - 客户端流式 RPC  
      即客户端用提供的一个数据流写入并发送一系列消息给服务端。一旦客户端完成消息写入，就等待服务端读取这些消息并返回应答。
      - 双向流式 RPC  
      即两边都可以分别通过一个读写数据流来发送一系列消息。这两个数据流操作是相互独立的，
      所以客户端和服务端能按其希望的任意顺序读写，例如：服务端可以在写应答前等待所有的客户端消息，或者它可以先读一个消息再写一个消息，或者是读写相结合的其他方式。每个数据流里消息的顺序会被保持。
##### 网络协议-http2
##### [etcd](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=MzU1OTIzOTE0Mw==&action=getalbum&album_id=1346747532856311809&scene=173&from_msgid=2247484327&from_itemidx=1&count=3&nolastread=1#wechat_redirect)
  
##### 序列化-protobuf
    - varint算法，二进制 
    - 打tag 数字 varint压缩
    
##### 分布式
  - [etcd在grpc重的实现](https://blog.csdn.net/weixin_44120629/article/details/108989750)
  - grpc+etcd
    - resolver 服务名称解析
      - 客户端是如何从一个域名/服务名，获取到其对应的实例ip，然后与之建立连接的呢？
        - resolver.Get(scheme) 通过传入实例名称或者配置etcd客户端
        - 底层http2连接对应的是一个grpc的stream，而stream的创建有两种方式，一种就是我们主动去创建一个stream池，这样当有请求需要发送时，我们可以直接使用我们创建好的stream
      - 运行过程中，如果后端的实例挂了，grpc如何感知到，并重新建立连接呢？
        - resetTransport, tryAllAddrs
          `首先，已经连接的后端发生了故障；
           然后，已经建立的http2client读到了连接goAway；
           再然后，resetTransport进入一次新的循环，重新获取解析结果；
           最后，resetTransport里通过新获取到的地址，重新建立连接`
        - resolver.watch 监听
  - 服务发现
    - etcd + resolver
    - 启动rpc服务-去建立租约 设置ttl
    - 服务发现怎么知道服务是否在线: 定时发送keepalive 心跳 续租 lease 
    - 客户端怎么知道某些rpc服务是否在线： watch 监听etcd的key value变化
        `etcd 启动时会注册 WatchServer[1], pb.WatchServer 用于处理 watch 请求
         接收 watch 请求
         每一个 watch 流都创建一个 serverWatchStream 结构体
         开启两个 goroutine, sendLoop 用于发送 watch 消息到流中，recvLoop 接受请求
         select 阻塞直到流关闭，或是超时退出。
         1.接收 watch 请求 recvLoop
         recvLoop 从 gRPCStream 读出 req, 然后分别处理类型为 CreateRequest, CancelRequest, ProgressRequest 的情况
         CreateRequest: 监听的可能是一个范围，所以构建 key 和 RangeEnd. 处理 StartRevision, 如果为 0, 那么使用当前 系统最新的 Rev+1. 调用 mvcc 层的 watchStream.Watch, 返回一个 watchid, 将这个 id封装到watchResponse,再将watchResponse 写到 ctrlStream
         CancelRequest: 还是调用 mvcc 层的 watchableStore.Cancel 取消订阅，然后清除状态信息
         ProgressRequest: broadcast 广播当前系统的 Rev 版本
         原文链接：https://blog.csdn.net/weixin_43916797/article/details/115591554`
        
  - 负载均衡
    - roundrobin 轮询调度算法
      - 从1到n的遍历服务器，然后重新开始循环。
    - weighted round robin 权重轮循
    - 动态轮询算法 动态负载均衡
    - 一致性哈希
    - 网络质量优先
    - 地理位置优先
  - 鉴权-身份认证： oauth
  - 安全
  - 日志
  - 监控
  
#### 其他 [rpcx](https://books.studygolang.com/go-rpc-programming-guide/part2/registry.html#inprocess) 分布式rpc框架

