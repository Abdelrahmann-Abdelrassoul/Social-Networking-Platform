# posts-service

## Initial behavior

* starts service,
* exposes `/health`,
* creates placeholder route group:

  * `POST /api/v1/posts`
  * `GET /api/v1/posts/{id}`
  * `PATCH /api/v1/posts/{id}`
  * `DELETE /api/v1/posts/{id}`