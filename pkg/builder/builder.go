package builder

import (
	"context"

	"github.com/mxpv/podsync/pkg/feed"
	"github.com/pkg/errors"

	"github.com/mxpv/podsync/pkg/model"
)

type Builder interface {
	Build(ctx context.Context, cfg *feed.Config) (*model.Feed, error)
}

func New(ctx context.Context, provider model.Provider, key string, downloader Downloader) (Builder, error) {
	switch provider {
	case model.ProviderYoutube:
		return NewYouTubeBuilder(key, downloader)
	case model.ProviderVimeo:
		return NewVimeoBuilder(ctx, key)
	case model.ProviderSoundcloud:
		return NewSoundcloudBuilder()
	case model.ProviderTwitch:
		return NewTwitchBuilder(key)
	// 新增：注册 Bilibili 数据源
    	case model.ProviderBilibili:
        	return NewBilibiliBuilder()  // 引用 bilibili.go 中的构造函数
	default:
		return nil, errors.Errorf("unsupported provider %q", provider)
	}
}
