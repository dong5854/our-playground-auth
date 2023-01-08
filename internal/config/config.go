package config

import (
	"log"
	"os"
)

func GetEnv(key string) string {
	rawEnv := os.Getenv(key)
	if len(rawEnv) == 0 {
		log.Println("empty environment :", key)
		return rawEnv
	}
	return rawEnv
}
