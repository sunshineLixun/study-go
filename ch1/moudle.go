package ch1

import (
	"fmt"

	"github.com/google/uuid"
)

func init() {
	fmt.Println("moddle init")
}

func Log() {
	fmt.Println(uuid.NewString())
}
