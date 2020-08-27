/**
 * Author:  Nyxvectar Yan
 * Repo:    gocy
 * Created: 07/22/2025
 */

package block

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func Book() {
	var Book1 Books
	var Book2 Books

	Book1.title = "Go in Action"
	Book1.author = "William Kennedy"
	Book1.subject = "Golang Programming"
	Book1.book_id = 445353

	Book2.title = "Get Programming with Go"
	Book2.author = "Nathan Youngman"
	Book2.subject = "Golang Programming"
	Book2.book_id = 531421

	fmt.Printf("Book 1 title : %s\n", Book1.title)
	fmt.Printf("Book 1 author : %s\n", Book1.author)
	fmt.Printf("Book 1 subject : %s\n", Book1.subject)
	fmt.Printf("Book 1 book_id : %d\n", Book1.book_id)

	fmt.Printf("\nBook 2 title : %s\n", Book2.title)
	fmt.Printf("Book 2 author : %s\n", Book2.author)
	fmt.Printf("Book 2 subject : %s\n", Book2.subject)
	fmt.Printf("Book 2 book_id : %d\n", Book2.book_id)
}
