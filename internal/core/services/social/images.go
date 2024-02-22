package social_service

import (
	"context"

	"github.com/google/uuid"
	social_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/social"
	deps_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/deps"
	social_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/social"
)

type imagesService struct {
	socialRepositories social_ports.ISocialRepositories
	deps               deps_ports.IDepsServices
}

func newImagesService(
	socialRepositories social_ports.ISocialRepositories,
	deps deps_ports.IDepsServices,
) social_ports.IImagesService {
	return &imagesService{
		socialRepositories: socialRepositories,
		deps:               deps,
	}
}

func (s *imagesService) UploadImage(ctx context.Context, image *social_domain.Image) (err error) {
	if err = s.deps.ValidatorService().ValidateStruct(image); err != nil {
		return
	}
	return s.socialRepositories.ImagesRepository().Create(ctx, image)
}

func (s *imagesService) DeleteImage(ctx context.Context, imageID uuid.UUID) (err error) {
	return s.socialRepositories.ImagesRepository().DeleteByID(ctx, imageID)
}

func (s *imagesService) UpdateImage(ctx context.Context, newImage *social_domain.Image) (err error) {
	if err = s.deps.ValidatorService().ValidateStruct(newImage); err != nil {
		return
	}
	return s.socialRepositories.ImagesRepository().Update(ctx, newImage)
}
