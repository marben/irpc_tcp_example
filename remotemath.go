package irpc_tcp_example

// RemoteMath interface is the only thing needed to define our protocol
// client and server code will be generated for us in remotemath_irpc.go file using the command:
// $ irpc remotemath.go
type RemoteMath interface {
	Add(a int, b int) (int, error)
}
