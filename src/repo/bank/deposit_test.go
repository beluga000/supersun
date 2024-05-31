package bank_test

import (
	"testing"

	"sunny.ksw.kr/inits"
	"sunny.ksw.kr/repo/bank"
)

func TestNoticeSearch(t *testing.T) {

	// .
	mongo_uri := "mongodb://localhost:27017/?maxPoolSize=20&w=majority"
	mongo_db := "local"
	// mongo_uri := "mongodb://localhost:27017"
	// mongo_db := "test" //

	inits.MongoInitDebug(mongo_uri, mongo_db)

	// [
	// 	"DGB대구은행",
	// 	"SH수협은행",
	// 	"NH농협은행",
	// 	"전북은행",
	// 	"경남은행",
	// 	"제주은행",
	// 	"IBK기업은행",
	// 	"광주은행",
	// 	"SC제일은행",
	// 	"부산은행",
	// 	"신한은행",
	// 	"우리은행",
	// 	"케이뱅크",
	// 	"KDB산업은행",
	// 	"KB국민은행",
	// 	"하나은행",
	// 	"카카오뱅크",
	// 	"토스뱅크"
	//   ]

	DGB대구은행 := "https://financial.pstatic.net/pie/common-bi/1.1.0/images/IS_DGB_Profile.png"
	SH수협은행 := "https://financial.pstatic.net/pie/common-bi/1.1.0/images/BK_SH_Profile.png"
	NH농협은행 := "https://financial.pstatic.net/pie/common-bi/1.1.0/images/BK_NH_Profile.png"
	전북은행 := "https://financial.pstatic.net/pie/common-bi/1.1.0/images/BK_JEONBUK_Profile.png"
	경남은행 := "https://financial.pstatic.net/pie/common-bi/1.1.0/images/BK_KYOUNGNAM_Profile.png"
	제주은행 := "https://financial.pstatic.net/pie/common-bi/1.1.0/images/BK_JEJU_Profile.png"
	IBK기업은행 := "https://financial.pstatic.net/pie/common-bi/1.1.0/images/BK_IBK_Profile.png"
	광주은행 := "https://financial.pstatic.net/pie/common-bi/1.1.0/images/BK_KWANGJU_Profile.png"
	SC제일은행 := "https://financial.pstatic.net/pie/common-bi/1.1.0/images/BK_SC_Profile.png"
	KB국민은행 := "https://financial.pstatic.net/pie/common-bi/1.1.0/images/BK_KB_Profile.png"
	신한은행 := "https://financial.pstatic.net/pie/common-bi/1.1.0/images/BK_SHINHAN_Profile.png"
	하나은행 := "https://financial.pstatic.net/pie/common-bi/1.1.0/images/BK_HANA_Profile.png"
	우리은행 := "https://financial.pstatic.net/pie/common-bi/1.1.0/images/BK_WOORI_Profile.png"
	KDB산업은행 := "https://financial.pstatic.net/pie/common-bi/1.1.0/images/BK_KDB_Profile.png"
	카카오뱅크 := "https://financial.pstatic.net/pie/common-bi/1.1.0/images/BK_KAKAO_Profile.png"
	토스뱅크 := "https://financial.pstatic.net/pie/common-bi/1.1.0/images/BK_TOSS_Profile.png"
	케이뱅크 := "https://financial.pstatic.net/pie/common-bi/1.1.0/images/BK_K_Profile.png"
	부산은행 := "https://financial.pstatic.net/pie/common-bi/1.1.0/images/BK_BUSAN_Profile.png"

	search := bank.SearchInstalment_Savings{}

	search.Finds()

	for _, v := range search.Instalment_Savingss {

		switch v.CompanyName {
		case "DGB대구은행":
			v.CompanyLogoURL = DGB대구은행
			v.Update()
		case "SH수협은행":
			v.CompanyLogoURL = SH수협은행
			v.Update()
		case "NH농협은행":
			v.CompanyLogoURL = NH농협은행
			v.Update()
		case "전북은행":
			v.CompanyLogoURL = 전북은행
			v.Update()
		case "경남은행":
			v.CompanyLogoURL = 경남은행
			v.Update()
		case "제주은행":
			v.CompanyLogoURL = 제주은행
			v.Update()
		case "IBK기업은행":
			v.CompanyLogoURL = IBK기업은행
			v.Update()
		case "광주은행":
			v.CompanyLogoURL = 광주은행
			v.Update()
		case "SC제일은행":
			v.CompanyLogoURL = SC제일은행
			v.Update()
		case "KB국민은행":
			v.CompanyLogoURL = KB국민은행
			v.Update()
		case "신한은행":
			v.CompanyLogoURL = 신한은행
			v.Update()
		case "하나은행":
			v.CompanyLogoURL = 하나은행
			v.Update()
		case "우리은행":
			v.CompanyLogoURL = 우리은행
			v.Update()
		case "KDB산업은행":
			v.CompanyLogoURL = KDB산업은행
			v.Update()
		case "카카오뱅크":
			v.CompanyLogoURL = 카카오뱅크
			v.Update()
		case "토스뱅크":
			v.CompanyLogoURL = 토스뱅크
			v.Update()
		case "케이뱅크":
			v.CompanyLogoURL = 케이뱅크
			v.Update()
		case "부산은행":
			v.CompanyLogoURL = 부산은행
			v.Update()
		}

	}

}
