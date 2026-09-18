package mesh

import "testing"

func TestTargetEqual(t *testing.T) {
	base := Target{Segments: []string{"orders", "create"}, Kind: KindRoute}
	cases := []struct {
		name  string
		other Target
		want  bool
	}{
		{name: "same segments and kind", other: Target{Segments: []string{"orders", "create"}, Kind: KindRoute}, want: true},
		{name: "metadata ignored", other: Target{Segments: []string{"orders", "create"}, Kind: KindRoute, Metadata: map[string]string{"k": "v"}}, want: true},
		{name: "different kind", other: Target{Segments: []string{"orders", "create"}, Kind: KindTopic}, want: false},
		{name: "different length", other: Target{Segments: []string{"orders"}, Kind: KindRoute}, want: false},
		{name: "one differing segment", other: Target{Segments: []string{"orders", "cancel"}, Kind: KindRoute}, want: false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := base.Equal(tc.other); got != tc.want {
				t.Fatalf("want %v, got %v", tc.want, got)
			}
			if got := tc.other.Equal(base); got != tc.want {
				t.Fatalf("reverse: want %v, got %v", tc.want, got)
			}
		})
	}
}

func TestKindString(t *testing.T) {
	cases := []struct {
		kind Kind
		want string
	}{
		{KindRoute, "route"},
		{KindTopic, "topic"},
		{Kind(99), "unknown"},
	}
	for _, tc := range cases {
		if got := tc.kind.String(); got != tc.want {
			t.Errorf("Kind(%d).String() = %q, want %q", int(tc.kind), got, tc.want)
		}
	}
}
