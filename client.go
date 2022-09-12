package astute

import (
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

func (c astuteClient) SaveTimesheet(params *SaveTimesheetParams) (SaveTimesheetResponse, error) {
	var res SaveTimesheetResponse

	reqTemplate := strings.TrimSpace(
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
    <date>{{.TimesheetDate}}</date>
    <{{.WeekdayTag}}_start>{{.StartTime}}</{{.WeekdayTag}}_start>
    <{{.WeekdayTag}}_finish>{{.EndTime}}</{{.WeekdayTag}}_finish>
    <{{.WeekdayTag}}_break>{{.BreakTime}}</{{.WeekdayTag}}_break>
    <{{.WeekdayTag}}_notes>{{.Notes}}</{{.WeekdayTag}}_notes>
  </tns:timesheetSave>
</q1:TimesheetSave>
</soap:Body>
</soap:Envelope>`,
	)

	st := strings.Split(params.StartTime.Format("15:04:05"), ":")
	et := strings.Split(params.EndTime.Format("15:04:05"), ":")

	templateData := struct {
		AuthParams
		UserParams
		ApiTransactionId string
		WeekdayTag       string
		TimesheetDate    string
		StartTime        string
		EndTime          string
		BreakTime        string
		Notes            string
	}{
		AuthParams:       c.AuthParams,
		UserParams:       params.UserParams,
		ApiTransactionId: uuid.New().String(),
		WeekdayTag:       getWeekdayTemplateForTime(params.StartTime),
		TimesheetDate:    params.StartTime.Format("2006-01-02"),
		StartTime:        fmt.Sprintf("%s%s", st[0], st[1]),
		EndTime:          fmt.Sprintf("%s%s", et[0], et[1]),
		BreakTime:        fmt.Sprintf("00%s", params.BreakTime),
		Notes:            params.Notes,
	}

	resp, err := c.B.Call(c.AuthParams.ApiUrl, "TimesheetSave", "urn:TimesheetSave", reqTemplate, templateData)
	if err != nil {
		return res, err
	}

	if resp.Code != http.StatusOK {
		return res, fmt.Errorf("response is not OK")
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

func (c astuteClient) SubmitTimesheet(params *SubmitTimesheetParams) (SaveTimesheetResponse, error) {
	var res SaveTimesheetResponse

	reqTemplate := strings.TrimSpace(
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
    <date>{{.TimesheetDate}}</date>
    <complete>{{.SubmissionTime}}</complete>
  </tns:timesheetSave>
</q1:TimesheetSave>
</soap:Body>
</soap:Envelope>`,
	)

	templateData := struct {
		AuthParams
		UserParams
		ApiTransactionId string
		TimesheetDate    string
		SubmissionTime   string
	}{
		AuthParams:       c.AuthParams,
		UserParams:       params.UserParams,
		ApiTransactionId: uuid.New().String(),
		TimesheetDate:    params.StartTime.Format("2006-01-02"),
		SubmissionTime:   params.SubmissionTime.String(),
	}

	resp, err := c.B.Call(c.AuthParams.ApiUrl, "TimesheetSave", "urn:TimesheetSave", reqTemplate, templateData)
	if err != nil {
		return res, err
	}

	if resp.Code != http.StatusOK {
		return res, fmt.Errorf("response is not OK")
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
