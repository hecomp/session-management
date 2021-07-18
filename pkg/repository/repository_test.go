package repository_test

import (
	"errors"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/hecomp/session-management/internal/models"
	"github.com/hecomp/session-management/pkg/in_memory/in_memoryfakes"
	. "github.com/hecomp/session-management/pkg/repository"
	"github.com/hecomp/session-management/pkg/test"
)

type RepositorySuite struct {
	repo         SessionMgmntRepository
	fakeMemStore *in_memoryfakes.FakeMemStore
}

var _ = Describe("Repository", func() {

	s := &RepositorySuite{}

	BeforeEach(func() {
		logger := test.GetLogger()
		s.fakeMemStore = new(in_memoryfakes.FakeMemStore)
		s.repo = NewSessionMgmntRepository(s.fakeMemStore, logger)
	})

	Describe("Create Session", func() {
		Context("Create()", func() {
			When("the API os called with TTL as param", func() {
				It("stores an unique sessionId in-memory store", func() {
					uniqueUUID := "90660b89-100e-4f8f-9801-2524df6fbe34"
					expiration := time.Now().Add(time.Second * time.Duration(40))
					s.fakeMemStore.CommitReturns(nil)
					err := s.repo.Create(uniqueUUID, expiration)
					Expect(err).To(BeNil())
				})
				It("error stores an unique sessionId in-memory store", func() {
					uniqueUUID := "90660b89-100e-4f8f-9801-2524df6fbe34"
					expiration := time.Now().Add(time.Second * time.Duration(40))
					s.fakeMemStore.CommitReturns(errors.New("Error commit"))
					err := s.repo.Create(uniqueUUID, expiration)
					Expect(err).ToNot(BeNil())
				})
			})
		})
	})

	Describe("Destroy Session", func() {
		Context("Destroy()", func() {
			When("the API os called with TTL as param", func() {
				It("destroy an unique sessionId in-memory store", func() {
					uniqueUUID := "90660b89-100e-4f8f-9801-2524df6fbe34"
					session := &models.DestroyRequest{
						SessionId: uniqueUUID,
					}
					s.fakeMemStore.DeleteReturns(nil)
					err := s.repo.Destroy(session)
					Expect(err).To(BeNil())
				})
				It("error destroy an unique sessionId in-memory store", func() {
					uniqueUUID := "90660b89-100e-4f8f-9801-2524df6fbe34"
					session := &models.DestroyRequest{
						SessionId: uniqueUUID,
					}
					s.fakeMemStore.DeleteReturns(errors.New("Error destroy"))
					err := s.repo.Destroy(session)
					Expect(err).ToNot(BeNil())
				})
			})
		})
	})

	Describe("Extend Session", func() {
		Context("Extend()", func() {
			When("the API os called with TTL as param", func() {
				It("extend an unique sessionId in-memory store", func() {
					uniqueUUID := "90660b89-100e-4f8f-9801-2524df6fbe34"
					session := &models.ExtendRequest{
						TTL: 100,
						SessionId: uniqueUUID,
					}
					s.fakeMemStore.ResetReturns([]byte(uniqueUUID), true, nil)
					err := s.repo.Extend(session)
					Expect(err).To(BeNil())
				})
				It("error extend an unique sessionId in-memory store", func() {
					uniqueUUID := "90660b89-100e-4f8f-9801-2524df6fbe34"
					session := &models.ExtendRequest{
						TTL: 100,
						SessionId: uniqueUUID,
					}
					s.fakeMemStore.ResetReturns([]byte(uniqueUUID), false, errors.New("Error destroy"))
					err := s.repo.Extend(session)
					Expect(err).ToNot(BeNil())
				})
			})
		})
	})

	Describe("Exist Session", func() {
		Context("Exist()", func() {
			When("the API os called with TTL as param", func() {
				It("exist an unique sessionId in-memory store", func() {
					uniqueUUID := "90660b89-100e-4f8f-9801-2524df6fbe34"
					s.fakeMemStore.FindReturns([]byte(uniqueUUID), true, nil)
					found, err := s.repo.Exist(uniqueUUID)
					Expect(err).To(BeNil())
					Expect(found).To(BeTrue())
				})
				It("error extend an unique sessionId in-memory store", func() {
					uniqueUUID := "90660b89-100e-4f8f-9801-2524df6fbe34"
					s.fakeMemStore.FindReturns([]byte(uniqueUUID), false, errors.New("Error Exist"))
					found, err := s.repo.Exist(uniqueUUID)
					Expect(err).ToNot(BeNil())
					Expect(found).ToNot(BeTrue())
				})
			})
		})
	})

	Describe("List Session", func() {
		Context("List()", func() {
			When("the API os called with TTL as param", func() {
				It("list an unique sessionId in-memory store", func() {
					expiration := time.Now().Add(time.Minute * time.Duration(5))
					uniqueUUID := "90660b89-100e-4f8f-9801-2524df6fbe34"
					sessionMap := map[string]models.Item{
						"90660b89-100e-4f8f-9801-2524df6fbe34": {
								Oject: []byte(uniqueUUID),
								Expiration: expiration.UnixNano(),
						},

					}
					s.fakeMemStore.ListReturns(sessionMap, nil)
					sessions, err := s.repo.List()
					sessionsList := test.ConvertMapToList(sessionMap)
					Expect(err).To(BeNil())
					Expect(sessions).To(Equal(sessionsList))
				})
				It("error list an unique sessionId in-memory store", func() {
					sessionMap := map[string]models.Item{}
					s.fakeMemStore.ListReturns(sessionMap, errors.New("Error list"))
					_, err := s.repo.List()
					Expect(err).ToNot(BeNil())
				})
			})
		})
	})
})
