package astute

import (
	"encoding/xml"
	"time"
)

type AstuteEntity interface {
	QueryTimesheetResponse | queryTimesheetUserResponse | saveTimesheetXmlResponse
}

type UserParams struct {
	UID    string
	UserId string
}

type SaveTimesheetParams struct {
	UserParams
	StartTime time.Time
	EndTime   time.Time
	BreakTime string
	Notes     string
}

type SaveTimesheetResponse struct {
	TimesheetId string
}

type SubmitTimesheetParams struct {
	UserParams
	StartTime      time.Time
	SubmissionTime time.Time
}

type saveTimesheetXmlResponse struct {
	XMLName       xml.Name `xml:"Envelope"`
	Text          string   `xml:",chardata"`
	EncodingStyle string   `xml:"encodingStyle,attr"`
	SOAPENV       string   `xml:"SOAP-ENV,attr"`
	Xsd           string   `xml:"xsd,attr"`
	Xsi           string   `xml:"xsi,attr"`
	SOAPENC       string   `xml:"SOAP-ENC,attr"`
	Tns           string   `xml:"tns,attr"`
	Body          struct {
		Text                  string `xml:",chardata"`
		TimesheetSaveResponse struct {
			Text     string `xml:",chardata"`
			Ns1      string `xml:"ns1,attr"`
			ParmsOut struct {
				Text       string `xml:",chardata"`
				Type       string `xml:"type,attr"`
				ResultCode struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr"`
				} `xml:"ResultCode"`
				Results struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr"`
				} `xml:"Results"`
			} `xml:"parms_out"`
		} `xml:"TimesheetSaveResponse"`
	} `xml:"Body"`
}

type QueryTimesheetResponse struct {
	XMLName       xml.Name `xml:"Envelope"`
	Text          string   `xml:",chardata"`
	EncodingStyle string   `xml:"encodingStyle,attr"`
	SOAPENV       string   `xml:"SOAP-ENV,attr"`
	Xsd           string   `xml:"xsd,attr"`
	Xsi           string   `xml:"xsi,attr"`
	SOAPENC       string   `xml:"SOAP-ENC,attr"`
	Tns           string   `xml:"tns,attr"`
	Body          struct {
		Text                   string `xml:",chardata"`
		TimesheetQueryResponse struct {
			Text     string `xml:",chardata"`
			Ns1      string `xml:"ns1,attr"`
			ParmsOut struct {
				Text       string `xml:",chardata"`
				Type       string `xml:"type,attr"`
				ResultCode struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr"`
				} `xml:"ResultCode"`
				Results struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr"`
					queryTimesheetUserResponse
				} `xml:"Results"`
			} `xml:"parms_out"`
		} `xml:"TimesheetQueryResponse"`
	} `xml:"Body"`
}

type queryTimesheetUserResponse struct {
	XMLName xml.Name `xml:"users"`
	Text    string   `xml:",chardata"`
	User    struct {
		Text                           string `xml:",chardata"`
		TSID                           string `xml:"TSID"`
		UID                            string `xml:"UID"`
		UserID                         string `xml:"user_id"`
		Date                           string `xml:"date"`
		DateTo                         string `xml:"date_to"`
		Sequence                       string `xml:"sequence"`
		Latest                         string `xml:"latest"`
		MonWork                        string `xml:"mon_work"`
		MonStart                       string `xml:"mon_start"`
		MonFinish                      string `xml:"mon_finish"`
		MonBreak                       string `xml:"mon_break"`
		MonLeaveType                   string `xml:"mon_leave_type"`
		MonLeave                       string `xml:"mon_leave"`
		MonLeaveHours                  string `xml:"mon_leave_hours"`
		MonNotes                       string `xml:"mon_notes"`
		MonPMPCID                      string `xml:"mon_PM_PCID"`
		TueWork                        string `xml:"tue_work"`
		TueStart                       string `xml:"tue_start"`
		TueFinish                      string `xml:"tue_finish"`
		TueBreak                       string `xml:"tue_break"`
		TueLeaveType                   string `xml:"tue_leave_type"`
		TueLeave                       string `xml:"tue_leave"`
		TueLeaveHours                  string `xml:"tue_leave_hours"`
		TueNotes                       string `xml:"tue_notes"`
		TuePMPCID                      string `xml:"tue_PM_PCID"`
		WedWork                        string `xml:"wed_work"`
		WedStart                       string `xml:"wed_start"`
		WedFinish                      string `xml:"wed_finish"`
		WedBreak                       string `xml:"wed_break"`
		WedLeaveType                   string `xml:"wed_leave_type"`
		WedLeave                       string `xml:"wed_leave"`
		WedLeaveHours                  string `xml:"wed_leave_hours"`
		WedNotes                       string `xml:"wed_notes"`
		WedPMPCID                      string `xml:"wed_PM_PCID"`
		ThuWork                        string `xml:"thu_work"`
		ThuStart                       string `xml:"thu_start"`
		ThuFinish                      string `xml:"thu_finish"`
		ThuBreak                       string `xml:"thu_break"`
		ThuLeaveType                   string `xml:"thu_leave_type"`
		ThuLeave                       string `xml:"thu_leave"`
		ThuLeaveHours                  string `xml:"thu_leave_hours"`
		ThuNotes                       string `xml:"thu_notes"`
		ThuPMPCID                      string `xml:"thu_PM_PCID"`
		FriWork                        string `xml:"fri_work"`
		FriStart                       string `xml:"fri_start"`
		FriFinish                      string `xml:"fri_finish"`
		FriBreak                       string `xml:"fri_break"`
		FriLeaveType                   string `xml:"fri_leave_type"`
		FriLeave                       string `xml:"fri_leave"`
		FriLeaveHours                  string `xml:"fri_leave_hours"`
		FriNotes                       string `xml:"fri_notes"`
		FriPMPCID                      string `xml:"fri_PM_PCID"`
		SatWork                        string `xml:"sat_work"`
		SatStart                       string `xml:"sat_start"`
		SatFinish                      string `xml:"sat_finish"`
		SatBreak                       string `xml:"sat_break"`
		SatLeaveType                   string `xml:"sat_leave_type"`
		SatLeave                       string `xml:"sat_leave"`
		SatLeaveHours                  string `xml:"sat_leave_hours"`
		SatNotes                       string `xml:"sat_notes"`
		SatPMPCID                      string `xml:"sat_PM_PCID"`
		SunWork                        string `xml:"sun_work"`
		SunStart                       string `xml:"sun_start"`
		SunFinish                      string `xml:"sun_finish"`
		SunBreak                       string `xml:"sun_break"`
		SunLeaveType                   string `xml:"sun_leave_type"`
		SunLeave                       string `xml:"sun_leave"`
		SunLeaveHours                  string `xml:"sun_leave_hours"`
		SunNotes                       string `xml:"sun_notes"`
		SunPMPCID                      string `xml:"sun_PM_PCID"`
		Notes                          string `xml:"notes"`
		Complete                       string `xml:"complete"`
		CompleteDst                    string `xml:"complete_dst"`
		Injury                         string `xml:"injury"`
		Approved                       string `xml:"approved"`
		ApprovedDst                    string `xml:"approved_dst"`
		ApprovedBy                     string `xml:"approved_by"`
		Paid                           string `xml:"paid"`
		PaidDst                        string `xml:"paid_dst"`
		ExportedAt                     string `xml:"exported_at"`
		ExportedInvoicenumber          string `xml:"exported_invoicenumber"`
		DidNotWork                     string `xml:"did_not_work"`
		DidNotWorkDst                  string `xml:"did_not_work_dst"`
		Lastupdated                    string `xml:"lastupdated"`
		Status                         string `xml:"status"`
		TimesheetOffsetStartingWeekday string `xml:"timesheetOffset_startingWeekday"`
		InputSource                    string `xml:"input_source"`
		CMCID                          string `xml:"CM_CID"`
		PayrollCMCID                   string `xml:"payroll_CM_CID"`
		ExcludePayroll                 string `xml:"exclude_payroll"`
		PayrunID                       string `xml:"payrun_id"`
		MasterTimesheetID              string `xml:"master_timesheet_id"`
		Remoteid                       string `xml:"remoteid"`
		Approver                       string `xml:"approver"`
		ApproverEmail                  string `xml:"approver_email"`
	} `xml:"user"`
}
