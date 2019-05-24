package influx

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/influxdata/influxdb1-client/v2"
	log "github.com/sirupsen/logrus"
)

func Connect(hostname *string, port *int) client.Client {
	log.Debug("Try to open connection to Influx")
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: fmt.Sprintf("http://%s:%d", *hostname, *port),
	})
	if err != nil {
		log.Fatalf("Error creating InfluxDB Client: %s", err.Error())
	}
	log.Debug("test connection to Influx")
	_, _, err = c.Ping(1000)
	if err != nil {
		log.Fatalf("error open connection to InfluxDB Client: %s", err.Error())
	}
	return c
}
