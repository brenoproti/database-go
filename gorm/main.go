package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    int
	ProductID int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	// Create with FK
	// c := Category{Name: "Eletronicos"}
	// db.Create(&c)

	// p := Product{
	// 	Name:       "Notebook",
	// 	Price:      8000.00,
	// 	CategoryID: c.ID,
	// }
	// db.Create(&p)

	// s := SerialNumber{
	// 	Number: 1234,
	// }
	// db.Create(&s)

	// Create Product
	// products := []Product{
	// 	{Name: "Notebook", Price: 5000.0},
	// 	{Name: "Mouse", Price: 100.0},
	// 	{Name: "Keyboard", Price: 200.0},
	// }
	// db.Create(&products)

	// Get one
	// var p Product
	// db.First(&p, 1)
	// db.First(&p, "name = ?", "Mouse")

	// Get all
	// var productsSel []Product
	// db.Find(&productsSel)

	// Get all with limit and offset
	// db.Limit(2).Offset(1).Find(&productsSel)
	// db.Find(&productsSel)

	// Get with where
	// db.Where("price > ?", 100).Find(&products)
	// db.Where("name LIKE ?", "%book%").Find(&products)

	// Get with preload
	// var products []Product
	// db.Preload("Category").Find(&products)

	// Get Has many
	// var categories []Category
	// db.Model(&Category{}).Preload("Products.SerialNumber").Find(&categories)

	// Update
	// var p Product
	// db.First(&p, 1)
	// p.Name = "New name"
	// db.Save(&p)

	// Delete
	// db.First(&p, 1)
	// db.Delete(&p)

}
