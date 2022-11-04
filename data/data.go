package data

import "github.com/Ahasannn/book-server-api/model"

var Books = []model.Book{
	{
		ID:    "4345935",
		Isbn:  "6573434",
		Title: "Book One",
		Author: &model.Author{
			Firstname: "Iftekhar", Lastname: "Fahim",
		},
	},
	{
		ID:    "4359395",
		Isbn:  "3543645",
		Title: "Book Two",
		Author: &model.Author{
			Firstname: "Ariful", Lastname: "Islam",
		},
	},
}