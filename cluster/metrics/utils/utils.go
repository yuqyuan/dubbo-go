/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package utils

import (
	"errors"
	"fmt"
)

import (
	"dubbo.apache.org/dubbo-go/v3/common"
)

var (
	ErrMetricsNotFound = errors.New("metrics not found")
)

func GetInvokerKey(url *common.URL) string {
	return url.Path
}

func GetInstanceKey(url *common.URL) string {
	return fmt.Sprintf("%s:%s", url.Ip, url.Port)
}

func ToFloat64(i interface{}) float64 {
	if i == nil {
		return 0
	}
	switch s := i.(type) {
	case float64:
		return s
	case float32:
		return float64(s)
	case int64:
		return float64(s)
	case int32:
		return float64(s)
	case int16:
		return float64(s)
	case int8:
		return float64(s)
	case uint:
		return float64(s)
	case uint64:
		return float64(s)
	case uint32:
		return float64(s)
	case uint16:
		return float64(s)
	case uint8:
		return float64(s)
	default:
		return 0
	}
}

func GetMethodMetricsKey(url *common.URL, methodName, key string) string {
	return fmt.Sprintf("%s.%s.%s.%s", GetInstanceKey(url), GetInvokerKey(url), methodName, key)
}