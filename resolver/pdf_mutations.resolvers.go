package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"go-template/internal/service/generatePdf"
	"sync"
)

var TEMPLATE_PATH = "pdftemplate/test.html"

// GeneratePDF is the resolver for the generatePdf field.
func (r *mutationResolver) GeneratePDF(ctx context.Context) (string, error) {

	exit := make(chan struct{})
	var wg sync.WaitGroup

	go func() {
		for i := 1; i <= 20; i++ {
			wg.Add(1)
			go func(i int) {

				htmlBody, err := generatePdf.ParseTemplate(TEMPLATE_PATH)

				if err != nil {
					fmt.Println(err)
				} else {
					err := generatePdf.GeneratePDF(htmlBody)

					if err != nil {
						fmt.Println(err)
					} else {
						fmt.Println("generated pdf", i)
					}
				}

				defer wg.Done()

			}(i)
		}
		wg.Wait()
		close(exit)
	}()

	for range exit {
		break
	}
	return "Pdf generated", nil
}
