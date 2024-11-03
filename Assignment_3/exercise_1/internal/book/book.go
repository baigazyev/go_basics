// internal/book/book.go
package book

// Book - структура для книги
type Book struct {
	ID     int
	Title  string
	Author string
	Year   int
}

// Хранилище для книг (вместо базы данных)
var books = make(map[int]Book)
var nextID = 1

// AddBook добавляет книгу в хранилище
func AddBook(title, author string, year int) Book {
	book := Book{
		ID:     nextID,
		Title:  title,
		Author: author,
		Year:   year,
	}
	books[nextID] = book
	nextID++
	return book
}

// GetBook возвращает книгу по ID
func GetBook(id int) (Book, bool) {
	book, exists := books[id]
	return book, exists
}

func GetAllBooks() []Book {
	var bookList []Book
	for _, book := range books {
		bookList = append(bookList, book)
	}
	return bookList
}

// UpdateBook обновляет книгу по ID
func UpdateBook(id int, title, author string, year int) (Book, bool) {
	book, exists := books[id]
	if exists {
		book.Title = title
		book.Author = author
		book.Year = year
		books[id] = book
	}
	return book, exists
}

// DeleteBook удаляет книгу по ID
func DeleteBook(id int) bool {
	_, exists := books[id]
	if exists {
		delete(books, id)
	}
	return exists
}
