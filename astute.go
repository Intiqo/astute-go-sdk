package astute

import "fmt"

// Client ... A client to interact with the Astute APIs
//
// See https://rhiannon-colleago.astutepayroll.info/rhiannon-colleago/api/ for API reference
type Client interface {
	// SaveTimesheet ... Save a timesheet
	//
	// Creates a new timesheet if none exists for the corresponding time period
	//
	SaveTimesheet(params *SaveTimesheetParams) (SaveTimesheetResponse, error)

	// SubmitTimesheet ... Submit a filled timesheet
	SubmitTimesheet(params *SubmitTimesheetParams) (SaveTimesheetResponse, error)
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
