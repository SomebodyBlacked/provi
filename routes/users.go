package routes

import (
	"net/http"
	"provi/database"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func Users(c *gin.Context) {
	db, err := database.Connect() // Use the Connect function from db.go
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		c.JSON(500, gin.H{"error": "Error querying the database"})
		return
	}
	defer rows.Close()

	var users []User
	columns, err := rows.Columns()
	if err != nil {
		c.JSON(500, gin.H{"error": "Error retrieving columns"})
		return
	}

	for rows.Next() {
		values := make([]interface{}, len(columns))
		valuePointers := make([]interface{}, len(columns))
		for i := range values {
			valuePointers[i] = &values[i]
		}

		if err := rows.Scan(valuePointers...); err != nil {
			c.JSON(500, gin.H{"error": "Error scanning row"})
			return
		}

		user := User{}
		for i, colName := range columns {
			val := values[i]
			switch colName {
			case "id":
				user.ID = val.(int64)
			case "username":
				usernameBytes := val.([]uint8)
				user.Username = string(usernameBytes)
			case "email":
				emailBytes := val.([]uint8)
				user.Email = string(emailBytes)
			}
		}

		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		c.JSON(500, gin.H{"error": "Error during rows iteration"})
		return
	}

	c.JSON(200, gin.H{
		"users": users,
	})
}
