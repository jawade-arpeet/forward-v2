package service

import (
	"context"
	"errors"
	"fmt"
	"forward/internal/dao"
	"forward/internal/errs"
	"forward/internal/repository"
	"strings"
)

type PlatformService struct {
	platformRepo *repository.PlatformRepository
}

func newPlatformService(platformRepo *repository.PlatformRepository) *PlatformService {
	return &PlatformService{platformRepo: platformRepo}
}

func (s *PlatformService) CreatePlatform(
	ctx context.Context,
	name string,
	logoURL string,
) error {
	name = strings.TrimSpace(name)
	logoURL = strings.TrimSpace(logoURL)

	_, err := s.platformRepo.CreatePlatform(ctx, name, logoURL)
	if err != nil {
		if errors.Is(err, errs.ErrPgUniqueViolation) {
			return errs.Conflict(
				errs.CodePlatformAlreadyExists,
				"platform already exists",
				errs.ErrPlatformAlreadyExists,
			)
		}

		return errs.Internal(fmt.Errorf("failed to create platform: %w", err))
	}

	return nil
}

func (s *PlatformService) GetAllPlatforms(ctx context.Context) ([]dao.PlatformInfo, error) {
	platforms, err := s.platformRepo.GetAllPlatforms(ctx)
	if err != nil {
		return nil, errs.Internal(
			fmt.Errorf("failed to retrieve platforms: %w", err),
		)
	}

	return platforms, nil
}
