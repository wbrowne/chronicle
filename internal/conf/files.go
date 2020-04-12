package conf

import (
	"os"
	"path/filepath"
)

var (
	CAFile             = configFile("ca.pem")
	ServerCertFile     = configFile("server.pem")
	ServerKeyFile      = configFile("server-key.pem")
	RootClientCertFile = configFile("root-client.pem")
	RootClientKeyFile  = configFile("root-client-key.pem")
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
