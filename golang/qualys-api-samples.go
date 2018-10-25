package golang

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"qualys-api-samples/golang/qualys"
	"strings"
)

var qauth64 string
var qApi *string
var ips *string
var assets *string
var hosts *string
var authFile *string
var authCli *string
func main() {
	//
	// QualysGuard API example: pull Vulnerability details from hosts scanned with Qualys
	//

	// Potentially useful guides to the Qualys API
	//
	// https://www.qualys.com/docs/qualys-api-vmpc-user-guide.pdf
	// https://community.qualys.com/docs/DOC-4523-qualys-api-client-examples#jive_content_id_Go_Language_Example
	// https://www.qualys.com/docs/qualys-asset-management-tagging-api-v2-user-guide.pdf
	// https://community.qualys.com/thread/18542-get-list-of-all-applications-on-authenticated-hosts-via-api
	//

	//
	// The Qualys structs were generated with Zek (https://github.com/miku/zek) using raw XML captured from exploring
	// the Qualys API manually:
	//
	//	`cat raw_output.xml | ./zek > golangstructname.go`
	//

	// Parse the CLI arguments, check for validity
	qApi = flag.String("ApiUrl", "m", "Point to the Qualys API endpoint (Review Qualys documentation to see which endpoint you should use)")
	ips = flag.String("HostIps", "", "Enter a comma-delimited list of host IP addresses that have been scanned by Qualys")
	assets = flag.String("AssetIds", "", "Enter a comma-delimited list of Qualys Asset IDs")
	hosts = flag.String("HostIds", "", "HostIDs passed as a comma-delimited list of Qualys host IDs")
	authFile = flag.String("AuthFile", "", "Auth file that contains username/password authentication. Should be in the format of `USERNAME:PASSWORD`\nTo generate a mock/example file, pass in a value of GENERATE_SAMPLE like so:\n\t-AuthFile=GENERATE_SAMPLE\n\t(a file named `credentials` will be created")
	authCli = flag.String("AuthCLI", "", "Can be used in place of the 'AuthFile' if you'd rather specify the username and password at the CLI.\nEx:\n\t-AuthCLI=USERNAME:PASSWORD\n\n\tQualys uses HTTP Basic authentication, which is why we take the credentials in this format")
	flag.Parse()

	if *authFile != "" && *authCli != ""{
		fmt.Println("Specify either AuthFile -or- AuthCLI. Both cannot be specified at once")
		os.Exit(127)
	}

	if (*ips != "" && *assets != "") || (*ips != "" && *hosts != "") || (*hosts != "" && *assets != ""){
		fmt.Println("Specify only ONE of the following CLI options:\n\tHostIps\n\tHostIds\n\tAssetIds")
		os.Exit(126)
	}

	if *authFile == "GENERATE_SAMPLE"{
		ioutil.WriteFile("credentials", []byte("username:password"), 0600)
		os.Exit(0)
	}

	// Parse authentication
	var authBytes []byte
	if *authFile != ""{
		authFileBytes, authFileBytesErr := ioutil.ReadFile(*authFile)
		if authFileBytesErr != nil{
			panic(authFileBytesErr)
		}

		authBytes = authFileBytes
	} else {
		authBytes = []byte(*authCli)
	}
	qauth64 = base64.StdEncoding.EncodeToString(authBytes)

	//
	// Get the list of Asset IDs
	var assetIdentifiers []string
	if *ips != ""{
		// Get the list of hosts by IP address
		hostsOutput := searchHostsByIps(*ips)
		var hostIds string
		for _, host := range hostsOutput.RESPONSE.HOSTLIST.HOST{
			hostIds += "," + host.ID.Text
		}
		hostIds = strings.TrimPrefix(hostIds, ",")

		// Convert the hosts into AssetIDs
		serviceResponse := searchHostsQps(hostIds)
		for _, hostAsset := range serviceResponse.Data.HostAsset{
			assetIdentifiers = append(assetIdentifiers, hostAsset.ID.Text)
		}

	} else if *hosts != ""{
		// Convert the hostIDs into AssetIDs
		serviceResponse := searchHostsQps(*hosts)
		for _, hostAsset := range serviceResponse.Data.HostAsset{
			assetIdentifiers = append(assetIdentifiers, hostAsset.ID.Text)
		}

	} else if *assets != ""{
		// Already have the Asset identifiers, just need to get them into a slice
		assetIdentifiers = strings.Split(*assets, ",")
	}

	//
	// Get detailed host information using the asset identifiers
	//		Build the list of QIDs
	var qids string
	for _, ai := range assetIdentifiers{
		hostDetail := hostAssetDetails(ai)
		for _, v := range hostDetail.Data.HostAsset.Vuln.List.HostAssetVuln{
			qids += v.Qid.Text + ","
			fmt.Println(v)
		}
	}
	qids = strings.TrimSuffix(qids, ",")

	//
	// Get the vulnerability details based on the list of QIDs
	//		Print the details
	details := vulnerabilityDetails(qids)
	for _, vDetails := range details.RESPONSE.VULNLIST.VULN{
		fmt.Println(vDetails)
	}

	// Print the raw list of qids
	fmt.Println("===")
	fmt.Println(qids)
}

//
// Take a comma-delimited list of Qualys QIDs and get the vulnerability details
func vulnerabilityDetails(qids string) qualys.KNOWLEDGEBASEVULNLISTOUTPUT{
	vulnDetailsRaw := qApiCallXml("GET", qauth64, "/api/2.0/fo/knowledge_base/vuln/?action=list&details=All&ids="+qids)
	
	var vulnlistdetails qualys.KNOWLEDGEBASEVULNLISTOUTPUT
	xml.Unmarshal(vulnDetailsRaw, &vulnlistdetails)
	
	return vulnlistdetails
}

//
// Returns detailed host information for a specific host
func hostAssetDetails(assetIdentifier string) qualys.ServiceResponseSingleAsset{
	hostAssetDetailsRaw := qApiCallXml("GET", qauth64, "/qps/rest/2.0/get/am/hostasset/"+assetIdentifier)

	var hostAssetDetailsServiceResponse qualys.ServiceResponseSingleAsset
	xml.Unmarshal(hostAssetDetailsRaw, &hostAssetDetailsServiceResponse)

	return hostAssetDetailsServiceResponse
}

//
// Get the HostListOutput from Qualys by searching via IP Addresses
func searchHostsByIps(ipAddresses string) qualys.HOSTLISTOUTPUT{
	// Remove any potential whitespace
	strings.Replace(ipAddresses, " ", "", -1)
	// Ensure there are no leading or trailing commas
	strings.Trim(ipAddresses, ",")

	// Get the raw host data back from Qualys
	hostsRawBytes := qApiCallXml("GET", qauth64, "/api/2.0/fo/asset/host/?action=list&ips="+ipAddresses)
	var hostsListOutput qualys.HOSTLISTOUTPUT
	xml.Unmarshal(hostsRawBytes, &hostsListOutput)

	return hostsListOutput
}

//
// Get a ServiceResponse which includes host details
func searchHostsQps(hostIds string) qualys.ServiceResponse{
	searchResults := qApiCallXml("POST", qauth64, "/qps/rest/2.0/search/am/hostasset", `<ServiceRequest>
  <filters>
    <Criteria field="qwebHostId" operator="IN">`+ hostIds + `</Criteria>
  </filters>
</ServiceRequest>`)

	var response qualys.ServiceResponse
	xml.Unmarshal(searchResults, &response)

	return response
}

//
// Bare-bones function that calls the specified API endpoint and returns the results as a raw byte array
func qApiCallXml(method string, auth64 string, url string, postbody ...string) []byte{
	// Get an HTTP Request
	qApiUrl := *qApi + url
	var qHttpReq *http.Request
	var qHttpReqErr error
	if len(postbody) > 0 {
		// Only add a postbody if one is supplied
		qHttpReq, qHttpReqErr = http.NewRequest(method, qApiUrl, bytes.NewBuffer([]byte(postbody[0])))
	} else {
		qHttpReq, qHttpReqErr = http.NewRequest(method, qApiUrl, nil)
	}
	if qHttpReqErr != nil{
		panic(qHttpReqErr)
	}

	// Populate the headers
	qHttpReq.Header.Add("X-requested-with", "qualys-integrator")
	qHttpReq.Header.Add("Authorization", "Basic " + auth64)

	qHttpResp, qHttpRespErr := http.DefaultClient.Do(qHttpReq)
	if qHttpRespErr != nil{
		panic(qHttpRespErr)
	}
	qHttpRespBytes, qHttpRespStrErr := ioutil.ReadAll(qHttpResp.Body)
	if qHttpRespStrErr != nil{
		panic(qHttpRespStrErr)
	}

	return qHttpRespBytes
}
