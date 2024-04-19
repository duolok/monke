package main

import (
    "fmt"
    "os"
    "os/user"
    "monke/repl"
)

func main() {
    user, err := user.Current()
    if err != nil {
        panic(err)
    }

    fmt.Printf("Hello %s! This is the Monke programming language! ", user.Username)
    fmt.Printf("Feel frere to tye in commands\n")
    repl.Start(os.Stdin, os.Stdout)
}
