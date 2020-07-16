package utils

import "testing"

func TestAliases(t *testing.T) {
	type args struct {
		text    string
		aliases []string
	}
	tests := []struct {
		name   string
		args   args
		wantOk bool
	}{
		{
			name: "Usual",
			args: args{
				text: "Test",
				aliases: []string{
					"tEsT",
					"not a test",
				},
			},
			wantOk: true,
		},
		{
			name: "Not Usual",
			args: args{
				text: "Test",
				aliases: []string{
					"well tEsTeD",
					"not a test",
				},
			},
			wantOk: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOk := Aliases(tt.args.text, tt.args.aliases...); gotOk != tt.wantOk {
				t.Errorf("Aliases() = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
