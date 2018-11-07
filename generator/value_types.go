package generator

import "fmt"

type PropertyValue interface {
	Parameters() []string
	IsSelector() bool
	AllowedValues() []string
	IsSetByFeatureFile() bool
}

type SimpleType interface {
	Parameters() []string
	IsSelector() bool
}

type SimpleString string

func (s SimpleString) Parameters() []string {
	return []string{fmt.Sprintf("%v", s)}
}

func (s SimpleString) IsSelector() bool {
	return false
}

func (s SimpleString) AllowedValues() []string {
	return nil
}

func (s SimpleString) IsSetByFeatureFile() bool {
	return false
}

type SimpleBoolean bool

func (s SimpleBoolean) Parameters() []string {
	return []string{fmt.Sprintf("%v", s)}
}

func (s SimpleBoolean) IsSelector() bool {
	return false
}

func (s SimpleBoolean) AllowedValues() []string {
	return nil
}

func (s SimpleBoolean) IsSetByFeatureFile() bool {
	return false
}

type SimpleInteger int

func (s SimpleInteger) Parameters() []string {
	return []string{fmt.Sprintf("%v", s)}
}

func (s SimpleInteger) IsSelector() bool {
	return false
}

func (s SimpleInteger) AllowedValues() []string {
	return nil
}

func (s SimpleInteger) IsSetByFeatureFile() bool {
	return false
}

type SelectorValue struct {
	Value string   `yaml:"value"`
	AV    []string `yaml:"-"`
}

func (s *SelectorValue) Parameters() []string {
	return []string{fmt.Sprintf("selector -> %s", s.Value)}
}
func (s *SelectorValue) IsSelector() bool {
	return true
}
func (s *SelectorValue) AllowedValues() []string {
	return s.AV
}
func (s *SelectorValue) IsSetByFeatureFile() bool {
	return false
}

type SimpleValue struct {
	Value            string   `yaml:"value"`
	AV               []string `yaml:"-"`
	SetByFeatureFile bool     `yaml:"-"`
}

func (s *SimpleValue) Parameters() []string {
	return []string{s.Value}
}

func (s *SimpleValue) IsSelector() bool {
	return false
}

func (s *SimpleValue) AllowedValues() []string {
	return s.AV
}

func (s *SimpleValue) IsSetByFeatureFile() bool {
	return s.SetByFeatureFile
}

type SimpleCredentialValueHolder struct {
	Value *SimpleCredentialValue `yaml:"value"`
}

func (s *SimpleCredentialValueHolder) Parameters() []string {
	return []string{s.Value.Password, s.Value.Identity}
}

func (s *SimpleCredentialValueHolder) IsSelector() bool {
	return false
}

func (s *SimpleCredentialValueHolder) AllowedValues() []string {
	return nil
}

func (s *SimpleCredentialValueHolder) IsSetByFeatureFile() bool {
	return false
}

type SecretValueHolder struct {
	Value *SecretValue `yaml:"value"`
}

func (s *SecretValueHolder) Parameters() []string {
	return []string{s.Value.Value}
}

func (s *SecretValueHolder) IsSelector() bool {
	return false
}

func (s *SecretValueHolder) AllowedValues() []string {
	return nil
}

func (s *SecretValueHolder) IsSetByFeatureFile() bool {
	return false
}

type SecretValue struct {
	Value string `yaml:"secret"`
}

func (s *SecretValue) Parameters() []string {
	return []string{s.Value}
}

func (s *SecretValue) IsSelector() bool {
	return false
}

func (s *SecretValue) AllowedValues() []string {
	return nil
}

func (s *SecretValue) IsSetByFeatureFile() bool {
	return false
}

type SimpleCredentialValue struct {
	Identity string `yaml:"identity"`
	Password string `yaml:"password"`
}

type CertificateValueHolder struct {
	Value *CertificateValue `yaml:"value"`
}

func (s *CertificateValueHolder) Parameters() []string {
	return []string{s.Value.CertPem, s.Value.CertPrivateKey}
}

func (s *CertificateValueHolder) IsSelector() bool {
	return false
}

func (s *CertificateValueHolder) AllowedValues() []string {
	return nil
}

func (s *CertificateValueHolder) IsSetByFeatureFile() bool {
	return false
}

type CertificateValue struct {
	CertPem        string `yaml:"cert_pem"`
	CertPrivateKey string `yaml:"private_key_pem"`
}

func (s *CertificateValue) Parameters() []string {
	return []string{s.CertPem, s.CertPrivateKey}
}

func (s *CertificateValue) IsSelector() bool {
	return false
}

func (s *CertificateValue) AllowedValues() []string {
	return nil
}

func (s *CertificateValue) IsSetByFeatureFile() bool {
	return false
}

type CollectionsPropertiesValueHolder struct {
	Value []map[string]SimpleType `yaml:"value"`
}

func (s *CollectionsPropertiesValueHolder) Parameters() []string {
	var parameters []string
	for _, paramMap := range s.Value {
		for _, param := range paramMap {
			parameters = append(parameters, param.Parameters()...)
		}
	}

	return parameters
}

func (s *CollectionsPropertiesValueHolder) IsSelector() bool {
	return false
}

func (s *CollectionsPropertiesValueHolder) AllowedValues() []string {
	return nil
}

func (s *CollectionsPropertiesValueHolder) IsSetByFeatureFile() bool {
	return false
}
