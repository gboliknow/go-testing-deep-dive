package math

import (
	"runtime"
	"testing"

	"example.com/testing/math/testutil"
)

// func TestMain(m *testing.M){
// 	code := m.Run()

//		os.Exit(code)
//	}
func Test_CalculateDiscount(t *testing.T) {
	got, _ := CalculateDiscount(100, 0.1, false)
	want := 10.0
	if got == want {
		t.Error("CalculateDiscount(100, 0.1, false)  failed: expected ", want, "got", got)
	}
}

func TestCalculateDiscount(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		price    float64
		discount float64
		isMember bool
		want     float64
		wantErr  bool
	}{
		{"no discount", 100, 0, false, 100, false},
		{"base discount only", 100, 10, false, 90, false},
		{"member extra discount", 100, 10, true, 85.5, false},
		{"max discount", 100, 100, false, 0, false},
		{"invalid negative price", -50, 10, false, 0, true},
		{"invalid discount > 100", 100, 150, false, 0, true},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got, err := CalculateDiscount(tc.price, tc.discount, tc.isMember)

			if tc.wantErr {
				testutil.MustFail(t, func() error { return err })
				return // stop here so we don't hit "unexpected error"
			}

			if err != nil {
				t.Fatalf("unexpected error :%v", err)
			}
			testutil.EqualFloat(t, got, tc.want, 0.001)
		})
	}

}

func TestCalculateDiscount_SkipOnWindows(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("skipping test on Windows due to float rounding difference")
	}
	got, _ := CalculateDiscount(100, 5, true)
	testutil.EqualFloat(t, got, 90.25, 0.001)
}

func BenchmarkCalculateDiscount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = CalculateDiscount(100, 10, true)
	}
}
