package awsgetsecret

import (
	"fmt"
	"time"

	"encoding/base64"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/project-flogo/core/activity"
)

const (
	ivAwsAccessKeyID     = "awsAccessKeyID"
	ivAwsSecretAccessKey = "awsSecretAccessKey"
	ivAwsSessionToken    = "awsSessionToken"
	ivAwsRegion          = "awsRegion"
	ivSecretARN          = "secretARN"
	ivAssumeRole         = "assumeRole"
	ivRoleARN            = "roleARN"
	ivRoleSessionName    = "roleSessionName"
	ovResult             = "secret"
)

type Activity struct {
	//metadata *activity.Metadata
}

func init() {
	_ = activity.Register(&Activity{}, New)
}

var activityMd = activity.ToMetadata(&Input{}, &Output{})

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// New create a new  activity
func New(ctx activity.InitContext) (activity.Activity, error) {

	ctx.Logger().Info("In New activity")

	act := &Activity{}
	return act, nil
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *Activity) Eval(context activity.Context) (done bool, err error) {

	// Get the values from Flogo
	context.Logger().Debug("Get Parms")
	awsSecretARN := context.GetInput(ivSecretARN).(string)
	awsRegion := context.GetInput(ivAwsRegion).(string)
	assumerole := context.GetInput(ivAssumeRole).(bool)

	// AWS Credentials, only if needed
	context.Logger().Debug("use Credentials")
	var awsAccessKeyID, awsSecretAccessKey, awsSessionToken = "", "", ""
	if context.GetInput(ivAwsAccessKeyID) != nil {
		awsAccessKeyID = context.GetInput(ivAwsAccessKeyID).(string)
	}
	if context.GetInput(ivAwsSecretAccessKey) != nil {
		awsSecretAccessKey = context.GetInput(ivAwsSecretAccessKey).(string)
	}
	if context.GetInput(ivAwsSessionToken) != nil {
		awsSessionToken = context.GetInput(ivAwsSessionToken).(string)
	}

	// Create a session with Credentials only if they are set
	context.Logger().Debug("Create Session")
	var awsSession *session.Session
	if awsAccessKeyID != "" && awsSecretAccessKey != "" {
		// Create new credentials using the accessKey and secretKey
		awsCredentials := credentials.NewStaticCredentials(awsAccessKeyID, awsSecretAccessKey, awsSessionToken)

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
		context.Logger().Debug("Assume role")
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
	context.Logger().Debug("Connect to Secrets manager")
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

	context.Logger().Debug("Result")
	context.Logger().Debug(result)
	context.Logger().Debug("Secret")
	context.Logger().Debug(secretString)
	// Set the output value in the context
	context.SetOutput(ovResult, secretString)

	return true, nil
}
