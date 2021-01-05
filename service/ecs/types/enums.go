// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AgentUpdateStatus string

// Enum values for AgentUpdateStatus
const (
	AgentUpdateStatusPending  AgentUpdateStatus = "PENDING"
	AgentUpdateStatusStaging  AgentUpdateStatus = "STAGING"
	AgentUpdateStatusStaged   AgentUpdateStatus = "STAGED"
	AgentUpdateStatusUpdating AgentUpdateStatus = "UPDATING"
	AgentUpdateStatusUpdated  AgentUpdateStatus = "UPDATED"
	AgentUpdateStatusFailed   AgentUpdateStatus = "FAILED"
)

// Values returns all known values for AgentUpdateStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AgentUpdateStatus) Values() []AgentUpdateStatus {
	return []AgentUpdateStatus{
		"PENDING",
		"STAGING",
		"STAGED",
		"UPDATING",
		"UPDATED",
		"FAILED",
	}
}

type AssignPublicIp string

// Enum values for AssignPublicIp
const (
	AssignPublicIpEnabled  AssignPublicIp = "ENABLED"
	AssignPublicIpDisabled AssignPublicIp = "DISABLED"
)

// Values returns all known values for AssignPublicIp. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AssignPublicIp) Values() []AssignPublicIp {
	return []AssignPublicIp{
		"ENABLED",
		"DISABLED",
	}
}

type CapacityProviderField string

// Enum values for CapacityProviderField
const (
	CapacityProviderFieldTags CapacityProviderField = "TAGS"
)

// Values returns all known values for CapacityProviderField. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (CapacityProviderField) Values() []CapacityProviderField {
	return []CapacityProviderField{
		"TAGS",
	}
}

type CapacityProviderStatus string

// Enum values for CapacityProviderStatus
const (
	CapacityProviderStatusActive   CapacityProviderStatus = "ACTIVE"
	CapacityProviderStatusInactive CapacityProviderStatus = "INACTIVE"
)

// Values returns all known values for CapacityProviderStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (CapacityProviderStatus) Values() []CapacityProviderStatus {
	return []CapacityProviderStatus{
		"ACTIVE",
		"INACTIVE",
	}
}

type CapacityProviderUpdateStatus string

// Enum values for CapacityProviderUpdateStatus
const (
	CapacityProviderUpdateStatusDeleteInProgress CapacityProviderUpdateStatus = "DELETE_IN_PROGRESS"
	CapacityProviderUpdateStatusDeleteComplete   CapacityProviderUpdateStatus = "DELETE_COMPLETE"
	CapacityProviderUpdateStatusDeleteFailed     CapacityProviderUpdateStatus = "DELETE_FAILED"
	CapacityProviderUpdateStatusUpdateInProgress CapacityProviderUpdateStatus = "UPDATE_IN_PROGRESS"
	CapacityProviderUpdateStatusUpdateComplete   CapacityProviderUpdateStatus = "UPDATE_COMPLETE"
	CapacityProviderUpdateStatusUpdateFailed     CapacityProviderUpdateStatus = "UPDATE_FAILED"
)

// Values returns all known values for CapacityProviderUpdateStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (CapacityProviderUpdateStatus) Values() []CapacityProviderUpdateStatus {
	return []CapacityProviderUpdateStatus{
		"DELETE_IN_PROGRESS",
		"DELETE_COMPLETE",
		"DELETE_FAILED",
		"UPDATE_IN_PROGRESS",
		"UPDATE_COMPLETE",
		"UPDATE_FAILED",
	}
}

type ClusterField string

// Enum values for ClusterField
const (
	ClusterFieldAttachments ClusterField = "ATTACHMENTS"
	ClusterFieldSettings    ClusterField = "SETTINGS"
	ClusterFieldStatistics  ClusterField = "STATISTICS"
	ClusterFieldTags        ClusterField = "TAGS"
)

// Values returns all known values for ClusterField. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ClusterField) Values() []ClusterField {
	return []ClusterField{
		"ATTACHMENTS",
		"SETTINGS",
		"STATISTICS",
		"TAGS",
	}
}

type ClusterSettingName string

// Enum values for ClusterSettingName
const (
	ClusterSettingNameContainerInsights ClusterSettingName = "containerInsights"
)

// Values returns all known values for ClusterSettingName. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ClusterSettingName) Values() []ClusterSettingName {
	return []ClusterSettingName{
		"containerInsights",
	}
}

type Compatibility string

// Enum values for Compatibility
const (
	CompatibilityEc2     Compatibility = "EC2"
	CompatibilityFargate Compatibility = "FARGATE"
)

// Values returns all known values for Compatibility. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (Compatibility) Values() []Compatibility {
	return []Compatibility{
		"EC2",
		"FARGATE",
	}
}

type Connectivity string

// Enum values for Connectivity
const (
	ConnectivityConnected    Connectivity = "CONNECTED"
	ConnectivityDisconnected Connectivity = "DISCONNECTED"
)

// Values returns all known values for Connectivity. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (Connectivity) Values() []Connectivity {
	return []Connectivity{
		"CONNECTED",
		"DISCONNECTED",
	}
}

type ContainerCondition string

// Enum values for ContainerCondition
const (
	ContainerConditionStart    ContainerCondition = "START"
	ContainerConditionComplete ContainerCondition = "COMPLETE"
	ContainerConditionSuccess  ContainerCondition = "SUCCESS"
	ContainerConditionHealthy  ContainerCondition = "HEALTHY"
)

// Values returns all known values for ContainerCondition. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ContainerCondition) Values() []ContainerCondition {
	return []ContainerCondition{
		"START",
		"COMPLETE",
		"SUCCESS",
		"HEALTHY",
	}
}

type ContainerInstanceField string

// Enum values for ContainerInstanceField
const (
	ContainerInstanceFieldTags ContainerInstanceField = "TAGS"
)

// Values returns all known values for ContainerInstanceField. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ContainerInstanceField) Values() []ContainerInstanceField {
	return []ContainerInstanceField{
		"TAGS",
	}
}

type ContainerInstanceStatus string

// Enum values for ContainerInstanceStatus
const (
	ContainerInstanceStatusActive             ContainerInstanceStatus = "ACTIVE"
	ContainerInstanceStatusDraining           ContainerInstanceStatus = "DRAINING"
	ContainerInstanceStatusRegistering        ContainerInstanceStatus = "REGISTERING"
	ContainerInstanceStatusDeregistering      ContainerInstanceStatus = "DEREGISTERING"
	ContainerInstanceStatusRegistrationFailed ContainerInstanceStatus = "REGISTRATION_FAILED"
)

// Values returns all known values for ContainerInstanceStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ContainerInstanceStatus) Values() []ContainerInstanceStatus {
	return []ContainerInstanceStatus{
		"ACTIVE",
		"DRAINING",
		"REGISTERING",
		"DEREGISTERING",
		"REGISTRATION_FAILED",
	}
}

type DeploymentControllerType string

// Enum values for DeploymentControllerType
const (
	DeploymentControllerTypeEcs        DeploymentControllerType = "ECS"
	DeploymentControllerTypeCodeDeploy DeploymentControllerType = "CODE_DEPLOY"
	DeploymentControllerTypeExternal   DeploymentControllerType = "EXTERNAL"
)

// Values returns all known values for DeploymentControllerType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DeploymentControllerType) Values() []DeploymentControllerType {
	return []DeploymentControllerType{
		"ECS",
		"CODE_DEPLOY",
		"EXTERNAL",
	}
}

type DeploymentRolloutState string

// Enum values for DeploymentRolloutState
const (
	DeploymentRolloutStateCompleted  DeploymentRolloutState = "COMPLETED"
	DeploymentRolloutStateFailed     DeploymentRolloutState = "FAILED"
	DeploymentRolloutStateInProgress DeploymentRolloutState = "IN_PROGRESS"
)

// Values returns all known values for DeploymentRolloutState. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DeploymentRolloutState) Values() []DeploymentRolloutState {
	return []DeploymentRolloutState{
		"COMPLETED",
		"FAILED",
		"IN_PROGRESS",
	}
}

type DesiredStatus string

// Enum values for DesiredStatus
const (
	DesiredStatusRunning DesiredStatus = "RUNNING"
	DesiredStatusPending DesiredStatus = "PENDING"
	DesiredStatusStopped DesiredStatus = "STOPPED"
)

// Values returns all known values for DesiredStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DesiredStatus) Values() []DesiredStatus {
	return []DesiredStatus{
		"RUNNING",
		"PENDING",
		"STOPPED",
	}
}

type DeviceCgroupPermission string

// Enum values for DeviceCgroupPermission
const (
	DeviceCgroupPermissionRead  DeviceCgroupPermission = "read"
	DeviceCgroupPermissionWrite DeviceCgroupPermission = "write"
	DeviceCgroupPermissionMknod DeviceCgroupPermission = "mknod"
)

// Values returns all known values for DeviceCgroupPermission. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DeviceCgroupPermission) Values() []DeviceCgroupPermission {
	return []DeviceCgroupPermission{
		"read",
		"write",
		"mknod",
	}
}

type EFSAuthorizationConfigIAM string

// Enum values for EFSAuthorizationConfigIAM
const (
	EFSAuthorizationConfigIAMEnabled  EFSAuthorizationConfigIAM = "ENABLED"
	EFSAuthorizationConfigIAMDisabled EFSAuthorizationConfigIAM = "DISABLED"
)

// Values returns all known values for EFSAuthorizationConfigIAM. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (EFSAuthorizationConfigIAM) Values() []EFSAuthorizationConfigIAM {
	return []EFSAuthorizationConfigIAM{
		"ENABLED",
		"DISABLED",
	}
}

type EFSTransitEncryption string

// Enum values for EFSTransitEncryption
const (
	EFSTransitEncryptionEnabled  EFSTransitEncryption = "ENABLED"
	EFSTransitEncryptionDisabled EFSTransitEncryption = "DISABLED"
)

// Values returns all known values for EFSTransitEncryption. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EFSTransitEncryption) Values() []EFSTransitEncryption {
	return []EFSTransitEncryption{
		"ENABLED",
		"DISABLED",
	}
}

type EnvironmentFileType string

// Enum values for EnvironmentFileType
const (
	EnvironmentFileTypeS3 EnvironmentFileType = "s3"
)

// Values returns all known values for EnvironmentFileType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EnvironmentFileType) Values() []EnvironmentFileType {
	return []EnvironmentFileType{
		"s3",
	}
}

type FirelensConfigurationType string

// Enum values for FirelensConfigurationType
const (
	FirelensConfigurationTypeFluentd   FirelensConfigurationType = "fluentd"
	FirelensConfigurationTypeFluentbit FirelensConfigurationType = "fluentbit"
)

// Values returns all known values for FirelensConfigurationType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (FirelensConfigurationType) Values() []FirelensConfigurationType {
	return []FirelensConfigurationType{
		"fluentd",
		"fluentbit",
	}
}

type HealthStatus string

// Enum values for HealthStatus
const (
	HealthStatusHealthy   HealthStatus = "HEALTHY"
	HealthStatusUnhealthy HealthStatus = "UNHEALTHY"
	HealthStatusUnknown   HealthStatus = "UNKNOWN"
)

// Values returns all known values for HealthStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (HealthStatus) Values() []HealthStatus {
	return []HealthStatus{
		"HEALTHY",
		"UNHEALTHY",
		"UNKNOWN",
	}
}

type IpcMode string

// Enum values for IpcMode
const (
	IpcModeHost IpcMode = "host"
	IpcModeTask IpcMode = "task"
	IpcModeNone IpcMode = "none"
)

// Values returns all known values for IpcMode. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (IpcMode) Values() []IpcMode {
	return []IpcMode{
		"host",
		"task",
		"none",
	}
}

type LaunchType string

// Enum values for LaunchType
const (
	LaunchTypeEc2     LaunchType = "EC2"
	LaunchTypeFargate LaunchType = "FARGATE"
)

// Values returns all known values for LaunchType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (LaunchType) Values() []LaunchType {
	return []LaunchType{
		"EC2",
		"FARGATE",
	}
}

type LogDriver string

// Enum values for LogDriver
const (
	LogDriverJsonFile    LogDriver = "json-file"
	LogDriverSyslog      LogDriver = "syslog"
	LogDriverJournald    LogDriver = "journald"
	LogDriverGelf        LogDriver = "gelf"
	LogDriverFluentd     LogDriver = "fluentd"
	LogDriverAwslogs     LogDriver = "awslogs"
	LogDriverSplunk      LogDriver = "splunk"
	LogDriverAwsfirelens LogDriver = "awsfirelens"
)

// Values returns all known values for LogDriver. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (LogDriver) Values() []LogDriver {
	return []LogDriver{
		"json-file",
		"syslog",
		"journald",
		"gelf",
		"fluentd",
		"awslogs",
		"splunk",
		"awsfirelens",
	}
}

type ManagedScalingStatus string

// Enum values for ManagedScalingStatus
const (
	ManagedScalingStatusEnabled  ManagedScalingStatus = "ENABLED"
	ManagedScalingStatusDisabled ManagedScalingStatus = "DISABLED"
)

// Values returns all known values for ManagedScalingStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ManagedScalingStatus) Values() []ManagedScalingStatus {
	return []ManagedScalingStatus{
		"ENABLED",
		"DISABLED",
	}
}

type ManagedTerminationProtection string

// Enum values for ManagedTerminationProtection
const (
	ManagedTerminationProtectionEnabled  ManagedTerminationProtection = "ENABLED"
	ManagedTerminationProtectionDisabled ManagedTerminationProtection = "DISABLED"
)

// Values returns all known values for ManagedTerminationProtection. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ManagedTerminationProtection) Values() []ManagedTerminationProtection {
	return []ManagedTerminationProtection{
		"ENABLED",
		"DISABLED",
	}
}

type NetworkMode string

// Enum values for NetworkMode
const (
	NetworkModeBridge NetworkMode = "bridge"
	NetworkModeHost   NetworkMode = "host"
	NetworkModeAwsvpc NetworkMode = "awsvpc"
	NetworkModeNone   NetworkMode = "none"
)

// Values returns all known values for NetworkMode. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (NetworkMode) Values() []NetworkMode {
	return []NetworkMode{
		"bridge",
		"host",
		"awsvpc",
		"none",
	}
}

type PidMode string

// Enum values for PidMode
const (
	PidModeHost PidMode = "host"
	PidModeTask PidMode = "task"
)

// Values returns all known values for PidMode. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (PidMode) Values() []PidMode {
	return []PidMode{
		"host",
		"task",
	}
}

type PlacementConstraintType string

// Enum values for PlacementConstraintType
const (
	PlacementConstraintTypeDistinctInstance PlacementConstraintType = "distinctInstance"
	PlacementConstraintTypeMemberOf         PlacementConstraintType = "memberOf"
)

// Values returns all known values for PlacementConstraintType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PlacementConstraintType) Values() []PlacementConstraintType {
	return []PlacementConstraintType{
		"distinctInstance",
		"memberOf",
	}
}

type PlacementStrategyType string

// Enum values for PlacementStrategyType
const (
	PlacementStrategyTypeRandom  PlacementStrategyType = "random"
	PlacementStrategyTypeSpread  PlacementStrategyType = "spread"
	PlacementStrategyTypeBinpack PlacementStrategyType = "binpack"
)

// Values returns all known values for PlacementStrategyType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PlacementStrategyType) Values() []PlacementStrategyType {
	return []PlacementStrategyType{
		"random",
		"spread",
		"binpack",
	}
}

type PlatformDeviceType string

// Enum values for PlatformDeviceType
const (
	PlatformDeviceTypeGpu PlatformDeviceType = "GPU"
)

// Values returns all known values for PlatformDeviceType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PlatformDeviceType) Values() []PlatformDeviceType {
	return []PlatformDeviceType{
		"GPU",
	}
}

type PropagateTags string

// Enum values for PropagateTags
const (
	PropagateTagsTaskDefinition PropagateTags = "TASK_DEFINITION"
	PropagateTagsService        PropagateTags = "SERVICE"
)

// Values returns all known values for PropagateTags. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PropagateTags) Values() []PropagateTags {
	return []PropagateTags{
		"TASK_DEFINITION",
		"SERVICE",
	}
}

type ProxyConfigurationType string

// Enum values for ProxyConfigurationType
const (
	ProxyConfigurationTypeAppmesh ProxyConfigurationType = "APPMESH"
)

// Values returns all known values for ProxyConfigurationType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ProxyConfigurationType) Values() []ProxyConfigurationType {
	return []ProxyConfigurationType{
		"APPMESH",
	}
}

type ResourceType string

// Enum values for ResourceType
const (
	ResourceTypeGpu                  ResourceType = "GPU"
	ResourceTypeInferenceAccelerator ResourceType = "InferenceAccelerator"
)

// Values returns all known values for ResourceType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ResourceType) Values() []ResourceType {
	return []ResourceType{
		"GPU",
		"InferenceAccelerator",
	}
}

type ScaleUnit string

// Enum values for ScaleUnit
const (
	ScaleUnitPercent ScaleUnit = "PERCENT"
)

// Values returns all known values for ScaleUnit. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (ScaleUnit) Values() []ScaleUnit {
	return []ScaleUnit{
		"PERCENT",
	}
}

type SchedulingStrategy string

// Enum values for SchedulingStrategy
const (
	SchedulingStrategyReplica SchedulingStrategy = "REPLICA"
	SchedulingStrategyDaemon  SchedulingStrategy = "DAEMON"
)

// Values returns all known values for SchedulingStrategy. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SchedulingStrategy) Values() []SchedulingStrategy {
	return []SchedulingStrategy{
		"REPLICA",
		"DAEMON",
	}
}

type Scope string

// Enum values for Scope
const (
	ScopeTask   Scope = "task"
	ScopeShared Scope = "shared"
)

// Values returns all known values for Scope. Note that this can be expanded in the
// future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Scope) Values() []Scope {
	return []Scope{
		"task",
		"shared",
	}
}

type ServiceField string

// Enum values for ServiceField
const (
	ServiceFieldTags ServiceField = "TAGS"
)

// Values returns all known values for ServiceField. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ServiceField) Values() []ServiceField {
	return []ServiceField{
		"TAGS",
	}
}

type SettingName string

// Enum values for SettingName
const (
	SettingNameServiceLongArnFormat           SettingName = "serviceLongArnFormat"
	SettingNameTaskLongArnFormat              SettingName = "taskLongArnFormat"
	SettingNameContainerInstanceLongArnFormat SettingName = "containerInstanceLongArnFormat"
	SettingNameAwsvpcTrunking                 SettingName = "awsvpcTrunking"
	SettingNameContainerInsights              SettingName = "containerInsights"
)

// Values returns all known values for SettingName. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (SettingName) Values() []SettingName {
	return []SettingName{
		"serviceLongArnFormat",
		"taskLongArnFormat",
		"containerInstanceLongArnFormat",
		"awsvpcTrunking",
		"containerInsights",
	}
}

type SortOrder string

// Enum values for SortOrder
const (
	SortOrderAsc  SortOrder = "ASC"
	SortOrderDesc SortOrder = "DESC"
)

// Values returns all known values for SortOrder. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (SortOrder) Values() []SortOrder {
	return []SortOrder{
		"ASC",
		"DESC",
	}
}

type StabilityStatus string

// Enum values for StabilityStatus
const (
	StabilityStatusSteadyState StabilityStatus = "STEADY_STATE"
	StabilityStatusStabilizing StabilityStatus = "STABILIZING"
)

// Values returns all known values for StabilityStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (StabilityStatus) Values() []StabilityStatus {
	return []StabilityStatus{
		"STEADY_STATE",
		"STABILIZING",
	}
}

type TargetType string

// Enum values for TargetType
const (
	TargetTypeContainerInstance TargetType = "container-instance"
)

// Values returns all known values for TargetType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (TargetType) Values() []TargetType {
	return []TargetType{
		"container-instance",
	}
}

type TaskDefinitionFamilyStatus string

// Enum values for TaskDefinitionFamilyStatus
const (
	TaskDefinitionFamilyStatusActive   TaskDefinitionFamilyStatus = "ACTIVE"
	TaskDefinitionFamilyStatusInactive TaskDefinitionFamilyStatus = "INACTIVE"
	TaskDefinitionFamilyStatusAll      TaskDefinitionFamilyStatus = "ALL"
)

// Values returns all known values for TaskDefinitionFamilyStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (TaskDefinitionFamilyStatus) Values() []TaskDefinitionFamilyStatus {
	return []TaskDefinitionFamilyStatus{
		"ACTIVE",
		"INACTIVE",
		"ALL",
	}
}

type TaskDefinitionField string

// Enum values for TaskDefinitionField
const (
	TaskDefinitionFieldTags TaskDefinitionField = "TAGS"
)

// Values returns all known values for TaskDefinitionField. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TaskDefinitionField) Values() []TaskDefinitionField {
	return []TaskDefinitionField{
		"TAGS",
	}
}

type TaskDefinitionPlacementConstraintType string

// Enum values for TaskDefinitionPlacementConstraintType
const (
	TaskDefinitionPlacementConstraintTypeMemberOf TaskDefinitionPlacementConstraintType = "memberOf"
)

// Values returns all known values for TaskDefinitionPlacementConstraintType. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (TaskDefinitionPlacementConstraintType) Values() []TaskDefinitionPlacementConstraintType {
	return []TaskDefinitionPlacementConstraintType{
		"memberOf",
	}
}

type TaskDefinitionStatus string

// Enum values for TaskDefinitionStatus
const (
	TaskDefinitionStatusActive   TaskDefinitionStatus = "ACTIVE"
	TaskDefinitionStatusInactive TaskDefinitionStatus = "INACTIVE"
)

// Values returns all known values for TaskDefinitionStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TaskDefinitionStatus) Values() []TaskDefinitionStatus {
	return []TaskDefinitionStatus{
		"ACTIVE",
		"INACTIVE",
	}
}

type TaskField string

// Enum values for TaskField
const (
	TaskFieldTags TaskField = "TAGS"
)

// Values returns all known values for TaskField. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (TaskField) Values() []TaskField {
	return []TaskField{
		"TAGS",
	}
}

type TaskSetField string

// Enum values for TaskSetField
const (
	TaskSetFieldTags TaskSetField = "TAGS"
)

// Values returns all known values for TaskSetField. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (TaskSetField) Values() []TaskSetField {
	return []TaskSetField{
		"TAGS",
	}
}

type TaskStopCode string

// Enum values for TaskStopCode
const (
	TaskStopCodeTaskFailedToStart        TaskStopCode = "TaskFailedToStart"
	TaskStopCodeEssentialContainerExited TaskStopCode = "EssentialContainerExited"
	TaskStopCodeUserInitiated            TaskStopCode = "UserInitiated"
)

// Values returns all known values for TaskStopCode. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (TaskStopCode) Values() []TaskStopCode {
	return []TaskStopCode{
		"TaskFailedToStart",
		"EssentialContainerExited",
		"UserInitiated",
	}
}

type TransportProtocol string

// Enum values for TransportProtocol
const (
	TransportProtocolTcp TransportProtocol = "tcp"
	TransportProtocolUdp TransportProtocol = "udp"
)

// Values returns all known values for TransportProtocol. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TransportProtocol) Values() []TransportProtocol {
	return []TransportProtocol{
		"tcp",
		"udp",
	}
}

type UlimitName string

// Enum values for UlimitName
const (
	UlimitNameCore       UlimitName = "core"
	UlimitNameCpu        UlimitName = "cpu"
	UlimitNameData       UlimitName = "data"
	UlimitNameFsize      UlimitName = "fsize"
	UlimitNameLocks      UlimitName = "locks"
	UlimitNameMemlock    UlimitName = "memlock"
	UlimitNameMsgqueue   UlimitName = "msgqueue"
	UlimitNameNice       UlimitName = "nice"
	UlimitNameNofile     UlimitName = "nofile"
	UlimitNameNproc      UlimitName = "nproc"
	UlimitNameRss        UlimitName = "rss"
	UlimitNameRtprio     UlimitName = "rtprio"
	UlimitNameRttime     UlimitName = "rttime"
	UlimitNameSigpending UlimitName = "sigpending"
	UlimitNameStack      UlimitName = "stack"
)

// Values returns all known values for UlimitName. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (UlimitName) Values() []UlimitName {
	return []UlimitName{
		"core",
		"cpu",
		"data",
		"fsize",
		"locks",
		"memlock",
		"msgqueue",
		"nice",
		"nofile",
		"nproc",
		"rss",
		"rtprio",
		"rttime",
		"sigpending",
		"stack",
	}
}