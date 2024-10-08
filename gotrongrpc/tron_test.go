package gotrongrpc

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/neatjson/neatjsons"
)

func TestNewGrpcClient(t *testing.T) {
	client := mustNewClient(NileNetGrpc)

	nodeInfo, err := client.grpcClient.GetNodeInfo()
	require.NoError(t, err)
	t.Log(neatjsons.S(nodeInfo))
}

func mustNewClient(address string) *Client {
	grpcClient, err := NewGrpcClient(address)
	if err != nil {
		panic(errors.WithMessage(err, "wrong"))
	}
	return NewClient(grpcClient)
}
