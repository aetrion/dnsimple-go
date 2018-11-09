package webhook

import (
	"github.com/dnsimple/dnsimple-go/dnsimple"
)

func switchEvent(name string, payload []byte) (Event, error) {
	var event Event

	switch name {
	case // account
		"account.update",
		"account.billing_settings_update",
		//"account.add_user",
		"account.remove_user":
		event = &AccountEvent{}
	case // contact
		"contact.create",
		"contact.delete",
		"contact.update":
		event = &ContactEvent{}
	case // dnssec
		"dnssec.rotation_complete",
		"dnssec.rotation_start":
		event = &DNSSECEvent{}
	case // domain
		"domain.auto_renewal_disable",
		"domain.auto_renewal_enable",
		"domain.create",
		"domain.delete",
		"domain.register",
		"domain.renew",
		"domain.delegation_change",
		"domain.registrant_change",
		"domain.resolution_disable",
		"domain.resolution_enable",
		"domain.transfer":					// TODO
		event = &DomainEvent{}
	case // email forward
		"email_forward.create",
		"email_forward.delete",
		"email_forward.update":
		event = &EmailForwardEvent{}
	case // webhook
		"webhook.create",
		"webhook.delete":
		event = &WebhookEvent{}
	case // whois privacy
		"whois_privacy.disable",
		"whois_privacy.enable",
		"whois_privacy.purchase",
		"whois_privacy.renew":				// TODO
		event = &WhoisPrivacyEvent{}
	case // zone
		"zone.create",
		"zone.delete":
		event = &ZoneEvent{}
	case // zone record
		"zone_record.create",
		"zone_record.delete",
		"zone_record.update":
		event = &ZoneRecordEvent{}
	default:
		event = &GenericEvent{}
	}

	return event, event.parse(payload)
}

//
// GenericEvent
//

// GenericEvent represents a generic event, where the data is a simple map of strings.
type GenericEvent struct {
	EventHeader
	Data interface{} `json:"data"`
}

func (e *GenericEvent) parse(payload []byte) error {
	e.payload = payload
	return unmashalEvent(payload, e)
}

// ParseGenericEvent unpacks the data into a GenericEvent.
func ParseGenericEvent(e *GenericEvent, payload []byte) error {
	return e.parse(payload)
}

//
// AccountEvent
//

// AccountEvent represents the base event sent for an account action.
type AccountEvent struct {
	EventHeader
	Data    *AccountEvent     `json:"data"`
	Account *dnsimple.Account `json:"account"`
}

// ParseAccountEvent unpacks the data into an AccountEvent.
func ParseAccountEvent(e *AccountEvent, payload []byte) error {
	return e.parse(payload)
}

func (e *AccountEvent) parse(payload []byte) error {
	e.payload, e.Data = payload, e
	return unmashalEvent(payload, e)
}

//
// ContactEvent
//

// ContactEvent represents the base event sent for a contact action.
type ContactEvent struct {
	EventHeader
	Data    *ContactEvent     `json:"data"`
	Contact *dnsimple.Contact `json:"contact"`
}

// ParseContactEvent unpacks the data into a ContactEvent.
func ParseContactEvent(e *ContactEvent, payload []byte) error {
	return e.parse(payload)
}

func (e *ContactEvent) parse(payload []byte) error {
	e.payload, e.Data = payload, e
	return unmashalEvent(payload, e)
}

//
// DNSSECEvent
//

// DNSSECEvent represents the base event sent for a DNSSEC action.
type DNSSECEvent struct {
	EventHeader
	Data                   *DNSSECEvent                     `json:"data"`
	DelegationSignerRecord *dnsimple.DelegationSignerRecord `json:"delegation_signer_record"`
	//DNSSECConfig           *dnsimple.DNSSECConfig           `json:"dnssec"`
}

// ParseDNSSECEvent unpacks the payload into a DNSSECEvent.
func ParseDNSSECEvent(e *DNSSECEvent, payload []byte) error {
	return e.parse(payload)
}

func (e *DNSSECEvent) parse(payload []byte) error {
	e.payload, e.Data = payload, e
	return unmashalEvent(payload, e)
}

//
// DomainEvent
//

// DomainEvent represents the base event sent for a domain action.
type DomainEvent struct {
	EventHeader
	Data       *DomainEvent         `json:"data"`
	Domain     *dnsimple.Domain     `json:"domain"`
	Registrant *dnsimple.Contact    `json:"registrant"`
	Delegation *dnsimple.Delegation `json:"name_servers"`
}

// ParseDomainEvent unpacks the payload into a DomainEvent.
func ParseDomainEvent(e *DomainEvent, payload []byte) error {
	return e.parse(payload)
}

func (e *DomainEvent) parse(payload []byte) error {
	e.payload, e.Data = payload, e
	return unmashalEvent(payload, e)
}

//
// EmailForwardEvent
//

// EmailForwardEvent represents the base event sent for an email forward action.
type EmailForwardEvent struct {
	EventHeader
	Data         *EmailForwardEvent     `json:"data"`
	EmailForward *dnsimple.EmailForward `json:"email_forward"`
}

// ParseEmailForwardEvent unpacks the payload into a EmailForwardEvent.
func ParseEmailForwardEvent(e *EmailForwardEvent, payload []byte) error {
	return e.parse(payload)
}

func (e *EmailForwardEvent) parse(payload []byte) error {
	e.payload, e.Data = payload, e
	return unmashalEvent(payload, e)
}

//
// WebhookEvent
//

// WebhookEvent represents the base event sent for a webhook action.
type WebhookEvent struct {
	EventHeader
	Data    *WebhookEvent     `json:"data"`
	Webhook *dnsimple.Webhook `json:"webhook"`
}

// ParseWebhookEvent unpacks the data into a WebhookEvent.
func ParseWebhookEvent(e *WebhookEvent, payload []byte) error {
	return e.parse(payload)
}

func (e *WebhookEvent) parse(payload []byte) error {
	e.payload, e.Data = payload, e
	return unmashalEvent(payload, e)
}

//
// WhoisPrivacyEvent
//

// WhoisPrivacyEvent represents the base event sent for a whois privacy action.
type WhoisPrivacyEvent struct {
	EventHeader
	Data         *WhoisPrivacyEvent     `json:"data"`
	Domain       *dnsimple.Domain       `json:"domain"`
	WhoisPrivacy *dnsimple.WhoisPrivacy `json:"whois_privacy"`
}

// ParseWhoisPrivacyEvent unpacks the data into a WhoisPrivacyEvent.
func ParseWhoisPrivacyEvent(e *WhoisPrivacyEvent, payload []byte) error {
	return e.parse(payload)
}

func (e *WhoisPrivacyEvent) parse(payload []byte) error {
	e.payload, e.Data = payload, e
	return unmashalEvent(payload, e)
}

//
// ZoneEvent
//

// ZoneEvent represents the base event sent for a zone action.
type ZoneEvent struct {
	EventHeader
	Data *ZoneEvent     `json:"data"`
	Zone *dnsimple.Zone `json:"zone"`
}

// ParseZoneEvent unpacks the data into a ZoneEvent.
func ParseZoneEvent(e *ZoneEvent, payload []byte) error {
	return e.parse(payload)
}

func (e *ZoneEvent) parse(payload []byte) error {
	e.payload, e.Data = payload, e
	return unmashalEvent(payload, e)
}

//
// ZoneRecordEvent
//

// ZoneRecordEvent represents the base event sent for a zone record action.
type ZoneRecordEvent struct {
	EventHeader
	Data       *ZoneRecordEvent     `json:"data"`
	ZoneRecord *dnsimple.ZoneRecord `json:"zone_record"`
}

// ParseZoneRecordEvent unpacks the data into a ZoneRecordEvent.
func ParseZoneRecordEvent(e *ZoneRecordEvent, payload []byte) error {
	return e.parse(payload)
}

func (e *ZoneRecordEvent) parse(payload []byte) error {
	e.payload, e.Data = payload, e
	return unmashalEvent(payload, e)
}
