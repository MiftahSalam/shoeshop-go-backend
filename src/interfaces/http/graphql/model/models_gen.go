// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type Product struct {
	Name         string    `json:"name"`
	Description  *string   `json:"description"`
	ImageURL     *string   `json:"imageUrl"`
	Rating       int       `json:"rating"`
	Price        float64   `json:"price"`
	NumReviews   int       `json:"numReviews"`
	CountInStock int       `json:"countInStock"`
	Reviews      []*Review `json:"reviews"`
}

type Review struct {
	Name    string `json:"name"`
	Rating  int    `json:"rating"`
	Comment string `json:"comment"`
	User    *User  `json:"user"`
}

type User struct {
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	IsAdmin   bool       `json:"isAdmin"`
	CreatedAt *time.Time `json:"createdAt"`
}
