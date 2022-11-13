package call

import (
	"github.com/hero-soft/web-scanner/pkg/talkgroup"
)

type Call struct {
	ID        string              `json:"id,omitempty"`
	Talkgroup talkgroup.Talkgroup `json:"talkgroup,omitempty"`
	Emergency bool                `json:"emergency,omitempty"`
	Priority  int                 `json:"priority,omitempty"`
	File      string              `json:"file,omitempty"`
}
