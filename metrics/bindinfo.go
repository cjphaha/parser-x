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

package metrics

import "github.com/prometheus/client_golang/prometheus"

// bindinfo metrics.
var (
	BindUsageCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "tidb",
			Subsystem: "bindinfo",
			Name:      "bind_usage_counter",
			Help:      "Counter of query using sql bind",
		}, []string{LableScope})

	BindTotalGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "tidb",
			Subsystem: "bindinfo",
			Name:      "bind_total_gauge",
			Help:      "Total number of sql bind",
		}, []string{LableScope, LblType})

	BindMemoryUsage = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "tidb",
			Subsystem: "bindinfo",
			Name:      "bind_memory_usage",
			Help:      "Memory usage of sql bind",
		}, []string{LableScope, LblType})
)
