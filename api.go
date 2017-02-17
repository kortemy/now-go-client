package now

// API represents the definition of Now API
type API interface {
	GetDeployments() (DeploymentsResponse, error)
	GetDeployment(deploymentID string) (Deployment, error)
	CreateDeployment(body Deployment) (Deployment, error)
	DeleteDeployment(deploymentID string) error
	GetFiles(deploymentID string) (FilesResponse, error)
	GetFile(deploymentID, fileID string) error
	GetDomains() (DomainsResponse, error)
	AddDomain(domain Domain) error
	DeleteDomain(domainName string) error
	GetDomainRecords(domainName string) (DNSRecordsResponse, error)
	AddDomainRecords(domainName string, record DNSRecord) error
	DeleteDomainRecord(domainName, recordID string) error
	GetCertificates(commonName string) (CertificatesResponse, error)
	CreateCertificate(domainNames []string) error
	RenewCertificate(domainNames []string) error
	ReplaceCertificate(commonName, certificate, key, chain string) error
	DeleteCertificate(commonName string) error
	GetAllAliases() (AliasesResponse, error)
	GetAliases(deploymentID string) (AliasesResponse, error)
	CreateAlias(alias Alias) (Alias, error)
	DeleteAlias(aliasID string) error
	GetSecrets() (SecretsResponse, error)
	CreateSecret(name, value string) (Secret, error)
	RenameSecret(idOrName, newName string) (Secret, error)
	DeleteSecret(idOrName string) error
}

// ErrorResponse matches data for error responses
type ErrorResponse struct {
	Err Error
}

// DeploymentsResponse matches data for getting array of Deployments
type DeploymentsResponse struct {
	Deployments []Deployment
}

// FilesResponse matches data for getting array of Files
type FilesResponse struct {
	Files []File
}

// DomainsResponse matches data for getting array of Domains
type DomainsResponse struct {
	Domains []Domain
}

// DNSRecordsResponse matches data for getting array of DNSRecords
type DNSRecordsResponse struct {
	Records []DNSRecord
}

// CertificatesResponse matches data for getting array of Certificates
type CertificatesResponse struct {
	Certs []Certificate
}

// AliasesResponse matches data for getting array of Aliases
type AliasesResponse struct {
	Aliases []Alias
}

// SecretsResponse matches data for getting array of Secrets
type SecretsResponse struct {
	Secrets []Secret
}
