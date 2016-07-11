package dnsimple

import (
	"fmt"
)

// Domain represents a list of name servers that correspond to a domain delegation.
type Delegation []string

// Vanity Name Server
type VanityNameServer struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	IPv4      string `json:"ipv4,omitempty"`
	IPv6      string `json:"ipv6,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

// WhoisPrivacyResponse represents a response from an API method that returns a delegation struct.
type DelegationResponse struct {
	Response
	Data *Delegation `json:"data"`
}

// Response for vanity name server enable and disable operations
type VanityDelegationResponse struct {
	Response
	Data []VanityNameServer `json:"data"`
}

// GetDomainDelegation gets the current delegated name servers for the domain.
//
// See https://developer.dnsimple.com/v2/registrar/delegation/#get
func (s *RegistrarService) GetDomainDelegation(accountID string, domainName string) (*DelegationResponse, error) {
	path := versioned(fmt.Sprintf("/%v/registrar/domains/%v/delegation", accountID, domainName))
	delegationResponse := &DelegationResponse{}

	resp, err := s.client.get(path, delegationResponse)
	if err != nil {
		return nil, err
	}

	delegationResponse.HttpResponse = resp
	return delegationResponse, nil
}

// ChangeDomainDelegation updates the delegated name severs for the domain.
//
// See https://developer.dnsimple.com/v2/registrar/delegation/#get
func (s *RegistrarService) ChangeDomainDelegation(accountID string, domainName string, newDelegation *Delegation) (*DelegationResponse, error) {
	path := versioned(fmt.Sprintf("/%v/registrar/domains/%v/delegation", accountID, domainName))
	delegationResponse := &DelegationResponse{}

	resp, err := s.client.put(path, newDelegation, delegationResponse)
	if err != nil {
		return nil, err
	}

	delegationResponse.HttpResponse = resp
	return delegationResponse, nil
}

// ChangeDomainDelegationToVanity enables vanity name servers for the given domain.
//
// See https://developer.dnsimple.com/v2/registrar/delegation/#delegateToVanity
func (s *RegistrarService) ChangeDomainDelegationToVanity(accountID string, domainName string, newDelegation *Delegation) (*VanityDelegationResponse, error) {
	path := versioned(fmt.Sprintf("/%v/registrar/domains/%v/delegation/vanity", accountID, domainName))
	delegationResponse := &VanityDelegationResponse{}

	resp, err := s.client.put(path, newDelegation, delegationResponse)
	if err != nil {
		return nil, err
	}

	delegationResponse.HttpResponse = resp
	return delegationResponse, nil
}

// ChangeDomainDelegationFromVanity disables vanity name servers for the given domain.
//
// See https://developer.dnsimple.com/v2/registrar/delegation/#dedelegateFromVanity
func (s *RegistrarService) ChangeDomainDelegationFromVanity(accountID string, domainName string) (*VanityDelegationResponse, error) {
	path := versioned(fmt.Sprintf("/%v/registrar/domains/%v/delegation/vanity", accountID, domainName))
	delegationResponse := &VanityDelegationResponse{}

	resp, err := s.client.delete(path, nil, nil)
	if err != nil {
		return nil, err
	}

	delegationResponse.HttpResponse = resp
	return delegationResponse, nil
}
