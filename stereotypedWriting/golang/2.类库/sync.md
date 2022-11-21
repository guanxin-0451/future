## sync 包
### sync.map-并发安全散列表
  - 功能&对比
    - 普通map 并发读写的情况下 map里的数据会被写乱， 之后就是garbage in和garbage out， 所以直接panic。
    - 对原生的map并发读写时，需要加锁。 
      - 对整个map加1把大锁.
      - 一个大map分成若干小map, 对key进行哈希， 只操作响应小的map.
    - sync.map 读，插入, 删除 都是常数级别时间复杂度。
    - sync.map 的0值是有效的，他在第一次使用后，不允许被复制。
    - cas 理论： CompareAndSet 乐观锁实现  
      是比较和交换，可以这样理解，在内存中，CAS有3个操作数，内存值V，旧的预期值A，要修改的新值B。当且仅当预期值A和内存值V相同时，
      将内存值V修改为B，否则什么都不做。
  - 使用
    - var m sync.Map
    - m.Store("key", value)
    - m.Range(func(key, value interface{}))
    - m.Load(key)
    - m.Delete(key)
    - m.LoadOrStore(key, value) 读取或写入
  - 数据结构
    - map
      ```
      type Map struct {
          // 互斥量mu 保护read和dirty
          mu Mutex
    
          // read contains the portion of the map's contents that are safe for
          // concurrent access (with or without mu held).
          //
          // The read field itself is always safe to load, but must only be stored with
          // mu held.
          //
          // Entries stored in read may be updated concurrently without mu, but updating
          // a previously-expunged entry requires that the entry be copied to the dirty
          // map and unexpunged with mu held.
          //Read包含了map中对于并发访问是安全的部分(无论是否持有mu)。 使用atomic.value 操作系统提供的cas原语，原子性修改
          //read本身总是可以安全加载的，但必须只在获取mu的情况下写入。
          //在read中的Entries stored 可能会在没有锁的情况下并发更新，但是要更新之前删除的entry，先将entry的状态expunged（已删除）改为nil（加锁）， entry复制到dirty map中。
          read atomic.Value // readOnly
    
          // dirty contains the portion of the map's contents that require mu to be
          // held. To ensure that the dirty map can be promoted to the read map quickly,
          // it also includes all of the non-expunged entries in the read map.
    
          // Expunged entries are not stored in the dirty map. An expunged entry in the
          // clean map must be unexpunged and added to the dirty map before a new value
          // can be stored to it.
          //
          // If the dirty map is nil, the next write to the map will initialize it by
          // making a shallow copy of the clean map, omitting stale entries.
          // Dirty包含了map中需要持有mu的部分内容（新写入的key）。为了确保脏映射可以快速提升为读映射，它还包括读映射中所有未删除的条目。
          //
          //删除的条目不会存储在脏映射中。在将新值存储到脏映射之前，必须把read中已删除的部分去掉。
          //
          //如果dirty map为nil，那么对map的下一次写入将通过创建新的clean map的浅拷贝来初始化它，省略被删除的的条目。
          dirty map[interface{}]*entry
    
          // misses counts the number of loads since the read map was last updated that
          // needed to lock mu to determine whether the key was present.
          //
          // Once enough misses have occurred to cover the cost of copying the dirty
          // map, the dirty map will be promoted to the read map (in the unamended
          // state) and the next store to the map will make a new dirty copy.
    
          // 每次从read中读取失败，misses计数+1（加锁）, 加到一定的阈值时，dirty提升为read. 下一个存储到来，会产生一个新的dirty.
          misses int
      }
      ```
    - readOnly
    ```
    // readOnly is an immutable struct stored atomically in the Map.read field.
    // map.read
    type readOnly struct {
        m       map[interface{}]*entry
        amended bool // true if the dirty map contains some key not in m.
    }
    ```
    - entry
      ```
      type entry struct {
        // p points to the interface{} value stored for the entry.
        //
        // If p == nil, the entry has been deleted and m.dirty == nil.
        //
        // If p == expunged, the entry has been deleted, m.dirty != nil, and the entry
        // is missing from m.dirty.
        //
        // Otherwise, the entry is valid and recorded in m.read.m[key] and, if m.dirty
        // != nil, in m.dirty[key].
        //
        // An entry can be deleted by atomic replacement with nil: when m.dirty is
        // next created, it will atomically replace nil with expunged and leave
        // m.dirty[key] unset.
        //
        // An entry's associated value can be updated by atomic replacement, provided
        // p != expunged. If p == expunged, an entry's associated value can be updated
        // only after first setting m.dirty[key] = e so that lookups using the dirty
        // map find the entry.
        p unsafe.Pointer // *interface{}
      }
      ```
      - read和dirty各自维护一套key，key指向的是同一个value（entry）
  - map.Store
    ```
    // Store sets the value for a key.
    func (m *Map) Store(key, value interface{}) {
        read, _ := m.read.Load().(readOnly)
        // 如果read中存在 直接修改entry
        if e, ok := read.m[key]; ok && e.tryStore(&value) {
            return
        }
        
        m.mu.Lock()
        read, _ = m.read.Load().(readOnly)
        if e, ok := read.m[key]; ok {
            // 如果 read map中 存在 该key
            if e.unexpungeLocked() {
                // 删除状态恢复
                // The entry was previously expunged, which implies that there is a
                // non-nil dirty map and this entry is not in it.
                // p == expunged（被删除） 则说明m.dirty != nil 且 m.dirty 不存在该key, 此时应该把删除状态恢复
                m.dirty[key] = e
            }
            e.storeLocked(&value)
        } else if e, ok := m.dirty[key]; ok {
            // dirty中有，read中没有，不用管read miss+1 并且更新下entry就行了
            e.storeLocked(&value)
        } else {
            // 俩map中都不存在该key
            // 如果dirtymap为空，则创建dirtymap， 并且浅拷贝read map 未删除的到dirty map
            // 更新read.amended 字段 标记下俩map有区别了
            // kv写入dirty map中，read不变
            if !read.amended {
                // We're adding the first new key to the dirty map.
                // Make sure it is allocated and mark the read-only map as incomplete.
                m.dirtyLocked()
                m.read.Store(readOnly{m: read.m, amended: true})
            }
            m.dirty[key] = newEntry(value)
        }
        m.mu.Unlock()
    }
    
    // tryStore stores a value if the entry has not been expunged.
    //
    // If the entry is expunged, tryStore returns false and leaves the entry
    // unchanged.
    func (e *entry) tryStore(i *interface{}) bool {
        for {
            p := atomic.LoadPointer(&e.p)
            if p == expunged {
                return false
            }
            if atomic.CompareAndSwapPointer(&e.p, p, unsafe.Pointer(i)) {
                return true
            }
        }
    }
    
    // unexpungeLocked ensures that the entry is not marked as expunged.
    // check entry是否被删除 cas
    // If the entry was previously expunged, it must be added to the dirty map
    // before m.mu is unlocked. 
    func (e *entry) unexpungeLocked() (wasExpunged bool) {
        return atomic.CompareAndSwapPointer(&e.p, expunged, nil)
    }
    
    // storeLocked unconditionally stores a value to the entry.
    // 
    // The entry must be known not to be expunged.
    func (e *entry) storeLocked(i *interface{}) {
        atomic.StorePointer(&e.p, unsafe.Pointer(i))
    }
    ```
  - map.Load
    ```
    / Load returns the value stored in the map for a key, or nil if no
    // value is present.
    // The ok result indicates whether value was found in the map.
    / Load返回存储在map中的一个键的值，如果没有，则返回nil 
    //值已经存在。
    // ok结果表示是否在map中找到value。
    func (m *Map) Load(key interface{}) (value interface{}, ok bool) {
        read, _ := m.read.Load().(readOnly)
        e, ok := read.m[key]
        if !ok && read.amended {
            m.mu.Lock()
            // Avoid reporting a spurious miss if m.dirty got promoted while we were
            // blocked on m.mu. (If further loads of the same key will not miss, it's
            // not worth copying the dirty map for this key.)
            read, _ = m.read.Load().(readOnly)
            e, ok = read.m[key]
            if !ok && read.amended {
                e, ok = m.dirty[key]
                // Regardless of whether the entry was present, record a miss: this key
                // will take the slow path until the dirty map is promoted to the read
                // map.
                // 记录miss值和控制dirty晋升
                m.missLocked()
            }
            m.mu.Unlock()
        }
        if !ok {
            return nil, false
        }
        return e.load()
    }
    
    func (e *entry) load() (value interface{}, ok bool) {
        p := atomic.LoadPointer(&e.p)
        if p == nil || p == expunged {
            return nil, false
        }
        return *(*interface{})(p), true
    }
    ```
    - map.Delete
      ```
      // Delete deletes the value for a key.
      func (m *Map) Delete(key interface{}) {
          read, _ := m.read.Load().(readOnly)
          e, ok := read.m[key]
          if !ok && read.amended {
              // read中没有且read和dirty不一致
              m.mu.Lock()
              read, _ = m.read.Load().(readOnly)
              e, ok = read.m[key]
              if !ok && read.amended {
                  delete(m.dirty, key)
              }
              m.mu.Unlock()
          }
          if ok {
              e.delete()
          }
      }
    
      func (e *entry) delete() (hadValue bool) {
          for {
              p := atomic.LoadPointer(&e.p)
              if p == nil || p == expunged {
                  return false
              }
      
              if atomic.CompareAndSwapPointer(&e.p, p, nil) {
                  return true
              }
          }
      }
    ```
  - map.Range 
    - 如果amended为true，说明不一致，不一致需要全部遍历一遍dirty，所以直接将dirty提升为read(需要加锁)，之后遍历read（不需要锁）.
### sync.map怎么做到并发安全的？
  - 读写拆分2个map
    - read atomic.Value
    - dirty map[interface{}]*entry
  - read，dirty 维护不同的key, 但是value都指向同1个entry
  - 读: 从read中读 如果read，dirty不一致，再从dirty读（直接加锁），miss+1， miss达到一定数量，dirty升级成read（StorePointer进行值的复制，复制到对应指针上）
  - 写: 需要注意删除的状态恢复。写dirty map 标记两个map不同。
  - 删:  read中没有且read和dirty不一致 删除dirty的key 及标记entry的状态为nil
  - range: 如果amended为true，说明不一致，不一致需要全部遍历一遍dirty，所以直接将dirty提升为read(需要加锁)，之后遍历read（不需要锁）.
  - 适合读多写少。