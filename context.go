package buildkit

type contextKey struct {
	name string
}

var paramKey = &contextKey{name: "params"}
