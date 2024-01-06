package uniflowpool

import "sync"

//EntitiesPool
/*
   This structure is simplify pool which add all new any objects to the tail of queue and get objects from head.
   When you will get objects from pool - the object will be deleted from pool
   Every kept object is  Entity, and it has next field
   **first is link to first structure of Entity
   **last is link to last structure of Entity
   **len is showing currently count objects inside the EntitiesPool
*/
type EntitiesPool struct {
	first *Entity
	last  *Entity

	muExternalLen, muExternalSet, muExternalGet, muExternalGetAll, muExternalGetCount, muInternal sync.Mutex

	len int
}

func NewEntitiesPool() *EntitiesPool {
	pool := EntitiesPool{}

	pool.first = nil
	pool.last = nil

	pool.len = 0

	return &pool
}

// Get one object from pool
// you can use it in multithreading mode
func (pool *EntitiesPool) Get() interface{} {
	pool.muExternalGet.Lock()
	if pool.first == nil {
		defer pool.muExternalGet.Unlock()
		return nil
	} else {
		resultEntity := pool.first
		if resultEntity.Next == nil {
			pool.first = nil
			pool.last = nil
			pool.len = 0
		} else {
			nextInterface := pool.first.Next
			nextEntity := nextInterface.(*Entity)
			nextEntity.Previous = nil
			pool.first = nextEntity
			pool.len -= 1
		}
		defer pool.muExternalGet.Unlock()
		return resultEntity.Value
	}
}

// GetAll all objects from pool
// you can use it in multithreading mode
func (pool *EntitiesPool) GetAll() []interface{} {
	pool.muExternalGetAll.Lock()
	result := make([]interface{}, pool.len-1)
	if pool.len == 0 {
		defer pool.muExternalGetAll.Unlock()
		return result
	}
	for {
		resultEntity := pool.Get()
		if resultEntity == nil {
			break
		}
		result = append(result, resultEntity)
	}
	defer pool.muExternalGetAll.Unlock()
	return result
}

// GetCount is back necessary count of objects from EntitiesPool if count will be more that
// pool has it will back array which will has size equals count, but some of the elements of array will be nil
// you can use it in multithreading mode
func (pool *EntitiesPool) GetCount(count int) []interface{} {
	pool.muExternalGetCount.Lock()
	result := make([]interface{}, count)

	if pool.len == 0 {
		pool.muExternalGetCount.Unlock()
		return result
	}
	for {
		if count == 0 {
			break
		}
		resultEntity := pool.Get()
		if resultEntity == nil {
			pool.muExternalGetCount.Unlock()
			break
		}
		result = append(result, resultEntity)
		count -= 1
	}

	pool.muExternalGetCount.Unlock()
	return result
}

// addElements just joins all any objects to EntitiesPool
// you can use it in multithreading mode
func (pool *EntitiesPool) addElement(element interface{}, wg *sync.WaitGroup) {
	pool.muInternal.Lock()
	defer wg.Done()

	if element != nil {
		newEntity := Entity{
			Value:    element,
			Next:     nil,
			Previous: nil,
		}

		if pool.len == 0 {
			pool.last = &newEntity
			pool.first = &newEntity
		} else {
			newEntity.Previous = pool.last
			pool.last.Next = &newEntity
			pool.last = &newEntity
		}

		pool.len += 1
	}

	pool.muInternal.Unlock()
}

// Set just joins all any objects to EntitiesPool
// you can use it in multithreading mode
func (pool *EntitiesPool) Set(value ...interface{}) {
	pool.muExternalSet.Lock()

	var wg sync.WaitGroup

	for _, val := range value {
		wg.Add(1)
		go pool.addElement(val, &wg)
	}

	wg.Wait()

	pool.muExternalSet.Unlock()
}

// Length just back information about currently count of objects inside EntitiesPool
// you can use it in multithreading mode
func (pool *EntitiesPool) Length() int {
	pool.muExternalLen.Lock()
	defer pool.muExternalLen.Unlock()
	return pool.len
}
