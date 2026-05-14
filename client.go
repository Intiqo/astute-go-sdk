package astute

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

type astuteClient struct {
	B          Backend
	AuthParams AuthParams
}

func (c astuteClient) QueryUser(params QueryUserParams) (QueryUserResponse, error) {
	var res QueryUserResponse

	reqTemplate := strings.TrimSpace(
		`<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:tns="urn:tsoIntegrator" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
<soap:Body>
<q1:UserQuery xmlns:q1="urn:UserQuery">
  <tns:userGet>
    <api_key>{{.ApiKey}}</api_key>
    <api_username>{{.ApiUsername}}</api_username>
    <api_password>{{.ApiPassword}}</api_password>
    <query>job_code like '%{{.JobCode}}%'</query>
  </tns:userGet>
</q1:UserQuery>
</soap:Body>
</soap:Envelope>`,
	)

	templateData := struct {
		AuthParams
		QueryUserParams
	}{
		AuthParams:      c.AuthParams,
		QueryUserParams: params,
	}

	resp, err := c.B.Call(c.AuthParams.ApiUrl, "UserQuery", "urn:UserQuery", reqTemplate, templateData)
	if err != nil {
		return res, err
	}

	if resp.Code != http.StatusOK {
		result, err := ParseResponse(resp.Data, faultResponse{})
		if err != nil {
			return res, nil
		}
		resText := result.Body.Fault.Faultstring.Text
		return res, fmt.Errorf(resText)
	}

	result, err := ParseResponse(resp.Data, queryUserXmlResponse{})
	if err != nil {
		return res, nil
	}

	xmlUsers := result.Body.UserQueryResponse.ParmsOut.Results.Text
	res, err = ParseResponse([]byte(xmlUsers), QueryUserResponse{})
	if err != nil {
		return res, nil
	}

	return res, nil
}

func (c astuteClient) QueryTimesheetByJob(params QueryTimesheetParams) (QueryTimesheetResponse, error) {
	var res QueryTimesheetResponse

	reqTemplate := strings.TrimSpace(
		`<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:tns="urn:tsoIntegrator" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
<soap:Body>
<q1:TimesheetQuery xmlns:q1="urn:TimesheetQuery">
  <tns:userGet>
    <api_key>{{.ApiKey}}</api_key>
    <api_username>{{.ApiUsername}}</api_username>
    <api_password>{{.ApiPassword}}</api_password>
    <query>UID = '{{.UID}}'</query>
  </tns:userGet>
</q1:TimesheetQuery>
</soap:Body>
</soap:Envelope>`,
	)

	templateData := struct {
		AuthParams
		QueryTimesheetParams
	}{
		AuthParams:           c.AuthParams,
		QueryTimesheetParams: params,
	}

	resp, err := c.B.Call(c.AuthParams.ApiUrl, "TimesheetQuery", "urn:TimesheetQuery", reqTemplate, templateData)
	if err != nil {
		return res, err
	}

	if resp.Code != http.StatusOK {
		result, err := ParseResponse(resp.Data, faultResponse{})
		if err != nil {
			return res, nil
		}
		resText := result.Body.Fault.Faultstring.Text
		return res, fmt.Errorf(resText)
	}

	result, err := ParseResponse(resp.Data, queryTimesheetXmlResponse{})
	if err != nil {
		return res, nil
	}

	xmlUsers := result.Body.TimesheetQueryResponse.ParmsOut.Results.Text
	res, err = ParseResponse([]byte(xmlUsers), QueryTimesheetResponse{})
	if err != nil {
		return res, nil
	}

	return res, nil
}

func (c astuteClient) QueryTimesheetById(id string) (QueryTimesheetResponse, error) {
	var res QueryTimesheetResponse

	reqTemplate := strings.TrimSpace(
		`<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:tns="urn:tsoIntegrator" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
<soap:Body>
<q1:TimesheetQuery xmlns:q1="urn:TimesheetQuery">
  <tns:userGet>
    <api_key>{{.ApiKey}}</api_key>
    <api_username>{{.ApiUsername}}</api_username>
    <api_password>{{.ApiPassword}}</api_password>
    <query>TSID = '{{.TSID}}'</query>
  </tns:userGet>
</q1:TimesheetQuery>
</soap:Body>
</soap:Envelope>`,
	)

	templateData := struct {
		AuthParams
		TSID string
	}{
		AuthParams: c.AuthParams,
		TSID:       id,
	}

	resp, err := c.B.Call(c.AuthParams.ApiUrl, "TimesheetQuery", "urn:TimesheetQuery", reqTemplate, templateData)
	if err != nil {
		return res, err
	}

	if resp.Code != http.StatusOK {
		result, err := ParseResponse(resp.Data, faultResponse{})
		if err != nil {
			return res, nil
		}
		resText := result.Body.Fault.Faultstring.Text
		return res, fmt.Errorf(resText)
	}

	result, err := ParseResponse(resp.Data, queryTimesheetXmlResponse{})
	if err != nil {
		return res, nil
	}

	xmlUsers := result.Body.TimesheetQueryResponse.ParmsOut.Results.Text
	res, err = ParseResponse([]byte(xmlUsers), QueryTimesheetResponse{})
	if err != nil {
		return res, nil
	}

	return res, nil
}

func (c astuteClient) SaveTimesheet(params *SaveTimesheetParams) (SaveTimesheetResponse, error) {
	var res SaveTimesheetResponse
	var reqTemplate string
	var resp *ClientResponse
	var err error

	// If the timesheet is marked as "Did not work", use the appropriate template
	if params.DidNotWork {
		templateData := struct {
			AuthParams
			UserParams
			TSID             string
			ApiTransactionId string
		}{
			AuthParams:       c.AuthParams,
			UserParams:       params.UserParams,
			TSID:             params.TSID,
			ApiTransactionId: uuid.New().String(),
		}
		reqTemplate = strings.TrimSpace(
			`<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:tns="urn:tsoIntegrator" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
<soap:Body>
<q1:TimesheetSave xmlns:q1="urn:TimesheetSave">
  <tns:timesheetSave>
    <api_key>{{.ApiKey}}</api_key>
    <api_username>{{.ApiUsername}}</api_username>
    <api_password>{{.ApiPassword}}</api_password>
    <api_transaction_id>{{.ApiTransactionId}}</api_transaction_id>
    <UID>{{.UID}}</UID>
    <user_id>{{.UserId}}</user_id>
		<TSID>{{.TSID}}</TSID>
		<did_not_work>1</did_not_work>
  </tns:timesheetSave>
</q1:TimesheetSave>
</soap:Body>
</soap:Envelope>`,
		)
		resp, err = c.B.Call(c.AuthParams.ApiUrl, "TimesheetSave", "urn:TimesheetSave", reqTemplate, templateData)
		if err != nil {
			return res, err
		}
	} else {
		reqTemplate = strings.TrimSpace(
			`<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:tns="urn:tsoIntegrator" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
<soap:Body>
<q1:TimesheetSave xmlns:q1="urn:TimesheetSave">
  <tns:timesheetSave>
    <api_key>{{.ApiKey}}</api_key>
    <api_username>{{.ApiUsername}}</api_username>
    <api_password>{{.ApiPassword}}</api_password>
    <api_transaction_id>{{.ApiTransactionId}}</api_transaction_id>
    <UID>{{.UID}}</UID>
    <user_id>{{.UserId}}</user_id>
	<TSID>{{.TSID}}</TSID>
    <date>{{.TimesheetDate}}</date>
		{{range $key, $value := .Days}}
			<{{.WeekdayTag}}_start>{{.StartTime}}</{{.WeekdayTag}}_start>
			<{{.WeekdayTag}}_finish>{{.EndTime}}</{{.WeekdayTag}}_finish>
			<{{.WeekdayTag}}_break>{{.BreakTime}}</{{.WeekdayTag}}_break>
			<{{.WeekdayTag}}_notes>{{.Notes}}</{{.WeekdayTag}}_notes>
		{{end}}
		{{if .Notes}}<notes>{{.Notes}}</notes>{{end}}
		{{if .Submit}}
			<complete>{{.SubmissionTime}}</complete>
			{{if .TriggerApprovalEmail}}
				<trigger_approval_email>1</trigger_approval_email>
			{{end}}
		{{end}}
  </tns:timesheetSave>
</q1:TimesheetSave>
</soap:Body>
</soap:Envelope>`,
		)

		days := make([]SaveTimesheetDayTemplateParams, 0)
		// Days may be empty when the caller only wants to submit (Submit=true) or update notes
		// on a timesheet whose entries were already populated by AddTimesheetShift. In that case
		// we use SubmissionTime (or now) for the <date> field instead of indexing Days[0].
		var tsStartTime time.Time
		if len(params.Days) > 0 {
			tsStartTime = params.Days[0].StartTime
		}
		for _, day := range params.Days {
			// Astute strictly requires that the start and end times be represented through 4 characters
			// in the format HHMM. If the time is less than 4 characters, we need to pad it with 0s.
			//
			// See https://api.astutepayroll.com/webservice/documentation/#type_timesheetSave for more details
			st := strings.Split(day.StartTime.Format("15:04:05"), ":")
			startTime := fmt.Sprintf("%s%s", st[0], st[1])
			et := strings.Split(day.EndTime.Format("15:04:05"), ":")
			endTime := fmt.Sprintf("%s%s", et[0], et[1])

			breakTime := "0000"

			// Astute strictly requires that the break time be of 4 characters
			//
			// See https://api.astutepayroll.com/webservice/documentation/#type_timesheetSave for more details
			if len(day.BreakTime) == 1 {
				breakTime = fmt.Sprintf("000%s", day.BreakTime)
			} else if len(day.BreakTime) == 2 {
				breakTime = fmt.Sprintf("00%s", day.BreakTime)
			} else if len(day.BreakTime) == 3 {
				breakTime = fmt.Sprintf("0%s", day.BreakTime)
			} else if len(day.BreakTime) == 4 {
				breakTime = day.BreakTime
			}

			buf := new(bytes.Buffer)
			_ = xml.EscapeText(buf, []byte(day.Notes))
			days = append(days, SaveTimesheetDayTemplateParams{
				WeekdayTag: getWeekdayTemplateForTime(day.StartTime),
				StartTime:  startTime,
				EndTime:    endTime,
				BreakTime:  breakTime,
				Notes:      buf.String(),
			})

			if day.StartTime.Before(tsStartTime) {
				tsStartTime = day.StartTime
			}
		}

		submissionTime := time.Now()
		if params.Submit && !params.SubmissionTime.IsZero() {
			submissionTime = params.SubmissionTime
		}

		if tsStartTime.IsZero() {
			tsStartTime = submissionTime
		}

		templateData := struct {
			AuthParams
			UserParams
			TSID                 string
			ApiTransactionId     string
			TimesheetDate        string
			Days                 []SaveTimesheetDayTemplateParams
			Submit               bool
			SubmissionTime       string
			Notes                string
			TriggerApprovalEmail bool
		}{
			AuthParams:           c.AuthParams,
			UserParams:           params.UserParams,
			TSID:                 params.TSID,
			ApiTransactionId:     uuid.New().String(),
			TimesheetDate:        tsStartTime.Format("2006-01-02"),
			Days:                 days,
			Submit:               params.Submit,
			SubmissionTime:       submissionTime.String(),
			Notes:                params.Notes,
			TriggerApprovalEmail: params.TriggerApprovalEmail,
		}
		resp, err = c.B.Call(c.AuthParams.ApiUrl, "TimesheetSave", "urn:TimesheetSave", reqTemplate, templateData)
		if err != nil {
			return res, err
		}
	}

	if resp.Code != http.StatusOK {
		result, err := ParseResponse(resp.Data, faultResponse{})
		if err != nil {
			return res, nil
		}
		resText := result.Body.Fault.Faultstring.Text
		return res, fmt.Errorf(resText)
	}

	result, err := ParseResponse(resp.Data, saveTimesheetXmlResponse{})
	if err != nil {
		return res, nil
	}

	resText := result.Body.TimesheetSaveResponse.ParmsOut.Results.Text

	if !strings.Contains(resText, "TSID:") {
		return res, fmt.Errorf(resText)
	}

	tsId := resText[21:]

	res = SaveTimesheetResponse{
		TimesheetId: tsId,
	}

	return res, nil
}

func (c astuteClient) AddTimesheetShift(params *AddTimesheetShiftParams) (AddTimesheetShiftResponse, error) {
	var res AddTimesheetShiftResponse

	reqTemplate := strings.TrimSpace(
		`<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:tns="urn:tsoIntegrator" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
<soap:Body>
<q1:TimesheetAddShift xmlns:q1="urn:TimesheetAddShift">
  <tns:timesheetAddShift>
    <api_key>{{.ApiKey}}</api_key>
    <api_username>{{.ApiUsername}}</api_username>
    <api_password>{{.ApiPassword}}</api_password>
    <api_transaction_id>{{.ApiTransactionId}}</api_transaction_id>
    <TSID>{{.TSID}}</TSID>
    {{if .UID}}<UID>{{.UID}}</UID>{{end}}
    {{if .UserId}}<user_id>{{.UserId}}</user_id>{{end}}
    <date>{{.Date}}</date>
    <start>{{.StartTime}}</start>
    {{if .EndTime}}<finish>{{.EndTime}}</finish>{{end}}
    {{if .BreakStart}}<break_start>{{.BreakStart}}</break_start>{{end}}
    {{if .BreakFinish}}<break_finish>{{.BreakFinish}}</break_finish>{{end}}
    {{if .Break2Start}}<break2_start>{{.Break2Start}}</break2_start>{{end}}
    {{if .Break2Finish}}<break2_finish>{{.Break2Finish}}</break2_finish>{{end}}
    {{if .Break3Start}}<break3_start>{{.Break3Start}}</break3_start>{{end}}
    {{if .Break3Finish}}<break3_finish>{{.Break3Finish}}</break3_finish>{{end}}
    {{if .BreakTime}}<break>{{.BreakTime}}</break>{{end}}
    {{if .Notes}}<notes>{{.Notes}}</notes>{{end}}
    {{if .JobCode}}<job_code>{{.JobCode}}</job_code>{{end}}
  </tns:timesheetAddShift>
</q1:TimesheetAddShift>
</soap:Body>
</soap:Envelope>`,
	)

	notesBuf := new(bytes.Buffer)
	_ = xml.EscapeText(notesBuf, []byte(params.Notes))

	templateData := struct {
		AuthParams
		UserParams
		TSID             string
		ApiTransactionId string
		Date             string
		StartTime        string
		EndTime          string
		BreakStart       string
		BreakFinish      string
		Break2Start      string
		Break2Finish     string
		Break3Start      string
		Break3Finish     string
		BreakTime        string
		Notes            string
		JobCode          string
	}{
		AuthParams:       c.AuthParams,
		UserParams:       params.UserParams,
		TSID:             params.TSID,
		ApiTransactionId: uuid.New().String(),
		Date:             params.Date.Format("2006-01-02"),
		StartTime:        formatTimeHHMM(params.StartTime),
		EndTime:          formatTimeHHMM(params.EndTime),
		BreakStart:       formatTimeHHMM(params.BreakStart),
		BreakFinish:      formatTimeHHMM(params.BreakFinish),
		Break2Start:      formatTimeHHMM(params.Break2Start),
		Break2Finish:     formatTimeHHMM(params.Break2Finish),
		Break3Start:      formatTimeHHMM(params.Break3Start),
		Break3Finish:     formatTimeHHMM(params.Break3Finish),
		BreakTime:        padBreakTime(params.BreakTime),
		Notes:            notesBuf.String(),
		JobCode:          params.JobCode,
	}

	resp, err := c.B.Call(c.AuthParams.ApiUrl, "TimesheetAddShift", "urn:TimesheetAddShift", reqTemplate, templateData)
	if err != nil {
		return res, err
	}

	if resp.Code != http.StatusOK {
		result, err := ParseResponse(resp.Data, faultResponse{})
		if err != nil {
			return res, nil
		}
		return res, fmt.Errorf("%s", result.Body.Fault.Faultstring.Text)
	}

	result, err := ParseResponse(resp.Data, addTimesheetShiftXmlResponse{})
	if err != nil {
		return res, nil
	}

	resultCode := strings.TrimSpace(result.Body.TimesheetAddShiftResponse.ParmsOut.ResultCode.Text)
	resultsText := strings.TrimSpace(result.Body.TimesheetAddShiftResponse.ParmsOut.Results.Text)

	// Astute returns a 200 even when the operation fails (e.g. "This shift conflicts with
	// another shift…"). Surface those as errors so callers don't proceed as if the shift saved.
	// We check both ResultCode (non-"S" indicates failure) and known failure phrases in Results.
	if isAddTimesheetShiftFailure(resultCode, resultsText) {
		return res, fmt.Errorf("%s", resultsText)
	}

	res = AddTimesheetShiftResponse{
		Result: resultsText,
	}

	return res, nil
}

func isAddTimesheetShiftFailure(resultCode, resultsText string) bool {
	if resultCode != "" && !strings.EqualFold(resultCode, "S") {
		return true
	}
	lower := strings.ToLower(resultsText)
	failurePhrases := []string{
		"cannot be saved",
		"please check your data",
		"conflicts with",
	}
	for _, p := range failurePhrases {
		if strings.Contains(lower, p) {
			return true
		}
	}
	return false
}

// Formats a time.Time as a 4-character HHMM string, or empty if t is zero.
// Astute expects shift start/finish times in this packed format.
func formatTimeHHMM(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format("1504")
}

// Left-pads a break-duration string to exactly 4 characters as required by Astute.
// An empty input is preserved as empty so the template can omit the tag.
func padBreakTime(bt string) string {
	switch len(bt) {
	case 0:
		return ""
	case 1:
		return "000" + bt
	case 2:
		return "00" + bt
	case 3:
		return "0" + bt
	case 4:
		return bt
	default:
		return "0000"
	}
}

// Helps in identifying the weekday for the given time
func getWeekdayTemplateForTime(startTime time.Time) string {
	weekDay := startTime.Weekday()
	weekDayTag := ""
	switch weekDay {
	case 0:
		weekDayTag = "sun"
	case 1:
		weekDayTag = "mon"
	case 2:
		weekDayTag = "tue"
	case 3:
		weekDayTag = "wed"
	case 4:
		weekDayTag = "thu"
	case 5:
		weekDayTag = "fri"
	case 6:
		weekDayTag = "sat"
	default:
	}

	return weekDayTag
}
