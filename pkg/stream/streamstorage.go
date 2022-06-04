package stream

import (
	"context"
	"encoding/binary"
	"fmt"
	"log"

	"github.com/pulsar-sigs/bk-client-go/pkg/stream/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type StorageClient struct {
}

func (s *StorageClient) Hello() {
	m := map[string]string{}
	md := metadata.New(m)

	i := int64(0)
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(i))

	fmt.Println(string(b))

	//set scid to grpc metadata
	// md.Set("bk-rt-sc-id-bin", string(b))
	ctx := metadata.NewOutgoingContext(context.TODO(), md)

	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:4181", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	streamName := "/stream1"
	namespaceName := "default"

	rangeService := proto.NewRootRangeServiceClient(conn)
	// createResp, err := rangeService.CreateStream(ctx, &proto.CreateStreamRequest{
	// 	Name:   streamName,
	// 	NsName: namespaceName,
	// })
	// if err != nil {
	// 	panic(err)
	// }
	// log.Println("createResp:", createResp)

	resp, err := rangeService.GetStream(ctx, &proto.GetStreamRequest{
		Id: &proto.GetStreamRequest_StreamName{
			StreamName: &proto.StreamName{
				StreamName:    streamName,
				NamespaceName: namespaceName,
			},
		},
	})
	if err != nil {
		panic(err)
	}
	log.Println("GetStream:", resp)
	streamId := resp.GetStreamProps().StreamId

	metaRangeService := proto.NewMetaRangeServiceClient(conn)
	metaresp, err3 := metaRangeService.GetActiveRanges(ctx, &proto.GetActiveRangesRequest{
		StreamId: streamId,
	})

	if err3 != nil {
		panic(err3)
	}
	log.Println(metaresp)
	rangeId := metaresp.GetRanges()[0].GetProps().GetRangeId()

	tableService := proto.NewTableServiceClient(conn)
	putResp, err2 := tableService.Put(ctx, &proto.PutRequest{
		Key:   []byte("hello"),
		Value: []byte("world"),
		Header: &proto.RoutingHeader{
			StreamId: streamId,
			RangeId:  rangeId,
		},
	})
	tableService.Put(ctx, &proto.PutRequest{
		Key:   []byte("hello2"),
		Value: []byte("world"),
		Header: &proto.RoutingHeader{
			StreamId: streamId,
			RangeId:  rangeId,
		},
	})

	incResp, err4 := tableService.Increment(ctx, &proto.IncrementRequest{
		Key:    []byte("hello3"),
		Amount: 1,
		Header: &proto.RoutingHeader{
			StreamId: streamId,
			RangeId:  rangeId,
		},
	})
	if err4 != nil {
		panic(err4)
	}
	log.Println("incResp:", incResp)

	if err2 != nil {
		panic(err2)
	}
	log.Println(putResp)

	rangeResp, err4 := tableService.Range(ctx, &proto.RangeRequest{
		Key:      []byte("hello"),
		RangeEnd: []byte("hello3"),
		Header: &proto.RoutingHeader{
			StreamId: streamId,
			RangeId:  rangeId,
		},
	})
	if err4 != nil {
		panic(err4)
	}
	log.Println("rangeResp", rangeResp)

	for _, v := range rangeResp.Kvs {
		if v.IsNumber {
			log.Println(v.GetNumberValue())
			log.Println(string(v.GetValue()))
		}
	}

	// 1. getActiveRanges
	// 2. put
	// 3. getTableStore
	// 4. getStream
}
