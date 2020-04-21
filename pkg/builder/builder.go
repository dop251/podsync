package builder

import (
	"context"

	"github.com/pkg/errors"

	"github.com/dop251/podsync/pkg/config"
	"github.com/dop251/podsync/pkg/model"
)

type Builder interface {
	Build(ctx context.Context, cfg *config.Feed) (*model.Feed, error)
}

func New(ctx context.Context, provider model.Provider, key string) (Builder, error) {
	switch provider {
	case model.ProviderYoutube:
		return NewYouTubeBuilder(key)
	case model.ProviderVimeo:
		return NewVimeoBuilder(ctx, key)
	default:
		return nil, errors.Errorf("unsupported provider %q", provider)
	}
}
