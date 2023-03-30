package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

/*
	requires SHA2 with 256 bytes 64 rounds

https://gchq.github.io/CyberChef/#recipe=SHA2('256',64,160)&input=Zm9vdGJhbGw
echo football | sha256sum  produces -> 205b60ee79914af6a09b897170b522c5e16366214b9a0735b4eb550f4b14a3c8

Use a password further down the list to compare sequential vs. concurrent job scheduler
https://gchq.github.io/CyberChef/#recipe=SHA2('256',64,160)&input=SlVTVElOMTk5Mw
JUSTIN1993 : 96306567b5683a463820ce53dd484eccc57c53aaa3c6d5f74018da1b8ef99815
*/

var checkHash string = "96306567b5683a463820ce53dd484eccc57c53aaa3c6d5f74018da1b8ef99815"

/*
wrap the hash creating objects into a function
function definitions require return at the top
*/
func hashString(inputString string) string {
	hashObject := sha256.New()
	hashObject.Write([]byte(inputString))
	calculatedHash := hashObject.Sum(nil)
	return hex.EncodeToString(calculatedHash)
}

func main() {
	/*
		https://yourbasic.org/golang/measure-execution-time/
		Found this sweet time measurement benchmark
	*/
	startTime := time.Now()

	fmt.Println("Trying Hash: " + checkHash)

	//open file and defer until main returns
	fileHandle, err := os.Open("rockyou-full.txt")
	if err != nil {
		panic(err)
	}
	defer fileHandle.Close()

	//set a counter just to see which line in your list works
	counterInt := 0

	//create scanner object and iterate foo
	scanner := bufio.NewScanner(fileHandle)
	var wg sync.WaitGroup
	//use the number of CPU cores available
	numCores := runtime.NumCPU()
	runtime.GOMAXPROCS(numCores)
	ch := make(chan string)

	for i := 0; i < numCores; i++ {
		wg.Add(1)
		go func() {
			for line := range ch {
				hashedEntry := hashString(line)
				if hashedEntry == checkHash {
					foundLineNum := strconv.Itoa(counterInt)
					fmt.Println("Found Password on line " + "(" + foundLineNum + "): " + line)
					break
				}
			}
			wg.Done()
		}()
	}

	for scanner.Scan() {
		counterInt++
		line := scanner.Text()
		ch <- line
	}

	close(ch)
	wg.Wait()

	//print out the nanoseconds spent
	duration := time.Since(startTime)
	fmt.Println("Runtime Nanoseconds: ")
	fmt.Println(duration.Nanoseconds())
}
