package runner

/**
    start/stop goroutine between iterations, universal version
    (c) Oleg Shevelev|mantyr@gmail.com
 */
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
