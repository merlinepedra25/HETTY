// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package reqlog_test

import (
	"context"
	"github.com/dstotijn/hetty/pkg/reqlog"
	"github.com/dstotijn/hetty/pkg/scope"
	"github.com/oklog/ulid"
	"sync"
)

// Ensure, that RepoMock does implement reqlog.Repository.
// If this is not the case, regenerate this file with moq.
var _ reqlog.Repository = &RepoMock{}

// RepoMock is a mock implementation of reqlog.Repository.
//
// 	func TestSomethingThatUsesRepository(t *testing.T) {
//
// 		// make and configure a mocked reqlog.Repository
// 		mockedRepository := &RepoMock{
// 			ClearRequestLogsFunc: func(ctx context.Context, projectID ulid.ULID) error {
// 				panic("mock out the ClearRequestLogs method")
// 			},
// 			FindRequestLogByIDFunc: func(ctx context.Context, id ulid.ULID) (reqlog.RequestLog, error) {
// 				panic("mock out the FindRequestLogByID method")
// 			},
// 			FindRequestLogsFunc: func(ctx context.Context, filter reqlog.FindRequestsFilter, scopeMoqParam *scope.Scope) ([]reqlog.RequestLog, error) {
// 				panic("mock out the FindRequestLogs method")
// 			},
// 			StoreRequestLogFunc: func(ctx context.Context, reqLog reqlog.RequestLog) error {
// 				panic("mock out the StoreRequestLog method")
// 			},
// 			StoreResponseLogFunc: func(ctx context.Context, reqLogID ulid.ULID, resLog reqlog.ResponseLog) error {
// 				panic("mock out the StoreResponseLog method")
// 			},
// 		}
//
// 		// use mockedRepository in code that requires reqlog.Repository
// 		// and then make assertions.
//
// 	}
type RepoMock struct {
	// ClearRequestLogsFunc mocks the ClearRequestLogs method.
	ClearRequestLogsFunc func(ctx context.Context, projectID ulid.ULID) error

	// FindRequestLogByIDFunc mocks the FindRequestLogByID method.
	FindRequestLogByIDFunc func(ctx context.Context, id ulid.ULID) (reqlog.RequestLog, error)

	// FindRequestLogsFunc mocks the FindRequestLogs method.
	FindRequestLogsFunc func(ctx context.Context, filter reqlog.FindRequestsFilter, scopeMoqParam *scope.Scope) ([]reqlog.RequestLog, error)

	// StoreRequestLogFunc mocks the StoreRequestLog method.
	StoreRequestLogFunc func(ctx context.Context, reqLog reqlog.RequestLog) error

	// StoreResponseLogFunc mocks the StoreResponseLog method.
	StoreResponseLogFunc func(ctx context.Context, reqLogID ulid.ULID, resLog reqlog.ResponseLog) error

	// calls tracks calls to the methods.
	calls struct {
		// ClearRequestLogs holds details about calls to the ClearRequestLogs method.
		ClearRequestLogs []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ProjectID is the projectID argument value.
			ProjectID ulid.ULID
		}
		// FindRequestLogByID holds details about calls to the FindRequestLogByID method.
		FindRequestLogByID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID ulid.ULID
		}
		// FindRequestLogs holds details about calls to the FindRequestLogs method.
		FindRequestLogs []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Filter is the filter argument value.
			Filter reqlog.FindRequestsFilter
			// ScopeMoqParam is the scopeMoqParam argument value.
			ScopeMoqParam *scope.Scope
		}
		// StoreRequestLog holds details about calls to the StoreRequestLog method.
		StoreRequestLog []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ReqLog is the reqLog argument value.
			ReqLog reqlog.RequestLog
		}
		// StoreResponseLog holds details about calls to the StoreResponseLog method.
		StoreResponseLog []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ReqLogID is the reqLogID argument value.
			ReqLogID ulid.ULID
			// ResLog is the resLog argument value.
			ResLog reqlog.ResponseLog
		}
	}
	lockClearRequestLogs   sync.RWMutex
	lockFindRequestLogByID sync.RWMutex
	lockFindRequestLogs    sync.RWMutex
	lockStoreRequestLog    sync.RWMutex
	lockStoreResponseLog   sync.RWMutex
}

// ClearRequestLogs calls ClearRequestLogsFunc.
func (mock *RepoMock) ClearRequestLogs(ctx context.Context, projectID ulid.ULID) error {
	if mock.ClearRequestLogsFunc == nil {
		panic("RepoMock.ClearRequestLogsFunc: method is nil but Repository.ClearRequestLogs was just called")
	}
	callInfo := struct {
		Ctx       context.Context
		ProjectID ulid.ULID
	}{
		Ctx:       ctx,
		ProjectID: projectID,
	}
	mock.lockClearRequestLogs.Lock()
	mock.calls.ClearRequestLogs = append(mock.calls.ClearRequestLogs, callInfo)
	mock.lockClearRequestLogs.Unlock()
	return mock.ClearRequestLogsFunc(ctx, projectID)
}

// ClearRequestLogsCalls gets all the calls that were made to ClearRequestLogs.
// Check the length with:
//     len(mockedRepository.ClearRequestLogsCalls())
func (mock *RepoMock) ClearRequestLogsCalls() []struct {
	Ctx       context.Context
	ProjectID ulid.ULID
} {
	var calls []struct {
		Ctx       context.Context
		ProjectID ulid.ULID
	}
	mock.lockClearRequestLogs.RLock()
	calls = mock.calls.ClearRequestLogs
	mock.lockClearRequestLogs.RUnlock()
	return calls
}

// FindRequestLogByID calls FindRequestLogByIDFunc.
func (mock *RepoMock) FindRequestLogByID(ctx context.Context, id ulid.ULID) (reqlog.RequestLog, error) {
	if mock.FindRequestLogByIDFunc == nil {
		panic("RepoMock.FindRequestLogByIDFunc: method is nil but Repository.FindRequestLogByID was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  ulid.ULID
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockFindRequestLogByID.Lock()
	mock.calls.FindRequestLogByID = append(mock.calls.FindRequestLogByID, callInfo)
	mock.lockFindRequestLogByID.Unlock()
	return mock.FindRequestLogByIDFunc(ctx, id)
}

// FindRequestLogByIDCalls gets all the calls that were made to FindRequestLogByID.
// Check the length with:
//     len(mockedRepository.FindRequestLogByIDCalls())
func (mock *RepoMock) FindRequestLogByIDCalls() []struct {
	Ctx context.Context
	ID  ulid.ULID
} {
	var calls []struct {
		Ctx context.Context
		ID  ulid.ULID
	}
	mock.lockFindRequestLogByID.RLock()
	calls = mock.calls.FindRequestLogByID
	mock.lockFindRequestLogByID.RUnlock()
	return calls
}

// FindRequestLogs calls FindRequestLogsFunc.
func (mock *RepoMock) FindRequestLogs(ctx context.Context, filter reqlog.FindRequestsFilter, scopeMoqParam *scope.Scope) ([]reqlog.RequestLog, error) {
	if mock.FindRequestLogsFunc == nil {
		panic("RepoMock.FindRequestLogsFunc: method is nil but Repository.FindRequestLogs was just called")
	}
	callInfo := struct {
		Ctx           context.Context
		Filter        reqlog.FindRequestsFilter
		ScopeMoqParam *scope.Scope
	}{
		Ctx:           ctx,
		Filter:        filter,
		ScopeMoqParam: scopeMoqParam,
	}
	mock.lockFindRequestLogs.Lock()
	mock.calls.FindRequestLogs = append(mock.calls.FindRequestLogs, callInfo)
	mock.lockFindRequestLogs.Unlock()
	return mock.FindRequestLogsFunc(ctx, filter, scopeMoqParam)
}

// FindRequestLogsCalls gets all the calls that were made to FindRequestLogs.
// Check the length with:
//     len(mockedRepository.FindRequestLogsCalls())
func (mock *RepoMock) FindRequestLogsCalls() []struct {
	Ctx           context.Context
	Filter        reqlog.FindRequestsFilter
	ScopeMoqParam *scope.Scope
} {
	var calls []struct {
		Ctx           context.Context
		Filter        reqlog.FindRequestsFilter
		ScopeMoqParam *scope.Scope
	}
	mock.lockFindRequestLogs.RLock()
	calls = mock.calls.FindRequestLogs
	mock.lockFindRequestLogs.RUnlock()
	return calls
}

// StoreRequestLog calls StoreRequestLogFunc.
func (mock *RepoMock) StoreRequestLog(ctx context.Context, reqLog reqlog.RequestLog) error {
	if mock.StoreRequestLogFunc == nil {
		panic("RepoMock.StoreRequestLogFunc: method is nil but Repository.StoreRequestLog was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		ReqLog reqlog.RequestLog
	}{
		Ctx:    ctx,
		ReqLog: reqLog,
	}
	mock.lockStoreRequestLog.Lock()
	mock.calls.StoreRequestLog = append(mock.calls.StoreRequestLog, callInfo)
	mock.lockStoreRequestLog.Unlock()
	return mock.StoreRequestLogFunc(ctx, reqLog)
}

// StoreRequestLogCalls gets all the calls that were made to StoreRequestLog.
// Check the length with:
//     len(mockedRepository.StoreRequestLogCalls())
func (mock *RepoMock) StoreRequestLogCalls() []struct {
	Ctx    context.Context
	ReqLog reqlog.RequestLog
} {
	var calls []struct {
		Ctx    context.Context
		ReqLog reqlog.RequestLog
	}
	mock.lockStoreRequestLog.RLock()
	calls = mock.calls.StoreRequestLog
	mock.lockStoreRequestLog.RUnlock()
	return calls
}

// StoreResponseLog calls StoreResponseLogFunc.
func (mock *RepoMock) StoreResponseLog(ctx context.Context, reqLogID ulid.ULID, resLog reqlog.ResponseLog) error {
	if mock.StoreResponseLogFunc == nil {
		panic("RepoMock.StoreResponseLogFunc: method is nil but Repository.StoreResponseLog was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		ReqLogID ulid.ULID
		ResLog   reqlog.ResponseLog
	}{
		Ctx:      ctx,
		ReqLogID: reqLogID,
		ResLog:   resLog,
	}
	mock.lockStoreResponseLog.Lock()
	mock.calls.StoreResponseLog = append(mock.calls.StoreResponseLog, callInfo)
	mock.lockStoreResponseLog.Unlock()
	return mock.StoreResponseLogFunc(ctx, reqLogID, resLog)
}

// StoreResponseLogCalls gets all the calls that were made to StoreResponseLog.
// Check the length with:
//     len(mockedRepository.StoreResponseLogCalls())
func (mock *RepoMock) StoreResponseLogCalls() []struct {
	Ctx      context.Context
	ReqLogID ulid.ULID
	ResLog   reqlog.ResponseLog
} {
	var calls []struct {
		Ctx      context.Context
		ReqLogID ulid.ULID
		ResLog   reqlog.ResponseLog
	}
	mock.lockStoreResponseLog.RLock()
	calls = mock.calls.StoreResponseLog
	mock.lockStoreResponseLog.RUnlock()
	return calls
}
