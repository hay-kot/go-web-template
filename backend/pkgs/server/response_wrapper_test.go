package server

import (
	"reflect"
	"testing"
)

func TestWrap(t *testing.T) {
	type args struct {
		namespace string
		data      interface{}
	}
	tests := []struct {
		name string
		args args
		want Wrapper
	}{
		{
			name: "check_basic_wrap",
			args: args{
				namespace: "test",
				data:      "test",
			},
			want: map[string]interface{}{
				"test": "test",
			},
		},
		{
			name: "check_basic_interface_wrap",
			args: args{
				namespace: "namespace",
				data: struct {
					Name string
					Data string
				}{
					Name: "Test Name",
					Data: "Test Data",
				},
			},
			want: map[string]interface{}{
				"namespace": struct {
					Name string
					Data string
				}{
					Name: "Test Name",
					Data: "Test Data",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Wrap(tt.args.namespace, tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Wrap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWrapper_Add(t *testing.T) {
	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name string
		w    Wrapper
		args args
		want Wrapper
	}{
		{
			name: "check_basic_add",
			w:    Wrap("data", "default"),
			args: args{
				key:   "test",
				value: "test",
			},
			want: map[string]interface{}{
				"data": "default",
				"test": "test",
			},
		},
		{
			name: "check_add_details",
			w:    Wrap("data", "default"),
			args: args{
				key:   "details",
				value: "test-details",
			},
			want: map[string]interface{}{
				"data":    "default",
				"details": "test-details",
			},
		},
		{
			name: "check_add_message",
			w:    Wrap("data", "default"),
			args: args{
				key:   "message",
				value: "test-message",
			},
			want: map[string]interface{}{
				"data":    "default",
				"message": "test-message",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.w.Add(tt.args.key, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Wrapper.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
