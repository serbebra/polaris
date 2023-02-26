// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"github.com/ethereum/go-ethereum/common"
	ethereumcorevm "github.com/ethereum/go-ethereum/core/vm"
	"math/big"
	ethcorevm "pkg.berachain.dev/stargazer/eth/core/vm"
	"sync"
)

// Ensure, that EVMLoggerMock does implement ethcorevm.EVMLogger.
// If this is not the case, regenerate this file with moq.
var _ ethcorevm.EVMLogger = &EVMLoggerMock{}

// EVMLoggerMock is a mock implementation of ethcorevm.EVMLogger.
//
//	func TestSomethingThatUsesEVMLogger(t *testing.T) {
//
//		// make and configure a mocked ethcorevm.EVMLogger
//		mockedEVMLogger := &EVMLoggerMock{
//			CaptureEndFunc: func(output []byte, gasUsed uint64, err error)  {
//				panic("mock out the CaptureEnd method")
//			},
//			CaptureEnterFunc: func(typ ethereumcorevm.OpCode, from common.Address, to common.Address, input []byte, gas uint64, value *big.Int)  {
//				panic("mock out the CaptureEnter method")
//			},
//			CaptureExitFunc: func(output []byte, gasUsed uint64, err error)  {
//				panic("mock out the CaptureExit method")
//			},
//			CaptureFaultFunc: func(pc uint64, op ethereumcorevm.OpCode, gas uint64, cost uint64, scope *ethereumcorevm.ScopeContext, depth int, err error)  {
//				panic("mock out the CaptureFault method")
//			},
//			CaptureStartFunc: func(env *ethereumcorevm.EVM, from common.Address, to common.Address, create bool, input []byte, gas uint64, value *big.Int)  {
//				panic("mock out the CaptureStart method")
//			},
//			CaptureStateFunc: func(pc uint64, op ethereumcorevm.OpCode, gas uint64, cost uint64, scope *ethereumcorevm.ScopeContext, rData []byte, depth int, err error)  {
//				panic("mock out the CaptureState method")
//			},
//			CaptureTxEndFunc: func(restGas uint64)  {
//				panic("mock out the CaptureTxEnd method")
//			},
//			CaptureTxStartFunc: func(gasLimit uint64)  {
//				panic("mock out the CaptureTxStart method")
//			},
//		}
//
//		// use mockedEVMLogger in code that requires ethcorevm.EVMLogger
//		// and then make assertions.
//
//	}
type EVMLoggerMock struct {
	// CaptureEndFunc mocks the CaptureEnd method.
	CaptureEndFunc func(output []byte, gasUsed uint64, err error)

	// CaptureEnterFunc mocks the CaptureEnter method.
	CaptureEnterFunc func(typ ethereumcorevm.OpCode, from common.Address, to common.Address, input []byte, gas uint64, value *big.Int)

	// CaptureExitFunc mocks the CaptureExit method.
	CaptureExitFunc func(output []byte, gasUsed uint64, err error)

	// CaptureFaultFunc mocks the CaptureFault method.
	CaptureFaultFunc func(pc uint64, op ethereumcorevm.OpCode, gas uint64, cost uint64, scope *ethereumcorevm.ScopeContext, depth int, err error)

	// CaptureStartFunc mocks the CaptureStart method.
	CaptureStartFunc func(env *ethereumcorevm.EVM, from common.Address, to common.Address, create bool, input []byte, gas uint64, value *big.Int)

	// CaptureStateFunc mocks the CaptureState method.
	CaptureStateFunc func(pc uint64, op ethereumcorevm.OpCode, gas uint64, cost uint64, scope *ethereumcorevm.ScopeContext, rData []byte, depth int, err error)

	// CaptureTxEndFunc mocks the CaptureTxEnd method.
	CaptureTxEndFunc func(restGas uint64)

	// CaptureTxStartFunc mocks the CaptureTxStart method.
	CaptureTxStartFunc func(gasLimit uint64)

	// calls tracks calls to the methods.
	calls struct {
		// CaptureEnd holds details about calls to the CaptureEnd method.
		CaptureEnd []struct {
			// Output is the output argument value.
			Output []byte
			// GasUsed is the gasUsed argument value.
			GasUsed uint64
			// Err is the err argument value.
			Err error
		}
		// CaptureEnter holds details about calls to the CaptureEnter method.
		CaptureEnter []struct {
			// Typ is the typ argument value.
			Typ ethereumcorevm.OpCode
			// From is the from argument value.
			From common.Address
			// To is the to argument value.
			To common.Address
			// Input is the input argument value.
			Input []byte
			// Gas is the gas argument value.
			Gas uint64
			// Value is the value argument value.
			Value *big.Int
		}
		// CaptureExit holds details about calls to the CaptureExit method.
		CaptureExit []struct {
			// Output is the output argument value.
			Output []byte
			// GasUsed is the gasUsed argument value.
			GasUsed uint64
			// Err is the err argument value.
			Err error
		}
		// CaptureFault holds details about calls to the CaptureFault method.
		CaptureFault []struct {
			// Pc is the pc argument value.
			Pc uint64
			// Op is the op argument value.
			Op ethereumcorevm.OpCode
			// Gas is the gas argument value.
			Gas uint64
			// Cost is the cost argument value.
			Cost uint64
			// Scope is the scope argument value.
			Scope *ethereumcorevm.ScopeContext
			// Depth is the depth argument value.
			Depth int
			// Err is the err argument value.
			Err error
		}
		// CaptureStart holds details about calls to the CaptureStart method.
		CaptureStart []struct {
			// Env is the env argument value.
			Env *ethereumcorevm.EVM
			// From is the from argument value.
			From common.Address
			// To is the to argument value.
			To common.Address
			// Create is the create argument value.
			Create bool
			// Input is the input argument value.
			Input []byte
			// Gas is the gas argument value.
			Gas uint64
			// Value is the value argument value.
			Value *big.Int
		}
		// CaptureState holds details about calls to the CaptureState method.
		CaptureState []struct {
			// Pc is the pc argument value.
			Pc uint64
			// Op is the op argument value.
			Op ethereumcorevm.OpCode
			// Gas is the gas argument value.
			Gas uint64
			// Cost is the cost argument value.
			Cost uint64
			// Scope is the scope argument value.
			Scope *ethereumcorevm.ScopeContext
			// RData is the rData argument value.
			RData []byte
			// Depth is the depth argument value.
			Depth int
			// Err is the err argument value.
			Err error
		}
		// CaptureTxEnd holds details about calls to the CaptureTxEnd method.
		CaptureTxEnd []struct {
			// RestGas is the restGas argument value.
			RestGas uint64
		}
		// CaptureTxStart holds details about calls to the CaptureTxStart method.
		CaptureTxStart []struct {
			// GasLimit is the gasLimit argument value.
			GasLimit uint64
		}
	}
	lockCaptureEnd     sync.RWMutex
	lockCaptureEnter   sync.RWMutex
	lockCaptureExit    sync.RWMutex
	lockCaptureFault   sync.RWMutex
	lockCaptureStart   sync.RWMutex
	lockCaptureState   sync.RWMutex
	lockCaptureTxEnd   sync.RWMutex
	lockCaptureTxStart sync.RWMutex
}

// CaptureEnd calls CaptureEndFunc.
func (mock *EVMLoggerMock) CaptureEnd(output []byte, gasUsed uint64, err error) {
	if mock.CaptureEndFunc == nil {
		panic("EVMLoggerMock.CaptureEndFunc: method is nil but EVMLogger.CaptureEnd was just called")
	}
	callInfo := struct {
		Output  []byte
		GasUsed uint64
		Err     error
	}{
		Output:  output,
		GasUsed: gasUsed,
		Err:     err,
	}
	mock.lockCaptureEnd.Lock()
	mock.calls.CaptureEnd = append(mock.calls.CaptureEnd, callInfo)
	mock.lockCaptureEnd.Unlock()
	mock.CaptureEndFunc(output, gasUsed, err)
}

// CaptureEndCalls gets all the calls that were made to CaptureEnd.
// Check the length with:
//
//	len(mockedEVMLogger.CaptureEndCalls())
func (mock *EVMLoggerMock) CaptureEndCalls() []struct {
	Output  []byte
	GasUsed uint64
	Err     error
} {
	var calls []struct {
		Output  []byte
		GasUsed uint64
		Err     error
	}
	mock.lockCaptureEnd.RLock()
	calls = mock.calls.CaptureEnd
	mock.lockCaptureEnd.RUnlock()
	return calls
}

// CaptureEnter calls CaptureEnterFunc.
func (mock *EVMLoggerMock) CaptureEnter(typ ethereumcorevm.OpCode, from common.Address, to common.Address, input []byte, gas uint64, value *big.Int) {
	if mock.CaptureEnterFunc == nil {
		panic("EVMLoggerMock.CaptureEnterFunc: method is nil but EVMLogger.CaptureEnter was just called")
	}
	callInfo := struct {
		Typ   ethereumcorevm.OpCode
		From  common.Address
		To    common.Address
		Input []byte
		Gas   uint64
		Value *big.Int
	}{
		Typ:   typ,
		From:  from,
		To:    to,
		Input: input,
		Gas:   gas,
		Value: value,
	}
	mock.lockCaptureEnter.Lock()
	mock.calls.CaptureEnter = append(mock.calls.CaptureEnter, callInfo)
	mock.lockCaptureEnter.Unlock()
	mock.CaptureEnterFunc(typ, from, to, input, gas, value)
}

// CaptureEnterCalls gets all the calls that were made to CaptureEnter.
// Check the length with:
//
//	len(mockedEVMLogger.CaptureEnterCalls())
func (mock *EVMLoggerMock) CaptureEnterCalls() []struct {
	Typ   ethereumcorevm.OpCode
	From  common.Address
	To    common.Address
	Input []byte
	Gas   uint64
	Value *big.Int
} {
	var calls []struct {
		Typ   ethereumcorevm.OpCode
		From  common.Address
		To    common.Address
		Input []byte
		Gas   uint64
		Value *big.Int
	}
	mock.lockCaptureEnter.RLock()
	calls = mock.calls.CaptureEnter
	mock.lockCaptureEnter.RUnlock()
	return calls
}

// CaptureExit calls CaptureExitFunc.
func (mock *EVMLoggerMock) CaptureExit(output []byte, gasUsed uint64, err error) {
	if mock.CaptureExitFunc == nil {
		panic("EVMLoggerMock.CaptureExitFunc: method is nil but EVMLogger.CaptureExit was just called")
	}
	callInfo := struct {
		Output  []byte
		GasUsed uint64
		Err     error
	}{
		Output:  output,
		GasUsed: gasUsed,
		Err:     err,
	}
	mock.lockCaptureExit.Lock()
	mock.calls.CaptureExit = append(mock.calls.CaptureExit, callInfo)
	mock.lockCaptureExit.Unlock()
	mock.CaptureExitFunc(output, gasUsed, err)
}

// CaptureExitCalls gets all the calls that were made to CaptureExit.
// Check the length with:
//
//	len(mockedEVMLogger.CaptureExitCalls())
func (mock *EVMLoggerMock) CaptureExitCalls() []struct {
	Output  []byte
	GasUsed uint64
	Err     error
} {
	var calls []struct {
		Output  []byte
		GasUsed uint64
		Err     error
	}
	mock.lockCaptureExit.RLock()
	calls = mock.calls.CaptureExit
	mock.lockCaptureExit.RUnlock()
	return calls
}

// CaptureFault calls CaptureFaultFunc.
func (mock *EVMLoggerMock) CaptureFault(pc uint64, op ethereumcorevm.OpCode, gas uint64, cost uint64, scope *ethereumcorevm.ScopeContext, depth int, err error) {
	if mock.CaptureFaultFunc == nil {
		panic("EVMLoggerMock.CaptureFaultFunc: method is nil but EVMLogger.CaptureFault was just called")
	}
	callInfo := struct {
		Pc    uint64
		Op    ethereumcorevm.OpCode
		Gas   uint64
		Cost  uint64
		Scope *ethereumcorevm.ScopeContext
		Depth int
		Err   error
	}{
		Pc:    pc,
		Op:    op,
		Gas:   gas,
		Cost:  cost,
		Scope: scope,
		Depth: depth,
		Err:   err,
	}
	mock.lockCaptureFault.Lock()
	mock.calls.CaptureFault = append(mock.calls.CaptureFault, callInfo)
	mock.lockCaptureFault.Unlock()
	mock.CaptureFaultFunc(pc, op, gas, cost, scope, depth, err)
}

// CaptureFaultCalls gets all the calls that were made to CaptureFault.
// Check the length with:
//
//	len(mockedEVMLogger.CaptureFaultCalls())
func (mock *EVMLoggerMock) CaptureFaultCalls() []struct {
	Pc    uint64
	Op    ethereumcorevm.OpCode
	Gas   uint64
	Cost  uint64
	Scope *ethereumcorevm.ScopeContext
	Depth int
	Err   error
} {
	var calls []struct {
		Pc    uint64
		Op    ethereumcorevm.OpCode
		Gas   uint64
		Cost  uint64
		Scope *ethereumcorevm.ScopeContext
		Depth int
		Err   error
	}
	mock.lockCaptureFault.RLock()
	calls = mock.calls.CaptureFault
	mock.lockCaptureFault.RUnlock()
	return calls
}

// CaptureStart calls CaptureStartFunc.
func (mock *EVMLoggerMock) CaptureStart(env *ethereumcorevm.EVM, from common.Address, to common.Address, create bool, input []byte, gas uint64, value *big.Int) {
	if mock.CaptureStartFunc == nil {
		panic("EVMLoggerMock.CaptureStartFunc: method is nil but EVMLogger.CaptureStart was just called")
	}
	callInfo := struct {
		Env    *ethereumcorevm.EVM
		From   common.Address
		To     common.Address
		Create bool
		Input  []byte
		Gas    uint64
		Value  *big.Int
	}{
		Env:    env,
		From:   from,
		To:     to,
		Create: create,
		Input:  input,
		Gas:    gas,
		Value:  value,
	}
	mock.lockCaptureStart.Lock()
	mock.calls.CaptureStart = append(mock.calls.CaptureStart, callInfo)
	mock.lockCaptureStart.Unlock()
	mock.CaptureStartFunc(env, from, to, create, input, gas, value)
}

// CaptureStartCalls gets all the calls that were made to CaptureStart.
// Check the length with:
//
//	len(mockedEVMLogger.CaptureStartCalls())
func (mock *EVMLoggerMock) CaptureStartCalls() []struct {
	Env    *ethereumcorevm.EVM
	From   common.Address
	To     common.Address
	Create bool
	Input  []byte
	Gas    uint64
	Value  *big.Int
} {
	var calls []struct {
		Env    *ethereumcorevm.EVM
		From   common.Address
		To     common.Address
		Create bool
		Input  []byte
		Gas    uint64
		Value  *big.Int
	}
	mock.lockCaptureStart.RLock()
	calls = mock.calls.CaptureStart
	mock.lockCaptureStart.RUnlock()
	return calls
}

// CaptureState calls CaptureStateFunc.
func (mock *EVMLoggerMock) CaptureState(pc uint64, op ethereumcorevm.OpCode, gas uint64, cost uint64, scope *ethereumcorevm.ScopeContext, rData []byte, depth int, err error) {
	if mock.CaptureStateFunc == nil {
		panic("EVMLoggerMock.CaptureStateFunc: method is nil but EVMLogger.CaptureState was just called")
	}
	callInfo := struct {
		Pc    uint64
		Op    ethereumcorevm.OpCode
		Gas   uint64
		Cost  uint64
		Scope *ethereumcorevm.ScopeContext
		RData []byte
		Depth int
		Err   error
	}{
		Pc:    pc,
		Op:    op,
		Gas:   gas,
		Cost:  cost,
		Scope: scope,
		RData: rData,
		Depth: depth,
		Err:   err,
	}
	mock.lockCaptureState.Lock()
	mock.calls.CaptureState = append(mock.calls.CaptureState, callInfo)
	mock.lockCaptureState.Unlock()
	mock.CaptureStateFunc(pc, op, gas, cost, scope, rData, depth, err)
}

// CaptureStateCalls gets all the calls that were made to CaptureState.
// Check the length with:
//
//	len(mockedEVMLogger.CaptureStateCalls())
func (mock *EVMLoggerMock) CaptureStateCalls() []struct {
	Pc    uint64
	Op    ethereumcorevm.OpCode
	Gas   uint64
	Cost  uint64
	Scope *ethereumcorevm.ScopeContext
	RData []byte
	Depth int
	Err   error
} {
	var calls []struct {
		Pc    uint64
		Op    ethereumcorevm.OpCode
		Gas   uint64
		Cost  uint64
		Scope *ethereumcorevm.ScopeContext
		RData []byte
		Depth int
		Err   error
	}
	mock.lockCaptureState.RLock()
	calls = mock.calls.CaptureState
	mock.lockCaptureState.RUnlock()
	return calls
}

// CaptureTxEnd calls CaptureTxEndFunc.
func (mock *EVMLoggerMock) CaptureTxEnd(restGas uint64) {
	if mock.CaptureTxEndFunc == nil {
		panic("EVMLoggerMock.CaptureTxEndFunc: method is nil but EVMLogger.CaptureTxEnd was just called")
	}
	callInfo := struct {
		RestGas uint64
	}{
		RestGas: restGas,
	}
	mock.lockCaptureTxEnd.Lock()
	mock.calls.CaptureTxEnd = append(mock.calls.CaptureTxEnd, callInfo)
	mock.lockCaptureTxEnd.Unlock()
	mock.CaptureTxEndFunc(restGas)
}

// CaptureTxEndCalls gets all the calls that were made to CaptureTxEnd.
// Check the length with:
//
//	len(mockedEVMLogger.CaptureTxEndCalls())
func (mock *EVMLoggerMock) CaptureTxEndCalls() []struct {
	RestGas uint64
} {
	var calls []struct {
		RestGas uint64
	}
	mock.lockCaptureTxEnd.RLock()
	calls = mock.calls.CaptureTxEnd
	mock.lockCaptureTxEnd.RUnlock()
	return calls
}

// CaptureTxStart calls CaptureTxStartFunc.
func (mock *EVMLoggerMock) CaptureTxStart(gasLimit uint64) {
	if mock.CaptureTxStartFunc == nil {
		panic("EVMLoggerMock.CaptureTxStartFunc: method is nil but EVMLogger.CaptureTxStart was just called")
	}
	callInfo := struct {
		GasLimit uint64
	}{
		GasLimit: gasLimit,
	}
	mock.lockCaptureTxStart.Lock()
	mock.calls.CaptureTxStart = append(mock.calls.CaptureTxStart, callInfo)
	mock.lockCaptureTxStart.Unlock()
	mock.CaptureTxStartFunc(gasLimit)
}

// CaptureTxStartCalls gets all the calls that were made to CaptureTxStart.
// Check the length with:
//
//	len(mockedEVMLogger.CaptureTxStartCalls())
func (mock *EVMLoggerMock) CaptureTxStartCalls() []struct {
	GasLimit uint64
} {
	var calls []struct {
		GasLimit uint64
	}
	mock.lockCaptureTxStart.RLock()
	calls = mock.calls.CaptureTxStart
	mock.lockCaptureTxStart.RUnlock()
	return calls
}