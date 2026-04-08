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
	// QueryTimesheetShift ... Query shifts for a given timesheet by TSID
	//
	// Queries shifts for a given timesheet by TSID. You can use this to query shifts for a timesheet record using the following supported parameters
	//
	QueryTimesheetShift(params QueryTimesheetShiftParams) (QueryTimesheetShiftResponse, error)
	// AddTimesheetShift ... Add a shift to an existing timesheet
	//
	// Adds a shift. You can use this to add a shift to an existing timesheet record using the following supported parameters
	//
	AddTimesheetShift(params *AddTimesheetShiftParams) (AddTimesheetShiftResponse, error)
	// DeleteTimesheetShift ... Delete a shift from a timesheet by TS_SID
	//
	// Deletes a shift from a timesheet by TS_SID. You can use this to delete a shift from a timesheet record using the following supported parameters
	//
	DeleteTimesheetShift(params DeleteTimesheetShiftParams) error
	// ApproveTimesheet ... Approve a timesheet. Once approved, a timesheet cannot be updated.
	//
	// Approves a timesheet. Once approved, a timesheet cannot be updated. You can use this to approve a timesheet record using the following supported parameters
	//
	ApproveTimesheet(params ApproveTimesheetParams) error
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
