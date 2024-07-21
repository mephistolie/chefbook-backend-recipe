package grpc

import "github.com/mephistolie/chefbook-backend-recipe/internal/config"

type Repository struct {
	Profile    *Profile
	Tag        *Tag
	Encryption *Encryption
}

func NewRepository(cfg *config.Config) (*Repository, error) {
	profileService, err := NewProfile(*cfg.ProfileService.Addr)
	if err != nil {
		return nil, err
	}
	tagService, err := NewTag(*cfg.TagService.Addr)
	if err != nil {
		return nil, err
	}
	encryptionService, err := NewEncryption(*cfg.EncryptionService.Addr)
	if err != nil {
		return nil, err
	}

	return &Repository{
		Profile:    profileService,
		Tag:        tagService,
		Encryption: encryptionService,
	}, nil
}

func (r *Repository) Stop() error {
	_ = r.Profile.Conn.Close()
	_ = r.Tag.Conn.Close()
	return nil
}
