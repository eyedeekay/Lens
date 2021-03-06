package client

import (
	"context"
	"fmt"

	"github.com/RTradeLtd/Lens"
	pb "github.com/RTradeLtd/Lens/models"
	"google.golang.org/grpc"
)

// Client is a lens client used to make requests to the Lens gRPC server
type Client struct {
	GC *grpc.ClientConn
	IC pb.IndexerAPIClient
}

// NewClient is used to generate our lens client
func NewClient(opts *lens.ConfigOpts, insecure bool) (*Client, error) {
	apiURL := fmt.Sprintf("%s:%s", opts.API.IP, opts.API.Port)
	var (
		gconn *grpc.ClientConn
		err   error
	)
	if insecure {
		gconn, err = grpc.Dial(apiURL, grpc.WithInsecure())
	}
	if err != nil {
		return nil, err
	}
	indexerConn := pb.NewIndexerAPIClient(gconn)
	return &Client{
		GC: gconn,
		IC: indexerConn,
	}, nil
}

// SubmitIndexRequest is used to submit an index request to lens
func (c *Client) SubmitIndexRequest(ctx context.Context, req *pb.IndexRequest) (*pb.IndexResponse, error) {
	return c.IC.SubmitIndexRequest(ctx, req)
}

// SubmitSearchRequest is used to submit a request to search through lens
func (c *Client) SubmitSearchRequest(ctx context.Context, req *pb.SearchRequest) (*pb.SearchResponse, error) {
	return c.IC.SubmitSearchRequest(ctx, req)
}
