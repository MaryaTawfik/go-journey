package main
import(
	"Library_Management/controllers"
	"Library_Management/models"
	"Library_Management/services"
)


func main() {
	library:=services.Library{
AllBooks:make(map[int]models.Book),
AllMember:make(map[int]models.Member),

	}
	member1 := models.Member{
		Id:            1,
		Name:          "Alice",
		BorrowedBook:  []models.Book{},
	}
	member2 := models.Member{
		Id:            1,
		Name:          "Marya",
		BorrowedBook:  []models.Book{},
	}

	book1:=models.Book{
		BookId: 123,
Title :"love ur self",
Author:" marya",
Status :"Available",
	}

	
	book2:=models.Book{
		BookId: 321,
Title :"ego",
Author:" ahmed",
Status :"Available",
	}
	library.AllMember[member1.Id] = member1
	library.AllMember[member2.Id] = member2
	library.AllBooks[book1.BookId]=book1
	library.AllBooks[book2.BookId]=book2
	controllers.DisplayMenu(&library)
}
