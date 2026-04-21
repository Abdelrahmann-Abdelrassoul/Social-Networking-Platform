# user-service

## Initial behavior

* starts service,
* exposes `/health`,
* creates placeholder route group:

  * `GET /api/v1/users/me`
  * `PATCH /api/v1/users/me`
  * `GET /api/v1/users/{id}`
  * `POST /api/v1/users/{id}/follow`
  * `DELETE /api/v1/users/{id}/follow`