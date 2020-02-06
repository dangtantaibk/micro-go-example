package strutil

import "testing"

func TestNormalizeUnicode(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TC01",
			args: args{s: "Đoàn quân"},
			want: "Doan quan",
		},
		{
			name: "TC02",
			args: args{s: "Tại sao lại như vậy"},
			want: "Tai sao lai nhu vay",
		},
		{
			name: "TC03",
			args: args{s: "đoàn quân"},
			want: "doan quan",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NormalizeUnicode(tt.args.s); got != tt.want {
				t.Errorf("NormalizeUnicode() = %v, want %v", got, tt.want)
			}
		})
	}
}
