You will do for each service:

1. open terminal in service folder
2. set env vars
3. run the service
4. open another terminal
5. call `/health`
6. call one placeholder route
7. stop service
8. run one negative test with bad config

Use **PowerShell** since you are on Windows.

---

# Before testing anything

## 1. Make sure Go works

Run:

```powershell
go version
```

## 2. Go to project root

## 3. Use two terminals

For each service:

* **Terminal 1** = run the service
* **Terminal 2** = send requests with `curl` or `Invoke-WebRequest`

---

# 1) API Gateway manual test

## Terminal 1 — run the service

```powershell
cd .\api-gateway

$env:SERVICE_NAME="api-gateway"
$env:APP_ENV="development"
$env:PORT="8080"
$env:LOG_LEVEL="info"
$env:HTTP_READ_TIMEOUT="10"
$env:HTTP_WRITE_TIMEOUT="10"
$env:HTTP_IDLE_TIMEOUT="60"
$env:AUTH_SERVICE_URL="http://localhost:8081"
$env:USERS_SERVICE_URL="http://localhost:8082"
$env:POSTS_SERVICE_URL="http://localhost:8083"
$env:FEED_SERVICE_URL="http://localhost:8084"
$env:NOTIFICATION_SERVICE_URL="http://localhost:8085"
$env:REDIS_HOST="localhost"
$env:REDIS_PORT="6379"
$env:JWT_SECRET="change-me"

go run ./cmd/server
```

## Expected result

You should see something like:

```text
starting api-gateway on port 8080
```

---

## Terminal 2 — health check

```powershell
curl http://localhost:8080/health
```

Expected:

```json
{"status":"ok","service":"api-gateway"}
```

---

## Terminal 2 — placeholder route test

```powershell
curl http://localhost:8080/api/v1/feed
```

Expected:

* status should be `501 Not Implemented`

If you want to inspect full response:

```powershell
try {
    Invoke-WebRequest http://localhost:8080/api/v1/feed -UseBasicParsing
} catch {
    $_.Exception.Response.StatusCode.value__
}
```

Expected:

* `501`

---

## Negative test

Go back to **Terminal 1**, stop service with:

```powershell
Ctrl + C
```

Then run:

```powershell
$env:PORT=""
go run ./cmd/server
```

Expected:

* service fails immediately

---

# 2) Auth Service manual test

## Terminal 1

```powershell
cd ..\auth-service

$env:SERVICE_NAME="auth-service"
$env:APP_ENV="development"
$env:PORT="8081"
$env:LOG_LEVEL="info"
$env:HTTP_READ_TIMEOUT="10"
$env:HTTP_WRITE_TIMEOUT="10"
$env:HTTP_IDLE_TIMEOUT="60"
$env:GOOGLE_CLIENT_ID="dummy-google-client-id"
$env:GOOGLE_CLIENT_SECRET="dummy-google-client-secret"
$env:GOOGLE_REDIRECT_URL="http://localhost:8081/api/v1/auth/callback"
$env:JWT_SECRET="change-me"
$env:JWT_EXPIRES_IN="24h"
$env:SESSION_TTL="24h"
$env:REDIS_HOST="localhost"
$env:REDIS_PORT="6379"

go run ./cmd/server
```

## Expected result

```text
starting auth-service on port 8081
```

---

## Terminal 2 — health check

```powershell
curl http://localhost:8081/health
```

Expected:

```json
{"status":"ok","service":"auth-service"}
```

---

## Terminal 2 — placeholder route

```powershell
curl http://localhost:8081/api/v1/auth/login
```

Expected:

* `501 Not Implemented`

Or:

```powershell
try {
    Invoke-WebRequest http://localhost:8081/api/v1/auth/login -UseBasicParsing
} catch {
    $_.Exception.Response.StatusCode.value__
}
```

Expected:

* `501`

---

## Negative test

Stop with `Ctrl + C`, then:

```powershell
$env:PORT=""
go run ./cmd/server
```

Expected:

* service fails

---

# 3) Users Service manual test

## Terminal 1

```powershell
cd ..\users-service

$env:SERVICE_NAME="users-service"
$env:APP_ENV="development"
$env:PORT="8082"
$env:LOG_LEVEL="info"
$env:HTTP_READ_TIMEOUT="10"
$env:HTTP_WRITE_TIMEOUT="10"
$env:HTTP_IDLE_TIMEOUT="60"
$env:DB_HOST="localhost"
$env:DB_PORT="5432"
$env:DB_NAME="users_db"
$env:DB_USER="postgres"
$env:DB_PASSWORD="postgres"
$env:DB_SSLMODE="disable"
$env:KAFKA_BROKERS="localhost:9092"
$env:KAFKA_TOPIC_USER_FOLLOWED="user.followed"

go run ./cmd/server
```

## Expected result

```text
starting users-service on port 8082
```

---

## Terminal 2 — health check

```powershell
curl http://localhost:8082/health
```

Expected:

```json
{"status":"ok","service":"users-service"}
```

---

## Terminal 2 — placeholder route 1

```powershell
curl http://localhost:8082/api/v1/users/me
```

Expected:

* `501`

## Terminal 2 — placeholder route 2

```powershell
curl http://localhost:8082/api/v1/users/123
```

Expected:

* `501`

## Terminal 2 — placeholder route 3

```powershell
try {
    Invoke-WebRequest http://localhost:8082/api/v1/users/123/follow -Method POST -UseBasicParsing
} catch {
    $_.Exception.Response.StatusCode.value__
}
```

Expected:

* `501`

---

## Negative test

Stop with `Ctrl + C`, then:

```powershell
$env:PORT=""
go run ./cmd/server
```

Expected:

* fail fast

---

# 4) Posts Service manual test

## Terminal 1

```powershell
cd ..\posts-service

$env:SERVICE_NAME="posts-service"
$env:APP_ENV="development"
$env:PORT="8083"
$env:LOG_LEVEL="info"
$env:HTTP_READ_TIMEOUT="10"
$env:HTTP_WRITE_TIMEOUT="10"
$env:HTTP_IDLE_TIMEOUT="60"
$env:DB_HOST="localhost"
$env:DB_PORT="5432"
$env:DB_NAME="posts_db"
$env:DB_USER="postgres"
$env:DB_PASSWORD="postgres"
$env:DB_SSLMODE="disable"
$env:KAFKA_BROKERS="localhost:9092"
$env:KAFKA_TOPIC_POST_CREATED="post.created"

go run ./cmd/server
```

## Expected result

```text
starting posts-service on port 8083
```

---

## Terminal 2 — health check

```powershell
curl http://localhost:8083/health
```

Expected:

```json
{"status":"ok","service":"posts-service"}
```

---

## Terminal 2 — placeholder routes

### Get post

```powershell
curl http://localhost:8083/api/v1/posts/1
```

Expected:

* `501`

### Create post

```powershell
try {
    Invoke-WebRequest http://localhost:8083/api/v1/posts -Method POST -UseBasicParsing
} catch {
    $_.Exception.Response.StatusCode.value__
}
```

Expected:

* `501`

### Update post

```powershell
try {
    Invoke-WebRequest http://localhost:8083/api/v1/posts/1 -Method PATCH -UseBasicParsing
} catch {
    $_.Exception.Response.StatusCode.value__
}
```

Expected:

* `501`

### Delete post

```powershell
try {
    Invoke-WebRequest http://localhost:8083/api/v1/posts/1 -Method DELETE -UseBasicParsing
} catch {
    $_.Exception.Response.StatusCode.value__
}
```

Expected:

* `501`

---

## Negative test

Stop service, then:

```powershell
$env:PORT=""
go run ./cmd/server
```

Expected:

* fail fast

---

# 5) Feed Service manual test

## Terminal 1

```powershell
cd ..\feed-service

$env:SERVICE_NAME="feed-service"
$env:APP_ENV="development"
$env:PORT="8084"
$env:LOG_LEVEL="info"
$env:HTTP_READ_TIMEOUT="10"
$env:HTTP_WRITE_TIMEOUT="10"
$env:HTTP_IDLE_TIMEOUT="60"
$env:REDIS_HOST="localhost"
$env:REDIS_PORT="6379"
$env:KAFKA_BROKERS="localhost:9092"
$env:KAFKA_TOPIC_POST_CREATED="post.created"
$env:KAFKA_TOPIC_USER_FOLLOWED="user.followed"

go run ./cmd/server
```

## Expected result

```text
starting feed-service on port 8084
```

---

## Terminal 2 — health check

```powershell
curl http://localhost:8084/health
```

Expected:

```json
{"status":"ok","service":"feed-service"}
```

---

## Terminal 2 — placeholder route

```powershell
curl http://localhost:8084/api/v1/feed
```

Expected:

* `501`

Or:

```powershell
try {
    Invoke-WebRequest http://localhost:8084/api/v1/feed -UseBasicParsing
} catch {
    $_.Exception.Response.StatusCode.value__
}
```

Expected:

* `501`

---

## Negative test

Stop service, then:

```powershell
$env:PORT=""
go run ./cmd/server
```

Expected:

* fail fast

---

# 6) Notification Service manual test

## Terminal 1

```powershell
cd ..\notification-service

$env:SERVICE_NAME="notification-service"
$env:APP_ENV="development"
$env:PORT="8085"
$env:LOG_LEVEL="info"
$env:HTTP_READ_TIMEOUT="10"
$env:HTTP_WRITE_TIMEOUT="10"
$env:HTTP_IDLE_TIMEOUT="60"
$env:DB_HOST="localhost"
$env:DB_PORT="5432"
$env:DB_NAME="notifications_db"
$env:DB_USER="postgres"
$env:DB_PASSWORD="postgres"
$env:DB_SSLMODE="disable"
$env:KAFKA_BROKERS="localhost:9092"
$env:KAFKA_TOPIC_USER_FOLLOWED="user.followed"
$env:KAFKA_TOPIC_POST_INTERACTED="post.interacted"

go run ./cmd/server
```

## Expected result

```text
starting notification-service on port 8085
```

---

## Terminal 2 — health check

```powershell
curl http://localhost:8085/health
```

Expected:

```json
{"status":"ok","service":"notification-service"}
```

---

## Terminal 2 — placeholder route

```powershell
curl http://localhost:8085/api/v1/notifications
```

Expected:

* `501`

Or:

```powershell
try {
    Invoke-WebRequest http://localhost:8085/api/v1/notifications -UseBasicParsing
} catch {
    $_.Exception.Response.StatusCode.value__
}
```

Expected:

* `501`

---

## Negative test

Stop service, then:

```powershell
$env:PORT=""
go run ./cmd/server
```

Expected:

* fail fast

---