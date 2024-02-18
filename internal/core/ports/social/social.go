package social_ports

import (
	"context"

	"github.com/google/uuid"
	social_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/social"
)

type IPostsRepository interface {

	// Post Table Commands //
	Create(ctx context.Context, post *social_domain.Post) (err error)
	Update(ctx context.Context, newPost *social_domain.Post) (err error)
	DeleteByUUID(ctx context.Context, postUUID uuid.UUID) (err error)
	// End Post Table Commands //

	// Post Table Queries //

	// End Post Table Queries //

}
