package awsgetsecret

import (
	"fmt"
	"time"

	"encoding/base64"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

const (
	ivAwsAccessKeyID     = "awsAccessKeyID"
	ivAwsSecretAccessKey = "awsSecretAccessKey"
	ivAwsRegion          = "awsRegion"
	ivSecretARN          = "secretARN"
	ivAssumeRole         = "assumeRole"
	ivRoleARN            = "roleARN"
	ivRoleSessionName    = "roleSessionName"
	ovResult             = "secret"
)

// log is the default package logger
var log = logger.GetLogger("activity-awsgetsecret")

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	// Get the values from Flogo
	log.Debug("Get Parms")
	awsSecretARN := context.GetInput(ivSecretARN).(string)
	awsRegion := context.GetInput(ivAwsRegion).(string)
	assumerole := context.GetInput(ivAssumeRole).(bool)

	// AWS Credentials, only if needed
	log.Debug("use Credentials")
	var awsAccessKeyID, awsSecretAccessKey = "", ""
	if context.GetInput(ivAwsAccessKeyID) != nil {
		awsAccessKeyID = context.GetInput(ivAwsAccessKeyID).(string)
	}
	if context.GetInput(ivAwsSecretAccessKey) != nil {
		awsSecretAccessKey = context.GetInput(ivAwsSecretAccessKey).(string)
	}

	// Create a session with Credentials only if they are set
	log.Debug("Create Session")
	var awsSession *session.Session
	if awsAccessKeyID != "" && awsSecretAccessKey != "" {
		// Create new credentials using the accessKey and secretKey
		awsCredentials := credentials.NewStaticCredentials(awsAccessKeyID, awsSecretAccessKey, "")

		// Create a new session with AWS credentials
		awsSession = session.Must(session.NewSession(&aws.Config{
			Credentials: awsCredentials,
			Region:      aws.String(awsRegion),
		}))
	} else {
		// Create a new session without AWS credentials
		awsSession = session.Must(session.NewSession(&aws.Config{
			Region: aws.String(awsRegion),
		}))
	}
	// AssumeRole processing
	if assumerole {
		log.Debug("Assume role")
		rolearn := context.GetInput(ivRoleARN).(string)
		rolesessioname := context.GetInput(ivRoleSessionName).(string)

		awsSession.Config.Credentials = stscreds.NewCredentials(awsSession, rolearn, func(p *stscreds.AssumeRoleProvider) {
			// if len(k.config.ExternalID) > 0 {
			// 	p.ExternalID = aws.String(k.config.ExternalID)
			// }
			p.RoleSessionName = rolesessioname
			p.Duration = time.Duration(900) * time.Second
		})
	}
	log.Debug("Connect to Secrets manager")
	svc := secretsmanager.New(awsSession,
		aws.NewConfig().WithRegion(awsRegion))
	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(awsSecretARN),
		VersionStage: aws.String("AWSCURRENT"), // VersionStage defaults to AWSCURRENT if unspecified
	}

	result, err := svc.GetSecretValue(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case secretsmanager.ErrCodeDecryptionFailure:
				// Secrets Manager can't decrypt the protected secret text using the provided KMS key.
				fmt.Println(secretsmanager.ErrCodeDecryptionFailure, aerr.Error())

			case secretsmanager.ErrCodeInternalServiceError:
				// An error occurred on the server side.
				fmt.Println(secretsmanager.ErrCodeInternalServiceError, aerr.Error())

			case secretsmanager.ErrCodeInvalidParameterException:
				// You provided an invalid value for a parameter.
				fmt.Println(secretsmanager.ErrCodeInvalidParameterException, aerr.Error())

			case secretsmanager.ErrCodeInvalidRequestException:
				// You provided a parameter value that is not valid for the current state of the resource.
				fmt.Println(secretsmanager.ErrCodeInvalidRequestException, aerr.Error())

			case secretsmanager.ErrCodeResourceNotFoundException:
				// We can't find the resource that you asked for.
				fmt.Println(secretsmanager.ErrCodeResourceNotFoundException, aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return false, err
	}

	// Decrypts secret using the associated KMS key.
	// Depending on whether the secret is a string or binary, one of these fields will be populated.
	var secretString string
	if result.SecretString != nil {
		secretString = *result.SecretString
	} else {
		decodedBinarySecretBytes := make([]byte, base64.StdEncoding.DecodedLen(len(result.SecretBinary)))
		len, err := base64.StdEncoding.Decode(decodedBinarySecretBytes, result.SecretBinary)
		if err != nil {
			fmt.Println("Base64 Decode Error:", err)
			return false, err
		}
		secretString = string(decodedBinarySecretBytes[:len])
	}

	log.Debug("Result")
	log.Debug(result)
	log.Debug("Secret")
	log.Debug(secretString)
	// Set the output value in the context
	context.SetOutput(ovResult, secretString)

	return true, nil
}
