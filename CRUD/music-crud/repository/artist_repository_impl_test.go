package repository

import (
	"context"
	"fmt"
	music_crud "music-crud"
	"music-crud/music-crud/entity"
	"testing"

	_ "github.com/lib/pq"
)

func TestArtistRepositoryInsert(t *testing.T) {
	artistRepository := NewArtistRepository(music_crud.GetConnection())

	ctx := context.Background()
	artistEntity := entity.Artist{
		Name: "Ronald Prote",
	}

	result, err := artistRepository.Insert(ctx, artistEntity)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	artistRepository := NewArtistRepository(music_crud.GetConnection())

	artist, err := artistRepository.FindById(context.Background(), 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(artist)
}

func TestGetAll(t *testing.T) {
	artistRepository := NewArtistRepository(music_crud.GetConnection())

	list, err := artistRepository.GetAll(context.Background())
	if err != nil {
		panic(err)
	}
	for _, artist := range list {
		fmt.Println(artist)
	}
}
