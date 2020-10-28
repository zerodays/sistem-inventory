package item

import (
	"database/sql"
	"github.com/zerodays/sistem-inventory/internal/database"
	"time"
)

// Represents company expense.
type Item struct {
	ID int `db:"id"`

	// Name of purchased item / service, cannot be null
	Name string `db:"name"`

	// Item description, can be null.
	Description string `db:"description"`

	// Time when item was purchased.
	DatePurchased time.Time `db:"date_purchased"`

	// Item price in cents.
	Price int `db:"price"`
}

// ForID returns user for specified ID.
func ForID(id int) (*Item, error) {
	// Query item.
	item := &Item{}
	query := `SELECT * FROM items WHERE id=$1`
	err := database.DB.Get(item, query, id)

	// Handle error
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (item *Item) Update(name, description string, datePurchased sql.NullTime, price int) error {
	update := `UPDATE items SET name=$2, description=$3, date_purchased=$4, price=$5 WHERE id=$1 RETURNING *`
	return database.DB.Get(item, update, item.ID, name, description, datePurchased, price)
}

func (item *Item) Insert() error {
	// Save item to database
	insert := `INSERT INTO items (name, description, date_purchased, price) VALUES ($1, $2, $3, $4) RETURNING *`
	err := database.DB.Get(item, insert, item.Name, item.Description, item.DatePurchased, item.Price)

	return err
}

func ListItems() ([]Item, error) {
	var items []Item

	query := `SELECT * FROM items ORDER BY date_purchased`
	err := database.DB.Select(&items, query)

	return items, err
}
