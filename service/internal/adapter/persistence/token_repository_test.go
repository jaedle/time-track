package persistence_test

import (
	"database/sql"
	"github.com/jaedle/time-track/service/internal/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jaedle/time-track/service/internal/adapter/persistence"
)

const localDataSource = "root:password@tcp(localhost:3307)/database"

const aTokenId = "d96cacde-f82f-401f-bdf3-a29fc2624582"
const anotherTokenId = "0fb1b3e3-9c41-45fc-b11e-6b2fc19a04af"
const yetAnotherTokenId = "2011697f-2807-4e45-9f06-f50c663e567c"

const anUserId = "a-user"
const anUserToken = "token-1"
const anUserAnotherToken = "token-2"

const anotherUserId = "another-user"
const anotherUserToken = "another-token"

const anUserIdWithNoTokens = "missing-user"


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
			Id:     aTokenId,
			UserId: anUserId,
			Token:  anUserToken,
		})
		Expect(err).NotTo(HaveOccurred())
		err = repo.Insert(model.Token{
			Id:     anotherTokenId,
			UserId: anUserId,
			Token:  anUserAnotherToken,
		})
		Expect(err).NotTo(HaveOccurred())

		err = repo.Insert(model.Token{
			Id:     yetAnotherTokenId,
			UserId: anotherUserId,
			Token:  anotherUserToken,
		})
		Expect(err).NotTo(HaveOccurred())

		Expect(repo.Size()).To(Equal(3))
		Expect(repo.FindForUser(anUserId)).To(ConsistOf(
			model.Token{
				Id:     aTokenId,
				UserId: anUserId,
				Token:  anUserToken,
			},
			model.Token{
				Id:     anotherTokenId,
				UserId: anUserId,
				Token:  anUserAnotherToken,
			},
		))

		Expect(repo.FindForUser(anotherUserId)).To(ConsistOf(
			model.Token{
				Id:     yetAnotherTokenId,
				UserId: anotherUserId,
				Token:  anotherUserToken,
			},
		))

		Expect(repo.FindForUser(anUserIdWithNoTokens)).To(BeEmpty())
	})

	It("requires a unique token id", func() {
		Expect(repo.Insert(model.Token{
			Id:     aTokenId,
			UserId: anUserId,
			Token:  anUserToken,
		})).NotTo(HaveOccurred())

		err := repo.Insert(model.Token{
			Id:     aTokenId,
			UserId: anUserId,
			Token:  anUserToken,
		})

		Expect(err).To(HaveOccurred())
		Expect(repo.Size()).To(Equal(1))
	})


	It("deletes token by id", func() {
		err := repo.Insert(model.Token{
			Id:     aTokenId,
			UserId: anUserId,
			Token:  anUserToken,
		})
		Expect(err).NotTo(HaveOccurred())
		err = repo.Insert(model.Token{
			Id:     anotherTokenId,
			UserId: anUserId,
			Token:  anUserAnotherToken,
		})
		Expect(err).NotTo(HaveOccurred())

		err = repo.Insert(model.Token{
			Id:     yetAnotherTokenId,
			UserId: anotherUserId,
			Token:  anotherUserToken,
		})
		Expect(err).NotTo(HaveOccurred())


		err = repo.Delete(aTokenId)


		Expect(err).To(Not(HaveOccurred()))
		Expect(repo.Size()).To(Equal(2))

		Expect(repo.FindForUser(anUserId)).To(ConsistOf(model.Token{
			Id:     anotherTokenId,
			UserId: anUserId,
			Token:  anUserAnotherToken,
		}))
		Expect(repo.FindForUser(anotherUserId)).To(ConsistOf(model.Token{
			Id:     yetAnotherTokenId,
			UserId: anotherUserId,
			Token:  anotherUserToken,
		}))
	})
})
