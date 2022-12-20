package repository

import (
	"database/sql"

	"testing"

	"github.com/projetosgo/exemploapi/entity"
	"github.com/stretchr/testify/suite"

	_ "github.com/go-sql-driver/mysql"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	Db *sql.DB
}

func (suite *UserRepositoryTestSuite) SetupSuite() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/usersdb")
	suite.NoError(err)
	_, err = db.Exec("CREATE TABLE users (id varchar(255) NOT NULL, name varchar(255) NOT NULL, email varchar(255) NOT NULL, PRIMARY KEY (id))")
	suite.Db = db
}

func (suite *UserRepositoryTestSuite) TearDownTest() {
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}

func (suite *UserRepositoryTestSuite) TestGivenAnOrder_WhenSave_ThenShouldSaveUser() {
	user, err := entity.NewUser("marcos", "marcos@email.com")
	suite.NoError(err)
	suite.NoError(user.GenerateID())
	suite.NoError(user.Validate())
	repo := NewUserRepository(suite.Db)
	err = repo.Save(user)
	suite.NoError(err)

	var userResult entity.User
	err = suite.Db.QueryRow("SELECT id, name, email from users where id = ?", user.ID).
		Scan(&userResult.ID, &userResult.Name, &userResult.Email)

	suite.NoError(err)
	suite.Equal(user.ID, userResult.ID)
	suite.Equal(user.Name, userResult.Name)
	suite.Equal(user.Email, userResult.Email)
}
