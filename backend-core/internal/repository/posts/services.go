package posts

import (
	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/model"
	"github.com/gin-gonic/gin"
)

type PostService interface {
	UpdatePost(ctx *gin.Context, input model.Post) (model.Post, error)
}

type postService struct {
	repo PostRepo
}

func NewPostService(repo PostRepo) PostService {
	return &postService{
		repo: repo,
	}
}

func (s *postService) UpdatePost(ctx *gin.Context, input model.Post) (model.Post, error) {

	updatedPost, err := s.repo.EditPost(*ctx, input)
	if err != nil {
		return model.Post{}, err
	}

	return updatedPost, nil
}
