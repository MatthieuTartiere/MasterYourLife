package main

import "master-your-life/activity-manager/cmd/manager"

var (
	BinaryName    string = "activity-manager"
	version       string = "v0.1.0"
	buildDateTime string = "unknown"
)

func main() {
	managerCmd.BinaryName = BinaryName
	managerCmd.BuildDateTime = buildDateTime
	managerCmd.RootCmd.Version = version
	_ = managerCmd.RootCmd.Execute()
}
