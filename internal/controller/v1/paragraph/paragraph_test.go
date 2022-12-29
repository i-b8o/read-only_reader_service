package paragraph_controller

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
	client := pb.NewChapterGRPCClient(conn)
	defer conn.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tests := []struct {
		input    uint64
		expected *pb.GetOneChapterResponse

		err error
	}{
		{
			input:    1,
			expected: &pb.GetOneChapterResponse{ID: 1, Name: "Имя первой записи", Num: "I", DocID: 1, OrderNum: 1, Paragraphs: []*pb.ReaderParagraph{&pb.ReaderParagraph{ID: 1, Num: 1, HasLinks: true, Class: "any-class", Content: "Содержимое <a id=\"dst101675\"></a> первого <a href='11111/a3a3a3/111'>параграфа</a>", ChapterID: 1}, &pb.ReaderParagraph{ID: 2, Num: 2, HasLinks: true, IsTable: true, IsNFT: true, Class: "any-class", Content: "Содержимое второго <a href='372952/4e92c731969781306ebd1095867d2385f83ac7af/335104'>пункта 5.14</a> параграфа", ChapterID: 1}, &pb.ReaderParagraph{ID: 3, Num: 3, HasLinks: true, Class: "any-class", Content: "<a id='335050'></a>Содержимое третьего параграфа<a href='/document/cons_doc_LAW_2875/'>таблицей N 2</a>.", ChapterID: 1}}},
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
		e, err := client.GetOne(ctx, req)
		if err != nil {
			t.Log(err)
		}
		assert.True(proto.Equal(test.expected, e), fmt.Sprintf("GetChapter(%v)=%v \nwant: %v", test.input, e, test.expected))
		if e != nil {
			assert.True(proto.Equal(test.expected.Paragraphs[0], e.Paragraphs[0]))
			assert.True(proto.Equal(test.expected.Paragraphs[1], e.Paragraphs[1]))
			assert.True(proto.Equal(test.expected.Paragraphs[2], e.Paragraphs[2]))
		}
		assert.Equal(test.err, err)
	}

}

func TestGetAll(t *testing.T) {
	assert := assert.New(t)

	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%s", "0.0.0.0", "30000"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}
	client := pb.NewChapterGRPCClient(conn)
	defer conn.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tests := []struct {
		input    uint64
		expected *pb.GetAllChaptersByDocIdResponse
		err      error
	}{
		{
			input:    1,
			expected: &pb.GetAllChaptersByDocIdResponse{Chapters: []*pb.ReaderChapter{&pb.ReaderChapter{ID: 1, Name: "Имя первой записи", Num: "I", OrderNum: 1}, &pb.ReaderChapter{ID: 2, Name: "Имя второй записи", Num: "II", OrderNum: 2}, &pb.ReaderChapter{ID: 3, Name: "Имя третьей записи", Num: "III", OrderNum: 3}}},
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
		req := &pb.GetAllChaptersByDocIdRequest{ID: test.input}
		e, err := client.GetAll(ctx, req)
		if err != nil {
			t.Log(err)
		}
		assert.True(proto.Equal(test.expected, e), fmt.Sprintf("GetChapter(%v)=%v \nwant: %v", test.input, e, test.expected))
		assert.Equal(test.err, err)
	}

}
