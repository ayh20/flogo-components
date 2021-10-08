package jwt

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/project-flogo/core/activity"
)

// JWT is an Activity that works with JWT Tokens
// It can create, verify and decrypt JWT tokens
// inputs : {input1, input2, datatype, comparemode}
// outputs: result (bool)
type Activity struct {
	metadata *activity.Metadata
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

	context.Logger().Debug("In Eval")

	in := &Input{}
	output := &Output{}
	err = context.GetInputObject(in)
	if err != nil {
		return false, err
	}
	// Get the runtime values
	sharedEncryptionKey := []byte(in.Secret)

	context.Logger().Debug(in.Mode, in.Header, in.Payload, in.Token, in.Secret, in.Algorithm)

	// Determine code path based on mode
	switch in.Mode {
	case "Decode":
		output.Claims = ""
		output.Token = ""

		// Parse takes the token string and a function for looking up the key. The latter is especially
		// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
		// head of the token to identify which key to use, but the parsed token (head and claims) is provided
		// to the callback, providing flexibility.

		token, err := jwt.Parse(in.Token, func(token *jwt.Token) (interface{}, error) {
			// sharedEncryptionKey contains a plain secret or a public key
			// we used the passed Algo name to determine the handling of the secret
			// returning the formatted ES/RS key or secret string
			return sharedEncryptionKey, nil
		})

		context.Logger().Info("Created Token")

		if err != nil {
			output.Claims = fmt.Errorf("Token Error: %v", err).Error()
			context.Logger().Info("Parse Failed - Token Error: ", err)
			return true, nil
		}

		context.Logger().Info("Check for Valid Token")
		// if the token is invalid then return a false
		if token.Valid {
			context.Logger().Info(token.Claims, token.Header)
		} else {
			output.Claims = fmt.Errorf("Token Invalid: %v", err).Error()
			context.Logger().Info("Token invalid: ", err)
			return true, nil
		}

		context.Logger().Info("Decode Claims")
		// Take the decoded claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			claimsJSON, _ := json.Marshal(claims)
			output.Claims = string(claimsJSON)
			context.Logger().Info("Valid Token, claims are: ", string(claimsJSON))
			output.Valid = true
			return true, nil
		} else {
			context.Logger().Info("Claims invalid: ", err)
			return true, nil
		}

	case "Verify":

		context.Logger().Info("In Verify - V0.1.0")

		// Set default responses
		output.Valid = false
		output.Claims = ""
		output.Token = ""

		// Parse takes the token string and a function for looking up the key. The latter is especially
		// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
		// head of the token to identify which key to use, but the parsed token (head and claims) is provided
		// to the callback, providing flexibility.

		token, err := jwt.Parse(in.Token, func(token *jwt.Token) (interface{}, error) {
			// sharedEncryptionKey contains a plain secret or a public key
			// we used the passed Algo name to determine the handling of the secret
			// returning the formatted ES/RS key or secret string
			if err != nil {
				output.Claims = fmt.Errorf("Parse Error: %v", err).Error()
				context.Logger().Info("Parse Failed - Parse Error: ", err)
				context.SetOutputObject(output)
				return true, nil
			}
			if isEs(in.Algorithm) {
				return jwt.ParseECPublicKeyFromPEM(sharedEncryptionKey)
			} else if isRs(in.Algorithm) {
				return jwt.ParseRSAPublicKeyFromPEM(sharedEncryptionKey)
			}
			return sharedEncryptionKey, nil
		})

		context.Logger().Info("Created Token")

		if err != nil {
			output.Claims = fmt.Errorf("Token Error: %v", err).Error()
			context.Logger().Info("Parse Failed - Token Error: ", err)
			context.SetOutputObject(output)
			return true, nil
		}

		context.Logger().Info("Check for Valid Token")
		// if the token is invalid then return a false
		if token.Valid {
			context.Logger().Info(token.Claims, token.Header)
		} else {
			output.Claims = fmt.Errorf("Token Invalid: %v", err).Error()
			context.Logger().Info("Token invalid: ", err)
			context.SetOutputObject(output)
			return true, nil
		}

		context.Logger().Info("Decode Claims")
		// Take the decoded claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			claimsJSON, _ := json.Marshal(claims)
			output.Claims = string(claimsJSON)
			context.Logger().Info("Valid Token, claims are: ", string(claimsJSON))
			output.Valid = true
			context.SetOutputObject(output)
			return true, nil
		} else {
			context.Logger().Info("Claims invalid: ", err)
			context.SetOutputObject(output)
			return true, nil
		}

		//return true, nil

	case "Sign":
		{
			// Take the inputed header, payload and secret to create a new JWT
			context.Logger().Debug("In Sign")

			var hdr map[string]interface{}
			claims := jwt.MapClaims{}

			// take the payload (claims) string and unmarshall it into a byte slice
			if err := json.Unmarshal([]byte(in.Payload), &claims); err != nil {
				context.Logger().Info("Invalid Payload: ", err)
				return false, err
			}
			context.Logger().Debug("Unmarshalled JSON payload", claims)

			// Take the header string and unmarshall
			if err := json.Unmarshal([]byte(in.Header), &hdr); err != nil {
				context.Logger().Info("Invalid Header: ", err)
				return false, err
			}
			context.Logger().Debug("Unmarshalled JSON header ", hdr)

			// get the alg value from the header
			alg := hdr["alg"].(string)

			// if the header and the passed algo method the same
			if in.Algorithm != alg {
				context.Logger().Info("Header algo doesn't match algorithm parm")
				return false, nil
			}

			// use the alg name to get the signing method
			signwith := jwt.GetSigningMethod(alg)
			context.Logger().Debug("signing: ", signwith)

			// get the tokens object (this creates the first two parts of the token, based on the determined values, rather that using the passed strings)
			token := jwt.NewWithClaims(signwith, claims)
			context.Logger().Debug("Token Object created", token)

			var key interface{}

			//  Depending on the algorithm type we need to convert  the format of the private string
			if isEs(alg) {
				key, err = jwt.ParseECPrivateKeyFromPEM(sharedEncryptionKey)
				if err != nil {
					context.Logger().Info("Bad ECDSA key", err)
					return false, err
				}
			} else if isRs(alg) {
				key, err = jwt.ParseRSAPrivateKeyFromPEM(sharedEncryptionKey)
				if err != nil {
					context.Logger().Info("Bad RSA key", err)
					return false, err
				}
			} else {
				key = sharedEncryptionKey
			}

			// Sign and get the complete encoded token as a string using the secret/key
			tokenString, err := token.SignedString(key)

			// if we don't have an error pass it back
			if err == nil {
				context.Logger().Debug("Token String created", tokenString)
				output.Token = tokenString
				context.SetOutputObject(output)
				return true, nil
			} else {
				context.Logger().Info("Signing error: ", err)
				return false, err
			}

		}
	}
	err = context.SetOutputObject(output)
	if err != nil {
		return false, err
	}

	return true, nil
}
func isEs(alg string) bool {
	return strings.HasPrefix(alg, "ES")
}

func isRs(alg string) bool {
	return strings.HasPrefix(alg, "RS")
}
