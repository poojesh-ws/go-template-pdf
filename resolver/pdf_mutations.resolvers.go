package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"go-template/internal/service/generatePdf"
)

var TEMPLATE_PATH = "pdftemplate/test.html"

// GeneratePDF is the resolver for the generatePdf field.
func (r *mutationResolver) GeneratePDF(ctx context.Context) (string, error) {

	htmlBody, err := generatePdf.ParseTemplate(TEMPLATE_PATH)

	if err != nil {
		return "", err
	}

	err = generatePdf.GeneratePDF(htmlBody)

	if err != nil {
		return "", err
	}

	return "Pdf generated", nil
}
