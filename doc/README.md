# How run app ?

For run this program you can use docker or normally run go program but before run program you can create a .env file to be like .evn.example.

If you want run without .env file use CLI help command.

```text
Usage of ./enpass <command>:
  -S string     set your ip:port (default "0.0.0.0:1111")
  -d            run debug mode
  -dsn string   database path (default "./db.sqlite")
  -lf uint      life time (default 10)
  -sk string    set secret key
```

Run with docker
> docker compose up

Run with Make
> make run

---

# Api Document

For Better request you can use postman file on postman folder.

## No Need Auth Endpoint

### Generate Random Password

Endpoints: ( GET )
> 0.0.0.0:1111/genRandomPass?size=< int >

Generate random password, if send none size or bad format size automatic Generate random password with length 10.

Response:

```json
{
    "password": "[NTEdJIJrQ"
}
```

### Login

Endpoints: ( POST )
> 0.0.0.0:1111/api/login

Login for get service and get password.

After login your password cached in app for < YOUR DELAY TIME SET IN ENV >.

If you login again cached password updated!

Request:

```json
{
    "password": "1111111111111112" // length only 16 or 32 character
}
```

Response: ( get jwt token )

```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTE0NTk2MDEsImlkIjoxfQ.WF-kqFAl-SjZTvXZq-x95FC9QB_KY0FuhykwXhDNk-o"
}
```

Note: this token most be set on header.
Header Name : < Authorization >

## Service Endpoint

### Create New Service

Endpoints: ( POST )
> 0.0.0.0:1111/api/service

This endpoint for create new service.

Service use for create new password.

Request:

```json
{
    "name" : "google.com"
}
```

Response:

```json
{
    "serviceid": 1,
    "name": "google.com"
}
```

### Update Service

Endpoints: ( PUT )
> 0.0.0.0:1111/api/service/:id

Need Parm **id**.

Request:

```json
{
    "name": "gooogleee.cooom"
}
```

Response:

```json
{
    "serviceid": 1,
    "name": "gooogleee.cooom"
}
```

### Get Service

Endpoints: ( GET )
> 0.0.0.0:1111/api/service/:id

Need Parm **id**.

Response:

```json
{
    "serviceid": 1,
    "name": "gooogleee.cooom"
}
```

### Get All Service

Endpoints: ( GET )
> 0.0.0.0:1111/api/service

Response:

```json
[
    {
        "serviceid": 1,
        "name": "gooogleee.cooom"
    }
]
```

### Delete Service

Endpoints: ( Delete )
> 0.0.0.0:1111/api/service/:id

Need Parm **id**.

Response:

```json
{
    "serviceid": 1,
    "name": "gooogleee.cooom"
}
```

## Password Endpoint

### Create Password

Endpoints: ( POST )
> 0.0.0.0:1111/api/password

Request:

```json
{
    "serviceid" : 1,
    "username" : "test",
    "password" : "test",
    "note": "test"
}
```

Response:

```json
{
    "passid": 1,
    "serviceid": 1,
    "username": "test",
    "password": "test",
    "note": "test"
}
```

### Delete Password

Endpoints: ( DELETE )
> 0.0.0.0:1111/api/password/:id

Need param < id >.

Response:

```json
{
    "passid": 1,
    "serviceid": 1,
    "username": "test",
    "password": "test",
    "note": "test"
}
```

### Update Password

Endpoints: ( PUT )
> 0.0.0.0:1111/api/password/:id

Need param < id >.

Request:

```json
{
    "serviceid" : 1,
    "username" : "test",
    "password" : "test",
    "note": "test"
}
```

Response:

```json
{
    "passid": 1,
    "serviceid": 1,
    "username": "test",
    "password": "test",
    "note": "test"
}
```

### Get Password

Endpoints: ( GET )
> 0.0.0.0:1111/api/password/:id

Need param < id >.

Response: ( with decrypt )

```json
{
    "passid": 1,
    "serviceid": 1,
    "username": "test",
    "password": "test",
    "note": "test"
}
```

### Get All Password

Endpoints: ( GET )
> 0.0.0.0:1111/api/password

Query decrypt is optional. ( type is bool )

Response:

```json
[
    {
        "passid": 2,
        "serviceid": 1,
        "username": "test",
        "password": "test",
        "note": "test"
    },
    {
        "passid": 3,
        "serviceid": 1,
        "username": "test",
        "password": "test",
        "note": "test"
    }
]
```

## User

### Update User Password

Endpoints: ( PUT )
> 0.0.0.0:1111/api/user

Request:

```json
{
    "old": "1111",
    "new": "2222"
}
```

Response:

```json
{
    "id": 1,
    "password": "1111111111111112"
}
```

## Errors

If program get errors, response with this Struct:

Status Code : 401, 500

```json
{
    "detail" : "error message"
}
```

Note: If you are request to some path then needed auth only get 401 error status code !
