## golang


### 常见问题

- 基础入门
  - 新手
    - [Golang开发新手常犯的50个错误](https://blog.csdn.net/gezhonglei2007/article/details/52237582)  
      1. 切片是指向同一片地址的  解决：cooy， 确定大小 dir1 := path[:sepIndex:sepIndex] //full slice expression  
      2. slice在添加元素前，与其它切片共享同一数据区域，修改会相互影响；但添加元素导致内存重新分配之后，不再指向原来的数据区域，修改元素，不再影响其它切片。
      
  - 数据类型
    - [连nil切片和空切片一不一样都不清楚？那BAT面试官只好让你回去等通知了。](https://mp.weixin.qq.com/s/myGJ4TrEoVGqLAN3tbZHMw)   
      make(0)切片指向的是同样的地址，append发生扩容后重新分配  
      var nil是0地址
      切片会扩容 注意函数传递是否有用
    - [golang面试题：字符串转成byte数组，会发生内存拷贝吗？](https://mp.weixin.qq.com/s/d80m0hgoKcHfKp4ZXH1M4A)  
      类型强转都会发生拷贝
      
    - [golang面试题：翻转含有中文、数字、英文字母的字符串](https://mp.weixin.qq.com/s/OIRPOszH-rTJp03AeRgnRQ)  
      rune
      
    - [golang面试题：拷贝大切片一定比小切片代价大吗？](https://mp.weixin.qq.com/s/hPYdiHYRufimyKT4FcW4HA)  
      浅拷贝一样  
    - struct 可以比较 在所有变量都是简单类型的时候
    - map不初始化使用会怎么样  
     Assignment to entry may panic because of 'nil' map 
    - map不初始化长度和初始化长度的区别  
      make：(t Type, size ...IntegerType)
        - make(map) 为空map分配足够的空间存放指定数量的元素，可以忽略大小。
    - map承载多大，大了怎么办  
      自动扩容 
    - map的iterator是否安全？能不能一边delete一边遍历？  
      安全，动态的获取
    - range  
      for range创建了每个元素的副本，而不是直接返回每个元素的引用，如果使用该值变量的地址作为指向每个元素的指针
    - 字符串不能改，那转成数组能改吗，怎么改
      转成[]byte 或者切片拼  
    - 怎么判断一个数组是否已经排序
      一次遍历比值  
    - 普通map如何不用锁解决协程安全问题
      通过channel顺序执行
    - array和slice的区别
      - array 定长，拷贝值
      - slice 可扩展 := 引用，但是扩容时会换地址
    - [golang面试题：json包变量不加tag会怎么样？](https://mp.weixin.qq.com/s/zZM_iLuopyenI0LD6VYZGw) 
    - 零切片、空切片、nil切片是什么
      1. 空切片需要分配  地址指向0
      2. nil切片 var定义的 nil切片 无地址
      3. 零切片 length是0 有空间 有地址
    - slice深拷贝和浅拷贝
      需要注意slice本身是浅拷贝
    - slice 扩容
      每逢 2^n 扩容 换地址
    - map触发扩容的时机，满足什么条件时扩容？
      map的数据量count大于(2^B)*6.5 类似
    - map扩容策略是什么
    - 自定义类型切片转字节切片和字节切片转回自动以类型切片
    - make和new什么区别  
      都是内存的分配（堆）
      - make只用于slice map channel的初始化
      - new用于类型的内存分配，并且内存值为0，返回的 指向类型的指针 返回的是引用类型本身
      - slice ，map，chanel创建的时候的几个参数什么含义  
        make：(t Type, size ...IntegerType)
        - make([]int, 0,10)分配一个底层数组  
        大小为10，并返回长度为0，容量为10的切片
        由底层数组支持。  
        - make(map) 为空map分配足够的空间存放指定数量的元素，可以忽略大小。
        - chan 是否有缓冲
    - 线程安全的map怎么实现
      	c.mu.Lock()
      	c.Map[key] = value
      	c.mu.Unlock()
   
  - 流程控制
    - [昨天那个在for循环里append元素的同事，今天还在么？](https://mp.weixin.qq.com/s/SHxcspmiKyPwPBbhfVxsGA) 
    - [golang面试官：for select时，如果通道已经关闭会怎么样？如果只有一个case呢？](https://mp.weixin.qq.com/s/lK6I353Iw08robqpmPB6-g) 
      - channel
        1. 不带缓存的channel
        2. 串联channel 多channel->
        3. 单方向channel  预定义类型 chan-<  int 只发送int的channel
        4. 带缓存的channel ch = make(chan string, 3)
        5. 无缓存的channel只有在receiver准备好后send才被执行，也就是说receiver的goroutine 和sender的goroutine 必须成对出现
        6. 如果无缓存，多余的写入的goroutine会泄漏，所以要使用有缓冲的chan
      - wg sync.WaitGroup  
        控制等待全部完成 wg.add(x) defer wg.Done() wg.Wait()
      - select 多路复用  
        多个case都满足是随机的，尽量保证平均  
        default为空的话 不停轮训，非阻塞。  
        channel的零值是nil  
        可以通过nil来激活/禁用case  
        channel 发送没接收会一直阻塞  
        带缓存的 缓存内不会阻塞
        要注意检查channel是否被消费，如果return 是个bug，泄漏go routine 不会被回收 
        goroutine 退出： 1. for range 退出 2. _,ok 然后赋值nil来退出。
      - 通道close 不判断ok，会死循环。
- 进阶
  - 包管理  （暂时不用看）
    [学go mod就够了！](https://studygolang.com/articles/27293)
  - 优化
    - 堆栈  [堆栈](https://www.cnblogs.com/sanmubai/p/13516885.html)
      - 栈（操作系统）：由操作系统自动分配释放 ，存放函数的参数值，局部变量的值等。其操作方式类似于数据结构中的栈。
        - 分段栈  
          1. 当 Goroutine 调用的函数层级或者局部变量需要的越来越多时，运行时会调用 
          runtime.morestack#go1.2 和 runtime.newstack#go1.2 创建一个新的栈空间，
          这些栈空间虽然不连续，但是当前 Goroutine 的多个栈空间会以链表的形式串联起来，运行时会通过指针找到连续的栈片段：
          2. 扩容麻烦 一个链表下的所有分段栈。分段栈机制虽然能够按需为当前 Goroutine 分配内存并且及时减少内存的占用，但是它也存在两个比较大的问题：
             1.如果当前 Goroutine 的栈几乎充满，那么任意的函数调用都会触发栈的扩容，当函数返回后又会触发栈的收缩，如果在一个循环中调用函数，栈的分配和释放就会造成巨大的额外开销，这被称为热分裂问题（Hot split）；
             2.一旦 Goroutine 使用的内存越过了分段栈的扩缩容阈值，运行时就会触发栈的扩容和缩容，带来额外的工作量；
        - 连续栈  
          连续栈可以解决分段栈中存在的两个问题，其核心原理就是每当程序的栈空间不足时，初始化一片更大的栈空间并将原栈中的所有值都迁移到新的栈中，新的局部变量或者函数调用就有了充足的内存空间。使用连续栈机制时，栈空间不足导致的扩容会经历以下几个步骤：  
          1. 在内存空间中分配更大的栈内存空间；
          2. 将旧栈中的所有内容复制到新的栈中；
          3. 将指向旧栈对应变量的指针重新指向新栈；
          4. 销毁并回收旧栈的内存空间；
        - 函数调用的时候检查当前堆栈的容量，如果快用净就新申请一块内存，并把栈指针指过去。当函数返回后，再把栈指针修改回来即可。
        
      - 堆（操作系统）： 一般由程序员分配释放， 若程序员不释放，程序结束时可能由OS回收，分配方式倒是类似于链表。
      - Go 语言的逃逸分析遵循以下两个不变性
        1.指向栈对象的指针不能存在于堆中； 操作系统的局部变量等等
        2.指向栈对象的指针不能在栈对象回收后存活；
     - 栈溢出： 递归 循环引用
      
    - 内存逃逸  
      能引起变量逃逸到堆上的典型情况：  
      - 在方法内把局部变量指针返回 局部变量原本应该在栈中分配，在栈中回收。但是由于返回时被外部引用，因此其生命周期大于栈，则溢出。
      - 发送指针或带有指针的值到 channel 中。 在编译时，是没有办法知道哪个 goroutine 会在 channel 上接收数据。所以编译器没法知道变量什么时候才会被释放。
      - 在一个切片上存储指针或带指针的值。 一个典型的例子就是 []*string 。这会导致切片的内容逃逸。尽管其后面的数组可能是在栈上分配的，但其引用的值一定是在堆上。
      - slice 的背后数组被重新分配了，因为 append 时可能会超出其容量( cap )。 slice 初始化的地方在编译时是可以知道的，它最开始会在栈上分配。如果切片背后的存储要基于运行时的数据进行扩充，就会在堆上分配。
      - 在 interface 类型上调用方法。 在 interface 类型上调用方法都是动态调度的 —— 方法的真正实现只能在运行时知道。想像一个 io.Reader 类型的变量 r , 调用 r.Read(b) 会使得 r 的值和切片b 的背后存储都逃逸掉，所以会在堆上分配。

    - [golang面试题：怎么避免内存逃逸？](https://mp.weixin.qq.com/s/4QAxGEr9KxtZXyfSG8VoCQ) 
    - [golang面试题：简单聊聊内存逃逸？](https://mp.weixin.qq.com/s/4YYR1eYFIFsNOaTxL4Q-eQ) 
    - [给大家丢脸了，用了三年golang，我还是没答对这道内存泄漏题](https://mp.weixin.qq.com/s/T6XXaFFyyOJioD6dqDJpFg)
    - 内存碎片化问题
    - chan相关的goroutine泄露的问题
    - string相关的goroutine泄露的问题
    - [你一定会遇到的内存回收策略导致的疑似内存泄漏的问题](https://colobu.com/2019/08/28/go-memory-leak-i-dont-think-so/)
    - sync.Pool的适用场景
    - go1.13sync.Pool对比go1.12版本优化点
    
    
  - 并发编程
    - [golang面试题：对已经关闭的的chan进行读写，会怎么样？为什么？](https://mp.weixin.qq.com/s/qm-8pvHBVRmLQQ4_DHc1Tw)   
      继续读完，然后一直读0值 死循环
    - [golang面试题：对未初始化的的chan进行读写，会怎么样？为什么？](https://mp.weixin.qq.com/s/ixJu0wrGXsCcGzveCqnr6A)    
      阻塞 不能阻塞就失败了
      
    - sync.map 的优缺点和使用场景
    - sync.Map的优化点

  - 高级特性
    - [golang面试题：能说说uintptr和unsafe.Pointer的区别吗？](https://mp.weixin.qq.com/s/IkOwh9bh36vK6JgN7b3KjA)
    - [golang 面试题：reflect（反射包）如何获取字段 tag？为什么 json 包不能导出私有变量的 tag？](https://mp.weixin.qq.com/s/WK9StkC3Jfy-o1dUqlo7Dg)
    - 协程和线程的差别
    - 垃圾回收的过程是怎么样的？
    - 什么是写屏障、混合写屏障，如何实现？
    - 开源库里会有一些类似下面这种奇怪的用法：`var _ io.Writer = (*myWriter)(nil)`，是为什么？
    - [GMP模型](https://zhuanlan.zhihu.com/p/261590663)  
      p调度M运行G
    - [动图图解，GMP里为什么要有P](https://mp.weixin.qq.com/s?__biz=MzAwMDAxNjU4Mg==&mid=2247484769&idx=1&sn=4d813bdf0977b3415db8faf4645ee216)
    - 协程之间是怎么调度的
    - 利用golang特性，设计一个QPS为500的服务器
    - 为什么gc会让程序变慢
    - 开多个线程和开多个协程会有什么区别
    - 两个interface{} 能不能比较  
      能： 动态类型和值都一样
    - 必须要手动对齐内存的情况
    - [go栈扩容和栈缩容，连续栈的缺点](https://segmentfault.com/a/1190000019570427)
    - golang怎么做代码优化
    - [golang隐藏技能:怎么访问私有成员](https://www.jianshu.com/p/7b3638b47845)
    
    
  - 问题排查
    - [trace](https://mp.weixin.qq.com/s?__biz=MzA4ODg0NDkzOA==&mid=2247487157&idx=1&sn=cbf1c87efe98433e07a2e58ee6e9899e&source=41#wechat_redirect) 
    - [pprof](https://mp.weixin.qq.com/s/d0olIiZgZNyZsO-OZDiEoA) 
  
  - 源码阅读
    - [sync.map](https://qcrao.com/2020/05/06/dive-into-go-sync-map/)
    - net/http 
      - [i/o timeout ， 希望你不要踩到这个net/http包的坑](https://mp.weixin.qq.com/s/UBiZp2Bfs7z1_mJ-JnOT1Q) 
    - [mutex](https://mp.weixin.qq.com/s/MntwgIJ2ynOAdwnypWUjZw)
    - [channel](https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-channel/?hmsr=toutiao.io&utm_medium=toutiao.io&utm_source=toutiao.io)
    - [context](https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-context/)
    - [select实现原理](https://mp.weixin.qq.com/s/9-eLJqYZrOpNLoiTGpBrKA)
    - main函数背后的启动过程
    - [内存管理](https://mp.weixin.qq.com/s?__biz=Mzg2MDU1Mjc3MQ==&mid=2247489860&idx=1&sn=2d3fa235f6768ad5a0c820b6241b9e99&source=41#wechat_redirect)
    - [GC垃圾回收](https://segmentfault.com/a/1190000020086769)
    - [timer](https://pengrl.com/p/62835/)

 
 #### [gmp](https://mp.weixin.qq.com/mp/homepage?__biz=MjM5MDUwNTQwMQ==&hid=1&sn=e47afe02b972f5296e1e3073982cf1b6&scene=1&devicetype=iOS13.7&version=18000e2a&lang=zh_CN&nettype=WIFI&ascene=7&session_us=gh_78aad54fff72&fontScale=100&wx_header=1)
   ![GMP调度图](https://cynthia-oss.oss-cn-beijing.aliyuncs.com/1634481001824.png)

   - G: goroutine
     - 结构:   
       - stack 
       - m 当前与g 绑定的m
       - sched goroutine 的运行现场
       - param wakeup时传入的参数
       - waitsince g被阻塞后的金丝时间
       - waitreason 被阻塞的原因
       - schedlink 链表 指向全剧队列里的下一个g
       - 
   - M: 工作线程  
       记录内核线程使用的栈信息。在执行调度代码时需要使用  
       使用用户goroutine代码时，用g自己的栈，调度会发生栈的切换
       当m没有工作的时候，在休眠前，会自旋的找工作
       - 结构
         - g0 : *g 栈空间
         - tls [6]uintptr 实现m与内核线程的绑定
         - curg 指向当前运行的g
         - p 当前绑定的p
         - spinning 当前处于自旋状态，正从其他线程偷工作
         - bolocked m 阻塞
       - 找可运行的goroutine流程:  
         先从本地队列找，定期会从全局队列找，最后实在没办法，就去别的 P 偷.
       - M 只有自旋和非自旋两种状态： 
         自旋的时候，会努力找工作；找不到的时候会进入非自旋状态，之后会休眠，直到有工作需要处理时，被其他工作线程唤醒，又进入自旋状态。
         
   - P: processor  
       为 M 的执行提供“上下文”，保存 M 执行 G 时的一些资源，
       例如本地可运行 G 队列，memeory cache 等。  
       一个 M 只有绑定 P 才能执行 goroutine，当 M 被阻塞时，
       整个 P 会被传递给其他 M ，或者说整个 P 被接管。
   - Go scheduler schedt   
       保存调度器的状态信息、全局的可运行 G 队列等。
       schedt全局只有一份实体
       - 启动流程
         ```从磁盘上读取可执行文件，加载到内存
            
            创建进程和主线程
            
            为主线程分配栈空间
            
            把由用户在命令行输入的参数拷贝到主线程的栈
            
            把主线程放入操作系统的运行队列等待被调度
         ```
   - Main Goroutine  
     - 调用newproc 启动过1个goroutine 传入func和func大小
   
   - goroutine阻塞时会发生什么[正在执行的goruntine发生阻塞，golang调度策略](https://blog.csdn.net/csdniter/article/details/112175118)
     - 阻塞在网络IO:  
       则通过netPoller（epoll实现）轮询，放开m，m去执行其他队列中的g。
       异步调用完成后，则被移回p的lrq中。
       ![](https://cynthia-oss.oss-cn-beijing.aliyuncs.com/1634482900127.png)
     - 同步系统调用阻塞（比如文件I/O） 会导致M阻塞：   
       G1 将进行同步系统调用以阻塞 M1。
       ![](https://cynthia-oss.oss-cn-beijing.aliyuncs.com/1634482984559.png)  
       调度器介入后：识别出 G1 已导致 M1 阻塞，此时，调度器将 M1 与 P 分离，同时也将 G1 带走。然后调度器引入新的 M2 来服务 P。此时，可以从 LRQ 中选择 G2 并在 M2 上进行上下文切换。
       ![](https://cynthia-oss.oss-cn-beijing.aliyuncs.com/1634482996211.png)
       阻塞的系统调用完成后：G1 可以移回 LRQ 并再次由 P 执行。如果这种情况再次发生，M1 将被放在旁边以备将来重复使用。
       ![](https://cynthia-oss.oss-cn-beijing.aliyuncs.com/1634483096306.png)
     - 如果在 Goroutine 去执行一个 sleep 操作，导致 M 被阻塞了： 
       Go 程序后台有一个监控线程 sysmon，它监控那些长时间运行的 G 任务然后设置可以强占的标识符，别的 Goroutine 就可以抢先进来执行。  
       只要下次这个 Goroutine 进行函数调用，那么就会被强占，同时也会保护现场，然后重新放入 P 的本地队列里面等待下次执行。
   - goroutine生命周期（main）：
      ![生命周期](https://cynthia-oss.oss-cn-beijing.aliyuncs.com/1634484024274.png)

   - 总结： https://zhuanlan.zhihu.com/p/261590663
     ![调度示意图](https://cynthia-oss.oss-cn-beijing.aliyuncs.com/1634483925771.png)
     - 调度过程：
       1. 创建G 保存在P的本地队列/全剧队列。
       2. P 唤醒1个M
       3. P 按照他的队列的G的执行顺序继续执行任务
       4. M 如果空闲，则找P的本地队列的G，如果没有，则找全局队列，如果还没有，就去别的P偷G。
       5. M 执行一个调度循环：调用G对象->执行->清理线程->继续寻找Goroutine。
       6. M 的栈保存在G对象。 M随时可能发生上下文切换。
       7. 如果G对象还没有被执行，M可以将G重新放到P的调度队列，等待下一次的调度执行。当调度执行时，M可以通过G的vdsoSP, vdsoPC 寄存器进行现场恢复。  
       8. 清理线程： G的调度是为了实现P/M的绑定，所以线程清理就是释放P上的G，让其他的G能够被调度。
 
 #### GC
- [gc](https://segmentfault.com/a/1190000020086769)
   - 常见的方式
     - 引用计数
     - 标记清除
   - stw
     找可清除的指针/对象的时候，stw（stop the word） 找指针是个很短的瞬间。
   - 三色标记
     - 第一阶段 gc开始 stop the world  
       初始状态所有内存是白色，进入标记队列是灰色
       1. stop the world
       2. 每个processor启动一个mark worker goroutine用于标记（用于第二阶段工作）
       3. 启动gc write barrier（记录一下后续在进行marking时被修改的指针）
       4. 找到所有roots（stack, heap, global vars）并加入标记队列
       5. start the world，进入第二阶段
     - 第二阶段 marking，start the world
       ![](https://cynthia-oss.oss-cn-beijing.aliyuncs.com/1634022937392.png)

     - GC触发条件内存大小阈值， 内存达到上次gc后的2倍达到定时时间 ，2m interval阈值是由一个gc percent的变量控制的,当新分配的内存占已在使用中的内存的比例超过gcprecent时就会触发。比如一次回收完毕后，内存的使用量为5M，那么下次回收的时机则是内存分配达到10M的时候。也就是说，并不是内存分配越多，垃圾回收频率越高。 如果一直达不到内存大小的阈值呢？这个时候GC就会被定时时间触发，比如一直达不到10M，那就定时（默认2min触发一次）触发一次GC保证资源的回收。
       链接：https://zhuanlan.zhihu.com/p/92210761
     - 写屏障： 会标记原值。
     ![完整的gc流程](https://cynthia-oss.oss-cn-beijing.aliyuncs.com/1634486274555.png)

 #### gin 实现[gin源码剖析](https://studygolang.com/articles/28836?fr=sidebar)
 [gin为什么用tree](https://blog.csdn.net/weixin_42265507/article/details/109083685)
 #### grpc [protobuf](https://www.jianshu.com/p/419efe983cb2)
 
### 资料
[切片](https://blog.csdn.net/weixin_42266173/article/details/81749949)
[快速入门汇总](https://mp.weixin.qq.com/mp/homepage?__biz=MjM5MDUwNTQwMQ==&hid=1&sn=e47afe02b972f5296e1e3073982cf1b6&scene=1&devicetype=iOS13.7&version=18000e2a&lang=zh_CN&nettype=WIFI&ascene=7&session_us=gh_78aad54fff72&fontScale=100&wx_header=1)

问题：创建一个len=0，空间为10的切片？


