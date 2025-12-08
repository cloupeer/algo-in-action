package stringz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Isomorphic(t *testing.T) {
	assert := assert.New(t)

	type args struct {
		s string
		t string
	}

	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{s: "egg", t: "add"},
			want: true,
		},
		{
			args: args{s: "foo", t: "bar"},
			want: false,
		},
		{
			args: args{s: "paper", t: "title"},
			want: true,
		},
	}

	for _, tt := range tests {
		actual := IsIsomorphic(tt.args.s, tt.args.t)
		assert.Equal(tt.want, actual)
	}
}
