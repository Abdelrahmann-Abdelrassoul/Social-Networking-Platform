package config

import (
    "log"
    "os"
    "strconv"
    "strings"
    "time"
)

type HTTPConfig struct {
    ReadTimeout  time.Duration
    WriteTimeout time.Duration
    IdleTimeout  time.Duration
}

type Config struct {
    ServiceName string
    AppEnv      string
    Port        string
    LogLevel    string
    HTTP        HTTPConfig

    GoogleClientID     string
    GoogleClientSecret string
    GoogleRedirectURL  string
    JWTSecret          string
    JWTExpiresIn       string
    SessionTTL         string
    RedisHost          string
    RedisPort          string

}

func Load() Config {
    cfg := Config{
        ServiceName: getEnv("SERVICE_NAME", "service"),
        AppEnv:      getEnv("APP_ENV", "development"),
        Port:        getEnv("PORT", "8080"),
        LogLevel:    getEnv("LOG_LEVEL", "info"),
        HTTP: HTTPConfig{
            ReadTimeout:  getDurationSeconds("HTTP_READ_TIMEOUT", 10),
            WriteTimeout: getDurationSeconds("HTTP_WRITE_TIMEOUT", 10),
            IdleTimeout:  getDurationSeconds("HTTP_IDLE_TIMEOUT", 60),
        },
    }


    cfg.GoogleClientID = getEnv("GOOGLE_CLIENT_ID", "")
    cfg.GoogleClientSecret = getEnv("GOOGLE_CLIENT_SECRET", "")
    cfg.GoogleRedirectURL = getEnv("GOOGLE_REDIRECT_URL", "")
    cfg.JWTSecret = getEnv("JWT_SECRET", "change-me")
    cfg.JWTExpiresIn = getEnv("JWT_EXPIRES_IN", "24h")
    cfg.SessionTTL = getEnv("SESSION_TTL", "24h")
    cfg.RedisHost = getEnv("REDIS_HOST", "localhost")
    cfg.RedisPort = getEnv("REDIS_PORT", "6379")

    validate(cfg)
    return cfg
}

func validate(cfg Config) {
    if strings.TrimSpace(cfg.Port) == "" {
        log.Fatal("PORT is required")
    }
}

func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok && strings.TrimSpace(value) != "" {
        return value
    }
    return fallback
}

func getDurationSeconds(key string, fallback int) time.Duration {
    raw := getEnv(key, strconv.Itoa(fallback))
    v, err := strconv.Atoi(raw)
    if err != nil {
        return time.Duration(fallback) * time.Second
    }
    return time.Duration(v) * time.Second
}
