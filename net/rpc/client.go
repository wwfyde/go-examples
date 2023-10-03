package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client, _ := rpc.DialHTTP("tcp", ":1234")
	var reply int
	client.Call("Demo.Name", "", &reply)

	fmt.Println(reply)

}
