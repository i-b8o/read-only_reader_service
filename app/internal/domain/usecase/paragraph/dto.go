package usecase_paragraph

import "regulations_service/internal/domain/entity"

type CreateParagraphsInput struct {
	Paragraphs []entity.Paragraph
}

type CreateParagraphsOutput struct {
	Message string
}
