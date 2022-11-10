package repository

import (
	"reflect"
	"testing"
)

func TestGetConn(t *testing.T) {
	tests := []struct {
		name    string
		want    *DB
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetConn()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetConn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetConn() = %v, want %v", got, tt.want)
			}
		})
	}
}
