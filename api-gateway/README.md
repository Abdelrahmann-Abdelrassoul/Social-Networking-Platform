# api-gateway

## Initial behavior

* starts an HTTP server,
* loads config,
* registers `/health`,
* registers placeholder route groups:

  * `/api/v1/auth`
  * `/api/v1/users`
  * `/api/v1/posts`
  * `/api/v1/feed`
  * `/api/v1/notifications`
* applies request ID, logging, and recovery middleware,
* shuts down gracefully.