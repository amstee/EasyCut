# Perms

[Previous page](../README.md)

This service allow to manage easy cut users permissions

It exposes three different endpoints :

* Server status : `GET /perms/status`
* Get User groups : `GET /perms/get/{user}`
* Update User groups : `PUT /perms/update/{user}`

## Service status

Send the service current status and version

#### Route : `/perms/status`

#### Type : `GET`

#### Response :

```json
{
  "status": "ok",
  "service": "perms",
  "version": "0.0.1"
}
```

## Get Groups

Get the user groups

#### Route : `/perms/get/{user}`

#### Type : `GET`

#### Headers : 

```json
{
  "authorization": "Bearer {{USER_TOKEN}}"
}
```

#### Response :

```json
[
  {
    "_id": "string",
    "name": "string",
    "description": "string"
  }
]
```

#### Response on error :

```json
{
  "message": "error message",
  "success": false
}
```

## Promote user

Add or delete a group to a user

#### Route : `/perms/update/{user}`

#### Type : `PUT`

#### Headers : 

```json
{
  "authorization": "Bearer {{USER_TOKEN}}"
}
```

#### Body : 

```json
{
	"groups": [
		{
			"name": "string",
			"delete": true
		}
	]
}
```

#### Response :

```json
[
    {
        "name": "string",
        "success": true
    }
]
```

#### Response on error :

```json
{
  "message": "error message",
  "success": false
}
```