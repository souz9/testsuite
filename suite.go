package testsuite

import (
	"reflect"
	"strings"
	"testing"
)

const MethodTestPrefix = "Test"

type Suite interface {
	BeforeAll(*testing.T)
	BeforeEach(*testing.T)
}

func Run(t *testing.T, suite Suite) {
	suite.BeforeAll(t)

	suiteT := reflect.TypeOf(suite)
	for i := 0; i < suiteT.NumMethod(); i++ {
		method := suiteT.Method(i)
		if name := method.Name; strings.HasPrefix(name, MethodTestPrefix) {
			name := strings.TrimPrefix(name, MethodTestPrefix)
			t.Run(name, func(t *testing.T) {
				suite.BeforeEach(t)

				in := []reflect.Value{reflect.ValueOf(suite), reflect.ValueOf(t)}
				method.Func.Call(in)
			})
		}
	}
}
