## context
### context 是什么
  - goroutine 的上下文
  - 取消信号  
    http请求取消
  - 超时时间  
    下游服务变慢，长耗时的goroutine变多，性能影响，可能导致服务不可用
  - 截止时间
  - k-v
  - 在go 内 不能直接杀死协程， 协程的关闭一般用channel和select的方式来控制。  
    而当处理一个请求产生很多协程，协程间互相关联，需要共享变量，可以被同时关闭等。这时用channel比较慢烦，通过context来实现。
  - ***context***用来解决goroutine之间的退出通知，元数据传递的问题的功能。
### 官方文档对使用的建议
  - 不要将context塞到结构体里，直接作为函数的第一参数。
  - 不要向函数传入1个nil的context。如果不知道传什么，可以传context todo
  - 不要把本该作为函数参数的类型赛道context中，context只承担共同的数据（框架层）
  - context是并发安全的  
    context.WithValue 写数据的时候必须创建新的子 context 不存在并发写。

### context 数据结构
  - context interface
    ```
    type Context interface {
        // Deadline returns the time when work done on behalf of this context
        // should be canceled. Deadline returns ok==false when no deadline is
        // set. Successive calls to Deadline return the same results.
        // Deadline返回代表该上下文完成的工作应该被取消的时间。如果没有设置截止日期，则返回ok==false。连续调用Deadline会返回相同的结果。
        Deadline() (deadline time.Time, ok bool)
    
        // Done returns a channel that's closed when work done on behalf of this
        // context should be canceled. Done may return nil if this context can
        // never be canceled. Successive calls to Done return the same value.
        //
        // WithCancel arranges for Done to be closed when cancel is called;
        // WithDeadline arranges for Done to be closed when the deadline
        // expires; WithTimeout arranges for Done to be closed when the timeout
        // elapses.
        //
    // Done返回一个通道，当代表该上下文所做的工作应被取消时，该通道将关闭。如果这个上下文永远不能被取消，Done可能会返回nil。连续调用Done将返回相同的值。 
    // 当取消被调用时，WithCancel安排Done被关闭;
    // WithDeadline安排在截止日期到期时完成;WithTimeout安排Done在超时时关闭
    //在select语句中使用Done:

        // Done is provided for use in select statements:
        //
        //  // Stream generates values with DoSomething and sends them to out
        //  // until DoSomething returns an error or ctx.Done is closed.
        //  func Stream(ctx context.Context, out chan<- Value) error {
        //  	for {
        //  		v, err := DoSomething(ctx)
        //  		if err != nil {
        //  			return err
        //  		}
        //  		select {
        //  		case <-ctx.Done():
        //  			return ctx.Err()
        //  		case out <- v:
        //  		}
        //  	}
        //  }
        //
        // See https://blog.golang.org/pipelines for more examples of how to use
        // a Done channel for cancelation.
        Done() <-chan struct{}
    
        // If Done is not yet closed, Err returns nil.
        // If Done is closed, Err returns a non-nil error explaining why:
        // Canceled if the context was canceled
        // or DeadlineExceeded if the context's deadline passed.
        // After Err returns a non-nil error, successive calls to Err return the same error.
        //如果Done未关闭，Err返回nil。 
        //如果Done被关闭，Err返回一个非nil错误解释原因:
        //如果上下文被取消，则取消
        //或DeadlineExceeded，如果上下文的截止日期过去了。
        //在Err返回非nil错误后，连续调用Err将返回相同的错误。
        Err() error
    
        // Value returns the value associated with this context for key, or nil
        // if no value is associated with key. Successive calls to Value with
        // the same key returns the same result.
        //
        // Use context values only for request-scoped data that transits
        // processes and API boundaries, not for passing optional parameters to
        // functions.
        //
        // A key identifies a specific value in a Context. Functions that wish
        // to store values in Context typically allocate a key in a global
        // variable then use that key as the argument to context.WithValue and
        // Context.Value. A key can be any type that supports equality;
        // packages should define keys as an unexported type to avoid
        // collisions.
        //
        // Value返回与此上下文关联的key值，如果key没有关联值，则返回nil。连续调用具有相同键的Value将返回相同的结果。 
        // 上下文值只能用于传递进程和API边界的请求范围内的数据，而不能用于向函数传递可选参数。键标识上下文中的特定值。希望在Context中存储值的函数通常会在全局变量中分配一个键，然后将该键用作Context的参数。WithValue和
        // Context.Value。键可以是任何支持相等的类型;
        // 包应该将键定义为要避免的未导出类型
        // Packages that define a Context key should provide type-safe accessors
        // for the values stored using that key:
        //
        // 	// Package user defines a User type that's stored in Contexts.
        // 	package user
        //
        // 	import "context"
        //
        // 	// User is the type of value stored in the Contexts.
        // 	type User struct {...}
        //
        // 	// key is an unexported type for keys defined in this package.
        // 	// This prevents collisions with keys defined in other packages.
        // 	type key int
        //
        // 	// userKey is the key for user.User values in Contexts. It is
        // 	// unexported; clients use user.NewContext and user.FromContext
        // 	// instead of using this key directly.
        // 	var userKey key
        //
        // 	// NewContext returns a new Context that carries value u.
        // 	func NewContext(ctx context.Context, u *User) context.Context {
        // 		return context.WithValue(ctx, userKey, u)
        // 	}
        //
        // 	// FromContext returns the User value stored in ctx, if any.
        // 	func FromContext(ctx context.Context) (*User, bool) {
        // 		u, ok := ctx.Value(userKey).(*User)
        // 		return u, ok
        // 	}
        Value(key interface{}) interface{}
    }
    ```

### 总结
  - context 是什么  
    - 做协程间通信用的*上下文*
    - 本质是个interface
    - context
    - canceler  
      支持cancel的interface 需要cancel的context需要实现该接口
    - cancelContext
      - 可以取消的context
    - timerContext
      - 定时取消的context
  - context 的结构
    - cancel ctx
      - Context
      - mu
      - done chan struct{} // 是一个chanel 负责取消通知， 懒汉式创建 
      - children map[canceler]struct{} 记录所有儿子
      - 不要将context塞到结构体里，直接作为函数的第一参数。因为删除的时候parent.Value()有类型check。 只会识别 cancelCtx 和 timerCtx
  - 如何取消？
    - 1.递归取消自己及所有的子节点
    - 2.删除自己，以及从父节点的children中清除掉。
    - 3.删除下游所有儿子
  - 如何取值？
    - 从当前context 顺着链路一路向上找父节点，是否有该key。
    - 存储的时候 多个context 同样的key 会共同存储，但是查找时，只会找到最近的一个
    - 所以，不建议业务数据使用context，非常不清晰
    