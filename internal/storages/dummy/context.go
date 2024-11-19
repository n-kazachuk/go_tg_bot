package dummy

import (
	"github.com/n-kazachuk/go_tg_bot/internal/domain/model"
	"sync"
)

type ContextStorage struct {
	storage map[int64]*model.UserContext
	mu      sync.RWMutex
}

func NewContextStorage() *ContextStorage {
	return &ContextStorage{
		storage: make(map[int64]*model.UserContext),
	}
}

func (c *ContextStorage) GetContext(userId int64) (*model.UserContext, error) {
	context, exists := c.storage[userId]

	if exists {
		return context, nil
	}

	newContext := &model.UserContext{
		ActiveCommand: "",
		Data:          nil,
	}

	c.storage[userId] = newContext

	return newContext, nil
}

func (c *ContextStorage) ClearContext(userID int64) error {
	delete(c.storage, userID)
	return nil
}
