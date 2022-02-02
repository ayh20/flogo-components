package awsgetsecret

import (
	"github.com/project-flogo/core/data/coerce"
)

// Input data structure
type Input struct {
	AwsAccessKeyID     string `md:"awsAccessKeyID"`
	AwsSecretAccessKey string `md:"awsSecretAccessKey"`
	AssumeRole         string `md:"assumeRole"`
	RoleARN            string `md:"roleARN"`
	RoleSessionName    string `md:"roleSessionName"`
	AwsRegion          string `md:"awsRegion"`
	SecretARN          string `md:"secretARN"`
}

//ToMap Input mapper
func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"awsAccessKeyID":     i.AwsAccessKeyID,
		"awsSecretAccessKey": i.AwsSecretAccessKey,
		"assumeRole":         i.AssumeRole,
		"roleARN":            i.RoleARN,
		"roleSessionName":    i.RoleSessionName,
		"awsRegion":          i.AwsRegion,
		"secretARN":          i.SecretARN,
	}
}

//FromMap Input from map
func (i *Input) FromMap(values map[string]interface{}) error {

	var err error

	i.AwsAccessKeyID, err = coerce.ToString(values["awsAccessKeyID"])
	if err != nil {
		return err
	}
	i.AwsSecretAccessKey, err = coerce.ToString(values["awsSecretAccessKey"])
	if err != nil {
		return err
	}
	i.AssumeRole, err = coerce.ToString(values["assumeRole"])
	if err != nil {
		return err
	}
	i.RoleARN, err = coerce.ToString(values["roleARN"])
	if err != nil {
		return err
	}
	i.RoleSessionName, err = coerce.ToString(values["roleSessionName"])
	if err != nil {
		return err
	}
	i.AwsRegion, err = coerce.ToString(values["awsRegion"])
	if err != nil {
		return err
	}
	i.SecretARN, err = coerce.ToString(values["secretARN"])
	if err != nil {
		return err
	}
	return nil
}

//Output data structure
type Output struct {
	Secret string `md:"secret"`
}

//ToMap Output mapper
func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"secret": o.Secret,
	}
}

//FromMap Output  from map
func (o *Output) FromMap(values map[string]interface{}) error {

	var err error

	o.Secret, err = coerce.ToString(values["secret"])
	if err != nil {
		return err
	}

	return nil
}
