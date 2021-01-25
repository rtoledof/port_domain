package port

import (
	"reflect"
	"testing"

	"portdb.io/src/proto/grpc"
)

func TestCreateRequest_Valid(t *testing.T) {
	type fields struct {
		Name        string
		City        string
		Country     string
		Alias       []string
		Regions     []string
		Coordinates []float32
		Province    string
		Timezone    string
		Unlocs      []string
		Code        string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &CreateRequest{
				Name:        tt.fields.Name,
				City:        tt.fields.City,
				Country:     tt.fields.Country,
				Alias:       tt.fields.Alias,
				Regions:     tt.fields.Regions,
				Coordinates: tt.fields.Coordinates,
				Province:    tt.fields.Province,
				Timezone:    tt.fields.Timezone,
				Unlocs:      tt.fields.Unlocs,
				Code:        tt.fields.Code,
			}
			if err := r.Valid(); (err != nil) != tt.wantErr {
				t.Errorf("Valid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUnmarshal(t *testing.T) {
	type args struct {
		pb *grpc.CreateRequest
	}
	tests := []struct {
		name string
		args args
		want *CreateRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Unmarshal(tt.args.pb); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unmarshal() = %v, want %v", got, tt.want)
			}
		})
	}
}
