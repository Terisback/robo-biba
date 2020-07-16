package utils

import "testing"

func TestGetIntColor(t *testing.T) {
	type args struct {
		embColor string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Red",
			args: args{
				embColor: "ff0000",
			},
			want: 16711680, // From online converter
		},
		{
			name: "Green",
			args: args{
				embColor: "00ff00",
			},
			want: 65280, // From online converter
		},
		{
			name: "Blue",
			args: args{
				embColor: "0000ff",
			},
			want: 255, // From online converter
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetIntColor(tt.args.embColor); got != tt.want {
				t.Errorf("GetIntColor() = %v, want %v", got, tt.want)
			}
		})
	}
}
