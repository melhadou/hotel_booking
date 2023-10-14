# Hotel reservation backend

## Project outline
- users -> book room from an hotel 
- admins -> going to check reservation/bookings 
- Authentication and authorization -> JWT tokens
- Hotels -> CRUD API -> JSON
- Rooms -> CRUD API -> JSON
- Scripts -> database management -> seeding, migration

## Resources
Installing mongodb client


```go
go get go.mongodb.org/mongo-driver/mongo
```
### gofiber 
- Installing Fiber:
```go
go get github.com/gofiber/fiber/v2
```
- Documentation:
```
https://gofiber.io
```

## Docker
### Installing mongodb as a Docker container
```bash
docker run --name mongodb -d mongo:latest -p 27017:27017
```
