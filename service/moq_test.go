// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package service

import (
	"context"
	"sync"
	"todo_app/entity"
	"todo_app/store"
)

// Ensure, that TaskAdderMock does implement TaskAdder.
// If this is not the case, regenerate this file with moq.
var _ TaskAdder = &TaskAdderMock{}

// TaskAdderMock is a mock implementation of TaskAdder.
//
//	func TestSomethingThatUsesTaskAdder(t *testing.T) {
//
//		// make and configure a mocked TaskAdder
//		mockedTaskAdder := &TaskAdderMock{
//			AddTaskFunc: func(ctx context.Context, db store.Execer, t *entity.Task) error {
//				panic("mock out the AddTask method")
//			},
//		}
//
//		// use mockedTaskAdder in code that requires TaskAdder
//		// and then make assertions.
//
//	}
type TaskAdderMock struct {
	// AddTaskFunc mocks the AddTask method.
	AddTaskFunc func(ctx context.Context, db store.Execer, t *entity.Task) error

	// calls tracks calls to the methods.
	calls struct {
		// AddTask holds details about calls to the AddTask method.
		AddTask []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Db is the db argument value.
			Db store.Execer
			// T is the t argument value.
			T *entity.Task
		}
	}
	lockAddTask sync.RWMutex
}

// AddTask calls AddTaskFunc.
func (mock *TaskAdderMock) AddTask(ctx context.Context, db store.Execer, t *entity.Task) error {
	if mock.AddTaskFunc == nil {
		panic("TaskAdderMock.AddTaskFunc: method is nil but TaskAdder.AddTask was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Db  store.Execer
		T   *entity.Task
	}{
		Ctx: ctx,
		Db:  db,
		T:   t,
	}
	mock.lockAddTask.Lock()
	mock.calls.AddTask = append(mock.calls.AddTask, callInfo)
	mock.lockAddTask.Unlock()
	return mock.AddTaskFunc(ctx, db, t)
}

// AddTaskCalls gets all the calls that were made to AddTask.
// Check the length with:
//
//	len(mockedTaskAdder.AddTaskCalls())
func (mock *TaskAdderMock) AddTaskCalls() []struct {
	Ctx context.Context
	Db  store.Execer
	T   *entity.Task
} {
	var calls []struct {
		Ctx context.Context
		Db  store.Execer
		T   *entity.Task
	}
	mock.lockAddTask.RLock()
	calls = mock.calls.AddTask
	mock.lockAddTask.RUnlock()
	return calls
}

// Ensure, that TaskListerMock does implement TaskLister.
// If this is not the case, regenerate this file with moq.
var _ TaskLister = &TaskListerMock{}

// TaskListerMock is a mock implementation of TaskLister.
//
//	func TestSomethingThatUsesTaskLister(t *testing.T) {
//
//		// make and configure a mocked TaskLister
//		mockedTaskLister := &TaskListerMock{
//			ListTasksFunc: func(ctx context.Context, db store.Queryer) (entity.Tasks, error) {
//				panic("mock out the ListTasks method")
//			},
//		}
//
//		// use mockedTaskLister in code that requires TaskLister
//		// and then make assertions.
//
//	}
type TaskListerMock struct {
	// ListTasksFunc mocks the ListTasks method.
	ListTasksFunc func(ctx context.Context, db store.Queryer) (entity.Tasks, error)

	// calls tracks calls to the methods.
	calls struct {
		// ListTasks holds details about calls to the ListTasks method.
		ListTasks []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Db is the db argument value.
			Db store.Queryer
		}
	}
	lockListTasks sync.RWMutex
}

// ListTasks calls ListTasksFunc.
func (mock *TaskListerMock) ListTasks(ctx context.Context, db store.Queryer) (entity.Tasks, error) {
	if mock.ListTasksFunc == nil {
		panic("TaskListerMock.ListTasksFunc: method is nil but TaskLister.ListTasks was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Db  store.Queryer
	}{
		Ctx: ctx,
		Db:  db,
	}
	mock.lockListTasks.Lock()
	mock.calls.ListTasks = append(mock.calls.ListTasks, callInfo)
	mock.lockListTasks.Unlock()
	return mock.ListTasksFunc(ctx, db)
}

// ListTasksCalls gets all the calls that were made to ListTasks.
// Check the length with:
//
//	len(mockedTaskLister.ListTasksCalls())
func (mock *TaskListerMock) ListTasksCalls() []struct {
	Ctx context.Context
	Db  store.Queryer
} {
	var calls []struct {
		Ctx context.Context
		Db  store.Queryer
	}
	mock.lockListTasks.RLock()
	calls = mock.calls.ListTasks
	mock.lockListTasks.RUnlock()
	return calls
}

// Ensure, that UserRegisterMock does implement UserRegister.
// If this is not the case, regenerate this file with moq.
var _ UserRegister = &UserRegisterMock{}

// UserRegisterMock is a mock implementation of UserRegister.
//
//	func TestSomethingThatUsesUserRegister(t *testing.T) {
//
//		// make and configure a mocked UserRegister
//		mockedUserRegister := &UserRegisterMock{
//			RegisterUserFunc: func(ctx context.Context, db store.Execer, u *entity.User) error {
//				panic("mock out the RegisterUser method")
//			},
//		}
//
//		// use mockedUserRegister in code that requires UserRegister
//		// and then make assertions.
//
//	}
type UserRegisterMock struct {
	// RegisterUserFunc mocks the RegisterUser method.
	RegisterUserFunc func(ctx context.Context, db store.Execer, u *entity.User) error

	// calls tracks calls to the methods.
	calls struct {
		// RegisterUser holds details about calls to the RegisterUser method.
		RegisterUser []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Db is the db argument value.
			Db store.Execer
			// U is the u argument value.
			U *entity.User
		}
	}
	lockRegisterUser sync.RWMutex
}

// RegisterUser calls RegisterUserFunc.
func (mock *UserRegisterMock) RegisterUser(ctx context.Context, db store.Execer, u *entity.User) error {
	if mock.RegisterUserFunc == nil {
		panic("UserRegisterMock.RegisterUserFunc: method is nil but UserRegister.RegisterUser was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Db  store.Execer
		U   *entity.User
	}{
		Ctx: ctx,
		Db:  db,
		U:   u,
	}
	mock.lockRegisterUser.Lock()
	mock.calls.RegisterUser = append(mock.calls.RegisterUser, callInfo)
	mock.lockRegisterUser.Unlock()
	return mock.RegisterUserFunc(ctx, db, u)
}

// RegisterUserCalls gets all the calls that were made to RegisterUser.
// Check the length with:
//
//	len(mockedUserRegister.RegisterUserCalls())
func (mock *UserRegisterMock) RegisterUserCalls() []struct {
	Ctx context.Context
	Db  store.Execer
	U   *entity.User
} {
	var calls []struct {
		Ctx context.Context
		Db  store.Execer
		U   *entity.User
	}
	mock.lockRegisterUser.RLock()
	calls = mock.calls.RegisterUser
	mock.lockRegisterUser.RUnlock()
	return calls
}

// Ensure, that UserGetterMock does implement UserGetter.
// If this is not the case, regenerate this file with moq.
var _ UserGetter = &UserGetterMock{}

// UserGetterMock is a mock implementation of UserGetter.
//
//	func TestSomethingThatUsesUserGetter(t *testing.T) {
//
//		// make and configure a mocked UserGetter
//		mockedUserGetter := &UserGetterMock{
//			GetUserFunc: func(ctx context.Context, db store.Queryer, name string) (*entity.User, error) {
//				panic("mock out the GetUser method")
//			},
//		}
//
//		// use mockedUserGetter in code that requires UserGetter
//		// and then make assertions.
//
//	}
type UserGetterMock struct {
	// GetUserFunc mocks the GetUser method.
	GetUserFunc func(ctx context.Context, db store.Queryer, name string) (*entity.User, error)

	// calls tracks calls to the methods.
	calls struct {
		// GetUser holds details about calls to the GetUser method.
		GetUser []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Db is the db argument value.
			Db store.Queryer
			// Name is the name argument value.
			Name string
		}
	}
	lockGetUser sync.RWMutex
}

// GetUser calls GetUserFunc.
func (mock *UserGetterMock) GetUser(ctx context.Context, db store.Queryer, name string) (*entity.User, error) {
	if mock.GetUserFunc == nil {
		panic("UserGetterMock.GetUserFunc: method is nil but UserGetter.GetUser was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Db   store.Queryer
		Name string
	}{
		Ctx:  ctx,
		Db:   db,
		Name: name,
	}
	mock.lockGetUser.Lock()
	mock.calls.GetUser = append(mock.calls.GetUser, callInfo)
	mock.lockGetUser.Unlock()
	return mock.GetUserFunc(ctx, db, name)
}

// GetUserCalls gets all the calls that were made to GetUser.
// Check the length with:
//
//	len(mockedUserGetter.GetUserCalls())
func (mock *UserGetterMock) GetUserCalls() []struct {
	Ctx  context.Context
	Db   store.Queryer
	Name string
} {
	var calls []struct {
		Ctx  context.Context
		Db   store.Queryer
		Name string
	}
	mock.lockGetUser.RLock()
	calls = mock.calls.GetUser
	mock.lockGetUser.RUnlock()
	return calls
}

// Ensure, that TokenGeneratorMock does implement TokenGenerator.
// If this is not the case, regenerate this file with moq.
var _ TokenGenerator = &TokenGeneratorMock{}

// TokenGeneratorMock is a mock implementation of TokenGenerator.
//
//	func TestSomethingThatUsesTokenGenerator(t *testing.T) {
//
//		// make and configure a mocked TokenGenerator
//		mockedTokenGenerator := &TokenGeneratorMock{
//			GenerateTokenFunc: func(ctx context.Context, u entity.User) ([]byte, error) {
//				panic("mock out the GenerateToken method")
//			},
//		}
//
//		// use mockedTokenGenerator in code that requires TokenGenerator
//		// and then make assertions.
//
//	}
type TokenGeneratorMock struct {
	// GenerateTokenFunc mocks the GenerateToken method.
	GenerateTokenFunc func(ctx context.Context, u entity.User) ([]byte, error)

	// calls tracks calls to the methods.
	calls struct {
		// GenerateToken holds details about calls to the GenerateToken method.
		GenerateToken []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// U is the u argument value.
			U entity.User
		}
	}
	lockGenerateToken sync.RWMutex
}

// GenerateToken calls GenerateTokenFunc.
func (mock *TokenGeneratorMock) GenerateToken(ctx context.Context, u entity.User) ([]byte, error) {
	if mock.GenerateTokenFunc == nil {
		panic("TokenGeneratorMock.GenerateTokenFunc: method is nil but TokenGenerator.GenerateToken was just called")
	}
	callInfo := struct {
		Ctx context.Context
		U   entity.User
	}{
		Ctx: ctx,
		U:   u,
	}
	mock.lockGenerateToken.Lock()
	mock.calls.GenerateToken = append(mock.calls.GenerateToken, callInfo)
	mock.lockGenerateToken.Unlock()
	return mock.GenerateTokenFunc(ctx, u)
}

// GenerateTokenCalls gets all the calls that were made to GenerateToken.
// Check the length with:
//
//	len(mockedTokenGenerator.GenerateTokenCalls())
func (mock *TokenGeneratorMock) GenerateTokenCalls() []struct {
	Ctx context.Context
	U   entity.User
} {
	var calls []struct {
		Ctx context.Context
		U   entity.User
	}
	mock.lockGenerateToken.RLock()
	calls = mock.calls.GenerateToken
	mock.lockGenerateToken.RUnlock()
	return calls
}