// Code generated from openapi/company-v3.json; DO NOT EDIT.

// The closed sets of values the Company API accepts, named by the document.
//
// Constants rather than a named type of their own, so one goes wherever the string goes:
// `clockster.UsersRoleEmployee` is `"employee"` and drops into the field that takes it. The
// fields stay strings, so a value the API starts answering with tomorrow reaches you as itself.
//
// Only what you send is ever closed. Every set here is on a query parameter or in a request body,
// and none of them is in an answer.

package clockster

// AttendanceInclude is what this field is allowed to be. One of "user", "location", "attachments".
const (
	AttendanceIncludeUser        = "user"
	AttendanceIncludeLocation    = "location"
	AttendanceIncludeAttachments = "attachments"
)

// AttendanceIncludeValues is every value of AttendanceInclude, in the order the document names
// them.
func AttendanceIncludeValues() []string {
	return []string{
		AttendanceIncludeUser,
		AttendanceIncludeLocation,
		AttendanceIncludeAttachments,
	}
}

// AttendanceSource is what this field is allowed to be. One of "device", "mobile", "frontend",
// "api", "system".
const (
	AttendanceSourceDevice   = "device"
	AttendanceSourceMobile   = "mobile"
	AttendanceSourceFrontend = "frontend"
	AttendanceSourceAPI      = "api"
	AttendanceSourceSystem   = "system"
)

// AttendanceSourceValues is every value of AttendanceSource, in the order the document names them.
func AttendanceSourceValues() []string {
	return []string{
		AttendanceSourceDevice,
		AttendanceSourceMobile,
		AttendanceSourceFrontend,
		AttendanceSourceAPI,
		AttendanceSourceSystem,
	}
}

// AttendanceStatus is what this field is allowed to be. One of "out", "in", "break".
const (
	AttendanceStatusOut   = "out"
	AttendanceStatusIn    = "in"
	AttendanceStatusBreak = "break"
)

// AttendanceStatusValues is every value of AttendanceStatus, in the order the document names them.
func AttendanceStatusValues() []string {
	return []string{
		AttendanceStatusOut,
		AttendanceStatusIn,
		AttendanceStatusBreak,
	}
}

// DepartmentsInclude is what this field is allowed to be. One of "managers".
const (
	DepartmentsIncludeManagers = "managers"
)

// DepartmentsIncludeValues is every value of DepartmentsInclude, in the order the document names
// them.
func DepartmentsIncludeValues() []string {
	return []string{
		DepartmentsIncludeManagers,
	}
}

// DocumentsEmploymentType is what this field is allowed to be. One of "full_time", "part_time",
// "irregular_hours", "contract_1", "contract_2", "apprenticeship", "traineeship", "piece_rate",
// "probation", "outstaffing".
const (
	DocumentsEmploymentTypeFullTime       = "full_time"
	DocumentsEmploymentTypePartTime       = "part_time"
	DocumentsEmploymentTypeIrregularHours = "irregular_hours"
	DocumentsEmploymentTypeContract1      = "contract_1"
	DocumentsEmploymentTypeContract2      = "contract_2"
	DocumentsEmploymentTypeApprenticeship = "apprenticeship"
	DocumentsEmploymentTypeTraineeship    = "traineeship"
	DocumentsEmploymentTypePieceRate      = "piece_rate"
	DocumentsEmploymentTypeProbation      = "probation"
	DocumentsEmploymentTypeOutstaffing    = "outstaffing"
)

// DocumentsEmploymentTypeValues is every value of DocumentsEmploymentType, in the order the
// document names them.
func DocumentsEmploymentTypeValues() []string {
	return []string{
		DocumentsEmploymentTypeFullTime,
		DocumentsEmploymentTypePartTime,
		DocumentsEmploymentTypeIrregularHours,
		DocumentsEmploymentTypeContract1,
		DocumentsEmploymentTypeContract2,
		DocumentsEmploymentTypeApprenticeship,
		DocumentsEmploymentTypeTraineeship,
		DocumentsEmploymentTypePieceRate,
		DocumentsEmploymentTypeProbation,
		DocumentsEmploymentTypeOutstaffing,
	}
}

// DocumentsInclude is what this field is allowed to be. One of "attachments", "signers",
// "labor_contract".
const (
	DocumentsIncludeAttachments   = "attachments"
	DocumentsIncludeSigners       = "signers"
	DocumentsIncludeLaborContract = "labor_contract"
)

// DocumentsIncludeValues is every value of DocumentsInclude, in the order the document names them.
func DocumentsIncludeValues() []string {
	return []string{
		DocumentsIncludeAttachments,
		DocumentsIncludeSigners,
		DocumentsIncludeLaborContract,
	}
}

// DocumentsParty is what this field is allowed to be. One of "employee", "counterparty".
const (
	DocumentsPartyEmployee     = "employee"
	DocumentsPartyCounterparty = "counterparty"
)

// DocumentsPartyValues is every value of DocumentsParty, in the order the document names them.
func DocumentsPartyValues() []string {
	return []string{
		DocumentsPartyEmployee,
		DocumentsPartyCounterparty,
	}
}

// DocumentsType is what this field is allowed to be. One of "passport", "cv", "diploma",
// "medical", "photo", "other", "medical_book", "employment_agreement",
// "termination_of_employment_agreement", "equipment_agreement", "application", "order",
// "supplementary_agreement", "job_description", "nda", "non_compete_agreement",
// "data_processing_agreement", "act_of_service_acceptance", "health_and_safety_briefing",
// "shift_schedule", "letter", "vacation_schedule", "contract", "agreement", "goods_release_note",
// "reconciliation_act", "return_to_supplier".
const (
	DocumentsTypePassport                         = "passport"
	DocumentsTypeCv                               = "cv"
	DocumentsTypeDiploma                          = "diploma"
	DocumentsTypeMedical                          = "medical"
	DocumentsTypePhoto                            = "photo"
	DocumentsTypeOther                            = "other"
	DocumentsTypeMedicalBook                      = "medical_book"
	DocumentsTypeEmploymentAgreement              = "employment_agreement"
	DocumentsTypeTerminationOfEmploymentAgreement = "termination_of_employment_agreement"
	DocumentsTypeEquipmentAgreement               = "equipment_agreement"
	DocumentsTypeApplication                      = "application"
	DocumentsTypeOrder                            = "order"
	DocumentsTypeSupplementaryAgreement           = "supplementary_agreement"
	DocumentsTypeJobDescription                   = "job_description"
	DocumentsTypeNda                              = "nda"
	DocumentsTypeNonCompeteAgreement              = "non_compete_agreement"
	DocumentsTypeDataProcessingAgreement          = "data_processing_agreement"
	DocumentsTypeActOfServiceAcceptance           = "act_of_service_acceptance"
	DocumentsTypeHealthAndSafetyBriefing          = "health_and_safety_briefing"
	DocumentsTypeShiftSchedule                    = "shift_schedule"
	DocumentsTypeLetter                           = "letter"
	DocumentsTypeVacationSchedule                 = "vacation_schedule"
	DocumentsTypeContract                         = "contract"
	DocumentsTypeAgreement                        = "agreement"
	DocumentsTypeGoodsReleaseNote                 = "goods_release_note"
	DocumentsTypeReconciliationAct                = "reconciliation_act"
	DocumentsTypeReturnToSupplier                 = "return_to_supplier"
)

// DocumentsTypeValues is every value of DocumentsType, in the order the document names them.
func DocumentsTypeValues() []string {
	return []string{
		DocumentsTypePassport,
		DocumentsTypeCv,
		DocumentsTypeDiploma,
		DocumentsTypeMedical,
		DocumentsTypePhoto,
		DocumentsTypeOther,
		DocumentsTypeMedicalBook,
		DocumentsTypeEmploymentAgreement,
		DocumentsTypeTerminationOfEmploymentAgreement,
		DocumentsTypeEquipmentAgreement,
		DocumentsTypeApplication,
		DocumentsTypeOrder,
		DocumentsTypeSupplementaryAgreement,
		DocumentsTypeJobDescription,
		DocumentsTypeNda,
		DocumentsTypeNonCompeteAgreement,
		DocumentsTypeDataProcessingAgreement,
		DocumentsTypeActOfServiceAcceptance,
		DocumentsTypeHealthAndSafetyBriefing,
		DocumentsTypeShiftSchedule,
		DocumentsTypeLetter,
		DocumentsTypeVacationSchedule,
		DocumentsTypeContract,
		DocumentsTypeAgreement,
		DocumentsTypeGoodsReleaseNote,
		DocumentsTypeReconciliationAct,
		DocumentsTypeReturnToSupplier,
	}
}

// LocationsInclude is what this field is allowed to be. One of "managers".
const (
	LocationsIncludeManagers = "managers"
)

// LocationsIncludeValues is every value of LocationsInclude, in the order the document names them.
func LocationsIncludeValues() []string {
	return []string{
		LocationsIncludeManagers,
	}
}

// PayrollPayslipsStatus is what this field is allowed to be. One of "draft", "approved", "paid".
const (
	PayrollPayslipsStatusDraft    = "draft"
	PayrollPayslipsStatusApproved = "approved"
	PayrollPayslipsStatusPaid     = "paid"
)

// PayrollPayslipsStatusValues is every value of PayrollPayslipsStatus, in the order the document
// names them.
func PayrollPayslipsStatusValues() []string {
	return []string{
		PayrollPayslipsStatusDraft,
		PayrollPayslipsStatusApproved,
		PayrollPayslipsStatusPaid,
	}
}

// SchedulesLeaveType is what this field is allowed to be. One of "annual", "unpaid", "sick",
// "unpaid_sick", "maternity", "paternity", "special", "day_off", "compensatory", "personal",
// "emergency", "unexcused_absence".
const (
	SchedulesLeaveTypeAnnual           = "annual"
	SchedulesLeaveTypeUnpaid           = "unpaid"
	SchedulesLeaveTypeSick             = "sick"
	SchedulesLeaveTypeUnpaidSick       = "unpaid_sick"
	SchedulesLeaveTypeMaternity        = "maternity"
	SchedulesLeaveTypePaternity        = "paternity"
	SchedulesLeaveTypeSpecial          = "special"
	SchedulesLeaveTypeDayOff           = "day_off"
	SchedulesLeaveTypeCompensatory     = "compensatory"
	SchedulesLeaveTypePersonal         = "personal"
	SchedulesLeaveTypeEmergency        = "emergency"
	SchedulesLeaveTypeUnexcusedAbsence = "unexcused_absence"
)

// SchedulesLeaveTypeValues is every value of SchedulesLeaveType, in the order the document names
// them.
func SchedulesLeaveTypeValues() []string {
	return []string{
		SchedulesLeaveTypeAnnual,
		SchedulesLeaveTypeUnpaid,
		SchedulesLeaveTypeSick,
		SchedulesLeaveTypeUnpaidSick,
		SchedulesLeaveTypeMaternity,
		SchedulesLeaveTypePaternity,
		SchedulesLeaveTypeSpecial,
		SchedulesLeaveTypeDayOff,
		SchedulesLeaveTypeCompensatory,
		SchedulesLeaveTypePersonal,
		SchedulesLeaveTypeEmergency,
		SchedulesLeaveTypeUnexcusedAbsence,
	}
}

// SchedulesType is what this field is allowed to be. One of "work", "free", "leave".
const (
	SchedulesTypeWork  = "work"
	SchedulesTypeFree  = "free"
	SchedulesTypeLeave = "leave"
)

// SchedulesTypeValues is every value of SchedulesType, in the order the document names them.
func SchedulesTypeValues() []string {
	return []string{
		SchedulesTypeWork,
		SchedulesTypeFree,
		SchedulesTypeLeave,
	}
}

// TasksInclude is what this field is allowed to be. One of "items", "managers", "user", "author".
const (
	TasksIncludeItems    = "items"
	TasksIncludeManagers = "managers"
	TasksIncludeUser     = "user"
	TasksIncludeAuthor   = "author"
)

// TasksIncludeValues is every value of TasksInclude, in the order the document names them.
func TasksIncludeValues() []string {
	return []string{
		TasksIncludeItems,
		TasksIncludeManagers,
		TasksIncludeUser,
		TasksIncludeAuthor,
	}
}

// TasksStatus is what this field is allowed to be. One of "created", "started", "paused",
// "completed", "incompleted", "pastdue".
const (
	TasksStatusCreated     = "created"
	TasksStatusStarted     = "started"
	TasksStatusPaused      = "paused"
	TasksStatusCompleted   = "completed"
	TasksStatusIncompleted = "incompleted"
	TasksStatusPastdue     = "pastdue"
)

// TasksStatusValues is every value of TasksStatus, in the order the document names them.
func TasksStatusValues() []string {
	return []string{
		TasksStatusCreated,
		TasksStatusStarted,
		TasksStatusPaused,
		TasksStatusCompleted,
		TasksStatusIncompleted,
		TasksStatusPastdue,
	}
}

// TimesheetsInclude is what this field is allowed to be. One of "actual", "variance", "user",
// "location", "department", "position".
const (
	TimesheetsIncludeActual     = "actual"
	TimesheetsIncludeVariance   = "variance"
	TimesheetsIncludeUser       = "user"
	TimesheetsIncludeLocation   = "location"
	TimesheetsIncludeDepartment = "department"
	TimesheetsIncludePosition   = "position"
)

// TimesheetsIncludeValues is every value of TimesheetsInclude, in the order the document names
// them.
func TimesheetsIncludeValues() []string {
	return []string{
		TimesheetsIncludeActual,
		TimesheetsIncludeVariance,
		TimesheetsIncludeUser,
		TimesheetsIncludeLocation,
		TimesheetsIncludeDepartment,
		TimesheetsIncludePosition,
	}
}

// UserFiltersInclude is what this field is allowed to be. One of "managers".
const (
	UserFiltersIncludeManagers = "managers"
)

// UserFiltersIncludeValues is every value of UserFiltersInclude, in the order the document names
// them.
func UserFiltersIncludeValues() []string {
	return []string{
		UserFiltersIncludeManagers,
	}
}

// UserRequestsInclude is what this field is allowed to be. One of "content", "user", "author".
const (
	UserRequestsIncludeContent = "content"
	UserRequestsIncludeUser    = "user"
	UserRequestsIncludeAuthor  = "author"
)

// UserRequestsIncludeValues is every value of UserRequestsInclude, in the order the document names
// them.
func UserRequestsIncludeValues() []string {
	return []string{
		UserRequestsIncludeContent,
		UserRequestsIncludeUser,
		UserRequestsIncludeAuthor,
	}
}

// UserRequestsStatus is what this field is allowed to be. One of "pending", "accepted",
// "rejected", "cancelled", "approval", "execution", "signing".
const (
	UserRequestsStatusPending   = "pending"
	UserRequestsStatusAccepted  = "accepted"
	UserRequestsStatusRejected  = "rejected"
	UserRequestsStatusCancelled = "cancelled"
	UserRequestsStatusApproval  = "approval"
	UserRequestsStatusExecution = "execution"
	UserRequestsStatusSigning   = "signing"
)

// UserRequestsStatusValues is every value of UserRequestsStatus, in the order the document names
// them.
func UserRequestsStatusValues() []string {
	return []string{
		UserRequestsStatusPending,
		UserRequestsStatusAccepted,
		UserRequestsStatusRejected,
		UserRequestsStatusCancelled,
		UserRequestsStatusApproval,
		UserRequestsStatusExecution,
		UserRequestsStatusSigning,
	}
}

// UserRequestsType is what this field is allowed to be. One of "leave", "work", "general",
// "finance".
const (
	UserRequestsTypeLeave   = "leave"
	UserRequestsTypeWork    = "work"
	UserRequestsTypeGeneral = "general"
	UserRequestsTypeFinance = "finance"
)

// UserRequestsTypeValues is every value of UserRequestsType, in the order the document names them.
func UserRequestsTypeValues() []string {
	return []string{
		UserRequestsTypeLeave,
		UserRequestsTypeWork,
		UserRequestsTypeGeneral,
		UserRequestsTypeFinance,
	}
}

// UsersEmployment is what this field is allowed to be. One of "full_time", "part_time",
// "irregular_hours", "contract_1", "contract_2", "apprenticeship", "traineeship", "piece_rate",
// "probation", "outstaffing".
const (
	UsersEmploymentFullTime       = "full_time"
	UsersEmploymentPartTime       = "part_time"
	UsersEmploymentIrregularHours = "irregular_hours"
	UsersEmploymentContract1      = "contract_1"
	UsersEmploymentContract2      = "contract_2"
	UsersEmploymentApprenticeship = "apprenticeship"
	UsersEmploymentTraineeship    = "traineeship"
	UsersEmploymentPieceRate      = "piece_rate"
	UsersEmploymentProbation      = "probation"
	UsersEmploymentOutstaffing    = "outstaffing"
)

// UsersEmploymentValues is every value of UsersEmployment, in the order the document names them.
func UsersEmploymentValues() []string {
	return []string{
		UsersEmploymentFullTime,
		UsersEmploymentPartTime,
		UsersEmploymentIrregularHours,
		UsersEmploymentContract1,
		UsersEmploymentContract2,
		UsersEmploymentApprenticeship,
		UsersEmploymentTraineeship,
		UsersEmploymentPieceRate,
		UsersEmploymentProbation,
		UsersEmploymentOutstaffing,
	}
}

// UsersGender is what this field is allowed to be. One of "male", "female", "other".
const (
	UsersGenderMale   = "male"
	UsersGenderFemale = "female"
	UsersGenderOther  = "other"
)

// UsersGenderValues is every value of UsersGender, in the order the document names them.
func UsersGenderValues() []string {
	return []string{
		UsersGenderMale,
		UsersGenderFemale,
		UsersGenderOther,
	}
}

// UsersInclude is what this field is allowed to be. One of "location", "locations", "department",
// "position", "user_filters", "dismissal", "meta".
const (
	UsersIncludeLocation    = "location"
	UsersIncludeLocations   = "locations"
	UsersIncludeDepartment  = "department"
	UsersIncludePosition    = "position"
	UsersIncludeUserFilters = "user_filters"
	UsersIncludeDismissal   = "dismissal"
	UsersIncludeMeta        = "meta"
)

// UsersIncludeValues is every value of UsersInclude, in the order the document names them.
func UsersIncludeValues() []string {
	return []string{
		UsersIncludeLocation,
		UsersIncludeLocations,
		UsersIncludeDepartment,
		UsersIncludePosition,
		UsersIncludeUserFilters,
		UsersIncludeDismissal,
		UsersIncludeMeta,
	}
}

// UsersLocale is what this field is allowed to be. One of "en", "ru", "kk", "uk", "id", "uz",
// "az", "fr", "vi", "zh".
const (
	UsersLocaleEn = "en"
	UsersLocaleRu = "ru"
	UsersLocaleKk = "kk"
	UsersLocaleUk = "uk"
	UsersLocaleID = "id"
	UsersLocaleUz = "uz"
	UsersLocaleAz = "az"
	UsersLocaleFr = "fr"
	UsersLocaleVi = "vi"
	UsersLocaleZh = "zh"
)

// UsersLocaleValues is every value of UsersLocale, in the order the document names them.
func UsersLocaleValues() []string {
	return []string{
		UsersLocaleEn,
		UsersLocaleRu,
		UsersLocaleKk,
		UsersLocaleUk,
		UsersLocaleID,
		UsersLocaleUz,
		UsersLocaleAz,
		UsersLocaleFr,
		UsersLocaleVi,
		UsersLocaleZh,
	}
}

// UsersRole is what this field is allowed to be. One of "admin", "employee".
const (
	UsersRoleAdmin    = "admin"
	UsersRoleEmployee = "employee"
)

// UsersRoleValues is every value of UsersRole, in the order the document names them.
func UsersRoleValues() []string {
	return []string{
		UsersRoleAdmin,
		UsersRoleEmployee,
	}
}

// UsersStatus is what this field is allowed to be. One of "active", "dismissed", "all".
const (
	UsersStatusActive    = "active"
	UsersStatusDismissed = "dismissed"
	UsersStatusAll       = "all"
)

// UsersStatusValues is every value of UsersStatus, in the order the document names them.
func UsersStatusValues() []string {
	return []string{
		UsersStatusActive,
		UsersStatusDismissed,
		UsersStatusAll,
	}
}

// WebhooksDeliveriesInclude is what this field is allowed to be. One of "payload".
const (
	WebhooksDeliveriesIncludePayload = "payload"
)

// WebhooksDeliveriesIncludeValues is every value of WebhooksDeliveriesInclude, in the order the
// document names them.
func WebhooksDeliveriesIncludeValues() []string {
	return []string{
		WebhooksDeliveriesIncludePayload,
	}
}

// WebhooksEvent is what this field is allowed to be. One of "user.created", "user.updated",
// "user.deleted", "user.restored", "user.purged", "location.created", "location.updated",
// "location.deleted", "department.created", "department.updated", "department.deleted",
// "position.created", "position.updated", "position.deleted", "task.created", "task.completed",
// "task.approved", "task.rejected", "task.deleted".
const (
	WebhooksEventUserCreated       = "user.created"
	WebhooksEventUserUpdated       = "user.updated"
	WebhooksEventUserDeleted       = "user.deleted"
	WebhooksEventUserRestored      = "user.restored"
	WebhooksEventUserPurged        = "user.purged"
	WebhooksEventLocationCreated   = "location.created"
	WebhooksEventLocationUpdated   = "location.updated"
	WebhooksEventLocationDeleted   = "location.deleted"
	WebhooksEventDepartmentCreated = "department.created"
	WebhooksEventDepartmentUpdated = "department.updated"
	WebhooksEventDepartmentDeleted = "department.deleted"
	WebhooksEventPositionCreated   = "position.created"
	WebhooksEventPositionUpdated   = "position.updated"
	WebhooksEventPositionDeleted   = "position.deleted"
	WebhooksEventTaskCreated       = "task.created"
	WebhooksEventTaskCompleted     = "task.completed"
	WebhooksEventTaskApproved      = "task.approved"
	WebhooksEventTaskRejected      = "task.rejected"
	WebhooksEventTaskDeleted       = "task.deleted"
)

// WebhooksEventValues is every value of WebhooksEvent, in the order the document names them.
func WebhooksEventValues() []string {
	return []string{
		WebhooksEventUserCreated,
		WebhooksEventUserUpdated,
		WebhooksEventUserDeleted,
		WebhooksEventUserRestored,
		WebhooksEventUserPurged,
		WebhooksEventLocationCreated,
		WebhooksEventLocationUpdated,
		WebhooksEventLocationDeleted,
		WebhooksEventDepartmentCreated,
		WebhooksEventDepartmentUpdated,
		WebhooksEventDepartmentDeleted,
		WebhooksEventPositionCreated,
		WebhooksEventPositionUpdated,
		WebhooksEventPositionDeleted,
		WebhooksEventTaskCreated,
		WebhooksEventTaskCompleted,
		WebhooksEventTaskApproved,
		WebhooksEventTaskRejected,
		WebhooksEventTaskDeleted,
	}
}
