package qualys

type HOST struct {
	Text string `xml:",chardata"`
	ID   struct {
		Text string `xml:",chardata"`
	} `xml:"ID"`
	IP struct {
		Text string `xml:",chardata"`
	} `xml:"IP"`
	TRACKINGMETHOD struct {
		Text string `xml:",chardata"`
	} `xml:"TRACKING_METHOD"`
	OS struct {
		Text string `xml:",chardata"`
	} `xml:"OS"`
	DNS struct {
		Text string `xml:",chardata"`
	} `xml:"DNS"`
	NETBIOS struct {
		Text string `xml:",chardata"`
	} `xml:"NETBIOS"`
	QGHOSTID struct {
		Text string `xml:",chardata"`
	} `xml:"QG_HOSTID"`
	LASTSCANDATETIME struct {
		Text string `xml:",chardata"`
	} `xml:"LAST_SCAN_DATETIME"`
	LASTVMSCANNEDDATE struct {
		Text string `xml:",chardata"`
	} `xml:"LAST_VM_SCANNED_DATE"`
	LASTVMSCANNEDDURATION struct {
		Text string `xml:",chardata"`
	} `xml:"LAST_VM_SCANNED_DURATION"`
	LASTVMAUTHSCANNEDDATE struct {
		Text string `xml:",chardata"`
	} `xml:"LAST_VM_AUTH_SCANNED_DATE"`
	LASTVMAUTHSCANNEDDURATION struct {
		Text string `xml:",chardata"`
	} `xml:"LAST_VM_AUTH_SCANNED_DURATION"`
	DETECTIONLIST struct {
		Text      string `xml:",chardata"`
		DETECTION []struct {
			Text string `xml:",chardata"`
			QID  struct {
				Text string `xml:",chardata"`
			} `xml:"QID"`
			TYPE struct {
				Text string `xml:",chardata"`
			} `xml:"TYPE"`
			SEVERITY struct {
				Text string `xml:",chardata"`
			} `xml:"SEVERITY"`
			PORT struct {
				Text string `xml:",chardata"`
			} `xml:"PORT"`
			PROTOCOL struct {
				Text string `xml:",chardata"`
			} `xml:"PROTOCOL"`
			SSL struct {
				Text string `xml:",chardata"`
			} `xml:"SSL"`
			RESULTS struct {
				Text string `xml:",chardata"`
			} `xml:"RESULTS"`
			STATUS struct {
				Text string `xml:",chardata"`
			} `xml:"STATUS"`
			FIRSTFOUNDDATETIME struct {
				Text string `xml:",chardata"`
			} `xml:"FIRST_FOUND_DATETIME"`
			LASTFOUNDDATETIME struct {
				Text string `xml:",chardata"`
			} `xml:"LAST_FOUND_DATETIME"`
			TIMESFOUND struct {
				Text string `xml:",chardata"`
			} `xml:"TIMES_FOUND"`
			LASTTESTDATETIME struct {
				Text string `xml:",chardata"`
			} `xml:"LAST_TEST_DATETIME"`
			LASTUPDATEDATETIME struct {
				Text string `xml:",chardata"`
			} `xml:"LAST_UPDATE_DATETIME"`
			ISIGNORED struct {
				Text string `xml:",chardata"`
			} `xml:"IS_IGNORED"`
			ISDISABLED struct {
				Text string `xml:",chardata"`
			} `xml:"IS_DISABLED"`
			LASTPROCESSEDDATETIME struct {
				Text string `xml:",chardata"`
			} `xml:"LAST_PROCESSED_DATETIME"`
			AFFECTRUNNINGKERNEL struct {
				Text string `xml:",chardata"`
			} `xml:"AFFECT_RUNNING_KERNEL"`
			LASTFIXEDDATETIME struct {
				Text string `xml:",chardata"`
			} `xml:"LAST_FIXED_DATETIME"`
		} `xml:"DETECTION"`
	} `xml:"DETECTION_LIST"`
}