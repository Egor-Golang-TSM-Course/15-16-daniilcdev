package grpcserver

import (
	"context"
	"fmt"
	"server-context/gRpcServer/protoGenerated"
)

type Storage interface {
	Get(string) ([]byte, error)
	Save(string, []byte) error
}

func NewStorage() Storage {
	return &fileStorage{}
}

type fileStorageServiceServer struct {
	protoGenerated.UnimplementedFileStorageServiceServer
	fs Storage
}

type fileStorage struct {
}

func (fs *fileStorage) Get(name string) ([]byte, error) {
	fmt.Printf("getting file %s\n", name)

	return []byte{}, nil
}

func (fs *fileStorage) Save(name string, data []byte) error {
	fmt.Printf("saving file %s\n", name)
	return nil
}

func (fss *fileStorageServiceServer) Get(ctx context.Context, msg *protoGenerated.LoadFileMessage) (*protoGenerated.FileData, error) {
	b, err := fss.fs.Get(msg.GetName())
	if err != nil {
		return nil, err
	}

	return &protoGenerated.FileData{Data: b}, nil
}

func (fss *fileStorageServiceServer) Save(ctx context.Context, msg *protoGenerated.SaveFileMessage) (*protoGenerated.OperationResultMessage, error) {
	err := fss.fs.Save(msg.GetName(), msg.GetData())
	if err != nil {
		return nil, err
	}

	return &protoGenerated.OperationResultMessage{Result: protoGenerated.Result_Success}, nil

}
