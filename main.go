package main

import (
	"flag"
	"fmt"

	"github.com/AbdouTlili/rmet-xapp/pkg/manager"
	"github.com/onosproject/onos-lib-go/pkg/certs"
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

var log = logging.GetLogger()

func main() {
	caPath := flag.String("caPath", "", "path to CA certificate")
	keyPath := flag.String("keyPath", "", "path to client private key")
	certPath := flag.String("certPath", "", "path to client certificate")
	e2tEndpoint := flag.String("e2tEndpoint", "onos-e2t:5150", "E2T service endpoint")
	ricActionID := flag.Int("ricActionID", 10, "RIC Action ID in E2 message")
	configPath := flag.String("configPath", "/etc/onos/config/config.json", "path to config.json file")
	grpcPort := flag.Int("grpcPort", 5150, "grpc Port number")
	smName := flag.String("smName", "e2sm-rmet", "Service model name in RAN function description")
	smVersion := flag.String("smVersion", "v1", "Service model version in RAN function description")

	ready := make(chan bool)

	flag.Parse()

	_, err := certs.HandleCertPaths(*caPath, *keyPath, *certPath, true)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Starting RMET xApp")
	cfg := manager.Config{
		CAPath:      *caPath,
		KeyPath:     *keyPath,
		CertPath:    *certPath,
		E2tEndpoint: *e2tEndpoint,
		GRPCPort:    *grpcPort,
		RicActionID: int32(*ricActionID),
		ConfigPath:  *configPath,
		SMName:      *smName,
		SMVersion:   *smVersion,
	}

	fmt.Println("rmet started !")

	mgr := manager.NewManager(cfg)
	log.Info("New manager created")
	mgr.Run()
	<-ready
}
