package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/johnrazeur/gin-boilerplate/config"
	_ "github.com/johnrazeur/gin-boilerplate/helpers"
	"github.com/johnrazeur/gin-boilerplate/models"
	"github.com/johnrazeur/gin-boilerplate/seeder"

	"github.com/stretchr/testify/suite"
)

type UserTestSuite struct {
	suite.Suite
	rec     *httptest.ResponseRecorder
	context *gin.Context
	app     *gin.Engine
}

func initDb() *gorm.DB {
	db, err := gorm.Open("sqlite3", "file::memory:?cache=shared")

	if err != nil {
		panic(err)
	}
	defer db.Close()
	return db
}

func (suite *UserTestSuite) SetupTest() {
	gin.SetMode(gin.ReleaseMode)
	config.Init()
	// config.Config.Server.Mode = "test"

	suite.rec = httptest.NewRecorder()
	suite.context, suite.app = gin.CreateTestContext(suite.rec)
	if err := seeder.Seed(models.DB().DB()); err != nil {
		panic(err)
	}
}

func (suite *UserTestSuite) TestSignup() {
	var (
		username = "test"
		email    = "test@gmail.com"
		password = "azertyuiop"
	)

	url := fmt.Sprintf(
		"/signup?username=%s&email=%s&password=%s&password2=%s",
		username,
		email,
		password,
		password,
	)
	suite.context.Request, _ = http.NewRequest(
		"POST",
		url,
		bytes.NewBufferString(""),
	)

	Signup(suite.context)
	suite.Equal(200, suite.rec.Code)

	var user models.User

	err := json.NewDecoder(suite.rec.Body).Decode(&user)

	if err != nil {
		suite.T().Fatalf("Unable to parse response from server %q into User structure, '%v'", suite.rec.Body, err)
	}
}

func (suite *UserTestSuite) TestLogin() {
	var (
		email    = "john@doe.com"
		password = "password"
	)

	url := fmt.Sprintf(
		"/login?&email=%s&password=%s",
		email,
		password,
	)
	suite.context.Request, _ = http.NewRequest(
		"POST",
		url,
		bytes.NewBufferString(""),
	)

	Login(suite.context)
	suite.Equal(200, suite.rec.Code)

	var user models.User

	err := json.NewDecoder(suite.rec.Body).Decode(&user)

	if err != nil {
		suite.T().Fatalf("Unable to parse response from server %q into User structure, '%v'", suite.rec.Body, err)
	}
}

func TestUserTestSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}
