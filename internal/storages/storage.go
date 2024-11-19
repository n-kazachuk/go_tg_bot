package storages

import "github.com/n-kazachuk/go_tg_bot/internal/domain/model"

type ContextStorage interface {
	GetContext(userId int64) (*model.UserContext, error)
	ClearContext(userID int64) error
}

type Storage struct {
	ContextStorage ContextStorage
}

func New(contextStorage ContextStorage) *Storage {
	return &Storage{
		ContextStorage: contextStorage,
	}
}
