package repositories

import (
	"database/sql"
	"github.com/Ruchanov/Golang_2023/assignment3/models"
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db}
}
func (r *BookRepository) GetBooks() []models.Book {
	var books []models.Book
	r.db.Find(&books)
	return books
}
func (r *BookRepository) GetBookByID(id int) (*models.Book, error) {
	var book models.Book
	err := r.db.Where("id = ?", id).First(&book).Error
	if err != nil {
		return &models.Book{}, err
	}
	return &book, nil
}
func (r *BookRepository) UpdateBookByID(id int, book *models.Book) (models.Book, error) {
	var updatedBook models.Book
	err := r.db.Model(&models.Book{}).Where("id = ?", id).Updates(book).First(&updatedBook).Error
	if err != nil {
		return models.Book{}, err
	}
	return updatedBook, nil
}
func (r *BookRepository) DeleteBook(book *models.Book) error {
	res := r.db.Delete(&book, book.Id)
	return res.Error
}
func (r *BookRepository) SearchByTitle(title string) ([]models.Book, error) {
	var books []models.Book
	err := r.db.Where("title LIKE ?", "%"+title+"%").Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}
func (r *BookRepository) CreateBook(book *models.Book) error {
	res := r.db.Model(&models.Book{}).Create(&book)
	return res.Error
}
func (r *BookRepository) GetBooksSortedByCost(order string) ([]models.Book, error) {
	var books []models.Book
	err := r.db.Order("cost " + order).Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func Iterator(rows *sql.Rows) (*[]models.Book, error) {
	var books []models.Book
	for rows.Next() {
		var book models.Book
		err := rows.Scan(&book.Id, &book.Title, &book.Description, &book.Cost)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return &books, nil
}
