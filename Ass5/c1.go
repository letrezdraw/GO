package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Checkpoint struct {
	mu           sync.Mutex
	arrived      int
	totalWorkers int
	barrier      chan bool
}

func NewCheckpoint(numWorkers int) *Checkpoint {
	return &Checkpoint{
		totalWorkers: numWorkers,
		barrier:      make(chan bool, numWorkers),
	}
}

func (cp *Checkpoint) Wait(workerID int) {
	cp.mu.Lock()
	cp.arrived++

	fmt.Printf("Worker %d arrived at checkpoint (arrived: %d/%d)\n",
		workerID, cp.arrived, cp.totalWorkers)

	if cp.arrived == cp.totalWorkers {
		fmt.Println("========================================")
		fmt.Println("All workers arrived! Releasing checkpoint...")
		fmt.Println("========================================")

		for i := 0; i < cp.totalWorkers; i++ {
			cp.barrier <- true
		}
	}
	cp.mu.Unlock()

	<-cp.barrier

	cp.mu.Lock()
	cp.arrived--
	cp.mu.Unlock()

	fmt.Printf("Worker %d proceeding to next round\n", workerID)
}

func worker(id int, checkpoint *Checkpoint, rounds int, wg *sync.WaitGroup) {
	defer wg.Done()

	for round := 1; round <= rounds; round++ {
		fmt.Printf("Worker %d: Starting round %d\n", id, round)

		workTime := time.Duration(rand.Intn(1000)+500) * time.Millisecond
		time.Sleep(workTime)

		fmt.Printf("Worker %d: Completed round %d work, going to checkpoint\n", id, round)

		checkpoint.Wait(id)
	}

	fmt.Printf("Worker %d: All rounds completed!\n", id)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	numWorkers := 4
	numRounds := 3

	fmt.Println("╔════════════════════════════════════════╗")
	fmt.Println("║   CHECKPOINT SYNCHRONIZATION PROBLEM   ║")
	fmt.Println("╚════════════════════════════════════════╝")
	fmt.Printf("Number of workers: %d\n", numWorkers)
	fmt.Printf("Number of rounds: %d\n", numRounds)
	fmt.Println()

	checkpoint := NewCheckpoint(numWorkers)
	var wg sync.WaitGroup

	fmt.Println("Starting all workers...")
	fmt.Println()

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, checkpoint, numRounds, &wg)
	}

	wg.Wait()

	fmt.Println()
	fmt.Println("════════════════════════════════════════")
	fmt.Println("All workers completed all rounds!")
	fmt.Println("════════════════════════════════════════")
}
