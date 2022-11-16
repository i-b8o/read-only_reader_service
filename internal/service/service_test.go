package service

import (
	"context"
	"fmt"
	"log"
	"testing"

	pb "github.com/i-b8o/read-only_contracts/pb/reader/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGetRegulation(t *testing.T) {
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%s", "0.0.0.0", "30000"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}
	client := pb.NewReaderGRPCClient(conn)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tests := []struct {
		id   uint64
		resp *pb.GetRegulationResponse
	}{
		{
			id:   1,
			resp: nil,
		},
		{
			id:   2,
			resp: &pb.GetRegulationResponse{Name: "Имя", Abbreviation: "Аббревиатура", Title: "Заголовок"},
		},
		{
			id:   3,
			resp: nil,
		},
	}

	for _, tt := range tests {
		getRegulationRequest := &pb.GetRegulationRequest{ID: tt.id}
		resp, err := client.GetRegulation(ctx, getRegulationRequest)
		if err != nil {
			t.Errorf("TestGetRegulation(%v) got unexpected error", err)
		}
		if resp != tt.resp {
			t.Errorf("GetRegulation(%v)=%v, wanted name |%v|", tt.id, resp, tt.resp)
		}
	}

}
