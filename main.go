package main

import (
	"fmt"
	"github.com/voyago/environment"
	"gocantoserver/foundation"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	env, err := environment.Make("server")

	if err != nil {
		panic("The given env [server/.env] is invalid.")
	}

	app, err := foundation.CreateApp(env)

	if err != nil {
		message, _ := fmt.Printf("There was an issue creating the App: %v", err)
		panic(message)
	}

	app.Start()

	db := app.GetConnection().GetDB()

	fmt.Println(db.Migrator().CurrentDatabase())

	// Migrate the schema
	//db.AutoMigrate(&Product{})
	//
	//// Create
	//db.Create(&Product{Code: "D42", Price: 100})
	//
	//// Read
	//var product Product
	//db.First(&product, 1) // find product with integer primary key
	////db.First(&product, "code = ?", "D42") // find product with code D42
	//
	//fmt.Println(product)

	// Update - update product's price to 200
	//db.Model(&product).Update("Price", 200)
	//// Update - update multiple fields
	//db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	//
	//// Delete - delete product
	//db.Delete(&product, 1)
}
