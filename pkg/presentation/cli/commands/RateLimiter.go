package commands

import (
	"github.com/Rototot/anti-brute-force/pkg/application/usecases"
	"github.com/Rototot/anti-brute-force/pkg/presentation/rest/controllers"
	"github.com/spf13/cobra"
	"net"
)

func NewLimiter(
	resetHandler controllers.ResetAttemptsHandler,
) *cobra.Command {
	var limiterCmd = &cobra.Command{
		Use:   "limiter",
		Short: "Limiter control",
	}

	limiterCmd.AddCommand(
		newRateLimiterReset(resetHandler),
	)

	return limiterCmd
}

func newRateLimiterReset(resetHandler controllers.ResetAttemptsHandler) *cobra.Command {
	var resetCmd = &cobra.Command{
		Use:     "reset",
		Short:   "Reset limits for ip and login",
		Example: `reset --login "login-123" --ip "192.168.1.1/8"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			ip, err := cmd.Flags().GetIP("ip")
			if err != nil {
				return err
			}

			login, err := cmd.Flags().GetString("login")
			if err != nil {
				return err
			}

			var useCase = usecases.ResetLoginAttempts{
				Login: login,
				IP:    ip,
			}

			if err := resetHandler.Execute(useCase); err != nil {
				return err
			}

			return nil
		},
	}

	resetCmd.Flags().String("login", "", "login value for reset")
	resetCmd.Flags().IP("ip", net.IP{}, "login value for reset")

	if err := resetCmd.MarkFlagRequired("login"); err != nil {
		panic(err)
	}
	if err := resetCmd.MarkFlagRequired("ip"); err != nil {
		panic(err)
	}

	return resetCmd
}
