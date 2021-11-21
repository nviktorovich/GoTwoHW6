package main

import (
	"fmt"
	"sync"
)

//Смоделировать ситуацию “гонки”, и проверить программу на наличии “гонки”


const max = 1000

func main() {
	var (
		wg sync.WaitGroup
		count = 1000
		mx = sync.Mutex{}
	)
	wg.Add(max)
	for i := 0; i < max ; i++ {
		go func() {
			defer wg.Done()
			mx.Lock()
			count--
			mx.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(count)
}
