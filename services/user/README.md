# User

[Previous page](../README.md)

This service allow third parties to manipulate auth0 users, is allows you to find users,
create users, ...
It also allows you to manipulate user's easy cut profile.

It exposes five different endpoints for auth0 users management :

## Service status

Send the service current status and version

#### Route : `/auth/status`

#### Type : `GET`

#### Response :

```json
{
  "status": "ok",
  "service": "user",
  "version": "0.0.1"
}
```

## Create user

Create a user both in auth0 and in easy-cut's storage

#### Route : `/user/create`

#### Type : `POST`

#### Headers : 

```json
{
  "content-type": "application/json"
}
```

#### Body :

```json
{
  "email": "YOUR_EMAIL",
  "password": "YOUR_PASSWORD",
  "email_verified": false,
  "verify_email": true,
  "app_metadata": {
    "authorization": {
      "groups": ["User"]
    }
  }
}
```

#### Response on error :

For the error check the status code of the response, you can find the info about those status codes
at this url : [https://auth0.com/docs/api/management/v2#!/Users/post_users]()

```json
{
    "email": "",
    "email_verified": false,
    "updated_at": "",
    "picture": "",
    "user_id": "",
    "identities": [],
    "created_at": "",
    "app_metadata": {
        "authorization": {
            "groups": []
        }
    }
}
```

#### Response on success :

Status code 201

```json
{
    "email": "test.user@gmail.com",
    "email_verified": false,
    "updated_at": "2018-10-19T16:06:10.388Z",
    "picture": "https://s.gravatar.com/avatar/32c786bfda3809addf2172ca299e0faa?s=480&r=pg&d=https%3A%2F%2Fcdn.auth0.com%2Favatars%2Fte.png",
    "user_id": "auth0|5bca00f240c7bc117a6e5d7e",
    "identities": [
        {
            "connection": "Username-Password-Authentication",
            "user_id": "5bca00f240c7bc117a6e5d7e",
            "provider": "auth0",
            "isSocial": false
        }
    ],
    "created_at": "2018-10-19T16:06:10.388Z",
    "app_metadata": {
        "authorization": {
            "groups": [
                "User"
            ]
        }
    }
}
```

## Update user

## Get user

## List users