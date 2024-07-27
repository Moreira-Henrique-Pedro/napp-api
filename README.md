# napp-api


## Requirements

* Golang: [Install Guide](https://golang.org/doc/install)
* PostgreSQL: [Install Guide](https://www.postgresql.org/)

## Environment Variables

Make sure to set the environment variables required by the application. You can create a `.env` file in the root directory with the following content:

DB_USERNAME=napp_db
DB_PASSWORD=napp_db_pass
DB_HOST=db
DB_PORT=5432
DB_NAME=stocks

# Running the Application Locally

To run the application locally, execute the following command:

```bash
go run main.go
```

## Running the Application Locally with Docker

To run the application using Docker, follow these steps:

1. **Ensure Docker is Installed**

   Make sure Docker and Docker Compose are installed on your machine. If not, you can download and install Docker Desktop from [here](https://www.docker.com/products/docker-desktop).

2. **Clone the Repository**

   Clone the repository to your local machine:

   ```bash
   git clone https://github.com/Moreira-Henrique-Pedro/napp-api
   ```

3. **Build and Run the Application**

    Use Docker Compose to build and run the application. Execute the following command:

    ```bash
    docker-compose up --build
    ```

    This command will build the Docker images and start the containers for both the application and the PostgreSQL database.

4. **Access the Application**

    Once the containers are up and running, you can access the application at http://localhost:3000.

# API Endpoints

Here are the available API endpoints:

GET /api/stock-service/ - Get all results
GET /api/stock-service/:id - Get result by ID
POST /api/stock-service/ - Create a new stock
PUT /api/stock-service/:id - Update a stock by ID
DELETE /api/stock-service/:id - Delete a stock by ID


