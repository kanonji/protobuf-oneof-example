package main

import (
	"fmt"
	"log"
	sample_v1 "protobuf-oneof/backend/gen/sample/v1"
)

func main() {
	// CreateFileのインスタンスを作成
	createFile := &sample_v1.CreateFile{
		Path:     "path/to/create",
		UploadId: "upload123",
		Etag:     "etag123",
		Sort:     1,
	}

	// UpdateFileのインスタンスを作成
	updateFile := &sample_v1.UpdateFile{
		Id:   2,
		Sort: 2,
	}

	// DeleteFileのインスタンスを作成
	deleteFile := &sample_v1.DeleteFile{
		Id: 1,
	}

	// FileOperationのインスタンスを作成
	fileOperations := []*sample_v1.FileOperation{
		{
			Operation: &sample_v1.FileOperation_Create{Create: createFile},
		},
		{
			Operation: &sample_v1.FileOperation_Update{Update: updateFile},
		},
		{
			Operation: &sample_v1.FileOperation_Delete{Delete: deleteFile},
		},
	}

	// Requestのインスタンスを作成
	request := &sample_v1.Request{
		EntityId:    1,
		Name:        "Sample",
		Description: "This is a sample request",
		Files:       fileOperations,
	}

	// Requestの内容を表示
	fmt.Printf("%+v\n", request)

	DecodeRequest(request)
}

func DecodeRequest(req *sample_v1.Request) {
	fmt.Println("Decoding Request...")

	for _, fileOp := range req.Files {
		switch op := fileOp.Operation.(type) {
		case *sample_v1.FileOperation_Create:
			fmt.Printf("Create File - Path: %s, Sort: %d, UploadId: %s, Etag: %s\n", op.Create.Path, op.Create.Sort, op.Create.UploadId, op.Create.Etag)
		case *sample_v1.FileOperation_Update:
			fmt.Printf("Update File - Id: %d, Sort: %d\n", op.Update.Id, op.Update.Sort)
		case *sample_v1.FileOperation_Delete:
			fmt.Printf("Delete File - Id: %d\n", op.Delete.Id)
		default:
			log.Println("Unknown operation")
		}
	}
}
