// Package amazons3 uploads or downloads files from Amazon Simple Storage Service (S3)
package sagemaker

import (
	"fmt"
	"time"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sagemaker"
	"github.com/aws/aws-sdk-go/service/sagemakerruntime"
)

const (
	ivAwsAccessKeyID     = "awsAccessKeyID"
	ivAwsSecretAccessKey = "awsSecretAccessKey"
	ivAwsRegion          = "awsRegion"
	ivEndPoint           = "endpointname"
	ivBody               = "body"
	ivContentType        = "contenttype"
	ivAssumeRole         = "assumeRole"
	ivRoleARN            = "roleARN"
	ivRoleSessionName    = "roleSessionName"
	ovResult             = "result"
)

const (
	ServiceName = "runtime.sagemaker" // Name of service.
	EndpointsID = ServiceName         // ID to lookup a service endpoint with.
	ServiceID   = "SageMaker Runtime" // ServiceID is a unique identifier of a specific service.
)

const (

	// ErrCodeInternalFailure for service response error code
	// "InternalFailure".
	//
	// An internal failure occurred.
	ErrCodeInternalFailure = "InternalFailure"

	// ErrCodeModelError for service response error code
	// "ModelError".
	//
	// Model (owned by the customer in the container) returned 4xx or 5xx error
	// code.
	ErrCodeModelError = "ModelError"

	// ErrCodeServiceUnavailable for service response error code
	// "ServiceUnavailable".
	//
	// The service is unavailable. Try your call again.
	ErrCodeServiceUnavailable = "ServiceUnavailable"

	// ErrCodeValidationError for service response error code
	// "ValidationError".
	//
	// Inspect your request and try again.
	ErrCodeValidationError = "ValidationError"
)

// log is the default package logger
var log = logger.GetLogger("activity-sagemaker")

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
	fmt.Println("Get Parms")
	endpoint := context.GetInput(ivEndPoint).(string)
	awsRegion := context.GetInput(ivAwsRegion).(string)
	body := context.GetInput(ivBody).(string)
	contenttype := context.GetInput(ivContentType).(string)
	assumerole := context.GetInput(ivAssumeRole).(bool)

	// AWS Credentials, only if needed
	fmt.Println("use Credentials")
	var awsAccessKeyID, awsSecretAccessKey = "", ""
	if context.GetInput(ivAwsAccessKeyID) != nil {
		awsAccessKeyID = context.GetInput(ivAwsAccessKeyID).(string)
	}
	if context.GetInput(ivAwsSecretAccessKey) != nil {
		awsSecretAccessKey = context.GetInput(ivAwsSecretAccessKey).(string)
	}

	// Create a session with Credentials only if they are set
	fmt.Println("Create Session")
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
		fmt.Println("assume role")
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
	fmt.Println("Connect to sagemaker runtime")
	// Create a SageMakerRuntime client from just a session.
	svc := sagemakerruntime.New(awsSession)

	fmt.Println("Connect to sagemaker")
	s := sagemaker.New(awsSession)
	dei := sagemaker.DescribeEndpointInput{}
	dei.SetEndpointName(endpoint)
	info, err := s.DescribeEndpoint(&dei)
	if err != nil {
		log.Info(err)
		fmt.Println(err)
		return false, err
	}
	log.Debug(info)
	fmt.Println("describe output")
	fmt.Println(info)

	//body := `{ "instances": [ { "start": "2018-03-13 00:00:00", "target": [100, 12] } ] }`
	params := sagemakerruntime.InvokeEndpointInput{}
	params.SetContentType(contenttype)
	params.SetBody([]byte(body))
	params.SetEndpointName(endpoint)
	fmt.Println("describe output")
	result, err := svc.InvokeEndpoint(&params)
	if err != nil {
		fmt.Println("Invoke Error")
		fmt.Println(err)
		log.Info(err)
		return false, err
	}
	fmt.Println("Result")
	resultBody := string(result.Body)
	fmt.Println(result)
	fmt.Println("Result Body")
	fmt.Println(resultBody)
	// Set the output value in the context
	context.SetOutput(ovResult, resultBody)

	return true, nil
}
