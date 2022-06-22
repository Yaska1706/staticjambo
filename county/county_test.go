package county

import (
	"testing"
)

func Test_countyCodeService_CountyName(t *testing.T) {
	var cService countyCodeService
	type args struct {
		code string
	}
	tests := []struct {
		name    string
		c       *countyCodeService
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "",
			args: args{
				code: "2",
			},
			c:       &cService,
			want:    "Kwale",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &countyCodeService{}
			got, err := c.CountyName(tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("countyCodeService.CountyName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("countyCodeService.CountyName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countyCodeService_CountyCode(t *testing.T) {
	var cService countyCodeService

	type args struct {
		name string
	}
	tests := []struct {
		name    string
		c       *countyCodeService
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				name: "Kwale",
			},
			c:       &cService,
			want:    "2",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &countyCodeService{}
			got, err := c.CountyCode(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("countyCodeService.CountyCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("countyCodeService.CountyCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
