package wskide

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
)

// VerboseFlag is flag for verbose
var VerboseFlag = kingpin.Flag("verbose", "Verbose").
	Short('v').Default("false").Bool()

var debugCmd = kingpin.Command("debug", "debug").Hidden()

// Main entrypoint for wskide
func Main() {
	cmd := kingpin.Parse()
	if *VerboseFlag {
		log.SetLevel(log.TraceLevel)
	}
	if !(whiskParse(cmd) || ideParse(cmd) || startParse(cmd)) {
		kingpin.Usage()
	}
}
