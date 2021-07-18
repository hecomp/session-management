package in_memory_test

import (
	"time"

	"github.com/go-kit/kit/log"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/hecomp/session-management/pkg/in_memory"
	"github.com/hecomp/session-management/pkg/test"
)

type InMemStoreSuite struct {
	mem *InMemStore
	logger log.Logger
}

var _ = Describe("InMemory", func() {

	s := &InMemStoreSuite{}

	BeforeEach(func() {
		s.logger = test.GetLogger()
		s.mem = NewInMemStore(0, s.logger)

	})

	Describe("Create session", func() {

		inMemResponse := "90660b89-100e-4f8f-9801-2524df6fbe34"

		Context("Commit()", func() {
			When("the API os called with TTL as param", func() {
				It("stores an unique sessionId in-memory store", func() {
					uniqueUUID := "90660b89-100e-4f8f-9801-2524df6fbe34"
					err := s.mem.Commit(uniqueUUID, []byte(uniqueUUID), time.Now().Add(time.Minute))
					sessionMap, _ := s.mem.List()
					Expect(err).To(BeNil())
					Expect(string(sessionMap[uniqueUUID].Oject)).To(Equal(inMemResponse))
				})
			})
		})
	})

	Describe("Find session", func() {
		inMemResponse := "90660b89-100e-4f8f-9801-2524df6fbe34"
		uniqueUUID := "90660b89-100e-4f8f-9801-2524df6fbe34"
		BeforeEach(func() {
			err := s.mem.Commit(uniqueUUID, []byte(uniqueUUID), time.Now().Add(time.Minute))
			Expect(err).To(BeNil())
		})

		Context("Find()", func() {
			When("the API os called with TTL as param", func() {
				It("finds sessionId in-memory store", func() {
					obj, _, err := s.mem.Find(uniqueUUID)
					Expect(err).To(BeNil())
					Expect(string(obj)).To(Equal(inMemResponse))
				})
				It("matches sessionId in the in-memory store", func() {
					_, found, err := s.mem.Find(uniqueUUID)
					Expect(err).To(BeNil())
					Expect(found).To(BeTrue())
				})
				It("does not matches sessionId in the in-memory store", func() {
					uniqueUUID2 := "90660b89-100e-4f8f-9801-2524df6fbe99"
					_, found, err := s.mem.Find(uniqueUUID2)
					Expect(err).To(BeNil())
					Expect(found).ToNot(BeTrue())
				})

				It("expired sessionId in the in-memory store", func() {
					er := s.mem.Commit(uniqueUUID, []byte(uniqueUUID), time.Now().Add(100*time.Millisecond))
					time.Sleep(101 * time.Millisecond)

					obj, found, err := s.mem.Find(uniqueUUID)
					Expect(err).To(BeNil())
					Expect(er).To(BeNil())
					Expect(obj).To(BeNil())
					Expect(found).ToNot(BeTrue())
				})
			})
		})
	})

	Describe("Destroy session", func() {
		uniqueUUID := "90660b89-100e-4f8f-9801-2524df6fbe34"
		BeforeEach(func() {
			err := s.mem.Commit(uniqueUUID, []byte(uniqueUUID), time.Now().Add(time.Minute))
			Expect(err).To(BeNil())
		})

		Context("Delete()", func() {
			When("the API os called", func() {
				It("deletes sessionId in-memory store", func() {
					err := s.mem.Delete(uniqueUUID)
					sessionMap := s.mem.Get()
					Expect(err).To(BeNil())
					Expect(string(sessionMap[uniqueUUID].Oject)).To(BeEmpty())
				})
			})
		})
	})

	Describe("Extend session", func() {
		inMemResponse := "90660b89-100e-4f8f-9801-2524df6fbe34"
		uniqueUUID := "90660b89-100e-4f8f-9801-2524df6fbe34"
		BeforeEach(func() {
			err := s.mem.Commit(uniqueUUID, []byte(uniqueUUID), time.Now().Add(time.Minute))
			Expect(err).To(BeNil())
		})

		Context("Reset()", func() {
			When("the API os called with TTL as param", func() {
				It("extend sessionId in-memory store", func() {
					expiration := time.Now().Add(time.Minute * time.Duration(5))

					session, _, err := s.mem.Reset(uniqueUUID, expiration)
					sessionMap := s.mem.Get()
					Expect(err).To(BeNil())
					Expect(string(session)).To(Equal(inMemResponse))
					Expect(sessionMap[uniqueUUID].Expiration).To(Equal(expiration.UnixNano()))
				})
				It("extend sessionId in-memory store not found", func() {
					uniqueUUID2 := "90660b89-100e-4f8f-9801-2524df6fbe99"

					expiration := time.Now().Add(time.Minute * time.Duration(5))
					_, found, err := s.mem.Reset(uniqueUUID2, expiration)
					Expect(err).To(BeNil())
					Expect(found).ToNot(BeTrue())
				})
				It("extend sessionId in-memory store expired", func() {
					uniqueUUID2 := "90660b89-100e-4f8f-9801-2524df6fbe88"
					er := s.mem.Commit(uniqueUUID, []byte(uniqueUUID), time.Now().Add(100*time.Millisecond))
					time.Sleep(101 * time.Millisecond)

					expiration := time.Now().Add(time.Minute * time.Duration(5))
					_, found, err := s.mem.Reset(uniqueUUID2, expiration)
					Expect(err).To(BeNil())
					Expect(er).To(BeNil())
					Expect(found).ToNot(BeTrue())
				})
			})
		})
	})

	Describe("List session", func() {
		uniqueUUID1 := "90660b89-100e-4f8f-9801-2524df6fbe34"
		uniqueUUID2 := "90660b89-100e-4f8f-9801-2524df6fbe99"
		uniqueUUID3 := "90660b89-100e-4f8f-9801-2524df6fbe88"
		BeforeEach(func() {
			err := s.mem.Commit(uniqueUUID1, []byte(uniqueUUID1), time.Now().Add(time.Minute))
			Expect(err).To(BeNil())
			err = s.mem.Commit(uniqueUUID2, []byte(uniqueUUID2), time.Now().Add(time.Minute))
			Expect(err).To(BeNil())
			err = s.mem.Commit(uniqueUUID3, []byte(uniqueUUID3), time.Now().Add(time.Minute))
			Expect(err).To(BeNil())
		})

		Context("Lit()", func() {
			When("the API os called", func() {
				It("deletes sessionId in-memory store", func() {
					sessions, err := s.mem.List()
					sessionMap := s.mem.Get()
					Expect(err).To(BeNil())
					Expect(sessions).To(Equal(sessionMap))
				})
			})
		})
	})

	Describe("Delete Session Expired", func() {
		uniqueUUID1 := "90660b89-100e-4f8f-9801-2524df6fbe34"
		uniqueUUID2 := "90660b89-100e-4f8f-9801-2524df6fbe99"
		uniqueUUID3 := "90660b89-100e-4f8f-9801-2524df6fbe88"
		BeforeEach(func() {
			s.mem = NewInMemStore(101 * time.Millisecond, s.logger)
			err := s.mem.Commit(uniqueUUID1, []byte(uniqueUUID1), time.Now().Add(time.Minute))
			Expect(err).To(BeNil())
			err = s.mem.Commit(uniqueUUID2, []byte(uniqueUUID2), time.Now().Add(100*time.Millisecond))
			Expect(err).To(BeNil())
			err = s.mem.Commit(uniqueUUID3, []byte(uniqueUUID3), time.Now().Add(time.Minute))
			Expect(err).To(BeNil())
		})

		Context("DeleteSessionExpired()", func() {
			When("the API os called", func() {
				It("deletes sessionId in-memory store", func() {
					sessionMap := s.mem.Get()
					time.Sleep(200 * time.Millisecond)

					Expect(string(sessionMap[uniqueUUID2].Oject)).To(BeEmpty())
				})
			})
		})
	})

})