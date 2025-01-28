package main

import (
	"fmt"
	"main/requests_babel"
	"os"
	"sync"
	"time"
)

type Container struct {
    mu       sync.Mutex
    f *os.File
}

func main() {
	debut := time.Now()
	wg := &sync.WaitGroup{}
	for i:=1; i<100; i++{
		dir := fmt.Sprintf("/mnt/d/dossiers/insa/dataset ELP/hex%d",i)
		fmt.Println(dir)
		file,_ := os.Create(dir)
		c := Container{
			f : file,
		}
		for w:=1 ; w<=4 ; w++{
			for s:=1 ; s<=5; s++{
				for v:=0 ; v<=32 ; v++ {
					wg.Add(1)
					go func(){
						defer wg.Done()
						a:=requests_babel.Get_book(1,w,s,v)
						c.mu.Lock()
						defer c.mu.Unlock()
						c.f.WriteString(a)
					}()
				}
				wg.Wait()
				fmt.Printf("%f percent of the way there \n",(float32(w-1)/4)*100+(float32(s)/5)*10)
			}
		}
		c.f.Sync()
		fmt.Println(time.Now().Sub(debut))
	}
}