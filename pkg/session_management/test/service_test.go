package test_test

import (
	"errors"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/hecomp/session-management/internal/models"
	"github.com/hecomp/session-management/pkg/repository/repositoryfakes"
	. "github.com/hecomp/session-management/pkg/session_management"
	"github.com/hecomp/session-management/pkg/test"
	. "github.com/hecomp/session-management/pkg/repository"
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

		//sessionMgmntResp := &SessionMgmntResponse{
		//	Message: "",
		//	Data: nil,
		//	Err: nil,
		//}
		Context("Create()", func() {
			When("the API os called with TTL as param", func() {
				It("stores an unique sessionId in-memory store", func() {
					req := &SessionRequest{
						TTL: 50,
					}
					s.fakeRepo.CreateReturns(nil)
					res, err := s.service.Create(req)
					Expect(err).To(BeNil())
					Expect(res.Message).To(Equal(CreateSessionSuccess))
				})
				It("error stores an unique sessionId in-memory store", func() {
					req := &SessionRequest{
						TTL: 50,
					}
					s.fakeRepo.CreateReturns(errors.New("Error create"))
					res, err := s.service.Create(req)
					Expect(err).ToNot(BeNil())
					Expect(res.Message).To(Equal(ErrEmpty.Error()))
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
				s.fakeRepo.DestroyReturns(nil)
				res, err := s.service.Destroy(session)
				Expect(err).To(BeNil())
				Expect(res.Message).To(Equal(DestroySessionSuccess))
			})
			It("error destroy an unique sessionId in-memory store", func() {
				uniqueUUID := "90660b89-100e-4f8f-9801-2524df6fbe34"
				session := &DestroyRequest{
					SessionId: uniqueUUID,
				}
				s.fakeRepo.DestroyReturns(errors.New("Error destroy"))
				res, err := s.service.Destroy(session)
				Expect(err).ToNot(BeNil())
				Expect(res.Message).To(Equal(ErrDestroy))
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
				s.fakeRepo.ExistReturns(true, nil)
				s.fakeRepo.ExtendReturns(nil)
				res, err := s.service.Extend(session)
				Expect(err).To(BeNil())
				Expect(res.Message).To(Equal(ExtendSessionSuccess))
			})
			It("error  session id empty extend an unique sessionId in-memory store", func() {
				uniqueUUID := "90660b89-100e-4f8f-9801-2524df6fbe34"
				session := &ExtendRequest{
					SessionId: uniqueUUID,
				}
				s.fakeRepo.ExistReturns(false, errors.New("Error  session id empty"))
				res, err := s.service.Extend(session)
				Expect(err).ToNot(BeNil())
				Expect(res.Message).To(Equal(ErrEmpty.Error()))
			})
			It("error not found empty extend an unique sessionId in-memory store", func() {
				uniqueUUID := "90660b89-100e-4f8f-9801-2524df6fbe34"
				session := &ExtendRequest{
					SessionId: uniqueUUID,
				}
				s.fakeRepo.ExistReturns(false, nil)

				res, err := s.service.Extend(session)
				Expect(err).ToNot(BeNil())
				Expect(res.Message).To(Equal(ErrNotFound.Error()))
			})
			It("error extend an unique sessionId in-memory store", func() {
				uniqueUUID := "90660b89-100e-4f8f-9801-2524df6fbe34"
				session := &ExtendRequest{
					SessionId: uniqueUUID,
				}
				s.fakeRepo.ExistReturns(true, nil)
				s.fakeRepo.ExtendReturns(errors.New("error extend session"))
				res, err := s.service.Extend(session)
				Expect(err).ToNot(BeNil())
				Expect(res.Message).To(Equal(ErrExtend))
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
				Expect(res.Message).To(Equal(ListSessionSuccess))
			})
			It("error list an unique sessionId in-memory store", func() {
				sessionMap := map[string]Item{}
				sessions := test.ConvertMapToList(sessionMap)
				s.fakeRepo.ListReturns(sessions, errors.New("Error destroy"))
				res, err := s.service.List()
				Expect(err).ToNot(BeNil())
				Expect(res.Message).To(Equal(ErrList))
			})
		})
	})

})
