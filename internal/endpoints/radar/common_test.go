package radar

import (
	"encoding/json"
	"testing"
)

func TestStringOrNumber_UnmarshalJSON(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want string
	}{
		{"string", `"23"`, "23"},
		{"int", `23`, "23"},
		{"float", `23.5`, "23.5"},
		{"null", `null`, ""},
		{"empty string", `""`, ""},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var v StringOrNumber
			if err := json.Unmarshal([]byte(tc.in), &v); err != nil {
				t.Fatalf("unmarshal %s: %v", tc.in, err)
			}
			if string(v) != tc.want {
				t.Fatalf("got %q, want %q", string(v), tc.want)
			}
		})
	}
}

func TestRDAProperties_UnmarshalsResolutionVersion(t *testing.T) {
	cases := []struct {
		name string
		body string
		want string
	}{
		{"numeric", `{"resolutionVersion": 23}`, "23"},
		{"string", `{"resolutionVersion": "23"}`, "23"},
		{"null", `{"resolutionVersion": null}`, ""},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var p RDAProperties
			if err := json.Unmarshal([]byte(tc.body), &p); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}
			if tc.want == "" {
				if p.ResolutionVersion != nil && string(*p.ResolutionVersion) != "" {
					t.Fatalf("expected nil/empty, got %q", *p.ResolutionVersion)
				}
				return
			}
			if p.ResolutionVersion == nil {
				t.Fatalf("expected %q, got nil", tc.want)
			}
			if string(*p.ResolutionVersion) != tc.want {
				t.Fatalf("got %q, want %q", *p.ResolutionVersion, tc.want)
			}
		})
	}
}

func TestStringOrNumber_MarshalJSON(t *testing.T) {
	cases := []struct {
		in   StringOrNumber
		want string
	}{
		{"", `null`},
		{"23", `23`},
		{"23.5", `23.5`},
		{"hello", `"hello"`},
	}
	for _, tc := range cases {
		t.Run(string(tc.in), func(t *testing.T) {
			b, err := json.Marshal(tc.in)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			if string(b) != tc.want {
				t.Fatalf("got %s, want %s", b, tc.want)
			}
		})
	}
}
