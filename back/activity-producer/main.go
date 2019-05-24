package main

import producerCmd "github.com/MasterYourLife/back/activity-producer/cmd/producer"

var (
	BinaryName    string = "activity-producer"
	version       string = "v0.1.0"
	buildDateTime string = "unknown"
)

func main() {
	producerCmd.BinaryName = BinaryName
	producerCmd.BuildDateTime = buildDateTime
	producerCmd.RootCmd.Version = version
	_ = producerCmd.RootCmd.Execute()
}
