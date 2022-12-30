package repository

import (
	"context"
	"music-crud/music-crud/entity"
)

type ArtistRepository interface {
	Insert(ctx context.Context, name entity.Artist) (entity.Artist, error)
	FindById(ctx context.Context, id int32) (entity.Artist, error)
	GetAll(ctx context.Context) ([]entity.Artist, error)
}
