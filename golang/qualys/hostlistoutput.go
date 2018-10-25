package qualys

import "encoding/xml"

type HOSTLISTOUTPUT struct {
	XMLName  xml.Name `xml:"HOST_LIST_OUTPUT"`
	Text     string   `xml:",chardata"`
	RESPONSE struct {
		Text     string `xml:",chardata"`
		DATETIME struct {
			Text string `xml:",chardata"`
		} `xml:"DATETIME"`
		HOSTLIST struct {
			Text string `xml:",chardata"`
			HOST []HOST `xml:"HOST"`
		} `xml:"HOST_LIST"`
	} `xml:"RESPONSE"`
}