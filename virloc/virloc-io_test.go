package virloc_test

import (
	"testing"

	"github.com/Victtor-Souza/virlocio/virloc"
)

func TestSerialize(t *testing.T) {
	type TestCase struct {
		test     string
		message  string
		expected error
	}

	testCases := []TestCase{
		{test: "Message RUV01",
			message:  ">RUV01102,NT003,271123183811+0000000+00000000000000700DE0000,04251140,0,0,0,0,0,0,0,0,0,0,0,4G:0,00000;ID=0675;#2FC1;*20<",
			expected: nil,
		},
		{test: "Message QTT",
			message:  ">RTT010100030036+0000000+000000000000009FFDE0000 20000000 000 13281142087703991127;ID=0675;#FFFF;*5C<",
			expected: nil,
		},
		{
			test:     "Message QGP",
			message:  ">RGP110715030802-3597296-062735570000000FF5F2500;ID=0675;#FFFF;*5C<",
			expected: nil,
		},
		{
			test:     "Message QSD",
			message:  ">RSD0035;ID=0675;#FFFF;*5C<",
			expected: nil,
		},
		{
			test:     "MESSAGE RUV00",
			message:  ">RUV00154,NT003,240622181404,04211377,34L10030,155B13,BR062,V1.2_CAN,VL08,0,0,0,89550537110004889777;ID=0081;#826E;*4F<",
			expected: nil,
		},
		{
			test:     "MESSAGE RUV03",
			message:  ">RUV03152,NT003,281123124031,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0;ID=0675;#4969;*53<",
			expected: nil,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.test, func(t *testing.T) {
			rp := virloc.NewVirlocReport(tC.message)
			if expected := virloc.Serialize(tC.message, rp); expected != tC.expected {
				t.Fatalf("%s FAILED | expected: %v, got %v", tC.test, tC.expected, expected)
				return
			}
		})
	}
}
