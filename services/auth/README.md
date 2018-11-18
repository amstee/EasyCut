# Auth

[Previous page](../README.md)

This service allow easy cut services or any other third party to validate a easy cut user connection and permissions.

It exposes six different endpoints :

## Service status

Send the service current status and version

#### Route : `/auth/status`

#### Type : `GET`

#### Response :

```json
{
  "status": "ok",
  "service": "auth",
  "version": "0.0.1"
}
```

## Check token

#### Route : `/auth/token`

#### Type : `GET`

#### Headers :

```json
{
  "Authorization": "Bearer [...token]"
}
```

#### Response on success :

```json
{
    "message": "token is valid",
    "success": true
}
```

#### Response on error :

```json
{
    "message": "error description",
    "success": false
}
```

## Check permissions

Allow to check Auth0 user permissions

#### Route `/auth/secure/permissions`

#### Type : `POST`

#### Headers :

```json
{
  "Authorization": "Bearer [...token]"
}
```

#### Body :

```json
{
  "scopes": ["read:users", "..."]
}
```

#### Response on success :

```json
{
  "scopes": {
    "read:users": true,
    "create:whatever": false
  }
}
```

#### Response on error :

```json
{
    "message": "error description",
    "success": false
}
```

## Check groups

Allow to check easy-cut users groups

#### Route `/auth/secure/groups`

#### Type : `POST`

#### Headers :

```json
{
  "Authorization": "Bearer [...token]"
}
```

#### Body :

```json
{
	"groups": ["Admin", "User", "Barber", "..."]
}
```

#### Response on success :

```json
{
  "groups": {
    "Admin": false,
    "User": true,
    "Barber": false
  }
}
```

#### Response on error :

```json
{
    "message": "error description",
    "success": false
}
```

## Match & Groups

Allow to check easy-cut users groups and verify that the user match its token

#### Route `/auth/secure/match/{user}`

#### Type : `POST`

#### Headers :

```json
{
  "Authorization": "Bearer [...token]"
}
```

#### Body :

```json
{
	"groups": ["Admin", "User", "Barber", "..."]
}
```

#### Response on success :

```json
{
  "groups": {
    "Admin": false,
    "User": true,
    "Barber": false
  }
}
```

#### Response on error :

```json
{
    "message": "error description",
    "success": false
}
```


## Extract

Extract info from jwt

#### Route `/auth/secure/extract`

#### Type : `GET`

#### Headers :

```json
{
  "Authorization": "Bearer [...token]"
}
```

#### Get params

```json
{
  "user": "true"
}
```

#### Response on success :

```json
{
  "user_id": "string"
}
```

#### Response on error :

```json
{
    "message": "error description",
    "success": false
}
```
