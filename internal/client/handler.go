package client

import (
	"bufio"
	"context"
	"fmt"
	"github.com/AlexanderNovichkov/graphql-example/internal/graph/gqlmodel"
	"github.com/machinebox/graphql"
	"text/tabwriter"
)

type Handler struct {
	gqlClient *graphql.Client
	reader    *bufio.Reader
	writer    *bufio.Writer
}

func NewHandler(gqlClient *graphql.Client, reader *bufio.Reader, writer *bufio.Writer) *Handler {
	return &Handler{
		gqlClient: gqlClient,
		reader:    reader,
		writer:    writer,
	}
}

func (h *Handler) PrintFinishedGames() {
	h.printGames(true)
}

func (h *Handler) PrintCurrentGames() {
	h.printGames(false)
}

func (h *Handler) PrintGameScoreboard() {
	gameId, err := h.readGameId()
	if err != nil {
		return
	}

	req := graphql.NewRequest(`
		query($gameId: ID!) {
			game(id: $gameId) {
				id
				players {
					name
					role
					isAlive
				}
				status
			  }
		}
	`)

	req.Var("gameId", gameId)

	result := struct {
		Game gqlmodel.Game
	}{}

	if err := h.gqlClient.Run(context.Background(), req, &result); err != nil {
		fmt.Fprintln(h.writer, "Error:", err)
		h.writer.Flush()
		return
	}

	fmt.Fprintln(h.writer, "Game status:", result.Game.Status.String())

	tw := tabwriter.NewWriter(h.writer, 1, 1, 1, ' ', 0)
	fmt.Fprintln(tw, "Name\tRole\tIs alive\t")
	for _, player := range result.Game.Players {
		fmt.Fprintf(tw, "%s\t%s\t%t\t\n", player.Name, player.Role, player.IsAlive)
	}
	tw.Flush()

	h.writer.Flush()
}

func (h *Handler) PrintGameComments() {
	gameId, err := h.readGameId()
	if err != nil {
		return
	}

	req := graphql.NewRequest(`
  		query($gameId: ID!) {
		  	game(id: $gameId) {
				comments{text}
			}
		}
	`)

	req.Var("gameId", gameId)

	result := struct {
		Game gqlmodel.Game
	}{}

	if err := h.gqlClient.Run(context.Background(), req, &result); err != nil {
		fmt.Fprintln(h.writer, "Error:", err)
		h.writer.Flush()
		return
	}

	fmt.Fprintln(h.writer, "Comments:")
	for _, comment := range result.Game.Comments {
		fmt.Fprintln(h.writer, comment.Text)
	}
	h.writer.Flush()
}

func (h *Handler) AddCommentToGame() {
	gameId, err := h.readGameId()
	if err != nil {
		return
	}

	fmt.Fprintln(h.writer, "Enter commentText:")
	h.writer.Flush()
	var commentText string
	_, err = h.reader.ReadString('\n')
	commentText, err = h.reader.ReadString('\n')

	if err != nil {
		fmt.Fprintln(h.writer, "Error:", err)
		h.writer.Flush()
		return
	}

	req := graphql.NewRequest(`
  		mutation($gameId: ID!, $commentText: String!) {
  			addComment(id:$gameId, comment: {text: $commentText}) {
				id
  			}
		}
	`)

	req.Var("gameId", gameId)
	req.Var("commentText", commentText)

	if err := h.gqlClient.Run(context.Background(), req, &struct{}{}); err != nil {
		fmt.Fprintln(h.writer, "Error:", err)
		h.writer.Flush()
		return
	}
	fmt.Fprintln(h.writer, "Comment added")
	h.writer.Flush()
}

func (h *Handler) printGames(isFinished bool) {
	req := graphql.NewRequest(`
		query($isFinished: Boolean) {
		  games(isFinished: $isFinished) {
			id
		  }
		}
	`)

	req.Var("isFinished", isFinished)

	result := struct {
		Games []*gqlmodel.Game
	}{}

	if err := h.gqlClient.Run(context.Background(), req, &result); err != nil {
		fmt.Fprintln(h.writer, "Error:", err)
		h.writer.Flush()
		return
	}

	if isFinished {
		fmt.Fprintln(h.writer, "Finished games IDs:")
	} else {
		fmt.Fprintln(h.writer, "Current games IDs:")
	}
	for _, game := range result.Games {
		fmt.Fprintln(h.writer, game.ID)
	}
	h.writer.Flush()
}

func (h *Handler) readGameId() (string, error) {
	fmt.Fprintln(h.writer, "Enter game ID:")
	h.writer.Flush()

	var gameId string
	if _, err := fmt.Fscan(h.reader, &gameId); err != nil {
		fmt.Fprintln(h.writer, "Error:", err)
		h.writer.Flush()
		return "", err
	}

	return gameId, nil
}
