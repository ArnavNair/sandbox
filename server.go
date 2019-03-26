package main

import (
	"context"
	"flag"
	"log"
	"net"

	pb "github.com/cpjudge/proto/submission"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/testdata"
)

var (
	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile   = flag.String("cert_file", "", "The TLS cert file")
	keyFile    = flag.String("key_file", "", "The TLS key file")
	jsonDBFile = flag.String("json_db_file", "", "A json file containing a list of features")
	serverAddr = flag.String("server_addr", "172.17.0.1:10000", "The server address in the format of host:port")
)

type sandboxServer struct{}

func (s *sandboxServer) SubmitCode(ctx context.Context, submission *pb.Submission) (*pb.CodeStatus, error) {

	submissionCode := RunSandbox(
		submission.GetTestcasesPath(),
		submission.GetSubmissionPath(),
		submission.GetLanguage(),
		submission.GetSubmissionId())
	codeStatus := &pb.CodeStatus{}
	switch submissionCode {
	case 145:
		codeStatus.CodeStatus = pb.SubmissionStatus_COMPILATION_ERROR
	case 146:
		codeStatus.CodeStatus = pb.SubmissionStatus_TIME_LIMIT_EXCEEDED
	default:
		codeStatus.CodeStatus = pb.SubmissionStatus_TO_BE_EVALUATED
	}

	return codeStatus, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", *serverAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	if *tls {
		if *certFile == "" {
			*certFile = testdata.Path("server1.pem")
		}
		if *keyFile == "" {
			*keyFile = testdata.Path("server1.key")
		}
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			log.Fatalf("Failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterSandboxServer(grpcServer, &sandboxServer{})
	grpcServer.Serve(lis)
}
