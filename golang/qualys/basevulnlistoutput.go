package qualys

import "encoding/xml"

type KNOWLEDGEBASEVULNLISTOUTPUT struct {
	XMLName  xml.Name `xml:"KNOWLEDGE_BASE_VULN_LIST_OUTPUT"`
	Text     string   `xml:",chardata"`
	RESPONSE struct {
		Text     string `xml:",chardata"`
		DATETIME struct {
			Text string `xml:",chardata"`
		} `xml:"DATETIME"`
		VULNLIST struct {
			Text string `xml:",chardata"`
			VULN []Vuln `xml:"VULN"`
		} `xml:"VULN_LIST"`
	} `xml:"RESPONSE"`
} 

