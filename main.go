package main

import (
	"fmt"
	"log"
	"net"
	"runtime"
	"time"

	pb "github.com/DesmondANIMUS/greeter/greet"
	"github.com/DesmondANIMUS/greeter/greetpackages/greethelper"
	"github.com/DesmondANIMUS/greeter/greetpackages/greetlinker"
	"github.com/DesmondANIMUS/greeter/greetpackages/greetmodel"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type server struct{}

func handleErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	lis, err := net.Listen("tcp", greetmodel.Port)
	handleErr(err)

	greethelper.Session = greethelper.ConnectDB()
	defer greethelper.Session.Close()

	s := grpc.NewServer()
	registerServices(s)

	fmt.Println("Server listening at ", greetmodel.Port)
	handleErr(s.Serve(lis))
}

func registerServices(s *grpc.Server) {
	pb.RegisterPingServer(s, server{})
}

func (s server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	start := time.Now()
	var person greetmodel.PersonModel
	person.Name = in.Name

	go greetlinker.AddUserLinker(person)
	data := fmt.Sprintf("Hey, %s %v", in.Name, time.Since(start))
	fmt.Println(time.Since(start))
	return &pb.HelloReply{Reply: data}, nil
}
