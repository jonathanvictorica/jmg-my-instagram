package services

type IMultimediaContentService interface {
	CreateMultimediaContent(multimediaContentRequest *MultimediaContentRequest) (string, error)
}

type MultimediaContentRequest struct {
}
