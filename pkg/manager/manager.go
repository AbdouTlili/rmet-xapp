package manager

import (
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

func NewManager(config Config) *Manager {
	//fmt.Printf("%#v", config)
	log.Infof("%#v", config)
	return &Manager{
		appConfig:        config,
		config:           config,
		measurementStore: "string",
		subManager:       "store",
	}
}

func (m Manager) Run() error {
	return nil
}

// Manager is an abstract struct for manager
type Manager struct {
	appConfig        Config
	config           Config
	measurementStore string
	subManager       string
}
