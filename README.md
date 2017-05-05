# API Server


Simple Rest API using gin(framework) & gorm(orm)


## Build

    go get ./...
    go build -o bin/server


## Run

    PORT=9000 bin/server

## Endpoint list

### Friends Resource

```
GET    /friends
GET    /friends/:id
POST   /friends
PUT    /friends/:id
DELETE /friends/:id
```

### Statuses Resource

```
GET    /statuses
GET    /statuses/:id
POST   /statuses
PUT    /statuses/:id
DELETE /statuses/:id
```

### Users Resource

```
GET    /users
GET    /users/:id
POST   /users
PUT    /users/:id
DELETE /users/:id
```

server runs at http://localhost:8080
