package conf

import (
	"os"
	"path/filepath"
)

var (
	CAFile               = configFile("ca.pem")

	ServerCertFile       = configFile("server.pem")
	ServerKeyFile        = configFile("server-key.pem")

	// root access
	RootClientCertFile   = configFile("root-client.pem")
	RootClientKeyFile    = configFile("root-client-key.pem")

	// no access
	NobodyClientCertFile = configFile("nobody-client.pem")
	NobodyClientKeyFile  = configFile("nobody-client-key.pem")

	ACLModelFile = configFile("model.conf")
	ACLPolicyFile = configFile("policy.csv")
)

func configFile(filename string) string {
	if os.Getenv("CONFIG_DIR") != "" {
		return filepath.Join(os.Getenv("CONFIG_DIR"), filename)
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return filepath.Join(homeDir, "dev/.chronicle", filename)
}
