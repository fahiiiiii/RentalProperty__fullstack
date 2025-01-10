# Property Listing Project

This project is a full-stack web application for listing rental properties in the United Arab Emirates. 

## Setup the Project

1. **Clone the repository**:
   ```sh
   git clone https://github.com/fahiiiiii/RentalProperty__fullstack/tree/main/property-listing
   cd RentalProperty__fullstack
   cd property-listing

 
***Create a .env file:***
Create a .env file in the root directory of your project.
Populate it with the necessary environment variables. You can refer to the dev.env file for the required variables and their descriptions.
***Running the Project***
Run the project:

Recommended:
    ```sh
    go run main.go

Alternative:
    ```sh
    bee run

This will start the server, fetch data, create tables, and store data into the database.

***Viewing the Database***
Access the Docker container:

    ```sh
    docker exec -it property-listing-db-1 /bin/bash
    
Connect to the PostgreSQL database:

    ```sh
    psql -U fahimah -d rental_db

Check the tables:
 SQL
\dt

Run sample queries:

SQL
SELECT PropertyName, City, Country FROM locations LIMIT 10;
-- or,
SELECT * FROM "locations" LIMIT 10;
SELECT * FROM rental_properties;
SELECT * FROM property_details;


Dependencies
Ensure you have the following dependencies installed:

Go
Bee
Docker
