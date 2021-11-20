package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"strconv"
	"sync"
	"time"
)

//Написать программу, которая использует мьютекс для безопасного доступа к
//данным из нескольких потоков. Выполните трассировку программы

var cnt int

var mx = sync.Mutex{}

var (
	readerChan1 = make(chan string, 100)
	readerChan2 = make(chan string, 100)
	writerChan  = make(chan string, 200)
	syncChan    = make(chan string)
)

func WP1(in chan string, out chan string, s chan string) {
	for {
		data, ok := <-in
		if !ok {
			s <- "ok"
			return
		}
		time.Sleep(time.Millisecond * 200)
		mx.Lock()
		cnt++
		out <- data + " cnt is: " + strconv.Itoa(cnt)
		mx.Unlock()
	}
}

func WP2(in chan string, out chan string, s chan string) {
	for {
		data, ok := <-in
		if !ok {
			s <- "ok"
			return
		}
		time.Sleep(time.Millisecond * 500)
		mx.Lock()
		cnt++
		out <- data + " cnt is: " + strconv.Itoa(cnt)
		mx.Unlock()
	}
}

func Closer(s chan string) {
	var cnt = 0
	//fmt.Println(cnt)
	for data := range s {
		if data == "ok" {
			cnt++
		}
		if cnt == 10 {
			close(writerChan)
		}
	}
}

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()


	for i := 0; i < 5; i++ {
	go WP1(readerChan1, writerChan, syncChan) // запускаем 5 потоков, для работы с каналом readerChan1
	}
	for i := 0; i < 5; i++ {
	go WP2(readerChan2, writerChan, syncChan) // запускаем 5 потоков, для работы с каналом readerChan2
	}
	go Closer(syncChan) // горутина принимает из канала сообщения и закрывает пишуший канал

	for i := 0; i < 100; i++ {
		data := "from first WP " + strconv.Itoa(i)
		readerChan1 <- data
		data2 := "from second WP " + strconv.Itoa(i)
		readerChan1 <- data2
	}
	close(readerChan1)
	close(readerChan2)

	for data := range writerChan {

		fmt.Println(data)
	}

}
