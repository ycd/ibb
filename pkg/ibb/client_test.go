package ibb

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

func TestClient_get(t *testing.T) {
	type fields struct {
		c     *http.Client
		debug bool
	}
	type args struct {
		ctx      context.Context
		resource string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				c:     tt.fields.c,
				debug: tt.fields.debug,
			}
			got, err := c.get(tt.args.ctx, tt.args.resource)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.get() = %v, want %v", got, tt.want)
			}
		})
	}
}
