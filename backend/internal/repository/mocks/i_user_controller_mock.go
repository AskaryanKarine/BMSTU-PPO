package mocks

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i DoramaSet/internal/interfaces/controller.IUserController -o ./i_user_controller_mock_test.go -n IUserControllerMock

import (
	"DoramaSet/internal/logic/model"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// IUserControllerMock implements IUserController
type IUserControllerMock struct {
	t minimock.Tester

	funcAuthByToken          func(token string) (up1 *model.User, err error)
	inspectFuncAuthByToken   func(token string)
	afterAuthByTokenCounter  uint64
	beforeAuthByTokenCounter uint64
	AuthByTokenMock          mIUserControllerMockAuthByToken

	funcLogin          func(username string, password string) (s1 string, err error)
	inspectFuncLogin   func(username string, password string)
	afterLoginCounter  uint64
	beforeLoginCounter uint64
	LoginMock          mIUserControllerMockLogin

	funcRegistration          func(record *model.User) (s1 string, err error)
	inspectFuncRegistration   func(record *model.User)
	afterRegistrationCounter  uint64
	beforeRegistrationCounter uint64
	RegistrationMock          mIUserControllerMockRegistration

	funcUpdateActive          func(token string) (err error)
	inspectFuncUpdateActive   func(token string)
	afterUpdateActiveCounter  uint64
	beforeUpdateActiveCounter uint64
	UpdateActiveMock          mIUserControllerMockUpdateActive
}

// NewIUserControllerMock returns a mock for IUserController
func NewIUserControllerMock(t minimock.Tester) *IUserControllerMock {
	m := &IUserControllerMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.AuthByTokenMock = mIUserControllerMockAuthByToken{mock: m}
	m.AuthByTokenMock.callArgs = []*IUserControllerMockAuthByTokenParams{}

	m.LoginMock = mIUserControllerMockLogin{mock: m}
	m.LoginMock.callArgs = []*IUserControllerMockLoginParams{}

	m.RegistrationMock = mIUserControllerMockRegistration{mock: m}
	m.RegistrationMock.callArgs = []*IUserControllerMockRegistrationParams{}

	m.UpdateActiveMock = mIUserControllerMockUpdateActive{mock: m}
	m.UpdateActiveMock.callArgs = []*IUserControllerMockUpdateActiveParams{}

	return m
}

type mIUserControllerMockAuthByToken struct {
	mock               *IUserControllerMock
	defaultExpectation *IUserControllerMockAuthByTokenExpectation
	expectations       []*IUserControllerMockAuthByTokenExpectation

	callArgs []*IUserControllerMockAuthByTokenParams
	mutex    sync.RWMutex
}

// IUserControllerMockAuthByTokenExpectation specifies expectation struct of the IUserController.AuthByToken
type IUserControllerMockAuthByTokenExpectation struct {
	mock    *IUserControllerMock
	params  *IUserControllerMockAuthByTokenParams
	results *IUserControllerMockAuthByTokenResults
	Counter uint64
}

// IUserControllerMockAuthByTokenParams contains parameters of the IUserController.AuthByToken
type IUserControllerMockAuthByTokenParams struct {
	token string
}

// IUserControllerMockAuthByTokenResults contains results of the IUserController.AuthByToken
type IUserControllerMockAuthByTokenResults struct {
	up1 *model.User
	err error
}

// Expect sets up expected params for IUserController.AuthByToken
func (mmAuthByToken *mIUserControllerMockAuthByToken) Expect(token string) *mIUserControllerMockAuthByToken {
	if mmAuthByToken.mock.funcAuthByToken != nil {
		mmAuthByToken.mock.t.Fatalf("IUserControllerMock.AuthByToken mock is already set by Set")
	}

	if mmAuthByToken.defaultExpectation == nil {
		mmAuthByToken.defaultExpectation = &IUserControllerMockAuthByTokenExpectation{}
	}

	mmAuthByToken.defaultExpectation.params = &IUserControllerMockAuthByTokenParams{token}
	for _, e := range mmAuthByToken.expectations {
		if minimock.Equal(e.params, mmAuthByToken.defaultExpectation.params) {
			mmAuthByToken.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmAuthByToken.defaultExpectation.params)
		}
	}

	return mmAuthByToken
}

// Inspect accepts an inspector function that has same arguments as the IUserController.AuthByToken
func (mmAuthByToken *mIUserControllerMockAuthByToken) Inspect(f func(token string)) *mIUserControllerMockAuthByToken {
	if mmAuthByToken.mock.inspectFuncAuthByToken != nil {
		mmAuthByToken.mock.t.Fatalf("Inspect function is already set for IUserControllerMock.AuthByToken")
	}

	mmAuthByToken.mock.inspectFuncAuthByToken = f

	return mmAuthByToken
}

// Return sets up results that will be returned by IUserController.AuthByToken
func (mmAuthByToken *mIUserControllerMockAuthByToken) Return(up1 *model.User, err error) *IUserControllerMock {
	if mmAuthByToken.mock.funcAuthByToken != nil {
		mmAuthByToken.mock.t.Fatalf("IUserControllerMock.AuthByToken mock is already set by Set")
	}

	if mmAuthByToken.defaultExpectation == nil {
		mmAuthByToken.defaultExpectation = &IUserControllerMockAuthByTokenExpectation{mock: mmAuthByToken.mock}
	}
	mmAuthByToken.defaultExpectation.results = &IUserControllerMockAuthByTokenResults{up1, err}
	return mmAuthByToken.mock
}

// Set uses given function f to mock the IUserController.AuthByToken method
func (mmAuthByToken *mIUserControllerMockAuthByToken) Set(f func(token string) (up1 *model.User, err error)) *IUserControllerMock {
	if mmAuthByToken.defaultExpectation != nil {
		mmAuthByToken.mock.t.Fatalf("Default expectation is already set for the IUserController.AuthByToken method")
	}

	if len(mmAuthByToken.expectations) > 0 {
		mmAuthByToken.mock.t.Fatalf("Some expectations are already set for the IUserController.AuthByToken method")
	}

	mmAuthByToken.mock.funcAuthByToken = f
	return mmAuthByToken.mock
}

// When sets expectation for the IUserController.AuthByToken which will trigger the result defined by the following
// Then helper
func (mmAuthByToken *mIUserControllerMockAuthByToken) When(token string) *IUserControllerMockAuthByTokenExpectation {
	if mmAuthByToken.mock.funcAuthByToken != nil {
		mmAuthByToken.mock.t.Fatalf("IUserControllerMock.AuthByToken mock is already set by Set")
	}

	expectation := &IUserControllerMockAuthByTokenExpectation{
		mock:   mmAuthByToken.mock,
		params: &IUserControllerMockAuthByTokenParams{token},
	}
	mmAuthByToken.expectations = append(mmAuthByToken.expectations, expectation)
	return expectation
}

// Then sets up IUserController.AuthByToken return parameters for the expectation previously defined by the When method
func (e *IUserControllerMockAuthByTokenExpectation) Then(up1 *model.User, err error) *IUserControllerMock {
	e.results = &IUserControllerMockAuthByTokenResults{up1, err}
	return e.mock
}

// AuthByToken implements IUserController
func (mmAuthByToken *IUserControllerMock) AuthByToken(token string) (up1 *model.User, err error) {
	mm_atomic.AddUint64(&mmAuthByToken.beforeAuthByTokenCounter, 1)
	defer mm_atomic.AddUint64(&mmAuthByToken.afterAuthByTokenCounter, 1)

	if mmAuthByToken.inspectFuncAuthByToken != nil {
		mmAuthByToken.inspectFuncAuthByToken(token)
	}

	mm_params := &IUserControllerMockAuthByTokenParams{token}

	// Record call args
	mmAuthByToken.AuthByTokenMock.mutex.Lock()
	mmAuthByToken.AuthByTokenMock.callArgs = append(mmAuthByToken.AuthByTokenMock.callArgs, mm_params)
	mmAuthByToken.AuthByTokenMock.mutex.Unlock()

	for _, e := range mmAuthByToken.AuthByTokenMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.up1, e.results.err
		}
	}

	if mmAuthByToken.AuthByTokenMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmAuthByToken.AuthByTokenMock.defaultExpectation.Counter, 1)
		mm_want := mmAuthByToken.AuthByTokenMock.defaultExpectation.params
		mm_got := IUserControllerMockAuthByTokenParams{token}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmAuthByToken.t.Errorf("IUserControllerMock.AuthByToken got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmAuthByToken.AuthByTokenMock.defaultExpectation.results
		if mm_results == nil {
			mmAuthByToken.t.Fatal("No results are set for the IUserControllerMock.AuthByToken")
		}
		return (*mm_results).up1, (*mm_results).err
	}
	if mmAuthByToken.funcAuthByToken != nil {
		return mmAuthByToken.funcAuthByToken(token)
	}
	mmAuthByToken.t.Fatalf("Unexpected call to IUserControllerMock.AuthByToken. %v", token)
	return
}

// AuthByTokenAfterCounter returns a count of finished IUserControllerMock.AuthByToken invocations
func (mmAuthByToken *IUserControllerMock) AuthByTokenAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmAuthByToken.afterAuthByTokenCounter)
}

// AuthByTokenBeforeCounter returns a count of IUserControllerMock.AuthByToken invocations
func (mmAuthByToken *IUserControllerMock) AuthByTokenBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmAuthByToken.beforeAuthByTokenCounter)
}

// Calls returns a list of arguments used in each call to IUserControllerMock.AuthByToken.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmAuthByToken *mIUserControllerMockAuthByToken) Calls() []*IUserControllerMockAuthByTokenParams {
	mmAuthByToken.mutex.RLock()

	argCopy := make([]*IUserControllerMockAuthByTokenParams, len(mmAuthByToken.callArgs))
	copy(argCopy, mmAuthByToken.callArgs)

	mmAuthByToken.mutex.RUnlock()

	return argCopy
}

// MinimockAuthByTokenDone returns true if the count of the AuthByToken invocations corresponds
// the number of defined expectations
func (m *IUserControllerMock) MinimockAuthByTokenDone() bool {
	for _, e := range m.AuthByTokenMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.AuthByTokenMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterAuthByTokenCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcAuthByToken != nil && mm_atomic.LoadUint64(&m.afterAuthByTokenCounter) < 1 {
		return false
	}
	return true
}

// MinimockAuthByTokenInspect logs each unmet expectation
func (m *IUserControllerMock) MinimockAuthByTokenInspect() {
	for _, e := range m.AuthByTokenMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to IUserControllerMock.AuthByToken with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.AuthByTokenMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterAuthByTokenCounter) < 1 {
		if m.AuthByTokenMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to IUserControllerMock.AuthByToken")
		} else {
			m.t.Errorf("Expected call to IUserControllerMock.AuthByToken with params: %#v", *m.AuthByTokenMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcAuthByToken != nil && mm_atomic.LoadUint64(&m.afterAuthByTokenCounter) < 1 {
		m.t.Error("Expected call to IUserControllerMock.AuthByToken")
	}
}

type mIUserControllerMockLogin struct {
	mock               *IUserControllerMock
	defaultExpectation *IUserControllerMockLoginExpectation
	expectations       []*IUserControllerMockLoginExpectation

	callArgs []*IUserControllerMockLoginParams
	mutex    sync.RWMutex
}

// IUserControllerMockLoginExpectation specifies expectation struct of the IUserController.Login
type IUserControllerMockLoginExpectation struct {
	mock    *IUserControllerMock
	params  *IUserControllerMockLoginParams
	results *IUserControllerMockLoginResults
	Counter uint64
}

// IUserControllerMockLoginParams contains parameters of the IUserController.Login
type IUserControllerMockLoginParams struct {
	username string
	password string
}

// IUserControllerMockLoginResults contains results of the IUserController.Login
type IUserControllerMockLoginResults struct {
	s1  string
	err error
}

// Expect sets up expected params for IUserController.Login
func (mmLogin *mIUserControllerMockLogin) Expect(username string, password string) *mIUserControllerMockLogin {
	if mmLogin.mock.funcLogin != nil {
		mmLogin.mock.t.Fatalf("IUserControllerMock.Login mock is already set by Set")
	}

	if mmLogin.defaultExpectation == nil {
		mmLogin.defaultExpectation = &IUserControllerMockLoginExpectation{}
	}

	mmLogin.defaultExpectation.params = &IUserControllerMockLoginParams{username, password}
	for _, e := range mmLogin.expectations {
		if minimock.Equal(e.params, mmLogin.defaultExpectation.params) {
			mmLogin.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmLogin.defaultExpectation.params)
		}
	}

	return mmLogin
}

// Inspect accepts an inspector function that has same arguments as the IUserController.Login
func (mmLogin *mIUserControllerMockLogin) Inspect(f func(username string, password string)) *mIUserControllerMockLogin {
	if mmLogin.mock.inspectFuncLogin != nil {
		mmLogin.mock.t.Fatalf("Inspect function is already set for IUserControllerMock.Login")
	}

	mmLogin.mock.inspectFuncLogin = f

	return mmLogin
}

// Return sets up results that will be returned by IUserController.Login
func (mmLogin *mIUserControllerMockLogin) Return(s1 string, err error) *IUserControllerMock {
	if mmLogin.mock.funcLogin != nil {
		mmLogin.mock.t.Fatalf("IUserControllerMock.Login mock is already set by Set")
	}

	if mmLogin.defaultExpectation == nil {
		mmLogin.defaultExpectation = &IUserControllerMockLoginExpectation{mock: mmLogin.mock}
	}
	mmLogin.defaultExpectation.results = &IUserControllerMockLoginResults{s1, err}
	return mmLogin.mock
}

// Set uses given function f to mock the IUserController.Login method
func (mmLogin *mIUserControllerMockLogin) Set(f func(username string, password string) (s1 string, err error)) *IUserControllerMock {
	if mmLogin.defaultExpectation != nil {
		mmLogin.mock.t.Fatalf("Default expectation is already set for the IUserController.Login method")
	}

	if len(mmLogin.expectations) > 0 {
		mmLogin.mock.t.Fatalf("Some expectations are already set for the IUserController.Login method")
	}

	mmLogin.mock.funcLogin = f
	return mmLogin.mock
}

// When sets expectation for the IUserController.Login which will trigger the result defined by the following
// Then helper
func (mmLogin *mIUserControllerMockLogin) When(username string, password string) *IUserControllerMockLoginExpectation {
	if mmLogin.mock.funcLogin != nil {
		mmLogin.mock.t.Fatalf("IUserControllerMock.Login mock is already set by Set")
	}

	expectation := &IUserControllerMockLoginExpectation{
		mock:   mmLogin.mock,
		params: &IUserControllerMockLoginParams{username, password},
	}
	mmLogin.expectations = append(mmLogin.expectations, expectation)
	return expectation
}

// Then sets up IUserController.Login return parameters for the expectation previously defined by the When method
func (e *IUserControllerMockLoginExpectation) Then(s1 string, err error) *IUserControllerMock {
	e.results = &IUserControllerMockLoginResults{s1, err}
	return e.mock
}

// Login implements IUserController
func (mmLogin *IUserControllerMock) Login(username string, password string) (s1 string, err error) {
	mm_atomic.AddUint64(&mmLogin.beforeLoginCounter, 1)
	defer mm_atomic.AddUint64(&mmLogin.afterLoginCounter, 1)

	if mmLogin.inspectFuncLogin != nil {
		mmLogin.inspectFuncLogin(username, password)
	}

	mm_params := &IUserControllerMockLoginParams{username, password}

	// Record call args
	mmLogin.LoginMock.mutex.Lock()
	mmLogin.LoginMock.callArgs = append(mmLogin.LoginMock.callArgs, mm_params)
	mmLogin.LoginMock.mutex.Unlock()

	for _, e := range mmLogin.LoginMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.s1, e.results.err
		}
	}

	if mmLogin.LoginMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmLogin.LoginMock.defaultExpectation.Counter, 1)
		mm_want := mmLogin.LoginMock.defaultExpectation.params
		mm_got := IUserControllerMockLoginParams{username, password}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmLogin.t.Errorf("IUserControllerMock.Login got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmLogin.LoginMock.defaultExpectation.results
		if mm_results == nil {
			mmLogin.t.Fatal("No results are set for the IUserControllerMock.Login")
		}
		return (*mm_results).s1, (*mm_results).err
	}
	if mmLogin.funcLogin != nil {
		return mmLogin.funcLogin(username, password)
	}
	mmLogin.t.Fatalf("Unexpected call to IUserControllerMock.Login. %v %v", username, password)
	return
}

// LoginAfterCounter returns a count of finished IUserControllerMock.Login invocations
func (mmLogin *IUserControllerMock) LoginAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmLogin.afterLoginCounter)
}

// LoginBeforeCounter returns a count of IUserControllerMock.Login invocations
func (mmLogin *IUserControllerMock) LoginBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmLogin.beforeLoginCounter)
}

// Calls returns a list of arguments used in each call to IUserControllerMock.Login.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmLogin *mIUserControllerMockLogin) Calls() []*IUserControllerMockLoginParams {
	mmLogin.mutex.RLock()

	argCopy := make([]*IUserControllerMockLoginParams, len(mmLogin.callArgs))
	copy(argCopy, mmLogin.callArgs)

	mmLogin.mutex.RUnlock()

	return argCopy
}

// MinimockLoginDone returns true if the count of the Login invocations corresponds
// the number of defined expectations
func (m *IUserControllerMock) MinimockLoginDone() bool {
	for _, e := range m.LoginMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.LoginMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterLoginCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcLogin != nil && mm_atomic.LoadUint64(&m.afterLoginCounter) < 1 {
		return false
	}
	return true
}

// MinimockLoginInspect logs each unmet expectation
func (m *IUserControllerMock) MinimockLoginInspect() {
	for _, e := range m.LoginMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to IUserControllerMock.Login with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.LoginMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterLoginCounter) < 1 {
		if m.LoginMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to IUserControllerMock.Login")
		} else {
			m.t.Errorf("Expected call to IUserControllerMock.Login with params: %#v", *m.LoginMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcLogin != nil && mm_atomic.LoadUint64(&m.afterLoginCounter) < 1 {
		m.t.Error("Expected call to IUserControllerMock.Login")
	}
}

type mIUserControllerMockRegistration struct {
	mock               *IUserControllerMock
	defaultExpectation *IUserControllerMockRegistrationExpectation
	expectations       []*IUserControllerMockRegistrationExpectation

	callArgs []*IUserControllerMockRegistrationParams
	mutex    sync.RWMutex
}

// IUserControllerMockRegistrationExpectation specifies expectation struct of the IUserController.Registration
type IUserControllerMockRegistrationExpectation struct {
	mock    *IUserControllerMock
	params  *IUserControllerMockRegistrationParams
	results *IUserControllerMockRegistrationResults
	Counter uint64
}

// IUserControllerMockRegistrationParams contains parameters of the IUserController.Registration
type IUserControllerMockRegistrationParams struct {
	record *model.User
}

// IUserControllerMockRegistrationResults contains results of the IUserController.Registration
type IUserControllerMockRegistrationResults struct {
	s1  string
	err error
}

// Expect sets up expected params for IUserController.Registration
func (mmRegistration *mIUserControllerMockRegistration) Expect(record *model.User) *mIUserControllerMockRegistration {
	if mmRegistration.mock.funcRegistration != nil {
		mmRegistration.mock.t.Fatalf("IUserControllerMock.Registration mock is already set by Set")
	}

	if mmRegistration.defaultExpectation == nil {
		mmRegistration.defaultExpectation = &IUserControllerMockRegistrationExpectation{}
	}

	mmRegistration.defaultExpectation.params = &IUserControllerMockRegistrationParams{record}
	for _, e := range mmRegistration.expectations {
		if minimock.Equal(e.params, mmRegistration.defaultExpectation.params) {
			mmRegistration.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmRegistration.defaultExpectation.params)
		}
	}

	return mmRegistration
}

// Inspect accepts an inspector function that has same arguments as the IUserController.Registration
func (mmRegistration *mIUserControllerMockRegistration) Inspect(f func(record *model.User)) *mIUserControllerMockRegistration {
	if mmRegistration.mock.inspectFuncRegistration != nil {
		mmRegistration.mock.t.Fatalf("Inspect function is already set for IUserControllerMock.Registration")
	}

	mmRegistration.mock.inspectFuncRegistration = f

	return mmRegistration
}

// Return sets up results that will be returned by IUserController.Registration
func (mmRegistration *mIUserControllerMockRegistration) Return(s1 string, err error) *IUserControllerMock {
	if mmRegistration.mock.funcRegistration != nil {
		mmRegistration.mock.t.Fatalf("IUserControllerMock.Registration mock is already set by Set")
	}

	if mmRegistration.defaultExpectation == nil {
		mmRegistration.defaultExpectation = &IUserControllerMockRegistrationExpectation{mock: mmRegistration.mock}
	}
	mmRegistration.defaultExpectation.results = &IUserControllerMockRegistrationResults{s1, err}
	return mmRegistration.mock
}

// Set uses given function f to mock the IUserController.Registration method
func (mmRegistration *mIUserControllerMockRegistration) Set(f func(record *model.User) (s1 string, err error)) *IUserControllerMock {
	if mmRegistration.defaultExpectation != nil {
		mmRegistration.mock.t.Fatalf("Default expectation is already set for the IUserController.Registration method")
	}

	if len(mmRegistration.expectations) > 0 {
		mmRegistration.mock.t.Fatalf("Some expectations are already set for the IUserController.Registration method")
	}

	mmRegistration.mock.funcRegistration = f
	return mmRegistration.mock
}

// When sets expectation for the IUserController.Registration which will trigger the result defined by the following
// Then helper
func (mmRegistration *mIUserControllerMockRegistration) When(record *model.User) *IUserControllerMockRegistrationExpectation {
	if mmRegistration.mock.funcRegistration != nil {
		mmRegistration.mock.t.Fatalf("IUserControllerMock.Registration mock is already set by Set")
	}

	expectation := &IUserControllerMockRegistrationExpectation{
		mock:   mmRegistration.mock,
		params: &IUserControllerMockRegistrationParams{record},
	}
	mmRegistration.expectations = append(mmRegistration.expectations, expectation)
	return expectation
}

// Then sets up IUserController.Registration return parameters for the expectation previously defined by the When method
func (e *IUserControllerMockRegistrationExpectation) Then(s1 string, err error) *IUserControllerMock {
	e.results = &IUserControllerMockRegistrationResults{s1, err}
	return e.mock
}

// Registration implements IUserController
func (mmRegistration *IUserControllerMock) Registration(record *model.User) (s1 string, err error) {
	mm_atomic.AddUint64(&mmRegistration.beforeRegistrationCounter, 1)
	defer mm_atomic.AddUint64(&mmRegistration.afterRegistrationCounter, 1)

	if mmRegistration.inspectFuncRegistration != nil {
		mmRegistration.inspectFuncRegistration(record)
	}

	mm_params := &IUserControllerMockRegistrationParams{record}

	// Record call args
	mmRegistration.RegistrationMock.mutex.Lock()
	mmRegistration.RegistrationMock.callArgs = append(mmRegistration.RegistrationMock.callArgs, mm_params)
	mmRegistration.RegistrationMock.mutex.Unlock()

	for _, e := range mmRegistration.RegistrationMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.s1, e.results.err
		}
	}

	if mmRegistration.RegistrationMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmRegistration.RegistrationMock.defaultExpectation.Counter, 1)
		mm_want := mmRegistration.RegistrationMock.defaultExpectation.params
		mm_got := IUserControllerMockRegistrationParams{record}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmRegistration.t.Errorf("IUserControllerMock.Registration got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmRegistration.RegistrationMock.defaultExpectation.results
		if mm_results == nil {
			mmRegistration.t.Fatal("No results are set for the IUserControllerMock.Registration")
		}
		return (*mm_results).s1, (*mm_results).err
	}
	if mmRegistration.funcRegistration != nil {
		return mmRegistration.funcRegistration(record)
	}
	mmRegistration.t.Fatalf("Unexpected call to IUserControllerMock.Registration. %v", record)
	return
}

// RegistrationAfterCounter returns a count of finished IUserControllerMock.Registration invocations
func (mmRegistration *IUserControllerMock) RegistrationAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmRegistration.afterRegistrationCounter)
}

// RegistrationBeforeCounter returns a count of IUserControllerMock.Registration invocations
func (mmRegistration *IUserControllerMock) RegistrationBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmRegistration.beforeRegistrationCounter)
}

// Calls returns a list of arguments used in each call to IUserControllerMock.Registration.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmRegistration *mIUserControllerMockRegistration) Calls() []*IUserControllerMockRegistrationParams {
	mmRegistration.mutex.RLock()

	argCopy := make([]*IUserControllerMockRegistrationParams, len(mmRegistration.callArgs))
	copy(argCopy, mmRegistration.callArgs)

	mmRegistration.mutex.RUnlock()

	return argCopy
}

// MinimockRegistrationDone returns true if the count of the Registration invocations corresponds
// the number of defined expectations
func (m *IUserControllerMock) MinimockRegistrationDone() bool {
	for _, e := range m.RegistrationMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.RegistrationMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterRegistrationCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcRegistration != nil && mm_atomic.LoadUint64(&m.afterRegistrationCounter) < 1 {
		return false
	}
	return true
}

// MinimockRegistrationInspect logs each unmet expectation
func (m *IUserControllerMock) MinimockRegistrationInspect() {
	for _, e := range m.RegistrationMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to IUserControllerMock.Registration with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.RegistrationMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterRegistrationCounter) < 1 {
		if m.RegistrationMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to IUserControllerMock.Registration")
		} else {
			m.t.Errorf("Expected call to IUserControllerMock.Registration with params: %#v", *m.RegistrationMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcRegistration != nil && mm_atomic.LoadUint64(&m.afterRegistrationCounter) < 1 {
		m.t.Error("Expected call to IUserControllerMock.Registration")
	}
}

type mIUserControllerMockUpdateActive struct {
	mock               *IUserControllerMock
	defaultExpectation *IUserControllerMockUpdateActiveExpectation
	expectations       []*IUserControllerMockUpdateActiveExpectation

	callArgs []*IUserControllerMockUpdateActiveParams
	mutex    sync.RWMutex
}

// IUserControllerMockUpdateActiveExpectation specifies expectation struct of the IUserController.UpdateActive
type IUserControllerMockUpdateActiveExpectation struct {
	mock    *IUserControllerMock
	params  *IUserControllerMockUpdateActiveParams
	results *IUserControllerMockUpdateActiveResults
	Counter uint64
}

// IUserControllerMockUpdateActiveParams contains parameters of the IUserController.UpdateActive
type IUserControllerMockUpdateActiveParams struct {
	token string
}

// IUserControllerMockUpdateActiveResults contains results of the IUserController.UpdateActive
type IUserControllerMockUpdateActiveResults struct {
	err error
}

// Expect sets up expected params for IUserController.UpdateActive
func (mmUpdateActive *mIUserControllerMockUpdateActive) Expect(token string) *mIUserControllerMockUpdateActive {
	if mmUpdateActive.mock.funcUpdateActive != nil {
		mmUpdateActive.mock.t.Fatalf("IUserControllerMock.UpdateActive mock is already set by Set")
	}

	if mmUpdateActive.defaultExpectation == nil {
		mmUpdateActive.defaultExpectation = &IUserControllerMockUpdateActiveExpectation{}
	}

	mmUpdateActive.defaultExpectation.params = &IUserControllerMockUpdateActiveParams{token}
	for _, e := range mmUpdateActive.expectations {
		if minimock.Equal(e.params, mmUpdateActive.defaultExpectation.params) {
			mmUpdateActive.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmUpdateActive.defaultExpectation.params)
		}
	}

	return mmUpdateActive
}

// Inspect accepts an inspector function that has same arguments as the IUserController.UpdateActive
func (mmUpdateActive *mIUserControllerMockUpdateActive) Inspect(f func(token string)) *mIUserControllerMockUpdateActive {
	if mmUpdateActive.mock.inspectFuncUpdateActive != nil {
		mmUpdateActive.mock.t.Fatalf("Inspect function is already set for IUserControllerMock.UpdateActive")
	}

	mmUpdateActive.mock.inspectFuncUpdateActive = f

	return mmUpdateActive
}

// Return sets up results that will be returned by IUserController.UpdateActive
func (mmUpdateActive *mIUserControllerMockUpdateActive) Return(err error) *IUserControllerMock {
	if mmUpdateActive.mock.funcUpdateActive != nil {
		mmUpdateActive.mock.t.Fatalf("IUserControllerMock.UpdateActive mock is already set by Set")
	}

	if mmUpdateActive.defaultExpectation == nil {
		mmUpdateActive.defaultExpectation = &IUserControllerMockUpdateActiveExpectation{mock: mmUpdateActive.mock}
	}
	mmUpdateActive.defaultExpectation.results = &IUserControllerMockUpdateActiveResults{err}
	return mmUpdateActive.mock
}

// Set uses given function f to mock the IUserController.UpdateActive method
func (mmUpdateActive *mIUserControllerMockUpdateActive) Set(f func(token string) (err error)) *IUserControllerMock {
	if mmUpdateActive.defaultExpectation != nil {
		mmUpdateActive.mock.t.Fatalf("Default expectation is already set for the IUserController.UpdateActive method")
	}

	if len(mmUpdateActive.expectations) > 0 {
		mmUpdateActive.mock.t.Fatalf("Some expectations are already set for the IUserController.UpdateActive method")
	}

	mmUpdateActive.mock.funcUpdateActive = f
	return mmUpdateActive.mock
}

// When sets expectation for the IUserController.UpdateActive which will trigger the result defined by the following
// Then helper
func (mmUpdateActive *mIUserControllerMockUpdateActive) When(token string) *IUserControllerMockUpdateActiveExpectation {
	if mmUpdateActive.mock.funcUpdateActive != nil {
		mmUpdateActive.mock.t.Fatalf("IUserControllerMock.UpdateActive mock is already set by Set")
	}

	expectation := &IUserControllerMockUpdateActiveExpectation{
		mock:   mmUpdateActive.mock,
		params: &IUserControllerMockUpdateActiveParams{token},
	}
	mmUpdateActive.expectations = append(mmUpdateActive.expectations, expectation)
	return expectation
}

// Then sets up IUserController.UpdateActive return parameters for the expectation previously defined by the When method
func (e *IUserControllerMockUpdateActiveExpectation) Then(err error) *IUserControllerMock {
	e.results = &IUserControllerMockUpdateActiveResults{err}
	return e.mock
}

// UpdateActive implements IUserController
func (mmUpdateActive *IUserControllerMock) UpdateActive(token string) (err error) {
	mm_atomic.AddUint64(&mmUpdateActive.beforeUpdateActiveCounter, 1)
	defer mm_atomic.AddUint64(&mmUpdateActive.afterUpdateActiveCounter, 1)

	if mmUpdateActive.inspectFuncUpdateActive != nil {
		mmUpdateActive.inspectFuncUpdateActive(token)
	}

	mm_params := &IUserControllerMockUpdateActiveParams{token}

	// Record call args
	mmUpdateActive.UpdateActiveMock.mutex.Lock()
	mmUpdateActive.UpdateActiveMock.callArgs = append(mmUpdateActive.UpdateActiveMock.callArgs, mm_params)
	mmUpdateActive.UpdateActiveMock.mutex.Unlock()

	for _, e := range mmUpdateActive.UpdateActiveMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmUpdateActive.UpdateActiveMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmUpdateActive.UpdateActiveMock.defaultExpectation.Counter, 1)
		mm_want := mmUpdateActive.UpdateActiveMock.defaultExpectation.params
		mm_got := IUserControllerMockUpdateActiveParams{token}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmUpdateActive.t.Errorf("IUserControllerMock.UpdateActive got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmUpdateActive.UpdateActiveMock.defaultExpectation.results
		if mm_results == nil {
			mmUpdateActive.t.Fatal("No results are set for the IUserControllerMock.UpdateActive")
		}
		return (*mm_results).err
	}
	if mmUpdateActive.funcUpdateActive != nil {
		return mmUpdateActive.funcUpdateActive(token)
	}
	mmUpdateActive.t.Fatalf("Unexpected call to IUserControllerMock.UpdateActive. %v", token)
	return
}

// UpdateActiveAfterCounter returns a count of finished IUserControllerMock.UpdateActive invocations
func (mmUpdateActive *IUserControllerMock) UpdateActiveAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmUpdateActive.afterUpdateActiveCounter)
}

// UpdateActiveBeforeCounter returns a count of IUserControllerMock.UpdateActive invocations
func (mmUpdateActive *IUserControllerMock) UpdateActiveBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmUpdateActive.beforeUpdateActiveCounter)
}

// Calls returns a list of arguments used in each call to IUserControllerMock.UpdateActive.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmUpdateActive *mIUserControllerMockUpdateActive) Calls() []*IUserControllerMockUpdateActiveParams {
	mmUpdateActive.mutex.RLock()

	argCopy := make([]*IUserControllerMockUpdateActiveParams, len(mmUpdateActive.callArgs))
	copy(argCopy, mmUpdateActive.callArgs)

	mmUpdateActive.mutex.RUnlock()

	return argCopy
}

// MinimockUpdateActiveDone returns true if the count of the UpdateActive invocations corresponds
// the number of defined expectations
func (m *IUserControllerMock) MinimockUpdateActiveDone() bool {
	for _, e := range m.UpdateActiveMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.UpdateActiveMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterUpdateActiveCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcUpdateActive != nil && mm_atomic.LoadUint64(&m.afterUpdateActiveCounter) < 1 {
		return false
	}
	return true
}

// MinimockUpdateActiveInspect logs each unmet expectation
func (m *IUserControllerMock) MinimockUpdateActiveInspect() {
	for _, e := range m.UpdateActiveMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to IUserControllerMock.UpdateActive with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.UpdateActiveMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterUpdateActiveCounter) < 1 {
		if m.UpdateActiveMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to IUserControllerMock.UpdateActive")
		} else {
			m.t.Errorf("Expected call to IUserControllerMock.UpdateActive with params: %#v", *m.UpdateActiveMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcUpdateActive != nil && mm_atomic.LoadUint64(&m.afterUpdateActiveCounter) < 1 {
		m.t.Error("Expected call to IUserControllerMock.UpdateActive")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *IUserControllerMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockAuthByTokenInspect()

		m.MinimockLoginInspect()

		m.MinimockRegistrationInspect()

		m.MinimockUpdateActiveInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *IUserControllerMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *IUserControllerMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockAuthByTokenDone() &&
		m.MinimockLoginDone() &&
		m.MinimockRegistrationDone() &&
		m.MinimockUpdateActiveDone()
}
