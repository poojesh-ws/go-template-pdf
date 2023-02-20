package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"go-template/internal/service/generatePdf"

	"github.com/google/uuid"
)

var TEMPLATE_PATH = "pdftemplate/test.html"

// GeneratePDF is the resolver for the generatePdf field.
func (r *mutationResolver) GeneratePDF(ctx context.Context) (string, error) {
	fileName := fmt.Sprintf("outputPdf/%s_%s.pdf", "testPdf", uuid.New().String())
	htmlBody, err := generatePdf.ParseTemplate(TEMPLATE_PATH)

	if err != nil {
		return "", err
	}

	go generatePdf.GeneratePDF(fileName, htmlBody)

	return fmt.Sprintf("Pdf generating for name %s", fileName), nil
}
