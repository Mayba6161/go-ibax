/*---------------------------------------------------------------------------------------------
 *  Copyright (c) IBAX. All rights reserved.
 *  See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/
package utils

import (
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// MockClock is an autogenerated mock type for the Clock type
type MockClock struct {
	mock.Mock
}

// Now provides a mock function with given fields:
func (_m *MockClock) Now() time.Time {
	ret := _m.Called()

	var r0 time.Time
