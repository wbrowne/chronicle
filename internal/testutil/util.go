package testutil

import (
	"os"
	"path/filepath"
)

var (
	CAFile               = configFile("ca.pem")
	ServerCertFile       = configFile("server.pem")
	ServerKeyFile        = configFile("server-key.pem")
	RootClientCertFile   = configFile("root-client.pem")
	RootClientKeyFile    = configFile("root-client-key.pem")
	NobodyClientCertFile = configFile("nobody-client.pem")
	NobodyClientKeyFile  = configFile("nobody-client-key.pem")
	ACLModelFile         = configFile("model.conf")
	ACLPolicyFile        = configFile("policy.csv")
)

func configFile(filename string) string {
	if os.Getenv("CONFIG_PATH") != "" {
		return filepath.Join(os.Getenv("CONFIG_PATH"), filename)
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return filepath.Join(homeDir, "dev", ".chronicle", filename)
}
