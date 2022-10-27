package main

import (
	"context"
	"log"

	desc "github.com/Nau077/golang-pet-first/pkg/note_v1"
	"google.golang.org/grpc"
)

const address = "localhost:50051"

func main() {
	con, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err.Error())
	}
	defer con.Close()

	client := desc.NewNoteServiceClient(con)

	res, err := client.CreateNote(context.Background(), &desc.CreateNoteRequest{
		NoteContent: &desc.NoteContent{
			Title:  "kfkf",
			Text:   "rr",
			Author: "kkk",
		},
	})
	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Println(res.Id)

}
