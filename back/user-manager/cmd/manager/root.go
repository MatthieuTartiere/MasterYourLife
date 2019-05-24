package userManagerCmd

import (
	"fmt"
	"strings"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	conf          Flags
	BuildDateTime string
	BinaryName    string
	Header        = `
   __  _______ __________     __  ______    _   _____   ________________ 
  / / / / ___// ____/ __ \   /  |/  /   |  / | / /   | / ____/ ____/ __ \
 / / / /\__ \/ __/ / /_/ /  / /|_/ / /| | /  |/ / /| |/ / __/ __/ / /_/ /
/ /_/ /___/ / /___/ _, _/  / /  / / ___ |/ /|  / ___ / /_/ / /___/ _, _/ 
\____//____/_____/_/ |_|  /_/  /_/_/  |_/_/ |_/_/  |_\____/_____/_/ |_|
`
)

type Flags struct {
	LogFile           *string
	LogLevel          *string
}

func dumpAppInfo() {
	log.WithFields(log.Fields{
		"log-file":            *conf.LogFile,
		"log-level":           *conf.LogLevel,

	}).Info("app info")
}

var RootCmd = &cobra.Command{
	Use: BinaryName,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("fin du job. au-revoir.")
	},
}

func init() {
	conf.LogFile = RootCmd.PersistentFlags().String("log-file", "", "if given logs are output in the given file")
	conf.LogLevel = RootCmd.PersistentFlags().String("log-level", "DEBUG", "log level")


	log.SetLevel(log.InfoLevel)
	cobra.OnInitialize(initLogger, dumpAppInfo)
}

// setup logger
func initLogger() {
	fmt.Println(Header)
	log.WithFields(log.Fields{
		"build_date_time": BuildDateTime,
		"version":         RootCmd.Version,
	}).Info(BinaryName)
	logLevel, err := log.ParseLevel(strings.ToLower(*conf.LogLevel))
	if err != nil {
		log.WithError(err).Warn("unknown log level, fallback to info")
		logLevel = log.InfoLevel
	}
	log.SetLevel(logLevel)
}
