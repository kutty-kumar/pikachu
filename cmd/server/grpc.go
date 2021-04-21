package main

import (
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/infobloxopen/atlas-app-toolkit/gateway"
	"github.com/infobloxopen/atlas-app-toolkit/requestid"
	"github.com/kutty-kumar/charminder/pkg"
	"github.com/kutty-kumar/ho_oh/pikachu_v1"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gLogger "gorm.io/gorm/logger"
	"log"
	"os"
	"pikachu/pkg/domain"
	r "pikachu/pkg/repository"
	"pikachu/pkg/svc"
	"time"
)

var (
	reg = prometheus.NewRegistry()
	grpcMetrics = grpc_prometheus.NewServerMetrics()
	createUserSuccessMetric = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "user_service_create_user_success_count",
		Help: "total number of successful invocations of create user method in user service",
	}, []string{"create_user_success_count"})
	createUserFailureMetric = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "user_service_create_user_failure_count",
		Help: "total number of failure invocations of create user method in user service",
	}, []string{"create_user_failure_count"})
)

func init(){
	reg.MustRegister(grpcMetrics, createUserSuccessMetric, createUserFailureMetric)
	createUserSuccessMetric.WithLabelValues("user_service")
	createUserFailureMetric.WithLabelValues("user_service")
}

func NewGRPCServer(logger *logrus.Logger, dbConnectionString string) (*grpc.Server, error) {
	grpcServer := grpc.NewServer(
		grpc.KeepaliveParams(
			keepalive.ServerParameters{
				Time:    time.Duration(viper.GetInt("heart_beat_config.keep_alive_time")) * time.Second,
				Timeout: time.Duration(viper.GetInt("heart_beat_config.keep_alive_timeout")) * time.Second,
			},
		),
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				// logging middleware
				grpc_logrus.UnaryServerInterceptor(logrus.NewEntry(logger)),

				// Request-Id interceptor
				requestid.UnaryServerInterceptor(),

				// Metrics middleware
				grpc_prometheus.UnaryServerInterceptor,

				// validation middleware
				grpc_validator.UnaryServerInterceptor(),

				// collection operators middleware
				gateway.UnaryServerInterceptor(),
			),
		),
	)


	dbLogger := gLogger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		gLogger.Config{
			SlowThreshold:              time.Second,   // Slow SQL threshold
			LogLevel:                   gLogger.Info, // Log level
			IgnoreRecordNotFoundError: true,           // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,          // Disable color
		},
	)
	// create new mysql database connection
	db, err := gorm.Open(mysql.Open(viper.GetString("database_config.dsn")), &gorm.Config{Logger: dbLogger})
	if err != nil {
		return nil, err
	}

	//dropTables(db)
	createTables(db)

	domainFactory := pkg.NewDomainFactory()
	domainFactory.RegisterMapping("user", func() pkg.Base {
		return &domain.User{}
	})
	domainFactory.RegisterMapping("identity", func() pkg.Base {
		return &domain.Identity{}
	})

	externalIdSetter := func(externalId string, base pkg.Base) pkg.Base {
		base.SetExternalId(externalId)
		return base
	}
	setterOption := pkg.WithExternalIdSetter(externalIdSetter)
	dbOption := pkg.WithDb(db)
	userBaseDao := pkg.NewBaseGORMDao(setterOption, pkg.WithCreator(domainFactory.GetMapping("user")), dbOption)

	identityBaseDao := pkg.NewBaseGORMDao(dbOption, pkg.WithCreator(domainFactory.GetMapping("identity")), setterOption)
	userAttributeBaseDao := pkg.NewBaseGORMDao(dbOption, pkg.WithCreator(domainFactory.GetMapping("user_attributes")), setterOption)
	identityRepository := r.NewIdentityGormRepository(identityBaseDao)
	userAttributeRepository := r.NewUserAttributeGormRepository(userAttributeBaseDao)
	// register service implementation with the grpcServer
	userBaseSvc := pkg.NewBaseSvc(userBaseDao)
	identityBaseSvc := pkg.NewBaseSvc(identityBaseDao)
	userAttributeBaseSvc := pkg.NewBaseSvc(userAttributeBaseDao)
	identityService := svc.NewIdentityService(identityBaseSvc, &identityRepository)
	userAttributeService := svc.NewUserAttributeService(userAttributeBaseSvc, &userAttributeRepository)
	userService := svc.NewUserService(userBaseSvc, identityService, userAttributeService)

	pikachu_v1.RegisterUserServiceServer(grpcServer, &userService)
	grpcMetrics.InitializeMetrics(grpcServer)
	return grpcServer, nil
}

func createTables(db *gorm.DB) {
	err := db.AutoMigrate(domain.User{}, domain.Identity{}, domain.Relation{}, domain.UserAttribute{})
	if err != nil {
		log.Fatalf("An error %v occurred while automigrating", err)
	}
}
