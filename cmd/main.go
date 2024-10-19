package main

import (
	"github.com/komarov-a-kubsu/GoLab5/internal/worker"
)

func main() {
	worker.RunWorkers(3)
}
