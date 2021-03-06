package config

import (
	"regexp"
	"time"
)

type AlertmanagerConfig struct {
	Name        string
	URI         string
	ExternalURI string `yaml:"external_uri" koanf:"external_uri"`
	Timeout     time.Duration
	Proxy       bool
	ReadOnly    bool `yaml:"readonly"`
	TLS         struct {
		CA                 string
		Cert               string
		Key                string
		InsecureSkipVerify bool `yaml:"insecureSkipVerify" koanf:"insecureSkipVerify"`
	}
	Headers map[string]string
}

type LinkDetectRules struct {
	Regex       string `yaml:"regex"`
	URITemplate string `yaml:"uriTemplate" koanf:"uriTemplate"`
}

type CustomLabelColor struct {
	Value         string         `yaml:"value"`
	ValueRegex    string         `yaml:"value_re" koanf:"value_re"`
	CompiledRegex *regexp.Regexp `yaml:"-"`
	Color         string         `yaml:"color"`
}

type CustomLabelColors map[string][]CustomLabelColor

type configSchema struct {
	Alertmanager struct {
		Interval    time.Duration
		Servers     []AlertmanagerConfig
		Name        string        `yaml:"-" koanf:"name"`
		Timeout     time.Duration `yaml:"-" koanf:"timeout"`
		URI         string        `yaml:"-" koanf:"uri"`
		ExternalURI string        `yaml:"-" koanf:"external_uri"`
		Proxy       bool          `yaml:"-" koanf:"proxy"`
		ReadOnly    bool          `yaml:"-" koanf:"readonly"`
	}
	AlertAcknowledgement struct {
		Enabled       bool
		Duration      time.Duration
		Author        string
		CommentPrefix string `yaml:"commentPrefix" koanf:"commentPrefix"`
	} `yaml:"alertAcknowledgement" koanf:"alertAcknowledgement"`
	Annotations struct {
		Default struct {
			Hidden bool
		}
		Hidden  []string
		Visible []string
		Keep    []string
		Strip   []string
	}
	Custom struct {
		CSS string
		JS  string
	}
	Debug   bool
	Filters struct {
		Default []string
	}
	Grid struct {
		Sorting struct {
			Order        string
			Reverse      bool
			Label        string
			CustomValues struct {
				Labels map[string]map[string]string
			} `yaml:"customValues" koanf:"customValues"`
		}
	} `yaml:"grid"`
	Karma struct {
		Name string
	}
	Labels struct {
		Keep  []string
		Strip []string
		Color struct {
			Custom CustomLabelColors
			Static []string
			Unique []string
		}
	}
	Listen struct {
		Address string
		Port    int
		Prefix  string
	}
	Log struct {
		Config    bool
		Level     string
		Format    string
		Timestamp bool
	}
	Receivers struct {
		Keep  []string
		Strip []string
	}
	Sentry struct {
		Private string
		Public  string
	}
	Silences struct {
		Comments struct {
			LinkDetect struct {
				Rules []LinkDetectRules `yaml:"rules"`
			} `yaml:"linkDetect" koanf:"linkDetect"`
		} `yaml:"comments"`
	} `yaml:"silences"`
	SilenceForm struct {
		Author struct {
			PopulateFromHeader struct {
				Header     string `yaml:"header" koanf:"header"`
				ValueRegex string `yaml:"value_re" koanf:"value_re"`
			} `yaml:"populate_from_header" koanf:"populate_from_header"`
		} `yaml:"author"`
		Strip struct {
			Labels []string
		}
	} `yaml:"silenceForm" koanf:"silenceForm"`
	UI struct {
		Refresh             time.Duration
		HideFiltersWhenIdle bool   `yaml:"hideFiltersWhenIdle" koanf:"hideFiltersWhenIdle"`
		ColorTitlebar       bool   `yaml:"colorTitlebar" koanf:"colorTitlebar"`
		Theme               string `yaml:"theme" koanf:"theme"`
		MinimalGroupWidth   int    `yaml:"minimalGroupWidth" koanf:"minimalGroupWidth"`
		AlertsPerGroup      int    `yaml:"alertsPerGroup" koanf:"alertsPerGroup"`
		CollapseGroups      string `yaml:"collapseGroups" koanf:"collapseGroups"`
	}
}
