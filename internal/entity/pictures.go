package entity

import "github.com/google/uuid"

type PictureUpload struct {
	PictureId   uuid.UUID
	PictureLink string
	UploadUrl   string
	FormData    map[string]string
	MaxSize     int64
}

type RecipePictureIds struct {
	Preview *uuid.UUID
	Cooking map[uuid.UUID][]uuid.UUID
}

func (p *RecipePictureIds) GetIds() []uuid.UUID {
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

type RecipePictures struct {
	Preview *string
	Cooking map[uuid.UUID][]string
}
