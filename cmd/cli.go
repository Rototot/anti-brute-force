package cmd

import (
	"github.com/Rototot/anti-brute-force/pkg/application/usecases"
	"github.com/Rototot/anti-brute-force/pkg/infrastructure/configurators"
	"github.com/Rototot/anti-brute-force/pkg/infrastructure/persistence/postgres"
	"github.com/Rototot/anti-brute-force/pkg/infrastructure/persistence/postgres/repositories"
	"github.com/Rototot/anti-brute-force/pkg/infrastructure/persistence/redis"
	repositories2 "github.com/Rototot/anti-brute-force/pkg/infrastructure/persistence/redis/repositories"
	"github.com/Rototot/anti-brute-force/pkg/presentation/cli/commands"
	"github.com/spf13/cobra"
)

func NewCLiCmd() *cobra.Command {
	// db
	pgConnection := postgres.NewConnection(configurators.NewPostgresConfig())

	// redis
	redisPool := redis.NewPool(configurators.NewRedisConfig())

	// domain
	// repositories
	ipBlackListRepository := repositories.NewBlackListIPRepository(pgConnection)
	ipWhiteListRepository := repositories.NewWhiteListIPRepository(pgConnection)

	bucketRepository := repositories2.NewBucketRepository(redisPool)

	// application
	//   cases
	caseAddToBlacklist := usecases.NewAddIPToBlacklistHandler(ipBlackListRepository)
	caseRemoveFromBlacklist := usecases.NewRemoveIPFromBlackListHandler(ipBlackListRepository)

	caseAddToWhitelist := usecases.NewAddIPToWhiteListHandler(ipWhiteListRepository)
	caseRemoveFromWhitelist := usecases.NewRemoveIPFromWhiteListHandler(ipWhiteListRepository)

	caseResetAttempts := usecases.NewResetLoginAttemptsHandler(bucketRepository)

	cliCmd := &cobra.Command{
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
