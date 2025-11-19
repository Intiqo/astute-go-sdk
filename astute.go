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
