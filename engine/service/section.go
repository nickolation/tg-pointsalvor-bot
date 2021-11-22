package service

import (
	"context"
	"log"

	sdk "github.com/nickolation/pointsalvor"
	"github.com/nickolation/tg-pointsalvor-bot/engine/external/repository"
)

type SectionService struct {
	repo *repository.Sections
}

func NewSectionService(repo *repository.Sections) *SectionService {
	return &SectionService{
		repo: repo,
	}
}

func (ss *SectionService) AddSection(ctx context.Context, token, name string, projId int) (*sdk.Section, error) {
	agent, err := sdk.NewAgent(token)
	if err != nil {
		log.Printf("error with linking the agent - [%s]", err.Error())
		return nil, err
	}

	section, err :=  agent.AddSection(ctx, sdk.NewSectionOpt{
		Project_id: projId,
		Name: name,
	})
	if err != nil {
		log.Printf("error with add the section - [%s]", err.Error())
		return nil, err
	}

	log.Printf("section is - [%v]", *section)
	return section, nil
}


func (ss *SectionService) DeleteSection() {

}
func (ss *SectionService) GetAllSections() {

}
func (ss *SectionService) GetLastSection() {

}
