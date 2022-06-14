package models

import (
	"github.com/sm0kernone/crm-common/pkg/config"
	"github.com/sm0kernone/crm-common/pkg/utils"
	"crypto/md5"
	"database/sql/driver"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/pariz/gountries"
	"gopkg.in/ugjka/go-tz.v2/tz"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var (
	Assignable    = "Assignable"
	NotAssignable = "Not Assignable"
)

type VoipStatus string

type Gender string

// all possible voip statuses
var (
	AnsweredCallVoipStatus VoipStatus = "Answered call"
	AttemptVoipStatus      VoipStatus = "Attempt"
	EfectiveCallVoipStatus VoipStatus = "Effective call"
)

var (
	ContactMeStatusReason                 = "Contact Me"
	ContactMeReassignStatusReason         = "Contact Me Reassign"
	ContactImmediatelyStatusReason        = "Contact Immediately"
	ContactMeImmediatelyStatusReason      = "Contact Me Immediately"
	DisqualifiedPoolStatusReason          = "Disqualified Pool"
	ReassignStatusReason                  = "Reassign"
	QualifiedStatusReason                 = "Qualified"
	DisqualifiedStatusReason              = "Disqualified"
	UnprocessedStatusReason               = "Unprocessed"
	ReassignDisqualifiedStatusReason      = "Reassign Disqualified"
	ReassignNoAnswerStatusReason          = "Reassign No Answer"
	ReassignContactMeStatusReason         = "Reassign Contact Me"
	ReassignContactMeNoAnswerStatusReason = "Reassign Contact Me No Answer"

	CallBackStatusReason         = "Call Back"
	CallBackInAWhileStatusReason = "Call In A While"
)

var (
	FemaleGender Gender = "Female"
	MaleGender   Gender = "Male"
)

type Leads struct {
	Id                            int        `sql:"id" json:"id,omitempty"`
	CrmId                         int        `sql:"crm_id" json:"crm_id,omitempty"`
	Email                         string     `sql:"email" json:"email,omitempty" validate:"required"`
	FirstName                     string     `sql:"first_name" json:"first_name,omitempty" validate:"required"`
	LastName                      string     `sql:"last_name" json:"last_name,omitempty"`
	CountryOfBirth                string     `sql:"country_of_birth" json:"country_of_birth,omitempty" validate:"required"`
	CountryOfResidence            string     `sql:"country_of_residence" json:"country_of_residence,omitempty" validate:"required"`
	StatusReason                  string     `sql:"status_reason" json:"status_reason,omitempty"`
	Birthday                      string     `sql:"birthday" json:"birthday,omitempty"`
	SpeakFrench                   string     `sql:"speak_french" json:"speak_french,omitempty"`
	AboveYears                    string     `sql:"above_years" json:"above_years,omitempty"`
	SpeakEnglish                  string     `sql:"speak_english" json:"speak_english,omitempty"`
	Education                     string     `sql:"education" json:"education,omitempty"`
	AvailableFunds                string     `sql:"available_funds" json:"available_funds,omitempty"`
	Campaign                      string     `sql:"campaign" json:"campaign,omitempty"`
	Keyword                       string     `sql:"keyword" json:"keyword,omitempty"`
	AdGroup                       string     `sql:"adgroup" json:"adgroup,omitempty"`
	AdId                          string     `sql:"adid" json:"adid,omitempty"`
	Source                        string     `sql:"source" json:"source,omitempty"`
	AdPosition                    string     `sql:"adposition" json:"adposition,omitempty"`
	Placement                     string     `sql:"placement" json:"placement,omitempty"`
	GclId                         string     `sql:"gclid" json:"gclid,omitempty"`
	Device                        string     `sql:"device" json:"device,omitempty"`
	GeoId                         string     `sql:"geoid" json:"geo_id,omitempty"`
	DeviceModel                   string     `sql:"devicemodel" json:"devicemodel,omitempty"`
	MatchType                     string     `sql:"matchtype" json:"matchtype,omitempty"`
	Phone                         string     `sql:"phone" json:"phone,omitempty" validate:"required"`
	Age                           int        `sql:"age" json:"age,omitempty"`
	Industry                      string     `sql:"industry" json:"industry,omitempty"`
	LeadStatus                    string     `sql:"lead_status" json:"lead_status,omitempty"`
	Rating                        string     `sql:"rating" json:"rating,omitempty"`
	Medium                        string     `sql:"medium" json:"medium,omitempty"`
	EmailCampaign                 string     `sql:"email_campaign" json:"email_campaign"`
	EmailTemplate                 string     `sql:"email_template" json:"email_template"`
	EmailSegment                  string     `sql:"email_segment" json:"email_segment"`
	LeadNumber                    string     `sql:"lead_number" json:"lead_number,omitempty"`
	AssignTime                    int        `sql:"assign_time" json:"assign_time,omitempty"`
	ContactMeTime                 int        `sql:"contact_me_time" json:"contact_me_time,omitempty"`
	Quote                         string     `sql:"quote" json:"quote,omitempty"`
	MainOccupation                string     `sql:"main_occupation"json:"main_occupation,omitempty"`
	ModifiedTime                  int        `sql:"modified_time" json:"modified_time,omitempty"`
	AssignedTo                    int        `sql:"assigned_to" json:"assigned_to,omitempty" validate:"required"`
	Touched                       string     `sql:"touched" json:"touched,omitempty"`
	CreatedTime                   int        `sql:"created_time" json:"created_time,omitempty"`
	ActivityTimeOne               int        `sql:"activity_time_1" json:"activity_time_1,omitempty"`
	ActivityTimeTwo               int        `sql:"activity_time_2" json:"activity_time_2,omitempty"`
	ActivityTimeThree             int        `sql:"activity_time_3" json:"activity_time_3,omitempty"`
	CampaignSource                string     `sql:"campaign_source" json:"campaign_source,omitempty"`
	MainLanguage                  string     `sql:"main_language" json:"main_language,omitempty"`
	MartialStatus                 string     `sql:"marital_status" json:"marital_status"`
	Gender                        Gender     `sql:"gender" json:"gender,omitempty"`
	DialerCampaign                string     `sql:"dialer_campaign" json:"dialer_campaign,omitempty"`
	OnlineStatus                  string     `sql:"online_status" json:"online_status,omitempty"`
	Dependent                     string     `sql:"dependent" json:"dependent,omitempty"`
	ContactMe                     bool       `sql:"contact_me,notnull" pg:",use_zero"json:"contact_me,omitempty"`
	SingleOrMarried               string     `sql:"single_or_married" json:"single_or_married,omitempty"`
	RangeOfMonthlySalary          string     `sql:"range_of_monthly_salary" json:"range_of_monthly_salary,omitempty"`
	Disqualified                  string     `sql:"disqualified" json:"disqualified,omitempty"`
	HaveAJob                      string     `sql:"have_a_job" json:"have_a_job,omitempty"`
	Attempt                       int        `sql:"attempt" json:"attempt,omitempty"`
	VisaType                      string     `sql:"visa_type"json:"visa_type,omitempty"`
	Nationality                   string     `sql:"nationality" json:"nationality,omitempty"`
	PreferredDestination          string     `sql:"preferred_destination" json:"preferred_destination,omitempty"`
	CurrentlyEmployed             string     `sql:"currently_employed" json:"currently_employed,omitempty"`
	Occupation                    string     `sql:"occupation" json:"occupation,omitempty"`
	YearsOfEmployment             string     `sql:"years_of_employment" json:"years_of_employment,omitempty"`
	CompletedHighSchool           string     `sql:"completed_high_school" json:"completed_high_school,omitempty"`
	OtherEducation                string     `sql:"other_education" json:"other_education,omitempty"`
	ChoseCurrency                 string     `sql:"choose_currency" json:"choose_currency,omitempty"`
	ExactBirthday                 string     `sql:"exact_birthday" json:"exact_birthday,omitempty"`
	CityYouWouldVisit             string     `sql:"city_you_would_visit" json:"city_you_would_visit,omitempty"`
	SalaryInCanada                string     `sql:"salary_in_canada" json:"salary_in_canada,omitempty"`
	PassengerOne                  string     `sql:"passenger_one" json:"passenger_one,omitempty"`
	PassengerTwo                  string     `sql:"passenger_two" json:"passenger_two,omitempty"`
	PassengerThree                string     `sql:"passenger_three" json:"passenger_three,omitempty"`
	ApplicationType               string     `sql:"application_type" json:"application_type,omitempty"`
	NetWorth                      string     `sql:"net_worth"json:"net_worth,omitempty"`
	NumberOfChildren              int        `sql:"number_of_children" json:"number_of_children,omitempty"`
	UpTo                          string     `sql:"up_to" json:"up_to,omitempty"`
	DateOfArrival                 string     `sql:"date_of_arrival" json:"date_of_arrival,omitempty"`
	NumberOfPassengers            int        `sql:"number_of_passengers" json:"number_of_passengers,omitempty"`
	UtmTerm                       string     `sql:"utm_term" json:"utm_term,omitempty"`
	PreviousStatus                string     `sql:"previous_status" json:"previous_status"`
	PreviousOwner                 string     `sql:"previous_owner" json:"previous_owner,omitempty"`
	ConversionRate                float64    `sql:"conversion_rate" json:"conversion_rate,omitempty"`
	FunnelStatusReason            string     `sql:"funnel_status_reason" json:"funnel_status_reason,omitempty"`
	CreatedAt                     int        `sql:"created_at" json:"created_at,omitempty"`
	UpdatedAt                     int        `sql:"updated_at" json:"updated_at,omitempty"`
	Mobile                        string     `sql:"mobile" json:"mobile,omitempty"`
	PaymentEmail                  string     `sql:"payment_email" json:"payment_email,omitempty"`
	PaymentTemplate               string     `sql:"payment_template" json:"payment_template,,omitempty"`
	SecondaryEmail                string     `sql:"secondary_email" json:"secondary_email,omitempty"`
	PaymentLinkStatusReason       string     `sql:"payment_link_status_reason" json:"payment_link_status_reason,omitempty"`
	CallAttempts                  int        `sql:"call_attempts" json:"call_attempts,omitempty"`
	Product                       string     `sql:"product" json:"product,omitempty"`
	Employed                      string     `sql:"employed" json:"employed,omitempty"`
	Password                      string     `sql:"password" json:"password,omitempty"`
	PreviousDisqualifiedStatus    string     `sql:"previous_disqualified_status" json:"previous_disqualified_status,omitempty"`
	VOIPStatus                    VoipStatus `sql:"voip_status" json:"voip_status,omitempty"`
	IWantTo                       string     `sql:"i_want_to" json:"i_want_to,omitempty"`
	AcceptGDPR                    bool       `sql:"accept_gdpr,notnull" pg:",use_zero" json:"accept_gdpr,omitempty"`
	ContactPreferencesInformation string     `sql:"contact_preferences_information" json:"contact_preferences_information,omitempty"`
	InformationAbout              string     `sql:"information_about" json:"information_about,omitempty"`
	BadExperience                 string     `sql:"bad_experience" json:"bad_experience,omitempty"`
	ReasonRating                  string     `sql:"reason_rating" json:"reason_rating,omitempty"`
	ContactPreferences            string     `sql:"contact_preferences" json:"contact_preferences"`
	PrivacyPolicy                 string     `sql:"privacy_policy" json:"privacy_policy"`
	LeavingAs                     string     `sql:"leaving_as" json:"leaving_as"`
	NotReceiveInformation         string     `sql:"not_receive_information" json:"not_receive_information"`
	ReasonText                    string     `sql:"reason_text" json:"reason_text,omitempty"`
	POBox                         string     `sql:"po_box" json:"po_box,omitempty"`
	NumberOfEmployees             string     `sql:"number_of_employees" json:"number_of_employees"`
	City                          string     `sql:"city" json:"city,omitempty"`
	State                         string     `sql:"state" json:"state,omitempty"`
	Country                       string     `sql:"country" json:"country,omitempty"`
	PostalCode                    string     `sql:"postal_code" json:"postal_code"`
	Fax                           string     `sql:"fax" json:"fax,omitempty"`
	Street                        string     `sql:"street" json:"street,omitempty"`
	Designation                   string     `sql:"designation" json:"designation,omitempty"`
	Company                       string     `sql:"company" json:"company,omitempty"`
	EmailOptOut                   bool       `sql:"email_opt_out,notnull" pg:",use_zero" json:"email_opt_out"`
	NumOfAttempts                 string     `sql:"num_of_attempts" json:"num_of_attempts"`
	AttemptTime                   int        `sql:"attempt_time" json:"attempt_time"`
	Assignable                    string     `sql:"assignable" json:"assignable,omitempty"`
	LastEventTime                 int        `sql:"last_event_time" json:"last_event_time"`
	YearsOfResidenceCountry       string     `sql:"years_residence_country" json:"years_of_residence_country"`
	VisitedCanada                 string     `sql:"visited_canada" json:"visited_canada"`
	FriendsFamilyCanada           string     `sql:"friends_family_canada" json:"friends_family_canada"`
	EductionLevel                 string     `sql:"eduction_level" json:"eduction_level"`
	AvgMonthlyIncome              string     `sql:"avg_monthly_income" json:"avg_monthly_income"`
	SavingsAmount                 string     `sql:"savings_amount" json:"savings_amount"`
	Children                      string     `sql:"children" json:"children,omitempty"`
	HowManyChildren               string     `sql:"how_many_children" json:"how_many_children"`
	HealthProblems                string     `sql:"health_problems" json:"health_problems"`
	WhatHealthProblems            string     `sql:"what_health_problems" json:"what_health_problems"`
	CriminalRecords               string     `sql:"criminal_records" json:"criminal_records"`
	WhatCriminalRecords           string     `sql:"what_criminal_records" json:"what_criminal_records"`
	WhyImmigrate                  string     `sql:"why_immigrate" json:"why_immigrate"`
	ThinkingAboutMove             string     `sql:"thinking_about_move" json:"thinking_about_move"`
	StartedSavingProcess          string     `sql:"started_saving_process" json:"started_saving_process"`
	ImmigrationProcess            string     `sql:"immigration_process" json:"immigration_process"`
	SwipedTime                    int        `sql:"swiped_time" json:"swiped_time"`
	AnswerVoipTime                int        `sql:"answer_voip_time" json:"answer_voip_time"`
	ActivityDate1                 int        `sql:"activity_date_1" json:"activity_date_1"`
	ActivityDate2                 int        `sql:"activity_date_2" json:"activity_date_2"`
	ActivityDate3                 int        `sql:"activity_date_3" json:"activity_date_3"`
	FbclId                        string     `sql:"fbclid" json:"fbclid,omitempty"`
	PreviousCampaignSource        string     `sql:"previous_campaign_source" json:"previous_campaign_source"`
	EffectiveDate                 int        `sql:"effective_date" json:"effective_date"`
	ScheduleTime                  int        `sql:"schedule_time" json:"schedule_time"`
	DisqualifiedTime              int        `sql:"disqualified_time" json:"disqualified_time"`
	PaymentBin                    string     `sql:"payment_bin" json:"payment_bin"`
	Deleted                       bool       `sql:"deleted,notnull" pg:",use_zero" json:"deleted,omitempty"`
	VisaStatusUS                  string     `sql:"visa_status_us" json:"visa_status_us"`
	Timezone                      string     `sql:"timezone" json:"timezone"`
	ReturnFields                  []string   `json:"return_fields"`
	PhoneCode                     string     `json:"phone_code"`
	PartnerCitizen                bool       `sql:"partner_citizen,notnull" json:"partner_citizen"`
	ContactStatus                 string     `sql:"contact_status" json:"contact_status,omitempty"`
	AutoAssign                    bool       `sql:"auto_assign,notnull" json:"auto_assign"`
	IsFake                        bool       `sql:"is_fake,notnull" json:"is_fake"`

	ChangedAttributes []string `sql:"-" json:"changed_attributes,omitempty"`
}

// Make the Leads struct implement the driver.Valuer interface. This method
// simply returns the JSON-encoded representation of the struct.
func (l Leads) Value() (driver.Value, error) {
	return json.Marshal(l)
}

// Make the Attrs struct implement the sql.Scanner interface. This method
// simply decodes a JSON-encoded value into the struct fields.
func (l *Leads) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &l)
}

func (l *Leads) GetHash() string {

	hash := md5.Sum([]byte(l.LeadNumber + l.Email))
	return hex.EncodeToString(hash[:])
}

func (l *Leads) GetUUID() string {

	hash := md5.Sum([]byte(l.LeadNumber))
	return hex.EncodeToString(hash[:])
}

func (l *Leads) TrimFields() {
	reqValuePtr := reflect.ValueOf(l)
	reqValue := reqValuePtr.Elem()
	for i := 0; i < reqValue.NumField(); i++ {
		field := reqValue.Field(i)

		// Ignore fields that don't have the same type as a string
		if field.Type() != reflect.TypeOf("") {
			continue
		}

		str := field.Interface().(string)
		str = strings.TrimSpace(str)
		field.SetString(str)
	}
}

func (l *Leads) SetEmail() {
	l.Email = strings.ToLower(l.Email)
}

func (l *Leads) SetBirthday() {
	if len(l.Birthday) == 4 {
		l.Birthday += "-01-01"
	}
}

func (l *Leads) SetAge() error {
	if len(l.Birthday) == 0 {
		l.Age = 20
		return nil
	}

	layout := "2006-01-02"
	birthdayToDate, err := time.Parse(layout, l.Birthday)
	if err != nil {
		return err
	}

	now := time.Now()

	diff := now.Sub(birthdayToDate)

	hours := int(diff.Hours())

	days := hours / 24

	l.Age = days / 365

	return nil
}

func (l *Leads) SetAssignable(crmAPI, secretKey string) {
	residenceCountryToLower := strings.ToLower(l.CountryOfResidence)
	birthCountryToLower := strings.ToLower(l.CountryOfBirth)

	data, _ := json.Marshal(map[string]string{
		"country_of_residence": residenceCountryToLower,
		"country_of_birth":     birthCountryToLower,
		"campaign_source":      l.CampaignSource,
		"email":                l.Email,
		"age":                  strconv.Itoa(l.Age),
	})

	payload := strings.NewReader(string(data))
	client := &http.Client{}
	req, _ := http.NewRequest("POST", crmAPI+"/api/assignable_settings/filter_lead", payload)

	req.Header.Add("secret_key", secretKey)
	req.Header.Add("Content-Type", "application/json")

	res, _ := client.Do(req)
	body, _ := ioutil.ReadAll(res.Body)

	result := map[string]bool{}
	_ = json.Unmarshal(body, &result)

	l.Assignable = NotAssignable
	if result["assignable"] == true {
		l.Assignable = Assignable
	}
	fmt.Println(l.Assignable)
}

func (l *Leads) SetCampaignSource() {
	if strings.ToLower(l.CountryOfResidence) == "united states" && strings.Contains(strings.ToLower(l.CampaignSource), "mdc") {
		l.CampaignSource = "Bcanadian - " + l.CampaignSource
	}
}

func (l *Leads) SetGender() {
	val := strings.Title(strings.ToLower(string(l.Gender)))
	if utils.InListString(val, []string{string(FemaleGender), string(MaleGender)}) {
		l.Gender = Gender(val)
		return
	}

	switch val {
	case "Man":
		l.Gender = MaleGender
		return
	case "Woman":
		l.Gender = FemaleGender
		return
	default:
		if l.Gender != "" {
			log.Println("Wrong Gender Format: ", l.Gender)
		}

		l.Gender = ""
		return
	}
}

func (l *Leads) NormalizeFields() {
	l.NormalizeStatusReason()
	l.Disqualified = utils.NormalizeField(l.Disqualified)
	l.ContactStatus = utils.NormalizeField(l.ContactStatus)
	l.normalizeSpeakLanguage()
}

func (l *Leads) NormalizeStatusReason() {
	l.StatusReason = utils.NormalizeField(l.StatusReason)
}

func (l *Leads) normalizeSpeakLanguage() {
	if l.SpeakEnglish != "" {
		l.SpeakEnglish = utils.NormalizeField(l.SpeakEnglish)

		if !utils.InListString(l.SpeakEnglish, []string{"Basic", "Moderate", "Proficient", "Poor"}) {
			log.Println("Found wrong statusReason on lead ", l.LeadNumber)
			l.SpeakEnglish = "Basic"
		}
	}

	if l.SpeakFrench != "" {
		l.SpeakFrench = utils.NormalizeField(l.SpeakFrench)

		if !utils.InListString(l.SpeakFrench, []string{"Basic", "Moderate", "Proficient", "Poor"}) {
			log.Println("Found wrong statusReason on lead ", l.LeadNumber)
			l.SpeakFrench = "Basic"
		}
	}
}

func (l *Leads) SetPhone() {
	l.Phone = utils.ReplaceRegex(`^00`, l.Phone)
	l.Phone = utils.ReplaceRegex(`^0`, l.Phone)
	l.Phone = l.PhoneCode + l.Phone

	l.Phone = "+" + strings.ReplaceAll(l.Phone, "+", "")
}

// SetCountry corrects Country, CountryOfBirth and CountryOfResidence fields to set them in db in manageable way
func (l *Leads) SetCountry(r *redis.Client) {
	issuedCountries := CountriesCorrections

	countries, err := GetListOfCountries(r)
	if err == nil {
		issuedCountries = countries.IssuedCountry
	}

	var countryChecked, countryOfBirthChecked, countryOfResidenceChecked bool

	var country string
	if countries != nil {
		// check whether country is correct
		countryChecked, country = utils.InCountryMap(l.Country, countries.Countries)
		if countryChecked {
			l.Country = country
		}

		countryOfBirthChecked, country = utils.InCountryMap(l.CountryOfBirth, countries.Countries)
		if countryOfBirthChecked {
			l.CountryOfBirth = country
		}

		countryOfResidenceChecked, country = utils.InCountryMap(l.CountryOfResidence, countries.Countries)
		if countryOfResidenceChecked {
			l.CountryOfResidence = country
		}

		// check whether countries are passed in iso format
		if val, ok := countries.Countries[strings.ToUpper(l.Country)]; ok {
			countryChecked = true
			l.Country = val
		}
		if val, ok := countries.Countries[strings.ToUpper(l.CountryOfResidence)]; ok {
			countryOfResidenceChecked = true
			l.CountryOfResidence = val
		}
		if val, ok := countries.Countries[strings.ToUpper(l.CountryOfBirth)]; ok {
			countryOfBirthChecked = true
			l.CountryOfBirth = val
		}
	}

	// do not loop over the map if each country field is already checked
	if !countryChecked || !countryOfBirthChecked || !countryOfResidenceChecked {
		// loop over all issued countries
		for correctCountry, incorrectCountries := range issuedCountries {
			// check whether country field needs to be updated
			if !countryChecked && l.Country == correctCountry {
				countryChecked = true
			} else {
				if utils.InListString(l.Country, incorrectCountries) {
					l.Country = correctCountry
					countryChecked = true
				}
			}

			// check whether country_of_birth field needs to be updated
			if !countryOfBirthChecked && l.CountryOfBirth == correctCountry {
				countryOfBirthChecked = true
			} else {
				if utils.InListString(l.CountryOfBirth, incorrectCountries) {
					l.CountryOfBirth = correctCountry
					countryOfBirthChecked = true
				}
			}

			// check whether country_of_residence field needs to be updated
			if !countryOfResidenceChecked && l.CountryOfResidence == correctCountry {
				countryOfResidenceChecked = true
			} else {
				if utils.InListString(l.CountryOfResidence, incorrectCountries) {
					l.CountryOfResidence = correctCountry
					countryOfResidenceChecked = true
				}
			}
		}
	}

	if !countryChecked {
		log.Println("INCORRECT country FIELD: ", l.Country)
		l.Country = utils.NormalizeCountryWithSpace(l.Country)
	}

	if !countryOfResidenceChecked {
		log.Println("INCORRECT country_of_residence FIELD: ", l.CountryOfResidence)
		l.CountryOfResidence = utils.NormalizeCountryWithSpace(l.CountryOfResidence)
	}

	if !countryOfBirthChecked {
		log.Println("INCORRECT country_of_birth FIELD: ", l.CountryOfBirth)
		l.CountryOfBirth = utils.NormalizeCountryWithSpace(l.CountryOfBirth)
	}
}

// GetListOfCountries retrieves list of countries from redis
func GetListOfCountries(r *redis.Client) (*config.CountriesList, error) {
	res, err := r.Get(config.CountriesListRedisKey).Result()
	if err != nil {
		return nil, err
	}

	countries := new(config.CountriesList)

	err = json.Unmarshal([]byte(res), &countries)
	if err != nil {
		return nil, err
	}

	return countries, nil
}

func (l *Leads) SetTimeZone() error {
	// get country data by its name
	country, err := gountries.New().FindCountryByName(strings.ToLower(l.CountryOfResidence))
	if err != nil {
		return err
	}

	// get TimeZone by coordinates
	zone, err := tz.GetZone(tz.Point{
		Lon: country.Coordinates.Longitude,
		Lat: country.Coordinates.Latitude,
	})

	loc, err := time.LoadLocation(zone[0])
	if err != nil {
		panic(err)
	}

	// getting offset to calculate GMT
	_, offset := time.Now().In(loc).Zone()

	var offsetHours = offset / 3600

	var gmtHours = strconv.Itoa(offsetHours) + ":00"

	if offsetHours > 0 {
		gmtHours = "+" + gmtHours
	}

	l.Timezone = fmt.Sprintf("(GMT %v) %v", gmtHours, zone[0])

	return nil
}

var notAssignableCampaignSources = []string{
	"immigration funnel cv2",
	"payment funnel cv v2",
	"payment funnel v2",
	"immigration funnel mdc-2",
	"bcanadian online funnel",
	"immigration funnel cv10",
	"new immi cv",
	"bcanadian online funnel v2",
	"test",
}

var notAssignableCountriesWithAge = []string{
	"malaysia",
	"singapore",
	"china",
	"hong kong",
	"japan",
	"vietnam",
	"thailand",
	"korea (south)",
	"cambodia",
	"saudi arabia",
	"philppines",
	"united arab emirates",
}

var notAssignableCountries = []string{
	"afghanistan",
	"angola",
	"azerbaijan",
	"bangladesh",
	"benin",
	"bhutan",
	"burundi",
	"cameroon",
	"chad",
	"congo",
	"cuba",
	"cyprus",
	"djibouti",
	"equatorial guinea",
	"eritrea",
	"ethiopia",
	"gabon",
	"gambia",
	"ghana",
	"guinea",
	"honduras",
	"india",
	"iran",
	"iraq",
	"ivory coast",
	"kazakhstan",
	"kosovo",
	"lesotho",
	"liberia",
	"madagascar",
	"malawi",
	"myanmar",
	"algeria",
	"nepal",
	"nicaragua",
	"niger",
	"nigeria",
	"pakistan",
	"rwanda",
	"sierra leone",
	"somalia",
	"sri lanka",
	"sudan",
	"swaziland",
	"togo",
	"turkmenistan",
	"uganda",
	"ukraine",
	"uzebekistan",
	"uzbekistan",
	"vanuatu",
	"venezuela",
	"yemen",
	"zimbabwe",
	"zambia",
	"canada",
	"usa",
	"morocco",
	"mozambique",
	"macedonia",
	"albania",
	"botswana",
	"kenya",
	"tanzania",
	"indonesia",
	"argentina",
	"dominican republic",
	"jamica",
	"south africa",
	"haiti",
	"panama",
	"namibia",
	"ghana",
	"somalia",
	"ethiopia",
	"belarus",
	"jamaica",
	"el salvador",
	"united states",
	"serbia",
	"israel",
}

var CountriesCorrections = map[string][]string{
	"Congo":                            {"Congo (DRC)", "Congo, The Democratic Republic Of The", "Congo, Democratic Republic of the", "Congo, Republic of the", "Congo, the Democratic Republic of the", "Congo, The Democratic Republic of the"},
	"Ivory Coast":                      {"Cote D'ivoire", "C&ocirc;te d'Ivoire", "Côte", "Côte d’Ivoi", "Cote d'Ivoire", "Cote d'Ivoire", "Cote D'Ivoire", "CoteD'Ivoire", "Côte d'Ivoire", "Côte d’Ivoire"},
	"Syria":                            {"Syrian Arab Republic"},
	"Aland Islands":                    {"Åland Islands"},
	"Cyprus":                           {"Akrotiri", "Cyprus (Republic of)"},
	"Angola":                           {"Ango"},
	"Australia":                        {"Austral"},
	"Bahamas":                          {"Baham"},
	"Bangladesh":                       {"Bangladesh,"},
	"India":                            {"Bassas da India", "India (Republic of)"},
	"Bhutan":                           {"Bhutan (Kingdom of)"},
	"Bolivia":                          {"Bolivia, Plurinational State of"},
	"Brazil":                           {"Brasil"},
	"Brunei Darussalam":                {"Brun"},
	"Myanmar":                          {"Burma", "Myanmar (Republic of the Union of)"},
	"Burundi":                          {"Burundi (Republic of)"},
	"Canada":                           {"Cana"},
	"Colombia":                         {"Colombia (Republic of)", "Colomb"},
	"Costa Rica":                       {"Costa Ri"},
	"Curacao":                          {"Curaçao", "Cura&ccedil;ao"},
	"Ecuador":                          {"Ecuad"},
	"France":                           {"Fran"},
	"Gambia":                           {"Gambia, The"},
	"Palestine":                        {"Gaza Strip", "Palestinian Territories", "Palestinian Territory, Occupied"},
	"Guernsey":                         {"Guerns"},
	"Guyana":                           {"Guya"},
	"Jordan":                           {"Hashemite Kingdom of Jordan", "West Bank", "WestBank"},
	"Honduras":                         {"Hondur"},
	"Hong Kong":                        {"Hong Kong S.A.R."},
	"Indonesia":                        {"Indones"},
	"Iran":                             {"Iran (Islamic Republic of)", "Iran, Islamic Republic Of", "Iraq (Republic of)"},
	"Ireland":                          {"Irela"},
	"Israel":                           {"Israel (State of)"},
	"Jamaica":                          {"Jamai"},
	"Kenya":                            {"Ken", "Kenya (Republic of)"},
	"Korea (South)":                    {"Korea South", "Korea", "Korea, South", "Korea,South", "South Korea"},
	"Korea (North)":                    {"Korea, Democratic People's Republic of", "Korea, Democratic People's Republic of", "Korea, North"}, // error in segm
	"Laos":                             {"Lao People's Democratic Republic", "Lao People\\'s Democratic Republic"},
	"Spain":                            {"l'Espagne"},
	"Liberia":                          {"Liber"},
	"Libya":                            {"Libyan Arab Jamahiriya"},
	"Macau":                            {"Macao"},
	"Macedonia":                        {"Macedonia, the former Yugoslav Republic of"},
	"Malaysia":                         {"Malays"},
	"Maldives":                         {"Maldiv", "Maldives (Republic of)"},
	"Mexico":                           {"Méjico", "México"}, // empty in segm
	"United States":                    {"Navassa Island", "United Stat", "United States of America"},
	"New Zealand":                      {"New Zeala"},
	"Nigeria":                          {"Nigeria (Federal Republic of)"},
	"China":                            {"Paracel Islands"},
	"Philippines":                      {"Philippin", "Philippines (Republic of the)"},
	"Portugal":                         {"Portug"},
	"Qatar":                            {"Qatar (State of)"},
	"Reunion":                          {"Réunion"},
	"Russia":                           {"Russian Federation"},
	"Rwanda":                           {"Rwan", "Rwanda (Republic of)"},
	"Saint Barthelemy":                 {"Saint Bartelemey"},
	"Saint Vincent and the Grenadines": {"Saint Vincent And Grenadines", "SaintVincentandtheGrenadines"},
	"Saudi Arabia":                     {"Saudi", "Saudi Arabia (Kingdom of)"},
	"Singapore":                        {"Singapo", "Singapore (Republic of)"},
	"Slovakia":                         {"Slovak Republic"},
	"South Africa":                     {"South Afri", "South Africa (Republic of)"},
	"South Georgia and the South Sandwich Islands": {"South Georgia And Sandwich Isl."},
	"Sudan (South)":            {"Sudan South", "South Sudan"},
	"Taiwan":                   {"Taiwan, Province of China"},
	"Tanzania":                 {"Tanzania (United Republic of)", "Tanzania, United Republic of"},
	"East Timor":               {"Timor-Leste"},
	"Trinidad and Tobago":      {"Trinidad and Toba"},
	"Turks and Caicos Islands": {"Turks and Caicos Islan", "Turks And Caicos Islands"},
	"Uganda":                   {"Uganda (Republic of)"},
	"Ukraine":                  {"Ukraine (Україна)"},
	"United Kingdom":           {"United Kingd", "United Kingdom of Great Britain and Northern Ireland"},
	"Venezuela":                {"Vénézuéla", "Venezuela, Bolivarian Republic of"},
	"Vietnam":                  {"Viet Nam", "vietnham"},
	"Virgin Islands British":   {"Virgin Islands, British"},
	"Zambia":                   {"Zamb"},
}
