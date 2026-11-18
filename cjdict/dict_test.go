package cjdict

import (
	"reflect"
	"strings"
	"sync"
	"testing"
)

var 哲理類 = "abcdefg"
var 筆劃類 = "hijklmn"
var 人體類 = "opqr"
var 字形類 = "stuvwy"

func TestGetChars(t *testing.T) {
	useTestCodes(t)

	if allChars := GetAllChars(); len(allChars) != 4 {
		t.Errorf("expected 4 chars, got %d", len(allChars))
	}
	if chars := GetChars(); len(chars) != 3 {
		t.Errorf("expected 3 chars, got %d", len(chars))
	}
}

func TestGetCandidate(t *testing.T) {
	useTestCodes(t)

	candidates := GetCandidates("abcdefg")

	want := []rune{'受', '鑫'}
	if !reflect.DeepEqual(candidates, want) {
		t.Errorf("expected %v, got %v", want, candidates)
	}
}

func TestMultiCode(t *testing.T) {
	useCodes(t, strings.Join([]string{
		"mkp,憨",
		"nkp,憨",
		"xmkp,憨",
		"xnkp,憨",
		"xyck,奕",
		"yck,奕",
	}, "\n"))
	cjmap := GetChars()
	want := map[rune]CjCode{
		'奕': {
			Codes:    []string{"xyck", "yck"},
			Bitmasks: []uint32{0x1800404, 0x1000404},
		},
		'憨': {
			Codes:    []string{"mkp", "nkp", "xmkp", "xnkp"},
			Bitmasks: []uint32{0x9400, 0xa400, 0x809400, 0x80a400},
		},
	}
	if !reflect.DeepEqual(cjmap, want) {
		t.Errorf("expected %v, got %v", want, cjmap)
	}
}

func TestGoldenData(t *testing.T) {
	if count := len(GetFreqs()); count != 5705 {
		t.Errorf("expected 5705 freqs, got %d (did you update golden data?)", count)
	}

	tests := map[string]struct {
		chars      map[rune]CjCode
		count      int
		multicodes int

		getCandidates func(string) []rune

		abc int
		hi  int
		op  int
		st  int
		ah  int
	}{
		`allChars`: {
			chars:      GetAllChars(),
			count:      20902,
			multicodes: 2288,

			getCandidates: GetAllCandidates,

			abc: 243,
			hi:  360,
			op:  58,
			st:  99,
			ah:  3206,
		},
		`chars`: {
			chars:      GetChars(),
			count:      5701,
			multicodes: 345,

			getCandidates: GetCandidates,

			abc: 66,
			hi:  108,
			op:  19,
			st:  39,
			ah:  877,
		},
	}

	for name, tt := range tests {
		if count := len(tt.chars); count != tt.count {
			t.Errorf("[%s] expected %d chars, got %d (did you update golden data?)", name, tt.count, count)
		}

		multiCodes := 0
		for _, cjcode := range tt.chars {
			if len(cjcode.Codes) > 1 {
				multiCodes++
			}
		}
		if multiCodes != tt.multicodes {
			t.Errorf("[%s] expected %d multi-code chars, got %d (did you update golden data?)", name, tt.multicodes, multiCodes)
		}

		if got := len(tt.getCandidates(哲理類)); got != tt.abc {
			t.Errorf("[%s] expected %d chars, got %d (did you update golden data?)", name, tt.abc, got)
		}
		if got := len(tt.getCandidates(筆劃類)); got != tt.hi {
			t.Errorf("[%s] expected %d chars, got %d (did you update golden data?)", name, tt.hi, got)
		}
		if got := len(tt.getCandidates(人體類)); got != tt.op {
			t.Errorf("[%s] expected %d chars, got %d (did you update golden data?)", name, tt.op, got)
		}
		if got := len(tt.getCandidates(字形類)); got != tt.st {
			t.Errorf("[%s] expected %d chars, got %d (did you update golden data?)", name, tt.st, got)
		}
		if got := len(tt.getCandidates(哲理類 + 筆劃類)); got != tt.ah {
			t.Errorf("[%s] expected %d chars, got %d (did you update golden data?)", name, tt.ah, got)
		}
	}

}

func useTestCodes(t *testing.T) {
	useCodes(t, strings.Join([]string{
		"abjj,暈,暉",
		"bbe,受",
		"ccc,鑫",
	}, "\n"))
	useFreqs(t, strings.Join([]string{
		"暈,1",
		"受,1",
		"鑫,1",
	}, "\n"))
}

func resetCaches() {
	freqsOnce, fullOnce, poolOnce = sync.Once{}, sync.Once{}, sync.Once{}
	cjFreqs, cjFullTable, cjTable = nil, nil, nil
}

func useCodes(t *testing.T, csv string) {
	t.Helper()
	origCodes := rawCodes

	t.Cleanup(func() {
		rawCodes = origCodes
		resetCaches()
	})
	rawCodes = []byte(csv)
	resetCaches()
}

func useFreqs(t *testing.T, csv string) {
	t.Helper()
	origFreqs := rawFreqs

	t.Cleanup(func() {
		rawFreqs = origFreqs
		resetCaches()
	})
	rawFreqs = []byte(csv)
	resetCaches()
}

func BenchmarkGetCandidates(b *testing.B) {
	GetChars() // warm the tables in order to measure only the query
	for b.Loop() {
		GetCandidates(哲理類 + 筆劃類)
	}
}

func BenchmarkGetAllCandidates(b *testing.B) {
	GetAllChars()
	for b.Loop() {
		GetAllCandidates(哲理類 + 筆劃類)
	}
}

func BenchmarkGetChars(b *testing.B) {
	for b.Loop() {
		resetCaches()
		GetChars()
	}
}
