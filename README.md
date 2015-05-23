# Runner

## Installation

  $ go get github.com/mantyr/runner
  
## Changelog

*    **2015-05-23** : Add `GetCommand` helper function to return message of channel in unlocked mode. Used to control goroutines. Start, stop and other command in return. [Oleg Shevelev][mantyr].

## Examples

```Go
package main

import (
    "github.com/mantyr/runner"
    "fmt"
    "time"
)

func run(name string, is_run bool, safe chan string) {
    var comm string

    fmt.Println(name)
    i := 0

    for {
        comm = runner.GetCommand(safe, is_run)
        if comm == "stop" {
            fmt.Println("stop")
            is_run = false
            continue
        } else if comm == "exit" {
            fmt.Println("full exit")
            return
        }

        fmt.Println("comm: ", comm)
        fmt.Println("i, ", i)
        is_run = true
        i++
    }
}

func main() {
    safe := make(chan string, 10)

    go run("test start/stop goroutine", true, safe)
    safe <- "stop"
    safe <- "start 1"
    safe <- "stop"
    safe <- "start 2"
    safe <- "stop"
    safe <- "stop"
    safe <- "stop"
    safe <- "start 3"
    safe <- "start 4"
    safe <- "stop"
    safe <- "start 5"
    safe <- "stop"

    time.Sleep(time.Second * 2)
    safe <- "exit"

    var input string
    fmt.Scanln(&input)
}
```

## Author

[Oleg Shevelev][mantyr]

[mantyr]: https://github.com/mantyr
