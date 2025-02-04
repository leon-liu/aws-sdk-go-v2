// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type DimensionValueType string

// Enum values for DimensionValueType
const (
	DimensionValueTypeVarchar DimensionValueType = "VARCHAR"
)

// Values returns all known values for DimensionValueType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DimensionValueType) Values() []DimensionValueType {
	return []DimensionValueType{
		"VARCHAR",
	}
}

type MeasureValueType string

// Enum values for MeasureValueType
const (
	MeasureValueTypeBigint  MeasureValueType = "BIGINT"
	MeasureValueTypeBoolean MeasureValueType = "BOOLEAN"
	MeasureValueTypeDouble  MeasureValueType = "DOUBLE"
	MeasureValueTypeVarchar MeasureValueType = "VARCHAR"
	MeasureValueTypeMulti   MeasureValueType = "MULTI"
)

// Values returns all known values for MeasureValueType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (MeasureValueType) Values() []MeasureValueType {
	return []MeasureValueType{
		"BIGINT",
		"BOOLEAN",
		"DOUBLE",
		"VARCHAR",
		"MULTI",
	}
}

type S3EncryptionOption string

// Enum values for S3EncryptionOption
const (
	S3EncryptionOptionSseS3  S3EncryptionOption = "SSE_S3"
	S3EncryptionOptionSseKms S3EncryptionOption = "SSE_KMS"
)

// Values returns all known values for S3EncryptionOption. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (S3EncryptionOption) Values() []S3EncryptionOption {
	return []S3EncryptionOption{
		"SSE_S3",
		"SSE_KMS",
	}
}

type ScalarMeasureValueType string

// Enum values for ScalarMeasureValueType
const (
	ScalarMeasureValueTypeBigint  ScalarMeasureValueType = "BIGINT"
	ScalarMeasureValueTypeBoolean ScalarMeasureValueType = "BOOLEAN"
	ScalarMeasureValueTypeDouble  ScalarMeasureValueType = "DOUBLE"
	ScalarMeasureValueTypeVarchar ScalarMeasureValueType = "VARCHAR"
)

// Values returns all known values for ScalarMeasureValueType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ScalarMeasureValueType) Values() []ScalarMeasureValueType {
	return []ScalarMeasureValueType{
		"BIGINT",
		"BOOLEAN",
		"DOUBLE",
		"VARCHAR",
	}
}

type ScalarType string

// Enum values for ScalarType
const (
	ScalarTypeVarchar             ScalarType = "VARCHAR"
	ScalarTypeBoolean             ScalarType = "BOOLEAN"
	ScalarTypeBigint              ScalarType = "BIGINT"
	ScalarTypeDouble              ScalarType = "DOUBLE"
	ScalarTypeTimestamp           ScalarType = "TIMESTAMP"
	ScalarTypeDate                ScalarType = "DATE"
	ScalarTypeTime                ScalarType = "TIME"
	ScalarTypeIntervalDayToSecond ScalarType = "INTERVAL_DAY_TO_SECOND"
	ScalarTypeIntervalYearToMonth ScalarType = "INTERVAL_YEAR_TO_MONTH"
	ScalarTypeUnknown             ScalarType = "UNKNOWN"
	ScalarTypeInteger             ScalarType = "INTEGER"
)

// Values returns all known values for ScalarType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ScalarType) Values() []ScalarType {
	return []ScalarType{
		"VARCHAR",
		"BOOLEAN",
		"BIGINT",
		"DOUBLE",
		"TIMESTAMP",
		"DATE",
		"TIME",
		"INTERVAL_DAY_TO_SECOND",
		"INTERVAL_YEAR_TO_MONTH",
		"UNKNOWN",
		"INTEGER",
	}
}

type ScheduledQueryRunStatus string

// Enum values for ScheduledQueryRunStatus
const (
	ScheduledQueryRunStatusAutoTriggerSuccess   ScheduledQueryRunStatus = "AUTO_TRIGGER_SUCCESS"
	ScheduledQueryRunStatusAutoTriggerFailure   ScheduledQueryRunStatus = "AUTO_TRIGGER_FAILURE"
	ScheduledQueryRunStatusManualTriggerSuccess ScheduledQueryRunStatus = "MANUAL_TRIGGER_SUCCESS"
	ScheduledQueryRunStatusManualTriggerFailure ScheduledQueryRunStatus = "MANUAL_TRIGGER_FAILURE"
)

// Values returns all known values for ScheduledQueryRunStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ScheduledQueryRunStatus) Values() []ScheduledQueryRunStatus {
	return []ScheduledQueryRunStatus{
		"AUTO_TRIGGER_SUCCESS",
		"AUTO_TRIGGER_FAILURE",
		"MANUAL_TRIGGER_SUCCESS",
		"MANUAL_TRIGGER_FAILURE",
	}
}

type ScheduledQueryState string

// Enum values for ScheduledQueryState
const (
	ScheduledQueryStateEnabled  ScheduledQueryState = "ENABLED"
	ScheduledQueryStateDisabled ScheduledQueryState = "DISABLED"
)

// Values returns all known values for ScheduledQueryState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ScheduledQueryState) Values() []ScheduledQueryState {
	return []ScheduledQueryState{
		"ENABLED",
		"DISABLED",
	}
}
