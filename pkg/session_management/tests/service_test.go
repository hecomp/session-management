package tests_test

import (
	"errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
	"time"

	. "github.com/hecomp/session-management/internal/models"
	. "github.com/hecomp/session-management/pkg/repository"
	"github.com/hecomp/session-management/pkg/repository/repositoryfakes"
	. "github.com/hecomp/session-management/pkg/session_management"
	"github.com/hecomp/session-management/pkg/test"
)

func TestService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Service Suite")
}

type ServiceSuite struct {
	service SessionMgmntService
	fakeRepo *repositoryfakes.FakeSessionMgmntRepository
}

var _ = Describe("Service", func() {

	s := &ServiceSuite{}

	BeforeEach(func() {
		logger := test.GetLogger()
		s.fakeRepo = new(repositoryfakes.FakeSessionMgmntRepository)
		s.service = NewService(s.fakeRepo, logger)
	})

	Describe("Create Session", func() {
		Context("Create()", func() {
			When("the API os called with TTL as param", func() {
				It("stores an unique sessionId in-memory store", func() {
					req := &SessionRequest{
						TTL: 50,
					}

					s.fakeRepo.CreateReturns(nil)
					sessionId, err := s.service.Create(req)
					Expect(err).To(BeNil())
					Expect(sessionId).ToNot(BeEmpty())
				})
				It("error stores an unique sessionId in-memory store", func() {
					req := &SessionRequest{
						TTL: 50,
					}
					s.fakeRepo.CreateReturns(errors.New("Error create"))
					sessionId, err := s.service.Create(req)
					Expect(err).ToNot(BeNil())
					Expect(sessionId).To(BeEmpty())
				})
			})
		})
	})

	Context("Destroy()", func() {
		When("the API os called with TTL as param", func() {
			It("destroy an unique sessionId in-memory store", func() {
				uniqueUUID := "90660b89-100e-4f8f-9801-2524df6fbe34"
				session := &DestroyRequest{
					SessionId: uniqueUUID,
				}
				s.fakeRepo.ExistReturns(true, nil)
				s.fakeRepo.DestroyReturns(nil)
				err := s.service.Destroy(session)
				Expect(err).To(BeNil())
			})
			It("error empty sessionId sent", func() {
				session := &DestroyRequest{
					SessionId: "",
				}
				s.fakeRepo.ExistReturns(true, errors.New("empty sessionId"))
				err := s.service.Destroy(session)
				Expect(err).ToNot(BeNil())
				Expect(err).To(Equal(ErrEmpty))
			})
			It("error exist an unique sessionId in-memory store", func() {
				uniqueUUID := "90660b89-100e-4f8f-9801-2524df6fbe34"
				session := &DestroyRequest{
					SessionId: uniqueUUID,
				}
				s.fakeRepo.ExistReturns(true, errors.New("error exist"))
				err := s.service.Destroy(session)
				Expect(err).ToNot(BeNil())
				Expect(err).To(Equal(ErrExist))
			})
			It("not found an unique sessionId in-memory store", func() {
				uniqueUUID := "90660b89-100e-4f8f-9801-2524df6fbe34"
				session := &DestroyRequest{
					SessionId: uniqueUUID,
				}
				s.fakeRepo.ExistReturns(false, nil)
				err := s.service.Destroy(session)
				Expect(err).ToNot(BeNil())
				Expect(err).To(Equal(ErrNotFound))
			})
			It("error destroy an unique sessionId in-memory store", func() {
				uniqueUUID := "90660b89-100e-4f8f-9801-2524df6fbe34"
				session := &DestroyRequest{
					SessionId: uniqueUUID,
				}
				s.fakeRepo.ExistReturns(true, nil)
				s.fakeRepo.DestroyReturns(errors.New("error destroy"))
				err := s.service.Destroy(session)
				Expect(err).ToNot(BeNil())
				Expect(err).To(Equal(ErrDestroy))
			})
		})
	})

	Context("Extend()", func() {
		When("the API os called with TTL as param", func() {
			It("extend an unique sessionId in-memory store", func() {
				uniqueUUID := "90660b89-100e-4f8f-9801-2524df6fbe34"
				session := &ExtendRequest{
					TTL: 100,
					SessionId: uniqueUUID,
				}
				s.fakeRepo.ExtendReturns(true, nil)
				err := s.service.Extend(session)
				Expect(err).To(BeNil())
			})
			It("error empty session id extend sent", func() {
				session := &ExtendRequest{
					SessionId: "",
				}
				s.fakeRepo.ExtendReturns(false, errors.New("error error empty session id"))
				err := s.service.Extend(session)
				Expect(err).ToNot(BeNil())
				Expect(err).To(Equal(ErrEmpty))
			})
			It("error not found empty extend an unique sessionId in-memory store", func() {
				uniqueUUID := "90660b89-100e-4f8f-9801-2524df6fbe34"
				session := &ExtendRequest{
					SessionId: uniqueUUID,
				}
				s.fakeRepo.ExtendReturns(false, nil)

				err := s.service.Extend(session)
				Expect(err).ToNot(BeNil())
				Expect(err).To(Equal(ErrNotFound))
			})
			It("error extend an unique sessionId in-memory store", func() {
				uniqueUUID := "90660b89-100e-4f8f-9801-2524df6fbe34"
				session := &ExtendRequest{
					SessionId: uniqueUUID,
				}
				s.fakeRepo.ExtendReturns(false, errors.New("error extend session"))
				err := s.service.Extend(session)
				Expect(err).ToNot(BeNil())
				Expect(err).To(Equal(ErrExtend))
			})
		})
	})

	Context("List()", func() {
		When("the API os called with TTL as param", func() {
			It("list an unique sessionId in-memory store", func() {
				expiration := time.Now().Add(time.Minute * time.Duration(5))
				uniqueUUID := "90660b89-100e-4f8f-9801-2524df6fbe34"

				sessionMap := map[string]Item{
					"90660b89-100e-4f8f-9801-2524df6fbe34": {
						Oject: []byte(uniqueUUID),
						Expiration: expiration.UnixNano(),
					},
				}
				sessions := test.ConvertMapToList(sessionMap)

				s.fakeRepo.ListReturns(sessions, nil)
				res, err := s.service.List()
				Expect(err).To(BeNil())
				Expect(res).To(Equal(sessions))
			})
			It("error list an unique sessionId in-memory store", func() {
				sessionMap := map[string]Item{}
				sessions := test.ConvertMapToList(sessionMap)
				s.fakeRepo.ListReturns(sessions, errors.New("Error destroy"))
				res, err := s.service.List()
				Expect(err).ToNot(BeNil())
				Expect(res).To(BeNil())
			})
		})
	})

})
