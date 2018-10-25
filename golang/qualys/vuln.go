package qualys

type Vuln struct {
	Text string `xml:",chardata"`
	QID  struct {
		Text string `xml:",chardata"`
	} `xml:"QID"`
	VULNTYPE struct {
		Text string `xml:",chardata"`
	} `xml:"VULN_TYPE"`
	SEVERITYLEVEL struct {
		Text string `xml:",chardata"`
	} `xml:"SEVERITY_LEVEL"`
	TITLE struct {
		Text string `xml:",chardata"`
	} `xml:"TITLE"`
	CATEGORY struct {
		Text string `xml:",chardata"`
	} `xml:"CATEGORY"`
	LASTSERVICEMODIFICATIONDATETIME struct {
		Text string `xml:",chardata"`
	} `xml:"LAST_SERVICE_MODIFICATION_DATETIME"`
	PUBLISHEDDATETIME struct {
		Text string `xml:",chardata"`
	} `xml:"PUBLISHED_DATETIME"`
	PATCHABLE struct {
		Text string `xml:",chardata"`
	} `xml:"PATCHABLE"`
	DIAGNOSIS struct {
		Text string `xml:",chardata"`
	} `xml:"DIAGNOSIS"`
	PCIFLAG struct {
		Text string `xml:",chardata"`
	} `xml:"PCI_FLAG"`
	DISCOVERY struct {
		Text   string `xml:",chardata"`
		REMOTE struct {
			Text string `xml:",chardata"`
		} `xml:"REMOTE"`
		AUTHTYPELIST struct {
			Text     string `xml:",chardata"`
			AUTHTYPE []struct {
				Text string `xml:",chardata"`
			} `xml:"AUTH_TYPE"`
		} `xml:"AUTH_TYPE_LIST"`
		ADDITIONALINFO struct {
			Text string `xml:",chardata"`
		} `xml:"ADDITIONAL_INFO"`
	} `xml:"DISCOVERY"`
	CONSEQUENCE struct {
		Text string `xml:",chardata"`
	} `xml:"CONSEQUENCE"`
	SOLUTION struct {
		Text string `xml:",chardata"`
	} `xml:"SOLUTION"`
	BUGTRAQLIST struct {
		Text    string `xml:",chardata"`
		BUGTRAQ []struct {
			Text string `xml:",chardata"`
			ID   struct {
				Text string `xml:",chardata"`
			} `xml:"ID"`
			URL struct {
				Text string `xml:",chardata"`
			} `xml:"URL"`
		} `xml:"BUGTRAQ"`
	} `xml:"BUGTRAQ_LIST"`
	SOFTWARELIST struct {
		Text     string `xml:",chardata"`
		SOFTWARE []struct {
			Text    string `xml:",chardata"`
			PRODUCT struct {
				Text string `xml:",chardata"`
			} `xml:"PRODUCT"`
			VENDOR struct {
				Text string `xml:",chardata"`
			} `xml:"VENDOR"`
		} `xml:"SOFTWARE"`
	} `xml:"SOFTWARE_LIST"`
	VENDORREFERENCELIST struct {
		Text            string `xml:",chardata"`
		VENDORREFERENCE []struct {
			Text string `xml:",chardata"`
			ID   struct {
				Text string `xml:",chardata"`
			} `xml:"ID"`
			URL struct {
				Text string `xml:",chardata"`
			} `xml:"URL"`
		} `xml:"VENDOR_REFERENCE"`
	} `xml:"VENDOR_REFERENCE_LIST"`
	CVELIST struct {
		Text string `xml:",chardata"`
		CVE  []struct {
			Text string `xml:",chardata"`
			ID   struct {
				Text string `xml:",chardata"`
			} `xml:"ID"`
			URL struct {
				Text string `xml:",chardata"`
			} `xml:"URL"`
		} `xml:"CVE"`
	} `xml:"CVE_LIST"`
	CORRELATION struct {
		Text     string `xml:",chardata"`
		EXPLOITS struct {
			Text     string `xml:",chardata"`
			EXPLTSRC []struct {
				Text    string `xml:",chardata"`
				SRCNAME struct {
					Text string `xml:",chardata"`
				} `xml:"SRC_NAME"`
				EXPLTLIST struct {
					Text  string `xml:",chardata"`
					EXPLT []struct {
						Text string `xml:",chardata"`
						REF  struct {
							Text string `xml:",chardata"`
						} `xml:"REF"`
						DESC struct {
							Text string `xml:",chardata"`
						} `xml:"DESC"`
						LINK struct {
							Text string `xml:",chardata"`
						} `xml:"LINK"`
					} `xml:"EXPLT"`
				} `xml:"EXPLT_LIST"`
			} `xml:"EXPLT_SRC"`
		} `xml:"EXPLOITS"`
	} `xml:"CORRELATION"`
}
