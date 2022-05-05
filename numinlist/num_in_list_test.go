package numinlist_test

import (
	"GoAlgos/numinlist"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNumInList(t *testing.T) {
	list := []int{3, 4, 12, 523, 34, 2, 3, 5}
	testCases := []struct {
		name        string
		want        bool
		params      []int
		find        int
		checkResult func(*testing.T, bool, bool)
	}{
		{
			name:   "ok found num",
			want:   true,
			params: list,
			find:   3,
			checkResult: func(t *testing.T, got bool, want bool) {
				require.True(t, got == want)
			},
		},
		{
			name:   "fail num not found",
			want:   false,
			params: list,
			find:   450,
			checkResult: func(t *testing.T, got bool, want bool) {
				require.True(t, got == want)
			},
		},
		{
			name:   "fail nil value",
			want:   false,
			params: nil,
			find:   450,
			checkResult: func(t *testing.T, got bool, want bool) {
				require.True(t, got == want)
			},
		},
		{
			name:   "fail empty slice",
			want:   false,
			params: []int{},
			find:   450,
			checkResult: func(t *testing.T, got bool, want bool) {
				require.True(t, got == want)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			got := numinlist.NumInSlice(tc.params, tc.find)
			log.Println("got: ", got, " want: ", tc.want)
			tc.checkResult(t, tc.want, got)
		})
	}
}
