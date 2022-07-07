package service

import (
	"context"
	entry "notes/internal/entry/domain"
)

type EntryService struct {
	repository entry.EntryRepository
}

func NewEntryService(repository entry.EntryRepository) EntryService {
	return EntryService{repository}
}

func (s EntryService) QueryEntries(ctx context.Context, userId string) ([]entry.Entry, error) {
	entries, err := s.repository.FindEntries(ctx, userId)
	return entries, err
}

func (s EntryService) QueryEntry(ctx context.Context, userId string, entryId string) (*entry.Entry, error) {
	entry, err := s.repository.FindEntryById(ctx, userId, entryId)
	return entry, err
}

func (s *EntryService) AddNewEntry(ctx context.Context, userId string, content string) error {
	err := s.repository.InsertEntry(ctx, userId, content)
	return err
}

func (s *EntryService) EditEntry(ctx context.Context, userId string, entryId string, content string) error {
	err := s.repository.UpdateEntry(ctx, userId, entryId, content)
	return err
}

func (s *EntryService) DeleteEntry(ctx context.Context, userId string, entryId string) error {
	err := s.repository.DeleteEntry(ctx, userId, entryId)
	return err
}
