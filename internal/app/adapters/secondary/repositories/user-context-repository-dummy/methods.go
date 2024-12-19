package user_context_repository_dummy

import "github.com/n-kazachuk/go_tg_bot/internal/app/domain/model"

func (c *UserContextRepository) GetContext(userId int64) (*model.UserContext, error) {
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

func (c *UserContextRepository) ClearContext(userID int64) error {
	delete(c.storage, userID)
	return nil
}
