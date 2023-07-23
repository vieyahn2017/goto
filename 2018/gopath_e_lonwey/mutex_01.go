package main
 
import (
    "fmt"
    "runtime"
    "sync"
)

type Counter struct {
    mu sync.Mutex
    x int64
}

func (c *Counter) Inc() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.x++
    fmt.Println(c.x)
}

func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())
    c := Counter{}
    var wait sync.WaitGroup
    wait.Add(4)
    for k := 4; k > 0; k-- {
        go func() {
            for i := 25; i > 0; i-- {
                c.Inc()
            }
            wait.Done()
        }()
    }
    wait.Wait()
    fmt.Println(c.x)
}