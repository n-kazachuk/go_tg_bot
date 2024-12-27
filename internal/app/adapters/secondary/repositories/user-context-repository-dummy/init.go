package user_context_repository_dummy

import (
	"github.com/n-kazachuk/go_tg_bot/internal/app/domain/user-context"
	"sync"
)

type UserContextRepository struct {
	storage map[int64]*user_context.UserContext
	mu      sync.RWMutex
}

func NewUserContextRepository() *UserContextRepository {
	return &UserContextRepository{
		storage: make(map[int64]*user_context.UserContext),
	}
}
