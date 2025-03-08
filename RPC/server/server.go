package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Calculator struct{}

type CArgs struct {
	A, B int
}

func (c *Calculator) Multiply(args *CArgs, reply *int) error {
	if args == nil {
		return errors.New("arguments cannot be nil")
	}
	*reply = args.A * args.B
	return nil
}

func (c *Calculator) Divide(args *CArgs, reply *int) error {
	if args.B == 0 {
		return errors.New("division by zero")
	}
	*reply = args.A / args.B
	return nil
}

type Greeter struct{}

type GArgs struct {
	H, B string
}

func (c *Greeter) SayHello(args *GArgs, reply *string) error {
	*reply = fmt.Sprintf("Hello, %s!", args.H)
	return nil
}

func (c *Greeter) SayBye(args *GArgs, reply *string) error {
	*reply = fmt.Sprintf("Bye, %s!", args.B)
	return nil
}

func main() {
	calc := new(Calculator)
	greet := new(Greeter)

	rpc.Register(calc)  // Register the Calculator service
	rpc.Register(greet) // Register the Greeter service

	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 8080...")

	for {
		fmt.Println("Waiting for connection...")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Connection error:", err)
			continue
		}
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn)) // Corrected JSON-RPC codec
	}
}
