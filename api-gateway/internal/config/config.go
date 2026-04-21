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

    AuthServiceURL         string
    UsersServiceURL        string
    PostsServiceURL        string
    FeedServiceURL         string
    NotificationServiceURL string
    RedisHost              string
    RedisPort              string
    JWTSecret              string

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


    cfg.AuthServiceURL = getEnv("AUTH_SERVICE_URL", "http://localhost:8081")
    cfg.UsersServiceURL = getEnv("USERS_SERVICE_URL", "http://localhost:8082")
    cfg.PostsServiceURL = getEnv("POSTS_SERVICE_URL", "http://localhost:8083")
    cfg.FeedServiceURL = getEnv("FEED_SERVICE_URL", "http://localhost:8084")
    cfg.NotificationServiceURL = getEnv("NOTIFICATION_SERVICE_URL", "http://localhost:8085")
    cfg.RedisHost = getEnv("REDIS_HOST", "localhost")
    cfg.RedisPort = getEnv("REDIS_PORT", "6379")
    cfg.JWTSecret = getEnv("JWT_SECRET", "change-me")

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
