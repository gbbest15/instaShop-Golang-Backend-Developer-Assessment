# instaShop-Golang-Backend-Developer-Assessment

**Documentation**

For detailed documentation on the project's API endpoints, please refer to the Postman collection file located at `instaShop/insta-Shop.postman_collection.json`. This file contains information on how to interact with the API, including request and response formats.

**Getting Started**

To run this project, you need to have Go installed on your machine. If you haven't installed Go yet, follow the instructions on the official Go website: https://golang.org/dl/

**Database Setup**

This project uses PostgreSQL as the database. You need to install PostgreSQL and create a database for this project. You can use the following command to create a database:
```sql
CREATE DATABASE instashop;
```
**Environment Variables**

You need to set the following environment variables to connect to your PostgreSQL database:
```go
DB_USER=your_username
DB_PASSWORD=your_password
DB_NAME=instashop
DB_PORT=5432
```
**Sample Code**

Here's a sample code snippet to connect to the database:
```go
package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// Define the connection string
	connStr := "user=your_username dbname=instashop sslmode=disable password=your_password"
	// Open the connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}
	defer db.Close()

	// Ping the database to verify the connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %q", err)
	}
	fmt.Println("Successfully connected to the database!")
}


