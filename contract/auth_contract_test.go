package contract

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

func TestNewValidateTokenRequest(t *testing.T) {
	const dummyToken = "token"
	request := ValidateTokenRequestContract{Token: dummyToken}
	reqBytes, _ := json.Marshal(request)
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name    string
		args    args
		want    *ValidateTokenRequestContract
		wantErr bool
	}{
		{
			name: "given correct request should return no error",
			args: args{r: &http.Request{
				Body: ioutil.NopCloser(bytes.NewBuffer(reqBytes)),
			}},
			want:    &ValidateTokenRequestContract{Token: dummyToken},
			wantErr: false,
		},
		{
			name: "given incorrect request should return error",
			args: args{r: &http.Request{
				Body: ioutil.NopCloser(bytes.NewBuffer([]byte("jfndjfn"))),
			}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "given invalid request should return error",
			args: args{r: &http.Request{
				Body: ioutil.NopCloser(bytes.NewBuffer([]byte(""))),
			}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewValidateTokenRequest(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewValidateTokenRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewValidateTokenRequest() got = %v, want %v", got, tt.want)
			}
		})
	}
}
