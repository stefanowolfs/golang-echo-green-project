package repository

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"green-echo/internal/app/green/domain"
)

type UserGorm struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

func (*UserGorm) TableName() string {
	return "users"
}

func (g *UserGorm) ToUser() *domain.User {
	return &domain.User{
		ID:   g.ID,
		Name: g.Name,
	}
}

func (g *UserGorm) FromUser(user *domain.User) {
	g.ID = user.ID
	g.Name = user.Name
}

type UserPostgresRepository struct {
	db *gorm.DB
}

var _ domain.UserRepository = &UserPostgresRepository{}

func NewUserPostgresRepository(db *gorm.DB) domain.UserRepository {
	return &UserPostgresRepository{
		db: db,
	}
}

func (r *UserPostgresRepository) GetByID(id uint) (*domain.User, error) {
	user := UserGorm{}

	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("database: user not found: %w", err)
		}
		return nil, fmt.Errorf("database: unexpected database error: %w", err)
	}

	return user.ToUser(), nil
}

func (r *UserPostgresRepository) NewUser(user *domain.User) (*domain.User, error) {
	entity := &UserGorm{}
	entity.FromUser(user)

	if err := r.db.Create(&entity).Error; err != nil {
		return nil, fmt.Errorf("database: unexpected database error: %w", err)
	}

	return entity.ToUser(), nil
}
