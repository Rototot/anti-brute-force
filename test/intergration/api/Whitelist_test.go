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

type WhitekListTestSuite struct {
	suite.Suite

	done context.CancelFunc
}

func (s *WhitekListTestSuite) SetupSuite() {
	_, done := StartApp()
	s.done = done
}

func (s *WhitekListTestSuite) SetupTest() {
	postgres.Clean()
	redis.Clean()
	fixtures.Load()
}

func (s *WhitekListTestSuite) TearDownAllSuite() {
	s.done()
}

func (s *WhitekListTestSuite) TestCreate() {

	s.Run("when ok", func() {
		apitest.New(). // configuration
				EnableNetworking().
				Post(fmt.Sprintf("%s/whitelist", baseUrl)).
				JSON(`{"subnet":"10.0.0.1/8"}`).
				Expect(s.T()).
				Status(http.StatusCreated).
				End()
	})

	s.Run("when invalid body", func() {
		apitest.New(). // configuration
				EnableNetworking().
				Post(fmt.Sprintf("%s/whitelist", baseUrl)). // request
				JSON(`{"subnet": "192-168-8-0"}`).
				Expect(s.T()).
				Status(http.StatusBadRequest).
				End()
	})
}

func (s *WhitekListTestSuite) TestDelete() {
	s.Run("when ok", func() {
		apitest.New(). // configuration
				EnableNetworking().
				Delete(fmt.Sprintf("%s/whitelist", baseUrl)). // request
				JSON(`{"subnet":"15.0.0.1/24"}`).
				Expect(s.T()).
				Status(http.StatusNoContent).
				End()
	})

	s.Run("when invali cidr", func() {
		apitest.New(). // configuration
				EnableNetworking().
				Delete(fmt.Sprintf("%s/whitelist", baseUrl)). // request
				Expect(s.T()).
				Status(http.StatusBadRequest).
				End()
	})
}

func (s *WhitekListTestSuite) TestIndex() {
	//s.Run("list", func() {
	//	apitest.New(). // configuration
	//		EnableNetworking().
	//		Get(fmt.Sprintf("%s/blacklist", baseUrl)). // request
	//		Expect(s.T()).
	//		Body(`[{"id":1, "subnet":""}]`).
	//		Status(http.StatusBadRequest).
	//		End()
	//})
}

func TestWhitekListTestSuite(t *testing.T) {
	suite.Run(t, new(WhitekListTestSuite))
}
