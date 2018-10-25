package qualys

import "encoding/xml"

type ServiceResponse struct {
	XMLName                   xml.Name `xml:"ServiceResponse"`
	Text                      string   `xml:",chardata"`
	Xsi                       string   `xml:"xsi,attr"`
	NoNamespaceSchemaLocation string   `xml:"noNamespaceSchemaLocation,attr"`
	ResponseCode              struct {
		Text string `xml:",chardata"`
	} `xml:"responseCode"`
	Count struct {
		Text string `xml:",chardata"`
	} `xml:"count"`
	HasMoreRecords struct {
		Text string `xml:",chardata"`
	} `xml:"hasMoreRecords"`
	Data struct {
		Text      string `xml:",chardata"`
		HostAsset []HostAsset `xml:"HostAsset"`
	} `xml:"data"`
}



type ServiceResponseSingleAsset struct {
	XMLName                   xml.Name `xml:"ServiceResponse"`
	Text                      string   `xml:",chardata"`
	Xsi                       string   `xml:"xsi,attr"`
	NoNamespaceSchemaLocation string   `xml:"noNamespaceSchemaLocation,attr"`
	ResponseCode              struct {
		Text string `xml:",chardata"`
	} `xml:"responseCode"`
	Count struct {
		Text string `xml:",chardata"`
	} `xml:"count"`
	HasMoreRecords struct {
		Text string `xml:",chardata"`
	} `xml:"hasMoreRecords"`
	Data struct {
		Text      string `xml:",chardata"`
		HostAsset HostAsset `xml:"HostAsset"`
	} `xml:"data"`
}

