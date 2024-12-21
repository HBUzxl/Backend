package models

type Case struct {
	Id uint `json:"id" gorm:"unique; AUTO_INCREMENT"` // ID

	CaseID        string `json:"caseID" gorm:"type:varchar(255); unique"` //病例号
	PatientName   string `json:"patientName"`                             //病人姓名
	PatientGender string `json:"patientGender"`                           //病人性别
	PatientAge    int    `json:"patientAge"`                              //病人年龄
	PatientPhone  string `json:"patientPhone"`                            //病人电话
	PatientType   string `json:"patientType"`                             //病人类型
	BiopsySite    string `json:"biopsySite"`                              //取材部位
	TissueCount   int    `json:"tissueCount"`                             //组织数量
	BarCode       string `json:"barCode"`                                 //条形码

	CheckupNo     string `json:"checkupNo"`     //体检编号
	ClinicalPhone string `json:"clinicalPhone"` //临床电话
	Hospital      string `json:"hospital"`      //送检医院
	SampleDate    string `json:"sampleDate"`    //采样日期
	ReceiveDate   string `json:"receiveDate"`   //接收日期

	PathologyType         string `json:"pathologyType"`         //病理类型
	InpatientNo           string `json:"inpatientNo"`           //门诊、住院号
	BedNo                 string `json:"bedNo"`                 //床号
	MaritalStatus         string `json:"maritalStatus"`         //婚姻状况
	PatientAddress        string `json:"patientAddress"`        //病人地址
	ClinicalDiagnosis     string `json:"clinicalDiagnosis"`     //临床诊断
	ClinicalData          string `json:"clinicalData"`          //临床资料
	GrossFinding          string `json:"grossFinding"`          //肉眼所见
	Immunohistochemistry  string `json:"immunohistochemistry"`  //免疫组化
	PathologicalDiagnosis string `json:"pathologicalDiagnosis"` //病理诊断
	Remarks               string `json:"remarks"`               //备注

	PrintCount int `json:"printCount"` //打印次数

	//病例状态
	CaseStatus string `json:"caseStatus"` //病例状态

	//关联专家
	ExpertID uint   `json:"expertID" gorm:"index"`                            //专家ID
	Expert   Expert `json:"expert" gorm:"foreignKey:ExpertID; references:Id"` //专家

	//关联切片
	Slices []Slice `json:"slices" gorm:"foreignKey:CaseID; references:CaseID"` //切片

	//关联附件
	Attachments []Attachment `json:"attachments" gorm:"foreignKey:CaseID; references:CaseID"` //附件
}
