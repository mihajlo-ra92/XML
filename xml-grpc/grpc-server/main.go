package main

import (
	"context"
	"log"
	"net"

	"github.com/mihajlo-ra92/demo-grpc/invoicer"
	"google.golang.org/grpc"
)
type myInvoicerServer struct{
	invoicer.UnimplementedInvoicerServer
}
func (s myInvoicerServer) Create(ctx context.Context, req *invoicer.CreateRequest)(*invoicer.CreateResponse, error) {
	return &invoicer.CreateResponse{
		Pdf: []byte(req.From),
		Docx: []byte("test"),
	},nil
}

func main() {
	lis, err := net.Listen("tcp",":8089")
	if err != nil {
		log.Fatalf("Cannot create listener: %s", err)
	}
	serverRegister :=grpc.NewServer()
	service := &myInvoicerServer{}
	invoicer.RegisterInvoicerServer(serverRegister, service)
	err = serverRegister.Serve(lis)
	if err != nil {
		log.Fatalf("Impossible to serve: %s", err)
	}
}