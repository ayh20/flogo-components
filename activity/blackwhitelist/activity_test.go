package blackwhitelist

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

func TestEvalSubtract(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	fmt.Println("#######   Testing blacklist")

	//test1
	tc.SetInput("whitelist", "")
	tc.SetInput("blacklist", "")
	tc.SetInput("ipaddress", "")
	act.Eval(tc)

	if tc.GetOutput("isOK") != nil {
		t.Fail()
	}
	res := tc.GetOutput("isOK")
	fmt.Println("Result should be: nil  is:", res)

	//test1
	tc.SetInput("whitelist", "")
	tc.SetInput("blacklist", "1.1.1.1,2.2.2.2,3.3.3.3")
	tc.SetInput("ipaddress", "4.4.4.4")
	act.Eval(tc)

	if tc.GetOutput("isOK") == false {
		t.Fail()
	}
	res = tc.GetOutput("isOK")
	fmt.Println("Result should be: True  is:", res)

	//test1
	tc.SetInput("whitelist", "")
	tc.SetInput("blacklist", "1.1.1.1,2.2.2.2,3.3.3.3")
	tc.SetInput("ipaddress", "1.1.1.1 ")
	act.Eval(tc)

	if tc.GetOutput("isOK") == true {
		t.Fail()
	}
	res = tc.GetOutput("isOK")
	fmt.Println("Result should be: false  is:", res)

	fmt.Println("#######   Testing whitelist")

	//test1
	tc.SetInput("whitelist", "1.1.1.1,2.2.2.2,3.3.3.3")
	tc.SetInput("blacklist", "")
	tc.SetInput("ipaddress", " 4.4.4.4")
	act.Eval(tc)

	if tc.GetOutput("isOK") == true {
		t.Fail()
	}
	res = tc.GetOutput("isOK")
	fmt.Println("Result should be: false  is:", res)

	//test1
	tc.SetInput("whitelist", " 1.1.1.1,2.2.2.2, 3.3.3.3")
	tc.SetInput("blacklist", "")
	tc.SetInput("ipaddress", "1.1.1.1")
	act.Eval(tc)

	if tc.GetOutput("isOK") == false {
		t.Fail()
	}
	res = tc.GetOutput("isOK")
	fmt.Println("Result should be: True  is:", res)

	//test1
	tc.SetInput("whitelist", "")
	tc.SetInput("blacklist", " 1.1.1.0/24 ,2.2.2.2 , 3.3.3.3")
	tc.SetInput("ipaddress", "1.1.1.1")
	act.Eval(tc)

	if tc.GetOutput("isOK") == true {
		t.Fail()
	}
	res = tc.GetOutput("isOK")
	fmt.Println("Result should be: False  is:", res)

	// masherylist v homeip
	fmt.Println("#######   Testing Mash Whitelist")
	// 64.94.14.0/27,64.94.228.128/28,216.52.39.0/24,216.52.244.96/27,216.133.249.0/24,23.23.79.128/25,107.22.159.192/28,54.82.131.0/25 ,75.101.137.168,75.101.142.168,75.101.146.168,75.101.141.43,75.101.129.141,174.129.251.74,174.129.251.80,50.18.151.192/28,50.112.119.192/28,54.193.255.0/25 ,204.236.130.149 ,204.236.130.201,204.236.130.207,176.34.239.192/28,54.247.111.192/26 ,54.93.255.128/27 ,54.252.79.192/27,54.251.88.0/27,69.71.111.140,69.71.111.141,207.126.59.91,207.126.59.94,165.254.103.205,165.254.103.203,70.34.228.92,70.34.228.93,4.53.108.203,4.53.108.205,208.72.116.130,208.72.116.131,200.85.152.87,200.85.152.89,200.155.158.42,200.155.158.43,187.45.223.91,187.45.223.93,165.254.103.205,165.254.103.203,213.130.49.203,213.130.49.205,213.198.94.38,213.198.94.39,212.72.53.203,212.72.53.205,87.236.193.132,87.236.193.137,93.94.105.60,93.94.105.75,103.19.90.28,103.19.90.29,103.15.105.253,103.15.105.254,103.248.191.19,123.100.230.144,123.100.230.146,123.100.230.148,123.100.230.150,110.50.254.174,110.50.254.177
	tc.SetInput("whitelist", "64.94.14.0/27,64.94.228.128/28,216.52.39.0/24,216.52.244.96/27,216.133.249.0/24,23.23.79.128/25,107.22.159.192/28,54.82.131.0/25 ,75.101.137.168,75.101.142.168,75.101.146.168,75.101.141.43,75.101.129.141,174.129.251.74,174.129.251.80,50.18.151.192/28,50.112.119.192/28,54.193.255.0/25 ,204.236.130.149 ,204.236.130.201,204.236.130.207,176.34.239.192/28,54.247.111.192/26 ,54.93.255.128/27 ,54.252.79.192/27,54.251.88.0/27,69.71.111.140,69.71.111.141,207.126.59.91,207.126.59.94,165.254.103.205,165.254.103.203,70.34.228.92,70.34.228.93,4.53.108.203,4.53.108.205,208.72.116.130,208.72.116.131,200.85.152.87,200.85.152.89,200.155.158.42,200.155.158.43,187.45.223.91,187.45.223.93,165.254.103.205,165.254.103.203,213.130.49.203,213.130.49.205,213.198.94.38,213.198.94.39,212.72.53.203,212.72.53.205,87.236.193.132,87.236.193.137,93.94.105.60,93.94.105.75,103.19.90.28,103.19.90.29,103.15.105.253,103.15.105.254,103.248.191.19,123.100.230.144,123.100.230.146,123.100.230.148,123.100.230.150,110.50.254.174,110.50.254.177")
	tc.SetInput("blacklist", "")
	tc.SetInput("ipaddress", "80.229.148.252")
	act.Eval(tc)

	if tc.GetOutput("isOK") == true {
		t.Fail()
	}
	res = tc.GetOutput("isOK")
	fmt.Println("Result should be: False  is:", res)

	tc.SetInput("whitelist", "64.94.14.0/27,64.94.228.128/28,216.52.39.0/24,216.52.244.96/27,216.133.249.0/24,23.23.79.128/25,107.22.159.192/28,54.82.131.0/25 ,75.101.137.168,75.101.142.168,75.101.146.168,75.101.141.43,75.101.129.141,174.129.251.74,174.129.251.80,50.18.151.192/28,50.112.119.192/28,54.193.255.0/25 ,204.236.130.149 ,204.236.130.201,204.236.130.207,176.34.239.192/28,54.247.111.192/26 ,54.93.255.128/27 ,54.252.79.192/27,54.251.88.0/27,69.71.111.140,69.71.111.141,207.126.59.91,207.126.59.94,165.254.103.205,165.254.103.203,70.34.228.92,70.34.228.93,4.53.108.203,4.53.108.205,208.72.116.130,208.72.116.131,200.85.152.87,200.85.152.89,200.155.158.42,200.155.158.43,187.45.223.91,187.45.223.93,165.254.103.205,165.254.103.203,213.130.49.203,213.130.49.205,213.198.94.38,213.198.94.39,212.72.53.203,212.72.53.205,87.236.193.132,87.236.193.137,93.94.105.60,93.94.105.75,103.19.90.28,103.19.90.29,103.15.105.253,103.15.105.254,103.248.191.19,123.100.230.144,123.100.230.146,123.100.230.148,123.100.230.150,110.50.254.174,110.50.254.177")
	tc.SetInput("blacklist", "")
	tc.SetInput("ipaddress", "107.22.159.195")
	act.Eval(tc)

	if tc.GetOutput("isOK") == false {
		t.Fail()
	}
	res = tc.GetOutput("isOK")
	fmt.Println("Result should be: true  is:", res)
}
