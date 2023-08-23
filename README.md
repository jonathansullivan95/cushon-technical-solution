# cushon-technical-solution

Implementation of Cushon technical assessment

## Setup

This command will spin up a local mysql database and initialise the schema:

    make start-db

The database has persistent storage so you can spin it down without affecting the data using this command:

    make stop-db

This command will build and run the application:

    make start

This runs on port 8080 by default. To run on a specific port you can run the following:

    go build
    ./cushon-technical-solution -port=9000

## Endpoints

- `/funds` - returns a list of available investment funds
- `/customeraccount/get` - returns a given customer's investment account
- `/customeraccount/investmentfund/add` - adds a given amount to a customer's investment fund

clear api definitions are listed in the api.md file
