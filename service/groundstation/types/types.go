// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// Details about an antenna demod decode Config used in a contact.
type AntennaDemodDecodeDetails struct {

	// Name of an antenna demod decode output node used in a contact.
	OutputNode *string
}

// Information about how AWS Ground Station should configure an antenna for
// downlink during a contact.
type AntennaDownlinkConfig struct {

	// Object that describes a spectral Config.
	//
	// This member is required.
	SpectrumConfig *SpectrumConfig
}

// Information about how AWS Ground Station should conﬁgure an antenna for downlink
// demod decode during a contact.
type AntennaDownlinkDemodDecodeConfig struct {

	// Information about the decode Config.
	//
	// This member is required.
	DecodeConfig *DecodeConfig

	// Information about the demodulation Config.
	//
	// This member is required.
	DemodulationConfig *DemodulationConfig

	// Information about the spectral Config.
	//
	// This member is required.
	SpectrumConfig *SpectrumConfig
}

// Information about the uplink Config of an antenna.
type AntennaUplinkConfig struct {

	// Information about the uplink spectral Config.
	//
	// This member is required.
	SpectrumConfig *UplinkSpectrumConfig

	// EIRP of the target.
	//
	// This member is required.
	TargetEirp *Eirp

	// Whether or not uplink transmit is disabled.
	TransmitDisabled *bool
}

// Details for certain Config object types in a contact.
//
// The following types satisfy this interface:
//  ConfigDetailsMemberEndpointDetails
//  ConfigDetailsMemberAntennaDemodDecodeDetails
type ConfigDetails interface {
	isConfigDetails()
}

// Information about the endpoint details.
type ConfigDetailsMemberEndpointDetails struct {
	Value EndpointDetails
}

func (*ConfigDetailsMemberEndpointDetails) isConfigDetails() {}

// Details for antenna demod decode Config in a contact.
type ConfigDetailsMemberAntennaDemodDecodeDetails struct {
	Value AntennaDemodDecodeDetails
}

func (*ConfigDetailsMemberAntennaDemodDecodeDetails) isConfigDetails() {}

// An item in a list of Config objects.
type ConfigListItem struct {

	// ARN of a Config.
	ConfigArn *string

	// UUID of a Config.
	ConfigId *string

	// Type of a Config.
	ConfigType ConfigCapabilityType

	// Name of a Config.
	Name *string
}

// Object containing the parameters of a Config. See the subtype definitions for
// what each type of Config contains.
//
// The following types satisfy this interface:
//  ConfigTypeDataMemberAntennaDownlinkConfig
//  ConfigTypeDataMemberTrackingConfig
//  ConfigTypeDataMemberDataflowEndpointConfig
//  ConfigTypeDataMemberAntennaDownlinkDemodDecodeConfig
//  ConfigTypeDataMemberAntennaUplinkConfig
//  ConfigTypeDataMemberUplinkEchoConfig
type ConfigTypeData interface {
	isConfigTypeData()
}

// Information about how AWS Ground Station should configure an antenna for
// downlink during a contact.
type ConfigTypeDataMemberAntennaDownlinkConfig struct {
	Value AntennaDownlinkConfig
}

func (*ConfigTypeDataMemberAntennaDownlinkConfig) isConfigTypeData() {}

// Object that determines whether tracking should be used during a contact executed
// with this Config in the mission profile.
type ConfigTypeDataMemberTrackingConfig struct {
	Value TrackingConfig
}

func (*ConfigTypeDataMemberTrackingConfig) isConfigTypeData() {}

// Information about the dataflow endpoint Config.
type ConfigTypeDataMemberDataflowEndpointConfig struct {
	Value DataflowEndpointConfig
}

func (*ConfigTypeDataMemberDataflowEndpointConfig) isConfigTypeData() {}

// Information about how AWS Ground Station should conﬁgure an antenna for downlink
// demod decode during a contact.
type ConfigTypeDataMemberAntennaDownlinkDemodDecodeConfig struct {
	Value AntennaDownlinkDemodDecodeConfig
}

func (*ConfigTypeDataMemberAntennaDownlinkDemodDecodeConfig) isConfigTypeData() {}

// Information about how AWS Ground Station should conﬁgure an antenna for uplink
// during a contact.
type ConfigTypeDataMemberAntennaUplinkConfig struct {
	Value AntennaUplinkConfig
}

func (*ConfigTypeDataMemberAntennaUplinkConfig) isConfigTypeData() {}

// Information about an uplink echo Config. Parameters from the
// AntennaUplinkConfig, corresponding to the specified AntennaUplinkConfigArn, are
// used when this UplinkEchoConfig is used in a contact.
type ConfigTypeDataMemberUplinkEchoConfig struct {
	Value UplinkEchoConfig
}

func (*ConfigTypeDataMemberUplinkEchoConfig) isConfigTypeData() {}

// Data describing a contact.
type ContactData struct {

	// UUID of a contact.
	ContactId *string

	// Status of a contact.
	ContactStatus ContactStatus

	// End time of a contact.
	EndTime *time.Time

	// Error message of a contact.
	ErrorMessage *string

	// Name of a ground station.
	GroundStation *string

	// Maximum elevation angle of a contact.
	MaximumElevation *Elevation

	// ARN of a mission profile.
	MissionProfileArn *string

	// Amount of time after a contact ends that you’d like to receive a CloudWatch
	// event indicating the pass has finished.
	PostPassEndTime *time.Time

	// Amount of time prior to contact start you’d like to receive a CloudWatch event
	// indicating an upcoming pass.
	PrePassStartTime *time.Time

	// Region of a contact.
	Region *string

	// ARN of a satellite.
	SatelliteArn *string

	// Start time of a contact.
	StartTime *time.Time

	// Tags assigned to a contact.
	Tags map[string]string
}

// Information about a dataflow edge used in a contact.
type DataflowDetail struct {

	// Dataflow details for the destination side.
	Destination *Destination

	// Error message for a dataflow.
	ErrorMessage *string

	// Dataflow details for the source side.
	Source *Source
}

// Information about a dataflow endpoint.
type DataflowEndpoint struct {

	// Socket address of a dataflow endpoint.
	Address *SocketAddress

	// Maximum transmission unit (MTU) size in bytes of a dataflow endpoint.
	Mtu *int32

	// Name of a dataflow endpoint.
	Name *string

	// Status of a dataflow endpoint.
	Status EndpointStatus
}

// Information about the dataflow endpoint Config.
type DataflowEndpointConfig struct {

	// Name of a dataflow endpoint.
	//
	// This member is required.
	DataflowEndpointName *string

	// Region of a dataflow endpoint.
	DataflowEndpointRegion *string
}

// Item in a list of DataflowEndpoint groups.
type DataflowEndpointListItem struct {

	// ARN of a dataflow endpoint group.
	DataflowEndpointGroupArn *string

	// UUID of a dataflow endpoint group.
	DataflowEndpointGroupId *string
}

// Information about the decode Config.
type DecodeConfig struct {

	// Unvalidated JSON of a decode Config.
	//
	// This member is required.
	UnvalidatedJSON *string
}

// Information about the demodulation Config.
type DemodulationConfig struct {

	// Unvalidated JSON of a demodulation Config.
	//
	// This member is required.
	UnvalidatedJSON *string
}

// Dataflow details for the destination side.
type Destination struct {

	// Additional details for a Config, if type is dataflow endpoint or antenna demod
	// decode.
	ConfigDetails ConfigDetails

	// UUID of a Config.
	ConfigId *string

	// Type of a Config.
	ConfigType ConfigCapabilityType

	// Region of a dataflow destination.
	DataflowDestinationRegion *string
}

// Object that represents EIRP.
type Eirp struct {

	// Units of an EIRP.
	//
	// This member is required.
	Units EirpUnits

	// Value of an EIRP. Valid values are between 20.0 to 50.0 dBW.
	//
	// This member is required.
	Value *float64
}

// Elevation angle of the satellite in the sky during a contact.
type Elevation struct {

	// Elevation angle units.
	//
	// This member is required.
	Unit AngleUnits

	// Elevation angle value.
	//
	// This member is required.
	Value *float64
}

// Information about the endpoint details.
type EndpointDetails struct {

	// A dataflow endpoint.
	Endpoint *DataflowEndpoint

	// Endpoint security details.
	SecurityDetails *SecurityDetails
}

// Object that describes the frequency.
type Frequency struct {

	// Frequency units.
	//
	// This member is required.
	Units FrequencyUnits

	// Frequency value. Valid values are between 2200 to 2300 MHz and 7750 to 8400 MHz
	// for downlink and 2025 to 2120 MHz for uplink.
	//
	// This member is required.
	Value *float64
}

// Object that describes the frequency bandwidth.
type FrequencyBandwidth struct {

	// Frequency bandwidth units.
	//
	// This member is required.
	Units BandwidthUnits

	// Frequency bandwidth value. AWS Ground Station currently has the following
	// bandwidth limitations:
	//
	// * For AntennaDownlinkDemodDecodeconfig, valid values are
	// between 125 kHz to 650 MHz.
	//
	// * For AntennaDownlinkconfig, valid values are
	// between 10 kHz to 54 MHz.
	//
	// * For AntennaUplinkConfig, valid values are between
	// 10 kHz to 54 MHz.
	//
	// This member is required.
	Value *float64
}

// Information about the ground station data.
type GroundStationData struct {

	// UUID of a ground station.
	GroundStationId *string

	// Name of a ground station.
	GroundStationName *string

	// Ground station Region.
	Region *string
}

// Item in a list of mission profiles.
type MissionProfileListItem struct {

	// ARN of a mission profile.
	MissionProfileArn *string

	// UUID of a mission profile.
	MissionProfileId *string

	// Name of a mission profile.
	Name *string

	// Region of a mission profile.
	Region *string
}

// Item in a list of satellites.
type SatelliteListItem struct {

	// A list of ground stations to which the satellite is on-boarded.
	GroundStations []string

	// NORAD satellite ID number.
	NoradSatelliteID int32

	// ARN of a satellite.
	SatelliteArn *string

	// UUID of a satellite.
	SatelliteId *string
}

// Information about endpoints.
type SecurityDetails struct {

	// ARN to a role needed for connecting streams to your instances.
	//
	// This member is required.
	RoleArn *string

	// The security groups to attach to the elastic network interfaces.
	//
	// This member is required.
	SecurityGroupIds []string

	// A list of subnets where AWS Ground Station places elastic network interfaces to
	// send streams to your instances.
	//
	// This member is required.
	SubnetIds []string
}

// Information about the socket address.
type SocketAddress struct {

	// Name of a socket address.
	//
	// This member is required.
	Name *string

	// Port of a socket address.
	//
	// This member is required.
	Port *int32
}

// Dataflow details for the source side.
type Source struct {

	// Additional details for a Config, if type is dataflow endpoint or antenna demod
	// decode.
	ConfigDetails ConfigDetails

	// UUID of a Config.
	ConfigId *string

	// Type of a Config.
	ConfigType ConfigCapabilityType

	// Region of a dataflow source.
	DataflowSourceRegion *string
}

// Object that describes a spectral Config.
type SpectrumConfig struct {

	// Bandwidth of a spectral Config. AWS Ground Station currently has the following
	// bandwidth limitations:
	//
	// * For AntennaDownlinkDemodDecodeconfig, valid values are
	// between 125 kHz to 650 MHz.
	//
	// * For AntennaDownlinkconfig valid values are
	// between 10 kHz to 54 MHz.
	//
	// * For AntennaUplinkConfig, valid values are between
	// 10 kHz to 54 MHz.
	//
	// This member is required.
	Bandwidth *FrequencyBandwidth

	// Center frequency of a spectral Config. Valid values are between 2200 to 2300 MHz
	// and 7750 to 8400 MHz for downlink and 2025 to 2120 MHz for uplink.
	//
	// This member is required.
	CenterFrequency *Frequency

	// Polarization of a spectral Config. Capturing both "RIGHT_HAND" and "LEFT_HAND"
	// polarization requires two separate configs.
	Polarization Polarization
}

// Object that determines whether tracking should be used during a contact executed
// with this Config in the mission profile.
type TrackingConfig struct {

	// Current setting for autotrack.
	//
	// This member is required.
	Autotrack Criticality
}

// Information about an uplink echo Config. Parameters from the
// AntennaUplinkConfig, corresponding to the specified AntennaUplinkConfigArn, are
// used when this UplinkEchoConfig is used in a contact.
type UplinkEchoConfig struct {

	// ARN of an uplink Config.
	//
	// This member is required.
	AntennaUplinkConfigArn *string

	// Whether or not an uplink Config is enabled.
	//
	// This member is required.
	Enabled *bool
}

// Information about the uplink spectral Config.
type UplinkSpectrumConfig struct {

	// Center frequency of an uplink spectral Config. Valid values are between 2025 to
	// 2120 MHz.
	//
	// This member is required.
	CenterFrequency *Frequency

	// Polarization of an uplink spectral Config. Capturing both "RIGHT_HAND" and
	// "LEFT_HAND" polarization requires two separate configs.
	Polarization Polarization
}

// UnknownUnionMember is returned when a union member is returned over the wire,
// but has an unknown tag.
type UnknownUnionMember struct {
	Tag   string
	Value []byte
}

func (*UnknownUnionMember) isConfigDetails()  {}
func (*UnknownUnionMember) isConfigTypeData() {}
