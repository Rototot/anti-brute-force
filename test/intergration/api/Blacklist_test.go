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

type BlackListTestSuite struct {
	suite.Suite

	done context.CancelFunc
}

func (s *BlackListTestSuite) SetupSuite() {
	_, done := StartApp()
	s.done = done
}

func (s *BlackListTestSuite) SetupTest() {
	postgres.Clean()
	redis.Clean()
	fixtures.Load()
}

func (s *BlackListTestSuite) TearDownAllSuite() {
	s.done()
}

func (s *BlackListTestSuite) TestCreate() {

	s.Run("when ok", func() {
		apitest.New(). // configuration
				EnableNetworking().
				Post(fmt.Sprintf("%s/blacklist", baseUrl)).
				JSON(`{"subnet":"10.0.0.1/8"}`).
				Expect(s.T()).
				Status(http.StatusCreated).
				End()
	})

	s.Run("when invalid body", func() {
		apitest.New(). // configuration
				EnableNetworking().
				Post(fmt.Sprintf("%s/blacklist", baseUrl)). // request
				JSON(`{"subnet": "192-168-8-0"}`).
				Expect(s.T()).
				Status(http.StatusBadRequest).
				End()
	})
}

func (s *BlackListTestSuite) TestDelete() {
	s.Run("when ok", func() {
		apitest.New(). // configuration
				EnableNetworking().
				Delete(fmt.Sprintf("%s/blacklist", baseUrl)). // request
				JSON(`{"subnet": "192.168.1.1/8"}`).
				Expect(s.T()).
				Status(http.StatusNoContent).
				End()
	})

	s.Run("when invali cidr", func() {
		apitest.New(). // configuration
				EnableNetworking().
				Delete(fmt.Sprintf("%s/blacklist", baseUrl)). // request
				JSON(`{"subnet": "192.168.8.0"}`).
				Expect(s.T()).
				Status(http.StatusBadRequest).
				End()
	})
}

func (s *BlackListTestSuite) TestIndex() {
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

func TestBlackListTestSuite(t *testing.T) {
	suite.Run(t, new(BlackListTestSuite))
}
