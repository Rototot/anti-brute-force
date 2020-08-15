package factories

import (
	"github.com/Rototot/anti-brute-force/pkg/domain/constants"
	"github.com/Rototot/anti-brute-force/pkg/domain/entities"
	"github.com/Rototot/anti-brute-force/pkg/domain/factories/mocks"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func TestBucketFactory_Create(t *testing.T) {
	const expectedIpCapacity = 1000
	const expectedLoginCapacity = 10
	const expectedPasswordCapacity = 100

	type fields struct {
		conf BucketConfigurator
	}
	type args struct {
		bType constants.BucketType
	}

	configurator := mocks.NewMockBucketConfigurator(gomock.NewController(t))
	configurator.EXPECT().IpBucketCapacity().Return(expectedIpCapacity).AnyTimes()
	configurator.EXPECT().LoginBucketCapacity().Return(expectedLoginCapacity).AnyTimes()
	configurator.EXPECT().PasswordBucketCapacity().Return(expectedPasswordCapacity).AnyTimes()

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.Bucket
		wantErr bool
	}{
		{
			name:   "",
			fields: fields{conf: configurator},
			args:   args{bType: constants.BucketTypeIp},
			want: &entities.Bucket{
				Capacity: expectedIpCapacity,
			},
		},
		{
			name:   "",
			fields: fields{conf: configurator},
			args:   args{bType: constants.BucketTypeLogin},
			want: &entities.Bucket{
				Capacity: expectedLoginCapacity,
			},
		},
		{
			name:   "",
			fields: fields{conf: configurator},
			args:   args{bType: constants.BucketTypePassword},
			want: &entities.Bucket{
				Capacity: expectedPasswordCapacity,
			},
		},
		{
			name:    "when unknown bucket type",
			fields:  fields{conf: configurator},
			args:    args{bType: constants.BucketType(100000000)},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &BucketFactory{
				conf: tt.fields.conf,
			}
			got, err := f.Create(tt.args.bType)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}
