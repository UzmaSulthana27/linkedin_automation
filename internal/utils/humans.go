package utils

import (
	"fmt"
	"time"
)

func HumanDelay() {
	fmt.Println("Simulating human behavior...")
	time.Sleep(2 * time.Second)
}
