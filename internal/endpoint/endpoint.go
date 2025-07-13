package endpoint

import (
	"context"

	"github.com/abhaypandey621/targeting-engine/internal/model"
	"github.com/abhaypandey621/targeting-engine/internal/service"
	"github.com/go-kit/kit/endpoint"
)

// MakeServeAdEndpoint creates the endpoint for campaign selection.
func MakeServeAdEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(model.AdRequest)
		if !ok {
			return nil, model.ErrBadRequest
		}
		resp, err := svc.ServeAd(ctx, &req)
		if err != nil {
			return nil, err
		}
		return model.AdResponse{Campaigns: resp}, nil
	}
}
