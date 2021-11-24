package service

import (
	"context"
	"log"

	sdk "github.com/nickolation/pointsalvor"
	"github.com/nickolation/tg-pointsalvor-bot/auth"
	"github.com/nickolation/tg-pointsalvor-bot/engine/external/repository"
)

type SectionService struct {
	auth auth.Auth
	repo repository.Sections
}

func NewSectionService(r repository.Sections, a auth.Auth) *SectionService {
	return &SectionService{
		repo: r,
		auth: a,
	}
}

func (ss *SectionService) AddSection(ctx context.Context, chatId int64, name string) error {
	opt, err := ss.auth.SignIn(ctx, chatId)
	if err != nil {
		//		test-log
		log.Printf("error with getting the token - [%s]", err.Error())

		return err
	}

	agent, err := sdk.NewAgent(opt.Token)
	if err != nil {
		log.Printf("error with linking the agent - [%s]", err.Error())
		return err
	}

	section, err :=  agent.AddSection(ctx, sdk.NewSectionOpt{
		Project_id: opt.ProjId,
		Name: name,
	})

	if err != nil {
		log.Printf("error with add the section - [%s]", err.Error())
		return err
	}

	if err := ss.repo.CasheSection(ctx, chatId, section); err != nil {
		//		test-log 
		log.Printf("error with add the section - [%s]", err.Error())

		return err
	}

	//		test-validation for section value 
	if section == nil {
		log.Printf("section is - [%v]", *section)
	}

	return  nil
}


func (ss *SectionService) DeleteSection() {

}

func (ss *SectionService) GetAllSections(ctx context.Context, chatId int64) ([]sdk.Section, error) {
	list, err :=  ss.repo.SelectAllSection(ctx, chatId)
	if err != nil {
		return nil, err
	}

	return list, nil
}


func (ss *SectionService) GetLastSection() {

}
