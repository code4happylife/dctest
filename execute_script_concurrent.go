// You can edit this code!
// Click here and start typing.
package main

import (
        "fmt"
        "os/exec"
        "strconv"
        "sync"
)

func main() {
        wg := &sync.WaitGroup{}
        for i := 1; i <= 1000; i++ {
                wg.Add(1)
                test_process := "/data/multi_test/hehe" + strconv.Itoa(i) + " &"
                fmt.Printf("%s\n", test_process)
                go func() {
                   cmd :=exec.Command("sh","-c",test_process)
                   cmd.Run()
                   wg.Done()
                }()
        }

        wg.Wait()
}
