package main

import "fmt"

type Book struct {
	bookid int
	title  string
	author string
}

func main() {
	var book1 Book
	book1.bookid = 1
	book1.author = "Louse Van Guido"
	book1.title = "Intro to go"

	/* print Book1 info */
	fmt.Printf("Book 1 title : %s\n", book1.title)
	fmt.Printf("Book 1 author : %s\n", book1.author)
	fmt.Printf("Book 1 book_id : %d\n", book1.bookid)
}
