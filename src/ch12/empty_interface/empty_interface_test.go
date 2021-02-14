package empty_interface

import (
	"fmt"
	"testing"
)

func run(p interface{}) {
	if v, ok := p.(int); ok {
		fmt.Printf("Integer:%v\n", v)
	} else if v, ok := p.(string); ok {
		fmt.Printf("String:%v\n", v)
	} else {
		fmt.Printf("Unknow Type\n")
	}

	switch v := p.(type) {
	case int:
		fmt.Printf("Integer:%v\n", v)
	case string:
		fmt.Printf("String:%v\n", v)
	default:
		fmt.Printf("Unknow Type\n")
	}
}

func TestRun(t *testing.T) {
	run(1)
	run("hello world")
	run(1.3)
}
