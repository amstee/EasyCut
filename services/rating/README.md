# Rating

[Previous page](../README.md)

This service allow third parties to rate barbers, salons and appointments

A rating is linked to either a user, a salon or an appointment.

This service exposes six different endpoints for ratings management :

* Server status : `GET /rating/status`
* Create salon : `POST /rating/rate`
* Get salon `GET /rating/get/{rating}`
* List salons : `GET /rating/list`
* Update salon : `PUT /rating/update/{rating}`
* Delete salon : `DELETE /rating/delete/{rating}`

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
