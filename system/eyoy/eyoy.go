package eyoy

const DEFAULT = 600

const FAILED_GET_USER = 700
const FAILED_CREATE_USER = 701
const FAILED_UPDATE_USESR = 702
const FAILED_DELETE_USESR = 703

var errorTitle = map[int]string{
	DEFAULT: "Default",

	FAILED_GET_USER:     "Failed Get User",
	FAILED_CREATE_USER:  "Failed Create User",
	FAILED_UPDATE_USESR: "Failed Update User",
	FAILED_DELETE_USESR: "Failed Delete User",
}

func StatusTitle(code int) string {
	title, ok := errorTitle[code]
	if !ok {
		return "Unknown Error"
	}

	return title
}
