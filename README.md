# graphql-example

## Run server:

```bash
docker-compose up 
```

GraphQL playground will be available at: http://localhost:8080.

## Run client:

```bash
go run cmd/client/main.go 
```

## Examples

### Mutations

```graphql
mutation{
    createGame(
        players: [
            {
                name: "player_1",
                role: CIVILIAN,
                isAlive: true
            },
            {
                name: "player_2"
                role: MAFIA
                isAlive: false
            },
        ],
        comments: [
            {
                text: "Some comment"
            }
        ],
        status:MafiaWon
    ) {
        id
        players {
            name
            role
            isAlive
        }
        comments {
            text
        }
        status
    }
}
```

```graphql
mutation {
    addComment(
        id: "8ff5e8d3-8206-4201-9ac8-b97ad31e320f"
        comment: { text: "comment1" }
    ) {
        id
    }
}
```

### Queries

```graphql
query{
    games(isFinished:true) {
        id
        status
    }
}

```

```graphql
query {
    game(id: "8ff5e8d3-8206-4201-9ac8-b97ad31e320d") {
        id
        players {
            name
            role
            isAlive
        }
        status
    }
}
```


