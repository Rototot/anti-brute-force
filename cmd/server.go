package cmd

import (
	"fmt"
	"github.com/Rototot/anti-brute-force/pkg/application/usecases"
	"github.com/Rototot/anti-brute-force/pkg/domain/factories"
	"github.com/Rototot/anti-brute-force/pkg/domain/services"
	"github.com/Rototot/anti-brute-force/pkg/infrastructure/configurators"
	"github.com/Rototot/anti-brute-force/pkg/infrastructure/persistence/postgres"
	"github.com/Rototot/anti-brute-force/pkg/infrastructure/persistence/postgres/repositories"
	"github.com/Rototot/anti-brute-force/pkg/infrastructure/persistence/redis"
	repositories2 "github.com/Rototot/anti-brute-force/pkg/infrastructure/persistence/redis/repositories"
	"github.com/Rototot/anti-brute-force/pkg/presentation/rest/controllers"
	"github.com/Rototot/anti-brute-force/pkg/presentation/rest/routers"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

const (
	defaultServerPort = 80
)

func newServerCmd() *cobra.Command {
	var serverCmd = &cobra.Command{
		Use:   "server",
		Short: "Start web API server",
		RunE: func(cmd *cobra.Command, args []string) error {
			port, err := cmd.Flags().GetInt("port")
			if err != nil {
				return err
			}

			return runHttpServer(port)
		},
	}

	serverCmd.Flags().Int("port", defaultServerPort, "listen server port")

	return serverCmd
}

func runHttpServer(port int) error {
	//db
	pgConnection := postgres.NewConnection(*configurators.NewPostgresConfig(viper.GetViper()))

	//redis
	redisPool := redis.NewPool(*configurators.NewRedisConfig(viper.GetViper()))

	// domain
	// repositories
	ipBlackListRepository := repositories.NewBlackListIPRepository(pgConnection)
	ipWhiteListRepository := repositories.NewWhiteListIPRepository(pgConnection)

	bucketRepository := repositories2.NewBucketRepository(redisPool)

	// domain services
	bucketFactory := factories.NewBucketFactory(configurators.NewBucketConfigurator(viper.GetViper()))
	ipGuard := services.NewIPGuard(ipWhiteListRepository, ipBlackListRepository)
	rateLimiter := services.NewBucketRateLimiter(bucketRepository)

	// application
	//   cases
	caseAddToBlacklist := usecases.NewAddIpToBlacklistHandler(ipBlackListRepository)
	caseRemoveFromBlacklist := usecases.NewRemoveIpFromBlackListHandler(ipBlackListRepository)

	caseAddToWhitelist := usecases.NewAddIPToWhiteListHandler(ipWhiteListRepository)
	caseRemoveFromWhitelist := usecases.NewRemoveIpFromWhiteListHandler(ipWhiteListRepository)

	caseCheckAttempt := usecases.NewCheckLoginAttemptHandler(bucketRepository, bucketFactory, ipGuard, rateLimiter)

	caseResetAttempts := usecases.NewResetLoginAttemptsHandler(bucketRepository, rateLimiter)

	// controllers
	blackListController := controllers.NewBlackListCrudController(
		validator.New(),
		caseAddToBlacklist,
		caseRemoveFromBlacklist,
	)

	whiteListControlller := controllers.NewWhiteListCrudController(
		validator.New(),
		caseAddToWhitelist,
		caseRemoveFromWhitelist,
	)

	rateLimiterController := controllers.NewRateLimitController(
		validator.New(),
		caseCheckAttempt,
		caseResetAttempts,
	)

	router := routers.NewRouter(
		whiteListControlller,
		blackListController,
		rateLimiterController,
	).Create()

	//
	server := &http.Server{
		Handler: router,
		Addr:    ":8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	var listenAddress = fmt.Sprintf(":%d", port)

	zap.S().Infof("\nStart listen addr %s\n", listenAddress)

	err := server.ListenAndServe()
	zap.S().Error(err)

	return err
}
