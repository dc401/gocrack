# gocrack
Example multi-threaded and single-threaded GoLang ability to crack SHA2-256 based hashes using wordlists
The example files are suffixed with main.go.multithreaded and main.go.singlethreaded accordingly. Remain one at a time and run as main.go or build it separately. 

## Usage
**Hash a password or pass a  SHA2-256 mode hash and modify the variable accordingly. For example, in the rockyou.txt wordlist:**

    /*
    
    requires SHA2 with 256 bytes 64 rounds
    https://gchq.github.io/CyberChef/#recipe=SHA2('256',64,160)&input=Zm9vdGJhbGw
    echo football | sha256sum produces -> 205b60ee79914af6a09b897170b522c5e16366214b9a0735b4eb550f4b14a3c8
    
      
    
    Use a password further down the list to compare sequential vs. concurrent job scheduler
    https://gchq.github.io/CyberChef/#recipe=SHA2('256',64,160)&input=SlVTVElOMTk5Mw
    JUSTIN1993 : 96306567b5683a463820ce53dd484eccc57c53aaa3c6d5f74018da1b8ef99815
    
    */
    
    var  checkHash  string = "96306567b5683a463820ce53dd484eccc57c53aaa3c6d5f74018da1b8ef99815" // EDIT THIS LINE

**Ensure you set the correct file name for the handler in the code:**

    //open file and defer until main returns
    
    fileHandle, err := os.Open("rockyou-full.txt") //EDIT THIS LINE
    
    if err != nil {
    
    panic(err)
    
    }
    
    defer fileHandle.Close()

**Runtime**

    go run ./main.go 
    go build gocrack-multithreaded ./main.go
   
## Performance
Based on my original single threaded code. I had ChatGPT-4 re-factor to include the use of concurrency (async) to include waitgroups, the go scheduler subroutines, and channels to pass the details back. 

The results regardless of file size on a MacOS 13.2 M1 Pro 2022 year make was a almost double main function exit runtime because of the channel spawn thread and wait queue overhead on a full unsorted rockyou.txt original list. 

This performance could be different if there were other more complex functions used, such as transposition, masking, and salting concatenation. 

In almost every single use of a dictionary iterated list only method of cracking a hash; the single thread was faster performing when used on the same system between a 77KB file vs. a 138 MB file with many more lines.

There's some great resources to learn about the concurrent processing in Go using the scheduler in the following resources:
 - https://www.ardanlabs.com/blog/2015/02/scheduler-tracing-in-go.html
 - https://www.developer.com/languages/go-scheduler/
 - https://levelup.gitconnected.com/how-does-golang-channel-works-6d66acd54753
 - https://www.practical-go-lessons.com/chap-30-concurrency


![enter image description here](https://raw.githubusercontent.com/dc401/gocrack/main/singlethread-vs-multithread-gocrack-perf.png)

## Concurrency with Channeling in Go
![enter image description here](https://raw.githubusercontent.com/dc401/gocrack/main/go-channel-scheduler-example.png)

## Disclaimer
No expressed or implicit warranty or liability of any kind. Run and re-factor at your own testing expense. 

This was just an quick coding interest in learning GoLang over Python for me.

## Author
Dennis Chow dchow[AT]xtecsystems.com March 29, 2023 - Single Threaded Example , Enhanced by Chat GPT-4 for the re-factored Multi-threaded example
