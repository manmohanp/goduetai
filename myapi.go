package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// DB is the database connection pool.
var DB *sql.DB

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		goDotEnvVariable("DB_HOST"), goDotEnvVariable("DB_PORT"), goDotEnvVariable("DB_USERNAME"), goDotEnvVariable("DB_PASSWORD"), goDotEnvVariable("DB_DBNAME"))

	log.Print("Connecting to database...")

	var err error
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Connected to database")

	//create table if dont exists
	log.Print("Creating table if not exists...")
	stmt, err := DB.Prepare("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name VARCHAR(255), email VARCHAR(255))")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	if _, err := stmt.Exec(); err != nil {
		log.Fatal(err)
	}
	log.Print("Table created")

	// Create a new router.
	r := gin.Default()

	// Define the routes.
	r.GET("/users", getUsers)
	r.GET("/users/:id", getUser)
	r.POST("/users", createUser)
	r.PUT("/users/:id", updateUser)
	r.DELETE("/users/:id", deleteUser)

	// Start the server.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8090"
	}
	log.Printf("Listening on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}

// getUsers gets all users from the database.
func getUsers(c *gin.Context) {
	// Get all users from the database.
	rows, err := DB.Query("SELECT * FROM users")
	if err != nil {
		c.Error(err)
		return
	}

	// Iterate over the rows and write them to the response.
	for rows.Next() {
		var id int
		var name string
		var email string
		if err := rows.Scan(&id, &name, &email); err != nil {
			c.Error(err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"id":    id,
			"name":  name,
			"email": email,
		})
	}
}

// getUser gets a user from the database.
func getUser(c *gin.Context) {
	// Get the user ID from the request.
	id := c.Param("id")

	// Get the user from the database.
	row := DB.QueryRow("SELECT * FROM users WHERE id = ?", id)
	var name string
	var email string
	if err := row.Scan(&id, &name, &email); err != nil {
		c.Error(err)
		return
	}

	// Write the user to the response.
	c.JSON(http.StatusOK, gin.H{
		"id":    id,
		"name":  name,
		"email": email,
	})
}

// createUser creates a user in the database.
func createUser(c *gin.Context) {
	var input CreateUserRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Print(err.Error())
		return
	}

	// Insert the user into the database.
	stmt, err := DB.Prepare("INSERT INTO users (name, email) VALUES ($1, $2)")
	if err != nil {
		c.Error(err)
		return
	}
	defer stmt.Close()
	if _, err := stmt.Exec(input.Name, input.Email); err != nil {
		c.Error(err)
		return
	}

	// Write a success message to the response.
	c.JSON(http.StatusOK, gin.H{
		"message": "User created",
	})
}

// updateUser updates a user in the database.
func updateUser(c *gin.Context) {
	// Get the user ID from the request.
	id := c.Param("id")

	var input CreateUserRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Print(err.Error())
		return
	}

	// Update the user in the database.
	stmt, err := DB.Prepare("UPDATE users SET name = $1, email = $2 WHERE id = $3")
	if err != nil {
		log.Print(err.Error())
		c.Error(err)
		return
	}
	defer stmt.Close()
	if _, err := stmt.Exec(input.Name, input.Email, id); err != nil {
		log.Print(err.Error())
		c.Error(err)
		return
	}

	// Write a success message to the response.
	c.JSON(http.StatusOK, gin.H{
		"message": "User updated",
	})
}

// deleteUser deletes a user from the database.
func deleteUser(c *gin.Context) {
	// Get the user ID from the request.
	id := c.Param("id")

	// Delete the user from the database.
	stmt, err := DB.Prepare("DELETE FROM users WHERE id = $1")
	if err != nil {
		c.Error(err)
		return
	}
	defer stmt.Close()
	if _, err := stmt.Exec(id); err != nil {
		c.Error(err)
		return
	}

	// Write a success message to the response.
	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted",
	})
}
