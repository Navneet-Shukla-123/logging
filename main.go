package main

import (
	"logging/utils"
)

func main() {
	utils.InitializeLogger()
	utils.Logger.Info("Hello world")
	utils.Logger.Error("Not able to reach blog.")
}
