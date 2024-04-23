package main

import "fmt"

// BookType represents the type of book
type BookType int

// Predefined Book types
const (
	HardCover BookType = iota
	SoftCover
	PaperBack
	EBook
)

// Book represents data about a book
type Book struct {
	Name      string
	Author    string
	PageCount int
	Type      BookType
}

// Library holds the collection of books
type Library struct {
	Collection []Book
}

// IterateBooks calls the given callback function
// for each book in the collection
func (l *Library) IterateBooks(f func(Book) error) {
	var err error
	for _, b := range l.Collection {
		err = f(b)
		if err != nil {
			fmt.Println("Error encountered:", err)
			break
		}
	}
}

// createIterator returns a BookIterator that can access the book
// collection on demand
func (l *Library) createIterator() iterator {
	return &BookIterator{
		books: l.Collection,
	}
}

// -------------------
// Create a Library structure to hold a set of Books
var lib *Library = &Library{
	Collection: []Book{
		{
			Name:      "Buku Bahasa Indonesia kelas 7",
			Author:    "Tatang sutatang",
			PageCount: 122,
			Type:      HardCover,
		},
		{
			Name:      "Buku MTK Bangun Datar SD",
			Author:    "Budi doremi",
			PageCount: 50,
			Type:      SoftCover,
		},
		{
			Name:      "Kisah bobo 1999",
			Author:    "Ka Seto",
			PageCount: 20,
			Type:      PaperBack,
		},

		{
			Name:      "Design Pattern 2030",
			Author:    "Lino Thomas",
			PageCount: 288,
			Type:      EBook,
		},
	},
}
