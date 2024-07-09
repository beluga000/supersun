package bank

type Recommand struct {

	// 기간
	Period string `json:"period" bson:"period"`
	// 목표금액
	TargetAmount int `json:"targetAmount" bson:"targetAmount"`
	// 주거래은행
	MainBank string `json:"mainBank" bson:"mainBank"`
	// 청년여부
	Youth string `json:"youth" bson:"youth"`
	// 월 납입 가능 금액
	MonthlyAmount int `json:"monthlyAmount" bson:"monthlyAmount"`
	// 사업자
	Business string `json:"business" bson:"business"`
	// 자녀
	Children string `json:"children" bson:"children"`
	// 취약계층
	VulnerableSocialGroup string `json:"vulnerableSocialGroup" bson:"vulnerableSocialGroup"`
	// 청년
	Young string `json:"young" bson:"young"`
	// 노인
	Old string `json:"old" bson:"old"`
	// 군인
	Soldier string `json:"soldier" bson:"soldier"`
}

type Recommand_Deposit struct {

	// 추천 적금 정보
	Deposit_Detail Deposit_Detail `json:"deposit_detail" bson:"deposit_detail"`
	// 납입 기간
	M_납입기간 int `json:"m_납입기간" bson:"m_납입기간"`
	// 월 납입금
	M_월납입금 int `json:"m_월납입금" bson:"m_월납입금"`
	// 납입 총 원금
	M_원금 int `json:"m_원금" bson:"m_원금"`
	// 납입 총 이자
	M_이자 int `json:"m_이자" bson:"m_이자"`
	// 납입 총 세금
	M_세금 int `json:"m_세금" bson:"m_세금"`
	// 만기금액
	M_만기금액 int `json:"m_만기금액" bson:"m_만기금액"`
}
