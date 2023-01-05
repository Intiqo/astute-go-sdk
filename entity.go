package astute

import (
	"encoding/xml"
	"time"
)

type PayrollEntity interface {
	queryUserXmlResponse | QueryUserResponse |
		queryTimesheetXmlResponse | QueryTimesheetResponse |
		saveTimesheetXmlResponse
}

type UserParams struct {
	UID    string
	UserId string
}

// Entities related to QueryUser

type QueryUserParams struct {
	JobCode string
}

type queryUserXmlResponse struct {
	XMLName       xml.Name `xml:"Envelope"`
	Text          string   `xml:",chardata"`
	EncodingStyle string   `xml:"encodingStyle,attr"`
	SOAPENV       string   `xml:"SOAP-ENV,attr"`
	Xsd           string   `xml:"xsd,attr"`
	Xsi           string   `xml:"xsi,attr"`
	SOAPENC       string   `xml:"SOAP-ENC,attr"`
	Tns           string   `xml:"tns,attr"`
	Body          struct {
		Text              string `xml:",chardata"`
		UserQueryResponse struct {
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
		} `xml:"UserQueryResponse"`
	} `xml:"Body"`
}

type QueryUserResponse struct {
	XMLName xml.Name `xml:"users"`
	Text    string   `xml:",chardata"`
	User    struct {
		Text                                          string `xml:",chardata"`
		UID                                           string `xml:"UID"`
		MID                                           string `xml:"MID"`
		ID                                            string `xml:"id"`
		Locked                                        string `xml:"locked"`
		NameFirst                                     string `xml:"name_first"`
		NameLast                                      string `xml:"name_last"`
		KnownAs                                       string `xml:"known_as"`
		DateBirth                                     string `xml:"date_birth"`
		EmploymentStartDate                           string `xml:"employment_start_date"`
		TerminationDate                               string `xml:"termination_date"`
		TerminationreasonID                           string `xml:"terminationreason_id"`
		TerminationNote                               string `xml:"termination_note"`
		Email                                         string `xml:"email"`
		EmailOptOut                                   string `xml:"email_opt_out"`
		PhoneMobile                                   string `xml:"phone_mobile"`
		PhoneCountryCode                              string `xml:"phone_country_code"`
		MobilePhone                                   string `xml:"mobile_phone"`
		MobileCountryCode                             string `xml:"mobile_country_code"`
		AddressStreet                                 string `xml:"address_street"`
		AddressLocality                               string `xml:"address_locality"`
		AddressRegion                                 string `xml:"address_region"`
		AddressPostcode                               string `xml:"address_postcode"`
		KinName                                       string `xml:"kin_name"`
		KinAddressStreet                              string `xml:"kin_address_street"`
		KinAddressLocality                            string `xml:"kin_address_locality"`
		KinAddressRegion                              string `xml:"kin_address_region"`
		KinAddressPostcode                            string `xml:"kin_address_postcode"`
		KinPhone                                      string `xml:"kin_phone"`
		KinEmail                                      string `xml:"kin_email"`
		KinRelationship                               string `xml:"kin_relationship"`
		Employer                                      string `xml:"employer"`
		WorkplaceName                                 string `xml:"workplace_name"`
		WorkplaceStreet                               string `xml:"workplace_street"`
		WorkplaceLocality                             string `xml:"workplace_locality"`
		WorkplaceRegion                               string `xml:"workplace_region"`
		WorkplacePostcode                             string `xml:"workplace_postcode"`
		InsuranceReference                            string `xml:"insurance_reference"`
		Logo                                          string `xml:"logo"`
		DateStart                                     string `xml:"date_start"`
		DateFinish                                    string `xml:"date_finish"`
		ManagerUserID                                 string `xml:"manager_user_id"`
		Manager2UserID                                string `xml:"manager2_user_id"`
		GroupManagerUserID                            string `xml:"group_manager_user_id"`
		RecruiterUserID                               string `xml:"recruiter_user_id"`
		Recruiter2UserID                              string `xml:"recruiter2_user_id"`
		Recruiter2MarginSplit                         string `xml:"recruiter2_margin_split"`
		Classification                                string `xml:"classification"`
		PayRate                                       string `xml:"pay_rate"`
		PayCurrency                                   string `xml:"pay_currency"`
		ChargeCurrency                                string `xml:"charge_currency"`
		PayType                                       string `xml:"pay_type"`
		PermissionGroup                               string `xml:"permission_group"`
		SubcontractorEntity                           string `xml:"subcontractor_entity"`
		PayOncosts                                    string `xml:"pay_oncosts"`
		PayChargeRate                                 string `xml:"pay_charge_rate"`
		PayRateInterval                               string `xml:"pay_rate_interval"`
		PayFrequency                                  string `xml:"pay_frequency"`
		PayNotes                                      string `xml:"pay_notes"`
		AccountName                                   string `xml:"account_name"`
		AccountBsb                                    string `xml:"account_bsb"`
		AccountNumber                                 string `xml:"account_number"`
		SecondaryAccountName                          string `xml:"secondary_account_name"`
		SecondaryAccountBsb                           string `xml:"secondary_account_bsb"`
		SecondaryAccountNumber                        string `xml:"secondary_account_number"`
		DepositAmount                                 string `xml:"deposit_amount"`
		DepositType                                   string `xml:"deposit_type"`
		TertiaryAccountName                           string `xml:"tertiary_account_name"`
		TertiaryAccountBsb                            string `xml:"tertiary_account_bsb"`
		TertiaryAccountNumber                         string `xml:"tertiary_account_number"`
		TertiaryDepositAmount                         string `xml:"tertiary_deposit_amount"`
		TertiaryDepositType                           string `xml:"tertiary_deposit_type"`
		SuperMemberReference                          string `xml:"super_member_reference"`
		Tfn                                           string `xml:"tfn"`
		TfnNotDeclaredReason                          string `xml:"tfn_not_declared_reason"`
		TfnPayeeTaxStatus                             string `xml:"tfn_payee_tax_status"`
		TfnSubmissionMethod                           string `xml:"tfn_submission_method"`
		TfnDueAt                                      string `xml:"tfn_due_at"`
		TfnSubmittedAtUtc                             string `xml:"tfn_submitted_at_utc"`
		TfnSubmittedBy                                string `xml:"tfn_submitted_by"`
		BillingEntity                                 string `xml:"billing_entity"`
		BillingAbn                                    string `xml:"billing_abn"`
		BillingContact                                string `xml:"billing_contact"`
		BillingEmail                                  string `xml:"billing_email"`
		TaxPayrollExempt                              string `xml:"tax_payroll_exempt"`
		AcceptTerms                                   string `xml:"accept_terms"`
		Ready                                         string `xml:"ready"`
		TaxTable                                      string `xml:"tax_table"`
		TaxVariationValue                             string `xml:"tax_variation_value"`
		TaxOffset                                     string `xml:"tax_offset"`
		EsctRate                                      string `xml:"esct_rate"`
		HasSlAgreement                                string `xml:"has_sl_agreement"`
		SlAgreementRate                               string `xml:"sl_agreement_rate"`
		AccountCode                                   string `xml:"account_code"`
		ExpenseAccountCode                            string `xml:"expense_account_code"`
		JobCode                                       string `xml:"job_code"`
		MyobJobCode                                   string `xml:"myob_job_code"`
		Remoteid                                      string `xml:"remoteid"`
		Lastupdated                                   string `xml:"lastupdated"`
		SubcontractorNogst                            string `xml:"subcontractor_nogst"`
		ParentUID                                     string `xml:"ParentUID"`
		DivisionManagerUID                            string `xml:"DivisionManagerUID"`
		JobTitle                                      string `xml:"job_title"`
		BID                                           string `xml:"BID"`
		BCID                                          string `xml:"BCID"`
		WID                                           string `xml:"WID"`
		RemoteidOut                                   string `xml:"remoteid_out"`
		SID                                           string `xml:"SID"`
		IntegrationPayprocessPendingGenerateSales     string `xml:"integration_payprocess_pending_generate_sales"`
		IntegrationPayprocessPendingGeneratePurchases string `xml:"integration_payprocess_pending_generate_purchases"`
		IntegrationExpprocessPendingGenerateSales     string `xml:"integration_expprocess_pending_generate_sales"`
		IntegrationExpprocessPendingGeneratePurchases string `xml:"integration_expprocess_pending_generate_purchases"`
		IntegrationPayprocessPendingOverride          string `xml:"integration_payprocess_pending_override"`
		IntegrationExpprocessPendingOverride          string `xml:"integration_expprocess_pending_override"`
		CategoryCode                                  string `xml:"category_code"`
		BillersPonum                                  string `xml:"billers_ponum"`
		ItemCode                                      string `xml:"item_code"`
		IsUserProfileReviewer                         string `xml:"isUserProfileReviewer"`
		CreationTimeGMT                               string `xml:"creationTimeGMT"`
		DismissedChromeWarning                        string `xml:"dismissedChromeWarning"`
		BullhornID                                    string `xml:"bullhornID"`
		BullhornIDPlacement                           string `xml:"bullhornID_Placement"`
		BullhornIDCustomer                            string `xml:"bullhornID_Customer"`
		BullhornIDApprover1                           string `xml:"bullhornID_Approver1"`
		BullhornIDApprover2                           string `xml:"bullhornID_Approver2"`
		BullhornIDRecruiter1                          string `xml:"bullhornID_Recruiter1"`
		TmpPCreport                                   string `xml:"tmp_PCreport"`
		VisaType                                      string `xml:"visa_type"`
		VisaExpiryDate                                string `xml:"visa_expiry_date"`
		VisaVerifiedConfirm                           string `xml:"visa_verified_confirm"`
		VisaVerifiedDate                              string `xml:"visa_verified_date"`
		VisaVerifiedBy                                string `xml:"visa_verified_by"`
		CMCID                                         string `xml:"CM_CID"`
		PayrollCMCID                                  string `xml:"payroll_CM_CID"`
		BrandingCMCID                                 string `xml:"branding_CM_CID"`
		MicropayEmployeeNumber                        string `xml:"micropay_employee_number"`
		SAPConsultantCode                             string `xml:"SAP_ConsultantCode"`
		LeaveManagementEnabled                        string `xml:"leaveManagement_enabled"`
		LeaveManagementApproverUID                    string `xml:"leaveManagement_approverUID"`
		TimesheetOffsetStartingWeekday                string `xml:"timesheetOffset_startingWeekday"`
		FastTrackIDEmployee                           string `xml:"FastTrack_ID_employee"`
		FastTrackIDJob                                string `xml:"FastTrack_ID_job"`
		FastTrackIDCustomer                           string `xml:"FastTrack_ID_customer"`
		ADPID                                         string `xml:"ADP_ID"`
		TimesheetInvoicing                            string `xml:"timesheet_invoicing"`
		TimesheetPayroll                              string `xml:"timesheet_payroll"`
		TimesheetEnabled                              string `xml:"timesheet_enabled"`
		TimesheetStartDate                            string `xml:"timesheet_start_date"`
		FixedpayitemEnabled                           string `xml:"fixedpayitem_enabled"`
		PayBasis                                      string `xml:"pay_basis"`
		AnnualSalary                                  string `xml:"annual_salary"`
		HoursPerWk                                    string `xml:"hours_per_wk"`
		TsType                                        string `xml:"ts_type"`
		TsTypePiecework                               string `xml:"ts_type_piecework"`
		SplitShifts                                   string `xml:"split_shifts"`
		PayCycle                                      string `xml:"pay_cycle"`
		BaseWagesRate                                 string `xml:"base_wages_rate"`
		BaseWage                                      string `xml:"base_wage"`
		NameMiddle                                    string `xml:"name_middle"`
		Gender                                        string `xml:"gender"`
		Engagement                                    string `xml:"engagement"`
		PayslipEmail                                  string `xml:"payslip_email"`
		InsuranceRate                                 string `xml:"insurance_rate"`
		InsuranceID                                   string `xml:"insurance_id"`
		InsuranceFallback                             string `xml:"insurance_fallback"`
		PayPayrollTax                                 string `xml:"pay_payroll_tax"`
		PamID                                         string `xml:"pam_id"`
		PreviousPayrollSystemID                       string `xml:"previous_payroll_system_id"`
		JobPamID                                      string `xml:"job_pam_id"`
		TimesheetExpense                              string `xml:"timesheet_expense"`
		OLOID                                         string `xml:"OL_OID"`
		ConsolidationCode                             string `xml:"consolidation_code"`
		CostCentre                                    string `xml:"cost_centre"`
		DebtorCode                                    string `xml:"debtor_code"`
		LiabilityAccountCode                          string `xml:"liability_account_code"`
		JobExpenseAccountCode                         string `xml:"job_expense_account_code"`
		JobIncomeAccountCode                          string `xml:"job_income_account_code"`
		SubentityAssigmentID                          string `xml:"subentityAssigment_id"`
		SubentityOneID                                string `xml:"subentityOne_id"`
		SubentityTwoID                                string `xml:"subentityTwo_id"`
		MasterTimesheetID                             string `xml:"master_timesheet_id"`
		FixedTimesheetStartDate                       string `xml:"fixed_timesheet_start_date"`
		MasterTimesheetAwardID                        string `xml:"master_timesheet_award_id"`
		MasterTimesheetPaycodeID                      string `xml:"master_timesheet_paycode_id"`
		LeaverequestEnabled                           string `xml:"leaverequest_enabled"`
		Recruiter3UserID                              string `xml:"recruiter3_user_id"`
		Recruiter3MarginSplit                         string `xml:"recruiter3_margin_split"`
		Recruiter4UserID                              string `xml:"recruiter4_user_id"`
		Recruiter4MarginSplit                         string `xml:"recruiter4_margin_split"`
		Recruiter5UserID                              string `xml:"recruiter5_user_id"`
		Recruiter5MarginSplit                         string `xml:"recruiter5_margin_split"`
		MasterTimesheetRate                           string `xml:"master_timesheet_rate"`
		MasterTimesheetUseBaseRate                    string `xml:"master_timesheet_use_base_rate"`
		ApproverGroupID                               string `xml:"ApproverGroupID"`
		TimesheetFinalisedData                        string `xml:"timesheet_finalised_data"`
		JobUsedForTimesheets                          string `xml:"job_used_for_timesheets"`
		FixedTimesheetEnabled                         string `xml:"fixed_timesheet_enabled"`
		TimesheetLength                               string `xml:"timesheet_length"`
		SuperAbn                                      string `xml:"super_abn"`
		SuperUsi                                      string `xml:"super_usi"`
		SuperName                                     string `xml:"super_name"`
		SuperProductName                              string `xml:"super_product_name"`
		SuperPhone                                    string `xml:"super_phone"`
		SuperAddress1                                 string `xml:"super_address1"`
		SuperAddress2                                 string `xml:"super_address2"`
		SuperAddress3                                 string `xml:"super_address3"`
		SuperState                                    string `xml:"super_state"`
		SuperPostcode                                 string `xml:"super_postcode"`
		SuperSmsf                                     string `xml:"super_smsf"`
		SuperEsa                                      string `xml:"super_esa"`
		SuperDocumentStorageID                        string `xml:"super_document_storage_id"`
		SuperDocumentType                             string `xml:"super_document_type"`
		SuperSpin                                     string `xml:"super_spin"`
		SuperPaymentBsb                               string `xml:"super_payment_bsb"`
		SuperPaymentAccount                           string `xml:"super_payment_account"`
		SuperPaymentReference                         string `xml:"super_payment_reference"`
		SuperApproverUID                              string `xml:"super_approver_UID"`
		SuperApproverDatetime                         string `xml:"super_approver_datetime"`
		SuperAccountName                              string `xml:"super_account_name"`
		IsKiwisaverEligible                           string `xml:"is_kiwisaver_eligible"`
		IsKiwisaverOptedout                           string `xml:"is_kiwisaver_optedout"`
		KiwisaverDateOptoutReceived                   string `xml:"kiwisaver_date_optout_received"`
		IsOnContributionHoliday                       string `xml:"is_on_contribution_holiday"`
		ContributionHolidayStartDate                  string `xml:"contribution_holiday_start_date"`
		ContributionHolidayEndDate                    string `xml:"contribution_holiday_end_date"`
		DirectDebitPayment                            string `xml:"direct_debit_payment"`
		UserCustomField1                              string `xml:"user_custom_field_1"`
		UserCustomField2                              string `xml:"user_custom_field_2"`
		UserCustomField3                              string `xml:"user_custom_field_3"`
		UserCustomField4                              string `xml:"user_custom_field_4"`
		UserCustomField5                              string `xml:"user_custom_field_5"`
		JobCustomField1                               string `xml:"job_custom_field_1"`
		JobCustomField2                               string `xml:"job_custom_field_2"`
		JobCustomField3                               string `xml:"job_custom_field_3"`
		JobCustomField4                               string `xml:"job_custom_field_4"`
		JobCustomField5                               string `xml:"job_custom_field_5"`
		PayconditionAID                               string `xml:"paycondition_AID"`
		PayconditionPCGID                             string `xml:"paycondition_PCGID"`
		HolidayGroupID                                string `xml:"holiday_group_id"`
		TaxTID                                        string `xml:"tax_TID"`
		SuperModificationDate                         string `xml:"super_modification_date"`
		InvoiceitemTypeIDTs                           string `xml:"invoiceitem_type_id_ts"`
		InvoiceitemTypeIDExp                          string `xml:"invoiceitem_type_id_exp"`
		AccountCodeExp                                string `xml:"account_code_exp"`
		JobInvoiceitemTypeIDTs                        string `xml:"job_invoiceitem_type_id_ts"`
		JobInvoiceitemTypeIDExp                       string `xml:"job_invoiceitem_type_id_exp"`
		JobAccountCodeExp                             string `xml:"job_account_code_exp"`
		GlGroupID                                     string `xml:"gl_group_id"`
		GlGroupFallback                               string `xml:"gl_group_fallback"`
		JobGlGroupID                                  string `xml:"job_gl_group_id"`
		JobGlGroupFallback                            string `xml:"job_gl_group_fallback"`
		JobGlEdit                                     string `xml:"job_gl_edit"`
		NotifyViaSms                                  string `xml:"notify_via_sms"`
		JobNotes                                      string `xml:"job_notes"`
		Archived                                      string `xml:"archived"`
		EmployeeFullTimeWeekMTSID                     string `xml:"employee_full_time_week_MTS_id"`
		JobFullTimeWeekMTSID                          string `xml:"job_full_time_week_MTS_id"`
		EmployeeExternalPayrollOutboundID             string `xml:"employee_external_payroll_outbound_id"`
		EmployeeExternalPayrollInboundID              string `xml:"employee_external_payroll_inbound_id"`
		EmployeeExternalInvoiceOutboundID             string `xml:"employee_external_invoice_outbound_id"`
		EmployeeExternalInvoiceInboundID              string `xml:"employee_external_invoice_inbound_id"`
		EmployeeExternalTimesheetOutboundID           string `xml:"employee_external_timesheet_outbound_id"`
		EmployeeExternalTimesheetInboundID            string `xml:"employee_external_timesheet_inbound_id"`
		JobExternalPayrollOutboundID                  string `xml:"job_external_payroll_outbound_id"`
		JobExternalPayrollInboundID                   string `xml:"job_external_payroll_inbound_id"`
		JobExternalInvoiceOutboundID                  string `xml:"job_external_invoice_outbound_id"`
		JobExternalInvoiceInboundID                   string `xml:"job_external_invoice_inbound_id"`
		JobExternalTimesheetOutboundID                string `xml:"job_external_timesheet_outbound_id"`
		JobExternalTimesheetInboundID                 string `xml:"job_external_timesheet_inbound_id"`
		AuthSecondFactor                              string `xml:"auth_second_factor"`
		AuthBackupEmail                               string `xml:"auth_backup_email"`
		PassportCountry                               string `xml:"passport_country"`
		Residency                                     string `xml:"residency"`
	} `xml:"user"`
}

// Entities related to QueryTimesheet

type QueryTimesheetParams struct {
	UID           string
	TimesheetDate string
}

type queryTimesheetXmlResponse struct {
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
				} `xml:"Results"`
			} `xml:"parms_out"`
		} `xml:"TimesheetQueryResponse"`
	} `xml:"Body"`
}

type QueryTimesheetResponse struct {
	XMLName xml.Name `xml:"users"`
	Text    string   `xml:",chardata"`
	User    []struct {
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

// Entities related to SaveTimesheet

type SaveTimesheetParams struct {
	UserParams
	TSID       string
	StartTime  time.Time
	EndTime    time.Time
	BreakTime  string
	Notes      string
	DidNotWork bool
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

type SaveTimesheetResponse struct {
	TimesheetId string
}

// Entities related to SubmitTimesheet

type SubmitTimesheetParams struct {
	UserParams
	TSID           string
	StartTime      time.Time
	SubmissionTime time.Time
}
