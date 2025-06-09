package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found, using system environment variables")
	}
}

func GetEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func GetIPLimits() map[string]string {
	ipLimitsStr := GetEnv("IP_LIMITS", "")
	ipLimits := make(map[string]string)

	if ipLimitsStr == "" {
		return ipLimits
	}

	pairs := strings.Split(ipLimitsStr, ",")
	for _, p := range pairs {
		kv := strings.Split(p, "=")
		if len(kv) == 2 {
			ipLimits[strings.TrimSpace(kv[0])] = strings.TrimSpace(kv[1])
		}
	}

	return ipLimits
}
