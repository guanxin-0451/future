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
      浅拷贝 一样  
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
      一个channel同时仅允许被一个goroutine读写，为简单起见，本章后续部分说明读写过程时不再涉及加锁和解锁。
    - 遍历map是无序的
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
    - sync.Pool的适用场景v
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
    - [go优先调度](https://www.cnblogs.com/dearplain/p/8276138.html)
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
       1. 创建G 保存在P的本地队列/全局队列。
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
     ```
        在开始gc是启动内存写屏障，所有修改对象的操作加入一个写屏障，通知gc将受影响的对象变为灰色；
        然后在第一次gc后，重新对gc过程中变灰的对象再扫描一次
     ```
     - 三色不变式  
       强三色不变式：黑色节点不允许引用白色节点，破坏了条件一。
       弱三色不变式：黑色节点允许引用白色节点，但是该白色节点必须有其他灰色节点的引用或间接引用，破坏了条件二。
     - 混合写屏障   
       栈区：
       栈上对象全部扫描标记为黑色（每个栈单独扫描，无需 STW 整个程序，停止单个扫描栈即可）；
       
       GC 期间，任何在栈上创建的新对象，均为黑色（不用再对栈重新扫描）；
       
       堆区：
       
       被删除的对象标记为灰色（删除写屏障）；
       
       被添加的对象标记为灰色（插入写屏障）。
       
   - 触发gc
     - 主动触发
     - 系统监控 超2分钟 自动触发
     - 步调算法控制的，预估下一次需要出发gc时，堆的大小。
   ![完整的gc流程](https://cynthia-oss.oss-cn-beijing.aliyuncs.com/1634486274555.png)

 #### gin 实现[gin源码剖析](https://studygolang.com/articles/28836?fr=sidebar)
 - 路由：前缀树
  - [gin为什么用tree](https://blog.csdn.net/weixin_42265507/article/details/109083685)
  - map 不能用动态参数
  - priority: 子孙路由总数， proority大的节点放在前面
 - engine
  - tree：路由
  - pool context池 减少gc
  - RouterGroup 提供设置
 
 #### grpc [protobuf](https://www.jianshu.com/p/419efe983cb2)

### go程序面试宝典
#### 逃逸分析
  - 垃圾回收
  - 逃逸分析把变量合理地分配到堆/栈上
    - 堆适合不可预知大小的内存分配，分配速度较慢，会形成内存碎片。
    - 栈通过push分配，并且会自动释放。
    - 通过逃逸分析，尽可能把不需要分配到堆上的变量分配到栈上，堆的变量少了，减少堆内存分配的开销，减少gc的压力。
  - 逃逸分析（逃逸到堆上）怎么完成的
    - 如果一个函数返回对一个变量的引用，这个变量就会发生逃逸。
    - go中的变量只有在编译器可以证明在函数返回后不会再被饮用的，才分配到栈上，否则则分配到堆上。
    - 编译器通过分析代码，决定分配变量。
      - 如果变量在函数外部没有被引用，则优先分配到栈上。
        - 需要的内存过大，超过了栈的存储能力，则分配到堆上。
      - 如果变量在函数外部存在引用，则必定放到堆上。
      - 函数参数为interface，没办法确认具体类型，也会逃逸到堆上。
  - go build gcflags '-m -l' main.go
    
  - golang的堆栈和c++的堆栈
    - c++ 堆栈： 操作系统层级的概念
    - go ： 操作系统分配的栈，被go本身消耗，用于调度器，垃圾回收，系统调用等。  
      用堆空间构造逻辑上的堆栈。
      - 栈可能进行深拷贝，将其整个复制到另一块空间上。
      - **指针的算数运算不能奏效**。
#### defer 延迟语句
  - 让函数/语句在当前函数执行完成后执行
  - defer在panic下也能正常的执行 start + 语句 + close 模式  如果出现panic 那就不能正常close      
  - defer压栈 后进先出 ： 因为后边的代码片可能依赖于前面的片，先结束前面的，会导致后面的失败。
  - defer后的函数在执行的时候，函数调用的参数会被保存起来，复制一份。如果是值，那就一致，如果是引用，值会变。
  - defer 函数定义时，对外部变量的引用有两种方式： 函数参数，闭包引用。
    - 函数参数会保持当时的现场。
    - **闭包defer func()和引用，都会是最新的值。**
  - return 后defer不能被注册。
  - return xxx := 先执行赋值 result = xxx, 再调用defer func() , return ;
  - **闭包是什么？**
    - 闭包=函数+引用环境
  - recover 要在defer调用的函数内写。 起到稳住主流程的作用
  
  
  - 为什么无法从父goroutine恢复子goroutine的panic？
    - goroutine被设计成一个独立的代码执行单元，有自己的执行栈，不和其他的goroutine共享任何数据。无法让goroutine有返回值。
    - 如果需要有全局的异常捕获，通过channel
    
#### 数据容器
##### 数组&切片
  - 数组 定长 l := [length]int{}
  - make 用于slice,map,chan, 
  - new 用于int， 数组，结构体等，返回指针。

##### map
  - sync.map 线程安全
  - 当map的key，value都不是指针，且size都小于128字节的情况下，会把bmap标记为不含指针。避免gc时扫描整个hmap。
  - map B := buckets的对数.  扩容时，buckets翻倍，按照2^B取hash值的后2^B位，进行rehash
  - bucket ：桶，每个桶存"同类的"key = 低B位相同的key，桶的每个槽记下高八位的值。 如果超过8个落入bucket， 则构建一个新的，放在overflow里。
  
  - key为什么是无序的： 扩容 
  
#### channel 通道
##### CSP
  - 控制并发数  
    ```
    var token = make(chan int, 3)
    
    for _, w = range work {
        go func() {
            defer func(){
                <- token
            }
            token <- 1
            w()
        }   
    }
    ```
    
  - 如何关闭channel？
    - 1 一个sender： sender处关闭即可。
    - 2 多个sender，一个receiver  receiver处发送一个关闭信号的channel发送给所有sender，senders监听到数据后，停止发送数据，不关闭channel，等待gc回收。
    - 3 多个sender，多个receiver 加中间人，接收所有关闭信号，再决定关闭。
  - 泄漏：goroutine操作channel后，处于发送或接受阻塞状态，channel处于满或空的状态，垃圾回收也不会回收此类资源。
  - panic 
    - 读写，关闭 nil channel写
#### 接口
  - go的itab中的fun字段是运行期间动态生成的。
  - duck typing 不是由继承自特定的类/接口决定，而是由它当前的方法和属性的集合决定的。
  - 如何用interface实现多态。
    - 一种类型具有多种类型的能力。
    - 允许不同的对象对同一消息做出灵活的反应。
    - 以一种通用的方式对待使用的对象。
    - 非动态语言必须通过继承和接口的方式来实现。
  - 接口转换的原理
    - 判断类型是否满足某个接口时： go将类型的方法集和接口所需要的方法集进行匹配。 先对方法字典排序，再比较 o(m+n)
  - 断言： 安全断言 xx, ok:= xxx.(类型)
  - 类实现String()方法，可以按String的返回进行fmt
  - 类型检查 var _ io.Writer = (*myWriter)(nil) 检查是否实现了接口
#### context 
  - context 用于解决goroutine之间退出通知，元数据传递的功能的问题。
  
#### error
  - 视error为值： 实现error接口的类型都是error
  - 检查并优雅的处理错误。
    - wrap
    - cause
  - 只处理错误一次  

#### 计时器
  - timer 
    - 四叉堆 o(log4N）
    - go 1.14 为每个p维护一个小顶堆
    
#### 反射
  - interface实现反射
  - reflect reflect.Type reflect.Value
  - 反射变量不能代表真实变量时 不可以直接操作
  - deepEqual 比较两个对象是否完全相同
  - 反射实现深拷贝
  
#### sync包
  - sync.WaitGroup 
    - 并发goroutine的控制 等待整体跑完。
    - 保证计数器不能为负值
    - 保证 Add() 方法全部调用完成之后再调用 Wait()
    - waitgroup 可以重复使用
    - atomic 原子操作代替锁, 提高并发性
    - 合并两个 int32 为一个 int64 提高读取存入数据性能
    - 对于不希望被复制的结构体， 可以使用 noCopy 字段
    - 通过信号量计数通知所有正在等待的goroutine。
  - sync.Pool 对象池
    - 保存临时取还的对象的池
    - 缓存对象，减少对gc的消耗
    - 协程安全: 
    - 多个goroutine使用同一个对象，避免频繁创建，销毁
    - gin 使用sync.pool
    - 实现：多个p对应多个存储切片，先从p对应的数组内拿（减少竞争），拿不到取其他的p拿，在拿不到，从victim的private拿 gc过程中，存储会转到victim上。
  - sync.map
    - 协程安全
    - 读多写少
    - 协程安全的map，搞了2个内部的map，1个read，1个dirty，dirty好像还是加锁了。read是单线程的，能保证原子性。

#### 调度机制
  - 可以理解为是用户级"线程"
    - 更小的栈内存消耗
    - 线程创建和销毁的消耗大，协程go runtime负责管理线程的创建销毁，消耗小。
  - main 怎么起来的
  ```
    一.go进程怎么让操作系统调度起来
     1.从磁盘读可执行文件（go build的那玩意），加载到内存。
     2.创建进程和主线程。
     3.给主线程分配栈空间（golang的程序执行本身要用栈空间，不给用户代码用）
     4.命令行参数加载到栈中。
     5.主线程放入系统的运行队列等待操作系统调度。
     
     二.go进程初始化（gmp）
     1.起个g0协程 万物开始的协程
     2.起个m0， 主线程绑定m0
     3.初始化所有的p，一般几个线程就几个p。
     三.main（）起来了
     1.创建个g
     2.放g队列里
     3.等着调度
   ```
#### 内存分配机制
  - 从操作系统申请，由页分配器负责
  - 为用户程序分配内存，对象分配器负责
    - 分配策略：
      - 顺序分配  
        用顺序表： 每个goroutine直接获得一页（逻辑》物理）内存。
      - 自由表分配（链表）
        自身对象所在的非托管内存（go自身所需的堆内存）
  
  - 回收用户程序所分配的内存，gc负责
  - 向操作系统归还深情的内存，拾荒器负责 scavenger
  - 对象分配
    - 大对象分配 直接页分配
    - 小对象分配<32kb 内存对齐 页分配
    - 微对象分配 < 16b 多个微对象合并 16b大小的块直接进行管理和释放。
  - 拾荒器：
    - 定时归还：主动 大量瞬时内存增长场景，无法解决。
    - 申请时直接占用： 被动策略。
    - 按照一定比例，以驻留内存比实际消耗内存高10%为目标，循序渐进的主动归还。

#### 垃圾回收
  - 根对象（根集合）
    - 垃圾回收器在标记过程中最先检查的对象
    - 包括：
      - 执行栈
      - 全局变量：编译期就能确定存在于整个生命周期的变量。
    - 常见的回收方式：
      - 追踪（tracing）  
        - 从根对象出发，根据对象间的引用信息，一步步推进，直到扫描完毕整个堆。  
        - go, java，js  
        - 标记清除：确定存活的对象标记，清扫可以回收的对象
        - 标记整理：解决内存碎片，尽量将对象整理到一块连续的内存中   
        - 增量式： 标记和清扫分批进行，近似实时。
        - 增量整理：
        - 分代式：根据存活时间的长短进行分类，进行分代假设。尽量从短时间存活的开始清除。
        - 三色清除：无分代， 不整理， 和用户代码并发执行。
          - go基于tcmalloc的现代内存分配算法，基本没有碎片问题。
          - go编译器 逃逸分析 大部分新生对象存储在栈上，可以直接回收。 不需要分代回收
        
      - 引用计数 reference counting
        python 
      
    - 内存泄漏
      - 根对象引用， 全局对象可能附着了某个变量，忽略了将其释放
      - goroutine 泄漏   不断产出新的goroutine 且不关闭， 内存不会释放。 比如go func(){select {}}
      - chan 一个goroutine 尝试向一个没有接收方的无缓冲的channel发送消息， 则该goroutine 永久的休眠，goroutine及执行栈都不会被释放。
      

 
### 资料
[切片](https://blog.csdn.net/weixin_42266173/article/details/81749949)
[快速入门汇总](https://mp.weixin.qq.com/mp/homepage?__biz=MjM5MDUwNTQwMQ==&hid=1&sn=e47afe02b972f5296e1e3073982cf1b6&scene=1&devicetype=iOS13.7&version=18000e2a&lang=zh_CN&nettype=WIFI&ascene=7&session_us=gh_78aad54fff72&fontScale=100&wx_header=1)
[go doc](https://golang.google.cn/doc/)
[go release history](https://golang.google.cn/doc/devel/release)
问题：创建一个len=0，空间为10的切片？


## 面试问题
### golang 最新版本做了些什么？历史版本？
Go1.14 - 2020 年 2 月：
现在 Go Module 已经可以用于生产环境，鼓励所有用户迁移到 Module。该版本支持嵌入具有重叠方法集的接口。性能方面做了较大的改进，包括：进一步提升 defer 性能、页分配器更高效，同时 timer 也更高效。
现在，Goroutine 支持异步抢占。
在 runtime.sighandler 函数中注册了 SIGURG 信号的处理函数 runtime.doSigPreempt，在触发垃圾回收的栈扫描时，调用函数挂起goroutine，并向M发送信号，M收到信号后，会让当前goroutine陷入休眠继续执行其他的goroutine。


[doc](https://golang.google.cn/doc/devel/release)

