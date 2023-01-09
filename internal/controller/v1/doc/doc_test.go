package doc_controller

import (
	"context"
	"fmt"
	"log"
	"testing"

	pb "github.com/i-b8o/read-only_contracts/pb/reader/v1"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

func TestGetOne(t *testing.T) {
	assert := assert.New(t)

	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%s", "0.0.0.0", "30000"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}
	client := pb.NewDocGRPCClient(conn)
	defer conn.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tests := []struct {
		input    uint64
		expected *pb.GetOneDocResponse
		err      error
	}{
		{
			input:    1,
			expected: &pb.GetOneDocResponse{Name: "Имя первой записи"},
			err:      nil,
		},
		{
			input:    0,
			expected: nil,
			err:      status.Errorf(codes.NotFound, "id was not found: 0"),
		},
		{
			input:    9999999999999999999,
			expected: nil,
			err:      status.Errorf(codes.NotFound, "id was not found: 9999999999999999999"),
		},
	}

	for _, test := range tests {
		req := &pb.GetOneDocRequest{ID: test.input}
		e, err := client.GetOne(ctx, req)
		if err != nil {
			t.Log(err)
		}
		assert.True(proto.Equal(test.expected, e), fmt.Sprintf("GetOneDoc(%v)=%v want: %v", test.input, e, test.expected))
		assert.Equal(test.err, err)
	}

}
