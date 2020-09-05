# Cart service
This is an implementation in Golang of a cart microservice extracted from a monolith
migration. The current version is a **work in progress** and it is not fully
completed.

## Status
### Things that are working.
- `/cart` endpoint.
- [API design](./api/swagger.yml)
- Description of the [architecture](./docs/README.md)
- CORS
- Logging.
- Error recovery.

### Things that are sort of working.
- Basic authentication: Currently, the list of users/passwords are hard coded.
- Database: The persistence layer is currently a SQLite database in memory with
    automatic migrations. It is good enough for development, but not for
    production.

### Things that are not working.
- `/cart/{cartId}/item` endpoint.
- Unit testing.
- Dockerfile.
- General polishing.


## Getting started
### Requirements
- Go 1.15
- SQLite 3.33.0

### How to build & run
```
cd /cmd/cart
go mod download
go build
./cart
```

### Some examples
**Unauthorised access**
```
❯ http GET http://localhost:3000/cart
HTTP/1.1 401 Unauthorized
Access-Control-Allow-Methods: POST,GET,POST,PUT,DELETE,HEAD,OPTIONS,PATCH
Access-Control-Allow-Origin: *
Content-Length: 67
Content-Type: application/json; charset=utf-8

{
    "description": "Unauthorised",
    "error": "Unauthorised",
    "description": "Unauthorised"
}
```

**Create a cart**
```
❯ http POST http://user:user@localhost:3000/cart
HTTP/1.1 201 Created
Access-Control-Allow-Methods: POST,GET,POST,PUT,DELETE,HEAD,OPTIONS,PATCH
Access-Control-Allow-Origin: *
Content-Length: 131
Content-Type: application/json; charset=utf-8

{
    "created_at": "2020-09-05T21:51:38.9030783+02:00",
    "id": 1,
    "items": null,
    "update_at": "2020-09-05T21:51:38.9030783+02:00",
    "user_id": 2
}
```

**Retrieve a cart**
```
❯ http GET http://user:user@localhost:3000/cart/1
HTTP/1.1 200 OK
Access-Control-Allow-Methods: GET,GET,POST,PUT,DELETE,HEAD,OPTIONS,PATCH
Access-Control-Allow-Origin: *
Content-Length: 131
Content-Type: application/json; charset=utf-8

{
    "created_at": "2020-09-05T21:51:38.9030783+02:00",
    "id": 1,
    "items": null,
    "update_at": "2020-09-05T21:51:38.9030783+02:00",
    "user_id": 2
}
```
