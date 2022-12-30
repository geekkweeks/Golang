package repository

import (
	"context"
	"database/sql"
	"errors"
	"music-crud/music-crud/entity"
	"strconv"
)

type artistRepositoryImpl struct {
	DB *sql.DB
}

func NewArtistRepository(db *sql.DB) ArtistRepository {
	return &artistRepositoryImpl{DB: db}
}

func (repository *artistRepositoryImpl) Insert(ctx context.Context, artist entity.Artist) (entity.Artist, error) {
	sqlScript := "INSERT INTO ARTISTS(NAME) VALUES($1)"
	_, err := repository.DB.ExecContext(ctx, sqlScript, artist.Name)
	if err != nil {
		panic(err)
	}
	return artist, err
}

func (repository *artistRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Artist, error) {
	sqlScript := "SELECT ID, NAME FROM ARTISTS WHERE ID = $1 LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, sqlScript, id)
	artistEntity := entity.Artist{}
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	if rows.Next() {
		rows.Scan(&artistEntity.Id, &artistEntity.Name)
		return artistEntity, nil
	} else {
		return artistEntity, errors.New("ID " + strconv.Itoa(int(id)) + " Not Found")
	}
}

func (repository *artistRepositoryImpl) GetAll(ctx context.Context) ([]entity.Artist, error) {
	sqlScript := "SELECT ID, NAME FROM ARTISTS"
	rows, err := repository.DB.QueryContext(ctx, sqlScript)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// artist collection
	var artists []entity.Artist
	for rows.Next() {
		artistEntity := entity.Artist{}
		rows.Scan(&artistEntity.Id, &artistEntity.Name)
		artists = append(artists, artistEntity)
	}
	return artists, nil
}
