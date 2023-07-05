package entity

import "github.com/google/uuid"

type PictureUpload struct {
	PictureId uuid.UUID
	URL       string
	FormData  map[string]string
	MaxSize   int64
}

type RecipePictures struct {
	Preview *uuid.UUID
	Cooking map[uuid.UUID][]uuid.UUID
}

func (p *RecipePictures) GetIds() []uuid.UUID {
	var ids []uuid.UUID

	if p.Preview != nil {
		ids = append(ids, *p.Preview)
	}

	if p.Cooking != nil {
		for _, pictures := range p.Cooking {
			ids = append(ids, pictures...)
		}
	}

	return ids
}
