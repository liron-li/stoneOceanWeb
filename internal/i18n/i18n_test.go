package i18n

import "testing"

func TestSupportedLanguagesHaveLocalizedVisibleText(t *testing.T) {
	criticalKeys := []string{
		"hero.eyebrow",
		"nav.formats",
		"formats.eyebrow",
		"app.nav.recovery",
		"app.currentTask",
		"labels.recommended",
		"labels.bestValue",
		"advantages.eyebrow",
		"workflow.eyebrow",
		"pricing.eyebrow",
		"download.eyebrow",
	}

	for _, lang := range languageDefinitions {
		if !Supported(lang.Code) {
			t.Fatalf("language %q is listed but not supported", lang.Code)
		}

		if lang.Code == "en" {
			continue
		}

		for _, key := range criticalKeys {
			if T(lang.Code, key) == en[key] {
				t.Fatalf("language %q still falls back to English for %q", lang.Code, key)
			}
		}
	}
}

func TestMatchAcceptLanguage(t *testing.T) {
	tests := []struct {
		name   string
		header string
		want   string
	}{
		{
			name:   "matches regional Chinese",
			header: "zh-CN,zh;q=0.9,en;q=0.8",
			want:   "zh",
		},
		{
			name:   "matches regional Spanish",
			header: "es-MX,es;q=0.8,en;q=0.6",
			want:   "es",
		},
		{
			name:   "honors q priority",
			header: "fr;q=0.6,de;q=0.9,en;q=0.7",
			want:   "de",
		},
		{
			name:   "falls back to English",
			header: "it-IT,it;q=0.9,pl;q=0.8",
			want:   "en",
		},
		{
			name:   "empty header falls back to English",
			header: "",
			want:   "en",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MatchAcceptLanguage(tt.header); got != tt.want {
				t.Fatalf("MatchAcceptLanguage(%q) = %q, want %q", tt.header, got, tt.want)
			}
		})
	}
}

func TestLanguagePathsAreStable(t *testing.T) {
	if DefaultLocale != "en" {
		t.Fatalf("DefaultLocale = %q, want en", DefaultLocale)
	}

	if Path("zh") != "/zh" {
		t.Fatalf("Path(%q) = %q, want /zh", "zh", Path("zh"))
	}

	if Path("en") != "/en" {
		t.Fatalf("Path(%q) = %q, want /en", "en", Path("en"))
	}

	if CheckoutPath("zh") != "/zh/checkout" {
		t.Fatalf("CheckoutPath(%q) = %q, want /zh/checkout", "zh", CheckoutPath("zh"))
	}

	if CheckoutPath("en") != "/en/checkout" {
		t.Fatalf("CheckoutPath(%q) = %q, want /en/checkout", "en", CheckoutPath("en"))
	}

	if CheckoutPath("unknown") != "/en/checkout" {
		t.Fatalf("CheckoutPath(%q) = %q, want /en/checkout", "unknown", CheckoutPath("unknown"))
	}

	if LicenseRecoveryPath("zh") != "/zh/license-recovery" {
		t.Fatalf("LicenseRecoveryPath(%q) = %q, want /zh/license-recovery", "zh", LicenseRecoveryPath("zh"))
	}

	if LicenseRecoveryPath("en") != "/en/license-recovery" {
		t.Fatalf("LicenseRecoveryPath(%q) = %q, want /en/license-recovery", "en", LicenseRecoveryPath("en"))
	}

	if LicenseRecoveryPath("unknown") != "/en/license-recovery" {
		t.Fatalf("LicenseRecoveryPath(%q) = %q, want /en/license-recovery", "unknown", LicenseRecoveryPath("unknown"))
	}
}
