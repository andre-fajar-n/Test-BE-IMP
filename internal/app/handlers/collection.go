package handlers

import (
	"context"
	"fmt"
	"test-be-IMP/gen/client"
	clientOpeartions "test-be-IMP/gen/client/operations"
	"test-be-IMP/gen/restapi/operations/coolection"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

func (h *handler) GetListCollections(ctx context.Context) (*coolection.ListCollectionOK, error) {
	data, err := h.getListCollectionFromClient(ctx)
	if err != nil {
		return nil, err
	}

	output := new(coolection.ListCollectionOK)
	for _, v := range data.Collections {
		output.Payload = append(output.Payload, &coolection.ListCollectionOKBodyItems0{
			Title:     v.Title,
			Thumbnail: v.Thumbnail,
		})
	}

	return output, nil
}

func (h *handler) getListCollectionFromClient(ctx context.Context) (*clientOpeartions.GetCatalogCollectionsOKBody, error) {
	// create the transport
	transport := httptransport.New("microdata.worldbank.org", "/index.php/api", nil)

	clientCollection := client.New(transport, strfmt.Default)

	resp, err := clientCollection.Operations.GetCatalogCollections(nil)
	if err != nil {
		fmt.Println("Error client:", err)
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("error client collection")
	}

	return resp.Payload, nil
}
