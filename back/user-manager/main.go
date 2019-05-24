package main

import userManagerCmd "github.com/MasterYourLife/back/user-manager/cmd/manager"

var (
	BinaryName    string = "user-manager"
	version       string = "v0.1.0"
	buildDateTime string = "unknown"
)

func main() {
	userManagerCmd.BinaryName = BinaryName
	userManagerCmd.BuildDateTime = buildDateTime
	userManagerCmd.RootCmd.Version = version
	_ = userManagerCmd.RootCmd.Execute()
}
