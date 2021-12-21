package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	start := time.Now()

	var t *time.Timer
	ch := make(chan bool)
	t = time.AfterFunc(randomDuration1(), func() {
		fmt.Println(time.Now().Sub(start))
		ch <- true
	})

	//label :
	//for {
	//	select {
	//		case <-ch: t.Reset(randomDuration1())
	//				break label
	//
	//	}
	//}

	for time.Since(start) < 5*time.Second {
		<-ch
		t.Reset(randomDuration1())
	}

	time.Sleep(5 * time.Second)
}

func randomDuration1() time.Duration {
	return time.Duration(rand.Int63n(1e9))
}
