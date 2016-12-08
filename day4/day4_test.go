package day4

import "testing"

var validateCases = []struct {
	name       string
	wantSector int
	wantValid  bool
}{
	{"aaaaa-bbb-z-y-x-123[abxyz]", 123, true},
	{"a-b-c-d-e-f-g-h-987[abcde]", 987, true},
	{"not-a-real-room-404[oarel]", 404, true},
	{"totally-real-room-200[decoy]", 200, false},
}

func TestValidate(t *testing.T) {
	for i, c := range validateCases {
		gotB, _, gotI := validate(c.name)
		if gotB != c.wantValid {
			t.Fatalf("Case %d: Wanted %v, got %v", i+1, c.wantValid, gotB)
		}
		if gotI != c.wantSector {
			t.Fatalf("Case %d: Wanted %v, got %v", i+1, c.wantSector, gotI)
		}
	}
}

var rot1Cases = []struct {
	encrypted string
	sector    int
	want      string
}{
	{"qzmt-zixmtkozy-ivhz", 343, "very encrypted name"},
}

func TestRotX(t *testing.T) {
	for i, c := range rot1Cases {
		got := rotX(c.encrypted, c.sector)
		if got != c.want {
			t.Fatalf("Case %d: Wanted %v, got %v", i+1, c.want, got)
		}
	}
}
