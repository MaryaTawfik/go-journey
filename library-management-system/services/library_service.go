package services

import (
	"Library_Management/models"
	"fmt"
)

type Library struct {
	AllBooks  map[int]models.Book
	AllMember map[int]models.Member
}

// Add
func (library *Library) AddBook(book models.Book) error {

	id := book.BookId

	// accept the id of the book
	// check if the id in AllBooks
	_, exists := library.AllBooks[id]
	if exists {

		return fmt.Errorf("this book are in current book list")

		// if true send a message this book available
		//else
	}
	book.Status = "Available"
	library.AllBooks[id] = book
	return nil
}

// change status to available
//  add the book in AllBooks

// Remove Book
func (library *Library) RemoveBook(id int) error {

	// accept the id of the book
	// check if the id of the book in AllBooks
	val, exists := library.AllBooks[id]
	if !exists {
		return fmt.Errorf("the book are not found")
	} else if val.Status == "Borrowed" {
		return fmt.Errorf("it allready borrowed you canot delete")
	}
	delete(library.AllBooks, id)
	return nil
	// if it there
	// check if the status is borrowed
	// if borrowed through an error that says it allready borrowed you canot delete
	// else
	//  remove it
	// else through an error not found
}

// Borrow Book
// accept the id of the member
func (library *Library) BorrowBook(bookid int, memberid int) error {
	// check if the id in allmember
	val, exists := library.AllMember[memberid]

	if !exists {
		return fmt.Errorf("user not found")
	}

	// if no through an error not found
	// else:
	valbook, exist := library.AllBooks[bookid]
	if !exist {
		return fmt.Errorf("book not found")
	}
	if valbook.Status == "Borrowed" {
		return fmt.Errorf("this book is already borrowed ")
	}
	valbook.Status = "Borrowed"
	library.AllBooks[bookid] = valbook
	val.BorrowedBook = append(val.BorrowedBook, valbook)
	library.AllMember[memberid] = val
	return nil
	// check if the id of the book are found in AllBooks
	// if yes check the status
	// if status available
	// make the status borrowed
	// add the book in BorrowedBook []Book
	// else
	// through an error the book is already borrowed by some one else

	// else
	// though an error the book are not found
}

// Return Book
// accept user id
func (library *Library) ReturnBook(bookID int, memberID int) error {

	book, exist := library.AllBooks[bookID]
	if !exist {
		return fmt.Errorf("you try to retutn  inavailable book")
	} else if book.Status == "Available" {
		return fmt.Errorf("book is not borrowed")
	}
	member, exists := library.AllMember[memberID]
	//check if the book id is in the the borrowednook of the user
	// if it found
	if !exists {
		return fmt.Errorf("member not found")
	}
	// book.Status = "Available"
	// member, _ := library.AllMember[memberID]
	foundidx := -1
	for idx, book := range member.BorrowedBook {
		if book.BookId == bookID {
			foundidx = idx
			break
		}
	}

	if foundidx != -1 {
		member.BorrowedBook = append(member.BorrowedBook[:foundidx], member.BorrowedBook[foundidx+1:]...)
	} else {
		return fmt.Errorf("the book is not found in your borrowed book list")
	}
	book.Status = "Available"
	library.AllBooks[bookID] = book
	library.AllMember[memberID] = member
	// remove the ook from BorrowedBook []Book of the user
	// changing the status of the book to available in book
	// else
	// through an erro that says its not in your list
	return nil
}

// ListAvailableBooks()

func (library *Library) ListAvailableBooks() []models.Book {
	var availablebooks []models.Book
	for _, val := range library.AllBooks {
		if val.Status == "Available" {
			availablebooks = append(availablebooks, val)
		}
	}
	return availablebooks
}

// creating empty arry to stor available books
// loop in side AllBooks map[int]Book check if i.status === available
// if available add to the array
// else continue(skip)
// return the arry

// ListBorrowedBooks(memberID)
func (library *Library) ListBorrowedBooks(memberID int) ([]models.Book,error) {
	member, exists := library.AllMember[memberID]
	if !exists {
		return nil, fmt.Errorf("member not found")
	}
	return member.BorrowedBook,nil
}

// accept member id
// check id the id found
// if found
// return all lListBorrowedBooks(memberID)
// else
// trough an error not found
