# Services

[Return to Homepage](../README.md)

This project is composed of a total of six services

## Auth

This service allow easy cut services or any other third party to validate a easy cut user connection and permissions.

It exposes four different endpoints :

* Server status : `GET /auth/status`
* Check token validity : `GET /auth/token` 
* Check User Permissions : `POST /auth/secured/permissions`
* Check user groups : `POST /auth/secured/groups`

To see more details about these endpoints check the [Auth service](auth/README.md)

## User

This service allow third parties to manipulate auth0 users, is allows you to find users,
create users, ...
It also allows you to manipulate user's easy cut profile.

It exposes five different endpoints for auth0 users management :

* Server status : `GET /user/status`
* Create user : `POST /user/create`
* List users : `GET /user/list`
* Get a user info : `GET /user/{user}`
* Update user : `PUT /user/update/{user}`

For more details about these endpoints check the [User service](user/README.md)

## Barber

To be developed

## Salon

To be developed

## Appointment

To be developed

## Rating

To be developed