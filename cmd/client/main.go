package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/AlexanderNovichkov/graphql-example/internal/client"
	"github.com/machinebox/graphql"
	"os"
)

type Action struct {
	Name    string
	Handler func()
}

func printMenu(actions []Action, writer *bufio.Writer) {
	for i, action := range actions {
		fmt.Fprintf(writer, "%d. %s\n", i+1, action.Name)
	}
	writer.Flush()
}

func handleInput(actions []Action, reader *bufio.Reader) {
	var actionId int
	_, err := fmt.Fscan(reader, &actionId)
	if err != nil || actionId < 1 || actionId > len(actions) {
		fmt.Println("Invalid input")
		return
	}
	actions[actionId-1].Handler()
}

var (
	endpoint = flag.String("endpoint", "http://localhost:8080/graphql", "graphql server endpoint")
)

func main() {
	flag.Parse()

	gqlClient := graphql.NewClient(*endpoint)
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	handler := client.NewHandler(gqlClient, reader, writer)

	var actions = []Action{
		{"Print finished games", handler.PrintFinishedGames},
		{"Print current games", handler.PrintCurrentGames},
		{"Print game scoreboard", handler.PrintGameScoreboard},
		{"Print game comments", handler.PrintGameComments},
		{"Add comment to game", handler.AddCommentToGame},
	}

	for {
		printMenu(actions, writer)
		handleInput(actions, reader)
		fmt.Fprintln(writer, "")
	}
}
