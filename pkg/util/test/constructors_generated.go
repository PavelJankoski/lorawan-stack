// Code generated by generate_constructors.go. DO NOT EDIT.

package test

import (
	pbtypes "github.com/gogo/protobuf/types"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/types"
)

type (
	// RootKeysOption transforms ttnpb.RootKeys and returns it.
	// Implemetations must be pure functions with no side-effects.
	RootKeysOption func(ttnpb.RootKeys) ttnpb.RootKeys

	// RootKeysOptionNamespace represents the namespace, on which various RootKeysOption are defined.
	RootKeysOptionNamespace struct{}
)

// WithRootKeyId returns a RootKeysOption, which returns a copy of ttnpb.RootKeys with RootKeyId set to v.
func (RootKeysOptionNamespace) WithRootKeyId(v string) RootKeysOption {
	return func(x ttnpb.RootKeys) ttnpb.RootKeys {
		x.RootKeyId = v
		return x
	}
}

// WithAppKey returns a RootKeysOption, which returns a copy of ttnpb.RootKeys with AppKey set to v.
func (RootKeysOptionNamespace) WithAppKey(v *ttnpb.KeyEnvelope) RootKeysOption {
	return func(x ttnpb.RootKeys) ttnpb.RootKeys {
		x.AppKey = v
		return x
	}
}

// WithNwkKey returns a RootKeysOption, which returns a copy of ttnpb.RootKeys with NwkKey set to v.
func (RootKeysOptionNamespace) WithNwkKey(v *ttnpb.KeyEnvelope) RootKeysOption {
	return func(x ttnpb.RootKeys) ttnpb.RootKeys {
		x.NwkKey = v
		return x
	}
}

// WithXXX_unrecognized returns a RootKeysOption, which returns a copy of ttnpb.RootKeys with XXX_unrecognized set to v.
func (RootKeysOptionNamespace) WithXXX_unrecognized(v []byte) RootKeysOption {
	return func(x ttnpb.RootKeys) ttnpb.RootKeys {
		x.XXX_unrecognized = v
		return x
	}
}

// Compose returns a functional composition of opts as a singular RootKeysOption.
func (RootKeysOptionNamespace) Compose(opts ...RootKeysOption) RootKeysOption {
	return func(x ttnpb.RootKeys) ttnpb.RootKeys {
		for _, opt := range opts {
			x = opt(x)
		}
		return x
	}
}

// Compose returns a functional composition of f and opts as a singular RootKeysOption.
func (f RootKeysOption) Compose(opts ...RootKeysOption) RootKeysOption {
	return func(x ttnpb.RootKeys) ttnpb.RootKeys {
		x = f(x)
		for _, opt := range opts {
			x = opt(x)
		}
		return x
	}
}

// RootKeysOptions is namespace containing ttnpb.RootKeys options.
var RootKeysOptions RootKeysOptionNamespace

// MakeRootKeys constructs a new ttnpb.RootKeys.
func MakeRootKeys(opts ...RootKeysOption) *ttnpb.RootKeys {
	v := RootKeysOptions.Compose(opts...)(baseRootKeys)
	return &v
}

type (
	// SessionKeysOption transforms ttnpb.SessionKeys and returns it.
	// Implemetations must be pure functions with no side-effects.
	SessionKeysOption func(ttnpb.SessionKeys) ttnpb.SessionKeys

	// SessionKeysOptionNamespace represents the namespace, on which various SessionKeysOption are defined.
	SessionKeysOptionNamespace struct{}
)

// WithSessionKeyId returns a SessionKeysOption, which returns a copy of ttnpb.SessionKeys with SessionKeyId set to v.
func (SessionKeysOptionNamespace) WithSessionKeyId(v []byte) SessionKeysOption {
	return func(x ttnpb.SessionKeys) ttnpb.SessionKeys {
		x.SessionKeyId = v
		return x
	}
}

// WithFNwkSIntKey returns a SessionKeysOption, which returns a copy of ttnpb.SessionKeys with FNwkSIntKey set to v.
func (SessionKeysOptionNamespace) WithFNwkSIntKey(v *ttnpb.KeyEnvelope) SessionKeysOption {
	return func(x ttnpb.SessionKeys) ttnpb.SessionKeys {
		x.FNwkSIntKey = v
		return x
	}
}

// WithSNwkSIntKey returns a SessionKeysOption, which returns a copy of ttnpb.SessionKeys with SNwkSIntKey set to v.
func (SessionKeysOptionNamespace) WithSNwkSIntKey(v *ttnpb.KeyEnvelope) SessionKeysOption {
	return func(x ttnpb.SessionKeys) ttnpb.SessionKeys {
		x.SNwkSIntKey = v
		return x
	}
}

// WithNwkSEncKey returns a SessionKeysOption, which returns a copy of ttnpb.SessionKeys with NwkSEncKey set to v.
func (SessionKeysOptionNamespace) WithNwkSEncKey(v *ttnpb.KeyEnvelope) SessionKeysOption {
	return func(x ttnpb.SessionKeys) ttnpb.SessionKeys {
		x.NwkSEncKey = v
		return x
	}
}

// WithAppSKey returns a SessionKeysOption, which returns a copy of ttnpb.SessionKeys with AppSKey set to v.
func (SessionKeysOptionNamespace) WithAppSKey(v *ttnpb.KeyEnvelope) SessionKeysOption {
	return func(x ttnpb.SessionKeys) ttnpb.SessionKeys {
		x.AppSKey = v
		return x
	}
}

// WithXXX_unrecognized returns a SessionKeysOption, which returns a copy of ttnpb.SessionKeys with XXX_unrecognized set to v.
func (SessionKeysOptionNamespace) WithXXX_unrecognized(v []byte) SessionKeysOption {
	return func(x ttnpb.SessionKeys) ttnpb.SessionKeys {
		x.XXX_unrecognized = v
		return x
	}
}

// Compose returns a functional composition of opts as a singular SessionKeysOption.
func (SessionKeysOptionNamespace) Compose(opts ...SessionKeysOption) SessionKeysOption {
	return func(x ttnpb.SessionKeys) ttnpb.SessionKeys {
		for _, opt := range opts {
			x = opt(x)
		}
		return x
	}
}

// Compose returns a functional composition of f and opts as a singular SessionKeysOption.
func (f SessionKeysOption) Compose(opts ...SessionKeysOption) SessionKeysOption {
	return func(x ttnpb.SessionKeys) ttnpb.SessionKeys {
		x = f(x)
		for _, opt := range opts {
			x = opt(x)
		}
		return x
	}
}

// SessionKeysOptions is namespace containing ttnpb.SessionKeys options.
var SessionKeysOptions SessionKeysOptionNamespace

// MakeSessionKeys constructs a new ttnpb.SessionKeys.
func MakeSessionKeys(opts ...SessionKeysOption) *ttnpb.SessionKeys {
	v := SessionKeysOptions.Compose(opts...)(baseSessionKeys)
	return &v
}

type (
	// SessionOption transforms ttnpb.Session and returns it.
	// Implemetations must be pure functions with no side-effects.
	SessionOption func(ttnpb.Session) ttnpb.Session

	// SessionOptionNamespace represents the namespace, on which various SessionOption are defined.
	SessionOptionNamespace struct{}
)

// WithDevAddr returns a SessionOption, which returns a copy of ttnpb.Session with DevAddr set to v.
func (SessionOptionNamespace) WithDevAddr(v []byte) SessionOption {
	return func(x ttnpb.Session) ttnpb.Session {
		x.DevAddr = v
		return x
	}
}

// WithKeys returns a SessionOption, which returns a copy of ttnpb.Session with Keys set to v.
func (SessionOptionNamespace) WithKeys(v *ttnpb.SessionKeys) SessionOption {
	return func(x ttnpb.Session) ttnpb.Session {
		x.Keys = v
		return x
	}
}

// WithLastFCntUp returns a SessionOption, which returns a copy of ttnpb.Session with LastFCntUp set to v.
func (SessionOptionNamespace) WithLastFCntUp(v uint32) SessionOption {
	return func(x ttnpb.Session) ttnpb.Session {
		x.LastFCntUp = v
		return x
	}
}

// WithLastNFCntDown returns a SessionOption, which returns a copy of ttnpb.Session with LastNFCntDown set to v.
func (SessionOptionNamespace) WithLastNFCntDown(v uint32) SessionOption {
	return func(x ttnpb.Session) ttnpb.Session {
		x.LastNFCntDown = v
		return x
	}
}

// WithLastAFCntDown returns a SessionOption, which returns a copy of ttnpb.Session with LastAFCntDown set to v.
func (SessionOptionNamespace) WithLastAFCntDown(v uint32) SessionOption {
	return func(x ttnpb.Session) ttnpb.Session {
		x.LastAFCntDown = v
		return x
	}
}

// WithLastConfFCntDown returns a SessionOption, which returns a copy of ttnpb.Session with LastConfFCntDown set to v.
func (SessionOptionNamespace) WithLastConfFCntDown(v uint32) SessionOption {
	return func(x ttnpb.Session) ttnpb.Session {
		x.LastConfFCntDown = v
		return x
	}
}

// WithStartedAt returns a SessionOption, which returns a copy of ttnpb.Session with StartedAt set to v.
func (SessionOptionNamespace) WithStartedAt(v *pbtypes.Timestamp) SessionOption {
	return func(x ttnpb.Session) ttnpb.Session {
		x.StartedAt = v
		return x
	}
}

// WithQueuedApplicationDownlinks returns a SessionOption, which returns a copy of ttnpb.Session with QueuedApplicationDownlinks set to vs.
func (SessionOptionNamespace) WithQueuedApplicationDownlinks(vs ...*ttnpb.ApplicationDownlink) SessionOption {
	return func(x ttnpb.Session) ttnpb.Session {
		x.QueuedApplicationDownlinks = vs
		return x
	}
}

// WithXXX_unrecognized returns a SessionOption, which returns a copy of ttnpb.Session with XXX_unrecognized set to v.
func (SessionOptionNamespace) WithXXX_unrecognized(v []byte) SessionOption {
	return func(x ttnpb.Session) ttnpb.Session {
		x.XXX_unrecognized = v
		return x
	}
}

// Compose returns a functional composition of opts as a singular SessionOption.
func (SessionOptionNamespace) Compose(opts ...SessionOption) SessionOption {
	return func(x ttnpb.Session) ttnpb.Session {
		for _, opt := range opts {
			x = opt(x)
		}
		return x
	}
}

// Compose returns a functional composition of f and opts as a singular SessionOption.
func (f SessionOption) Compose(opts ...SessionOption) SessionOption {
	return func(x ttnpb.Session) ttnpb.Session {
		x = f(x)
		for _, opt := range opts {
			x = opt(x)
		}
		return x
	}
}

// SessionOptions is namespace containing ttnpb.Session options.
var SessionOptions SessionOptionNamespace

// MakeSession constructs a new ttnpb.Session.
func MakeSession(opts ...SessionOption) *ttnpb.Session {
	v := SessionOptions.Compose(opts...)(baseSession)
	return &v
}

type (
	// MACStateOption transforms ttnpb.MACState and returns it.
	// Implemetations must be pure functions with no side-effects.
	MACStateOption func(ttnpb.MACState) ttnpb.MACState

	// MACStateOptionNamespace represents the namespace, on which various MACStateOption are defined.
	MACStateOptionNamespace struct{}
)

// WithCurrentParameters returns a MACStateOption, which returns a copy of ttnpb.MACState with CurrentParameters set to v.
func (MACStateOptionNamespace) WithCurrentParameters(v *ttnpb.MACParameters) MACStateOption {
	return func(x ttnpb.MACState) ttnpb.MACState {
		x.CurrentParameters = v
		return x
	}
}

// WithDesiredParameters returns a MACStateOption, which returns a copy of ttnpb.MACState with DesiredParameters set to v.
func (MACStateOptionNamespace) WithDesiredParameters(v *ttnpb.MACParameters) MACStateOption {
	return func(x ttnpb.MACState) ttnpb.MACState {
		x.DesiredParameters = v
		return x
	}
}

// WithDeviceClass returns a MACStateOption, which returns a copy of ttnpb.MACState with DeviceClass set to v.
func (MACStateOptionNamespace) WithDeviceClass(v ttnpb.Class) MACStateOption {
	return func(x ttnpb.MACState) ttnpb.MACState {
		x.DeviceClass = v
		return x
	}
}

// WithLorawanVersion returns a MACStateOption, which returns a copy of ttnpb.MACState with LorawanVersion set to v.
func (MACStateOptionNamespace) WithLorawanVersion(v ttnpb.MACVersion) MACStateOption {
	return func(x ttnpb.MACState) ttnpb.MACState {
		x.LorawanVersion = v
		return x
	}
}

// WithLastConfirmedDownlinkAt returns a MACStateOption, which returns a copy of ttnpb.MACState with LastConfirmedDownlinkAt set to v.
func (MACStateOptionNamespace) WithLastConfirmedDownlinkAt(v *pbtypes.Timestamp) MACStateOption {
	return func(x ttnpb.MACState) ttnpb.MACState {
		x.LastConfirmedDownlinkAt = v
		return x
	}
}

// WithLastDevStatusFCntUp returns a MACStateOption, which returns a copy of ttnpb.MACState with LastDevStatusFCntUp set to v.
func (MACStateOptionNamespace) WithLastDevStatusFCntUp(v uint32) MACStateOption {
	return func(x ttnpb.MACState) ttnpb.MACState {
		x.LastDevStatusFCntUp = v
		return x
	}
}

// WithPingSlotPeriodicity returns a MACStateOption, which returns a copy of ttnpb.MACState with PingSlotPeriodicity set to v.
func (MACStateOptionNamespace) WithPingSlotPeriodicity(v *ttnpb.PingSlotPeriodValue) MACStateOption {
	return func(x ttnpb.MACState) ttnpb.MACState {
		x.PingSlotPeriodicity = v
		return x
	}
}

// WithPendingApplicationDownlink returns a MACStateOption, which returns a copy of ttnpb.MACState with PendingApplicationDownlink set to v.
func (MACStateOptionNamespace) WithPendingApplicationDownlink(v *ttnpb.ApplicationDownlink) MACStateOption {
	return func(x ttnpb.MACState) ttnpb.MACState {
		x.PendingApplicationDownlink = v
		return x
	}
}

// WithQueuedResponses returns a MACStateOption, which returns a copy of ttnpb.MACState with QueuedResponses set to vs.
func (MACStateOptionNamespace) WithQueuedResponses(vs ...*ttnpb.MACCommand) MACStateOption {
	return func(x ttnpb.MACState) ttnpb.MACState {
		x.QueuedResponses = vs
		return x
	}
}

// WithPendingRequests returns a MACStateOption, which returns a copy of ttnpb.MACState with PendingRequests set to vs.
func (MACStateOptionNamespace) WithPendingRequests(vs ...*ttnpb.MACCommand) MACStateOption {
	return func(x ttnpb.MACState) ttnpb.MACState {
		x.PendingRequests = vs
		return x
	}
}

// WithQueuedJoinAccept returns a MACStateOption, which returns a copy of ttnpb.MACState with QueuedJoinAccept set to v.
func (MACStateOptionNamespace) WithQueuedJoinAccept(v *ttnpb.MACState_JoinAccept) MACStateOption {
	return func(x ttnpb.MACState) ttnpb.MACState {
		x.QueuedJoinAccept = v
		return x
	}
}

// WithPendingJoinRequest returns a MACStateOption, which returns a copy of ttnpb.MACState with PendingJoinRequest set to v.
func (MACStateOptionNamespace) WithPendingJoinRequest(v *ttnpb.MACState_JoinRequest) MACStateOption {
	return func(x ttnpb.MACState) ttnpb.MACState {
		x.PendingJoinRequest = v
		return x
	}
}

// WithRxWindowsAvailable returns a MACStateOption, which returns a copy of ttnpb.MACState with RxWindowsAvailable set to v.
func (MACStateOptionNamespace) WithRxWindowsAvailable(v bool) MACStateOption {
	return func(x ttnpb.MACState) ttnpb.MACState {
		x.RxWindowsAvailable = v
		return x
	}
}

// WithRecentUplinks returns a MACStateOption, which returns a copy of ttnpb.MACState with RecentUplinks set to vs.
func (MACStateOptionNamespace) WithRecentUplinks(vs ...*ttnpb.MACState_UplinkMessage) MACStateOption {
	return func(x ttnpb.MACState) ttnpb.MACState {
		x.RecentUplinks = vs
		return x
	}
}

// WithRecentDownlinks returns a MACStateOption, which returns a copy of ttnpb.MACState with RecentDownlinks set to vs.
func (MACStateOptionNamespace) WithRecentDownlinks(vs ...*ttnpb.DownlinkMessage) MACStateOption {
	return func(x ttnpb.MACState) ttnpb.MACState {
		x.RecentDownlinks = vs
		return x
	}
}

// WithLastNetworkInitiatedDownlinkAt returns a MACStateOption, which returns a copy of ttnpb.MACState with LastNetworkInitiatedDownlinkAt set to v.
func (MACStateOptionNamespace) WithLastNetworkInitiatedDownlinkAt(v *pbtypes.Timestamp) MACStateOption {
	return func(x ttnpb.MACState) ttnpb.MACState {
		x.LastNetworkInitiatedDownlinkAt = v
		return x
	}
}

// WithRejectedAdrDataRateIndexes returns a MACStateOption, which returns a copy of ttnpb.MACState with RejectedAdrDataRateIndexes set to vs.
func (MACStateOptionNamespace) WithRejectedAdrDataRateIndexes(vs ...ttnpb.DataRateIndex) MACStateOption {
	return func(x ttnpb.MACState) ttnpb.MACState {
		x.RejectedAdrDataRateIndexes = vs
		return x
	}
}

// WithRejectedAdrTxPowerIndexes returns a MACStateOption, which returns a copy of ttnpb.MACState with RejectedAdrTxPowerIndexes set to vs.
func (MACStateOptionNamespace) WithRejectedAdrTxPowerIndexes(vs ...uint32) MACStateOption {
	return func(x ttnpb.MACState) ttnpb.MACState {
		x.RejectedAdrTxPowerIndexes = vs
		return x
	}
}

// WithRejectedFrequencies returns a MACStateOption, which returns a copy of ttnpb.MACState with RejectedFrequencies set to vs.
func (MACStateOptionNamespace) WithRejectedFrequencies(vs ...uint64) MACStateOption {
	return func(x ttnpb.MACState) ttnpb.MACState {
		x.RejectedFrequencies = vs
		return x
	}
}

// WithLastDownlinkAt returns a MACStateOption, which returns a copy of ttnpb.MACState with LastDownlinkAt set to v.
func (MACStateOptionNamespace) WithLastDownlinkAt(v *pbtypes.Timestamp) MACStateOption {
	return func(x ttnpb.MACState) ttnpb.MACState {
		x.LastDownlinkAt = v
		return x
	}
}

// WithRejectedDataRateRanges returns a MACStateOption, which returns a copy of ttnpb.MACState with RejectedDataRateRanges set to v.
func (MACStateOptionNamespace) WithRejectedDataRateRanges(v map[uint64]*ttnpb.MACState_DataRateRanges) MACStateOption {
	return func(x ttnpb.MACState) ttnpb.MACState {
		x.RejectedDataRateRanges = v
		return x
	}
}

// WithLastAdrChangeFCntUp returns a MACStateOption, which returns a copy of ttnpb.MACState with LastAdrChangeFCntUp set to v.
func (MACStateOptionNamespace) WithLastAdrChangeFCntUp(v uint32) MACStateOption {
	return func(x ttnpb.MACState) ttnpb.MACState {
		x.LastAdrChangeFCntUp = v
		return x
	}
}

// WithXXX_unrecognized returns a MACStateOption, which returns a copy of ttnpb.MACState with XXX_unrecognized set to v.
func (MACStateOptionNamespace) WithXXX_unrecognized(v []byte) MACStateOption {
	return func(x ttnpb.MACState) ttnpb.MACState {
		x.XXX_unrecognized = v
		return x
	}
}

// Compose returns a functional composition of opts as a singular MACStateOption.
func (MACStateOptionNamespace) Compose(opts ...MACStateOption) MACStateOption {
	return func(x ttnpb.MACState) ttnpb.MACState {
		for _, opt := range opts {
			x = opt(x)
		}
		return x
	}
}

// Compose returns a functional composition of f and opts as a singular MACStateOption.
func (f MACStateOption) Compose(opts ...MACStateOption) MACStateOption {
	return func(x ttnpb.MACState) ttnpb.MACState {
		x = f(x)
		for _, opt := range opts {
			x = opt(x)
		}
		return x
	}
}

// MACStateOptions is namespace containing ttnpb.MACState options.
var MACStateOptions MACStateOptionNamespace

// MakeMACState constructs a new ttnpb.MACState.
func MakeMACState(opts ...MACStateOption) *ttnpb.MACState {
	v := MACStateOptions.Compose(opts...)(baseMACState)
	return &v
}

type (
	// EndDeviceIdentifiersOption transforms ttnpb.EndDeviceIdentifiers and returns it.
	// Implemetations must be pure functions with no side-effects.
	EndDeviceIdentifiersOption func(ttnpb.EndDeviceIdentifiers) ttnpb.EndDeviceIdentifiers

	// EndDeviceIdentifiersOptionNamespace represents the namespace, on which various EndDeviceIdentifiersOption are defined.
	EndDeviceIdentifiersOptionNamespace struct{}
)

// WithDeviceId returns a EndDeviceIdentifiersOption, which returns a copy of ttnpb.EndDeviceIdentifiers with DeviceId set to v.
func (EndDeviceIdentifiersOptionNamespace) WithDeviceId(v string) EndDeviceIdentifiersOption {
	return func(x ttnpb.EndDeviceIdentifiers) ttnpb.EndDeviceIdentifiers {
		x.DeviceId = v
		return x
	}
}

// WithApplicationIds returns a EndDeviceIdentifiersOption, which returns a copy of ttnpb.EndDeviceIdentifiers with ApplicationIds set to v.
func (EndDeviceIdentifiersOptionNamespace) WithApplicationIds(v *ttnpb.ApplicationIdentifiers) EndDeviceIdentifiersOption {
	return func(x ttnpb.EndDeviceIdentifiers) ttnpb.EndDeviceIdentifiers {
		x.ApplicationIds = v
		return x
	}
}

// WithDevEui returns a EndDeviceIdentifiersOption, which returns a copy of ttnpb.EndDeviceIdentifiers with DevEui set to v.
func (EndDeviceIdentifiersOptionNamespace) WithDevEui(v *types.EUI64) EndDeviceIdentifiersOption {
	return func(x ttnpb.EndDeviceIdentifiers) ttnpb.EndDeviceIdentifiers {
		x.DevEui = v
		return x
	}
}

// WithJoinEui returns a EndDeviceIdentifiersOption, which returns a copy of ttnpb.EndDeviceIdentifiers with JoinEui set to v.
func (EndDeviceIdentifiersOptionNamespace) WithJoinEui(v *types.EUI64) EndDeviceIdentifiersOption {
	return func(x ttnpb.EndDeviceIdentifiers) ttnpb.EndDeviceIdentifiers {
		x.JoinEui = v
		return x
	}
}

// WithDevAddr returns a EndDeviceIdentifiersOption, which returns a copy of ttnpb.EndDeviceIdentifiers with DevAddr set to v.
func (EndDeviceIdentifiersOptionNamespace) WithDevAddr(v *types.DevAddr) EndDeviceIdentifiersOption {
	return func(x ttnpb.EndDeviceIdentifiers) ttnpb.EndDeviceIdentifiers {
		x.DevAddr = v
		return x
	}
}

// WithXXX_unrecognized returns a EndDeviceIdentifiersOption, which returns a copy of ttnpb.EndDeviceIdentifiers with XXX_unrecognized set to v.
func (EndDeviceIdentifiersOptionNamespace) WithXXX_unrecognized(v []byte) EndDeviceIdentifiersOption {
	return func(x ttnpb.EndDeviceIdentifiers) ttnpb.EndDeviceIdentifiers {
		x.XXX_unrecognized = v
		return x
	}
}

// Compose returns a functional composition of opts as a singular EndDeviceIdentifiersOption.
func (EndDeviceIdentifiersOptionNamespace) Compose(opts ...EndDeviceIdentifiersOption) EndDeviceIdentifiersOption {
	return func(x ttnpb.EndDeviceIdentifiers) ttnpb.EndDeviceIdentifiers {
		for _, opt := range opts {
			x = opt(x)
		}
		return x
	}
}

// Compose returns a functional composition of f and opts as a singular EndDeviceIdentifiersOption.
func (f EndDeviceIdentifiersOption) Compose(opts ...EndDeviceIdentifiersOption) EndDeviceIdentifiersOption {
	return func(x ttnpb.EndDeviceIdentifiers) ttnpb.EndDeviceIdentifiers {
		x = f(x)
		for _, opt := range opts {
			x = opt(x)
		}
		return x
	}
}

// EndDeviceIdentifiersOptions is namespace containing ttnpb.EndDeviceIdentifiers options.
var EndDeviceIdentifiersOptions EndDeviceIdentifiersOptionNamespace

// MakeEndDeviceIdentifiers constructs a new ttnpb.EndDeviceIdentifiers.
func MakeEndDeviceIdentifiers(opts ...EndDeviceIdentifiersOption) *ttnpb.EndDeviceIdentifiers {
	v := EndDeviceIdentifiersOptions.Compose(opts...)(baseEndDeviceIdentifiers)
	return &v
}

type (
	// EndDeviceOption transforms ttnpb.EndDevice and returns it.
	// Implemetations must be pure functions with no side-effects.
	EndDeviceOption func(ttnpb.EndDevice) ttnpb.EndDevice

	// EndDeviceOptionNamespace represents the namespace, on which various EndDeviceOption are defined.
	EndDeviceOptionNamespace struct{}
)

// WithIds returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with Ids set to v.
func (EndDeviceOptionNamespace) WithIds(v *ttnpb.EndDeviceIdentifiers) EndDeviceOption {
	return func(x ttnpb.EndDevice) ttnpb.EndDevice {
		x.Ids = v
		return x
	}
}

// WithCreatedAt returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with CreatedAt set to v.
func (EndDeviceOptionNamespace) WithCreatedAt(v *pbtypes.Timestamp) EndDeviceOption {
	return func(x ttnpb.EndDevice) ttnpb.EndDevice {
		x.CreatedAt = v
		return x
	}
}

// WithUpdatedAt returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with UpdatedAt set to v.
func (EndDeviceOptionNamespace) WithUpdatedAt(v *pbtypes.Timestamp) EndDeviceOption {
	return func(x ttnpb.EndDevice) ttnpb.EndDevice {
		x.UpdatedAt = v
		return x
	}
}

// WithName returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with Name set to v.
func (EndDeviceOptionNamespace) WithName(v string) EndDeviceOption {
	return func(x ttnpb.EndDevice) ttnpb.EndDevice {
		x.Name = v
		return x
	}
}

// WithDescription returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with Description set to v.
func (EndDeviceOptionNamespace) WithDescription(v string) EndDeviceOption {
	return func(x ttnpb.EndDevice) ttnpb.EndDevice {
		x.Description = v
		return x
	}
}

// WithAttributes returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with Attributes set to v.
func (EndDeviceOptionNamespace) WithAttributes(v map[string]string) EndDeviceOption {
	return func(x ttnpb.EndDevice) ttnpb.EndDevice {
		x.Attributes = v
		return x
	}
}

// WithVersionIds returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with VersionIds set to v.
func (EndDeviceOptionNamespace) WithVersionIds(v *ttnpb.EndDeviceVersionIdentifiers) EndDeviceOption {
	return func(x ttnpb.EndDevice) ttnpb.EndDevice {
		x.VersionIds = v
		return x
	}
}

// WithServiceProfileId returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with ServiceProfileId set to v.
func (EndDeviceOptionNamespace) WithServiceProfileId(v string) EndDeviceOption {
	return func(x ttnpb.EndDevice) ttnpb.EndDevice {
		x.ServiceProfileId = v
		return x
	}
}

// WithNetworkServerAddress returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with NetworkServerAddress set to v.
func (EndDeviceOptionNamespace) WithNetworkServerAddress(v string) EndDeviceOption {
	return func(x ttnpb.EndDevice) ttnpb.EndDevice {
		x.NetworkServerAddress = v
		return x
	}
}

// WithNetworkServerKekLabel returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with NetworkServerKekLabel set to v.
func (EndDeviceOptionNamespace) WithNetworkServerKekLabel(v string) EndDeviceOption {
	return func(x ttnpb.EndDevice) ttnpb.EndDevice {
		x.NetworkServerKekLabel = v
		return x
	}
}

// WithApplicationServerAddress returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with ApplicationServerAddress set to v.
func (EndDeviceOptionNamespace) WithApplicationServerAddress(v string) EndDeviceOption {
	return func(x ttnpb.EndDevice) ttnpb.EndDevice {
		x.ApplicationServerAddress = v
		return x
	}
}

// WithApplicationServerKekLabel returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with ApplicationServerKekLabel set to v.
func (EndDeviceOptionNamespace) WithApplicationServerKekLabel(v string) EndDeviceOption {
	return func(x ttnpb.EndDevice) ttnpb.EndDevice {
		x.ApplicationServerKekLabel = v
		return x
	}
}

// WithApplicationServerId returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with ApplicationServerId set to v.
func (EndDeviceOptionNamespace) WithApplicationServerId(v string) EndDeviceOption {
	return func(x ttnpb.EndDevice) ttnpb.EndDevice {
		x.ApplicationServerId = v
		return x
	}
}

// WithJoinServerAddress returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with JoinServerAddress set to v.
func (EndDeviceOptionNamespace) WithJoinServerAddress(v string) EndDeviceOption {
	return func(x ttnpb.EndDevice) ttnpb.EndDevice {
		x.JoinServerAddress = v
		return x
	}
}

// WithLocations returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with Locations set to v.
func (EndDeviceOptionNamespace) WithLocations(v map[string]*ttnpb.Location) EndDeviceOption {
	return func(x ttnpb.EndDevice) ttnpb.EndDevice {
		x.Locations = v
		return x
	}
}

// WithPicture returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with Picture set to v.
func (EndDeviceOptionNamespace) WithPicture(v *ttnpb.Picture) EndDeviceOption {
	return func(x ttnpb.EndDevice) ttnpb.EndDevice {
		x.Picture = v
		return x
	}
}

// WithSupportsClassB returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with SupportsClassB set to v.
func (EndDeviceOptionNamespace) WithSupportsClassB(v bool) EndDeviceOption {
	return func(x ttnpb.EndDevice) ttnpb.EndDevice {
		x.SupportsClassB = v
		return x
	}
}

// WithSupportsClassC returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with SupportsClassC set to v.
func (EndDeviceOptionNamespace) WithSupportsClassC(v bool) EndDeviceOption {
	return func(x ttnpb.EndDevice) ttnpb.EndDevice {
		x.SupportsClassC = v
		return x
	}
}

// WithLorawanVersion returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with LorawanVersion set to v.
func (EndDeviceOptionNamespace) WithLorawanVersion(v ttnpb.MACVersion) EndDeviceOption {
	return func(x ttnpb.EndDevice) ttnpb.EndDevice {
		x.LorawanVersion = v
		return x
	}
}

// WithLorawanPhyVersion returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with LorawanPhyVersion set to v.
func (EndDeviceOptionNamespace) WithLorawanPhyVersion(v ttnpb.PHYVersion) EndDeviceOption {
	return func(x ttnpb.EndDevice) ttnpb.EndDevice {
		x.LorawanPhyVersion = v
		return x
	}
}

// WithFrequencyPlanId returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with FrequencyPlanId set to v.
func (EndDeviceOptionNamespace) WithFrequencyPlanId(v string) EndDeviceOption {
	return func(x ttnpb.EndDevice) ttnpb.EndDevice {
		x.FrequencyPlanId = v
		return x
	}
}

// WithMinFrequency returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with MinFrequency set to v.
func (EndDeviceOptionNamespace) WithMinFrequency(v uint64) EndDeviceOption {
	return func(x ttnpb.EndDevice) ttnpb.EndDevice {
		x.MinFrequency = v
		return x
	}
}

// WithMaxFrequency returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with MaxFrequency set to v.
func (EndDeviceOptionNamespace) WithMaxFrequency(v uint64) EndDeviceOption {
	return func(x ttnpb.EndDevice) ttnpb.EndDevice {
		x.MaxFrequency = v
		return x
	}
}

// WithSupportsJoin returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with SupportsJoin set to v.
func (EndDeviceOptionNamespace) WithSupportsJoin(v bool) EndDeviceOption {
	return func(x ttnpb.EndDevice) ttnpb.EndDevice {
		x.SupportsJoin = v
		return x
	}
}

// WithResetsJoinNonces returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with ResetsJoinNonces set to v.
func (EndDeviceOptionNamespace) WithResetsJoinNonces(v bool) EndDeviceOption {
	return func(x ttnpb.EndDevice) ttnpb.EndDevice {
		x.ResetsJoinNonces = v
		return x
	}
}

// WithRootKeys returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with RootKeys set to v.
func (EndDeviceOptionNamespace) WithRootKeys(v *ttnpb.RootKeys) EndDeviceOption {
	return func(x ttnpb.EndDevice) ttnpb.EndDevice {
		x.RootKeys = v
		return x
	}
}

// WithNetId returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with NetId set to v.
func (EndDeviceOptionNamespace) WithNetId(v []byte) EndDeviceOption {
	return func(x ttnpb.EndDevice) ttnpb.EndDevice {
		x.NetId = v
		return x
	}
}

// WithMacSettings returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with MacSettings set to v.
func (EndDeviceOptionNamespace) WithMacSettings(v *ttnpb.MACSettings) EndDeviceOption {
	return func(x ttnpb.EndDevice) ttnpb.EndDevice {
		x.MacSettings = v
		return x
	}
}

// WithMacState returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with MacState set to v.
func (EndDeviceOptionNamespace) WithMacState(v *ttnpb.MACState) EndDeviceOption {
	return func(x ttnpb.EndDevice) ttnpb.EndDevice {
		x.MacState = v
		return x
	}
}

// WithPendingMacState returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with PendingMacState set to v.
func (EndDeviceOptionNamespace) WithPendingMacState(v *ttnpb.MACState) EndDeviceOption {
	return func(x ttnpb.EndDevice) ttnpb.EndDevice {
		x.PendingMacState = v
		return x
	}
}

// WithSession returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with Session set to v.
func (EndDeviceOptionNamespace) WithSession(v *ttnpb.Session) EndDeviceOption {
	return func(x ttnpb.EndDevice) ttnpb.EndDevice {
		x.Session = v
		return x
	}
}

// WithPendingSession returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with PendingSession set to v.
func (EndDeviceOptionNamespace) WithPendingSession(v *ttnpb.Session) EndDeviceOption {
	return func(x ttnpb.EndDevice) ttnpb.EndDevice {
		x.PendingSession = v
		return x
	}
}

// WithLastDevNonce returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with LastDevNonce set to v.
func (EndDeviceOptionNamespace) WithLastDevNonce(v uint32) EndDeviceOption {
	return func(x ttnpb.EndDevice) ttnpb.EndDevice {
		x.LastDevNonce = v
		return x
	}
}

// WithUsedDevNonces returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with UsedDevNonces set to vs.
func (EndDeviceOptionNamespace) WithUsedDevNonces(vs ...uint32) EndDeviceOption {
	return func(x ttnpb.EndDevice) ttnpb.EndDevice {
		x.UsedDevNonces = vs
		return x
	}
}

// WithLastJoinNonce returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with LastJoinNonce set to v.
func (EndDeviceOptionNamespace) WithLastJoinNonce(v uint32) EndDeviceOption {
	return func(x ttnpb.EndDevice) ttnpb.EndDevice {
		x.LastJoinNonce = v
		return x
	}
}

// WithLastRjCount_0 returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with LastRjCount_0 set to v.
func (EndDeviceOptionNamespace) WithLastRjCount_0(v uint32) EndDeviceOption {
	return func(x ttnpb.EndDevice) ttnpb.EndDevice {
		x.LastRjCount_0 = v
		return x
	}
}

// WithLastRjCount_1 returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with LastRjCount_1 set to v.
func (EndDeviceOptionNamespace) WithLastRjCount_1(v uint32) EndDeviceOption {
	return func(x ttnpb.EndDevice) ttnpb.EndDevice {
		x.LastRjCount_1 = v
		return x
	}
}

// WithLastDevStatusReceivedAt returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with LastDevStatusReceivedAt set to v.
func (EndDeviceOptionNamespace) WithLastDevStatusReceivedAt(v *pbtypes.Timestamp) EndDeviceOption {
	return func(x ttnpb.EndDevice) ttnpb.EndDevice {
		x.LastDevStatusReceivedAt = v
		return x
	}
}

// WithPowerState returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with PowerState set to v.
func (EndDeviceOptionNamespace) WithPowerState(v ttnpb.PowerState) EndDeviceOption {
	return func(x ttnpb.EndDevice) ttnpb.EndDevice {
		x.PowerState = v
		return x
	}
}

// WithBatteryPercentage returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with BatteryPercentage set to v.
func (EndDeviceOptionNamespace) WithBatteryPercentage(v *pbtypes.FloatValue) EndDeviceOption {
	return func(x ttnpb.EndDevice) ttnpb.EndDevice {
		x.BatteryPercentage = v
		return x
	}
}

// WithDownlinkMargin returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with DownlinkMargin set to v.
func (EndDeviceOptionNamespace) WithDownlinkMargin(v int32) EndDeviceOption {
	return func(x ttnpb.EndDevice) ttnpb.EndDevice {
		x.DownlinkMargin = v
		return x
	}
}

// WithQueuedApplicationDownlinks returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with QueuedApplicationDownlinks set to vs.
func (EndDeviceOptionNamespace) WithQueuedApplicationDownlinks(vs ...*ttnpb.ApplicationDownlink) EndDeviceOption {
	return func(x ttnpb.EndDevice) ttnpb.EndDevice {
		x.QueuedApplicationDownlinks = vs
		return x
	}
}

// WithFormatters returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with Formatters set to v.
func (EndDeviceOptionNamespace) WithFormatters(v *ttnpb.MessagePayloadFormatters) EndDeviceOption {
	return func(x ttnpb.EndDevice) ttnpb.EndDevice {
		x.Formatters = v
		return x
	}
}

// WithProvisionerId returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with ProvisionerId set to v.
func (EndDeviceOptionNamespace) WithProvisionerId(v string) EndDeviceOption {
	return func(x ttnpb.EndDevice) ttnpb.EndDevice {
		x.ProvisionerId = v
		return x
	}
}

// WithProvisioningData returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with ProvisioningData set to v.
func (EndDeviceOptionNamespace) WithProvisioningData(v *pbtypes.Struct) EndDeviceOption {
	return func(x ttnpb.EndDevice) ttnpb.EndDevice {
		x.ProvisioningData = v
		return x
	}
}

// WithMulticast returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with Multicast set to v.
func (EndDeviceOptionNamespace) WithMulticast(v bool) EndDeviceOption {
	return func(x ttnpb.EndDevice) ttnpb.EndDevice {
		x.Multicast = v
		return x
	}
}

// WithClaimAuthenticationCode returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with ClaimAuthenticationCode set to v.
func (EndDeviceOptionNamespace) WithClaimAuthenticationCode(v *ttnpb.EndDeviceAuthenticationCode) EndDeviceOption {
	return func(x ttnpb.EndDevice) ttnpb.EndDevice {
		x.ClaimAuthenticationCode = v
		return x
	}
}

// WithSkipPayloadCrypto returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with SkipPayloadCrypto set to v.
func (EndDeviceOptionNamespace) WithSkipPayloadCrypto(v bool) EndDeviceOption {
	return func(x ttnpb.EndDevice) ttnpb.EndDevice {
		x.SkipPayloadCrypto = v
		return x
	}
}

// WithSkipPayloadCryptoOverride returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with SkipPayloadCryptoOverride set to v.
func (EndDeviceOptionNamespace) WithSkipPayloadCryptoOverride(v *pbtypes.BoolValue) EndDeviceOption {
	return func(x ttnpb.EndDevice) ttnpb.EndDevice {
		x.SkipPayloadCryptoOverride = v
		return x
	}
}

// WithActivatedAt returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with ActivatedAt set to v.
func (EndDeviceOptionNamespace) WithActivatedAt(v *pbtypes.Timestamp) EndDeviceOption {
	return func(x ttnpb.EndDevice) ttnpb.EndDevice {
		x.ActivatedAt = v
		return x
	}
}

// WithLastSeenAt returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with LastSeenAt set to v.
func (EndDeviceOptionNamespace) WithLastSeenAt(v *pbtypes.Timestamp) EndDeviceOption {
	return func(x ttnpb.EndDevice) ttnpb.EndDevice {
		x.LastSeenAt = v
		return x
	}
}

// WithXXX_unrecognized returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with XXX_unrecognized set to v.
func (EndDeviceOptionNamespace) WithXXX_unrecognized(v []byte) EndDeviceOption {
	return func(x ttnpb.EndDevice) ttnpb.EndDevice {
		x.XXX_unrecognized = v
		return x
	}
}

// Compose returns a functional composition of opts as a singular EndDeviceOption.
func (EndDeviceOptionNamespace) Compose(opts ...EndDeviceOption) EndDeviceOption {
	return func(x ttnpb.EndDevice) ttnpb.EndDevice {
		for _, opt := range opts {
			x = opt(x)
		}
		return x
	}
}

// Compose returns a functional composition of f and opts as a singular EndDeviceOption.
func (f EndDeviceOption) Compose(opts ...EndDeviceOption) EndDeviceOption {
	return func(x ttnpb.EndDevice) ttnpb.EndDevice {
		x = f(x)
		for _, opt := range opts {
			x = opt(x)
		}
		return x
	}
}

// EndDeviceOptions is namespace containing ttnpb.EndDevice options.
var EndDeviceOptions EndDeviceOptionNamespace

// MakeEndDevice constructs a new ttnpb.EndDevice.
func MakeEndDevice(opts ...EndDeviceOption) *ttnpb.EndDevice {
	v := EndDeviceOptions.Compose(opts...)(baseEndDevice)
	return &v
}
