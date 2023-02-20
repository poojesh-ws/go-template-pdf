package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"go-template/gqlmodels"
	"os"
)

// CheckPDFGenerated is the resolver for the checkPdfGenerated field.
func (r *queryResolver) CheckPDFGenerated(ctx context.Context, fileName string) (*gqlmodels.PDFGeneratedStatus, error) {
	_, err := os.Stat(fileName)

	if err != nil {
		return nil, fmt.Errorf("Pdf does not exist")
	}

	return &gqlmodels.PDFGeneratedStatus{
		Status: "Pdf has been created",
	}, nil
}

// Query returns gqlmodels.QueryResolver implementation.
func (r *Resolver) Query() gqlmodels.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
