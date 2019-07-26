package main

import (
	"fmt"
	"net/http"
	// _ "net/http/pprof"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"counter/counter"
	"counter/fs"
)

var ctr counter.CounterService
var file fs.FileService
var mu sync.Mutex

func handlerIncrement(w http.ResponseWriter, request *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	ctr.Increment()
	fmt.Fprint(w, "counter: ", ctr.Count())
}

func getCounter(w http.ResponseWriter, request *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	fmt.Fprint(w, "counter: ", ctr.Count())
}

func timer() {
	for _ = range time.NewTicker(1 * time.Second).C {
		mu.Lock()
		ctr.Add()
		mu.Unlock()
	}
}

// file "cache.txt" contain how many request each second received during the previous 60 seconds.
func main() {
	// initialize: read and parse file into []int
	file = fs.NewFileService("cache.txt")
	t, err := file.Read()
	if err != nil || len(t) < 60 {
		// default []int
		c := [60]int{}
		t = c[:]
	}

	// create counter with previous state
	ctr = counter.NewCounter(t)

	http.HandleFunc("/", handlerIncrement)

	http.HandleFunc("/count", getCounter)

	// start time tick
	go timer()

	errs := make(chan error, 2)
	go func() {
		errs <- http.ListenAndServe(":8080", nil)
	}()
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	fmt.Println(" terminated: ", <-errs)
	fmt.Println("saving state to file")

	// save state before exit
	mu.Lock()
	_ = file.Save(ctr.Get())
	mu.Unlock()
	fmt.Println("exit")
}
