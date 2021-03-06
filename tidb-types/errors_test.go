// Licensed to Apache Software Foundation (ASF) under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Apache Software Foundation (ASF) licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//
// Copyright 2019 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package types

import (
	"github.com/dubbogo/parser/mysql"
	"github.com/dubbogo/parser/terror"
	. "github.com/pingcap/check"
)

type testErrorSuite struct{}

var _ = Suite(testErrorSuite{})

func (s testErrorSuite) TestError(c *C) {
	kvErrs := []*terror.Error{
		ErrInvalidDefault,
		ErrDataTooLong,
		ErrIllegalValueForType,
		ErrTruncated,
		ErrOverflow,
		ErrDivByZero,
		ErrTooBigDisplayWidth,
		ErrTooBigFieldLength,
		ErrTooBigSet,
		ErrTooBigScale,
		ErrTooBigPrecision,
		ErrBadNumber,
		ErrInvalidFieldSize,
		ErrMBiggerThanD,
		ErrWarnDataOutOfRange,
		ErrDuplicatedValueInType,
		ErrDatetimeFunctionOverflow,
		ErrCastAsSignedOverflow,
		ErrCastNegIntAsUnsigned,
		ErrInvalidYearFormat,
		ErrInvalidYear,
		ErrTruncatedWrongVal,
		ErrInvalidWeekModeFormat,
		ErrWrongValue,
	}
	for _, err := range kvErrs {
		code := err.ToSQLError().Code
		c.Assert(code != mysql.ErrUnknown && code == uint16(err.Code()), IsTrue, Commentf("err: %v", err))
	}
}
