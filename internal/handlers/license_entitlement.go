package handlers

import (
	"crypto/ed25519"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"os"
	"strings"
	"time"

	"stone-ocean-web/internal/store"
)

const licenseSigningPrivateKeyEnv = "LICENSE_SIGNING_PRIVATE_KEY"

type licenseSigner struct {
	privateKey ed25519.PrivateKey
	keyID      string
}

type licenseEntitlement struct {
	Schema            string              `json:"schema"`
	LicenseID         uint                `json:"licenseId"`
	ActivationID      uint                `json:"activationId"`
	DeviceIDHash      string              `json:"deviceIdHash"`
	Status            store.LicenseStatus `json:"status"`
	PlanCode          string              `json:"planCode"`
	PlanName          string              `json:"planName"`
	Kind              store.LicenseKind   `json:"kind"`
	Features          []string            `json:"features"`
	MaxActivations    int                 `json:"maxActivations"`
	IssuedAt          string              `json:"issuedAt"`
	ExpiresAt         *string             `json:"expiresAt"`
	EntitlementIssued string              `json:"entitlementIssuedAt"`
	NextCheckAt       string              `json:"nextCheckAt"`
	OfflineGraceUntil string              `json:"offlineGraceUntil"`
	SigningKeyID      string              `json:"signingKeyId,omitempty"`
}

func newLicenseSignerFromEnv() *licenseSigner {
	raw := strings.TrimSpace(os.Getenv(licenseSigningPrivateKeyEnv))
	if raw == "" {
		return nil
	}
	keyBytes, err := base64.StdEncoding.DecodeString(raw)
	if err != nil {
		return nil
	}
	switch len(keyBytes) {
	case ed25519.SeedSize:
		return newLicenseSigner(ed25519.NewKeyFromSeed(keyBytes))
	case ed25519.PrivateKeySize:
		return newLicenseSigner(ed25519.PrivateKey(keyBytes))
	default:
		return nil
	}
}

func newLicenseSigner(privateKey ed25519.PrivateKey) *licenseSigner {
	publicKey, ok := privateKey.Public().(ed25519.PublicKey)
	if !ok {
		return nil
	}
	sum := sha256.Sum256(publicKey)
	return &licenseSigner{
		privateKey: privateKey,
		keyID:      hex.EncodeToString(sum[:8]),
	}
}

func (s *licenseSigner) KeyID() string {
	if s == nil {
		return ""
	}
	return s.keyID
}

func (s *licenseSigner) Sign(entitlement licenseEntitlement) (string, error) {
	entitlement.SigningKeyID = s.keyID
	payload, err := json.Marshal(entitlement)
	if err != nil {
		return "", err
	}
	signature := ed25519.Sign(s.privateKey, payload)
	return base64.RawURLEncoding.EncodeToString(payload) + "." + base64.RawURLEncoding.EncodeToString(signature), nil
}

func newLicenseEntitlement(result *store.LicenseActivationResult, now, nextCheckAt, offlineGraceUntil time.Time) licenseEntitlement {
	var expiresAt *string
	if result.License.ExpiresAt != nil {
		formatted := result.License.ExpiresAt.Format(time.RFC3339)
		expiresAt = &formatted
	}

	return licenseEntitlement{
		Schema:            "recoverease-license-v1",
		LicenseID:         result.License.ID,
		ActivationID:      result.Activation.ID,
		DeviceIDHash:      result.Activation.DeviceIDHash,
		Status:            result.License.Status,
		PlanCode:          result.License.LicensePlan.Code,
		PlanName:          result.License.LicensePlan.Name,
		Kind:              result.License.LicensePlan.Kind,
		Features:          licenseFeatures(result.License.LicensePlan.Kind),
		MaxActivations:    result.License.LicensePlan.MaxActivations,
		IssuedAt:          result.License.IssuedAt.Format(time.RFC3339),
		ExpiresAt:         expiresAt,
		EntitlementIssued: now.Format(time.RFC3339),
		NextCheckAt:       nextCheckAt.Format(time.RFC3339),
		OfflineGraceUntil: offlineGraceUntil.Format(time.RFC3339),
	}
}

func licenseFeatures(kind store.LicenseKind) []string {
	features := []string{
		"smart_recovery",
		"advanced_recovery",
		"gpu_acceleration",
	}
	if kind == store.LicenseKindLifetime {
		features = append(features, "lifetime_license")
	}
	return features
}
