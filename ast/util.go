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
// Copyright 2018 PingCAP, Inc.
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

package ast

import "math"

// UnspecifiedSize is unspecified size.
const (
	UnspecifiedSize = math.MaxUint64
)

// IsReadOnly checks whether the input ast is readOnly.
func IsReadOnly(node Node) bool {
	switch st := node.(type) {
	case *SelectStmt:
		if st.LockTp == SelectLockForUpdate || st.LockTp == SelectLockForUpdateNoWait {
			return false
		}

		checker := readOnlyChecker{
			readOnly: true,
		}

		node.Accept(&checker)
		return checker.readOnly
	case *ExplainStmt:
		return !st.Analyze || IsReadOnly(st.Stmt)
	case *DoStmt, *ShowStmt:
		return true
	case *UnionStmt:
		for _, sel := range node.(*UnionStmt).SelectList.Selects {
			if !IsReadOnly(sel) {
				return false
			}
		}
		return true
	default:
		return false
	}
}

// readOnlyChecker checks whether a query's ast is readonly, if it satisfied
// 1. selectstmt;
// 2. need not to set var;
// it is readonly statement.
type readOnlyChecker struct {
	readOnly bool
}

// Enter implements Visitor interface.
func (checker *readOnlyChecker) Enter(in Node) (out Node, skipChildren bool) {
	switch node := in.(type) {
	case *VariableExpr:
		// like func rewriteVariable(), this stands for SetVar.
		if !node.IsSystem && node.Value != nil {
			checker.readOnly = false
			return in, true
		}
	}
	return in, false
}

// Leave implements Visitor interface.
func (checker *readOnlyChecker) Leave(in Node) (out Node, ok bool) {
	return in, checker.readOnly
}
