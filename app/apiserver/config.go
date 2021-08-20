package apiserver

import (
	"sumi/app/store/sqlstore"
	"sumi/app/utils"
	"sumi/app/utils/modes"
)


type Config struct {
	Port        string
	LogLevel    string
	DatabaseUrl string
	Mode     string
	SignInKey   string
	ServerUrl string
}

func LoadConfig() *Config {
	utils.NewConfigFactory().Load()

	mode := modes.Release

	//port := os.Getenv("PORT")
	//if len(port) <= 0 {
	//	port = "8000"
	//}

	//logLevel := os.Getenv("LOG_LEVEL")
	//if len(logLevel) <= 0 {
	//	logLevel = modes.Debug
	//}

	//productionEnv := os.Getenv("PRODUCTION")
	//if len(productionEnv) > 0 {
	//	mode = productionEnv
	//}

	//signInKey := os.Getenv("SIGN_IN_KEY")
	//if len(signInKey) <= 0 {
	//	panic("No SIGN_IN_KEY in environment")
	//}

	//serverUrl := os.Getenv("SERVER_URL")
	//if len(serverUrl) <= 0 {
	//	panic("No SERVER_URL in environment")
	//}

	storeConfig := sqlstore.LoadConfig(mode)

	return &Config{
		//Port:        port,
		//LogLevel:    logLevel,
		DatabaseUrl: storeConfig.DatabaseUrl(),
		Mode:     mode,
		//SignInKey:   signInKey,
		//ServerUrl: serverUrl,
	}
}
