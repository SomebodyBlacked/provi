package routes

import (
	"database/sql"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// Product representa la estructura de un producto.
type Product struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func Products(c *gin.Context) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/provi")
	if err != nil {
		c.JSON(500, gin.H{"error": "Error connecting to the database"})
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		c.JSON(500, gin.H{"error": "Error querying the database"})
		return
	}
	defer rows.Close()

	var products []Product
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

		product := Product{}
		for i, colName := range columns {
			val := values[i]
			switch colName {
			case "id":
				product.ID = val.(int64)
			case "name":
				nameBytes := val.([]uint8)
				product.Name = string(nameBytes)
			case "description":
				descBytes := val.([]uint8)
				product.Description = string(descBytes)
			case "price":
				priceBytes := val.([]uint8)
				priceStr := string(priceBytes)
				price, err := strconv.ParseFloat(priceStr, 64)
				if err != nil {
					c.JSON(500, gin.H{"error": "Error converting price to float"})
					return
				}
				product.Price = price
			}
		}

		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		c.JSON(500, gin.H{"error": "Error during rows iteration"})
		return
	}

	c.JSON(200, gin.H{
		"products": products,
	})
}
