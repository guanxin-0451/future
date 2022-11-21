## 通道
### csp是什么
  - 解释:  
    Communicating sequential processes(顺序通信)
  - golang 通过通信来实现内存共享，而不是共享内存。
  - 尽量使用channel,把goroutine 当作免费的资源，随意使用。
### 通道的应用
  - 停止信号  
    通过channel使接收方停止
  - 定时任务  
    和定时器结合
    ```
    ticker := time.Tick(1 *time.Second)
    ...
    case <- ticker:
        do()
    ```
  - 解耦生产方消费房
  - 控制并发goroutine数量
    - 构建1个带缓冲的channel "token", 并发的goroutine 通过获取和归还token,来控制并发数量
### 数据结构
  - hchan
  ```
        // src/runtime/chan.go
        type hchan struct {
            qcount   uint           // total data in the queue channel元素数量
            dataqsiz uint           // size of the circular queue  底层 circular queue 长度
            buf      unsafe.Pointer // points to an array of dataqsiz elements  指向底层 circular queue 的指针 只针对有缓存的chan
            elemsize uint16         // chan 中元素大小
            closed   uint32         // chan 是否已关闭
            elemtype *_type // element type   chan中元素类型
            sendx    uint   // send index     已发送元素在循环数组内的索引
            recvx    uint   // receive index  已接收元素在循环数组中的索引
            recvq    waitq  // list of recv waiters  等待接收的goroutine队列
            sendq    waitq  // list of send waiters   等待发送的goroutine队列

            // hchan的所有字段及sudogs（waitq中的sudo g 本质使 goroutine）中的部分字段加锁
            // waitq使sudogs的双向链表  sudog是对goroutine的封装
            // lock protects all fields in hchan, as well as several
            // fields in sudogs blocked on this channel.
            //
            // Do not change another G's status while holding this lock
            // (in particular, do not ready a G), as this can deadlock
            // with stack shrinking.
            
            lock mutex
        }
   ```
  - 主要字段
    - buf  
      底层循环数组的指针 存储channel元素
    - recvq  
      sudog(goroutine的封装)的双向链表 等待接收的goroutine队列
    - sendq  
      sudog的双向链表 等待发送的goroutine队列
    - lock  
      保证每个读channel或者写channel的操作都是原子操作
### 操作
  - 创建
    ```
    ch1 := make(chan int)
    ch2 := make(chan int, 10)
    
    // src/runtime/chan.go
    func makechan(t *chantype, size int) *hchan {
        elem := t.elem

        // compiler checks this but be safe.
        // 检查channel的size 及内存对齐
        if elem.size >= 1<<16 {
            throw("makechan: invalid channel element type")
        }
        if hchanSize%maxAlign != 0 || elem.align > maxAlign {
            throw("makechan: bad alignment")
        }
        // 计算chan需要的内存大小（类型的size *数量），计算是否溢出
        mem, overflow := math.MulUintptr(elem.size, uintptr(size))
        if overflow || mem > maxAlloc-hchanSize || size < 0 {
            panic(plainError("makechan: size out of range"))
        }

        // Hchan does not contain pointers interesting for GC when elements stored in buf do not contain pointers.
        // buf中的元素不包含指针时，gc不会回收hchan(不关心其他指针)。
        // buf points into the same allocation, elemtype is persistent.
        // buf 指向相同的分配（所有chan共享同一个buf?），元素类型持久。
        // SudoG's are referenced from their owning thread so they can't be collected.
        // sudoG 被所属线程饮用 他们不能被收集
        // TODO(dvyukov,rlh): Rethink when collector can move allocated objects.
        // todo: 何时gc可以移动已分配的对象？
        var c *hchan
        switch {
        case mem == 0:
            // Queue or element size is zero.
            // 容量为0
            // 非缓冲型的，buf没用，直接指向chan的起始地址处
            // 缓冲型的, 说明元素无指针且元素类型为strcut{}
            // 因为只会用到接收和发送游标， 不会真正复制东西到c.buf， 否则会覆盖chan
            c = (*hchan)(mallocgc(hchanSize, nil, true))
            // Race detector uses this location for synchronization.
            c.buf = c.raceaddr()
        case elem.kind&kindNoPointers != 0:
            // Elements do not contain pointers.
            // 不包含指针 gc不会扫描chan中的元素
            // Allocate hchan and buf in one call.
            // 只进行1次分配
            c = (*hchan)(mallocgc(hchanSize+mem, nil, true))
            c.buf = add(unsafe.Pointer(c), hchanSize)
        default:
            // Elements contain pointers.
            // chan和buf2次分配内存
            c = new(hchan)
            c.buf = mallocgc(mem, elem, true)
        }
        c.elemsize = uint16(elem.size)
        c.elemtype = elem
        c.dataqsiz = uint(size)

        if debugChan {
            print("makechan: chan=", c, "; elemsize=", elem.size, "; elemalg=", elem.alg, "; dataqsiz=", size, "\n")
        }
        return c
    }
    ```
    - 容量为0 buf直接指向chan的起始地址处， 不会真正复制东西到c.buf， 否则会覆盖chan
    - 带缓冲区但是不包含指针， gc不需要扫描chan中元素，只进行一次分配
    - 带缓冲区，内容为指针: 进行2次分配
    
  - 接收： 从channel 中读
    ```
    // chanrecv receives on channel c and writes the received data to ep.
    // 从c上接收写到ep
    // ep may be nil, in which case received data is ignored.
    // ep可能为nil，在这种情况下接收的数据将被忽略。
    // If block == false and no elements are available, returns (false, false).
    // 如果block == false且没有可用的元素，则返回(false, false)。
    // Otherwise, if c is closed, zeros *ep and returns (true, false).
    // ||如果c是关闭的，则 将 *ep清空并返回(true, false)。
    // Otherwise, fills in *ep with an element and returns (true, true).
    // ||用元素填充*ep并返回(true, true)。
    // A non-nil ep must point to the heap or the caller's stack.
    // 非空的ep必须指向堆或调用方的堆栈。
    func chanrecv(c *hchan, ep unsafe.Pointer, block bool) (selected, received bool) {
    // raceenabled: don't need to check ep, as it is always on the stack
    // or is new memory allocated by reflect.

        if debugChan {
            print("chanrecv: chan=", c, "\n")
        }

        if c == nil {
            if !block {
                return
            }
            gopark(nil, nil, waitReasonChanReceiveNilChan, traceEvGoStop, 2)
            throw("unreachable")
        }

        // Fast path: check for failed non-blocking operation without acquiring the lock.
        //
        // After observing that the channel is not ready for receiving, we observe that the
        // channel is not closed. Each of these observations is a single word-sized read
        // (first c.sendq.first or c.qcount, and second c.closed).
        // Because a channel cannot be reopened, the later observation of the channel
        // being not closed implies that it was also not closed at the moment of the
        // first observation. We behave as if we observed the channel at that moment
        // and report that the receive cannot proceed.
        //
        // The order of operations is important here: reversing the operations can lead to
        // incorrect behavior when racing with a close.
        //快速路径:检查未获得锁的非阻塞操作是否失败。
        //在观察到信道还没有准备好接收之后，我们观察到通道未关闭。每一个观察都是一个单词大小的阅读量
        //(第一个c.sendq.first或c.qcount，第二个c.closed)。
        // 因为通道不能重新打开，所以后来观察到的通道没有关闭意味着在第一次观察时它也没有关闭。我们的行为就像我们在那一刻观察到了通道，并报告接收无法继续。
        //操作的顺序很重要:反转操作可以导致当运行结束时的错误行为。
        if !block && (c.dataqsiz == 0 && c.sendq.first == nil ||
            c.dataqsiz > 0 && atomic.Loaduint(&c.qcount) == 0) &&
            atomic.Load(&c.closed) == 0 {
            // 因为channel不可能被重复打开，如果前一个是未关闭的，就可以结束了
            return
        }

        var t0 int64
        if blockprofilerate > 0 {
            t0 = cputicks()
        }

        lock(&c.lock)

        if c.closed != 0 && c.qcount == 0 {
            if raceenabled {
                raceacquire(c.raceaddr())
            }
            unlock(&c.lock)
            if ep != nil {
                typedmemclr(c.elemtype, ep)
            }
            return true, false
        }

        if sg := c.sendq.dequeue(); sg != nil {
            // Found a waiting sender. If buffer is size 0, receive value
            // directly from sender. Otherwise, receive from head of queue
            // and add sender's value to the tail of the queue (both map to
            // the same buffer slot because the queue is full).
            // 有等待发送队列 说明buf满了。
            // 发现一个等待的发送者。如果缓冲区大小为0，则直接从发送端接收值。 直接进行内存复制，接收头部，写尾部。
            // 否则，
            return true, true
        }

        if c.qcount > 0 {
            // Receive directly from queue
            qp := chanbuf(c, c.recvx)
            if raceenabled {
                raceacquire(qp)
                racerelease(qp)
            }
            if ep != nil {
                typedmemmove(c.elemtype, ep, qp)
            }
            typedmemclr(c.elemtype, qp)
            c.recvx++
            if c.recvx == c.dataqsiz {
                c.recvx = 0
            }
            c.qcount--
            unlock(&c.lock)
            return true, true
        }

        if !block {
            unlock(&c.lock)
            return false, false
        }

        // no sender available: block on this channel.
        gp := getg()
        mysg := acquireSudog()
        mysg.releasetime = 0
        if t0 != 0 {
            mysg.releasetime = -1
        }
        // No stack splits between assigning elem and enqueuing mysg
        // on gp.waiting where copystack can find it.
        mysg.elem = ep
        mysg.waitlink = nil
        gp.waiting = mysg
        mysg.g = gp
        mysg.isSelect = false
        mysg.c = c
        gp.param = nil
        c.recvq.enqueue(mysg)
        goparkunlock(&c.lock, waitReasonChanReceive, traceEvGoBlockRecv, 3)

        // someone woke us up
        if mysg != gp.waiting {
            throw("G waiting list is corrupted")
        }
        gp.waiting = nil
        if mysg.releasetime > 0 {
            blockevent(mysg.releasetime-t0, 2)
        }
        closed := gp.param == nil
        gp.param = nil
        mysg.c = nil
        releaseSudog(mysg)
        return true, !closed
    }

    // recv processes a receive operation on a full channel c.
    // There are 2 parts:
    // 1) The value sent by the sender sg is put into the channel
    //    and the sender is woken up to go on its merry way.
    // 2) The value received by the receiver (the current G) is
    //    written to ep.
    // For synchronous channels, both values are the same.
    // For asynchronous channels, the receiver gets its data from
    // the channel buffer and the sender's data is put in the
    // channel buffer.
    // Channel c must be full and locked. recv unlocks c with unlockf.
    // sg must already be dequeued from c.
    // A non-nil ep must point to the heap or the caller's stack.
    func recv(c *hchan, sg *sudog, ep unsafe.Pointer, unlockf func(), skip int) {
        if c.dataqsiz == 0 {
        if raceenabled {
            racesync(c, sg)
        }
        if ep != nil {
        // copy data from sender
            recvDirect(c.elemtype, sg, ep)
        }
        } else {
        // Queue is full. Take the item at the
        // head of the queue. Make the sender enqueue
        // its item at the tail of the queue. Since the
        // queue is full, those are both the same slot.
            qp := chanbuf(c, c.recvx)
            if raceenabled {
                raceacquire(qp)
                racerelease(qp)
                raceacquireg(sg.g, qp)
                racereleaseg(sg.g, qp)
            }
            // copy data from queue to receiver
            if ep != nil {
                typedmemmove(c.elemtype, ep, qp)
            }
            // copy data from sender to queue
            typedmemmove(c.elemtype, qp, sg.elem)
            c.recvx++
            if c.recvx == c.dataqsiz {
                c.recvx = 0
            }
            c.sendx = c.recvx // c.sendx = (c.sendx+1) % c.dataqsiz
        }
        sg.elem = nil
        gp := sg.g
        unlockf()
        gp.param = unsafe.Pointer(sg)
        if sg.releasetime != 0 {
            sg.releasetime = cputicks()
        }
        goready(gp, skip+1)
    }
    ```
    - 操作流程:
      - 1.快速通道，非阻塞调用，不用获取锁 检查锁失败 atomic.Load(&c.closed) == 0 channel不能被重复打开。如果是为关
      - 非阻塞调用：
        - 1.加锁
        - 2.读buf（需要判断buf是否已满）
      - 阻塞调用:
        - 调起gopark挂起goroutine
    - 读：  
      读已经关闭的 chan，能一直读到内容，但是读到的内容根据通道内关闭前是否有元素而不同。
      如果 chan 关闭前，buffer 内有元素还未读，会正确读到 chan 内的值，且返回的第二个 bool 值为 true；
      如果 chan 关闭前，buffer 内有元素已经被读完，chan 内无值，返回 channel 元素的零值，第二个 bool 值为 false。
    - 写：  
      写已经关闭的 chan 会 panic。
    
    - block=true: 阻塞模式 会调起gopark挂起goroutine
    - channel是nil的情况下，读的goroutine会被阻塞
  - 发送: 
      ```
      func chansend(c *hchan, ep unsafe.Pointer, block bool, callerpc uintptr) bool {
        if c == nil {
            if !block {
                return false
            }
            gopark(nil, nil, waitReasonChanSendNilChan, traceEvGoStop, 2)
            throw("unreachable")
        }
    
        if debugChan {
            print("chansend: chan=", c, "\n")
        }
    
        if raceenabled {
            racereadpc(c.raceaddr(), callerpc, funcPC(chansend))
        }
    
        // Fast path: check for failed non-blocking operation without acquiring the lock.
        //
        // After observing that the channel is not closed, we observe that the channel is
        // not ready for sending. Each of these observations is a single word-sized read
        // (first c.closed and second c.recvq.first or c.qcount depending on kind of channel).
        // Because a closed channel cannot transition from 'ready for sending' to
        // 'not ready for sending', even if the channel is closed between the two observations,
        // they imply a moment between the two when the channel was both not yet closed
        // and not ready for sending. We behave as if we observed the channel at that moment,
        // and report that the send cannot proceed.
        //
        // It is okay if the reads are reordered here: if we observe that the channel is not
        // ready for sending and then observe that it is not closed, that implies that the
        // channel wasn't closed during the first observation.
        if !block && c.closed == 0 && ((c.dataqsiz == 0 && c.recvq.first == nil) ||
            (c.dataqsiz > 0 && c.qcount == c.dataqsiz)) {
            return false
        }
    
        var t0 int64
        if blockprofilerate > 0 {
            t0 = cputicks()
        }
    
        lock(&c.lock)
    
        if c.closed != 0 {
            unlock(&c.lock)
            panic(plainError("send on closed channel"))
        }
    
        if sg := c.recvq.dequeue(); sg != nil {
            // Found a waiting receiver. We pass the value we want to send
            // directly to the receiver, bypassing the channel buffer (if any).
            send(c, sg, ep, func() { unlock(&c.lock) }, 3)
            return true
        }
    
        if c.qcount < c.dataqsiz {
            // Space is available in the channel buffer. Enqueue the element to send.
            qp := chanbuf(c, c.sendx)
            if raceenabled {
                raceacquire(qp)
                racerelease(qp)
            }
            typedmemmove(c.elemtype, qp, ep)
            c.sendx++
            if c.sendx == c.dataqsiz {
                c.sendx = 0
            }
            c.qcount++
            unlock(&c.lock)
            return true
        }
    
        if !block {
            unlock(&c.lock)
            return false
        }
    
        // Block on the channel. Some receiver will complete our operation for us.
        gp := getg()
        mysg := acquireSudog()
        mysg.releasetime = 0
        if t0 != 0 {
            mysg.releasetime = -1
        }
        // No stack splits between assigning elem and enqueuing mysg
        // on gp.waiting where copystack can find it.
        mysg.elem = ep
        mysg.waitlink = nil
        mysg.g = gp
        mysg.isSelect = false
        mysg.c = c
        gp.waiting = mysg
        gp.param = nil
        c.sendq.enqueue(mysg)
        goparkunlock(&c.lock, waitReasonChanSend, traceEvGoBlockSend, 3)
        // Ensure the value being sent is kept alive until the
        // receiver copies it out. The sudog has a pointer to the
        // stack object, but sudogs aren't considered as roots of the
        // stack tracer.
        KeepAlive(ep)
    
        // someone woke us up.
        if mysg != gp.waiting {
            throw("G waiting list is corrupted")
        }
        gp.waiting = nil
        if gp.param == nil {
            if c.closed == 0 {
                throw("chansend: spurious wakeup")
            }
            panic(plainError("send on closed channel"))
        }
        gp.param = nil
        if mysg.releasetime > 0 {
            blockevent(mysg.releasetime-t0, 2)
        }
        mysg.c = nil
        releaseSudog(mysg)
        return true
      }
      ```
    - 有 goroutine 阻塞在 channel recv 队列上，此时缓存队列为空，直接将消息发送给 reciever goroutine,只产生一次复制。
    - 当 channel 缓存队列有剩余空间时，将数据放到队列里，等待接收，接收后总共产生两次复制。
    - 当 channel 缓存队列已满时，将当前 goroutine 加入 send 队列并阻塞。
  - 关闭
    - recvq和sendq中分别保存了阻塞的发送者和接收者。
    - 关闭channel后，对于等待接收者而言，会收到1个相应类型的零值。
    - 对于等待发送者来说, 会panic。
    - 通道关闭后，ok为false读到的才无效，否则还会有缓冲区中的数据。
    - 如何优雅地关闭通道？
      - 关闭1个close的channel会导致panic。
      - 向1个close的channel发送消息会导致panic。
      - 1 一个sender： sender处关闭即可。
      - 2 多个sender，一个receiver  receiver处发送一个关闭信号的channel发送给所有sender，senders监听到数据后，停止发送数据，不关闭channel，等待gc回收。
      - 3 多个sender，多个receiver 加中间人，接收所有关闭信号，再决定关闭。
### 收发数据的本质
  - 本质都是值的复制
### happens before
  - 如果事件a和事件b存在happened-before 关系。即a->b 那么a，b完成后的结果一定要体现出这种关系。
  - go中主协程退出不回等待其他协程。
  - 现代编译器/cpu 会做编译器重排，内存重排。
  - channel发送的happens before
    - 第n个send一定在happens-before 第 n 个 receive finished.
    - 容量为m的缓冲型channel, 第n个reveive 一定happens-before 第n+m 个send finished。
### 泄漏
  - 泄漏：goroutine操作channel后，处于发送或接受阻塞状态，channel处于满或空的状态，垃圾回收也不会回收此类资源。
### panic
  - 向1个关闭的channel写
  - 关闭1个nil/closed channel。
  - 读写一个nil channel 会被无限阻塞。

### 通道操作情况总结

| 操作 | nil channel | closed channel | not nil not closed channel |
| --- | --- | --- | --- |
| close | panic | panic | 正常关闭 |
| 读 <- ch |阻塞  | 对应类型的0值 缓冲型会继续读|  阻塞或正常读取数据, 缓冲型channel为空或非缓冲型channel没有等待的发送者时会阻塞 |
| 写 ch <- | 阻塞 | panic | 阻塞或正常写入数据, 非缓冲型没等待的接收者或者有缓冲的buf满了 会阻塞 |


### 参考
  - 《go程序员面试宝典》
