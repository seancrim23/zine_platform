package server

type Config struct {
	ServerPort   uint16
	GCPProjectId string
	//TODO add any needed config for gc storage
}

func LoadConfig() Config {
	cfg := Config{
		ServerPort:   3000,
		GCPProjectId: "replace-me-with-env",
	}

	return cfg
}
