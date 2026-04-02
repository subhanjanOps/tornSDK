package httpclient

import "strings"

// Selections builds Torn "selections" query parameter values.
type Selections struct {
	parts []string
}

// NewSelections creates an empty selection builder.
func NewSelections(parts ...string) *Selections {
	s := &Selections{}
	s.Add(parts...)
	return s
}

// Add appends non-empty trimmed selection parts.
func (s *Selections) Add(parts ...string) *Selections {
	if s == nil {
		return s
	}

	for _, p := range parts {
		if t := strings.TrimSpace(p); t != "" {
			s.parts = append(s.parts, t)
		}
	}

	return s
}

// String returns the comma-joined selections or an empty string.
func (s *Selections) String() string {
	if s == nil || len(s.parts) == 0 {
		return ""
	}
	return strings.Join(s.parts, ",")
}

// Apply sets the selections on the provided request.
func (s *Selections) Apply(r *Request) *Request {
	if r == nil {
		return nil
	}
	if s == nil || s.String() == "" {
		return r.SetQuery("selections")
	}
	return r.SetQuery("selections", s.String())
}
