package utils

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var (
	once     sync.Once
	instance ConfigFactory
)

type ConfigFactory struct {
	isLoaded bool
}

//Load ..
func (configFactory *ConfigFactory) Load() {
	if !configFactory.isLoaded {
		loadEnv()
		configFactory.isLoaded = true
	}
}

func loadEnv() {
	// Если приложение в продакшен моде -
	// Оно есть в переменных среды и мы его грузим сразу без загрузки .env

	prodEnv := os.Getenv("PRODUCTION")
	isProductionSetBeforeEnvLoad := len(prodEnv) > 0

	if !isProductionSetBeforeEnvLoad {
		err := godotenv.Load(".env")
		if err != nil {
			log.Println("Error loading .env file")
		}
	}
}

//NewConfigFactory ...
func NewConfigFactory() *ConfigFactory {
	once.Do(func() {
		instance = ConfigFactory{
			isLoaded: false,
		}
	})

	return &instance
}
