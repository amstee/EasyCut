# Salon

[Previous page](../README.md)

This service allow third parties manipulate easy cut salons, with this service you can create, update, ... salons.

A salon is linked to a user, meaning this user is the manager.

This services exposes six different endpoints for salons management :

* Server status : `GET /salon/status`
* Create salon : `POST /salon/create/`
* Get salon `GET /salon/get/{salon}`
* List salons : `GET /salon/list`
* Update salon : `PUT /salon/update/{salon}`
* Delete salon : `DELETE /salon/delete/{salon}`

## Service status

Send the service current status and version

#### Route : `/salon/status`

#### Type : `GET`

#### Response :

```json
{
  "status": "ok",
  "service": "salon",
  "version": "0.0.1"
}
```

## Create salon

Create a salon

#### Route : `/salon/create/`

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
    "name": "string",
    "address": "string",
    "employee_number": 10,
    "barbers": ["string"],
    "website": "string",
    "created": "created",
    "updated": "updated"
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
    "_id": "string",
    "user_id": "string",
    "name": "string",
    "address": "string",
    "employee_number": 10,
    "barbers": ["string"],
    "website": "string",
    "created": "created",
    "updated": "updated"
}
```

## Get salon

Get a salon info

#### Route : `/salon/get/{salon}`

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
    "_id": "string",
    "user_id": "string",
    "name": "string",
    "address": "string",
    "employee_number": 10,
    "barbers": ["string"],
    "website": "string",
    "created": "created",
    "updated": "updated"
}
```

## salons list

Get many salons

#### Route : `/salon/list`

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
  "address": "te*",
  "min_emp": 0,
  "max_emp": 10
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
        "_id": "string",
        "user_id": "string",
        "name": "string",
        "address": "string",
        "employee_number": 10,
        "barbers": ["string"],
        "website": "string",
        "created": "created",
        "updated": "updated"
    }
]
```

## Update salon

Update a salon's info

#### Route : `/salon/update/{salon}`

#### Type : `PUT`

#### Permissions : salon

#### Headers : 

```json
{
  "authorization": "Bearer {{USER_TOKEN}}"
}
```

#### Body :

```json
{
    "name": "string",
    "address": "string",
    "employee_number": 10,
    "barbers": ["string"],
    "website": "string",
    "created": "created",
    "updated": "updated"
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
    "_id": "string",
    "user_id": "string",
    "name": "string",
    "address": "string",
    "employee_number": 10,
    "barbers": ["string"],
    "website": "string",
    "created": "created",
    "updated": "updated"
}
```

## Delete salon

Delete a salon

#### Route : `/salon/delete/{salon}`

#### Type : `DELETE`

#### Permissions : salon

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
  "message": "salon deleted",
  "success": true
}
```
