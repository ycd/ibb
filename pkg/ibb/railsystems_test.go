package ibb

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

func TestClient_DailyMaximumJourneyByRailSystem(t *testing.T) {
	type fields struct {
		c     *http.Client
		debug bool
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *RailSystemsDailyMaximumJourneys
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
			got, err := c.DailyMaximumJourneyByRailSystem(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.DailyMaximumJourneyByRailSystem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.DailyMaximumJourneyByRailSystem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_RailSystemsTimeline(t *testing.T) {
	type fields struct {
		c     *http.Client
		debug bool
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *RailSystemsTimeline
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
			got, err := c.RailSystemsTimeline(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.RailSystemsTimeline() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.RailSystemsTimeline() = %v, want %v", got, tt.want)
			}
		})
	}
}
