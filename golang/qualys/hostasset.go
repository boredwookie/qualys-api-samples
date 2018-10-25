package qualys

type HostAsset struct {
	Text string `xml:",chardata"`
	ID   struct {
		Text string `xml:",chardata"`
	} `xml:"id"`
	Name struct {
		Text string `xml:",chardata"`
	} `xml:"name"`
	Created struct {
		Text string `xml:",chardata"`
	} `xml:"created"`
	Modified struct {
		Text string `xml:",chardata"`
	} `xml:"modified"`
	Type struct {
		Text string `xml:",chardata"`
	} `xml:"type"`
	Tags struct {
		Text string `xml:",chardata"`
		List struct {
			Text      string `xml:",chardata"`
			TagSimple []struct {
				Text string `xml:",chardata"`
				ID   struct {
					Text string `xml:",chardata"`
				} `xml:"id"`
				Name struct {
					Text string `xml:",chardata"`
				} `xml:"name"`
			} `xml:"TagSimple"`
		} `xml:"list"`
	} `xml:"tags"`
	SourceInfo struct {
		Text string `xml:",chardata"`
		List struct {
			Text                 string `xml:",chardata"`
			Ec2AssetSourceSimple struct {
				Text    string `xml:",chardata"`
				AssetId struct {
					Text string `xml:",chardata"`
				} `xml:"assetId"`
				Type struct {
					Text string `xml:",chardata"`
				} `xml:"type"`
				FirstDiscovered struct {
					Text string `xml:",chardata"`
				} `xml:"firstDiscovered"`
				LastUpdated struct {
					Text string `xml:",chardata"`
				} `xml:"lastUpdated"`
				Ec2InstanceTags struct {
					Text string `xml:",chardata"`
					Tags struct {
						Text string `xml:",chardata"`
						List struct {
							Text    string `xml:",chardata"`
							EC2Tags []struct {
								Text string `xml:",chardata"`
								Key  struct {
									Text string `xml:",chardata"`
								} `xml:"key"`
								Value struct {
									Text string `xml:",chardata"`
								} `xml:"value"`
							} `xml:"EC2Tags"`
						} `xml:"list"`
					} `xml:"tags"`
				} `xml:"ec2InstanceTags"`
				ReservationId struct {
					Text string `xml:",chardata"`
				} `xml:"reservationId"`
				AvailabilityZone struct {
					Text string `xml:",chardata"`
				} `xml:"availabilityZone"`
				PrivateDnsName struct {
					Text string `xml:",chardata"`
				} `xml:"privateDnsName"`
				InstanceId struct {
					Text string `xml:",chardata"`
				} `xml:"instanceId"`
				InstanceType struct {
					Text string `xml:",chardata"`
				} `xml:"instanceType"`
				CreatedDate struct {
					Text string `xml:",chardata"`
				} `xml:"createdDate"`
				InstanceState struct {
					Text string `xml:",chardata"`
				} `xml:"instanceState"`
				GroupId struct {
					Text string `xml:",chardata"`
				} `xml:"groupId"`
				GroupName struct {
					Text string `xml:",chardata"`
				} `xml:"groupName"`
				SpotInstance struct {
					Text string `xml:",chardata"`
				} `xml:"spotInstance"`
				AccountId struct {
					Text string `xml:",chardata"`
				} `xml:"accountId"`
				SubnetId struct {
					Text string `xml:",chardata"`
				} `xml:"subnetId"`
				VpcId struct {
					Text string `xml:",chardata"`
				} `xml:"vpcId"`
				Region struct {
					Text string `xml:",chardata"`
				} `xml:"region"`
				Zone struct {
					Text string `xml:",chardata"`
				} `xml:"zone"`
				ImageId struct {
					Text string `xml:",chardata"`
				} `xml:"imageId"`
				PrivateIpAddress struct {
					Text string `xml:",chardata"`
				} `xml:"privateIpAddress"`
				MonitoringEnabled struct {
					Text string `xml:",chardata"`
				} `xml:"monitoringEnabled"`
			} `xml:"Ec2AssetSourceSimple"`
		} `xml:"list"`
	} `xml:"sourceInfo"`
	QwebHostId struct {
		Text string `xml:",chardata"`
	} `xml:"qwebHostId"`
	Os struct {
		Text string `xml:",chardata"`
	} `xml:"os"`
	DnsHostName struct {
		Text string `xml:",chardata"`
	} `xml:"dnsHostName"`
	Address struct {
		Text string `xml:",chardata"`
	} `xml:"address"`
	TrackingMethod struct {
		Text string `xml:",chardata"`
	} `xml:"trackingMethod"`
	NetworkInterface struct {
		Text string `xml:",chardata"`
		List struct {
			Text               string `xml:",chardata"`
			HostAssetInterface []struct {
				Text     string `xml:",chardata"`
				Hostname struct {
					Text string `xml:",chardata"`
				} `xml:"hostname"`
				InterfaceId struct {
					Text string `xml:",chardata"`
				} `xml:"interfaceId"`
				InterfaceName struct {
					Text string `xml:",chardata"`
				} `xml:"interfaceName"`
				Type struct {
					Text string `xml:",chardata"`
				} `xml:"type"`
				Address struct {
					Text string `xml:",chardata"`
				} `xml:"address"`
			} `xml:"HostAssetInterface"`
		} `xml:"list"`
	} `xml:"networkInterface"`
	IsDockerHost struct {
		Text string `xml:",chardata"`
	} `xml:"isDockerHost"`
	LastVulnScan struct {
		Text string `xml:",chardata"`
	} `xml:"lastVulnScan"`
	OpenPort struct {
		Text string `xml:",chardata"`
		List struct {
			Text              string `xml:",chardata"`
			HostAssetOpenPort []struct {
				Text string `xml:",chardata"`
				Port struct {
					Text string `xml:",chardata"`
				} `xml:"port"`
				Protocol struct {
					Text string `xml:",chardata"`
				} `xml:"protocol"`
				ServiceId struct {
					Text string `xml:",chardata"`
				} `xml:"serviceId"`
				ServiceName struct {
					Text string `xml:",chardata"`
				} `xml:"serviceName"`
			} `xml:"HostAssetOpenPort"`
		} `xml:"list"`
	} `xml:"openPort"`
	Vuln struct {
		Text string `xml:",chardata"`
		List struct {
			Text          string `xml:",chardata"`
			HostAssetVuln []struct {
				Text string `xml:",chardata"`
				Qid  struct {
					Text string `xml:",chardata"`
				} `xml:"qid"`
				HostInstanceVulnId struct {
					Text string `xml:",chardata"`
				} `xml:"hostInstanceVulnId"`
				FirstFound struct {
					Text string `xml:",chardata"`
				} `xml:"firstFound"`
				LastFound struct {
					Text string `xml:",chardata"`
				} `xml:"lastFound"`
			} `xml:"HostAssetVuln"`
		} `xml:"list"`
	} `xml:"vuln"`
}
