# OAUth2.0 Golang template

#### All APIs need some sort of authentication (unless they're public of course, which seems to be fairly uncommon) and so when your developing an API from scratch, one of the first things youre gonna have to implement is your auth solution.

#### Unfortunately however, there does seem to be a lot of complexity involved in choosing the right solution and integrating it properly...


#### So I created this repository, a project template with the following features

- Simple REST API using the popular [gin](https://github.com/gin-gonic/gin) framework for golang.
- OAuth2.0 authentication provided by Auth0

- All endpoints */api/(host|exhibition|host|user)* require JWT authentication 

- Additionally, the GET /api/host endpoint requires a token with the `read:host` permission.

### If this meets your use case or similar, hopefully you can clone this template to simplify the process of getting your API up and running

## How to run it
- Copy the file named .env.example file into a new file named .env
  - `cp .env .env.example`
- Fill out the values in .env with values taken from your Auth0 account (see below for an example Auth0 setup)
- Start the server
```
# From the repository root
go mod download
go run .
```

## Auth0 Setup
- To use this template youll need to create an account with Auth0 and configure this API to validate the JWT using your OAuth server

1. Go to [https://auth0.com/](https://auth0.com/) and create an account
    - At the time of writing, I created a free account which allows me to create up to 1000 tokens a month
2. From your [dashboard](https://manage.auth0.com/dashboard/) create 2 new applications
    - An API application (This is your API)
        - Add the permissions your API uses here (e.g `read:host`, `read:user`...)
        - Copy the  value of the `audience` setting (the same as `Identifier`) into your `AUTH0_AUDIENCE` environment variable set in the .env file
    - A Machine to Machine (M2M) application (this is your API client)
        - When creating the client application, you should be prompted with a question for which permissions to grant it. Add all (or some, as desired) of the permissions you previously created in the API application.
        - Copy the values of Domain, Client ID, and Client Secret into the `AUTH0_CLIENT_ID`, `AUTH0_DOMAIN`, `AUTH0_CLIENT_SECRET` environment variableS set in the .env file
    - Authorise the API Client to use the API application (you can do this from the 'Machine to Machine applications tab in the API application settings)

## Thats it! Hope it works for you 🍀

