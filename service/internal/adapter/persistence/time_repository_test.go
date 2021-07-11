package persistence_test

import (
	"database/sql"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jaedle/time-track/service/internal/adapter/persistence"
)

var _ = Describe("TimeRepository", func() {
	var database *sql.DB
	var repo *persistence.TimeRepository
	BeforeEach(func() {
		var err error
		database, err = persistence.NewDatabase("root:password@tcp(localhost:3307)/database")
		Expect(err).NotTo(HaveOccurred())
		Expect(database.Ping()).NotTo(HaveOccurred())

		_, err = database.Exec("DROP TABLE IF EXISTS times")
		Expect(err).NotTo(HaveOccurred())

		repo = persistence.NewTimeRepository(database)
		err = repo.Init()
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		_ = database.Close()
	})

	It("is initially empty", func() {
		Expect(repo.Size()).To(Equal(0))
	})

	It("persists times", func() {
		err := repo.Insert("some-description")
		Expect(err).NotTo(HaveOccurred())

		err = repo.Insert("another-description")
		Expect(err).NotTo(HaveOccurred())

		Expect(repo.Size()).To(Equal(2))
	})

})
