package repository

import (
	"encoding/json"
	"github.com/AlexanderNovichkov/graphql-example/internal/db/model"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateGame(game *model.Game) error {
	gameJson, err := json.Marshal(game)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(`
		INSERT INTO games (id, game)
		VALUES ($1 ,$2)
	`, game.ID, gameJson)

	return err
}

func (r *Repository) UpdateGame(game *model.Game) error {
	gameJson, err := json.Marshal(game)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(`
		UPDATE games
		SET game = $1
		WHERE id = $2
	`, gameJson, game.ID)
	return err
}

func (r *Repository) GetGame(id uuid.UUID) (*model.Game, error) {
	var gameJson []byte

	err := r.db.Get(&gameJson, `
		SELECT game
		FROM games
		WHERE id = $1
	`, id)

	var game model.Game

	err = json.Unmarshal(gameJson, &game)
	if err != nil {
		return nil, err
	}

	return &game, err
}

func (r *Repository) GetGames(isFinished *bool) ([]*model.Game, error) {
	query := `
		SELECT game
		FROM games
	`

	if isFinished != nil {
		if *isFinished {
			query += `
				WHERE (game ->> 'status') != 'NotFinished'
			`
		} else {
			query += `
				WHERE (game ->> 'status') = 'NotFinished'
			`
		}
	}

	var gamesJson [][]byte
	err := r.db.Select(&gamesJson, query)

	var games []*model.Game
	for _, gameJson := range gamesJson {
		var game model.Game
		err = json.Unmarshal(gameJson, &game)
		if err != nil {
			return nil, err
		}
		games = append(games, &game)
	}

	return games, err
}
