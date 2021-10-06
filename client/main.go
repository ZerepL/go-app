package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/Fhoust/Go-app/goApp"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:3000"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	addUser := pb.NewAddUserClient(conn)
	getUser := pb.NewGetUserClient(conn)
	updateUser := pb.NewUpdateUserClient(conn)
	deleteUser := pb.NewDeleteUserClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()


	user, err := addUser.AddNewUser(ctx, &pb.User{Name: name})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Printf("Added %s", user.GetName())

	user, err = getUser.GetUserInfo(ctx, &pb.User{Id: user.GetId()})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Printf("The id of %s is %d", user.GetName(), user.GetId())

	user, err = updateUser.UpdateOneUser(ctx, &pb.User{Id: user.GetId(), Name: "New Potato"})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Printf("Updating %d to %s", user.GetId() ,"New Potato")

	user, err = getUser.GetUserInfo(ctx, &pb.User{Id: user.GetId()})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Printf("In database %d has the name %s", user.GetId(), user.GetName())

	user, err = deleteUser.DeleteOldUser(ctx, &pb.User{Id: user.GetId()})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Printf("%d was deleted", user.GetId())

	user, err = getUser.GetUserInfo(ctx, &pb.User{Id: user.GetId()})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Printf("Now database has this value for %d: %s", user.GetId(), user.GetName())
}
