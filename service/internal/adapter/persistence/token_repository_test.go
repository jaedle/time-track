package persistence_test

import (
	"database/sql"
	"github.com/jaedle/time-track/service/internal/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jaedle/time-track/service/internal/adapter/persistence"
)

const localDataSource = "root:password@tcp(localhost:3307)/database"

const anUserId = "a-user"
const anUserToken = "token-1"
const anUserAnotherToken = "token-2"

const anotherUserId = "another-user"
const anotherUserToken = "another-token"


var _ = Describe("TokenRepository", func() {
	var database *sql.DB
	var repo *persistence.TokenRepository

	BeforeEach(func() {
		var err error
		database, err = persistence.NewDatabase(localDataSource)
		Expect(err).NotTo(HaveOccurred())
		Expect(database.Ping()).NotTo(HaveOccurred())

		_, err = database.Exec("DROP TABLE IF EXISTS tokens")
		Expect(err).NotTo(HaveOccurred())

		repo = persistence.NewTokenRepository(database)
		err = repo.Init()
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		_ = database.Close()
	})

	It("is initially empty", func() {
		Expect(repo.Size()).To(Equal(0))
	})

	It("persists tokens", func() {
		err := repo.Insert(model.Token{
			UserId: anUserId,
			Token:  anUserToken,
		})
		Expect(err).NotTo(HaveOccurred())
		err = repo.Insert(model.Token{
			UserId: anUserId,
			Token:  anUserAnotherToken,
		})
		Expect(err).NotTo(HaveOccurred())

		err = repo.Insert(model.Token{
			UserId: anotherUserId,
			Token:  anotherUserToken,
		})
		Expect(err).NotTo(HaveOccurred())

		Expect(repo.Size()).To(Equal(3))
		Expect(repo.FindForUser(anUserId)).To(ConsistOf(
			model.Token{
				UserId: anUserId,
				Token:  anUserToken,
			},
			model.Token{
				UserId: anUserId,
				Token:  anUserAnotherToken,
			},
		))

		Expect(repo.FindForUser(anotherUserId)).To(ConsistOf(
			model.Token{
				UserId: anotherUserId,
				Token:  anotherUserToken,
			},
		))

		Expect(repo.FindForUser("missing-user")).To(BeEmpty())
	})

})
