# Wallet API 

This is JSON API in Golang to get the balance and manage credit/debit operations on the players wallets. 
For example, get  the balance of the wallet with id 123,
or to credit/debit the wallet with id 456 for 10.00 €.

## Technologies
- Gin
- MySQL
- Gorm
- Redis

Below are the endpoints available

##Endpoints
Ping Endpoint: pings the API and returns a "pong" message if API is healthy

`GET /api/v1/ping`

Seed Endpoint: endpoint to seed dummy players and wallets

`GET /api/v1/seed`

Auth Endpoint: authenticates a user and return a JWT Token used to access protected routes

`POST /api/v1/login`


### Protected Routes
These routes are protected by the `AuthorizeJWT` Middleware and will require valid JWT tokens to access the endpoints

Balance Endpoint : retrieves the balance of a given wallet with wallet_id

`GET /api/v1/wallets/{wallet_id}/balance`

Credit Endpoint : credits the requested wallet with wallet_id

`POST / api/v1/wallets/{wallet_id}/credit`

Debit Endpoint: debits money from a given wallet with wallet_id

`POST / api/v1/wallets/{wallet_id}/debit `
