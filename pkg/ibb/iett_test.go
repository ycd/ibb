package ibb

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

func TestClient_IETTTicketPrices(t *testing.T) {
	type fields struct {
		c     *http.Client
		debug bool
	}
	type args struct {
		ctx  context.Context
		year int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *TicketPrices
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
			got, err := c.IETTTicketPrices(tt.args.ctx, tt.args.year)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.IETTTicketPrices() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.IETTTicketPrices() = %v, want %v", got, tt.want)
			}
		})
	}
}
