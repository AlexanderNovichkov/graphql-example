package converter

import (
	"github.com/AlexanderNovichkov/graphql-example/internal/db/model"
	"github.com/AlexanderNovichkov/graphql-example/internal/graph/gqlmodel"
)

func PlayerInputToPlayerDB(playerInput *gqlmodel.PlayerInput) *model.Player {
	return &model.Player{
		Name:    playerInput.Name,
		Role:    playerInput.Role.String(),
		IsAlive: playerInput.IsAlive,
	}
}

func PlayersInputToPlayersDB(playersInput []*gqlmodel.PlayerInput) []*model.Player {
	players := make([]*model.Player, len(playersInput))
	for i, playerInput := range playersInput {
		players[i] = PlayerInputToPlayerDB(playerInput)
	}
	return players
}

func PlayerDBToPlayer(playerDB *model.Player) *gqlmodel.Player {
	return &gqlmodel.Player{
		Name:    playerDB.Name,
		Role:    gqlmodel.Role(playerDB.Role),
		IsAlive: playerDB.IsAlive,
	}
}

func PlayersDBToPlayers(playersDB []*model.Player) []*gqlmodel.Player {
	players := make([]*gqlmodel.Player, len(playersDB))
	for i, playerDB := range playersDB {
		players[i] = PlayerDBToPlayer(playerDB)
	}
	return players
}

func CommentDBToComment(commentDB *model.Comment) *gqlmodel.Comment {
	return &gqlmodel.Comment{
		Text: commentDB.Text,
	}
}

func CommentsDBToComments(commentsDB []*model.Comment) []*gqlmodel.Comment {
	comments := make([]*gqlmodel.Comment, len(commentsDB))
	for i, commentDB := range commentsDB {
		comments[i] = CommentDBToComment(commentDB)
	}
	return comments
}

func CommentInputToCommentDB(commentInput *gqlmodel.CommentInput) *model.Comment {
	return &model.Comment{
		Text: commentInput.Text,
	}
}

func CommentsInputToCommentsDB(commentsInput []*gqlmodel.CommentInput) []*model.Comment {
	comments := make([]*model.Comment, len(commentsInput))
	for i, commentInput := range commentsInput {
		comments[i] = CommentInputToCommentDB(commentInput)
	}
	return comments
}

func GameDBToGame(gameDB *model.Game) *gqlmodel.Game {
	return &gqlmodel.Game{
		ID:       gameDB.ID.String(),
		Players:  PlayersDBToPlayers(gameDB.Players),
		Comments: CommentsDBToComments(gameDB.Comments),
		Status:   gqlmodel.GameStatus(gameDB.Status),
	}
}

func GamesDBToGames(gamesDB []*model.Game) []*gqlmodel.Game {
	games := make([]*gqlmodel.Game, len(gamesDB))
	for i, gameDB := range gamesDB {
		games[i] = GameDBToGame(gameDB)
	}
	return games
}
