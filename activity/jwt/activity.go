package jwt

import (
	"encoding/json"
	"strings"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/dgrijalva/jwt-go"
)

// activityLog is the default logger for the Log Activity
var activityLog = logger.GetLogger("activity-JWT")

const (
	ivToken   = "token"
	ivHeader  = "header"
	ivPayload = "payload"
	ivSecret  = "secret"
	ivMode    = "mode"
	ivAlgo    = "algorithm"

	ovToken = "token"
	ovValid = "valid"
)

func init() {
	activityLog.SetLogLevel(logger.InfoLevel)
}

// JWT is an Activity that works with JWT Tokens
// It can create, verify and decrypt JWT tokens
// inputs : {input1, input2, datatype, comparemode}
// outputs: result (bool)
type JWT struct {
	metadata *activity.Metadata
}

// NewActivity creates a new AppActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &JWT{metadata: metadata}
}

// Metadata returns the activity's metadata
func (a *JWT) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *JWT) Eval(context activity.Context) (done bool, err error) {

	activityLog.Debug("In Eval")
	// Get the runtime values
	tokenstring, _ := context.GetInput(ivToken).(string)
	header, _ := context.GetInput(ivHeader).(string)
	payload, _ := context.GetInput(ivPayload).(string)
	secret, _ := context.GetInput(ivSecret).(string)
	sharedEncryptionKey := []byte(secret)
	mode, _ := context.GetInput(ivMode).(string)
	algo, _ := context.GetInput(ivAlgo).(string)

	activityLog.Debug(mode, header, payload, tokenstring, secret)

	// Determine code path based on mode
	switch mode {
	case "Decrypt":
		//  Do we need a decrypt and what does it logicall mean anyway !

		/* 		key := []byte(secret)
		   		tok, err := jwt.ParseEncrypted(token)
		   		if err != nil {
		   			activityLog.Info("Parse Error")
		   			activityLog.Info(err.Error)
		   			context.SetOutput("valid", false)
		   			return false, err
		   		}

		   		out := jwt.Claims{}

		   		if err := tok.Claims(key, &out); err != nil {
		   			activityLog.Info("claims error")
		   			activityLog.Info(err.Error)
		   			context.SetOutput("valid", false)
		   			return false, err
		   		}
		   		//out2 := jwt.NestedJSONWebToken.
		   		activityLog.Info(fmt.Sprintf("iss: %s, sub: %s, ID: %s \n", out.Issuer, out.Subject, out.ID))
		   		context.SetOutput("valid", true) */
		return true, nil

	case "Verify":

		activityLog.Debug("In Verify")

		// Parse takes the token string and a function for looking up the key. The latter is especially
		// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
		// head of the token to identify which key to use, but the parsed token (head and claims) is provided
		// to the callback, providing flexibility.

		token, err := jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
			// sharedEncryptionKey contains a plain secret or a public key
			// we used the passed Algo name to determine the handling of the secret
			// returning the formatted ES/RS key or secret string
			if isEs(algo) {
				return jwt.ParseECPublicKeyFromPEM(sharedEncryptionKey)
			} else if isRs(algo) {
				return jwt.ParseRSAPublicKeyFromPEM(sharedEncryptionKey)
			}
			return sharedEncryptionKey, nil
		})

		if err != nil {
			activityLog.Debug("Parse Failed: ", err)
			context.SetOutput(ovValid, false)
			return false, err
		}

		// if the token is invalid then return a false
		if token.Valid {
			activityLog.Debug(token.Claims, token.Header)
		} else {
			activityLog.Info("Token invalid: ", err)
			context.SetOutput(ovValid, false)
			context.SetOutput(ovToken, "")
			return false, nil
		}

		// Take the decoded claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			activityLog.Debug("Valid Token, claims are: ", claims)
			context.SetOutput(ovValid, true)
			return true, nil
		} else {
			activityLog.Info("Token invalid: ", err)
			context.SetOutput(ovValid, false)
		}
		context.SetOutput(ovToken, "")
		return true, nil

	case "Sign":
		{
			// Take the inputed header, payload and secret to create a new JWT
			activityLog.Debug("In Sign")

			var hdr map[string]interface{}
			claims := jwt.MapClaims{}

			// take the payload (claims) string and unmarshall it into a byte slice
			if err := json.Unmarshal([]byte(payload), &claims); err != nil {
				activityLog.Info("Invalid Payload: ", err)
				return false, err
			}
			activityLog.Debug("Unmarshalled JSON payload", claims)

			// Take the header string and unmarshall
			if err := json.Unmarshal([]byte(header), &hdr); err != nil {
				activityLog.Info("Invalid Header: ", err)
				return false, err
			}
			activityLog.Debug("Unmarshalled JSON header ", hdr)

			// get the alg value from the header
			alg := hdr["alg"].(string)

			// if the header and the passed algo method the same
			if algo != alg {
				activityLog.Info("Header algo doesn't match algorithm parm")
				return false, nil
			}

			// use the alg name to get the signing method
			signwith := jwt.GetSigningMethod(alg)
			activityLog.Debug("signing: ", signwith)

			// get the tokens object (this creates the first two parts of the token, based on the determined values, rather that using the passed strings)
			token := jwt.NewWithClaims(signwith, claims)
			activityLog.Debug("Token Object created", token)

			var key interface{}

			//  Depending on the algorithm type we need to convert  the format of the private string
			if isEs(alg) {
				key, err = jwt.ParseECPrivateKeyFromPEM(sharedEncryptionKey)
				if err != nil {
					activityLog.Info("Bad ECDSA key", err)
					return false, err
				}
			} else if isRs(alg) {
				key, err = jwt.ParseRSAPrivateKeyFromPEM(sharedEncryptionKey)
				if err != nil {
					activityLog.Info("Bad RSA key", err)
					return false, err
				}
			} else {
				key = sharedEncryptionKey
			}

			// Sign and get the complete encoded token as a string using the secret/key
			tokenString, err := token.SignedString(key)

			// if we don't have an error pass it back
			if err == nil {
				activityLog.Debug("Token String created", tokenString)
				context.SetOutput(ovToken, tokenString)
				return true, nil
			} else {
				activityLog.Info("Signing error: ", err)
				return false, err
			}

		}
	}

	return true, nil
}
func isEs(alg string) bool {
	return strings.HasPrefix(alg, "ES")
}

func isRs(alg string) bool {
	return strings.HasPrefix(alg, "RS")
}
