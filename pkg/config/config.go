package config

import (
	"github.com/sm0kernone/crm-common/pkg/aws"
	"github.com/sm0kernone/crm-common/pkg/curl"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/go-redis/redis"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	CommonPath            = "/common.%v.yaml"
	ConfigPath            = "/%v/conf.%v.yaml"
	CountriesListRedisKey = "countries_list"
)

// LoadFromFile loads config from file
func loadFromFile(path string) (*Configuration, error) {
	fmt.Printf("loading config file: %s\n", path)
	var cfg = new(Configuration)
	var files []string

	root := "./"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fmt.Println(file)
	}

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading config file, %s", err)
	}

	if err := yaml.Unmarshal(bytes, cfg); err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %v", err)
	}

	fmt.Println("here")

	return cfg, err
}

func loadCommonFromS3(cfg *Configuration, awsSession *session.Session) error {
	bytes, err := aws.Download(awsSession, cfg.AwsConfigs.Bucket, fmt.Sprintf(CommonPath, cfg.Common.Environment))
	if err != nil {
		return fmt.Errorf("error reading config file from s3, %s", err)
	}

	if err := yaml.Unmarshal(bytes, cfg); err != nil {
		return fmt.Errorf("unable to decode into struct, %v", err)
	}

	return nil
}

func loadFromS3(cfg *Configuration, awsSession *session.Session, msName string) error {

	bytes, err := aws.Download(awsSession, cfg.AwsConfigs.Bucket, fmt.Sprintf(ConfigPath, msName, cfg.Common.Environment))
	if err != nil {
		return fmt.Errorf("error reading config file from s3, %s", err)
	}

	if err := yaml.Unmarshal(bytes, cfg); err != nil {
		return fmt.Errorf("unable to decode into struct, %v", err)
	}

	return nil
}

// LoadConfigs function to load configs for service,
// params:
// path - path to local file to configs for aws
// msName - name of microservice - folderName in aws bucket with configs for this ms - set "" if there is no separate config for this ms
func LoadConfigs(path, msName string) (*Configuration, error) {
	configs, err := loadFromFile(path)
	if err != nil {
		return nil, err
	}

	fmt.Println(configs.AwsConfigs.Region)

	awsSession := aws.GetAwsSession(configs.AwsConfigs.Region, configs.AwsConfigs.AccessKeyId, configs.AwsConfigs.SecretKey)

	err = loadCommonFromS3(configs, awsSession)
	if err != nil {
		return nil, err
	}

	if len(msName) > 0 {
		err = loadFromS3(configs, awsSession, msName)
		if err != nil {
			return nil, err
		}
	}

	return configs, nil
}

func SetCountriesList(client *redis.Client, filePath string) error {
	resp, err := loadCountriesList(filePath)
	if err != nil {
		return err
	}

	return setCountriesToRedis(client, resp)
}

// loadCountriesList loads CountriesList from s3
func loadCountriesList(path string) ([]byte, error) {
	return curl.Get(path)
}

// setCountriesToRedis sets countries to redis
func setCountriesToRedis(client *redis.Client, countries []byte) error {
	return client.Set(CountriesListRedisKey, string(countries), 0).Err()
}

type Configuration struct {
	LeadMS                 LeadsCreatorMS         `yaml:"leads_ms"`
	QueryMS                QueryExecutorMS        `yaml:"query_ms"`
	SIPMs                  SIPMs                  `yaml:"sip_ms"`
	CallActivityMS         CallActivityMS         `yaml:"call_activity_ms"`
	AssignerMS             AssignerMS             `yaml:"assigner_ms"`
	CsrPointsMS            CsrPointsMS            `yaml:"csr_points_ms"`
	NotificationSystem     NotificationSystem     `yaml:"notification_system"`
	ScheduledNotifications ScheduledNotifications `yaml:"scheduled_notifications"`
	DashboardAPI           DashboardAPI           `yaml:"dashboard_api"`
	Common                 Common                 `yaml:"common"`
	AwsConfigs             AwsConfigs             `yaml:"aws_configs"`
	NsqAuth                NSQAuth                `yaml:"nsq_auth"`
	ExponeaReport          ExponeaReport          `yaml:"exponea_report"`
	DashboardCacheUpdater  DashboardCacheUpdater  `yaml:"dashboard_cache_updater"`
	DashboardAggregator    DashboardAggregator    `yaml:"dashboard_aggregator"`
	FreshworkMS            FreshworkMS            `yaml:"freshwork_ms"`
}

type AwsConfigs struct {
	AccessKeyId   string `yaml:"access_key_id"`
	SecretKey     string `yaml:"secret_key"`
	Region        string `yaml:"region"`
	Bucket        string `yaml:"bucket"`
	CloudFrontURL string `yaml:"cloud_front_url"`
}

type Common struct {
	Newrelic             Newrelic      `yaml:"newrelic"`
	DbPSNs               []string      `yaml:"db_psns"`
	HistoryDbPSN         string        `yaml:"history_psn"`
	Redis                Redis         `yaml:"redis"`
	CallActivityCloner   ClonerWorker  `yaml:"call_activity_cloner"`
	ApplicationsCloner   ClonerWorker  `yaml:"applications_cloner"`
	DocumentsCloner      ClonerWorker  `yaml:"documents_cloner"`
	LeadsCloner          ClonerWorker  `yaml:"leads_cloner"`
	OpportunitiesCloner  ClonerWorker  `yaml:"opportunities_cloner"`
	VisasCloner          ClonerWorker  `yaml:"visas_cloner"`
	RequirementsCloner   ClonerWorker  `yaml:"requirements_cloner"`
	PaymentHistoryCloner ClonerWorker  `yaml:"payment_history_cloner"`
	TicketsCloner        ClonerWorker  `yaml:"tickets_cloner"`
	InstallmentsCloner   ClonerWorker  `yaml:"installments_cloner"`
	UsersCloner          ClonerWorker  `yaml:"users_cloner"`
	Environment          string        `yaml:"environment"`
	Nsq                  NSQ           `yaml:"nsq"`
	ReviseWorker         ReviseWorker  `yaml:"revise_worker"`
	PaymentWorker        PaymentWorker `yaml:"payment_worker"`
	AssignSystemURL      string        `yaml:"assign_system_url"`
	Segmentation         Segmentation  `yaml:"segmentation"`
}

type ReviseWorker struct {
	Kafka KafkaConfigs `yaml:"kafka"`
}

type PaymentWorker struct {
	NSQ NSQ `yaml:"nsq"`
}

type NotificationSystem struct {
	Newrelic      Newrelic   `yaml:"newrelic"`
	WorkersInPool int        `yaml:"workers_in_pool"`
	Nsq           NSQ        `yaml:"nsq"`
	Server        Server     `yaml:"server"`
	Centrifugo    Centrifugo `yaml:"centrifugo"`
	TTL           string     `yaml:"ttl"`
	SecretKey     string     `yaml:"secret_key"`
}

type ScheduledNotifications struct {
	Newrelic      Newrelic   `yaml:"newrelic"`
	WorkersInPool int        `yaml:"workers_in_pool"`
	Nsq           NSQ        `yaml:"nsq"`
	Server        Server     `yaml:"server"`
	Centrifugo    Centrifugo `yaml:"centrifugo"`
	TTL           string     `yaml:"ttl"`
	SecretKey     string     `yaml:"secret_key"`
	GetTopic      string     `yaml:"get_topic"`
}

type DashboardAPI struct {
	Newrelic   Newrelic   `yaml:"newrelic"`
	Server     Server     `yaml:"server"`
	GrpcClient GRPCClient `yaml:"grpc_client"`
	Redis      Redis      `yaml:"redis"`
}

type NSQAuth struct {
	Server Server `yaml:"server"`
}

type ExponeaReport struct {
	Mandrill Mandrill `yaml:"mandrill"`
}

type DashboardCacheUpdater struct {
	Nsq   NSQ    `yaml:"nsq"`
	Redis Redis  `yaml:"redis"`
	DbPsn string `yaml:"db_psn"`
}

type DashboardAggregator struct {
	Redis           Redis    `yaml:"redis"`
	CrmRedisAddress string   `yaml:"crm_redis_address"`
	DbPsn           string   `yaml:"db_psn"`
	HistoryDBPsn    string   `yaml:"history_db_psn"`
	GrpcPort        int      `yaml:"grpc_port"`
	Newrelic        Newrelic `yaml:"newrelic"`
}

type FreshworkMS struct {
	Server            Server   `yaml:"server"`
	CrmAPI            string   `yaml:"crm_api"`
	SecretKey         string   `yaml:"secret_key"`
	Newrelic          Newrelic `yaml:"newrelic"`
	FreshworkUrl      string   `yaml:"freshwork_url"`
	FreshworkUsername string   `yaml:"freshwork_username"`
	FreshworkPassword string   `yaml:"freshwork_password"`
}

type CsrPointsMS struct {
	Newrelic      Newrelic `yaml:"newrelic"`
	WorkersInPool int      `yaml:"workers_in_pool"`
}

type Centrifugo struct {
	URL    string `yaml:"url"`
	ApiKey string `yaml:"api_key"`
}

type NSQ struct {
	Host                             string `yaml:"host"`
	ProducerPort                     string `yaml:"producer_port"`
	ConsumerPort                     string `yaml:"consumer_port"`
	Topic                            string `yaml:"topic"`
	CsrTopic                         string `yaml:"csr_topic"`
	EventTopic                       string `yaml:"event_topic"`
	LeadsClonerTopic                 string `yaml:"leads_cloner_topic"`
	ApplicationsClonerTopic          string `yaml:"applications_cloner_topic"`
	OpportunitiesClonerTopic         string `yaml:"opportunities_cloner_topic"`
	OpportunitiesClonerCacheChannel  string `yaml:"opportunities_cloner_cache"`
	CallActivityClonerTopic          string `yaml:"call_activity_cloner_topic"`
	DocumentsClonerTopic             string `yaml:"documents_cloner_topic"`
	VisasClonerTopic                 string `yaml:"visas_cloner_topic"`
	RequirementsClonerTopic          string `yaml:"requirements_cloner_topic"`
	PaymentHistoryClonerTopic        string `yaml:"payment_history_cloner_topic"`
	TicketsClonerTopic               string `yaml:"tickets_cloner_topic"`
	InstallmentsClonerTopic          string `yaml:"installments_cloner_topic"`
	UsersClonerTopic                 string `yaml:"users_cloner_topic"`
	PaymentHistoryClonerCacheChannel string `yaml:"payment_history_cloner_cache"`
	RefundsTopic                     string `yaml:"refunds_topic"`
	RefundsCacheChannel              string `yaml:"refunds_cache"`
	PartialsTopic                    string `yaml:"partials_topic"`
	PartialsCacheChannel             string `yaml:"partials_cache"`
	PaymentHistoryWorkerTopic        string `yaml:"payment_history_worker_topic"`
	Channel                          string `yaml:"channel"`
	SecretKey                        string `yaml:"secret_key"`
}

type Segmentation struct {
	BaseURL            string `yaml:"base_url"`
	ApiKey             string `yaml:"api_key"`
	ApiSecret          string `yaml:"api_secret"`
	CallScheduleURL    string `yaml:"call_schedule_url"`
	ApproveScheduleURL string `yaml:"approve_schedule_url"`
	AddToCalendarURL   string `yaml:"add_to_calendar_url"`
}

type KafkaConfigs struct {
	BrokerProducer string   `yaml:"broker_producer"`
	Brokers        []string `yaml:"brokers"`
	GroupID        string   `yaml:"group_id"`
	Topic          string   `yaml:"topic"`
	MinBytes       int      `yaml:"min_bytes"`
	MaxBytes       int      `yaml:"max_bytes"`
	MaxWait        int      `yaml:"max_wait"`
}

type Redis struct {
	Address  string `yaml:"address"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
}

type LeadsCreatorMS struct {
	Server       Server      `yaml:"server"`
	Newrelic     Newrelic    `yaml:"newrelic"`
	Neverbounce  Neverbounce `yaml:"neverbounce"`
	CountriesURL string      `yaml:"countries_url"`
	CrmAPI       string      `yaml:"crm_api"`
	SecretKey    string      `yaml:"secret_key"`
}

type QueryExecutorMS struct {
	Server   Server   `yaml:"server"`
	DbPSNs   []string `yaml:"db_psns"`
	Newrelic Newrelic `yaml:"newrelic"`
}

type SIPMs struct {
	Server                 Server     `yaml:"server"`
	CallActivityGrpcClient GRPCClient `yaml:"call_activity_grpc_client"`
	Sip                    Sip        `yaml:"sip"`
	CRM                    NewCRM     `yaml:"crm"`
	Newrelic               Newrelic   `yaml:"newrelic"`
	Centrifugo             Centrifugo `yaml:"centrifugo"`
	SecretKey              string     `yaml:"secret_key"`
}

type CallActivityMS struct {
	GrpcPort int      `yaml:"grpc_port"`
	Newrelic Newrelic `yaml:"newrelic"`
	Sip      Sip      `yaml:"sip"`
	SipMsURL string   `yaml:"sip_ms_url"`
	CRM      NewCRM   `yaml:"crm"`
}

type AssignerMS struct {
	Server   Server   `yaml:"server"`
	Newrelic Newrelic `yaml:"newrelic"`
	Redis    Redis    `yaml:"redis"`
}

type ClonerWorker struct {
	WorkersInPool    int               `yaml:"workers_in_pool"`
	Recipients       []Recipient       `yaml:"recipients"`
	ExponiaEvents    []ExponiaEvent    `yaml:"exponia_events"`
	Evaluations      []EvaluationEvent `yaml:"evaluations"`
	Newrelic         Newrelic          `yaml:"newrelic"`
	CustomRecipients []CustomRecipient `yaml:"custom_recipients"`
}

type ExponiaEvent struct {
	EventType   ExponiaEventType  `yaml:"event_type"`
	Url         string            `yaml:"url"`
	Headers     map[string]string `yaml:"headers"`
	ContentType string            `yaml:"content_type"`
	Method      string            `yaml:"method"`
	Category    string            `yaml:"category"`
}

type CustomRecipient struct {
	EventType   string `yaml:"event_type"`
	Url         string `yaml:"url"`
	Method      string `yaml:"method"`
	ContentType string `yaml:"content_type"`
}

type ExponiaEventType string

type CustomEventType string

type EvaluationEventType string

var (
	CustomerExponiaEvent         ExponiaEventType = "customer"
	ConsentExponeaEvent          ExponiaEventType = "consent"
	PurchaseExponiaEvent         ExponiaEventType = "purchase"
	PurchaseItemExponiaEvent     ExponiaEventType = "purchase_item"
	CallActivityExponeaEvent     ExponiaEventType = "call_activity"
	ResultSendExponeaEvent       ExponiaEventType = "result_sent"
	DocumentExponeaEvent         ExponiaEventType = "document"
	InstallmentPieceExponeaEvent ExponiaEventType = "installments_piece"
	VisaRCICAssignExoineaEvent   ExponiaEventType = "visa_rcic_assign"
	SentToRCICExponeaEvent       ExponiaEventType = "sent_to_rcic"
	PurchaseInstallmentEvent     ExponiaEventType = "purchase_installment"
	UploadResultExponeaEvent     ExponiaEventType = "upload_result"
)

var (
	IeltsEmailCustomEvent CustomEventType = "ielts_email"
)

var (
	CreateEvaluationEventType     EvaluationEventType = "create"
	CreateUserEvaluationEventType EvaluationEventType = "create_user"
	UpdateUserEvaluationEventType EvaluationEventType = "update_user"
	DeleteUserEvaluationEventType EvaluationEventType = "delete_user"
)

var ExponiaEventsSlice = []string{string(CustomerExponiaEvent), string(PurchaseExponiaEvent), string(PurchaseItemExponiaEvent), string(ConsentExponeaEvent),
	string(CallActivityExponeaEvent), string(ResultSendExponeaEvent), string(DocumentExponeaEvent), string(InstallmentPieceExponeaEvent), string(VisaRCICAssignExoineaEvent), string(SentToRCICExponeaEvent), string(PurchaseInstallmentEvent), string(UploadResultExponeaEvent)}

type Recipient struct {
	URL         string            `yaml:"url"`
	Method      string            `yaml:"method"`
	ContentType string            `yaml:"content_type"`
	Headers     map[string]string `yaml:"headers"`
}

type EvaluationEvent struct {
	URL         string              `yaml:"url"`
	Method      string              `yaml:"method"`
	ContentType string              `yaml:"content_type"`
	Headers     map[string]string   `yaml:"headers"`
	Evaluation  string              `yaml:"evaluation"`
	Type        EvaluationEventType `yaml:"type"`
}

type NewCRM struct {
	CallActivityAPI string `yaml:"call_activity_api"`
	SecretKey       string `yaml:"secret_key"`
}

// Server holds data necessary for server configuration
type Server struct {
	Port         string `yaml:"port"`
	Debug        bool   `yaml:"debug"`
	ReadTimeout  int    `yaml:"read_timeout_seconds"`
	WriteTimeout int    `yaml:"write_timeout_seconds"`
}

type Mandrill struct {
	ApiKey     string   `yaml:"api_key"`
	FromEmail  string   `yaml:"from_email"`
	FromName   string   `yaml:"from_name"`
	Recipients []string `yaml:"recipients"`
}

type Newrelic struct {
	AppName string `yaml:"app_name"`
	License string `yaml:"license"`
}

type Sip struct {
	ApiKey string `yaml:"api_key"`
}

type GRPCClient struct {
	Host                     string `yaml:"host"`
	Port                     int    `yaml:"port"`
	ConnectionBackoffTimeout string `yaml:"connection_backoff_timeout"`
}

type CountriesList struct {
	Countries     map[string]string   `json:"countries"`
	Nationality   []string            `json:"natianality"`
	IssuedCountry map[string][]string `json:"issuedCountry"`
}

type Neverbounce struct {
	BaseURL string `yaml:"base_url"`
	ApiKey  string `yaml:"api_key"`
}
