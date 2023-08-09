// Code generated by protoc-gen-go-json. DO NOT EDIT.
// versions:
// - protoc-gen-go-json v1.5.1
// - protoc             v4.22.2
// source: ttn/lorawan/v3/gateway_services.proto

package ttnpb

import (
	golang "github.com/TheThingsIndustries/protoc-gen-go-json/golang"
	jsonplugin "github.com/TheThingsIndustries/protoc-gen-go-json/jsonplugin"
)

// MarshalProtoJSON marshals the PullGatewayConfigurationRequest message to JSON.
func (x *PullGatewayConfigurationRequest) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.GatewayIds != nil || s.HasField("gateway_ids") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("gateway_ids")
		x.GatewayIds.MarshalProtoJSON(s.WithField("gateway_ids"))
	}
	if x.FieldMask != nil || s.HasField("field_mask") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("field_mask")
		if x.FieldMask == nil {
			s.WriteNil()
		} else {
			golang.MarshalLegacyFieldMask(s, x.FieldMask)
		}
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the PullGatewayConfigurationRequest to JSON.
func (x *PullGatewayConfigurationRequest) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the PullGatewayConfigurationRequest message from JSON.
func (x *PullGatewayConfigurationRequest) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "gateway_ids", "gatewayIds":
			if s.ReadNil() {
				x.GatewayIds = nil
				return
			}
			x.GatewayIds = &GatewayIdentifiers{}
			x.GatewayIds.UnmarshalProtoJSON(s.WithField("gateway_ids", true))
		case "field_mask", "fieldMask":
			s.AddField("field_mask")
			if s.ReadNil() {
				x.FieldMask = nil
				return
			}
			v := golang.UnmarshalFieldMask(s)
			if s.Err() != nil {
				return
			}
			x.FieldMask = v
		}
	})
}

// UnmarshalJSON unmarshals the PullGatewayConfigurationRequest from JSON.
func (x *PullGatewayConfigurationRequest) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the AssertGatewayRightsRequest message to JSON.
func (x *AssertGatewayRightsRequest) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if len(x.GatewayIds) > 0 || s.HasField("gateway_ids") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("gateway_ids")
		s.WriteArrayStart()
		var wroteElement bool
		for _, element := range x.GatewayIds {
			s.WriteMoreIf(&wroteElement)
			element.MarshalProtoJSON(s.WithField("gateway_ids"))
		}
		s.WriteArrayEnd()
	}
	if x.RequiredRights != nil || s.HasField("required_rights") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("required_rights")
		x.RequiredRights.MarshalProtoJSON(s.WithField("required_rights"))
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the AssertGatewayRightsRequest to JSON.
func (x *AssertGatewayRightsRequest) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the AssertGatewayRightsRequest message from JSON.
func (x *AssertGatewayRightsRequest) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "gateway_ids", "gatewayIds":
			s.AddField("gateway_ids")
			if s.ReadNil() {
				x.GatewayIds = nil
				return
			}
			s.ReadArray(func() {
				if s.ReadNil() {
					x.GatewayIds = append(x.GatewayIds, nil)
					return
				}
				v := &GatewayIdentifiers{}
				v.UnmarshalProtoJSON(s.WithField("gateway_ids", false))
				if s.Err() != nil {
					return
				}
				x.GatewayIds = append(x.GatewayIds, v)
			})
		case "required_rights", "requiredRights":
			if s.ReadNil() {
				x.RequiredRights = nil
				return
			}
			x.RequiredRights = &Rights{}
			x.RequiredRights.UnmarshalProtoJSON(s.WithField("required_rights", true))
		}
	})
}

// UnmarshalJSON unmarshals the AssertGatewayRightsRequest from JSON.
func (x *AssertGatewayRightsRequest) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the AssertGatewayRightsResponse message to JSON.
func (x *AssertGatewayRightsResponse) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if len(x.GatewayIds) > 0 || s.HasField("gateway_ids") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("gateway_ids")
		s.WriteArrayStart()
		var wroteElement bool
		for _, element := range x.GatewayIds {
			s.WriteMoreIf(&wroteElement)
			element.MarshalProtoJSON(s.WithField("gateway_ids"))
		}
		s.WriteArrayEnd()
	}
	if x.HasRequiredRights || s.HasField("has_required_rights") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("has_required_rights")
		s.WriteBool(x.HasRequiredRights)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the AssertGatewayRightsResponse to JSON.
func (x *AssertGatewayRightsResponse) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the AssertGatewayRightsResponse message from JSON.
func (x *AssertGatewayRightsResponse) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "gateway_ids", "gatewayIds":
			s.AddField("gateway_ids")
			if s.ReadNil() {
				x.GatewayIds = nil
				return
			}
			s.ReadArray(func() {
				if s.ReadNil() {
					x.GatewayIds = append(x.GatewayIds, nil)
					return
				}
				v := &GatewayIdentifiers{}
				v.UnmarshalProtoJSON(s.WithField("gateway_ids", false))
				if s.Err() != nil {
					return
				}
				x.GatewayIds = append(x.GatewayIds, v)
			})
		case "has_required_rights", "hasRequiredRights":
			s.AddField("has_required_rights")
			x.HasRequiredRights = s.ReadBool()
		}
	})
}

// UnmarshalJSON unmarshals the AssertGatewayRightsResponse from JSON.
func (x *AssertGatewayRightsResponse) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}
