// Code generated by counterfeiter. DO NOT EDIT.
package inputconfigfakes

import (
	sync "sync"

	atc "github.com/concourse/concourse/atc"
	algorithm "github.com/concourse/concourse/atc/db/algorithm"
	inputconfig "github.com/concourse/concourse/atc/scheduler/inputmapper/inputconfig"
)

type FakeTransformer struct {
	TransformInputConfigsStub        func(*algorithm.VersionsDB, string, []atc.JobInput) (algorithm.InputConfigs, error)
	transformInputConfigsMutex       sync.RWMutex
	transformInputConfigsArgsForCall []struct {
		arg1 *algorithm.VersionsDB
		arg2 string
		arg3 []atc.JobInput
	}
	transformInputConfigsReturns struct {
		result1 algorithm.InputConfigs
		result2 error
	}
	transformInputConfigsReturnsOnCall map[int]struct {
		result1 algorithm.InputConfigs
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTransformer) TransformInputConfigs(arg1 *algorithm.VersionsDB, arg2 string, arg3 []atc.JobInput) (algorithm.InputConfigs, error) {
	var arg3Copy []atc.JobInput
	if arg3 != nil {
		arg3Copy = make([]atc.JobInput, len(arg3))
		copy(arg3Copy, arg3)
	}
	fake.transformInputConfigsMutex.Lock()
	ret, specificReturn := fake.transformInputConfigsReturnsOnCall[len(fake.transformInputConfigsArgsForCall)]
	fake.transformInputConfigsArgsForCall = append(fake.transformInputConfigsArgsForCall, struct {
		arg1 *algorithm.VersionsDB
		arg2 string
		arg3 []atc.JobInput
	}{arg1, arg2, arg3Copy})
	fake.recordInvocation("TransformInputConfigs", []interface{}{arg1, arg2, arg3Copy})
	fake.transformInputConfigsMutex.Unlock()
	if fake.TransformInputConfigsStub != nil {
		return fake.TransformInputConfigsStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.transformInputConfigsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeTransformer) TransformInputConfigsCallCount() int {
	fake.transformInputConfigsMutex.RLock()
	defer fake.transformInputConfigsMutex.RUnlock()
	return len(fake.transformInputConfigsArgsForCall)
}

func (fake *FakeTransformer) TransformInputConfigsArgsForCall(i int) (*algorithm.VersionsDB, string, []atc.JobInput) {
	fake.transformInputConfigsMutex.RLock()
	defer fake.transformInputConfigsMutex.RUnlock()
	argsForCall := fake.transformInputConfigsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeTransformer) TransformInputConfigsReturns(result1 algorithm.InputConfigs, result2 error) {
	fake.TransformInputConfigsStub = nil
	fake.transformInputConfigsReturns = struct {
		result1 algorithm.InputConfigs
		result2 error
	}{result1, result2}
}

func (fake *FakeTransformer) TransformInputConfigsReturnsOnCall(i int, result1 algorithm.InputConfigs, result2 error) {
	fake.TransformInputConfigsStub = nil
	if fake.transformInputConfigsReturnsOnCall == nil {
		fake.transformInputConfigsReturnsOnCall = make(map[int]struct {
			result1 algorithm.InputConfigs
			result2 error
		})
	}
	fake.transformInputConfigsReturnsOnCall[i] = struct {
		result1 algorithm.InputConfigs
		result2 error
	}{result1, result2}
}

func (fake *FakeTransformer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.transformInputConfigsMutex.RLock()
	defer fake.transformInputConfigsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeTransformer) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ inputconfig.Transformer = new(FakeTransformer)