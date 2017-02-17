package now

import (
	"time"
)

// Error model definition
type Error struct {
	Code    string
	Message string
	URL     string
}

// Deployment model definition
type Deployment struct {
	UID   string
	Name  string
	URL   string
	State string
}

// File model definition
type File struct {
	UID      string
	Type     string
	Name     string
	Children []File
}

// Domain model definition
type Domain struct {
	UID        string
	Name       string
	IsExternal bool
	Aliases    []Alias
}

// DNSRecord model definition (Record in now API)
type DNSRecord struct {
	ID         string
	Slug       string
	Type       string
	Name       string
	Value      string
	MXPriority int
}

// Certificate model definition (Cert in now API)
type Certificate struct {
	UID        string
	CN         string
	AutoRenew  bool
	Expiration time.Time
}

// Alias model definition
type Alias struct {
	UID          string
	Alias        string
	DeploymentID string
}

// Secret model definition
type Secret struct {
	UID   string
	Name  string
	Value string
}
