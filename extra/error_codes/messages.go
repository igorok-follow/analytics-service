package error_codes

const (
	INTERNAL            string = "internal server error"
	BADREQUEST          string = "bad request"
	TOKENINCORRECT      string = "token incorrect"
	TOKENLIFETIME       string = "token lifetime expired"
	USERISEXIST         string = "user is exist"
	USERNOTREGISTERED   string = "user not registered"
	INVALIDPASSWD       string = "invalid password"
	TOKENCREATIONERROR  string = "token creation error"
	INVALIDREFRESHTOKEN string = "invalid refresh_token"
	INVALIDCODE         string = "invalid verification code"
)
