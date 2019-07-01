package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func filosofo(i int, turnoBuff chan int, tenedores []sync.Mutex, cantFilosofos int, ready *sync.WaitGroup) {
	ready.Done()
	ready.Wait()
	for {
		turnoBuff <- i
		fmt.Println("#", i, "hambriento...")
		tenedores[i].Lock()
		tenedores[(i+1)%cantFilosofos].Lock()
		fmt.Println("#", i, "comiendo...")
		n := (rand.Intn(100))
		time.Sleep(time.Duration(n) * time.Millisecond)
		tenedores[i].Unlock()
		tenedores[(i+1)%cantFilosofos].Unlock()
		fmt.Println("#", i, "pensando...")
		<-turnoBuff
	}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	cantFilosofos := 5
	turnoBuff := make(chan int, cantFilosofos-1)
	tenedores := make([]sync.Mutex, cantFilosofos)
	var ready sync.WaitGroup
	ready.Add(cantFilosofos)
	for i := 0; i < cantFilosofos; i++ {
		go filosofo(i, turnoBuff, tenedores, cantFilosofos, &ready)
	}
	time.Sleep(500 * time.Millisecond)
}
