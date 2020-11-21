package postgres

import (
	model "github.com/Salaton/screening-test/graph/model"
	"golang.org/x/crypto/bcrypt"
)

// This file contains all functions that are used in the postgres.go file
//Loops over the items and map them to the Item Struct
func loopOverItems(itemsInput []*model.ItemInput) []*model.Item {
	var items []*model.Item
	for _, itemInput := range itemsInput {
		items = append(items, &model.Item{
			Name:     itemInput.Name,
			Quantity: itemInput.Quantity,
		})
	}
	return items
}

// HashPassword to hash the password since we dont want to store
// a password string in the database
func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

// CheckPasswordHash that compares the password and generated hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
