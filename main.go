package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	mandibleConf "github.com/Imgur/mandible/config"
	processors "github.com/Imgur/mandible/imageprocessor"
	mandible "github.com/Imgur/mandible/server"
)

func path_exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func main() {
	configFile := os.Getenv("MANDIBLE_CONF")
	jpegMiniLicenseFile := os.Getenv("JPEGMINI_SETTINGS_FILE")

	config := mandibleConf.NewConfiguration(configFile)

	var server *mandible.Server
	var stats mandible.RuntimeStats

	if config.DatadogEnabled {
		var err error
		stats, err = mandible.NewDatadogStats(config.DatadogHostname)
		if err != nil {
			log.Printf("Invalid Datadog Hostname: %s", config.DatadogHostname)
			os.Exit(1)
		}
		log.Println("Stats init success")
	} else {
		stats = &mandible.DiscardStats{}
	}

	if config.JpegMiniEnabled {
		if !(path_exists("/etc/jpegmini/jpegmini.cfg") || path_exists(jpegMiniLicenseFile)) {
			log.Printf("JPEGMini is enabled, but license file does not exist")
			os.Exit(1)
		}
	}

	if os.Getenv("AUTHENTICATION_HMAC_KEY") != "" {
		key := []byte(os.Getenv("AUTHENTICATION_HMAC_KEY"))
		auth := mandible.NewHMACAuthenticatorSHA256(key)
		server = mandible.NewAuthenticatedServer(config, processors.EverythingStrategy, auth, stats)
	} else {
		server = mandible.NewServer(config, processors.EverythingStrategy, stats)
	}

	muxer := http.NewServeMux()
	server.Configure(muxer)

	port := fmt.Sprintf(":%d", server.Config.Port)

	log.Printf("Listening on Port: %s", port)

	stats.LogStartup()
	http.ListenAndServe(port, muxer)
}
