// Copyright 2020 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package eventpb

import (
	"reflect"
	"strings"

	"github.com/cockroachdb/cockroach/pkg/util/log/logpb"
)

// GetEventTypeName retrieves the system.eventlog type name for the given payload.
func GetEventTypeName(event EventPayload) string {
	// This logic takes the type names and converts from CamelCase to snake_case.
	typeName := reflect.TypeOf(event).Elem().Name()
	var res strings.Builder
	res.WriteByte(typeName[0] + 'a' - 'A')
	for i := 1; i < len(typeName); i++ {
		if typeName[i] >= 'A' && typeName[i] <= 'Z' {
			res.WriteByte('_')
			res.WriteByte(typeName[i] + 'a' - 'A')
		} else {
			res.WriteByte(typeName[i])
		}
	}
	return res.String()
}

// EventPayload is implemented by CommonEventDetails.
type EventPayload interface {
	CommonDetails() *CommonEventDetails
	LoggingChannel() logpb.Channel
}

// CommonDetails implements the EventWithCommonPayload interface.
func (m *CommonEventDetails) CommonDetails() *CommonEventDetails { return m }

// EventWithCommonSQLPayload is implemented by CommonSQLEventDetails.
type EventWithCommonSQLPayload interface {
	EventPayload
	CommonSQLDetails() *CommonSQLEventDetails
}

// CommonSQLDetails implements the EventWithCommonSQLPayload interface.
func (m *CommonSQLEventDetails) CommonSQLDetails() *CommonSQLEventDetails { return m }

// EventWithCommonSchemaChangePayload is implemented by CommonSchemaChangeDetails.
type EventWithCommonSchemaChangePayload interface {
	EventPayload
	CommonSchemaChangeDetails() *CommonSchemaChangeEventDetails
}

// CommonSchemaChangeDetails implements the EventWithCommonSchemaChangePayload interface.
func (m *CommonSchemaChangeEventDetails) CommonSchemaChangeDetails() *CommonSchemaChangeEventDetails {
	return m
}

var _ EventWithCommonSchemaChangePayload = (*FinishSchemaChange)(nil)
var _ EventWithCommonSchemaChangePayload = (*ReverseSchemaChange)(nil)
var _ EventWithCommonSchemaChangePayload = (*FinishSchemaChangeRollback)(nil)
