package sumlist

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSumlist(t *testing.T) {
	testCases := []struct {
		name        string
		params      []int
		want        int
		checkResult func(t *testing.T, got int, want int)
	}{
		{
			name:   "ok result",
			params: []int{2, 3, 5, 2, 5},
			want:   17,
			checkResult: func(t *testing.T, got int, want int) {
				require.Equal(t, want, got)
			},
		},
		{
			name:   "ok zero",
			params: []int{0},
			want:   0,
			checkResult: func(t *testing.T, got, want int) {
				require.Equal(t, got, want)
			},
		},
		{
			name:   "fail not match want",
			params: []int{1, 2, 2, 5},
			want:   11,
			checkResult: func(t *testing.T, got, want int) {
				require.NotEqual(t, want, got)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			total := sumList(tc.params)
			tc.checkResult(t, total, tc.want)
		})

	}
}
