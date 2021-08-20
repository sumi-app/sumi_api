package sqlstore

import (
	"os"
	"sumi/app/utils"
	"sumi/app/utils/modes"
)

type StoreConfig struct {
	databaseUrl string
}

func (s *StoreConfig) DatabaseUrl() string {
	return s.databaseUrl
}

func LoadConfig(mode string) *StoreConfig {
	utils.NewConfigFactory().Load()
	var databaseUrl string
	if mode == modes.Profile{
		databaseUrl = os.Getenv("TEST_DATABASE_URL")
	} else if mode == modes.Release {
		databaseUrl = os.Getenv("DATABASE_URL")
	}

	if len(databaseUrl) <= 0 {
		panic("Key DATABASE_URL not Found in environment")
	}

	return &StoreConfig{
		databaseUrl: databaseUrl,
	}
}
