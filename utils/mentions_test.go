package utils

import (
	"fmt"
	"testing"
)

const (
	id      = 281037696225247233
	otherID = 296264598577741825
)

func TestGetIDFromArg(t *testing.T) {
	type args struct {
		arg string
	}
	tests := []struct {
		name   string
		args   args
		wantId uint64
		wantOk bool
	}{
		{
			name:   "DefaultID",
			args:   args{arg: fmt.Sprintf("%d", id)},
			wantId: id,
			wantOk: true,
		},
		{
			name:   "IDWithArgs",
			args:   args{arg: fmt.Sprintf("%d online %d", id, otherID)},
			wantId: id,
			wantOk: true,
		},
		{
			name:   "DefaultMention",
			args:   args{arg: fmt.Sprintf("<@%d>", id)},
			wantId: id,
			wantOk: true,
		},
		{
			name:   "MentionWithArgs",
			args:   args{arg: fmt.Sprintf("<@%d> online %d", id, otherID)},
			wantId: id,
			wantOk: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotId, gotOk := GetIDFromArg(tt.args.arg)
			if gotId != tt.wantId {
				t.Errorf("GetIDFromArg() gotId = %v, want %v", gotId, tt.wantId)
			}
			if gotOk != tt.wantOk {
				t.Errorf("GetIDFromArg() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
