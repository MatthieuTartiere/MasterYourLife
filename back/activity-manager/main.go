package main

import managerCmd "github.com/MasterYourLife/back/activity-manager/cmd/manager"

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
