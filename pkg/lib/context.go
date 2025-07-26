package lib

import (
	"context"

	"github.com/bagusyanuar/go-erp/pkg/constant"
	"github.com/google/uuid"
)

func GetUserIDSafe(ctx context.Context) (uuid.UUID, bool) {
	val := ctx.Value(constant.UserIDKey)
	userID, ok := val.(uuid.UUID)
	return userID, ok
}
