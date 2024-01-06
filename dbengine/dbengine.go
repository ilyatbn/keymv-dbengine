package dbengine

import (
	"log"
	"os"
	"fmt"
	"github.com/gocql/gocql"
	dbengine "github.com/ilyatbn/keymv-proto/dbengine"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	// "google.golang.org/grpc/metadata"
)

type Server struct {
	dbengine.UnimplementedDBEngineServer
}

const (
    dbSession string = "db"
	Keyspace string = "keymv"
)

type User struct {
	id gocql.UUID
	org gocql.UUID
	role gocql.UUID
	password string
}

func (s *Server) GetUserLogonData(ctx context.Context, in *dbengine.EmailReq) (*dbengine.UserLogonData, error) {
	logger := log.New(os.Stdout, in.RequestId +" ", log.LstdFlags|log.Lmsgprefix)
	logger.Println("Received userinfo:"+in.Email)
	var user User
	session := ctx.Value(dbSession).(*gocql.Session)
	if session == nil {
		return nil, status.Error(codes.Internal, "no database connection found")
	}
	query := fmt.Sprintf("select org, role, userid, password from users WHERE email='%v'",in.Email)
	err := session.Query(query).Scan(&user.org, &user.role, &user.id, &user.password)
	if err != nil {
		logger.Printf("Error getting data from Scylla: %v", err)
		return nil, err
	}
	return &dbengine.UserLogonData{ Org: user.org.String(), Role: user.role.String(),Password: user.password, Id: user.id.String() }, nil
}


