package model

import "github.com/google/uuid"

type Comment struct {
	Text string `json:"text"`
}

type Game struct {
	ID       uuid.UUID  `json:"id"`
	Players  []*Player  `json:"players"`
	Comments []*Comment `json:"comments"`
	Status   string     `json:"status"`
}

type Player struct {
	Name    string `json:"name"`
	Role    string `json:"role"`
	IsAlive bool   `json:"isAlive"`
}
