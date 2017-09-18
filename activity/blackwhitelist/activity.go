package blackwhitelist

import (
	"net"
	"strings"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// activityLog is the default logger for the Log Activity
var activityLog = logger.GetLogger("activity-blackwhitelist")

const (
	ivWhiteList = "whitelist"
	ivBlackList = "blacklist"
	ivIPAddress = "ipaddress"

	ovResult = "isOK"
)

func init() {
	activityLog.SetLogLevel(logger.InfoLevel)
}

// blackwhitelist is a utility that checks an ip against a whitelist and a blacklist
// it works as follows:
// The blacklist takes precident ... if a blacklist is supplied then if the input ip
// is found then it's rejected and the result is false
// The whitelist, if supplied is then checked and if the ip is found the OK and the result is OK
//
// origin datatype and a compare mode ... ie "=" or ">"
// inputs : {input1, input2, datatype, comparemode}
// outputs: result (bool)
type blackwhitelist struct {
	metadata *activity.Metadata
}

// NewActivity creates a new AppActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &blackwhitelist{metadata: metadata}
}

// Metadata returns the activity's metadata
func (a *blackwhitelist) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *blackwhitelist) Eval(context activity.Context) (done bool, err error) {

	// Get the runtime values
	wl, _ := context.GetInput(ivWhiteList).(string)
	bl, _ := context.GetInput(ivBlackList).(string)
	ip, _ := context.GetInput(ivIPAddress).(string)
	ip = strings.TrimSpace(ip)

	// check we have a blacklist or a whitelist
	if (wl == "") && (bl == "") {
		activityLog.Info("No Whitelist or blacklist passed")
		return false, nil
	}

	// if we have a blacklist, check against the blacklist and reject if found.
	if bl != "" {
		bls := strings.Split(bl, ",")

		for j := 0; j < len(bls); j++ {
			// get the ipstring and trim space (consider more here)
			listip := strings.TrimSpace(bls[j])
			//log it
			activityLog.Debug("Blacklist value ", listip, " compared to: ", net.ParseIP(ip))
			// Do we have a CIDR format range
			if strings.Contains(listip, "/") { // Handle address ranges
				_, cidrnet, err := net.ParseCIDR(listip)
				// throw error if not valid format
				if err != nil {
					activityLog.Info("Parse Error on: ", listip, " ", err)
					return false, err
				}
				// if found in the range is an error
				if cidrnet.Contains(net.ParseIP(ip)) {
					activityLog.Debug("Found in blacklist")
					context.SetOutput(ovResult, false)
					return true, nil
				}
			} else {
				if net.ParseIP(listip).Equal(net.ParseIP(ip)) {
					// found in bl is not OK
					activityLog.Debug("Found in blacklist")
					context.SetOutput(ovResult, false)
					return true, nil
				}
			}
		}
	}

	// if we have a whitelist we need to check it's on the list
	if wl != "" {
		wls := strings.Split(wl, ",")
		for j := 0; j < len(wls); j++ {
			// get the ipstring and trim space (consider more here)
			listip := strings.TrimSpace(wls[j])
			//log it
			activityLog.Debug("Whitelist value ", listip, " compared to: ", net.ParseIP(ip))
			// Do we have a CIDR format range
			if strings.Contains(listip, "/") { // Handle address ranges
				_, cidrnet, err := net.ParseCIDR(listip)
				// throw error if not valid format
				if err != nil {
					activityLog.Info("Parse Error on: ", listip, " ", err)
					return false, err
				}
				// if found in the range is an error
				if cidrnet.Contains(net.ParseIP(ip)) {
					// found in wl is OK
					context.SetOutput(ovResult, true)
					return true, nil
				}
			} else {
				if net.ParseIP(listip).Equal(net.ParseIP(ip)) {
					// found in wl is OK
					context.SetOutput(ovResult, true)
					return true, nil
				}
			}
		}
		// not found in wl is NOT OK
		activityLog.Debug("Not in whitelist")
		context.SetOutput(ovResult, false)
		return true, nil
	}

	// if we drop through it means we only had a blacklist and our IP wasnt found which is OK
	context.SetOutput(ovResult, true)
	return true, nil
}
