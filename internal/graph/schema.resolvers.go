package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/AlexanderNovichkov/graphql-example/internal/db/model"
	"github.com/AlexanderNovichkov/graphql-example/internal/graph/converter"
	"github.com/AlexanderNovichkov/graphql-example/internal/graph/gqlgenerated"
	"github.com/AlexanderNovichkov/graphql-example/internal/graph/gqlmodel"
	"github.com/google/uuid"
)

func (r *mutationResolver) CreateGame(ctx context.Context, players []*gqlmodel.PlayerInput, comments []*gqlmodel.CommentInput, status gqlmodel.GameStatus) (*gqlmodel.Game, error) {
	Game := model.Game{
		ID:       uuid.New(),
		Players:  converter.PlayersInputToPlayersDB(players),
		Comments: converter.CommentsInputToCommentsDB(comments),
		Status:   status.String(),
	}

	err := r.Repository.CreateGame(&Game)
	if err != nil {
		return nil, err
	}

	return converter.GameDBToGame(&Game), nil
}

func (r *mutationResolver) UpdateGame(ctx context.Context, id string, players []*gqlmodel.PlayerInput, comments []*gqlmodel.CommentInput, status gqlmodel.GameStatus) (*gqlmodel.Game, error) {
	idParsed, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	Game := model.Game{
		ID:       idParsed,
		Players:  converter.PlayersInputToPlayersDB(players),
		Comments: converter.CommentsInputToCommentsDB(comments),
		Status:   status.String(),
	}

	err = r.Repository.UpdateGame(&Game)
	if err != nil {
		return nil, err
	}
	return converter.GameDBToGame(&Game), nil
}

func (r *mutationResolver) AddComment(ctx context.Context, id string, comment *gqlmodel.CommentInput) (*gqlmodel.Game, error) {
	idParsed, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	game, err := r.Repository.GetGame(idParsed)

	game.Comments = append(game.Comments, converter.CommentInputToCommentDB(comment))

	err = r.Repository.UpdateGame(game)
	if err != nil {
		return nil, err
	}

	return converter.GameDBToGame(game), nil
}

func (r *queryResolver) Game(ctx context.Context, id string) (*gqlmodel.Game, error) {
	idParsed, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	Game, err := r.Repository.GetGame(idParsed)
	if err != nil {
		return nil, err
	}

	return converter.GameDBToGame(Game), nil
}

func (r *queryResolver) Games(ctx context.Context, isFinished *bool) ([]*gqlmodel.Game, error) {
	Games, err := r.Repository.GetGames(isFinished)
	if err != nil {
		return nil, err
	}

	return converter.GamesDBToGames(Games), nil
}

// Mutation returns gqlgenerated.MutationResolver implementation.
func (r *Resolver) Mutation() gqlgenerated.MutationResolver { return &mutationResolver{r} }

// Query returns gqlgenerated.QueryResolver implementation.
func (r *Resolver) Query() gqlgenerated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
