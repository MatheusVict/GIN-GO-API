# Gin user authentication API

This an API Golang with Gin, that makes jwt authentication and user crud using facade and SOLID pattern

### Technologies used

Exemplo:
* [Golang](https://github.com/golang/go)
* [Gin](https://www.google.com/url?sa=t&rct=j&q=&esrc=s&source=web&cd=&cad=rja&uact=8&ved=2ahUKEwjL_Pn72oyCAxV7uZUCHWuxC24QFnoECAsQAQ&url=https%3A%2F%2Fgin-gonic.com%2F&usg=AOvVaw3NpIRVflxdDy9iVoKMyoRA&opi=89978449)
* [JWT](https://jwt.io/)
* [MongoDB](https://www.google.com/url?sa=t&rct=j&q=&esrc=s&source=web&cd=&cad=rja&uact=8&ved=2ahUKEwixiJ7H2IyCAxW1NTUKHVcUDLgQFnoECAYQAQ&url=https%3A%2F%2Fwww.mongodb.com%2F&usg=AOvVaw2ODprZpw7B6asXJGVwvA-G&opi=89978449)
* [Gomock-uber](https://github.com/uber-go/mock)

## Necessary Dependencies and Versions

* Go - Versão: 1.20

### Getting starter

To run at golang instace use


```
go get
```
Then

```
go run *.go
```

To run with docker user

```
docker build -t gin_api .
```
and

```
docker run -d gin_api
```

It'll return:
```bash
[GIN-debug] Listening and serving HTTP on :8080
```
After that, you should take a look on [localhos:8080](http://localhost:8080)

## How to run the tests

```
go test -v ./...
```

## Http routes

### Register:
> POST

```/createUser```

Body:

```json
{
  "email": "user@email.com",
  "password": "123456@!",
  "name": "user 1",
  "age": 10
}
```
Return:

```json
{
  "id": "6534242206c1e14b2bbb73a8",
  "email": "user@email.com",
  "password": "123456@!",
  "name": "user 1",
  "age": 10
}
```

### Login:
> POST

```/login```

Body:

```json
{
  "email": "user@email.com",
  "password": "123456@!"
}
```
Return:
```curl
HTTP/1.1 200 OK
Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZ2UiOjEwLCJlbWFpbCI6ImphZnNkbDJlQG1kc2YuY20iLCJleHAiOjE2OTgxNzcxNzEsImlkIjoiNjUzNDI0MjIwNmMxZTE0YjJiYmI3M2E4IiwibmFtZSI6ImFhYWEifQ.I1-CNGfiiPn1SzEHFfxYpBBBNvf7gvdq5a8DbKoqRQM
Content-Type: application/json; charset=utf-8
Date: Mon, 23 Oct 2023 19:52:51 GMT
Content-Length: 83

{
  "id":"6534242206c1e14b2bbb73a8",
  "email":"jafsdl2e@mdsf.cm",
  "name":"aaaa","age":10
}
```

#### Every route down here needs a Header:
```
Authorization: Bearer your_token_here
```

### Get an user by id:
> GET

```/getUserById/{userId}```

Return:

```json
{
  "id": "6534242206c1e14b2bbb73a8",
  "email": "user@email.com",
  "password": "123456@!",
  "name": "user 1",
  "age": 10
}
```

### Get an user by email:
> GET

```/getUserByEmail/{userEmail}```

Return:

```json
{
  "id": "6534242206c1e14b2bbb73a8",
  "email": "user@email.com",
  "password": "123456@!",
  "name": "user 1",
  "age": 10
}
```


### Update user by id:
> POST

```/updateUser/{userId}```

Body:

```json
{
  "name": "user 1",
  "age": 10
}
```
Return:

```
status 200 OK
```

### Delete an user by id:
> GET

```/deleteUser/{userId}```

Return:

```
status 200 OK
```

## Problems faced

### Problem 1:
Error to connect to mongoDB
* Check your ```.env``` file it must be at root

### Problem 2:
Project does not run
* Check if you used command ```go get``` before of all

