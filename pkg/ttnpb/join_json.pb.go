// Code generated by protoc-gen-go-json. DO NOT EDIT.
// versions:
// - protoc-gen-go-json v1.4.0
// - protoc             v3.9.1
// source: lorawan-stack/api/join.proto

package ttnpb

import (
	gogo "github.com/TheThingsIndustries/protoc-gen-go-json/gogo"
	jsonplugin "github.com/TheThingsIndustries/protoc-gen-go-json/jsonplugin"
	types "go.thethings.network/lorawan-stack/v3/pkg/types"
)

// MarshalProtoJSON marshals the JoinRequest message to JSON.
func (x *JoinRequest) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if len(x.RawPayload) > 0 || s.HasField("raw_payload") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("raw_payload")
		s.WriteBytes(x.RawPayload)
	}
	if x.Payload != nil || s.HasField("payload") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("payload")
		x.Payload.MarshalProtoJSON(s.WithField("payload"))
	}
	if len(x.DevAddr) > 0 || s.HasField("dev_addr") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("dev_addr")
		types.MarshalHEXBytes(s.WithField("dev_addr"), x.DevAddr)
	}
	if x.SelectedMacVersion != 0 || s.HasField("selected_mac_version") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("selected_mac_version")
		x.SelectedMacVersion.MarshalProtoJSON(s)
	}
	if len(x.NetId) > 0 || s.HasField("net_id") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("net_id")
		types.MarshalHEXBytes(s.WithField("net_id"), x.NetId)
	}
	if x.DownlinkSettings != nil || s.HasField("downlink_settings") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("downlink_settings")
		x.DownlinkSettings.MarshalProtoJSON(s.WithField("downlink_settings"))
	}
	if x.RxDelay != 0 || s.HasField("rx_delay") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("rx_delay")
		x.RxDelay.MarshalProtoJSON(s)
	}
	if x.CfList != nil || s.HasField("cf_list") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("cf_list")
		x.CfList.MarshalProtoJSON(s.WithField("cf_list"))
	}
	if len(x.CorrelationIds) > 0 || s.HasField("correlation_ids") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("correlation_ids")
		s.WriteStringArray(x.CorrelationIds)
	}
	if x.ConsumedAirtime != nil || s.HasField("consumed_airtime") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("consumed_airtime")
		if x.ConsumedAirtime == nil {
			s.WriteNil()
		} else {
			gogo.MarshalDuration(s, x.ConsumedAirtime)
		}
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the JoinRequest to JSON.
func (x *JoinRequest) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the JoinRequest message from JSON.
func (x *JoinRequest) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "raw_payload", "rawPayload":
			s.AddField("raw_payload")
			x.RawPayload = s.ReadBytes()
		case "payload":
			if s.ReadNil() {
				x.Payload = nil
				return
			}
			x.Payload = &Message{}
			x.Payload.UnmarshalProtoJSON(s.WithField("payload", true))
		case "dev_addr", "devAddr":
			s.AddField("dev_addr")
			x.DevAddr = types.Unmarshal4Bytes(s.WithField("dev_addr", false))
		case "selected_mac_version", "selectedMacVersion":
			s.AddField("selected_mac_version")
			x.SelectedMacVersion.UnmarshalProtoJSON(s)
		case "net_id", "netId":
			s.AddField("net_id")
			x.NetId = types.Unmarshal3Bytes(s.WithField("net_id", false))
		case "downlink_settings", "downlinkSettings":
			if s.ReadNil() {
				x.DownlinkSettings = nil
				return
			}
			x.DownlinkSettings = &DLSettings{}
			x.DownlinkSettings.UnmarshalProtoJSON(s.WithField("downlink_settings", true))
		case "rx_delay", "rxDelay":
			s.AddField("rx_delay")
			x.RxDelay.UnmarshalProtoJSON(s)
		case "cf_list", "cfList":
			if s.ReadNil() {
				x.CfList = nil
				return
			}
			x.CfList = &CFList{}
			x.CfList.UnmarshalProtoJSON(s.WithField("cf_list", true))
		case "correlation_ids", "correlationIds":
			s.AddField("correlation_ids")
			if s.ReadNil() {
				x.CorrelationIds = nil
				return
			}
			x.CorrelationIds = s.ReadStringArray()
		case "consumed_airtime", "consumedAirtime":
			s.AddField("consumed_airtime")
			if s.ReadNil() {
				x.ConsumedAirtime = nil
				return
			}
			v := gogo.UnmarshalDuration(s)
			if s.Err() != nil {
				return
			}
			x.ConsumedAirtime = v
		}
	})
}

// UnmarshalJSON unmarshals the JoinRequest from JSON.
func (x *JoinRequest) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the JoinResponse message to JSON.
func (x *JoinResponse) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if len(x.RawPayload) > 0 || s.HasField("raw_payload") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("raw_payload")
		s.WriteBytes(x.RawPayload)
	}
	if x.SessionKeys != nil || s.HasField("session_keys") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("session_keys")
		x.SessionKeys.MarshalProtoJSON(s.WithField("session_keys"))
	}
	if x.Lifetime != nil || s.HasField("lifetime") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("lifetime")
		if x.Lifetime == nil {
			s.WriteNil()
		} else {
			gogo.MarshalDuration(s, x.Lifetime)
		}
	}
	if len(x.CorrelationIds) > 0 || s.HasField("correlation_ids") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("correlation_ids")
		s.WriteStringArray(x.CorrelationIds)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the JoinResponse to JSON.
func (x *JoinResponse) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the JoinResponse message from JSON.
func (x *JoinResponse) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "raw_payload", "rawPayload":
			s.AddField("raw_payload")
			x.RawPayload = s.ReadBytes()
		case "session_keys", "sessionKeys":
			if s.ReadNil() {
				x.SessionKeys = nil
				return
			}
			x.SessionKeys = &SessionKeys{}
			x.SessionKeys.UnmarshalProtoJSON(s.WithField("session_keys", true))
		case "lifetime":
			s.AddField("lifetime")
			if s.ReadNil() {
				x.Lifetime = nil
				return
			}
			v := gogo.UnmarshalDuration(s)
			if s.Err() != nil {
				return
			}
			x.Lifetime = v
		case "correlation_ids", "correlationIds":
			s.AddField("correlation_ids")
			if s.ReadNil() {
				x.CorrelationIds = nil
				return
			}
			x.CorrelationIds = s.ReadStringArray()
		}
	})
}

// UnmarshalJSON unmarshals the JoinResponse from JSON.
func (x *JoinResponse) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}
