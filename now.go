package now

import "fmt"

// New initializes Now client
func New(token string) *Now {
	client := Client{
		Token:   token,
		BaseURL: "https://api.zeit.co/now",
	}
	return &Now{client}
}

// Now represents the API
type Now struct {
	client Client
}

// GetDeployments returns all deployments for your account
func (n *Now) GetDeployments() (DeploymentsResponse, error) {
	result := DeploymentsResponse{}
	err := n.client.Call("GET", "/deployments", nil, &result)
	return result, err
}

// GetDeployment fetches a single deployment by id
func (n *Now) GetDeployment(deploymentID string) (Deployment, error) {
	result := Deployment{}
	err := n.client.Call("GET", fmt.Sprintf("/deployments/%s", deploymentID), nil, &result)
	return result, err
}

// CreateDeployment creates a new deployment
func (n *Now) CreateDeployment(body Deployment) (Deployment, error) {
	result := Deployment{}
	err := n.client.Call("POST", "/deployments", body, &result)
	return result, err
}

// DeleteDeployment removes a deployment with a given id
func (n *Now) DeleteDeployment(deploymentID string) error {
	err := n.client.Call("DELETE", fmt.Sprintf("/deployments/%s", deploymentID), nil, nil)
	return err
}

//
// // GetFiles will return all files for a given deployment id
// func (n *Now) GetFiles(deploymentID string) (FilesResponse, error) {
//
// }
//
// // GetFile returns a file specified by deployment and file id
// func (n *Now) GetFile(deploymentID, fileID string) error {
//
// }
//
// // GetDomains returns all domains on your account
// func (n *Now) GetDomains() (DomainsResponse, error) {
//
// }
//
// // AddDomain adds a new domain
// func (n *Now) AddDomain(domain Domain) error {
//
// }
//
// // DeleteDomain will remove the domain with a given name
// func (n *Now) DeleteDomain(domainName string) error {
//
// }
//
// // GetDomainRecords returns DNSRecordsResponse for a given domain name
// func (n *Now) GetDomainRecords(domainName string) (DNSRecordsResponse, error) {
//
// }
//
// // AddDomainRecords sets a DNSRecord for a given domain name
// func (n *Now) AddDomainRecords(domainName string, record DNSRecord) error {
//
// }
//
// // DeleteDomainRecord deletes a domain record for a given domain name and record id
// func (n *Now) DeleteDomainRecord(domainName, recordID string) error {
//
// }
//
// // GetCertificates returns all registered certificates
// func (n *Now) GetCertificates(commonName string) (CertificatesResponse, error) {
//
// }
//
// // CreateCertificate creates a new certificate for given domain names
// func (n *Now) CreateCertificate(domainNames []string) error {
//
// }
//
// // RenewCertificate will renew an existing certificate for given domain names
// func (n *Now) RenewCertificate(domainNames []string) error {
//
// }
//
// // ReplaceCertificate will update the cert with new certificate, key and CA chain
// func (n *Now) ReplaceCertificate(commonName, certificate, key, chain string) error {
//
// }
//
// // DeleteCertificate will delete the certificate by a given name
// func (n *Now) DeleteCertificate(commonName string) error {
//
// }
//
// // GetAllAliases will return all aliases set on your account
// func (n *Now) GetAllAliases() (AliasesResponse, error) {
//
// }
//
// // GetAliases will return all aliases set for a given deployment
// func (n *Now) GetAliases(deploymentID string) (AliasesResponse, error) {
//
// }
//
// // CreateAlias will create a new alias
// func (n *Now) CreateAlias(alias Alias) (Alias, error) {
//
// }
//
// // DeleteAlias will delete alias with a given alias id
// func (n *Now) DeleteAlias(aliasID string) error {
//
// }
//
// // GetSecrets will return all set secrets
// func (n *Now) GetSecrets() (SecretsResponse, error) {
//
// }
//
// // CreateSecret will create a new secret with a given name and value
// func (n *Now) CreateSecret(name, value string) (Secret, error) {
//
// }
//
// // RenameSecret will replace old secret name with a new name
// func (n *Now) RenameSecret(idOrName, newName string) (Secret, error) {
//
// }
//
// // DeleteSecret will delete a secret with a given id or name
// func (n *Now) DeleteSecret(idOrName string) error {
//
// }
