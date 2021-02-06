package zonotools

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestEncodeJson(t *testing.T) {
	type args struct {
		jsonString string
		data       interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := EncodeJson(tt.args.jsonString, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("EncodeJson() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	t.Run("normal", func(t *testing.T) {
		s := struct {
			Id   int    `json:"id"`
			Name string `json:"name"`
		}{}
		err := EncodeJson(`{"id":123, "name":"test"}`, &s)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, 123, s.Id)
		assert.Equal(t, "test", s.Name)
	})
	t.Run("fail", func(t *testing.T) {
		s := struct {
			Id   int    `json:"id"`
			Name string `json:"name"`
		}{}
		err := EncodeJson(`{"id":123, "name":"test"}`, s)
		if err == nil {
			t.Fail()
		}
	})
}

func TestDecodeJson(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DecodeJson(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeJson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodeJson() got = %v, want %v", got, tt.want)
			}
		})
	}
	t.Run("normal", func(t *testing.T) {
		s := struct {
			Id   int    `json:"id"`
			Name string `json:"name"`
		}{123, "test"}
		data, err := DecodeJson(s)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, []byte(`{"id":123,"name":"test"}`), data)
	})
	t.Run("normal", func(t *testing.T) {
		s := struct {
			Id   int    `json:"id"`
			Name string `json:"name"`
		}{123, "test"}
		data, err := DecodeJson(&s)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, []byte(`{"id":123,"name":"test"}`), data)
	})
	t.Run("normal", func(t *testing.T) {
		m := map[string]interface{}{"id": 123, "name": "test"}
		data, err := DecodeJson(m)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, []byte(`{"id":123,"name":"test"}`), data)
	})
}
