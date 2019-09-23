package management

const (
	DuoGuardianFactor              GuardianFactorType = "duo"
	EmailGuardianFactor            GuardianFactorType = "email"
	OtpGuardianFactor              GuardianFactorType = "otp"
	PushNotificationGuardianFactor GuardianFactorType = "push-notification"
	SmsGuardianFactor              GuardianFactorType = "sms"
)

type GuardianFactorType string

type GuardianFactor struct {
	// States if this factor is enabled
	Enabled *bool `json:"enabled,omitempty"`
}

type GuardianFactorStatus struct {
	// States if this factor is enabled
	Enabled *bool `json:"enabled,omitempty"`

	// Guardian Factor name
	Name *string `json:"name,omitempty"`

	// For factors with trial limits (e.g. SMS) states if those limits have been exceeded
	TrialExpired *bool `json:"trial_expired,omitempty"`
}

type GuardianFactorPushNotificationSnsConfig struct {
	// AWS Access Key ID
	AwsAccessKeyID *string `json:"aws_access_key_id,omitempty"`

	// AWS Secret Access Key ID
	AwsSecretAccessKeyID *string `json:"aws_secret_access_key,omitempty"`

	// AWS Region
	AwsRegion *string `json:"aws_region,omitempty"`

	// SNS APNS Platform Application ARN
	SnsApnsPlatformApplicationArn *string `json:"sns_apns_platform_application_arn,omitempty"`

	// SNS GCM Platform Application ARN
	SnsGcmPlatformApplicationArn *string `json:"sns_gcm_platform_application_arn,omitempty"`
}

type GuardianFactorSmsTwilioConfig struct {
	// From number
	From *string `json:"from,omitempty"`

	// Copilot SID
	MessagingServiceSid *string `json:"messaging_service_sid,omitempty"`

	// Twilio Authentication token
	AuthToken *string `json:"auth_token,omitempty"`

	// Twilio SID
	Sid *string `json:"sid,omitempty"`
}

type GuardianManager struct {
	m *Management
}

func NewGuardianManager(m *Management) *GuardianManager {
	return &GuardianManager{m}
}

func (gm *GuardianManager) ListFactorsAndStatuses() (gf []*GuardianFactorStatus, err error) {
	err = gm.m.get(gm.m.uri("guardian/factors"), gf)
	return
}

func (gm *GuardianManager) UpdateFactor(name GuardianFactorType, gf *GuardianFactor) error {
	return gm.m.put(gm.m.uri("guardian/factors", string(name)), gf)
}

func (gm *GuardianManager) GetFactorPushNotificationSnsConfig() (sc *GuardianFactorPushNotificationSnsConfig, err error) {
	err = gm.m.get(gm.m.uri("guardian/factors/push-notification/providers/sns"), sc)
	return
}

func (gm *GuardianManager) UpdateFactorPushNotificationSnsConfig(sc *GuardianFactorPushNotificationSnsConfig) error {
	return gm.m.put(gm.m.uri("guardian/factors/push-notification/providers/sns"), sc)
}

func (gm *GuardianManager) GetFactorSmsTwilioConfig() (tc *GuardianFactorSmsTwilioConfig, err error) {
	err = gm.m.get(gm.m.uri("guardian/factors/sms/providers/twilio"), tc)
	return
}

func (gm *GuardianManager) UpdateFactorSmsTwilioConfig(tc *GuardianFactorSmsTwilioConfig) error {
	return gm.m.put(gm.m.uri("guardian/factors/sms/providers/twilio"), tc)
}
