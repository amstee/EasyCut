# Rating

[Previous page](../README.md)

This service allow third parties to rate barbers, ratings and appointments

A rating is linked to either a user, a rating or an appointment.

This service exposes six different endpoints for ratings management :

* Server status : `GET /rating/status`
* Create rating : `POST /rating/rate`
* Get rating `GET /rating/get/{rating}`
* List ratings : `GET /rating/list`
* Update rating : `PUT /rating/update/{rating}`
* Delete rating : `DELETE /rating/delete/{rating}`

## Service status

Send the service current status and version

#### Route : `/rating/status`

#### Type : `GET`

#### Response :

```json
{
  "status": "ok",
  "service": "rating",
  "version": "0.0.1"
}
```

## Create rating

Create a rating

#### Route : `/rating/rate`

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
  "user": "string",
  "barber": "string",
  "salon": "string",
  "comment": "string",
  "stars": 0,
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
    "target_id": "string",
    "target_type": "string",
    "comment": "string",
    "stars": 0,
    "created": "date",
    "updated": "date"
}
```

## Get rating

Get a rating info

#### Route : `/rating/get/{rating}`

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
    "target_id": "string",
    "target_type": "string",
    "comment": "string",
    "stars": 0,
    "created": "date",
    "updated": "date"
}
```

## ratings list

Get many ratings

#### Route : `/rating/list`

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
  "target_id": "te*",
  "target_type": "salon",
  "min_stars": 1,
  "max_stars": 5
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
        "target_id": "string",
        "target_type": "string",
        "comment": "string",
        "stars": 0,
        "created": "date",
        "updated": "date"
    }
]
```

## Update rating

Update a rating's info

#### Route : `/rating/update/{rating}`

#### Type : `PUT`

#### Permissions : rating

#### Headers : 

```json
{
  "authorization": "Bearer {{USER_TOKEN}}"
}
```

#### Body :

```json
{
  "comment": "string",
  "stars": 0,
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

Status code 200

```json
{
    "_id": "string",
    "target_id": "string",
    "target_type": "string",
    "comment": "string",
    "stars": 0,
    "created": "date",
    "updated": "date"
}
```

## Delete rating

Delete a rating

#### Route : `/rating/delete/{rating}`

#### Type : `DELETE`

#### Permissions : rating

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
  "message": "rating deleted",
  "success": true
}
```
