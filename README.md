# EnPass

Very Simple Personal Encrypted Password Manager!

## How Work This Program ??

This program use [gin framework](https://gin-gonic.com/) for backend api.

For run this program you can use docker or normally run go program but before run program you need create a .env file to be like .evn.example.

Run with docker
> docker compose up

Run with go
> go mod tidy  
> go run main.go

## Api Help

For Better request you can use postman file on postman folder.

Endpoints:
> 0.0.0.0:1111/genRandomPass?size=< int >

Generate random password, if send none size or bad format size automatic Generate random password with length 10.

Response:

```json
{
    "password": "[NTEdJIJrQ"
}
```

---

Endpoints:
> 0.0.0.0:1111/api/service/:id

Get service with service id.

Response:

```json
{
    "ID": 1,
    "Name": "gooogleee.cooom",
    "CreatedAt": "2023-06-09T22:20:01.54512953+03:30",
    "UpdatedAt": "2023-06-09T22:27:59.655135728+03:30"
}
```

---

Endpoints:
> 0.0.0.0:1111/api/pass/:id?decrypt=true

Get password with password id, if want decrypt values set flag decrypt to true.

```json
{
    "Note": "aaaaa",
    "Password": "aaaaa",
    "UserName": "aaaaa"
}
```

### TODO

[ ] Add more endpoint  
[ ] make uint test
