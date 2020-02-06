package cryptoutil

import (
	"crypto/rand"
	"encoding/hex"
	"io"
	"reflect"
	"testing"
)

func TestAESGCMEncryptAndDecrypt(t *testing.T) {
	type args struct {
		key   []byte
		nonce []byte
		data  []byte
	}
	data := []byte("ExampleDataForTestOnly")
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	nonce2, _ := hex.DecodeString("1d77707779945c1bd2e6a60a")
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "TC1",
			args: args{
				key:   []byte("This-Is-TheKey-With-32Characters"),
				nonce: nonce,
				data:  data,
			},
			want:    data,
			wantErr: false,
		},
		{
			name: "TC2",
			args: args{
				key:   []byte("This-Is-TheKeyWith-31Characters"),
				nonce: nonce,
				data:  data,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "TC3",
			args: args{
				key:   []byte("Arm2nCLDm4vyuJ5GyKCjHJCshAXG7dZ2"),
				nonce: nonce2,
				data:  []byte("https://google.com"),
			},
			want:    []byte("https://google.com"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, err := AESGCMEncrypt(tt.args.key, tt.args.nonce, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("AESGCMEncrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			got, err := AESGCMDecrypt(tt.args.key, tt.args.nonce, got1)
			if (err != nil) != tt.wantErr {
				t.Errorf("AESGCMEncrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AESGCMEncrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}
