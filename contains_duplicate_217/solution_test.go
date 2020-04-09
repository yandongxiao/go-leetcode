package contains_duplicate_217

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_containsDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "contains duplicate",
			args: args{
				nums: []int{
					66, 2, 3, 66,
				},
			},
			want: true,
		},
		{
			name: "contains duplicate",
			args: args{
				nums: []int{
					0, 4, 5, 0, 3, 6,
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, containsDuplicate(tt.args.nums))
		})
	}
}
