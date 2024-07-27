package rpc

import (
	"context"
	"io"
	"os"

	pb "github.com/tcm5343/circular-dependency-detector/protos/filestream"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewClient(address string) (pb.FileStreamClient, *grpc.ClientConn, error) {
	conn, err := grpc.NewClient(
		"172.17.0.1:8080",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, nil, err
	}

	// send file and command to alloy-analyzer-service
	client := pb.NewFileStreamClient(conn)
	return client, conn, nil
}

func UploadFile(client pb.FileStreamClient, filePath string, command string) (pb.FileStream_UploadAndAnalyzeClient, error) {
	stream, err := client.UploadAndAnalyze(context.Background())
	if err != nil {
		return nil, err
	}

	alloyProps := &pb.AlloyProperties{
		Command: command,
	}
	if err := stream.Send(&pb.FileStreamRequest{
		Request: &pb.FileStreamRequest_Props{Props: alloyProps},
	}); err != nil {
		return nil, err
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	buf := make([]byte, 1024)
	seq := int32(0)
	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		if err := stream.Send(&pb.FileStreamRequest{
			Request: &pb.FileStreamRequest_Chunk{
				Chunk: &pb.FileChunk{
					Content:  buf[:n],
					Sequence: seq,
				},
			},
		}); err != nil {
			return nil, err
		}
		seq++
	}

	if err := stream.CloseSend(); err != nil {
		return nil, err
	}

	return stream, nil
}
