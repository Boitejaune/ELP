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
	for i:=13; i<=13; i++{
		dir := fmt.Sprintf("/mnt/d/dossiers/insa/dataset ELP/hex%d",i)
		fmt.Println(dir)
		file,_ := os.Create(dir)
		c := Container{
			f : file,
		}
		for w:=1 ; w<=4 ; w++{
			for s:=1 ; s<=5; s++{
				order := 1
				for v:=1 ; v<=32 ; v++ {
					wg.Add(1)
					go func(){
						defer wg.Done()
						a:=requests_babel.Get_book(i,w,s,v)
						for order != v{
							
						}
						c.mu.Lock()
						c.f.WriteString(a)
						c.mu.Unlock()
						order += 1
					}()
				}
				wg.Wait()
				fmt.Printf("%f percent of the way there \n",(float32(w-1)/4)*100+(float32(s))*5)
			}
		}
		c.f.Sync()
		c.f.Close()
		fmt.Println(time.Now().Sub(debut))
	}
}