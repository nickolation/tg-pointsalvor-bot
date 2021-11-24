package service

import (
	"context"

	"github.com/nickolation/tg-pointsalvor-bot/engine/external/repository"
)

type Actor struct {
	Committer repository.Commits
}

func NewActor(com *repository.Commits) *Actor {
	return &Actor{
		Committer: *com,
	}
}

func (act *Actor) ActWrite(ctx context.Context, chatId int64, model string) error {
	if err := act.Committer.CommitWrite(ctx, chatId, model); err != nil {
		return err
	}

	return nil
}

func (act *Actor) ActView(ctx context.Context, chatId int64) error {

	return nil
}


func (act *Actor) ActClose(ctx context.Context, chatId int64) error {

	return nil
}


