# GO API Template

This is my goto template for creating a new API in Go. It includes a basic structure for the project, logging, 
configuration, and database setup. It's a slimmed down version of the template we use for our Go services at my job with some changes for new libraries
or patterns I prefer. 

## Getting Started
1. Clone the repo to a directory with your project name
2. Update `go.mod` file with your project name
3. Update `docker-compose.yml` file with your project name in the relevant places (DB name, app name, etc)
4. Update module name in `.mockery.yml` to match your project name
5. Update variables in the Makefile to match your project name ($PROJNAME and $DB_CONNECTION_URI for example)
6. Rename `.env-template` to `.env` and set appropriate values
7. The template sets up Users as an example in domain, service, and data directories. Update or delete as you see fit. Delete the migrations and start new ones
8. Go build something cool!

## Libraries
* Router - [Chi](https://github.com/go-chi/chi)
* Logging - [Zerolog](https://github.com/rs/zerolog)
* .ENV file support - [godotenv](https://github.com/joho/godotenv)
* Reading config into structs - [cleanenv](https://github.com/ilyakaznacheev/cleanenv)
* DB extensions - [SQLX](https://github.com/jmoiron/sqlx)
* DB migrations - [go-migrate](https://github.com/golang-migrate/migrate)
* Mocking - [mockery](https://github.com/vektra/mockery)
* Testify - [testify](https://github.com/stretchr/testify)
 
## Middleware
* chi's RequestID - adds a request ID to the context
* chi's RealIP - adds the real IP to the context
* custom RequestResponseLogger - logs the request and response. Also adds requestID to each log entry for correlating logs
* chi's Heartbeat - adds a /ping endpoint to check if the server is up
* custom Recoverer - recovers from panics and logs the error. Custom so that error is logged with requestID as well as message and string stack trace to make alerting on and viewing in log aggregators easier
* chi's Compress - compresses the response
* chi's Timeout - adds a timeout to the request context