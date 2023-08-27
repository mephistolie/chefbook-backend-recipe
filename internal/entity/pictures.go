package entity

import "github.com/google/uuid"

type PictureUpload struct {
	PictureLink string
	UploadUrl   string
	FormData    map[string]string
	MaxSize     int64
}

type RecipePictures struct {
	Preview *string
	Cooking map[uuid.UUID][]string
}

type RecipePictureIds struct {
	Preview *uuid.UUID
	Cooking map[uuid.UUID][]uuid.UUID
}
