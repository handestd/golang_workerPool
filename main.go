package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rappers := []string{"Dế Choắt", "GDucky", "MCK", "TLinh", "Rtee"}
	numBeat := 100

	wg := &sync.WaitGroup{}
	beatChannel := make(chan string, 10)

	// produceBeat
	go produceBeat(numBeat, beatChannel)

	// Write rap song
	for _, rapper := range rappers {
		wg.Add(1)
		go writeSong(wg, rapper, beatChannel)
	}

	wg.Wait()

}

func writeSong(wg *sync.WaitGroup, rapper string, beatChannel chan string) {
	defer wg.Done()

	for {
		beat, ok := <-beatChannel
		if !ok {
			fmt.Printf("Hết beat. %s ra khỏi phòng thu\n", rapper)
			return
		}

		fmt.Printf("Rapper %s sử dụng beat %s\n", rapper, beat)
		wait()
	}
}

func produceBeat(numBeat int, beatChannel chan string) {
	for i := 1; i <= numBeat; i++ {
		wait()
		fmt.Println("Tạo ra beat số: ", i)
		beatChannel <- fmt.Sprintf("Beat %d", i)
	}

	close(beatChannel)
}

func wait() {
	time.Sleep(time.Second * time.Duration(rand.Intn(2)))
}
