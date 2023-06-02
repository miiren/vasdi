package job

import (
	"fmt"
)

type helloJob struct{}

func (h *helloJob) Run() {
	fmt.Println("hello in FunCall ~ ")
}
