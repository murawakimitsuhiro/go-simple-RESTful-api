package note

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/murawakimitsuhiro/go-simple-RESTful-api/models"
	"github.com/murawakimitsuhiro/go-simple-RESTful-api/repository"
)

func NewGormRepo(Conn *gorm.DB) repository.NoteRepo {
	return &mysqlNoteRepo{
		Conn: Conn,
	}
}

type mysqlNoteRepo struct {
	Conn *gorm.DB
}

func (m *mysqlNoteRepo) Fetch(ctx context.Context, num uint) ([]models.Note, error) {
	var notes []models.Note

	err := m.Conn.Find(&notes).Error

	return notes, err
}

func (m *mysqlNoteRepo) GetByID(tx context.Context, id uint) (*models.Note, error) {
	note := &models.Note{}

	return note, m.Conn.First(note, id).Error
}

func (m *mysqlNoteRepo) Create(tx context.Context, n *models.Note) (uint, error) {
	return n.ID, m.Conn.Create(n).Error
}

func (m *mysqlNoteRepo) Update(tx context.Context, n *models.Note) (*models.Note, error) {
	note := models.Note{}
	m.Conn.First(&note, n.ID)

	n.CreatedAt = note.CreatedAt

	err := m.Conn.Save(n).Error

	return n, err
}

func (m *mysqlNoteRepo) Delete(tx context.Context, id uint) (bool, error) {
	if err := m.Conn.Where("ID = ?", id).Delete(models.Note{}).Error; err != nil {
		return false, err
	}

	return true, nil
}
