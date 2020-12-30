package main
import (
	"context"
	"fmt"
	"log"
	pb "github.com/pistatium/re_ark/protos"
	"google.golang.org/grpc"
)
func main() {
	//sampleなのでwithInsecure
	conn, err := grpc.Dial("127.0.0.1:19003", grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer conn.Close()
	client := pb.NewUserClient(conn)
	message := &pb.GetUserRequest{UserId: "Alice"}
	res, err := client.GetUser(context.TODO(), message)
	fmt.Printf("result:%#v \n", res)
	fmt.Printf("error::%#v \n", err)
}
