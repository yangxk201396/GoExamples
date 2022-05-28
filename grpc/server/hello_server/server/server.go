package server

import (
	"crypto/tls"
	"log"
	"net"
	"net/http"
	"path"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/yang201396/GoExamples/grpc/pkg/util"
	pb "github.com/yang201396/GoExamples/grpc/proto"
)

var (
	Port           string
	CertServerName string
	CertPemPath    string
	CertKeyPath    string
	SwaggerDir     string
	EndPoint       string

	tlsConfig *tls.Config
)

func Run() (err error) {
	EndPoint = ":" + Port
	tlsConfig = util.GetTLSConfig(CertPemPath, CertKeyPath)

	conn, err := net.Listen("tcp", EndPoint)
	if err != nil {
		log.Printf("TCP Listen err:%v\n", err)
	}

	srv := newServer()

	log.Printf("gRPC and https listen on: %s\n", Port)

	if err = srv.Serve(util.NewTLSListener(conn, tlsConfig)); err != nil {
		log.Printf("ListenAndServe err: %v\n", err)
	}

	return err
}

func newServer() *http.Server {
	grpcServer := newGrpc()
	gwmux, err := newGateway()
	if err != nil {
		panic(err)
	}
	httpmux, err := newHttpHandler()
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", gwmux)
	mux.Handle("/http/", httpmux)
	mux.HandleFunc("/swagger/", serveSwaggerFile)
	mux.Handle("/openapiv2/", openAPIServer(SwaggerDir))

	return &http.Server{
		Addr:      EndPoint,
		Handler:   util.GrpcHandlerFunc(grpcServer, mux),
		TLSConfig: tlsConfig,
	}
}

func newGrpc() *grpc.Server {
	creds, err := credentials.NewServerTLSFromFile(CertPemPath, CertKeyPath)
	if err != nil {
		panic(err)
	}

	opts := []grpc.ServerOption{
		grpc.Creds(creds),
	}

	server := grpc.NewServer(opts...)

	pb.RegisterHelloWorldServer(server, NewHelloService())

	return server
}

func newGateway() (http.Handler, error) {
	ctx := context.Background()
	creds, err := credentials.NewClientTLSFromFile(CertPemPath, CertServerName)
	if err != nil {
		return nil, err
	}
	opts := []grpc.DialOption{grpc.WithTransportCredentials(creds)}

	gwmux := runtime.NewServeMux()
	if err := pb.RegisterHelloWorldHandlerFromEndpoint(ctx, gwmux, EndPoint, opts); err != nil {
		return nil, err
	}

	return gwmux, nil
}

func newHttpHandler() (http.Handler, error) {
	svr := pb.NewHelloWorldHTTPConverter(NewHelloService())
	return svr.SayHelloWorld(nil), nil
}

func serveSwaggerFile(w http.ResponseWriter, r *http.Request) {
	if !strings.HasSuffix(r.URL.Path, "swagger.json") {
		log.Printf("Not Found: %s", r.URL.Path)
		http.NotFound(w, r)
		return
	}

	p := strings.TrimPrefix(r.URL.Path, "/swagger/")
	p = path.Join(SwaggerDir, p)

	log.Printf("serveSwaggerFile path: %s", p)

	http.ServeFile(w, r, p)
}

func openAPIServer(dir string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasSuffix(r.URL.Path, ".swagger.json") {
			log.Printf("Not Found: %s", r.URL.Path)
			http.NotFound(w, r)
			return
		}

		log.Printf("openAPIServer path: %s", r.URL.Path)
		p := strings.TrimPrefix(r.URL.Path, "/openapiv2/")
		p = path.Join(dir, p)
		http.ServeFile(w, r, p)
	}
}
