// Golang Runner
// 2015 (c) Oleg Shevelev|mantyr@gmail.com
package runner

import (
    "log"
)

// For start/stop goroutine between iterations, universal version
func GetCommand(safe chan string, is_run bool) string {
    select {
        case comm := <- safe:
            return comm
        default:
            if !is_run {
                comm := <- safe // лочим пока не получим команду
                return comm
            }
            return "next"
    }
}

// return last command from channel or "next" or will wait for the first command
func GetLastCommand(safe chan string, is_run bool) string {
    comm := "next"
    for {
        select {
            case comm = <- safe:
                is_run = true
            default:
                if !is_run {
                    comm = <- safe
                }
                return comm
        }
    }
}


func Recovery(message string, f func()) {
    if r := recover(); r != nil {
        log.Println("Recovered in ", message, r)
        f()
    }
}