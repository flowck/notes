package entry

import "context"

type EntryRepository interface {
	InsertEntry(ctx context.Context, userId string, content string) error
	DeleteEntry(ctx context.Context, userId string, id string) error
	UpdateEntry(ctx context.Context, userId string, id string, content string) error
	FindEntryById(ctx context.Context, userId string, id string) (*Entry, error)
	FindEntries(ctx context.Context, userId string) ([]Entry, error)
}
