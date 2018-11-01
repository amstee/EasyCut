# Barber

[Previous page](../README.md)

This service allow third parties to manipuate easy cut barbers, with this service you can
create, update, ... barbers.

A barber is linked to a user, meaning that they share the same id.

This service exposes six different endpoints for barbers management :

* Server status : `GET /barber/status`
* Create barber : `POST /barber/create/{user}`
* Get barber `GET /barber/get/{user}`
* List barbers : `GET /barber/list`
* Update barber : `PUT /barber/update/{user}`
* Delete barber : `DELETE /barber/delete/{user}`

## Service status

Send the service current status and version

#### Route : `/barber/status`

#### Type : `GET`

#### Response :

```json
{
  "status": "ok",
  "service": "barber",
  "version": "0.0.1"
}
```

## Create barber

Create a barber

#### Route : `/barber/create/{user}`

#### Type : `POST`

#### Permissions : User

#### Headers : 

```json
{
  "content-type": "application/json",
  "authorization": "Bearer {{USER_TOKEN}}"
}
```

#### Body :

```json
{
  "experience": "string",
  "style": "string",
  "created": "string",
  "updated": "string"
}
```

#### Response on error :

```json
{
  "message": "error message",
  "success": false
}
```

#### Response on success :

Status code 201

```json
{
    "user": {
      "user ...": "data ..."
    },
    "barber": {
      "experience": "string",
      "style": "string",
      "created": "string",
      "updated": "string"
    }
}
```

## Get Barber

Get a barber info

#### Route : `/barber/get/{user}`

#### Type : `GET`

#### Permissions : User

#### Headers : 

```json
{
  "authorization": "Bearer {{USER_TOKEN}}"
}
```

#### Response on error :

```json
{
  "message": "error message",
  "success": false
}
```

#### Response on success :

Status code 200

```json
{
    "user": {
      "user ...": "data ..."
    },
    "barber": {
      "experience": "string",
      "style": "string",
      "created": "string",
      "updated": "string"
    }
}
```

## Barbers list

Get a barbers list

#### Route : `/barber/list`

#### Type : `GET`

#### Permissions : User

#### Headers : 

```json
{
  "authorization": "Bearer {{USER_TOKEN}}"
}
```

#### URL Params

```json
{
  "email": "jer*",
  "nickname": "jer*",
  "username": "ams*"
}
```

#### Response on error :

```json
{
  "message": "error message",
  "success": false
}
```

#### Response on success :

Status code 200

```json
[
    {
      "user": {
          "user ...": "data ..."
       },
      "barber": {
          "experience": "string",
          "style": "string",
          "created": "string",
          "updated": "string"
        }
    }
]
```

## Update barber

Update a barber's info

#### Route : `/barber/update/{user}`

#### Type : `PUT`

#### Permissions : Barber

#### Headers : 

```json
{
  "authorization": "Bearer {{USER_TOKEN}}"
}
```

#### Body :

```json
{
    "barber": {
      "experience": "string",
      "style": "string",
      "created": "string",
      "updated": "string"
    }
}
```

#### Response on error :

```json
{
  "message": "error message",
  "success": false
}
```

#### Response on success :

Status code 200

```json
{
    "barber": {
      "experience": "string",
      "style": "string",
      "created": "string",
      "updated": "string"
    }
}
```

## Delete barber

Update a barber's info

#### Route : `/barber/delete/{user}`

#### Type : `DELETE`

#### Permissions : Barber

#### Headers : 

```json
{
  "authorization": "Bearer {{USER_TOKEN}}"
}
```

#### Response on error :

```json
{
  "message": "error message",
  "success": false
}
```

#### Response on success :

Status code 200

```json
{
  "message": "barber deleted",
  "success": true
}
```
