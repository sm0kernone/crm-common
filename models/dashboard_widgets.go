package models

import (
	"github.com/sm0kernone/crm-common/pkg/utils"
	"errors"
	"fmt"
)

// FEWidgetType is the name of widget on frontend side
type FEWidgetType string

var (
	GeneralWidgetType     FEWidgetType = "general"
	ShiftWidgetType       FEWidgetType = "shift"
	SourceSplitWidgetType FEWidgetType = "source_split"
	PaymentWidgetType     FEWidgetType = "payment"
	IncomeWidgetType      FEWidgetType = "income"
	ResultWidgetType      FEWidgetType = "result"
)

var AllowedWidgetTypes = []FEWidgetType{GeneralWidgetType, ShiftWidgetType, SourceSplitWidgetType, PaymentWidgetType,
	IncomeWidgetType, ResultWidgetType}

type WidgetShift struct {
	ShiftName  string     `json:"shift_name"`
	ShiftType  ShiftType  `json:"shift_type"`
	ShiftHours ShiftHours `json:"shift_hours"`
}

func (w *WidgetShift) Validate() error {
	if !w.ShiftType.IsValid() {
		return fmt.Errorf("%v shiftType is not valid; allowed shiftTypes: %v", w.ShiftType, AllShiftTypes)
	}

	if !w.ShiftHours.IsValid() {
		return fmt.Errorf("%v shiftHour is not valid; allowed shiftHours: %v", w.ShiftType, AllShiftHours)
	}

	return nil
}

// AdditionalWidgetParams public struct of widget
type AdditionalWidgetParams struct {
	Shifts               []WidgetShift            `json:"shifts,omitempty"`
	ShiftsA              []WidgetShift            `json:"shifts_a,omitempty"`
	ShiftsB              []WidgetShift            `json:"shifts_b,omitempty"`
	SpecificUsers        []int                    `json:"specific_users,omitempty"`
	CampaignSources      SourceCampaignSourcesReq `json:"campaign_sources,omitempty"`
	ProductTypes         SourceProductTypesReq    `json:"product_types,omitempty"`
	Sections             []IncomeWidgetSection    `json:"sections,omitempty"`
	RefundTypes          []RefundTypeWidgetParam  `json:"refunds,omitempty"`
	CalculationType      CalculationType          `json:"calculation_type,omitempty"`
	CalculationTypCustom CalculationTypeCustomReq `json:"calculation_type_custom,omitempty"`
	CustomAName          string                   `json:"custom_a_name"`
	CustomBName          string                   `json:"custom_b_name"`
}

type CalculationTypeCustomReq struct {
	From int `json:"from"`
	To   int `json:"to"`
}
type CalculationType string

var (
	CalculationTypeMonth  CalculationType = "month"
	CalculationTypeCustom CalculationType = "custom"
)

var AllWidgetCalculationTypes = []string{string(CalculationTypeMonth), string(CalculationTypeCustom)}

type RefundTypeWidgetParam string

var (
	RefundTypeWidgetFull       RefundTypeWidgetParam = "full"
	RefundTypeWidgetPartial    RefundTypeWidgetParam = "partial"
	RefundTypeWidgetChargeback RefundTypeWidgetParam = "chargeback"
)

var AllWidgetRefundTypes = []string{string(RefundTypeWidgetFull), string(RefundTypeWidgetPartial), string(RefundTypeWidgetChargeback)}

type SourceCampaignSourcesReq map[SourceCampaignSources]CampaignSourceConditions

type CampaignSourceConditions struct {
	In    []string `json:"in"`
	NotIn []string `json:"not_in"`
}

type SourceCampaignSources string

var (
	CVCampaignSource  SourceCampaignSources = "CV"
	MDCCampaignSource SourceCampaignSources = "MDC"
	BCCampaignSource  SourceCampaignSources = "BC"
)

type SourceProductType string

var (
	OnlineSourceProduct  SourceProductType = "online"
	OfflineSourceProduct SourceProductType = "offline"
	OtherSourceProduct   SourceProductType = "other"
)

type SourceProductTypesReq map[SourceProductType][]string

type IncomeWidgetSection string

var (
	IncomeWidgetSectionGross      IncomeWidgetSection = "gross"
	IncomeWidgetSectionNet        IncomeWidgetSection = "net"
	IncomeWidgetSectionRefund     IncomeWidgetSection = "refund"
	IncomeWidgetSectionChargeback IncomeWidgetSection = "chargeback"
	IncomeWidgetSectionBank       IncomeWidgetSection = "bank"
	IncomeWidgetSectionWin        IncomeWidgetSection = "win"
)

var AllIncomeWidgetSections = []string{string(IncomeWidgetSectionGross), string(IncomeWidgetSectionNet), string(IncomeWidgetSectionRefund), string(IncomeWidgetSectionChargeback),
	string(IncomeWidgetSectionBank), string(IncomeWidgetSectionWin)}

func (d *DashboardWidgets) ValidateIncomeWidgetSections() error {
	for _, section := range d.AdditionalParams.Sections {
		if !utils.InListString(string(section), AllIncomeWidgetSections) {
			return fmt.Errorf("%v section is not valid, allowed sections %v",
				section, AllIncomeWidgetSections)
		}
	}

	return nil
}

type WidgetPosition struct {
	Cols    int `sql:"cols" json:"cols" label:"Columns"`
	Rows    int `sql:"rows" json:"rows" label:"Rows"`
	MinRows int `json:"min_rows"`
	MinCols int `json:"min_cols"`
	PosX    int `sql:"pos_x" json:"pos_x" label:"Position X"`
	PosY    int `sql:"pos_y" json:"pos_y" label:"Position Y"`
}

type WidgetParamsType string

var (
	ParamsTypeDefault           WidgetParamsType = "default"
	ParamsTypeDefaultGross      WidgetParamsType = "default_gross"
	ParamsTypeDefaultNet        WidgetParamsType = "default_net"
	ParamsTypeShiftsGross       WidgetParamsType = "shifts_gross"
	ParamsTypeShiftsNet         WidgetParamsType = "shifts_net"
	ParamsTypeDefaultRefunds    WidgetParamsType = "refunds"
	ParamsTypeDefaultPrediction WidgetParamsType = "predictions"
	ParamsTypeDefaultWin        WidgetParamsType = "win"
	ParamsTypeSplitFirst        WidgetParamsType = "split_first"
	ParamsTypeSplitSecond       WidgetParamsType = "split_second"
	ParamsTypeTransaction       WidgetParamsType = "transaction"
	ParamsTypeApproveRate       WidgetParamsType = "approve_rate"
	ParamsTypeSwipeRate         WidgetParamsType = "swipe_rate"
	ParamsTypeRefundRate        WidgetParamsType = "refund_rate"
)

var AllowedWidgetParamsTypes = []WidgetParamsType{ParamsTypeDefault, ParamsTypeDefaultNet, ParamsTypeDefaultGross, ParamsTypeDefaultWin,
	ParamsTypeShiftsGross, ParamsTypeShiftsNet, ParamsTypeDefaultRefunds, ParamsTypeDefaultPrediction, ParamsTypeSplitFirst, ParamsTypeSplitSecond, ParamsTypeTransaction,
	ParamsTypeApproveRate, ParamsTypeSwipeRate, ParamsTypeRefundRate}

var WidgetTypeToParamsType = map[FEWidgetType][]WidgetParamsType{
	GeneralWidgetType:     {ParamsTypeDefaultGross, ParamsTypeDefaultNet, ParamsTypeShiftsGross, ParamsTypeShiftsNet, ParamsTypeDefaultRefunds, ParamsTypeDefaultPrediction, ParamsTypeDefaultWin},
	ShiftWidgetType:       {ParamsTypeDefault},
	SourceSplitWidgetType: {ParamsTypeSplitFirst, ParamsTypeSplitSecond},
	PaymentWidgetType:     {ParamsTypeTransaction, ParamsTypeApproveRate, ParamsTypeSwipeRate, ParamsTypeRefundRate},
	IncomeWidgetType:      {ParamsTypeDefault},
	ResultWidgetType:      {ParamsTypeDefault},
}

func ValidateWidgetParamsType(widgetType FEWidgetType, paramsType WidgetParamsType) bool {
	for _, allowedParamsType := range WidgetTypeToParamsType[widgetType] {
		if allowedParamsType == paramsType {
			return true
		}
	}

	return false
}

type DashboardWidgets struct {
	Base
	Name             string                 `json:"name" sql:"name" label:"Name" type:"string" validate:"required"`
	DashboardID      int                    `json:"dashboard_id" sql:"widget_id"`
	WidgetType       FEWidgetType           `json:"widget_type" sql:"widget_type" label:"Widget Type" type:"string" validate:"required"`
	ParamsType       WidgetParamsType       `json:"params_type" sql:"params_type" label:"Parameters Type"`
	AdditionalParams AdditionalWidgetParams `json:"additional_params" pg:"additional" sql:"additional" validate:"omitempty"`
	Position         WidgetPosition         `json:"position" sql:"position" validate:"required"`
	BackgroundColor  string                 `json:"background_color" sql:"background_color" validate:"required"`
}

func (d *DashboardWidgets) ValidateCalculationType() error {
	if !utils.InListString(string(d.AdditionalParams.CalculationType), AllWidgetCalculationTypes) {
		return fmt.Errorf("%v is not allowed as calculationType, allowed calculationTypes %v",
			d.AdditionalParams.CalculationType, AllWidgetCalculationTypes)
	}

	if d.AdditionalParams.CalculationType == CalculationTypeCustom {
		if d.AdditionalParams.CalculationTypCustom.From < 1 || d.AdditionalParams.CalculationTypCustom.From > 31 ||
			d.AdditionalParams.CalculationTypCustom.To < 1 || d.AdditionalParams.CalculationTypCustom.To > 31 {
			return errors.New("calculation custom date ranges are not valid")
		}
	}

	return nil
}
