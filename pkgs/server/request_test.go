package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type TestStruct struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func TestDecode(t *testing.T) {
	type args struct {
		r   *http.Request
		val interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "check_error",
			args: args{
				r: &http.Request{
					Body: http.NoBody,
				},
				val: make(map[string]interface{}),
			},
			wantErr: true,
		},
		{
			name: "check_success",
			args: args{
				r: httptest.NewRequest("POST", "/", strings.NewReader(`{"name":"test","data":"test"}`)),
				val: TestStruct{
					Name: "test",
					Data: "test",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Decode(tt.args.r, &tt.args.val); (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
