// Code generated by mockery v2.13.0-beta.1. DO NOT EDIT.

package mocks

import (
	big "math/big"

	mock "github.com/stretchr/testify/mock"

	pg "github.com/smartcontractkit/chainlink/core/services/pg"
)

// ContractSubmitter is an autogenerated mock type for the ContractSubmitter type
type ContractSubmitter struct {
	mock.Mock
}

// Submit provides a mock function with given fields: roundID, submission, qopts
func (_m *ContractSubmitter) Submit(roundID *big.Int, submission *big.Int, qopts ...pg.QOpt) error {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, roundID, submission)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(*big.Int, *big.Int, ...pg.QOpt) error); ok {
		r0 = rf(roundID, submission, qopts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type NewContractSubmitterT interface {
	mock.TestingT
	Cleanup(func())
}

// NewContractSubmitter creates a new instance of ContractSubmitter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewContractSubmitter(t NewContractSubmitterT) *ContractSubmitter {
	mock := &ContractSubmitter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
