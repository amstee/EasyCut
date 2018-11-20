# Services

[Return to Homepage](../README.md)

This project is composed of a total of six services

## Auth

This service allow easy cut services or any other third party to validate a easy cut user connection and permissions.

It exposes six different endpoints :

* Server status : `GET /auth/status`
* Check token validity : `GET /auth/token` 
* Check User Permissions : `POST /auth/secure/permissions`
* Check user groups : `POST /auth/secure/groups`
* Check user groups and token : `POST /auth/secure/match/{user}`
* Extract info from token : `GET /auth/secure/extract`

To see more details about these endpoints check the [Auth service](auth/README.md)

## Perms

This service allow to manage easy cut users permissions

It exposes three different endpoints :

* Server status : `GET /perms/status`
* Get User groups : `GET /perms/get/{user}`
* Update User groups : `PUT /perms/update/{user}`

To see more details about these endpoints check the [Perms service](perms/README.md)

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

This service allow third parties to manipulate easy cut barbers, with this service you can
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

This service allow third parties to manipulate easy cut salons, with this service you can create, update, ... salons.

A salon is linked to a user, meaning this user is the manager.

This services exposes six different endpoints for salons management :

* Server status : `GET /salon/status`
* Create salon : `POST /salon/create/`
* Get salon `GET /salon/get/{salon}`
* List salons : `GET /salon/list`
* Update salon : `PUT /salon/update/{salon}`
* Delete salon : `DELETE /salon/delete/{salon}`

For more details about these endpoints check the [Salon service](salon/README.md)

## Rating

This service allow third parties to rate barbers, salons and appointments

A rating is linked to either a user, a salon or an appointment.

This service exposes six different endpoints for ratings management :

* Server status : `GET /rating/status`
* Create rating : `POST /rating/rate`
* Get rating `GET /rating/get/{rating}`
* List ratings : `GET /rating/list`
* Update rating : `PUT /rating/update/{rating}`
* Delete rating : `DELETE /rating/delete/{rating}`

For more details about these endpoints check the [Rating service](rating/README.md)

## Appointment

This service allow third parties to create appointments

An appointment is linked to two users.

This service exposes six different endpoints for appointments management :

* Server status : `GET /appointment/status`
* Create appointment : `POST /appointment/schedule`
* Get appointment `GET /appointment/get/{appointment}`
* List appointments : `GET /appointment/list`
* Update appointment : `PUT /appointment/update/{appointment}`
* Delete appointment : `DELETE /appointment/delete/{appointment}`

For more details about these endpoints check the [Appointment service](appointment/README.md)
