package controller

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	pb "github.com/i-b8o/read-only_contracts/pb/reader/v1"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestGetRegulation(t *testing.T) {
	assert := assert.New(t)

	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%s", "0.0.0.0", "30000"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}
	client := pb.NewReaderGRPCClient(conn)
	defer conn.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tests := []struct {
		input    uint64
		expected *pb.GetOneRegulationResponse
		err      error
	}{
		{
			input:    1,
			expected: &pb.GetOneRegulationResponse{Name: "Имя первой записи", Abbreviation: "Аббревиатура первой записи", Title: "Заголовок первой записи"},
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
		req := &pb.GetOneRegulationRequest{ID: test.input}
		e, err := client.GetOneRegulation(ctx, req)
		if err != nil {
			t.Log(err)
		}
		assert.True(proto.Equal(test.expected, e), fmt.Sprintf("GetOneRegulation(%v)=%v want: %v", test.input, e, test.expected))
		assert.Equal(test.err, err)
	}

}

func TestGetChapter(t *testing.T) {
	assert := assert.New(t)

	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%s", "0.0.0.0", "30000"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}
	client := pb.NewReaderGRPCClient(conn)
	defer conn.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	dateString := "2023-01-01"
	date, _ := time.Parse("2006-01-02", dateString)

	tests := []struct {
		input    uint64
		expected *pb.GetOneChapterResponse

		err error
	}{
		{
			input:    1,
			expected: &pb.GetOneChapterResponse{ID: 1, Name: "Имя первой записи", Num: "I", RegulationID: 1, OrderNum: 1, Paragraphs: []*pb.ReaderParagraph{&pb.ReaderParagraph{ID: 1, Num: 1, Class: "any-class", Content: "Содержимое первого параграфа", ChapterID: 1}, &pb.ReaderParagraph{ID: 2, Num: 2, HasLinks: true, IsTable: true, IsNFT: true, Class: "any-class", Content: "Содержимое второго параграфа", ChapterID: 1}, &pb.ReaderParagraph{ID: 3, Num: 3, Class: "any-class", Content: "Содержимое третьего параграфа", ChapterID: 1}}, UpdatedAt: timestamppb.New(date)},
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
		req := &pb.GetOneChapterRequest{ID: test.input}
		e, err := client.GetOneChapter(ctx, req)
		if err != nil {
			t.Log(err)
		}
		assert.True(proto.Equal(test.expected, e), fmt.Sprintf("GetChapter(%v)=%v \nwant: %v", test.input, e, test.expected))
		assert.Equal(test.err, err)
	}

}

func TestGetAllChapters(t *testing.T) {
	assert := assert.New(t)

	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%s", "0.0.0.0", "30000"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}
	client := pb.NewReaderGRPCClient(conn)
	defer conn.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tests := []struct {
		input    uint64
		expected *pb.GetAllChaptersByRegulationIdResponse
		err      error
	}{
		{
			input:    1,
			expected: &pb.GetAllChaptersByRegulationIdResponse{Chapters: []*pb.ReaderChapter{&pb.ReaderChapter{ID: 1, Name: "Имя первой записи", Num: "I", OrderNum: 1}, &pb.ReaderChapter{ID: 2, Name: "Имя второй записи", Num: "II", OrderNum: 2}, &pb.ReaderChapter{ID: 3, Name: "Имя третьей записи", Num: "III", OrderNum: 3}}},
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
			err:      status.Errorf(codes.Unknown, "9999999999999999999 is greater than maximum value for Int4"),
		},
	}

	for _, test := range tests {
		req := &pb.GetAllChaptersByRegulationIdRequest{ID: test.input}
		e, err := client.GetAllChaptersByRegulationId(ctx, req)
		if err != nil {
			t.Log(err)
		}
		assert.True(proto.Equal(test.expected, e), fmt.Sprintf("GetChapter(%v)=%v \nwant: %v", test.input, e, test.expected))
		assert.Equal(test.err, err)
	}

}
