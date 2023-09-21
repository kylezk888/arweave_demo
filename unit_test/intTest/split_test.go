package intTest

import (
	"testing"
)

func TestSplitAll(t *testing.T) {
	t.Parallel()
	//assert := assert.New(t)

	type args struct {
		s   string
		sep string
	}

	tests := []struct {
		name       string
		args       args
		wantResult []string
	}{
		{"test1", args{"a:b:c", ":"}, []string{"a", "b", "c"}},
		{"test2", args{"a:b:c", ","}, []string{"a:b:c"}},
		{"test3", args{"abcd", "bc"}, []string{"a", "c"}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			//res := testcase.Split(tt.args.s, tt.args.sep)
			//assert.Equal(tt.wantResult, res, "not equal")
			//require.Equal(t, tt.wantResult, res, "not equal")
			//if !reflect.DeepEqual(tt.wantResult, res) {
			//	t.Error("note: not equal")
			//}
		})

	}
}
