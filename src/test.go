package main

import (
	"fmt"
	"./utils"
	"./filesystem"
)
func main() {
    e,_:= utils.PathExists("/d/")
    fmt.Println(e)
    filesystem.CreateMirrorDisk()
    

}