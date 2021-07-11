package persistence_test

import (
	"database/sql"
	"github.com/jaedle/time-track/service/internal/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jaedle/time-track/service/internal/adapter/persistence"
)

var _ = Describe("TokenRepository", func() {
	var database *sql.DB
	var repo *persistence.TokenRepository

	BeforeEach(func() {
		var err error
		database, err = persistence.NewDatabase("root:password@tcp(localhost:3307)/database")
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
			UserId: "a-user",
			Token:  "token-1",
		})
		Expect(err).NotTo(HaveOccurred())
		err = repo.Insert(model.Token{
			UserId: "a-user",
			Token:  "token-2",
		})
		Expect(err).NotTo(HaveOccurred())

		err = repo.Insert(model.Token{
			UserId: "another-user",
			Token:  "another-token",
		})
		Expect(err).NotTo(HaveOccurred())

		Expect(repo.Size()).To(Equal(3))
		Expect(repo.FindForUser("a-user")).To(ConsistOf(
			model.Token{
				UserId: "a-user",
				Token:  "token-1",
			},
			model.Token{
				UserId: "a-user",
				Token:  "token-2",
			},
		))

		Expect(repo.FindForUser("another-user")).To(ConsistOf(
			model.Token{
				UserId: "another-user",
				Token:  "another-token",
			},
		))

		Expect(repo.FindForUser("missing-user")).To(BeEmpty())
	})

})
