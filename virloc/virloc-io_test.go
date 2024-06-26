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
			message:  ">RUV01102,NT003,130524144211-2368988-04656462000072300C20001,04302843,0,0,0,231201,170097245,0,0,0,55,0,1,4G:1,00000;ID=2990;#373D;*25<",
			expected: nil,
		},
		{
			test:     "MESSAGE RUV03",
			message:  ">RUV03150,NT003,030822184257,5F,000,65,69584,120455002,2312,86,10,29,44895330,0,40,36,0,0,0,0,0,0,0,0,0;ID=0679;#0030;*54<",
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
			test:     "MESSAGE RUV02",
			message:  ">RUV02108,NT003,240622184842,11111,22302,33333,44444,55555,66666,77777,88888,00361,600000,300000;ID=0081;#8020;*43<",
			expected: nil,
		},
		{
			test:     "DEFAULT MESSAGE",
			message:  ">RSXP01;ID=0081;#8020;*43<",
			expected: nil,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.test, func(t *testing.T) {
			rp, _ := virloc.NewVirlocReport(tC.message)
			if expected := virloc.Serialize(tC.message, rp); expected != tC.expected {
				t.Fatalf("%s FAILED | expected: %v, got %v", tC.test, tC.expected, expected)
				return
			}
		})
	}
}

func Test_GetSerialNumberId(t *testing.T) {
	type TestCase struct {
		test     string
		message  string
		expected string
	}

	testCases := []TestCase{
		{
			test:     "Test SerialNumber",
			message:  ">RUV01102,NT003,271123183811+0000000+00000000000000700DE0000,04251140,0,0,0,0,0,0,0,0,0,0,0,4G:0,00000;ID=0675;#2FC1;*20<",
			expected: "0675",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			rp, _ := virloc.NewVirlocReport(tc.message)
			if rp.GetSerialNumber() != tc.expected {
				t.Fatalf("%s FAILED | expected %v, got %v", tc.test, tc.expected, rp.GetSerialNumber())
				return
			}
		})
	}
}

func Test_CalculateCheckSum(t *testing.T) {
	type TestCase struct {
		test             string
		message          string
		expectedCheckSum string
	}

	testCases := []TestCase{
		{
			test:             "Test checksum",
			message:          ">SSXP10;ID=0675;#FFFF;",
			expectedCheckSum: "1B",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			ck := virloc.CalculateChecksum(tc.message)

			if ck != tc.expectedCheckSum {
				t.Fatalf("FAILED %s | Expected %s, got %s", tc.test, tc.expectedCheckSum, ck)
			}
		})
	}
}
