#### 

- [锁](https://aobing.blog.csdn.net/article/details/104691668)

- sync、fsync、fdatasync

    - sync系统调用：将所有修改过的缓冲区排入写队列，然后就返回了，它并不等实际的写磁盘的操作结束。所以它的返回并不能保证数据的安全性。通常会有一个update系统守护进程每隔30s调用一次sync。命令sync(1)也是调用sync函数。

    - fsync系统调用：需要你在入参的位置上传递给他一个fd，然后系统调用就会对这个fd指向的文件起作用。fsync会确保一直到写磁盘操作结束才会返回，所以当你的程序使用这个函数并且它成功返回时，就说明数据肯定已经安全的落盘了。所以fsync适合数据库这种程序。

    - fdatasync系统调用：和fsync类似但是它只会影响文件的一部分，因为除了文件中的数据之外，fsync还会同步文件的属性。
#### 内存
Tcmalloc概述

tcmalloc将内存请求分为两类，大对象请求和小对象请求，大对象为>=32K的对象。

tcmalloc会为每个线程分配本地缓存，小对象请求可以直接从本地缓存获取，如果没有空闲内存，则从central heap中一次性获取一连串小对象。

tcmalloc对于小内存，按8的整数次倍分配，对于大内存，按4K的整数次倍分配。

当某个线程缓存中所有对象的总大小超过2MB的时候，会进行垃圾收集。垃圾收集阈值会自动根据线程数量的增加而减少，这样就不会因为程序有大量线程而过度浪费内存。



Tcmalloc优势

1 速度快

ptmalloc在一台2.8GHz的P4机器上（对于小对象）执行一次malloc及free大约需要300纳秒。而TCMalloc的版本同样的操作大约只需要50纳秒。

2 降低了锁争用

TCMalloc减少了多线程程序中的锁争用情况。对于小对象，几乎已经达到了零争用。对于大对象，TCMalloc尝试使用粒度较好和有效的自旋锁。

3 节省内存

分配N个8字节对象使用大约8N * 1.01字节的空间。而ptmalloc2中每个对象都使用了一个四字节的头。

尽管ptmalloc2也实现了per-thread缓存，，第一个线程申请的内存，释放的时候还是必须放到第一个线程池中（不可移动），这样可能导致大量内存浪费。

ptmalloc2 also reduces lock contention by using per-thread arenas but there is a big problem with ptmalloc2's use of per-thread arenas. In ptmalloc2 memory can never move from one arena to another. This can lead to huge amounts of wasted space.

tcmalloc跨线程归还内存，是因为所有线程公用了底层的一个分配器?，所以跨线程归还是无需加锁的。
#### 线程 协程 进程
- 线程定义  
  线程是进程的基本执行单元，一个进程的所有任务都在线程中执行  
  进程要想执行任务，必须得有线程，进程至少要有一条线程  
  程序启动会默认开启一条线程，这条线程被称为主线程或 UI 线程

- 进程定义  
 进程是指在系统中正在运行的一个应用程序  
 每个进程之间是独立的，每个进程均运行在其专用的且受保护的内存
- 进程与线程的区别
  - 进程
    - 概念：运行中的程序
    - 状态： 运行 阻塞 就绪 挂起
    - 控制结构： pcb 进程控制块: 进程标识符. cpu相关信息， 资源分配清单
      - pcb组成队列： 阻塞队列，就绪队列。
    - 通信
      - 管道
      - fifo管道
      - 消息队列：内核中
        - 每个消息体有一个最大长度的限制，并且队列所包含消息体的总长度也是有上限的，这是其中一个不足之处。
          另一个缺点是消息队列通信过程中存在用户态和内核态之间的数据拷贝问题。进程往消息队列写入数据时，会发送用户态拷贝数据到内核态的过程，同理读取数据时会发生从内核态到用户态拷贝数据的过程。
      - 信号量
      - 共享内存
    - 上下文切换
    `CPU 上下文切换就是先把前一个任务的 CPU 上下文（CPU 寄存器和程序计数器）保存起来，然后加载新任务的上下文到这些寄存器和程序计数器，
    最后再跳转到程序计数器所指的新位置，运行新任务。
    为了保证所有进程可以得到公平调度，CPU 时间被划分为一段段的时间片，这些时间片再被轮流分配给各个进程。这样，当某个进程的时间片耗尽了，就会被系统挂起，切换到其它正在等待 CPU 的进程运行； 
    进程在系统资源不足（比如内存不足）时，要等到资源满足后才可以运行，这个时候进程也会被挂起，并由系统调度其他进程运行；
    当进程通过睡眠函数 sleep 这样的方法将自己主动挂起时，自然也会重新调度；
    当有优先级更高的进程运行时，为了保证高优先级进程的运行，当前进程会被挂起，由高优先级进程来运行；
    发生硬件中断时，CPU 上的进程会被中断挂起，转而执行内核中的中断服务程序；
    `
  - 线程
  - 同一个进程内多个线程之间可以共享代码段、数据段、打开的文件等资源，但每个线程都有独立一套的寄存器和栈，这样可以确保线程的控制流是相对独立的。
  - 地址空间：同一进程的线程共享本进程的地址空间，而进程之间则是独立的地址空间。  
  - 资源拥有：同一进程内的线程共享本进程的资源（如内存、I/O、cpu等），但是进程之间的资源是独立的。  
  - 一个进程崩溃后，在保护模式下不会对其他进程产生影响，但是一个线程崩溃整个进程都死掉。所以多进程要比多线程健壮。  
  进程切换时，消耗的资源大，效率高。所以涉及到频繁的切换时，使用线程要好于进程。同样如果要求同时进行并且又要共享某些变量的并发操作，只能用线程不能用进程
  执行过程：每个独立的进程有一个程序运行的入口、顺序执行序列和程序入口。但是线程不能独立执行，必须依存在应用程序中，由应用程序提供多个线程执行控制。  
  - 线程是处理器调度的基本单位，但是进程不是。

协程是用户态的线程下的资源。

- 进程间通信
```
一、管道

管道，通常指无名管道，是 UNIX 系统IPC最古老的形式。

特点：

它是半双工的（即数据只能在一个方向上流动），具有固定的读端和写端。

它只能用于具有亲缘关系的进程之间的通信（也是父子进程或者兄弟进程之间）。

它可以看成是一种特殊的文件，对于它的读写也可以使用普通的read、write 等函数。但是它不是普通的文件，并不属于其他任何文件系统，并且只存在于内存中。

二、FIFO

FIFO，也称为命名管道，它是一种文件类型。

1、特点

FIFO可以在无关的进程之间交换数据，与无名管道不同。

FIFO有路径名与之相关联，它以一种特殊设备文件形式存在于文件系统中。

三、消息队列

消息队列，是消息的链接表，存放在内核中。一个消息队列由一个标识符（即队列ID）来标识。

特点

消息队列是面向记录的，其中的消息具有特定的格式以及特定的优先级。

消息队列独立于发送与接收进程。进程终止时，消息队列及其内容并不会被删除。

消息队列可以实现消息的随机查询,消息不一定要以先进先出的次序读取,也可以按消息的类型读取。

四、信号量

信号量（semaphore）与已经介绍过的 IPC 结构不同，它是一个计数器。信号量用于实现进程间的互斥与同步，而不是用于存储进程间通信数据。

特点

信号量用于进程间同步，若要在进程间传递数据需要结合共享内存。

信号量基于操作系统的 PV 操作，程序对信号量的操作都是原子操作。

每次对信号量的 PV 操作不仅限于对信号量值加 1 或减 1，而且可以加减任意正整数。

支持信号量组。

五、共享内存

共享内存（Shared Memory），指两个或多个进程共享一个给定的存储区。

特点

共享内存是最快的一种 IPC，因为进程是直接对内存进行存取。

因为多个进程可以同时操作，所以需要进行同步。

信号量+共享内存通常结合在一起使用，信号量用来同步对共享内存的访问。
```


单线程，多线程
```https://cloud.tencent.com/developer/article/1120615
   
   1）以前一直有个误区，以为：高性能服务器 一定是多线程来实现的
   
   原因很简单因为误区二导致的：多线程 一定比 单线程 效率高，其实不然！
   
   在说这个事前希望大家都能对 CPU 、 内存 、 硬盘的速度都有了解了，这样可能理解得更深刻一点，不了解的朋友点：CPU到底比内存跟硬盘快多少
   
   2）redis 核心就是 如果我的数据全都在内存里，我单线程的去操作 就是效率最高的，为什么呢，因为多线程的本质就是 CPU 模拟出来多个线程的情况，这种模拟出来的情况就有一个代价，就是上下文的切换，对于一个内存的系统来说，它没有上下文的切换就是效率最高的。redis 用 单个CPU 绑定一块内存的数据，然后针对这块内存的数据进行多次读写的时候，都是在一个CPU上完成的，所以它是单线程处理这个事。在内存的情况下，这个方案就是最佳方案 —— 阿里 沈询
   
   因为一次CPU上下文的切换大概在 1500ns 左右。
   
   从内存中读取 1MB 的连续数据，耗时大约为 250us，假设1MB的数据由多个线程读取了1000次，那么就有1000次时间上下文的切换，
   
   那么就有1500ns * 1000 = 1500us ，我单线程的读完1MB数据才250us ,你光时间上下文的切换就用了1500us了，我还不算你每次读一点数据 的时间，
   
   3）那什么时候用多线程的方案呢？
   
   【IOPS（Input/Output Operations Per Second）是一个用于计算机存储设备（如硬盘（HDD）、固态硬盘（SSD）或存储区域网络（SAN））性能测试的量测方式】
   
   【吞吐量是指对网络、设备、端口、虚电路或其他设施，单位时间内成功地传送数据的数量（以比特、字节、分组等测量）】
   
   答案是：下层的存储等慢速的情况。比如磁盘
   
   内存是一个 IOPS 非常高的系统，因为我想申请一块内存就申请一块内存，销毁一块内存我就销毁一块内存，内存的申请和销毁是很容易的。而且内存是可以动态的申请大小的。
   
   磁盘的特性是：IPOS很低很低，但吞吐量很高。这就意味着，大量的读写操作都必须攒到一起，再提交到磁盘的时候，性能最高。为什么呢？
   
   如果我有一个事务组的操作（就是几个已经分开了的事务请求，比如写读写读写，这么五个操作在一起），在内存中，因为IOPS非常高，我可以一个一个的完成，但是如果在磁盘中也有这种请求方式的话，
   
   我第一个写操作是这样完成的：我先在硬盘中寻址，大概花费10ms，然后我读一个数据可能花费1ms然后我再运算（忽略不计），再写回硬盘又是10ms ，总共21ms
   
   第二个操作去读花了10ms, 第三个又是写花费了21ms ,然后我再读10ms, 写21ms ，五个请求总共花费83ms，这还是最理想的情况下，这如果在内存中，大概1ms不到。
   
   所以对于磁盘来说，它吞吐量这么大，那最好的方案肯定是我将N个请求一起放在一个buff里，然后一起去提交。
   
   方法就是用异步：将请求和处理的线程不绑定，请求的线程将请求放在一个buff里，然后等buff快满了，处理的线程再去处理这个buff。然后由这个buff 统一的去写入磁盘，或者读磁盘，这样效率就是最高。java里的 IO不就是这么干的么~
   
   对于慢速设备，这种处理方式就是最佳的，慢速设备有磁盘，网络 ，SSD 等等，
   
   多线程 ，异步的方式处理这些问题非常常见，大名鼎鼎的netty 就是这么干的。
   
   终于把 redis 为什么是单线程说清楚了，把什么时候用单线程跟多线程也说清楚了，其实也是些很简单的东西，只是基础不好的时候，就真的尴尬。。。。
   
   4）补一发大师语录：来说说，为何单核cpu绑定一块内存效率最高
   
   “我们不能任由操作系统负载均衡，因为我们自己更了解自己的程序，所以我们可以手动地为其分配CPU核，而不会过多地占用CPU”，默认情况下单线程在进行系统调用的时候会随机使用CPU内核，为了优化Redis，我们可以使用工具为单线程绑定固定的CPU内核，减少不必要的性能损耗！
   
   redis作为单进程模型的程序，为了充分利用多核CPU，常常在一台server上会启动多个实例。而为了减少切换的开销，有必要为每个实例指定其所运行的CPU。
   
   Linux 上 taskset 可以将某个进程绑定到一个特定的CPU。你比操作系统更了解自己的程序，为了避免调度器愚蠢的调度你的程序，或是为了在多线程程序中避免缓存失效造成的开销。
   
   顺便再提一句：redis 的瓶颈在网络上 。。。。
```

#### 多路复用
  - 5种io模型
    - 阻塞 IO 模型：硬件到系统内核，阻塞。系统内核到程序空间，阻塞。
    - 非阻塞 IO 模型：硬件到系统内核，轮询阻塞。系统内核到程序空间，阻塞。
    - 复用 IO 模型：硬件到系统内核，多流轮询阻塞。系统内核到程序空间，阻塞。
    - 信号驱动 IO 模型：硬件到系统内核，信号回调不阻塞。系统内核到程序空间，阻塞。
    - 异步 IO 模型：硬件到系统内核，信号回调不阻塞。系统内核到程序空间，信号回调不阻塞。
  - 一 基本概念
    - 用户空间和内核空间
      ```现在操作系统都是采用虚拟存储器，那么对32位操作系统而言，它的寻址空间（虚拟存储空间）为4G（2的32次方）。操作系统的核心是内核，
      独立于普通的应用程序，可以访问受保护的内存空间，也有访问底层硬件设备的所有权限。为了保证用户进程不能直接操作内核（kernel），
      保证内核的安全，操心系统将虚拟空间划分为两部分，一部分为内核空间，一部分为用户空间。针对linux操作系统而言，
      将最高的1G字节（从虚拟地址0xC0000000到0xFFFFFFFF），供内核使用，称为内核空间，而将较低的3G字节（从虚拟地址0x00000000到0xBFFFFFFF），
      供各个进程使用，称为用户空间。
      ```
    - 进程切换
      为了控制进程的执行，内核必须有能力挂起正在CPU上运行的进程，并恢复以前挂起的某个进程的执行。
      这种行为被称为进程切换。因此可以说，任何进程都是在操作系统内核的支持下运行的，是与内核紧密相关的。  
      切换进程消耗资源  
      从一个进程的运行转到另一个进程上运行，这个过程中经过下面这些变化：  
      1. 保存处理机上下文，包括程序计数器和其他寄存器。
      2. 更新PCB信息。
      3. 把进程的PCB移入相应的队列，如就绪、在某事件阻塞等队列。
      4. 选择另一个进程执行，并更新其PCB。
      5. 更新内存管理的数据结构。
      6. 恢复处理机上下文。
    - 进程阻塞
      正在执行的进程，由于期待的某些事件未发生，如请求系统资源失败、等待某种操作的完成、新数据尚未到达或无新工作做等，
      则由系统自动执行阻塞原语(Block)，使自己由运行状态变为阻塞状态。可见，进程的阻塞是进程自身的一种主动行为，
      也因此只有处于运行态的进程（获得CPU），才可能将其转为阻塞状态。当进程进入阻塞状态，是不占用CPU资源的。
    - 文件描述符
      指向内核为每一个进程所维护的该进程打开文件的记录表。当程序打开一个现有文件或者创建一个新文件时，内核向进程返回一个文件描述符。
    - 缓存I/O
      缓存 I/O 又被称作标准 I/O，大多数文件系统的默认 I/O 操作都是缓存 I/O。在 Linux 的缓存 I/O 机制中，
      操作系统会将 I/O 的数据缓存在文件系统的页缓存（ page cache ）中，也就是说，数据会先被拷贝到操作系统内核的缓冲区中，
      然后才会从操作系统内核的缓冲区拷贝到应用程序的地址空间。  
      Sendfile 提供了一种减少用户空间转到内核空间的操作
  - 二 IO 模式  
    当一个read操作发生时，它会经历两个阶段：  
    1. 等待数据准备 (Waiting for the data to be ready)
    2. 将数据从内核拷贝到进程中 (Copying the data from the kernel to the process)  
    正式因为这两个阶段，linux系统产生了下面五种网络模式的方案。  
    - 阻塞 I/O（blocking IO）  
      获取数据的时候 一直阻塞 直到接收到
    - 非阻塞 I/O（nonblocking IO）
      获取数据时 如果数据还没有完全在内核准备好，会返回error，可以不停重试
    - I/O 多路复用（ IO multiplexing）
      当用户进程调用了select，那么整个进程会被block，而同时，kernel会“监视”所有select负责的socket，当任何一个socket中的数据准备好了
      ，select就会返回。这个时候用户进程再调用read操作，将数据从kernel拷贝到用户进程。
    - 信号驱动 I/O（ signal driven IO）
    - 异步 I/O（asynchronous IO）
      用户进程发起read操作之后，立刻就可以开始去做其它的事。而另一方面，从kernel的角度，当它受到一个asynchronous read之后，
      首先它会立刻返回，所以不会对用户进程产生任何block。然后，kernel会等待数据准备完成，然后将数据拷贝到用户内存，
      当这一切都完成之后，kernel会给用户进程发送一个signal，告诉它read操作完成了。
      
  - 三 IO多路复用  
    多路复用+非阻塞IO有一点：非阻塞只是在读写数据的阶段是非阻塞的，在调用select，poll，epoll的监听阶段还是阻塞的  
    - select
      ```
      int select(int maxfd, fd_set *readset, fd_set *writeset, fd_set *exceptset, const struct timeval *timeout);
      ```
      - Maxfd: 文件描述符的范围，比待检的最大文件描述符大1
      - Readfds：被读监控的文件描述符集
      - Writefds：被写监控的文件描述符集
      - Exceptfds：被异常监控的文件描述符集
      - timeval: 定时器  
        - 如果是nil，表示如果没有i/o select一直等待
        - 如果是非0的值 表示等待固定的一段时间后 从select阻塞调用中返回。
        - 将 tv_sec 和 tv_usec 都设置成 0 表示不等待，检测完立刻返回。
       
      宏:操作向量描述符(a[maxfd-1], ..., a[1], a[0])
        1. void FD_ZERO(fd_set *fdset); 将所有元素设置成0
        2. void FD_SET(int fd, fd_set *fdset); 将a[fd] 设置成1
        3. void FD_CLR(int fd, fd_set *fdset); a[fd] 设置成0
        4. int FD_ISSET(int fd, fd_set *fdset); 判断a[fd]是0还是1  
      
      总结: select 阻塞的获得一个含有若干个待读/写的块的向量，标记了哪个块可读写，然后用户程序去check每个块是否有数据。  
           有文件描述符的个数限制：1024  
           **每次阻塞 读完数据之后，需要重新设置待写入的描述符集合 性能差**
           等待队列和revents在一起
    - poll
        ```
        int poll(struct pollfd *fds, unsigned long nfds, int timeout);
        ```
        - pollfd 数组 
          ```struct pollfd{
                int fd; 描述符
                short events; 待检测的事件类型 ，通过二进制掩码来表示多个不同地事件
                short revents; 每次检测后不会修改传入值， 结果保留在revents
          }
          ```
        - nfds 数组fds的大小 向poll申请的事件检测的个数
        - timeout 如果是一个 <0 的数，表示在有事件发生之前永远等待;如果是 0，表示不阻塞进程，立即 返回;如果是一个 >0 的数，表示 poll 调用方等待指定的毫秒数后返回。
        - 和 select的区别:
          - 区别1：select使用的是定长数组，而poll是通过用户自定义数组长度的形式（pollfd[]）
          - 区别2：select只支持最大fd < 1024，如果单个进程的文件句柄数超过1024，select就不能用了。poll在接口上无限制，
            考虑到每次都要拷贝到内核，一般文件句柄多的情况下建议用epoll。
          - 区别3：select由于使用的是位运算，所以select需要分别设置read/write/error fds的掩码。
          而poll是通过设置数据结构中fd和event参数来实现read/write，比如读为POLLIN，写为POLLOUT，出错为POLLERR：
          - 区别4：**select中fd_set是被内核和用户共同修改的，所以要么每次FD_CLR再FD_SET，**
          **要么备份一份memcpy进去。而poll中用户修改的是events，系统修改的是revents。所以参考muduo的代码，
          都不需要自己去清除revents，从而使得代码更加简洁。**
          - 区别5：select的timeout使用的是struct timeval *timeout，poll的timeout单位是int。
          - 区别6：select使用的是绝对时间，poll使用的是相对时间。
          - 区别7：select的精度是微秒（timeval的分度），poll的精度是毫秒。   
          - 区别8：select的timeout为NULL时表示无限等待，否则是指定的超时目标时间；poll的timeout为-1表示无限等待。所以有用select来实现usleep的。  
          - 区别9：理论上poll可以监听更多的事件
    - epoll[原理](https://blog.csdn.net/armlinuxww/article/details/92803381)
      - 结构
        - int epoll_create(int size): 创建epoll实例，用来调用epoll_ctl和epoll_wait  
          内核可以动态分配需要的内核数据结构
        - int epoll_ctl(int epfd, int op, int fd, struct epoll_evenet *event): 增加或删除监控的事件。
        - int epoll_wait(int epfd, struct epoll_event *events, int maxevents, int timeout): 调用者进程被挂起， 等待内核I/O事件的分发  
      - 数据结构:   
        包含了 Lock、MTX、WQ(等待队列)与 Rdlist 等成员，其中 Rdlist 和 RBR 是我们所关心的。
        - 就绪列表： rdList 采用双向链表，可以快速插入删除。
        - 索引结构： 红黑树（RBR）
      - 对比select  
        Select 低效的原因：
          1. 是将“维护等待队列”和“阻塞进程”两个步骤合二为一。
          2. 只能1个个遍历fds
        ![epoll对比select](https://cynthia-oss.oss-cn-beijing.aliyuncs.com/1634455305983.png)
      - 为什么快
        1. 维护队列和阻塞拆开 epoll_ctl 负责维护等待队列 epoll_wait负责阻塞 避免了每次都要从用户态把数据传给内核。
        2. 内核维护了就绪队列 有红黑树索引 
      - 三者对比
        ![epoll select poll对比](https://cynthia-oss.oss-cn-beijing.aliyuncs.com/1634457114578.png)

          
    
    
#### linux
 - [cpu负载问题](https://mp.weixin.qq.com/s/24vBHgtw5efC9V9yYqknNg)
   - 通过uptime，w或者top命令看到CPU的平均负载。
     - Load Average 系统平均负载：当前系统正在运行的和处于等待运行的进程数之和
     `负载的3个数字，比如上图的4.86，5.28，5.00，分别代表系统在过去的1分钟，5分钟，15分钟内的系统平均负载。`
     - cpu利用率：当前正在运行的进程实时占用CPU的百分比，他是对一段时间内CPU使用状况的统计
   - 负载高，利用率低怎么办？
     `CPU负载很高，利用率却很低，说明处于等待状态的任务很多，负载越高，代表可能很多僵死的进程。通常这种情况是IO密集型的任务，大量请求在请求相同的IO，导致任务队列堆积。`
     - ps -axjf查看是否存在状态为D+状态的进程，这个状态指的就是不可中断的睡眠状态的进程。处于这个状态的进程无法终止，也无法自行退出，只能通过恢复其依赖的资源或者重启系统来解决。
   - 负载很低，利用率却很高
     - top命令找到使用率最高的任务，定位到去看看就行了。如果代码没有问题，那么过段时间CPU使用率就会下降的。
   - CPU使用率达到100%呢？怎么排查？
     - golang delve 堆栈跟踪
 - 常见命令：
   - 文件/目录相关
     - 增删改查
       - 增 touch 创建文件
       - 删除 rm 
       - 改 
         - mv 
         - cp
         - chmod r=4，w=2，x=1
         - chown 用于修改文件和目录的所有者和所属组。一般用法chown user 文件用于修改文件所有者，chown user:user 文件修改文件所有者和组，冒号前面是所有者，后面是组。
         - zip gzip tar
       - 查 
         - ls -a -l
         - cat 一次性展示
         - more
         - less
         - tail -fn 100 展示最后一百行
   - 网络
     - netstat 
       `netstat命令以符号形式显示各种与网络相关的数据结构的内容。有多种输出格式，具体取决于显示信息的选项。该命令的第一种形式显示每个协议的活动套接字列表。第二种形式根据选择的选项显示其他网络数据结构之一的内容。使用第三种形式，并指定等待间隔，netstat将在配置的网络接口上连续显示有关数据包流量的信息。第四种形式显示指定协议或地址族的统计信息。如果指定了等待间隔，将显示最近间隔秒的协议信息。第五种形式显示指定协议或地址族的每个接口的统计信息。第六种形式显示mbuf（9）统计信息。第七种形式显示指定地址系列的路由表。第八种形式显示路由统计信息。`  
     - ifconfig 
       `查本机ip等`
     - hostname
       `主机名用于显示系统的DNS名称，并显示或设置其主机名或NIS域名。`
     - curl
     