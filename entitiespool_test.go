package uniflowpool

import (
	"github.com/google/uuid"
	"log"
	"sync"
	"testing"
	"time"
)

func TestEntitiesPool_Set(t *testing.T) {
	pool := NewEntitiesPool()

	log.Println("Testing method Set in EntityPool structure")
	log.Println("Length of pool:", pool.Length())
	var wg sync.WaitGroup

	i := 0
	for {
		if i == 5 {
			break
		}
		wg.Add(1)
		go func(w *sync.WaitGroup) {
			pool.Set(uuid.NewString(), uuid.NewString(), uuid.NewString(), uuid.NewString(), uuid.NewString())
			w.Done()
		}(&wg)
		i++
	}

	wg.Wait()

	result := 25
	if pool.Length() != result {
		t.Errorf("Method is not correct working: need %d have %d ", result, pool.Length())
		return
	}

	log.Println("Length of pool:", pool.Length())
	log.Println("Method is correct and was SUCCESSFUL done ")
}

func TestEntitiesPool_Get(t *testing.T) {
	pool := NewEntitiesPool()

	log.Println("Testing method Get one in EntityPool structure")
	log.Println("Length of pool:", pool.Length())
	var wg sync.WaitGroup

	i := 0
	for {
		if i == 5 {
			break
		}
		wg.Add(1)
		go func(w *sync.WaitGroup) {
			pool.Set(uuid.NewString(), uuid.NewString(), uuid.NewString(), uuid.NewString(), uuid.NewString())
			w.Done()
		}(&wg)
		i++
	}

	wg.Wait()
	log.Println("Length of pool:", pool.Length())

	i = 0
	for {
		if i == 20 {
			break
		}
		go pool.Get()
		i += 1
	}

	time.Sleep(1 * time.Second)
	result := 5
	if pool.Length() != result {
		t.Errorf("Method is not correct working: need %d have %d ", result, pool.Length())
		return
	}

	log.Println("Length of pool:", pool.Length())
	log.Println("Method is correct and was SUCCESSFUL done ")
}

func TestEntitiesPool_GetAll(t *testing.T) {
	pool := NewEntitiesPool()

	log.Println("Testing method GetAll in EntityPool structure")
	log.Println("Length of pool:", pool.Length())
	var wg sync.WaitGroup

	i := 0
	for {
		if i == 5 {
			break
		}
		wg.Add(1)
		go func(w *sync.WaitGroup) {
			pool.Set(uuid.NewString(), uuid.NewString(), uuid.NewString(), uuid.NewString(), uuid.NewString())
			w.Done()
		}(&wg)
		i++
	}

	wg.Wait()
	log.Println("Length of pool:", pool.Length())

	pool.GetAll()

	result := 0
	if pool.Length() != result {
		t.Errorf("Method is not correct working: need %d have %d ", result, pool.Length())
		return
	}

	log.Println("Length of pool:", pool.Length())
	log.Println("Method is correct and was SUCCESSFUL done ")
}

func TestEntitiesPool_GetCount(t *testing.T) {
	pool := NewEntitiesPool()

	log.Println("Testing method GetCount in EntityPool structure")
	log.Println("Length of pool:", pool.Length())
	var wg sync.WaitGroup

	i := 0
	for {
		if i == 5 {
			break
		}
		wg.Add(1)
		go func(w *sync.WaitGroup) {
			pool.Set(uuid.NewString(), uuid.NewString(), uuid.NewString(), uuid.NewString(), uuid.NewString())
			w.Done()
		}(&wg)
		i++
	}

	wg.Wait()
	log.Println("Length of pool:", pool.Length())

	pool.GetCount(10)

	result := 15
	if pool.Length() != result {
		t.Errorf("Method is not correct working: need %d have %d ", result, pool.Length())
		return
	}

	log.Println("Length of pool:", pool.Length())
	log.Println("Method is correct and was SUCCESSFUL done ")
}
