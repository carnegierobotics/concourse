// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/concourse/atc/builds"
	"github.com/concourse/atc/db"
	"github.com/concourse/atc/event"
)

type FakeBuildsDB struct {
	GetBuildStub        func(buildID int) (builds.Build, error)
	getBuildMutex       sync.RWMutex
	getBuildArgsForCall []struct {
		buildID int
	}
	getBuildReturns struct {
		result1 builds.Build
		result2 error
	}
	GetBuildEventsStub        func(buildID int) ([]db.BuildEvent, error)
	getBuildEventsMutex       sync.RWMutex
	getBuildEventsArgsForCall []struct {
		buildID int
	}
	getBuildEventsReturns struct {
		result1 []db.BuildEvent
		result2 error
	}
}

func (fake *FakeBuildsDB) GetBuild(buildID int) (builds.Build, error) {
	fake.getBuildMutex.Lock()
	fake.getBuildArgsForCall = append(fake.getBuildArgsForCall, struct {
		buildID int
	}{buildID})
	fake.getBuildMutex.Unlock()
	if fake.GetBuildStub != nil {
		return fake.GetBuildStub(buildID)
	} else {
		return fake.getBuildReturns.result1, fake.getBuildReturns.result2
	}
}

func (fake *FakeBuildsDB) GetBuildCallCount() int {
	fake.getBuildMutex.RLock()
	defer fake.getBuildMutex.RUnlock()
	return len(fake.getBuildArgsForCall)
}

func (fake *FakeBuildsDB) GetBuildArgsForCall(i int) int {
	fake.getBuildMutex.RLock()
	defer fake.getBuildMutex.RUnlock()
	return fake.getBuildArgsForCall[i].buildID
}

func (fake *FakeBuildsDB) GetBuildReturns(result1 builds.Build, result2 error) {
	fake.GetBuildStub = nil
	fake.getBuildReturns = struct {
		result1 builds.Build
		result2 error
	}{result1, result2}
}

func (fake *FakeBuildsDB) GetBuildEvents(buildID int) ([]db.BuildEvent, error) {
	fake.getBuildEventsMutex.Lock()
	fake.getBuildEventsArgsForCall = append(fake.getBuildEventsArgsForCall, struct {
		buildID int
	}{buildID})
	fake.getBuildEventsMutex.Unlock()
	if fake.GetBuildEventsStub != nil {
		return fake.GetBuildEventsStub(buildID)
	} else {
		return fake.getBuildEventsReturns.result1, fake.getBuildEventsReturns.result2
	}
}

func (fake *FakeBuildsDB) GetBuildEventsCallCount() int {
	fake.getBuildEventsMutex.RLock()
	defer fake.getBuildEventsMutex.RUnlock()
	return len(fake.getBuildEventsArgsForCall)
}

func (fake *FakeBuildsDB) GetBuildEventsArgsForCall(i int) int {
	fake.getBuildEventsMutex.RLock()
	defer fake.getBuildEventsMutex.RUnlock()
	return fake.getBuildEventsArgsForCall[i].buildID
}

func (fake *FakeBuildsDB) GetBuildEventsReturns(result1 []db.BuildEvent, result2 error) {
	fake.GetBuildEventsStub = nil
	fake.getBuildEventsReturns = struct {
		result1 []db.BuildEvent
		result2 error
	}{result1, result2}
}

var _ event.BuildsDB = new(FakeBuildsDB)
