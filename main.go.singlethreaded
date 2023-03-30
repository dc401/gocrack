package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
  "strconv"
  "time"
)

/* requires SHA2 with 256 bytes 64 rounds
https://gchq.github.io/CyberChef/#recipe=SHA2('256',64,160)&input=Zm9vdGJhbGw
echo football | sha256sum  produces -> 205b60ee79914af6a09b897170b522c5e16366214b9a0735b4eb550f4b14a3c8

TL;DR - Basically, SHA1-256 mode vs. SHA2-256 mode
*/
var checkHash string = "6382deaf1f5dc6e792b76db4a4a7bf2ba468884e000b25e7928e621e27fb23cb"

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
	for scanner.Scan() {
		line := scanner.Text()
    counterInt++
    //you can use := vs define var foo type format inside a function struct
		hashedEntry := hashString(line)
		//fmt.Println("Entry: " + line +" : " + hashedEntry )
		if hashedEntry == checkHash {
      foundLineNum := strconv.Itoa(counterInt)
			fmt.Println("Found Password on line " +"("+ foundLineNum +"): " +line)
      //exits iteration subroutine back to main
			break 
		}

	}
  //print out the nanoseconds spent
  duration := time.Since(startTime)
  fmt.Println("Runtime Nanoseconds: ")
  fmt.Println(duration.Nanoseconds())
}
