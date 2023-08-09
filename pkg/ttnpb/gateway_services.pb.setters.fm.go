// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

import fmt "fmt"

func (dst *PullGatewayConfigurationRequest) SetFields(src *PullGatewayConfigurationRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "gateway_ids":
			if len(subs) > 0 {
				var newDst, newSrc *GatewayIdentifiers
				if (src == nil || src.GatewayIds == nil) && dst.GatewayIds == nil {
					continue
				}
				if src != nil {
					newSrc = src.GatewayIds
				}
				if dst.GatewayIds != nil {
					newDst = dst.GatewayIds
				} else {
					newDst = &GatewayIdentifiers{}
					dst.GatewayIds = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.GatewayIds = src.GatewayIds
				} else {
					dst.GatewayIds = nil
				}
			}
		case "field_mask":
			if len(subs) > 0 {
				return fmt.Errorf("'field_mask' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FieldMask = src.FieldMask
			} else {
				dst.FieldMask = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *AssertGatewayRightsRequest) SetFields(src *AssertGatewayRightsRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "gateway_ids":
			if len(subs) > 0 {
				return fmt.Errorf("'gateway_ids' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.GatewayIds = src.GatewayIds
			} else {
				dst.GatewayIds = nil
			}
		case "required_rights":
			if len(subs) > 0 {
				var newDst, newSrc *Rights
				if (src == nil || src.RequiredRights == nil) && dst.RequiredRights == nil {
					continue
				}
				if src != nil {
					newSrc = src.RequiredRights
				}
				if dst.RequiredRights != nil {
					newDst = dst.RequiredRights
				} else {
					newDst = &Rights{}
					dst.RequiredRights = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.RequiredRights = src.RequiredRights
				} else {
					dst.RequiredRights = nil
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *AssertGatewayRightsResponse) SetFields(src *AssertGatewayRightsResponse, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "gateway_ids":
			if len(subs) > 0 {
				return fmt.Errorf("'gateway_ids' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.GatewayIds = src.GatewayIds
			} else {
				dst.GatewayIds = nil
			}
		case "has_required_rights":
			if len(subs) > 0 {
				return fmt.Errorf("'has_required_rights' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.HasRequiredRights = src.HasRequiredRights
			} else {
				var zero bool
				dst.HasRequiredRights = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}
