package co

import (
	"strings"
)

type MsgEx struct {

	// .
	Success bool

	//
	Failure bool

	Code         string
	Message      string
	OrgMessage__ string

	ValidateErr [][]string

	//
	RtnData interface{}

	ResultsCount *ResultsCount `json:"results_count" bson:"results_count"`

	// 에러 위치 ??  deplicated // message와 orgmessage를 비교해서 에러 위치를 파악하는게 좋을 듯
	Loc string
}

// .deplicated
func ErrorPass(messages ...string) MsgEx {

	// message_ := ""

	return MsgEx{

		Success: false,

		Failure: true,

		Code: "500",

		//OrgMessage: strings.Join(messages, " "),

		Message: strings.Join(messages, " "),

		//Loc: loc,
	}

}

// .
func SuccessPass(msg string) MsgEx {

	return MsgEx{

		Success: true,

		Failure: false,

		Code: "100",

		Message: msg,

		//OrgMessage: "",
	}

}

func SuccessPassMsg(msg string) MsgEx {

	return MsgEx{

		Success: true,

		Failure: false,

		Code: "100",

		Message: msg,

		//OrgMessage: "",
	}

}

func SuccessMsg(loc string, msg string) MsgEx {

	return MsgEx{

		Success: true,

		Failure: false,

		Code: "100",

		Message: msg,

		//OrgMessage: "",
		Loc: loc,
	}

}

func ErrorMsg(loc string, msg string) MsgEx {

	return MsgEx{

		Success: false,

		Failure: true,

		Code: "100",

		Message: msg,

		Loc: loc,
	}

}

// .
type ResultsCount struct {

	//
	Total int `json:"total" bson:"total"`

	Success int `json:"success" bson:"success"`

	Error int `json:"error" bson:"error"`
}
