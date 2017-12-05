// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Site site
// swagger:model site

type Site struct {

	// account name
	AccountName string `json:"account_name,omitempty"`

	// account slug
	AccountSlug string `json:"account_slug,omitempty"`

	// admin url
	AdminURL string `json:"admin_url,omitempty"`

	// build settings
	BuildSettings *SiteBuildSettings `json:"build_settings,omitempty"`

	// capabilities
	Capabilities map[string]interface{} `json:"capabilities,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// custom domain
	CustomDomain string `json:"custom_domain,omitempty"`

	// deploy hook
	DeployHook string `json:"deploy_hook,omitempty"`

	// deploy url
	DeployURL string `json:"deploy_url,omitempty"`

	// domain aliases
	DomainAliases []string `json:"domain_aliases"`

	// force ssl
	ForceSsl bool `json:"force_ssl,omitempty"`

	// git provider
	GitProvider string `json:"git_provider,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// managed dns
	ManagedDNS bool `json:"managed_dns,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// notification email
	NotificationEmail string `json:"notification_email,omitempty"`

	// password
	Password string `json:"password,omitempty"`

	// plan
	Plan string `json:"plan,omitempty"`

	// published deploy
	PublishedDeploy *Deploy `json:"published_deploy,omitempty"`

	// screenshot url
	ScreenshotURL string `json:"screenshot_url,omitempty"`

	// session id
	SessionID string `json:"session_id,omitempty"`

	// ssl
	Ssl bool `json:"ssl,omitempty"`

	// ssl url
	SslURL string `json:"ssl_url,omitempty"`

	// state
	State string `json:"state,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`

	// url
	URL string `json:"url,omitempty"`

	// user id
	UserID string `json:"user_id,omitempty"`
}

/* polymorph site account_name false */

/* polymorph site account_slug false */

/* polymorph site admin_url false */

/* polymorph site build_settings false */

/* polymorph site capabilities false */

/* polymorph site created_at false */

/* polymorph site custom_domain false */

/* polymorph site deploy_hook false */

/* polymorph site deploy_url false */

/* polymorph site domain_aliases false */

/* polymorph site force_ssl false */

/* polymorph site git_provider false */

/* polymorph site id false */

/* polymorph site managed_dns false */

/* polymorph site name false */

/* polymorph site notification_email false */

/* polymorph site password false */

/* polymorph site plan false */

/* polymorph site published_deploy false */

/* polymorph site screenshot_url false */

/* polymorph site session_id false */

/* polymorph site ssl false */

/* polymorph site ssl_url false */

/* polymorph site state false */

/* polymorph site updated_at false */

/* polymorph site url false */

/* polymorph site user_id false */

// Validate validates this site
func (m *Site) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuildSettings(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDomainAliases(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePublishedDeploy(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Site) validateBuildSettings(formats strfmt.Registry) error {

	if swag.IsZero(m.BuildSettings) { // not required
		return nil
	}

	if m.BuildSettings != nil {

		if err := m.BuildSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("build_settings")
			}
			return err
		}
	}

	return nil
}

func (m *Site) validateDomainAliases(formats strfmt.Registry) error {

	if swag.IsZero(m.DomainAliases) { // not required
		return nil
	}

	return nil
}

func (m *Site) validatePublishedDeploy(formats strfmt.Registry) error {

	if swag.IsZero(m.PublishedDeploy) { // not required
		return nil
	}

	if m.PublishedDeploy != nil {

		if err := m.PublishedDeploy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("published_deploy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Site) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Site) UnmarshalBinary(b []byte) error {
	var res Site
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SiteBuildSettings site build settings
// swagger:model SiteBuildSettings

type SiteBuildSettings struct {

	// allowed branches
	AllowedBranches []string `json:"allowed_branches"`

	// branch
	Branch string `json:"branch,omitempty"`

	// cmd
	Cmd string `json:"cmd,omitempty"`

	// dir
	Dir string `json:"dir,omitempty"`
}

/* polymorph SiteBuildSettings allowed_branches false */

/* polymorph SiteBuildSettings branch false */

/* polymorph SiteBuildSettings cmd false */

/* polymorph SiteBuildSettings dir false */

// Validate validates this site build settings
func (m *SiteBuildSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllowedBranches(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SiteBuildSettings) validateAllowedBranches(formats strfmt.Registry) error {

	if swag.IsZero(m.AllowedBranches) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SiteBuildSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SiteBuildSettings) UnmarshalBinary(b []byte) error {
	var res SiteBuildSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
