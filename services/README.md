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

This service allow third parties to manipulate auth0 users, it allows you to find users,
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

For more details about these endpoints check the [Barber service](barber/README.md)

## Salon

To be developed

## Appointment

To be developed

## Rating

To be developed