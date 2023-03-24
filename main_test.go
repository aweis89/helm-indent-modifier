package main

import (
	"reflect"
	"testing"
)

func Test_processIndent(t *testing.T) {
	t.Parallel()
	type args struct {
		content   []byte
		inc       int
		dec       int
		startLine int
		endLine   int
	}
	tests := map[string]struct {
		args    args
		want    []byte
		wantErr bool
	}{
		"should inc indent": {
			args: args{
				content:   []byte("this:\n {{ toYaml .Foo | indent 10 }}"),
				inc:       3,
				dec:       0,
				startLine: 1,
				endLine:   2,
			},
			want:    []byte("this:\n {{ toYaml .Foo | indent 13 }}"),
			wantErr: false,
		},
		"should inc nindent": {
			args: args{
				content:   []byte("this:\n {{ toYaml .Foo | nindent 10 }}"),
				inc:       3,
				dec:       0,
				startLine: 1,
				endLine:   2,
			},
			want:    []byte("this:\n {{ toYaml .Foo | nindent 13 }}"),
			wantErr: false,
		},
	}
	for name, tt := range tests {
		name, tt := name, tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got, err := processIndent(tt.args.content, tt.args.inc, tt.args.dec, tt.args.startLine, tt.args.endLine)
			if (err != nil) != tt.wantErr {
				t.Errorf("processIndent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("processIndent() = %s, want %s", got, tt.want)
			}
		})
	}
}
