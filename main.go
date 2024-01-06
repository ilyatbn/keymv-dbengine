package main

import (
	"context"
	"log"
	"net"
	// "fmt"
	"github.com/gocql/gocql"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	db "github.com/ilyatbn/keymv-dbengine/dbengine"
	dbengine "github.com/ilyatbn/keymv-proto/dbengine"
	"google.golang.org/grpc"
)

const (
    dbSession string = "db"
	Keyspace string = "keymv"
)

func scyllasession() *gocql.Session {
	cluster := db.CreateCluster(gocql.Quorum, Keyspace,"localhost")
	session, err := gocql.NewSession(*cluster)
	if err != nil {
		log.Fatalf("unable to connect to scylla: %v", err)
	}
	return session
}

func main() {
	db_session := scyllasession()
	_, err:=db_session.KeyspaceMetadata(Keyspace)
	if err!=nil {
		log.Fatalf("Error in database connection check:%v",err)
	}
	
	listenPort := ":49010"
	lis, err := net.Listen("tcp4", listenPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	dbEngineServer := db.Server{}

	grpcServer := grpc.NewServer(
		grpc.ChainStreamInterceptor(
			DBStreamServerInterceptor(db_session),
		),
		grpc.ChainUnaryInterceptor(
			DBUnaryServerInterceptor(db_session),
		),
	)
	
	dbengine.RegisterDBEngineServer(grpcServer, &dbEngineServer)
	log.Printf("Listening on 0.0.0.0"+listenPort)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func DBUnaryServerInterceptor(session *gocql.Session) grpc.UnaryServerInterceptor {
    return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {	
		newctx:=context.WithValue(ctx, dbSession, session)
        return handler(newctx, req)
    }
}

func DBStreamServerInterceptor(session *gocql.Session) grpc.StreamServerInterceptor {
    return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
        wrapped := grpc_middleware.WrapServerStream(stream)
        wrapped.WrappedContext = context.WithValue(stream.Context(), dbSession, session)
        return handler(srv, wrapped)
    }
}