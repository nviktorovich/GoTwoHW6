package main

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
)
//Написать многопоточную программу, в которой будет использоваться явный вызов
//планировщика. Выполните трассировку программы

const path = "temp.txt"
var wg = sync.WaitGroup{}
var mx = sync.Mutex{}

func makeFile(p string) (*os.File, error) {
	file, err := os.Create(p)
	if err != nil{
		err := errors.New("не удалось создать файл")
		return nil, err
	}
	return file, nil
}

func writeManyStrings(f *os.File)  {
	defer wg.Done()
	for i:=0; i<10000; i++{
		mx.Lock()
		f.WriteString("hello\n")
		mx.Unlock()
	}
}

func writeSymbolString(f *os.File) {
	defer wg.Done()
	for i:=0; i<10; i++{
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
	trace.Start(os.Stderr)
	defer trace.Stop()

	f, err := makeFile(path)
	defer f.Close()
	if err != nil {
		fmt.Println("ошибка создания файла", err)
		os.Exit(1)
	}
	wg.Add(2)
	go writeManyStrings(f)
	go writeSymbolString(f)
	wg.Wait()
fmt.Println("done")



}