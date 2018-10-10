# Easy Cut Server

## Project organization

### Services

[Services Directory](services/)

Each service is located in its own directory associated with its name, the API endpoints are matching
with their directory name. Example : Create account --> auth/create/account

Each service is defined by a Dockerfile allowing us to easily deploy each service

### Infrastructure

[Infrastructure Directory](infra/)

The infrastructure directory contains the logic implemented for our projet deployment