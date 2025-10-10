package irpc_tcp_example

import "time"

type StringTool interface {
	Reverse(in string) (string, error)
	Repeat(in string, n int) (string, error)
	TimeToStr(t time.Time) (string, error)
}
