package config

import (
	"flag"
	"os"

	"github.com/lunny/log"
)

// LogConfig is
func LogConfig() *os.File {
	f, err := os.OpenFile("api.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		log.Errorf("cannot open 'api.log', (%s)", err.Error())
		flag.Usage()
		os.Exit(-1)
	}
	defer f.Close()
	return f
}
