package contexts

import (
	"context"
	"os/user"
)

func GetSessionUser(ctx context.Context) *user.User {
	if sessionUser, ok := ctx.Value("session-user").(*user.User); ok {
		return sessionUser
	}
	return nil
}
