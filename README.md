# GO API Template

This is my goto template for creating a new API in Go. It includes a basic structure for the project, logging, 
configuration, and database setup. It's a slimmed down version of the template we use for our Go services at my job but made simpler for my side projects. 

## Getting Started
1. Clone the repo to a directory with your project name
2. Update `docker-compose.yml` file with your project name in the relevant places (DB name, app name, etc)
3. Update variables in the Makefile to match your project name ($PROJNAME and $DB_CONNECTION_URI for example)
4. Rename `.env-template` to `.env` and set appropriate values
5. The template sets up Users as an example in domain, service, and data directories. Update or delete as you see fit. Delete the migrations and start new ones
6. Go build something cool!

## Components
* Router - [Chi](https://github.com/go-chi/chi)
* Logging - [Zerolog](https://github.com/rs/zerolog)
* .ENV file support - [godotenv](https://github.com/joho/godotenv)
* Reading config into structs - [cleanenv](https://github.com/ilyakaznacheev/cleanenv)
* DB extensions - [SQLX](https://github.com/jmoiron/sqlx)
* DB migrations - [go-migrate](https://github.com/golang-migrate/migrate)
 
## Middleware
* chi's RequestID - adds a request ID to the context
* chi's RealIP - adds the real IP to the context
* custom RequestResponseLogger - logs the request and response. Also adds requestID to each log entry for correlating logs
* chi's Heartbeat - adds a /ping endpoint to check if the server is up
* custom Recoverer - recovers from panics and logs the error. Custom so that error is logged with requestID as well as message and string stack trace to make alerting on and viewing in log aggregators easier
* chi's Compress - compresses the response
* chi's Timeout - adds a timeout to the request context