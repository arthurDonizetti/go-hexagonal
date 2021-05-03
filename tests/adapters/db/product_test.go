package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/arthurDonizetti/go-hexagonal/adapters/db"
	"github.com/arthurDonizetti/go-hexagonal/application"
	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products (
		"id" string,
		"name" string,
		"price" float,
		"status" string
	);`

	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `INSERT INTO products VALUES(
		"any_uuid", "any_name", 0, "disabled"
	);`

	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func TestProductDb_Get(t *testing.T) {
	setUp()
	defer Db.Close()

	productDb := db.NewProductDb(Db)
	product, err := productDb.Get("any_uuid")
	require.Nil(t, err)
	require.Equal(t, "any_name", product.GetName())
	require.Equal(t, 0.0, product.GetPrice())
	require.Equal(t, application.DISABLED, product.GetStatus())
}

func TestProductDb_Save(t *testing.T) {
	setUp()
	defer Db.Close()
	productDb := db.NewProductDb(Db)

	product := application.NewProduct()
	product.Name = "any_name"
	product.Price = 0
	product.Status = application.DISABLED

	productResult, err := productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.GetName(), productResult.GetName())
	require.Equal(t, product.GetPrice(), productResult.GetPrice())
	require.Equal(t, product.GetStatus(), productResult.GetStatus())
	require.Equal(t, product.GetID(), productResult.GetID())

	product.Status = application.ENABLED
	productResult, err = productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.GetStatus(), productResult.GetStatus())
}
