package app

import (
	"context"
	"fmt"
	"net"
	postgressql "read-only_reader_service/internal/adapters/db/postgresql"
	"read-only_reader_service/internal/config"
	chapter_controller "read-only_reader_service/internal/controller/v1/chapter"
	paragraph_controller "read-only_reader_service/internal/controller/v1/paragraph"
	regulation_controller "read-only_reader_service/internal/controller/v1/regulation"
	service "read-only_reader_service/internal/domain/service"
	chapter_usecase "read-only_reader_service/internal/domain/usecase/chapter"
	paragraph_usecase "read-only_reader_service/internal/domain/usecase/paragraph"
	regulation_usecase "read-only_reader_service/internal/domain/usecase/regulation"

	"read-only_reader_service/pkg/client/postgresql"
	"time"

	pb "github.com/i-b8o/read-only_contracts/pb/reader/v1"

	"github.com/i-b8o/logging"
	"google.golang.org/grpc"
)

type App struct {
	cfg        *config.Config
	grpcServer *grpc.Server
	logger     logging.Logger
}

func NewApp(ctx context.Context, config *config.Config) (App, error) {
	logger := logging.GetLogger(config.AppConfig.LogLevel)

	logger.Print("Postgres initializing")
	pgConfig := postgresql.NewPgConfig(
		config.PostgreSQL.Username, config.PostgreSQL.Password,
		config.PostgreSQL.Host, config.PostgreSQL.Port, config.PostgreSQL.Database,
	)

	pgClient, err := postgresql.NewClient(context.Background(), 5, time.Second*5, pgConfig)
	if err != nil {
		logger.Fatal(err)
	}
	regAdapter := postgressql.NewRegulationStorage(pgClient)
	chapterAdapter := postgressql.NewChapterStorage(pgClient)
	paragraphAdapter := postgressql.NewParagraphStorage(pgClient)

	regulationService := service.NewRegulationService(regAdapter)
	chapterService := service.NewChapterService(chapterAdapter)
	paragraphService := service.NewParagraphService(paragraphAdapter)

	regulationUsecase := regulation_usecase.NewRegulationUsecase(regulationService)
	chapterUsecase := chapter_usecase.NewChapterUsecase(chapterService, paragraphService)
	paragraphUsecase := paragraph_usecase.NewParagraphUsecase(paragraphService)

	regulationController := regulation_controller.NewRegulationGRPCService(regulationUsecase)
	chapterController := chapter_controller.NewChapterGRPCService(chapterUsecase)
	paragraphController := paragraph_controller.NewParagraphGRPCService(paragraphUsecase)
	// read ca's cert, verify to client's certificate
	// homeDir, err := os.UserHomeDir()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// caPem, err := ioutil.ReadFile(homeDir + "/certs/ca-cert.pem")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // create cert pool and append ca's cert
	// certPool := x509.NewCertPool()
	// if !certPool.AppendCertsFromPEM(caPem) {
	// 	log.Fatal(err)
	// }

	// // read server cert & key
	// serverCert, err := tls.LoadX509KeyPair(homeDir+"/certs/server-cert.pem", homeDir+"/certs/server-key.pem")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // configuration of the certificate what we want to
	// conf := &tls.Config{
	// 	Certificates: []tls.Certificate{serverCert},
	// 	ClientAuth:   tls.RequireAndVerifyClientCert,
	// 	ClientCAs:    certPool,
	// }

	// //create tls certificate
	// tlsCredentials := credentials.NewTLS(conf)

	// grpcServer := grpc.NewServer(grpc.Creds(tlsCredentials))
	// pb.RegisterReadOnlyRegulationGRPCServer(grpcServer, regulationGrpcService)
	logger.Print("grpc server initializing")
	grpcServer := grpc.NewServer()
	pb.RegisterRegulationGRPCServer(grpcServer, regulationController)
	pb.RegisterChapterGRPCServer(grpcServer, chapterController)
	pb.RegisterParagraphGRPCServer(grpcServer, paragraphController)

	return App{cfg: config, grpcServer: grpcServer, logger: logger}, nil
}

func (a *App) Run(ctx context.Context) error {
	address := fmt.Sprintf("%s:%d", a.cfg.GRPC.IP, a.cfg.GRPC.Port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	a.logger.Printf("started server on %s", address)
	return a.grpcServer.Serve(listener)
}
