// Code generated by protoc-gen-go-form. DO NOT EDIT.
// Source: metric.proto

package pb

import (
	json "encoding/json"
	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
	structpb "google.golang.org/protobuf/types/known/structpb"
	url "net/url"
	strconv "strconv"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/urlenc" package it is being compiled against.
var _ urlenc.URLValuesUnmarshaler = (*QueryWithInfluxFormatRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*QueryWithInfluxFormatResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*Serie)(nil)
var _ urlenc.URLValuesUnmarshaler = (*Row)(nil)
var _ urlenc.URLValuesUnmarshaler = (*Result)(nil)
var _ urlenc.URLValuesUnmarshaler = (*QueryWithTableFormatRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*QueryWithTableFormatResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*TableResult)(nil)
var _ urlenc.URLValuesUnmarshaler = (*TableColumn)(nil)
var _ urlenc.URLValuesUnmarshaler = (*TableRow)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GeneralQueryRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GeneralQueryResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*Filter)(nil)

// QueryWithInfluxFormatRequest implement urlenc.URLValuesUnmarshaler.
func (m *QueryWithInfluxFormatRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "start":
				m.Start = vals[0]
			case "end":
				m.End = vals[0]
			case "statement":
				m.Statement = vals[0]
			}
		}
	}
	return nil
}

// QueryWithInfluxFormatResponse implement urlenc.URLValuesUnmarshaler.
func (m *QueryWithInfluxFormatResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// Serie implement urlenc.URLValuesUnmarshaler.
func (m *Serie) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "name":
				m.Name = vals[0]
			case "columns":
				m.Columns = vals
			}
		}
	}
	return nil
}

// Row implement urlenc.URLValuesUnmarshaler.
func (m *Row) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// Result implement urlenc.URLValuesUnmarshaler.
func (m *Result) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "statement_id":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.StatementId = val
			}
		}
	}
	return nil
}

// QueryWithTableFormatRequest implement urlenc.URLValuesUnmarshaler.
func (m *QueryWithTableFormatRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "start":
				m.Start = vals[0]
			case "end":
				m.End = vals[0]
			case "statement":
				m.Statement = vals[0]
			}
		}
	}
	return nil
}

// QueryWithTableFormatResponse implement urlenc.URLValuesUnmarshaler.
func (m *QueryWithTableFormatResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &TableResult{}
				}
			case "data.interval":
				if m.Data == nil {
					m.Data = &TableResult{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.Interval = val
			}
		}
	}
	return nil
}

// TableResult implement urlenc.URLValuesUnmarshaler.
func (m *TableResult) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "interval":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Interval = val
			}
		}
	}
	return nil
}

// TableColumn implement urlenc.URLValuesUnmarshaler.
func (m *TableColumn) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "key":
				m.Key = vals[0]
			case "name":
				m.Name = vals[0]
			case "flag":
				m.Flag = vals[0]
			}
		}
	}
	return nil
}

// TableRow implement urlenc.URLValuesUnmarshaler.
func (m *TableRow) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// GeneralQueryRequest implement urlenc.URLValuesUnmarshaler.
func (m *GeneralQueryRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "ql":
				m.Ql = vals[0]
			case "statement":
				m.Statement = vals[0]
			case "format":
				m.Format = vals[0]
			}
		}
	}
	return nil
}

// GeneralQueryResponse implement urlenc.URLValuesUnmarshaler.
func (m *GeneralQueryResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if len(vals) > 1 {
					var list []interface{}
					for _, text := range vals {
						var v interface{}
						err := json.NewDecoder(strings.NewReader(text)).Decode(&v)
						if err != nil {
							list = append(list, v)
						} else {
							list = append(list, text)
						}
					}
					val, _ := structpb.NewList(list)
					m.Data = structpb.NewListValue(val)
				} else {
					var v interface{}
					err := json.NewDecoder(strings.NewReader(vals[0])).Decode(&v)
					if err != nil {
						val, _ := structpb.NewValue(v)
						m.Data = val
					} else {
						m.Data = structpb.NewStringValue(vals[0])
					}
				}
			}
		}
	}
	return nil
}

// Filter implement urlenc.URLValuesUnmarshaler.
func (m *Filter) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "key":
				m.Key = vals[0]
			case "op":
				m.Op = vals[0]
			case "value":
				if len(vals) > 1 {
					var list []interface{}
					for _, text := range vals {
						var v interface{}
						err := json.NewDecoder(strings.NewReader(text)).Decode(&v)
						if err != nil {
							list = append(list, v)
						} else {
							list = append(list, text)
						}
					}
					val, _ := structpb.NewList(list)
					m.Value = structpb.NewListValue(val)
				} else {
					var v interface{}
					err := json.NewDecoder(strings.NewReader(vals[0])).Decode(&v)
					if err != nil {
						val, _ := structpb.NewValue(v)
						m.Value = val
					} else {
						m.Value = structpb.NewStringValue(vals[0])
					}
				}
			}
		}
	}
	return nil
}
