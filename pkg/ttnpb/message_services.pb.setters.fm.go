// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

import fmt "fmt"

func (dst *EncodeDownlinkMessageRequest) SetFields(src *EncodeDownlinkMessageRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "ids":
			if len(subs) > 0 {
				var newDst, newSrc *EndDeviceIdentifiers
				if src != nil {
					newSrc = &src.EndDeviceIdentifiers
				}
				newDst = &dst.EndDeviceIdentifiers
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.EndDeviceIdentifiers = src.EndDeviceIdentifiers
				} else {
					var zero EndDeviceIdentifiers
					dst.EndDeviceIdentifiers = zero
				}
			}
		case "end_device_version_ids":
			if len(subs) > 0 {
				var newDst, newSrc *EndDeviceVersionIdentifiers
				if src != nil {
					newSrc = &src.EndDeviceVersionIds
				}
				newDst = &dst.EndDeviceVersionIds
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.EndDeviceVersionIds = src.EndDeviceVersionIds
				} else {
					var zero EndDeviceVersionIdentifiers
					dst.EndDeviceVersionIds = zero
				}
			}
		case "message":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationDownlink
				if src != nil {
					newSrc = &src.Message
				}
				newDst = &dst.Message
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Message = src.Message
				} else {
					var zero ApplicationDownlink
					dst.Message = zero
				}
			}
		case "formatter":
			if len(subs) > 0 {
				return fmt.Errorf("'formatter' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Formatter = src.Formatter
			} else {
				var zero PayloadFormatter
				dst.Formatter = zero
			}
		case "parameter":
			if len(subs) > 0 {
				return fmt.Errorf("'parameter' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Parameter = src.Parameter
			} else {
				var zero string
				dst.Parameter = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *DecodeUplinkMessageRequest) SetFields(src *DecodeUplinkMessageRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "ids":
			if len(subs) > 0 {
				var newDst, newSrc *EndDeviceIdentifiers
				if src != nil {
					newSrc = &src.EndDeviceIdentifiers
				}
				newDst = &dst.EndDeviceIdentifiers
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.EndDeviceIdentifiers = src.EndDeviceIdentifiers
				} else {
					var zero EndDeviceIdentifiers
					dst.EndDeviceIdentifiers = zero
				}
			}
		case "end_device_version_ids":
			if len(subs) > 0 {
				var newDst, newSrc *EndDeviceVersionIdentifiers
				if src != nil {
					newSrc = &src.EndDeviceVersionIds
				}
				newDst = &dst.EndDeviceVersionIds
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.EndDeviceVersionIds = src.EndDeviceVersionIds
				} else {
					var zero EndDeviceVersionIdentifiers
					dst.EndDeviceVersionIds = zero
				}
			}
		case "message":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationUplink
				if src != nil {
					newSrc = &src.Message
				}
				newDst = &dst.Message
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Message = src.Message
				} else {
					var zero ApplicationUplink
					dst.Message = zero
				}
			}
		case "formatter":
			if len(subs) > 0 {
				return fmt.Errorf("'formatter' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Formatter = src.Formatter
			} else {
				var zero PayloadFormatter
				dst.Formatter = zero
			}
		case "parameter":
			if len(subs) > 0 {
				return fmt.Errorf("'parameter' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Parameter = src.Parameter
			} else {
				var zero string
				dst.Parameter = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *DecodeDownlinkMessageRequest) SetFields(src *DecodeDownlinkMessageRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "ids":
			if len(subs) > 0 {
				var newDst, newSrc *EndDeviceIdentifiers
				if src != nil {
					newSrc = &src.EndDeviceIdentifiers
				}
				newDst = &dst.EndDeviceIdentifiers
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.EndDeviceIdentifiers = src.EndDeviceIdentifiers
				} else {
					var zero EndDeviceIdentifiers
					dst.EndDeviceIdentifiers = zero
				}
			}
		case "end_device_version_ids":
			if len(subs) > 0 {
				var newDst, newSrc *EndDeviceVersionIdentifiers
				if src != nil {
					newSrc = &src.EndDeviceVersionIds
				}
				newDst = &dst.EndDeviceVersionIds
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.EndDeviceVersionIds = src.EndDeviceVersionIds
				} else {
					var zero EndDeviceVersionIdentifiers
					dst.EndDeviceVersionIds = zero
				}
			}
		case "message":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationDownlink
				if src != nil {
					newSrc = &src.Message
				}
				newDst = &dst.Message
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Message = src.Message
				} else {
					var zero ApplicationDownlink
					dst.Message = zero
				}
			}
		case "formatter":
			if len(subs) > 0 {
				return fmt.Errorf("'formatter' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Formatter = src.Formatter
			} else {
				var zero PayloadFormatter
				dst.Formatter = zero
			}
		case "parameter":
			if len(subs) > 0 {
				return fmt.Errorf("'parameter' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Parameter = src.Parameter
			} else {
				var zero string
				dst.Parameter = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}
