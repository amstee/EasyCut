# appointment

[Previous page](../README.md)

This service allow third parties to create appointments

An appointment is linked to two users.

This service exposes six different endpoints for appointments management :

* Server status : `GET /appointment/status`
* Create salon : `POST /appointment/schedule`
* Get salon `GET /appointment/get/{appointment}`
* List salons : `GET /appointment/list`
* Update salon : `PUT /appointment/update/{appointment}`
* Delete salon : `DELETE /appointment/delete/{appointment}`

## Service status

Send the service current status and version

#### Route : `/appointment/status`

#### Type : `GET`

#### Response :

```json
{
  "status": "ok",
  "service": "appointment",
  "version": "0.0.1"
}
```

## Schedule an appointment

Schedule an appointment

#### Route : `/appointment/schedule`

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
  "user_id": "string",
  "barber_id": "string",
  "description": "string",
  "date": "date",
  "duration": 0,
  "created": "date",
  "updated": "date"
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
  "barber_id": "string",
  "description": "string",
  "date": "date",
  "duration": 0,
  "created": "date",
  "updated": "date"
}
```

## Get appointment

Get a appointment info

#### Route : `/appointment/get/{appointment}`

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
  "barber_id": "string",
  "description": "string",
  "date": "date",
  "duration": 0,
  "created": "date",
  "updated": "date"
}
```

## appointments list

Get many appointments

#### Route : `/appointment/list`

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
  "user_id": "string",
  "barber_id": "string",
  "min_date": "date",
  "max_date": "date"
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
      "barber_id": "string",
      "description": "string",
      "date": "date",
      "duration": 0,
      "created": "date",
      "updated": "date"
    }
]
```

## Update appointment

Update a appointment's info

#### Route : `/appointment/update/{appointment}`

#### Type : `PUT`

#### Permissions : appointment

#### Headers : 

```json
{
  "authorization": "Bearer {{USER_TOKEN}}"
}
```

#### Body :

```json
{
  "description": "string",
  "duration": 0,
  "date": "date"
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
  "barber_id": "string",
  "description": "string",
  "date": "date",
  "duration": 0,
  "created": "date",
  "updated": "date"
}
```

## Delete appointment

Delete a appointment

#### Route : `/appointment/delete/{appointment}`

#### Type : `DELETE`

#### Permissions : appointment

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
  "message": "appointment deleted",
  "success": true
}
```
