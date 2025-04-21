package service

import "github.com/ScruffyPete/gologbook/internal/domain"

type EntryService struct {
	repo domain.EntryRepository
}

type CreateEntryInput struct {
	Body string `json:"body"`
}

// func (s *EntryService) List
