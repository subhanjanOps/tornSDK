package forum

import (
	"context"
	"fmt"
	"net/url"

	"github.com/subhanjanOps/tornSDK/internal/rawapi"
)

type Service struct {
	client rawapi.Requester
}

func NewService(c rawapi.Requester) *Service {
	return &Service{client: c}
}

func (s *Service) GetForumAllThreads(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "forum/threads", query)
}

func (s *Service) GetForumCategories(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "forum/categories", query)
}

func (s *Service) GetForumGeneric(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "forum", query)
}

func (s *Service) GetForumLookup(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "forum/lookup", query)
}

func (s *Service) GetForumThread(ctx context.Context, threadId string, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, fmt.Sprintf("forum/%s/thread", threadId), query)
}

func (s *Service) GetForumThreadPosts(ctx context.Context, threadId string, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, fmt.Sprintf("forum/%s/posts", threadId), query)
}

func (s *Service) GetForumThreads(ctx context.Context, categoryIds string, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, fmt.Sprintf("forum/%s/threads", categoryIds), query)
}

func (s *Service) GetForumTimestamp(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "forum/timestamp", query)
}
