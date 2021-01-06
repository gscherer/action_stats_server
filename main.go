package main

import (
    "fmt"
    "flag"
    "github.com/gscherer/action_stats/action_stats"
)

func main() {
    portPtr := flag.Int("port", 8080, "The port on which to start the server")
    flag.Parse()
    action_stats.StartServer(fmt.Sprintf(":%d", *portPtr))
}


