// Code generated by thriftrw v0.4.0
// @generated

// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package keyvalue

import (
	"errors"
	"fmt"
	"go.uber.org/thriftrw/wire"
	"go.uber.org/yarpc/examples/thrift/keyvalue/kv"
	"strings"
)

type GetValueArgs struct {
	Key *string `json:"key,omitempty"`
}

func (v *GetValueArgs) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Key != nil {
		w, err = wire.NewValueString(*(v.Key)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *GetValueArgs) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.Key = &x
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (v *GetValueArgs) String() string {
	var fields [1]string
	i := 0
	if v.Key != nil {
		fields[i] = fmt.Sprintf("Key: %v", *(v.Key))
		i++
	}
	return fmt.Sprintf("GetValueArgs{%v}", strings.Join(fields[:i], ", "))
}

func (v *GetValueArgs) MethodName() string {
	return "getValue"
}

func (v *GetValueArgs) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

var GetValueHelper = struct {
	Args           func(key *string) *GetValueArgs
	IsException    func(error) bool
	WrapResponse   func(string, error) (*GetValueResult, error)
	UnwrapResponse func(*GetValueResult) (string, error)
}{}

func init() {
	GetValueHelper.Args = func(key *string) *GetValueArgs {
		return &GetValueArgs{Key: key}
	}
	GetValueHelper.IsException = func(err error) bool {
		switch err.(type) {
		case *kv.ResourceDoesNotExist:
			return true
		default:
			return false
		}
	}
	GetValueHelper.WrapResponse = func(success string, err error) (*GetValueResult, error) {
		if err == nil {
			return &GetValueResult{Success: &success}, nil
		}
		switch e := err.(type) {
		case *kv.ResourceDoesNotExist:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for GetValueResult.DoesNotExist")
			}
			return &GetValueResult{DoesNotExist: e}, nil
		}
		return nil, err
	}
	GetValueHelper.UnwrapResponse = func(result *GetValueResult) (success string, err error) {
		if result.DoesNotExist != nil {
			err = result.DoesNotExist
			return
		}
		if result.Success != nil {
			success = *result.Success
			return
		}
		err = errors.New("expected a non-void result")
		return
	}
}

type GetValueResult struct {
	Success      *string                  `json:"success,omitempty"`
	DoesNotExist *kv.ResourceDoesNotExist `json:"doesNotExist,omitempty"`
}

func (v *GetValueResult) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Success != nil {
		w, err = wire.NewValueString(*(v.Success)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}
	if v.DoesNotExist != nil {
		w, err = v.DoesNotExist.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	if i != 1 {
		return wire.Value{}, fmt.Errorf("GetValueResult should have exactly one field: got %v fields", i)
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _ResourceDoesNotExist_Read(w wire.Value) (*kv.ResourceDoesNotExist, error) {
	var v kv.ResourceDoesNotExist
	err := v.FromWire(w)
	return &v, err
}

func (v *GetValueResult) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.Success = &x
				if err != nil {
					return err
				}
			}
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.DoesNotExist, err = _ResourceDoesNotExist_Read(field.Value)
				if err != nil {
					return err
				}
			}
		}
	}
	count := 0
	if v.Success != nil {
		count++
	}
	if v.DoesNotExist != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("GetValueResult should have exactly one field: got %v fields", count)
	}
	return nil
}

func (v *GetValueResult) String() string {
	var fields [2]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", *(v.Success))
		i++
	}
	if v.DoesNotExist != nil {
		fields[i] = fmt.Sprintf("DoesNotExist: %v", v.DoesNotExist)
		i++
	}
	return fmt.Sprintf("GetValueResult{%v}", strings.Join(fields[:i], ", "))
}

func (v *GetValueResult) MethodName() string {
	return "getValue"
}

func (v *GetValueResult) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
