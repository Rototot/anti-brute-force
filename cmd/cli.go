package cmd

import (
	"github.com/Rototot/anti-brute-force/pkg/application/usecases"
	"github.com/Rototot/anti-brute-force/pkg/domain/services"
	"github.com/Rototot/anti-brute-force/pkg/infrastructure/configurators"
	"github.com/Rototot/anti-brute-force/pkg/infrastructure/persistence/postgres"
	"github.com/Rototot/anti-brute-force/pkg/infrastructure/persistence/postgres/repositories"
	"github.com/Rototot/anti-brute-force/pkg/infrastructure/persistence/redis"
	repositories2 "github.com/Rototot/anti-brute-force/pkg/infrastructure/persistence/redis/repositories"
	"github.com/Rototot/anti-brute-force/pkg/presentation/cli/commands"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

func newCLiCmd() *cobra.Command {

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
	rateLimiter := services.NewBucketRateLimiter(bucketRepository)

	// application
	//   cases
	caseAddToBlacklist := usecases.NewAddIpToBlacklistHandler(ipBlackListRepository)
	caseRemoveFromBlacklist := usecases.NewRemoveIpFromBlackListHandler(ipBlackListRepository)

	caseAddToWhitelist := usecases.NewAddIPToWhiteListHandler(ipWhiteListRepository)
	caseRemoveFromWhitelist := usecases.NewRemoveIpFromWhiteListHandler(ipWhiteListRepository)

	caseResetAttempts := usecases.NewResetLoginAttemptsHandler(bucketRepository, rateLimiter)

	var cliCmd = &cobra.Command{
		Use:   "cli",
		Short: "cli app control",
		Long:  "Choose command for continue",
	}

	cliCmd.AddCommand(
		commands.NewBlackList(caseAddToBlacklist, caseRemoveFromBlacklist),
		commands.NewWhitelist(caseAddToWhitelist, caseRemoveFromWhitelist),
		commands.NewLimiter(caseResetAttempts),
	)

	return cliCmd
}
