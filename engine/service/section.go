package service

import "github.com/nickolation/tg-pointsalvor-bot/engine/external/repository"

type SectionService struct {
	repo *repository.Sections
}

func NewSectionService(repo *repository.Sections) *SectionService {
	return &SectionService{
		repo: repo,
	}
}

func (ss *SectionService) AddSection() {

}
func (ss *SectionService) DeleteSection() {

}
func (ss *SectionService) GetAllSections() {

}
func (ss *SectionService) GetLastSection() {

}
