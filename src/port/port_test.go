package port

import (
	"context"
	"reflect"
	"testing"

	"portdb.io/src/models"
	"portdb.io/src/proto/grpc"
)

func Test_service_Fetch(t *testing.T) {
	type fields struct {
		repo                                 models.PortDomainRepository
		UnimplementedPortDomainServiceServer grpc.UnimplementedPortDomainServiceServer
	}
	type args struct {
		ctx     context.Context
		request *grpc.FetchRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *grpc.Port
		wantErr bool
	}{
		{
			name: "fetch port",
			fields: fields{
				repo: &testRepo{
					port: &models.Port{
						Code:     "test_code",
						Name:     "test_name",
						City:     "test_city",
						Country:  "test_country",
						Province: "test_province",
					},
				},
			},
			args: args{
				ctx: nil,
				request: &grpc.FetchRequest{
					Code: "test_code",
				},
			},
			want: &grpc.Port{
				Code:     "test_code",
				Name:     "test_name",
				City:     "test_city",
				Country:  "test_country",
				Province: "test_province",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo:                                 tt.fields.repo,
				UnimplementedPortDomainServiceServer: tt.fields.UnimplementedPortDomainServiceServer,
			}
			got, err := s.Fetch(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Fetch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Fetch() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_Store(t *testing.T) {
	type fields struct {
		repo                                 models.PortDomainRepository
		UnimplementedPortDomainServiceServer grpc.UnimplementedPortDomainServiceServer
	}
	type args struct {
		ctx     context.Context
		request *grpc.CreateRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *grpc.Port
		wantErr bool
	}{
		{
			name: "store port",
			fields: fields{
				repo: &testRepo{},
			},
			args: args{
				ctx: nil,
				request: &grpc.CreateRequest{
					Code:     "test_code",
					Name:     "test_name",
					City:     "test_city",
					Country:  "test_country",
					Province: "test_province",
				},
			},
			want: &grpc.Port{
				Code:     "test_code",
				Name:     "test_name",
				City:     "test_city",
				Country:  "test_country",
				Province: "test_province",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo:                                 tt.fields.repo,
				UnimplementedPortDomainServiceServer: tt.fields.UnimplementedPortDomainServiceServer,
			}
			got, err := s.Store(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Store() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Store() got = %v, want %v", got, tt.want)
			}
		})
	}
}
