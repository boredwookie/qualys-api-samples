package qualys

import "encoding/xml"

type HOSTLISTVMDETECTIONOUTPUT struct {
	XMLName  xml.Name `xml:"HOST_LIST_VM_DETECTION_OUTPUT"`
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
		WARNING struct {
			Text string `xml:",chardata"`
			CODE struct {
				Text string `xml:",chardata"`
			} `xml:"CODE"`
			TEXT struct {
				Text string `xml:",chardata"`
			} `xml:"TEXT"`
			URL struct {
				Text string `xml:",chardata"`
			} `xml:"URL"`
		} `xml:"WARNING"`
	} `xml:"RESPONSE"`
} 

