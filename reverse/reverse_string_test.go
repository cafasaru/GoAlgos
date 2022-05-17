package reverse

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReverse(t *testing.T) {
	testCases := []struct {
		name        string
		want        string
		checkResult func(t *testing.T, want string, got string)
	}{
		{
			name: "ok reversed",
			want: "olleh",
			checkResult: func(t *testing.T, want string, got string) {
				require.Equal(t, want, got)
			},
		},

	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			word := "hello"

			got, err := reverse(word)
			require.NoError(t, err)

			tc.checkResult(t, tc.want, got)
		})
	}
}
