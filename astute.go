package astute

import "fmt"

// Client ... A client to interact with the Astute APIs
//
// See https://rhiannon-colleago.astutepayroll.info/rhiannon-colleago/api/ for API reference
type Client interface {
	// QueryTimesheetByJob ... Query a timesheet by it's corresponding job
	//
	// You can use this to query a timesheet record using the following supported parameters—
	//
	// UID.
	//
	//
	QueryTimesheetByJob(params QueryTimesheetParams) (QueryTimesheetResponse, error)
	// QueryTimesheetById ... Query a timesheet by an ID.
	QueryTimesheetById(id string) (QueryTimesheetResponse, error)
	// QueryUser ... Query a user
	//
	// You can use this to query a user record using the following supported parameters
	//
	// - Job Code
	//
	//
	QueryUser(params QueryUserParams) (QueryUserResponse, error)
	// SaveTimesheet ... Save a timesheet
	//
	// Creates a new timesheet if none exists for the corresponding time period
	//
	SaveTimesheet(params *SaveTimesheetParams) (SaveTimesheetResponse, error)
	// AddTimesheetShift ... Add a shift to an existing timesheet
	//
	// See https://api.astutepayroll.com/webservice/documentation/#type_TimesheetAddShift
	AddTimesheetShift(params *AddTimesheetShiftParams) (AddTimesheetShiftResponse, error)
	// DeleteTimesheetShift ... Delete a shift previously added via AddTimesheetShift
	//
	// Identified by TS_SID (shift ID returned by AddTimesheetShift) and TSID (timesheet ID).
	//
	// See https://api.astutepayroll.com/webservice/documentation/#type_TimesheetDeleteShift
	DeleteTimesheetShift(params *DeleteTimesheetShiftParams) (DeleteTimesheetShiftResponse, error)
	// ListTimesheetShifts ... List every shift currently recorded on a timesheet, returning
	// the TS_SID of each. Useful when re-emitting a group through TimesheetAddShift after
	// the timesheet was previously populated via TimesheetSave's weekday-tag form: those
	// day-level entries are otherwise un-addressable for delete (no TS_SID is returned by
	// TimesheetSave). Pair with DeleteTimesheetShift to wipe a timesheet before re-adding.
	//
	// See https://api.astutepayroll.com/webservice/documentation/#method_TimesheetQueryShift
	ListTimesheetShifts(tsid string) ([]string, error)
}

func NewClient(params AuthParams) (Client, error) {
	var c astuteClient
	if params.ApiUrl == "" || params.ApiKey == "" || params.ApiUsername == "" || params.ApiPassword == "" {
		return c, fmt.Errorf("api url, key, username & password must not be empty")
	}
	b := NewBackend()
	c = astuteClient{
		B:          b,
		AuthParams: params,
	}
	return c, nil
}
