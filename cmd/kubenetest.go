package main

import (
	"math/rand"
	"time"

	"github.com/terloo/kubenetest/cmd/app"
)

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))
	cmd := app.NewKubenetestCommand()
	cmd.Execute()
}
