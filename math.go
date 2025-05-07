package irpc_tcp_example

// Math interface is the only thing needed to define our protocol
// client and server code will be generated for us in math_irpc.go file using the command:
// $ irpc math.go
type Math interface {
	Add(a, b int) (int, error)
}
