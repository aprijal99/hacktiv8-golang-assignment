package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mtx sync.Mutex

	data1 := []interface{}{"bisa1", "bisa2", "bisa3"}
	data2 := []interface{}{"coba1", "coba2", "coba3"}

	for i := 1; i < 5; i++ {
		wg.Add(2)
		go loopData(data1, i, &wg, &mtx)
		go loopData(data2, i, &wg, &mtx)
	}

	wg.Wait()
}

func loopData(data []interface{}, idx int, wg *sync.WaitGroup, mtx *sync.Mutex) {
	mtx.Lock()
	fmt.Println(data, idx)
	mtx.Unlock()

	wg.Done()
}
