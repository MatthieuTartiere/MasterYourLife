package main

import "master-your-life/goals-manager/cmd/manager"


var (
	BinaryName    string = "goals-manager"
	version       string = "v0.1.0"
	buildDateTime string = "unknown"
)

func main() {
	goalsManagerCmd.BinaryName = BinaryName
	goalsManagerCmd.BuildDateTime = buildDateTime
	goalsManagerCmd.RootCmd.Version = version
	_ = goalsManagerCmd.RootCmd.Execute()
}
