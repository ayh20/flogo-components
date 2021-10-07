package jwt

import (
	"fmt"
	"testing"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/support/test"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {

	ref := activity.GetRef(&JWT{})
	act := activity.Get(ref)

	assert.NotNil(t, act)
}

func TestDecrypt(t *testing.T) {

	act := &JWT{}
	tc := test.NewActivityContext(act.Metadata())

	//fmt.Println("#######   Testing JWT Decrypt")

	/* 	//test1
	   	tc.SetInput("token", `eyJhbGciOiJkaXIiLCJlbmMiOiJBMTI4R0NNIn0..jg45D9nmr6-8awml.z-zglLlEw9MVkYHi-Znd9bSwc-oRGbqKzf9WjXqZxno.kqji2DiZHZmh-1bLF6ARPw`)
	   	tc.SetInput("secret", "itsa16bytesecret")
	   	tc.SetInput("mode", "Decrypt")
	   	tc.SetInput("algorithm", "algorithm")
	   	act.Eval(tc)

	   	if tc.GetOutput("result") == nil {
	   		t.Fail()
	   	} */

	//test2
	fmt.Println("===> Test2")
	tc.SetInput("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ")
	tc.SetInput("secret", "secret")
	tc.SetInput("mode", "Verify")
	tc.SetInput("algorithm", "HS256")
	act.Eval(tc)

	if tc.GetOutput("valid") == nil {
		fmt.Println("******** Test Failed  ********")
		t.Fail()
	} else {
		fmt.Println("******** Result: ", tc.GetOutput("valid"), tc.GetOutput("claims"))
	}

	//test2a
	fmt.Println("===> Test2A")
	tc.SetInput("token", "xeyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ")
	tc.SetInput("secret", "secret")
	tc.SetInput("mode", "Verify")
	tc.SetInput("algorithm", "HS256")
	act.Eval(tc)

	if tc.GetOutput("valid") == nil {
		fmt.Println("******** Test Failed  ********")
		t.Fail()
	} else {
		fmt.Println("******** Result: ", tc.GetOutput("valid"), tc.GetOutput("claims"))
	}

	//test2b
	fmt.Println("===> Test2B")
	tc.SetInput("token", "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsIng1dCI6Imk2bEdrM0ZaenhSY1ViMkMzbkVRN3N5SEpsWSIsImtpZCI6Imk2bEdrM0ZaenhSY1ViMkMzbkVRN3N5SEpsWSJ9.eyJhdWQiOiJlZjFkYTlkNC1mZjc3LTRjM2UtYTAwNS04NDBjM2Y4MzA3NDUiLCJpc3MiOiJodHRwczovL3N0cy53aW5kb3dzLm5ldC9mYTE1ZDY5Mi1lOWM3LTQ0NjAtYTc0My0yOWYyOTUyMjIyOS8iLCJpYXQiOjE1MzcyMzMxMDYsIm5iZiI6MTUzNzIzMzEwNiwiZXhwIjoxNTM3MjM3MDA2LCJhY3IiOiIxIiwiYWlvIjoiQVhRQWkvOElBQUFBRm0rRS9RVEcrZ0ZuVnhMaldkdzhLKzYxQUdyU091TU1GNmViYU1qN1hPM0libUQzZkdtck95RCtOdlp5R24yVmFUL2tES1h3NE1JaHJnR1ZxNkJuOHdMWG9UMUxrSVorRnpRVmtKUFBMUU9WNEtjWHFTbENWUERTL0RpQ0RnRTIyMlRJbU12V05hRU1hVU9Uc0lHdlRRPT0iLCJhbXIiOlsid2lhIl0sImFwcGlkIjoiNzVkYmU3N2YtMTBhMy00ZTU5LTg1ZmQtOGMxMjc1NDRmMTdjIiwiYXBwaWRhY3IiOiIwIiwiZW1haWwiOiJBYmVMaUBtaWNyb3NvZnQuY29tIiwiZmFtaWx5X25hbWUiOiJMaW5jb2xuIiwiZ2l2ZW5fbmFtZSI6IkFiZSAoTVNGVCkiLCJpZHAiOiJodHRwczovL3N0cy53aW5kb3dzLm5ldC83MmY5ODhiZi04NmYxLTQxYWYtOTFhYi0yZDdjZDAxMjIyNDcvIiwiaXBhZGRyIjoiMjIyLjIyMi4yMjIuMjIiLCJuYW1lIjoiYWJlbGkiLCJvaWQiOiIwMjIyM2I2Yi1hYTFkLTQyZDQtOWVjMC0xYjJiYjkxOTQ0MzgiLCJyaCI6IkkiLCJzY3AiOiJ1c2VyX2ltcGVyc29uYXRpb24iLCJzdWIiOiJsM19yb0lTUVUyMjJiVUxTOXlpMmswWHBxcE9pTXo1SDNaQUNvMUdlWEEiLCJ0aWQiOiJmYTE1ZDY5Mi1lOWM3LTQ0NjAtYTc0My0yOWYyOTU2ZmQ0MjkiLCJ1bmlxdWVfbmFtZSI6ImFiZWxpQG1pY3Jvc29mdC5jb20iLCJ1dGkiOiJGVnNHeFlYSTMwLVR1aWt1dVVvRkFBIiwidmVyIjoiMS4wIn0.D3H6pMUtQnoJAGq6AHd")
	tc.SetInput("secret", "secret")
	tc.SetInput("mode", "Verify")
	tc.SetInput("algorithm", "RS256")
	act.Eval(tc)

	if tc.GetOutput("valid") == nil {
		fmt.Println("******** Test Failed  ********")
		t.Fail()
	} else {
		fmt.Println("******** Result: ", tc.GetOutput("valid"), tc.GetOutput("claims"))
	}

	//test3
	fmt.Println("===> Test3")
	tc.SetInput("header", `{"typ":"JWT","alg":"HS256"}`)
	tc.SetInput("payload", `{"foo":"bar","nbf":1444478400}`)
	tc.SetInput("secret", "secret")
	tc.SetInput("mode", "Sign")
	tc.SetInput("algorithm", "HS256")
	act.Eval(tc)

	if tc.GetOutput("token") == nil {
		fmt.Println("******** Test Failed  ********")
		t.Fail()
	} else {
		fmt.Println("******** Result: ", tc.GetOutput("token"))
	}

	//Test 4 - validate returned token
	fmt.Println("===> Test4")
	var lasttoken = tc.GetOutput("token").(string)
	tc.SetInput("token", lasttoken)
	tc.SetInput("secret", "secret")
	tc.SetInput("mode", "Verify")
	tc.SetInput("algorithm", "HS256")
	act.Eval(tc)

	if tc.GetOutput("valid") == nil {
		fmt.Println("******** Test Failed  ********")
		t.Fail()
	} else {
		fmt.Println("******** Result: ", tc.GetOutput("valid"), tc.GetOutput("claims"))
	}

	fmt.Println("===> Test5")
	tc.SetInput("token", lasttoken+"x")
	tc.SetInput("secret", "secret")
	tc.SetInput("mode", "Verify")
	tc.SetInput("algorithm", "HS256")
	act.Eval(tc)

	if tc.GetOutput("token") == nil {
		fmt.Println("******** Test Failed  ********")
		t.Fail()
	}

	//test6
	fmt.Println("===> Test6")
	tc.SetInput("header", `{"typ":"JWT","alg":"RS256"}`)
	tc.SetInput("algorithm", "RS256")
	tc.SetInput("payload", `{"foo":"bar","nbf":1444478400}`)
	tc.SetInput("secret", `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEA4f5wg5l2hKsTeNem/V41fGnJm6gOdrj8ym3rFkEU/wT8RDtn
SgFEZOQpHEgQ7JL38xUfU0Y3g6aYw9QT0hJ7mCpz9Er5qLaMXJwZxzHzAahlfA0i
cqabvJOMvQtzD6uQv6wPEyZtDTWiQi9AXwBpHssPnpYGIn20ZZuNlX2BrClciHhC
PUIIZOQn/MmqTD31jSyjoQoV7MhhMTATKJx2XrHhR+1DcKJzQBSTAGnpYVaqpsAR
ap+nwRipr3nUTuxyGohBTSmjJ2usSeQXHI3bODIRe1AuTyHceAbewn8b462yEWKA
Rdpd9AjQW5SIVPfdsz5B6GlYQ5LdYKtznTuy7wIDAQABAoIBAQCwia1k7+2oZ2d3
n6agCAbqIE1QXfCmh41ZqJHbOY3oRQG3X1wpcGH4Gk+O+zDVTV2JszdcOt7E5dAy
MaomETAhRxB7hlIOnEN7WKm+dGNrKRvV0wDU5ReFMRHg31/Lnu8c+5BvGjZX+ky9
POIhFFYJqwCRlopGSUIxmVj5rSgtzk3iWOQXr+ah1bjEXvlxDOWkHN6YfpV5ThdE
KdBIPGEVqa63r9n2h+qazKrtiRqJqGnOrHzOECYbRFYhexsNFz7YT02xdfSHn7gM
IvabDDP/Qp0PjE1jdouiMaFHYnLBbgvlnZW9yuVf/rpXTUq/njxIXMmvmEyyvSDn
FcFikB8pAoGBAPF77hK4m3/rdGT7X8a/gwvZ2R121aBcdPwEaUhvj/36dx596zvY
mEOjrWfZhF083/nYWE2kVquj2wjs+otCLfifEEgXcVPTnEOPO9Zg3uNSL0nNQghj
FuD3iGLTUBCtM66oTe0jLSslHe8gLGEQqyMzHOzYxNqibxcOZIe8Qt0NAoGBAO+U
I5+XWjWEgDmvyC3TrOSf/KCGjtu0TSv30ipv27bDLMrpvPmD/5lpptTFwcxvVhCs
2b+chCjlghFSWFbBULBrfci2FtliClOVMYrlNBdUSJhf3aYSG2Doe6Bgt1n2CpNn
/iu37Y3NfemZBJA7hNl4dYe+f+uzM87cdQ214+jrAoGAXA0XxX8ll2+ToOLJsaNT
OvNB9h9Uc5qK5X5w+7G7O998BN2PC/MWp8H+2fVqpXgNENpNXttkRm1hk1dych86
EunfdPuqsX+as44oCyJGFHVBnWpm33eWQw9YqANRI+pCJzP08I5WK3osnPiwshd+
hR54yjgfYhBFNI7B95PmEQkCgYBzFSz7h1+s34Ycr8SvxsOBWxymG5zaCsUbPsL0
4aCgLScCHb9J+E86aVbbVFdglYa5Id7DPTL61ixhl7WZjujspeXZGSbmq0Kcnckb
mDgqkLECiOJW2NHP/j0McAkDLL4tysF8TLDO8gvuvzNC+WQ6drO2ThrypLVZQ+ry
eBIPmwKBgEZxhqa0gVvHQG/7Od69KWj4eJP28kq13RhKay8JOoN0vPmspXJo1HY3
CKuHRG+AP579dncdUnOMvfXOtkdM4vk0+hWASBQzM9xzVcztCa+koAugjVaLS9A+
9uQoqEeVNTckxx0S2bYevRy7hGQmUJTyQm3j1zEUR5jpdbL83Fbq
-----END RSA PRIVATE KEY-----`)
	tc.SetInput("mode", "Sign")
	act.Eval(tc)

	if tc.GetOutput("token") == nil {
		fmt.Println("******** Test Failed  ********")
		t.Fail()
	} else {
		fmt.Println("******** Result: ", tc.GetOutput("token"))
	}

	//Test 7 - validate returned token
	fmt.Println("===> Test7")
	lasttoken = tc.GetOutput("token").(string)
	tc.SetInput("token", lasttoken)
	//tc.SetInput("secret", "secret")
	tc.SetInput("header", "")
	tc.SetInput("algorithm", "RS256")
	tc.SetInput("payload", "")
	tc.SetInput("mode", "Verify")
	tc.SetInput("algorithm", "RS256")
	tc.SetInput("secret", `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA4f5wg5l2hKsTeNem/V41
fGnJm6gOdrj8ym3rFkEU/wT8RDtnSgFEZOQpHEgQ7JL38xUfU0Y3g6aYw9QT0hJ7
mCpz9Er5qLaMXJwZxzHzAahlfA0icqabvJOMvQtzD6uQv6wPEyZtDTWiQi9AXwBp
HssPnpYGIn20ZZuNlX2BrClciHhCPUIIZOQn/MmqTD31jSyjoQoV7MhhMTATKJx2
XrHhR+1DcKJzQBSTAGnpYVaqpsARap+nwRipr3nUTuxyGohBTSmjJ2usSeQXHI3b
ODIRe1AuTyHceAbewn8b462yEWKARdpd9AjQW5SIVPfdsz5B6GlYQ5LdYKtznTuy
7wIDAQAB
-----END PUBLIC KEY-----`)
	act.Eval(tc)

	if tc.GetOutput("valid") == nil {
		fmt.Println("******** Test Failed  ********")
		t.Fail()
	} else {
		fmt.Println("******** Result: ", tc.GetOutput("valid"), tc.GetOutput("claims"))
	}

	//Test 8 - validate token
	fmt.Println("===> Test8")
	lasttoken = tc.GetOutput("token").(string)
	tc.SetInput("token", "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsIng1dCI6IllNRUxIVDBndmIwbXhvU0RvWWZvbWpxZmpZVSIsImtpZCI6IllNRUxIVDBndmIwbXhvU0RvWWZvbWpxZmpZVSJ9.eyJhdWQiOiJodHRwOi8vZmxvZ28udGVzdDEiLCJpc3MiOiJodHRwczovL3N0cy53aW5kb3dzLm5ldC80N2VjMGM4Yy0yMDdkLTQ3NzQtOGNjOS02ZDNmOGRhMzE3OGUvIiwiaWF0IjoxNTg1MTQ5MTE1LCJuYmYiOjE1ODUxNDkxMTUsImV4cCI6MTU4NTE1MzAxNSwiYWlvIjoiNDJkZ1lIZ1NGaXlwdTFVL3lyZDZYV2YvdFZvNUFBPT0iLCJhcHBpZCI6IjNhNjRjNGQ1LTQ1M2ItNGU1Ni05ZGIyLTJmNGFiMzk2OWQ5ZSIsImFwcGlkYWNyIjoiMSIsImlkcCI6Imh0dHBzOi8vc3RzLndpbmRvd3MubmV0LzQ3ZWMwYzhjLTIwN2QtNDc3NC04Y2M5LTZkM2Y4ZGEzMTc4ZS8iLCJvaWQiOiIzNGNkNGQ4Ni1kMmM0LTRkMjQtOWU2My0yY2ZjMmJkMDU4OWYiLCJzdWIiOiIzNGNkNGQ4Ni1kMmM0LTRkMjQtOWU2My0yY2ZjMmJkMDU4OWYiLCJ0aWQiOiI0N2VjMGM4Yy0yMDdkLTQ3NzQtOGNjOS02ZDNmOGRhMzE3OGUiLCJ1dGkiOiJZSExUekVzbW5VYUlyc1c1R1htTEFBIiwidmVyIjoiMS4wIn0.ZAPvx8urJ1vWbUh86lNfzM4I9G2AVIVF522ev2m_ZkLyxo_m6SrXtAHAyxI3U72xt1Q0Z8MyjUSSXnLG5SuyWPzftSUzAY0wRCHLl-Yl9HZPsp258qaW0I5j0UM70oqeEhFnK5Y_gc9HC8IoZnDs71kuIM5vlceKuSYRWhb-XCBUlteUlJ9rO1GxyrCoruxppU-zdqo-c4e8qbQW0zqOA6iavdYyoflWOhaAKRTlOPbb-MpBL_Zt5WvKNi13JCD09ps3mdokCM1dyEo72rUefoSlDZbA4JIaveVz2e2Iq0PL6FcdEaxNkXvr_mIEl0Rj7DElCSdYYa3m0z64cSJVIQ")
	//tc.SetInput("secret", "secret")
	tc.SetInput("mode", "Decode")
	tc.SetInput("algorithm", "")
	tc.SetInput("secret", "")
	act.Eval(tc)

	if tc.GetOutput("valid") == nil {
		fmt.Println("******** Test Failed  ********")
		t.Fail()
	} else {
		fmt.Println("******** Result: ", tc.GetOutput("valid"), tc.GetOutput("claims"))
	}
}
