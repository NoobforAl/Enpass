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
> 0.0.0.0:1111/api/createService

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
> 0.0.0.0:1111/api/updateService/:id

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
> 0.0.0.0:1111/api/allService

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
> 0.0.0.0:1111/api/deleteService/:id

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
> 0.0.0.0:1111/api/createPass

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
> 0.0.0.0:1111/api/deletePass/:id

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
> 0.0.0.0:1111/api/updatePass/:id

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
> 0.0.0.0:1111/api/pass/:id/?decrypt=true

Need param < id >.

Query decrypt is optional. ( type is bool )

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

Response: ( without decrypt )

```json
{
    "passid": 1,
    "serviceid": 1,
    "username": "BNA6+NOivbSzguaDA4rBexCPd16EI90DpmM2xhlf0a64",
    "password": "Yt6+TBTGu/bBa/BPUK2pKca/KlhpHqJuoLWisbUDA/fd",
    "note": "wA5p1S66MErvIj8DkDs0t4MiFctjT/RniI4QcHzeLlJb"
}
```

### Get All Password

Endpoints: ( GET )
> 0.0.0.0:1111/api/allPass?decrypt=false

Query decrypt is optional. ( type is bool )

Response:

```json
[
    {
        "passid": 2,
        "serviceid": 1,
        "username": "YYWxe0TX0dw8u9ATKrFxgEy8em5uUonKJG4ZdUb3rbk=",
        "password": "n/inlMMBcZVtylKHX8JU/neU7ay6mnsz4uX7hprFYCM=",
        "note": "Y/tldFPQezjLgTKWGs2Wj86TeadoPzra4Sd9U5VORvQ="
    },
    {
        "passid": 3,
        "serviceid": 1,
        "username": "9zgI29jjasskZiG8A6Yj5qYCYr8B7S24cMPmbBEcAxQ=",
        "password": "Lu1lIawtqH2fRseHQxk4bFScS166upj0gCwMCdrM+S0=",
        "note": "eMygc7T+26Ssrfx82YwsV68XbUES4sLsl25mPkGOIOM="
    }
]
```

## User

### Update User Password

Endpoints: ( PUT )
> 0.0.0.0:1111/api/updateUser

Request:

```json
{
    "old": "1111111111111111",
    "new": "1111111111111112"
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
