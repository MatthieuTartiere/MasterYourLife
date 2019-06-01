package managerCmd

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"strings"
)

var (
	conf          Flags
	BuildDateTime string
	BinaryName    string
	Header        = `
 █████╗  ██████╗████████╗    ███╗   ███╗ █████╗ ███╗   ██╗ █████╗  ██████╗ ███████╗██████╗ 
██╔══██╗██╔════╝╚══██╔══╝    ████╗ ████║██╔══██╗████╗  ██║██╔══██╗██╔════╝ ██╔════╝██╔══██╗
███████║██║        ██║       ██╔████╔██║███████║██╔██╗ ██║███████║██║  ███╗█████╗  ██████╔╝
██╔══██║██║        ██║       ██║╚██╔╝██║██╔══██║██║╚██╗██║██╔══██║██║   ██║██╔══╝  ██╔══██╗
██║  ██║╚██████╗   ██║██╗    ██║ ╚═╝ ██║██║  ██║██║ ╚████║██║  ██║╚██████╔╝███████╗██║  ██║
╚═╝  ╚═╝ ╚═════╝   ╚═╝╚═╝    ╚═╝     ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝╚═╝  ╚═╝ ╚═════╝ ╚══════╝╚═╝  ╚═╝
`
)

type Flags struct {
	LogFile           *string
	LogLevel          *string
	InfluxServerPort  *int
	InfluxHostname    *string
}

func dumpAppInfo() {
	log.WithFields(log.Fields{
		"log-file":            *conf.LogFile,
		"log-level":           *conf.LogLevel,
		"influx-port":         *conf.InfluxServerPort,
		"influx-hostname":     *conf.InfluxHostname,
	}).Info("app info")
}

var RootCmd = &cobra.Command{
	Use: BinaryName,
	Run: func(cmd *cobra.Command, args []string) {
		//client := influx.Connect(conf.InfluxHostname, conf.InfluxServerPort)
		// call processing here
		//client.Close()
		fmt.Println("fin du job. au-revoir.")
	},
}

func init() {
	conf.LogFile = RootCmd.PersistentFlags().String("log-file", "", "if given logs are output in the given file")
	conf.LogLevel = RootCmd.PersistentFlags().String("log-level", "DEBUG", "log level")
	conf.InfluxServerPort = RootCmd.PersistentFlags().Int("influx-port", 8086, "influx server port")
	conf.InfluxHostname = RootCmd.PersistentFlags().String("influx-hostname", "127.0.0.1", "influx server hostname")

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
