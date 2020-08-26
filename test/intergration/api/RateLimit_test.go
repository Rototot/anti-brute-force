//+build e2e

package api

import (
	"context"
	"fmt"
	"github.com/Rototot/anti-brute-force/test/utils/fixtures"
	"github.com/Rototot/anti-brute-force/test/utils/postgres"
	"github.com/Rototot/anti-brute-force/test/utils/redis"
	"github.com/steinfletcher/apitest"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type RateLimitTestSuite struct {
	suite.Suite

	done context.CancelFunc
}

func (s *RateLimitTestSuite) SetupSuite() {
	_, done := StartApp()
	s.done = done
}

func (s *RateLimitTestSuite) SetupTest() {
	postgres.Clean()
	redis.Clean()
	fixtures.Load()
}

func (s *RateLimitTestSuite) TearDownAllSuite() {
	s.done()
}

func (s *RateLimitTestSuite) TestAttempt() {

	s.Run("when ok", func() {
		apitest.New(). // configuration
				EnableNetworking().
				Post(fmt.Sprintf("%s/login/attempt", baseUrl)).
				JSON(`{"login":"test_login", "ip": "15.0.0.1", "password": "test_pass"}`).
				Expect(s.T()).
				Body(`{"data":true}`).
				Status(http.StatusOK).
				End()
	})

	s.Run("when invalid body", func() {
		apitest.New(). // configuration
				EnableNetworking().
				Post(fmt.Sprintf("%s/login/attempt", baseUrl)). // request
				JSON(`{"login":"test_login", "ip": "15.0.0.1/2", "password": "test_pass"}`).
				Expect(s.T()).
				Status(http.StatusBadRequest).
				End()
	})
}

func (s *RateLimitTestSuite) TestReset() {
	s.Run("when ok", func() {
		apitest.New(). // configuration
				EnableNetworking().
				Delete(fmt.Sprintf("%s/login/attempt", baseUrl)). // request
				JSON(`{"login":"test_login", "ip": "15.0.0.1"}`).
				Expect(s.T()).
				Status(http.StatusNoContent).
				End()
	})

	s.Run("when invalid cidr", func() {
		apitest.New(). // configuration
				EnableNetworking().
				Delete(fmt.Sprintf("%s/login/attempt", baseUrl)). // request
				JSON(`{"login":"test_login", "ip": "15.0.0.1/2"}`).
				Expect(s.T()).
				Status(http.StatusBadRequest).
				End()
	})
}

func TestRateLimitTestSuite(t *testing.T) {
	suite.Run(t, new(RateLimitTestSuite))
}
