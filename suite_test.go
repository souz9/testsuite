package testsuite

import "testing"

func TestSuite(t *testing.T) {
	s := &S{}
	Run(t, s)

	if s.beforeAll == 1 && s.beforeEach == 2 && s.testA == 1 && s.testB == 1 {
	} else {
		t.Fatalf("BeforeAll=%v BeforeEach=%v TestA=%v TestB=%v",
			s.beforeAll, s.beforeEach, s.testA, s.testB)
	}
}

type S struct {
	beforeAll, beforeEach, testA, testB int
}

func (s *S) BeforeAll(t *testing.T)  { s.beforeAll++ }
func (s *S) BeforeEach(t *testing.T) { s.beforeEach++ }

func (s *S) TestA(t *testing.T) { s.testA++ }
func (s *S) TestB(t *testing.T) { s.testB++ }
