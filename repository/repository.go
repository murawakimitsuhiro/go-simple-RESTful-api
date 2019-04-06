package repository

import (
	"context"

	"github.com/murawakimitsuhiro/go-simple-RESTful-api/models"
)

type NoteRepo interface {
	Fetch(ctx context.Context, num uint) ([]models.Note, error)
	GetByID(ctx context.Context, id uint) (*models.Note, error)
	Create(ctx context.Context, n *models.Note) (uint, error)
	Update(ctx context.Context, n *models.Note) (*models.Note, error)
	Delete(ctx context.Context, id uint) (bool, error)
}
