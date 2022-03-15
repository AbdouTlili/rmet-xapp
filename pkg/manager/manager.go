package manager

import (
	"fmt"

	"github.com/onosproject/onos-lib-go/pkg/logging"
)

var log = logging.GetLogger()

// Config is a manager configuration
type Config struct {
	CAPath      string
	KeyPath     string
	CertPath    string
	E2tEndpoint string
	GRPCPort    int
	RicActionID int32
	ConfigPath  string
	SMName      string
	SMVersion   string
}

func NewManager(config Config) error {
	fmt.Printf("%#v", config)
	log.Debugf("%#v", config)
	return nil
}

func (m Manager) Run() {

}

// Manager is an abstract struct for manager
type Manager struct {
	appConfig        Config
	config           Config
	measurementStore string
	subManager       string
}
