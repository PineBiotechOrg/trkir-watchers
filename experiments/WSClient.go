package main

import (
    "fmt"
    "strings"
)


func main() {
    rpiIps := strings.Split(rpiIps, separator)
    mouseIds := strings.Split(mouseIds, separator)

    for ind, rpiIp := range rpiIps {
        mouseId := mouseIds[ind]
        fmt.Printf("Connecting to %s\n", rpiIp)
        readRpi(rpiIp, mouseId)
        fmt.Printf("Done\n")
    }
}
