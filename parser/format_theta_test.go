package parser

import "testing"

var thetaSlice01 = []string{
	"(0,2720) ; CL",
	"(0,2650) ; V2",
	"(0,7730) ; V3",
	"(0,3410)",
	"(0,0.232) FIX ; KA transit",
}

var thetaSlice01Cleaned = []string{
	"(0,2720)      ; CL",
	"(0,2650)      ; V2",
	"(0,7730)      ; V3",
	"(0,3410)      ; <NEED COMMENT>",
	"(0,0.232) FIX ; KA transit",
}

func TestFormattingThetaBlock(t *testing.T) {

	formattedData := FormatThetaBlock(thetaSlice01)
	for i, val := range formattedData {
		if val != thetaSlice01Cleaned[i] {
			t.Log("GOT: ", val, " EXPECTED: ", thetaSlice01Cleaned[i])
			t.Fail()
		}
	}
}
