package grpc

import "github.com/mephistolie/chefbook-backend-recipe/internal/config"

type Repository struct {
	Profile  *Profile
	Category *Category
	Tag      *Tag
}

func NewRepository(cfg *config.Config) (*Repository, error) {
	profileService, err := NewProfile(*cfg.ProfileService.Addr)
	if err != nil {
		return nil, err
	}
	categoryService, err := NewCategory(*cfg.CategoryService.Addr)
	if err != nil {
		return nil, err
	}
	tagService, err := NewTag(*cfg.CategoryService.Addr)
	if err != nil {
		return nil, err
	}

	return &Repository{
		Profile:  profileService,
		Category: categoryService,
		Tag:      tagService,
	}, nil
}

func (r *Repository) Stop() error {
	_ = r.Profile.Conn.Close()
	_ = r.Tag.Conn.Close()
	_ = r.Category.Conn.Close()
	return nil
}
