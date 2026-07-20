package controllers

import (
	"Library_Management/models"
	"Library_Management/services" 
	"fmt"
	
)
func AddBookController(library services.LibraryManager) {
	fmt.Println("enter bookID")
	var bookid int
	fmt.Scan(& bookid)
	fmt.Println("enter book name")
	title:=""
	fmt.Scan(& title)
	fmt.Println("enter the name of the author")
	author:=""
	fmt.Scan(& author)
	book:=models.Book{
		BookId : bookid,
		Title : title,
		Author : author,
		Status : "",
			}
		err :=library.AddBook(book)
		if err != nil{
			fmt.Printf("Failed to add book: %v\n", err)
			return}
			fmt.Println("You added a book successfully!")

		
}

func RemoveBookController(library services.LibraryManager)  {
	fmt.Println("enter book id")
	var id int
	fmt.Scanln(&id)
	err:=library.RemoveBook(id)
	if err != nil{
		fmt.Printf("failed to remove a book : %v\n",err)
		return
	}
fmt.Println("book removed successfully")
	
}

func BorrowBookController(library services.LibraryManager)  {
	fmt.Println("enter your id")
	var memberid int
	fmt.Scanln(&memberid)
	fmt.Println("enter book id")
	var bookid int
	fmt.Scanln(&bookid)
	 err := library.BorrowBook(bookid,memberid)
	 if err != nil{
		fmt.Printf("task is faild:%v\n",err)
		return
	 }
	 fmt.Println("task completed successfully")
	
}
func ReturnBookController(library services.LibraryManager){
	fmt.Println("enter your id")
	var memberid int
	fmt.Scanln(&memberid)
	fmt.Println("enter book id")
	var bookid int
	fmt.Scanln(&bookid)
	err:=library.ReturnBook(bookid,memberid)
	if err != nil{
		fmt.Printf("task is faild:%v\n",err)
		return
	 }
	 fmt.Println("task completed successfully")
	

}

func ListAvailableBookController(library services.LibraryManager)  {
	book:=library.ListAvailableBooks()
	if len(book)==0{
		fmt.Println("No books are currently available in the library.")
		return
	}
	for _,b := range book{
		fmt.Printf("ID:%d   |  Title:%s   |  Author:%s\n",b.BookId,b.Title,b.Author )
	}

}

func ListBorrowedBook(library services.LibraryManager)  {
	fmt.Println("enter your id")
	var memberid int
	fmt.Scanln(&memberid)
	books,err := library.ListBorrowedBooks(memberid)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	if len(books) == 0 {
		fmt.Println("This member hasn't borrowed any books.")
		return
	}

	
	for _, book := range books {
		fmt.Printf("ID: %d | Title: %s | Author: %s\n", book.BookId, book.Title, book.Author)
	}
	
}

func DisplayMenu(library services.LibraryManager) {
exit := false
for !exit{
	fmt.Println("\n--- LIBRARY MANAGEMENT SYSTEM ---")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books")
		fmt.Println("7. Exit")
		fmt.Print("Enter your choice: ")
	var c int
	fmt.Scanln(&c)
	switch c {
	case 1:
		AddBookController(library)
	case 2:
		RemoveBookController(library)
	case 3:
		BorrowBookController(library)
	case 4:
		ReturnBookController(library)
	case 5:
		ListAvailableBookController(library)
	case 6:
		ListBorrowedBook(library)
	case 7:
		exit=true
	default:
		fmt.Println("Invalid choice")
	}

	
			

	
}
fmt.Println("Goodbye!")}