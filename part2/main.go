package main

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"sync"
)

const path = "temp.txt"
var wg = sync.WaitGroup{}
var mx = sync.Mutex{}
//var wg2 = sync.WaitGroup{}

func makeFile(p string) (*os.File, error) {
	file, err := os.Create(p)
	if err != nil{
		err := errors.New("не удалось создать файл")
		return nil, err
	}
	return file, nil
}

func writeManyStrings(f *os.File)  {
	//defer f.Close()
	defer wg.Done()
	for i:=0; i<10000; i++{
		fmt.Println("here1")
		//runtime.Gosched()
		mx.Lock()
		f.WriteString("hello\n")
		mx.Unlock()
	}
}

func writeSymbolString(f *os.File) {
	//defer f.Close()
	defer wg.Done()
	for i:=0; i<10; i++{
		fmt.Println("here2")
		runtime.Gosched()
		mx.Lock()
		_, err := f.WriteString("!@!@$#@#!@#!@$!#!@#!@#!@#!@!@#@!@#\n")
		if err != nil{
			fmt.Println(err)
		}
		mx.Unlock()
	}
}

func main()  {
	f, err := makeFile(path)
	defer f.Close()
	if err != nil {
		fmt.Println("ошибка создания файла", err)
		os.Exit(1)
	}
	wg.Add(2)
	go writeManyStrings(f)
	//wg1.Wait()
	//wg2.Add(1)
	go writeSymbolString(f)
	wg.Wait()
fmt.Println("done")



}