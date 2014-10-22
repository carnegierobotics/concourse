// This file was generated by counterfeiter
package fakes

import (
	"sync"
	"time"

	"github.com/concourse/atc/builds"
	"github.com/concourse/atc/db"
	"github.com/concourse/atc/scheduler"
)

type FakeTrackerDB struct {
	SaveBuildEventStub        func(buildID int, event db.BuildEvent) error
	saveBuildEventMutex       sync.RWMutex
	saveBuildEventArgsForCall []struct {
		buildID int
		event   db.BuildEvent
	}
	saveBuildEventReturns struct {
		result1 error
	}
	SaveBuildStartTimeStub        func(buildID int, startTime time.Time) error
	saveBuildStartTimeMutex       sync.RWMutex
	saveBuildStartTimeArgsForCall []struct {
		buildID   int
		startTime time.Time
	}
	saveBuildStartTimeReturns struct {
		result1 error
	}
	SaveBuildEndTimeStub        func(buildID int, startTime time.Time) error
	saveBuildEndTimeMutex       sync.RWMutex
	saveBuildEndTimeArgsForCall []struct {
		buildID   int
		startTime time.Time
	}
	saveBuildEndTimeReturns struct {
		result1 error
	}
	SaveBuildInputStub        func(buildID int, vr builds.VersionedResource) error
	saveBuildInputMutex       sync.RWMutex
	saveBuildInputArgsForCall []struct {
		buildID int
		vr      builds.VersionedResource
	}
	saveBuildInputReturns struct {
		result1 error
	}
	SaveBuildOutputStub        func(buildID int, vr builds.VersionedResource) error
	saveBuildOutputMutex       sync.RWMutex
	saveBuildOutputArgsForCall []struct {
		buildID int
		vr      builds.VersionedResource
	}
	saveBuildOutputReturns struct {
		result1 error
	}
	SaveBuildStatusStub        func(buildID int, status builds.Status) error
	saveBuildStatusMutex       sync.RWMutex
	saveBuildStatusArgsForCall []struct {
		buildID int
		status  builds.Status
	}
	saveBuildStatusReturns struct {
		result1 error
	}
}

func (fake *FakeTrackerDB) SaveBuildEvent(buildID int, event db.BuildEvent) error {
	fake.saveBuildEventMutex.Lock()
	fake.saveBuildEventArgsForCall = append(fake.saveBuildEventArgsForCall, struct {
		buildID int
		event   db.BuildEvent
	}{buildID, event})
	fake.saveBuildEventMutex.Unlock()
	if fake.SaveBuildEventStub != nil {
		return fake.SaveBuildEventStub(buildID, event)
	} else {
		return fake.saveBuildEventReturns.result1
	}
}

func (fake *FakeTrackerDB) SaveBuildEventCallCount() int {
	fake.saveBuildEventMutex.RLock()
	defer fake.saveBuildEventMutex.RUnlock()
	return len(fake.saveBuildEventArgsForCall)
}

func (fake *FakeTrackerDB) SaveBuildEventArgsForCall(i int) (int, db.BuildEvent) {
	fake.saveBuildEventMutex.RLock()
	defer fake.saveBuildEventMutex.RUnlock()
	return fake.saveBuildEventArgsForCall[i].buildID, fake.saveBuildEventArgsForCall[i].event
}

func (fake *FakeTrackerDB) SaveBuildEventReturns(result1 error) {
	fake.SaveBuildEventStub = nil
	fake.saveBuildEventReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeTrackerDB) SaveBuildStartTime(buildID int, startTime time.Time) error {
	fake.saveBuildStartTimeMutex.Lock()
	fake.saveBuildStartTimeArgsForCall = append(fake.saveBuildStartTimeArgsForCall, struct {
		buildID   int
		startTime time.Time
	}{buildID, startTime})
	fake.saveBuildStartTimeMutex.Unlock()
	if fake.SaveBuildStartTimeStub != nil {
		return fake.SaveBuildStartTimeStub(buildID, startTime)
	} else {
		return fake.saveBuildStartTimeReturns.result1
	}
}

func (fake *FakeTrackerDB) SaveBuildStartTimeCallCount() int {
	fake.saveBuildStartTimeMutex.RLock()
	defer fake.saveBuildStartTimeMutex.RUnlock()
	return len(fake.saveBuildStartTimeArgsForCall)
}

func (fake *FakeTrackerDB) SaveBuildStartTimeArgsForCall(i int) (int, time.Time) {
	fake.saveBuildStartTimeMutex.RLock()
	defer fake.saveBuildStartTimeMutex.RUnlock()
	return fake.saveBuildStartTimeArgsForCall[i].buildID, fake.saveBuildStartTimeArgsForCall[i].startTime
}

func (fake *FakeTrackerDB) SaveBuildStartTimeReturns(result1 error) {
	fake.SaveBuildStartTimeStub = nil
	fake.saveBuildStartTimeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeTrackerDB) SaveBuildEndTime(buildID int, startTime time.Time) error {
	fake.saveBuildEndTimeMutex.Lock()
	fake.saveBuildEndTimeArgsForCall = append(fake.saveBuildEndTimeArgsForCall, struct {
		buildID   int
		startTime time.Time
	}{buildID, startTime})
	fake.saveBuildEndTimeMutex.Unlock()
	if fake.SaveBuildEndTimeStub != nil {
		return fake.SaveBuildEndTimeStub(buildID, startTime)
	} else {
		return fake.saveBuildEndTimeReturns.result1
	}
}

func (fake *FakeTrackerDB) SaveBuildEndTimeCallCount() int {
	fake.saveBuildEndTimeMutex.RLock()
	defer fake.saveBuildEndTimeMutex.RUnlock()
	return len(fake.saveBuildEndTimeArgsForCall)
}

func (fake *FakeTrackerDB) SaveBuildEndTimeArgsForCall(i int) (int, time.Time) {
	fake.saveBuildEndTimeMutex.RLock()
	defer fake.saveBuildEndTimeMutex.RUnlock()
	return fake.saveBuildEndTimeArgsForCall[i].buildID, fake.saveBuildEndTimeArgsForCall[i].startTime
}

func (fake *FakeTrackerDB) SaveBuildEndTimeReturns(result1 error) {
	fake.SaveBuildEndTimeStub = nil
	fake.saveBuildEndTimeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeTrackerDB) SaveBuildInput(buildID int, vr builds.VersionedResource) error {
	fake.saveBuildInputMutex.Lock()
	fake.saveBuildInputArgsForCall = append(fake.saveBuildInputArgsForCall, struct {
		buildID int
		vr      builds.VersionedResource
	}{buildID, vr})
	fake.saveBuildInputMutex.Unlock()
	if fake.SaveBuildInputStub != nil {
		return fake.SaveBuildInputStub(buildID, vr)
	} else {
		return fake.saveBuildInputReturns.result1
	}
}

func (fake *FakeTrackerDB) SaveBuildInputCallCount() int {
	fake.saveBuildInputMutex.RLock()
	defer fake.saveBuildInputMutex.RUnlock()
	return len(fake.saveBuildInputArgsForCall)
}

func (fake *FakeTrackerDB) SaveBuildInputArgsForCall(i int) (int, builds.VersionedResource) {
	fake.saveBuildInputMutex.RLock()
	defer fake.saveBuildInputMutex.RUnlock()
	return fake.saveBuildInputArgsForCall[i].buildID, fake.saveBuildInputArgsForCall[i].vr
}

func (fake *FakeTrackerDB) SaveBuildInputReturns(result1 error) {
	fake.SaveBuildInputStub = nil
	fake.saveBuildInputReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeTrackerDB) SaveBuildOutput(buildID int, vr builds.VersionedResource) error {
	fake.saveBuildOutputMutex.Lock()
	fake.saveBuildOutputArgsForCall = append(fake.saveBuildOutputArgsForCall, struct {
		buildID int
		vr      builds.VersionedResource
	}{buildID, vr})
	fake.saveBuildOutputMutex.Unlock()
	if fake.SaveBuildOutputStub != nil {
		return fake.SaveBuildOutputStub(buildID, vr)
	} else {
		return fake.saveBuildOutputReturns.result1
	}
}

func (fake *FakeTrackerDB) SaveBuildOutputCallCount() int {
	fake.saveBuildOutputMutex.RLock()
	defer fake.saveBuildOutputMutex.RUnlock()
	return len(fake.saveBuildOutputArgsForCall)
}

func (fake *FakeTrackerDB) SaveBuildOutputArgsForCall(i int) (int, builds.VersionedResource) {
	fake.saveBuildOutputMutex.RLock()
	defer fake.saveBuildOutputMutex.RUnlock()
	return fake.saveBuildOutputArgsForCall[i].buildID, fake.saveBuildOutputArgsForCall[i].vr
}

func (fake *FakeTrackerDB) SaveBuildOutputReturns(result1 error) {
	fake.SaveBuildOutputStub = nil
	fake.saveBuildOutputReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeTrackerDB) SaveBuildStatus(buildID int, status builds.Status) error {
	fake.saveBuildStatusMutex.Lock()
	fake.saveBuildStatusArgsForCall = append(fake.saveBuildStatusArgsForCall, struct {
		buildID int
		status  builds.Status
	}{buildID, status})
	fake.saveBuildStatusMutex.Unlock()
	if fake.SaveBuildStatusStub != nil {
		return fake.SaveBuildStatusStub(buildID, status)
	} else {
		return fake.saveBuildStatusReturns.result1
	}
}

func (fake *FakeTrackerDB) SaveBuildStatusCallCount() int {
	fake.saveBuildStatusMutex.RLock()
	defer fake.saveBuildStatusMutex.RUnlock()
	return len(fake.saveBuildStatusArgsForCall)
}

func (fake *FakeTrackerDB) SaveBuildStatusArgsForCall(i int) (int, builds.Status) {
	fake.saveBuildStatusMutex.RLock()
	defer fake.saveBuildStatusMutex.RUnlock()
	return fake.saveBuildStatusArgsForCall[i].buildID, fake.saveBuildStatusArgsForCall[i].status
}

func (fake *FakeTrackerDB) SaveBuildStatusReturns(result1 error) {
	fake.SaveBuildStatusStub = nil
	fake.saveBuildStatusReturns = struct {
		result1 error
	}{result1}
}

var _ scheduler.TrackerDB = new(FakeTrackerDB)
