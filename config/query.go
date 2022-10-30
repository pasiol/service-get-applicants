package config

import (
	"strings"

	"github.com/beevik/etree"
	pq "github.com/pasiol/gopq"
)

// Applicant struct
type Applicant struct {
	ID                         string
	Email                      string
	LastName                   string
	NickName                   string
	FirtsName                  string
	PersonalID                 string
	School                     string
	Class                      string
	Exam                       string
	StartDate                  string
	LastRegistrationDate       string
	Memo                       string
	Campus                     string
	ContactPersonAddress       string
	ContactPerson              string
	ContactPersonPostalAddress string
	Couch                      string
	Education                  string
	CostCentre                 string
	Code5                      string
	Gender                     string
	Nationally                 string
	NativelLanguage            string
	HomeCountry                string
	StreetAddress              string
	PostalAddress              string
	CellPhone                  string
	DataShieldStartDate        string
	DataShieldEndDate          string
	WorkExperience             string
	PrimaryEducation           string
	FormerEducation            string
	FormerEducationYear        string
	Remark3                    string
	Remark1                    string
	HealthRequirements         string
	Ban                        string
	BanningSchool              string
	CompulsoryEducationYear    string
	Code10                     string
	ApplyDate                  string
	ApplicantID                string
	ApplyEducation             string
	HomeCounty                 string
	Password                   string
	TeachingLanguage           string
}

type UserAccount struct {
	Email      string
	FirstName  string
	NickName   string
	LastName   string
	Password   string
	CellPhone  string
	PersonalID string
}

func mapApplicantData(data []string) Applicant {
	a := Applicant{}

	a.ID = data[0]
	a.Email = data[1]
	a.LastName = capWords(data[2])
	a.NickName = capWords(data[3])
	a.FirtsName = capWords(data[4])
	a.PersonalID = data[5]
	a.School = data[6]
	a.Class = data[7]
	a.Exam = data[8]
	a.StartDate = data[9]
	a.LastRegistrationDate = data[10]
	a.Memo = data[11]
	a.Campus = data[12]
	a.ContactPersonAddress = data[13]
	a.ContactPerson = data[14]
	a.ContactPersonPostalAddress = data[15]
	a.Couch = data[16]
	a.Education = data[17]
	a.CostCentre = data[18]
	a.Code5 = data[19]
	a.Gender = data[20]
	a.Nationally = data[21]
	a.NativelLanguage = data[22]
	a.HomeCountry = data[23]
	a.StreetAddress = capWords(data[24])
	a.PostalAddress = strings.ToTitle(data[25])
	a.CellPhone = data[26]
	a.DataShieldStartDate = data[27]
	a.DataShieldEndDate = data[28]
	a.WorkExperience = data[29]
	a.PrimaryEducation = data[30]
	a.FormerEducation = data[31]
	a.FormerEducationYear = data[32]
	a.Remark3 = data[33]
	a.Remark1 = data[34]
	a.HealthRequirements = data[35]
	a.Ban = data[36]
	a.BanningSchool = data[37]
	a.CompulsoryEducationYear = data[38]
	a.Code10 = data[39]
	a.ApplyDate = data[40]
	a.ApplicantID = data[41]
	a.ApplyEducation = data[42]
	a.HomeCounty = data[43]
	a.Password = data[44]
	a.TeachingLanguage = data[45]

	return a
}

func mapUserAccountData(data []string) UserAccount {
	u := UserAccount{}
	u.Email = data[1]
	u.LastName = capWords(data[2])
	u.FirstName = capWords(data[4])
	u.NickName = capWords(data[3])
	u.Password = data[44]
	u.CellPhone = data[26]
	u.PersonalID = data[5]

	return u
}

func capWords(words string) string {
	splitted := strings.Split(words, " ")
	var capitalized string
	for _, word := range splitted {
		if len(word) > 1 {
			capitalized = capitalized + strings.ToUpper(word[0:1]) + word[1:] + " "
		} else {
			capitalized = capitalized + strings.ToUpper(word) + " "
		}
	}
	return capitalized[:len(words)]
}

// Applicants query
func Applicants() pq.PrimusQuery {
	pq := pq.PrimusQuery{}
	pq.Charset = "UTF-8"
	pq.Database = "hakijat"
	pq.Sort = ""
	pq.Search = ""
	pq.Header = "id;email;sukunimi;kutsumanimi;etunimet;hetu;koulu;koulutus;tutkinto;aloituspäivä;viimeinen ilmoittautumispäivä;yhteyshenkilön_lisätietoja;yksikkö;yhteyshenkilön_lähiosoite;Yhteyshenkilön_nimi;yhteyshenkilön_postiosoite;ryhmävalmentaja;koulutusala;kustannuspaikka_alue;koodi_5;sukupuoli;kansalaisuus;äidinkieli;kotimaa;katuosoite;postinumero;matkapuhelin;turvakielto_aloitus;turvakielto_lopetus;työkokemus;pohjakoulutus;aikaisempi_koulutus;aikaisempi_koulutus_vuosi;huom_3;huom_1;terveydentilavaatimus;este;oppilaitos_peruuttanut;oppivelvollisuus_vuosi;koodi_10;hakuajankohta;hakijan_korttinumero;haettu_koulutus;kotikunta;salasana;opetuskieli"
	pq.Data = "#DATA{V1};#DATA{K33};#DATA{K2};#DATA{K4};#DATA{K3};#DATA{K5};Riverian hakijat;#DATA{K20064^V1};#DATA{K20064^K32^V1};#DATA{K200};#DATA{K12107+14};#DATA{K904, CHROK(\" 1234567890qwertyuiopåasdfghjklöäzxcvbnm,.()-:/\\\",1,1)};#DATA{K20064^K33^K2};#DATA{K901, , CHROK(\" 1234567890qwertyuiopåasdfghjklöäzxcvbnm,.()-:/\\\",1,1)};#DATA{K900, CHROK(\" 1234567890qwertyuiopåasdfghjklöäzxcvbnm,.()-:/\\\",1,1)};#DATA{K902, CHROK(\" 1234567890qwertyuiopåasdfghjklöäzxcvbnm,.()-:/\\\",1,1)};#DATA{K152^V1};#DATA{K20064^K25^V1};#DATA{K20064^K557[1]^V1};#DATA{K434};#DATA{K6};#DATA{K50^V1};#DATA{K43^V1};#DATA{K52^V1};#DATA{K20, CHROK(\" 1234567890qwertyuiopåasdfghjklöäzxcvbnm,.()-:/\\\",1,1)};#DATA{K21};#DATA{K23, CHROK(\" 1234567890+\",1,1)};#DATA{K1740};#DATA{K1741};#DATA{K614, CHROK(\" 1234567890qwertyuiopåasdfghjklöäzxcvbnm,.()-:/\\\",1,1)};#DATA{K210^V1};#DATA{K218, CHROK(\" 1234567890qwertyuiopåasdfghjklöäzxcvbnm,.()-:/\\\",1,1)};#DATA{K3448};#DATA{K507, CHROK(\" 1234567890qwertyuiopåasdfghjklöäzxcvbnm,.()-:/\\\",1,1)};#DATA{K505, CHROK(\" 1234567890qwertyuiopåasdfghjklöäzxcvbnm,.()-:/\\\",1,1)};#DATA{K1317};#DATA{K3270};#DATA{K342};#DATA{K131};#DATA{K2859, CHROK(\" 1234567890qwertyuiopåasdfghjklöäzxcvbnm,.()-:/\\\",1,1)};#DATA{K967};#DATA{V1};#DATA{K11964^V1};#DATA{K60^V1};#DATA{K125};#DATA{K44^V1}"
	pq.Footer = ""

	return pq
}

// ArchieveXML generator
func ArchieveXML(id string) (string, error) {
	archieveDoc := etree.NewDocument()
	archieveDoc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)
	primusquery := archieveDoc.CreateElement("PRIMUSQUERY_IMPORT")
	primusquery.CreateAttr("ARCHIVEMODE", "0")
	primusquery.CreateAttr("CREATEIFNOTFOUND", "0")
	identity := archieveDoc.CreateElement("IDENTITY")
	identity.CreateText("service-get-applicants")
	card := archieveDoc.CreateElement("CARD")
	card.CreateAttr("FIND", id)
	archieve := card.CreateElement("ARKISTO")
	archieve.CreateText("Kyllä")
	archieveDoc.Indent(2)
	xmlAsString, _ := archieveDoc.WriteToString()
	filename, err := pq.CreateTMPFile(pq.StringWithCharset(128)+".xml", xmlAsString)
	if err != nil {
		return "", err
	}
	return filename, nil
}

// ApplicantXML func
func ApplicantXML(data []string) (string, error) {
	applicant := mapApplicantData(data)
	applicantDoc := etree.NewDocument()
	applicantDoc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)
	primusquery := applicantDoc.CreateElement("PRIMUSQUERY_IMPORT")
	primusquery.CreateAttr("ARCHIVEMODE", "0")
	primusquery.CreateAttr("CREATEIFNOTFOUND", "1")
	identity := primusquery.CreateElement("IDENTITY")
	identity.CreateText("service-get-applicants")
	card := primusquery.CreateElement("CARD")
	card.CreateAttr("FIND", applicant.ID)
	noLDAP := card.CreateElement("EILDAP")
	noLDAP.CreateText("Kyllä")
	userName := card.CreateElement("UUSITUNNUS")
	userName.CreateAttr("CMD", "MODIFY")
	userName.CreateAttr("LINE", "1")
	userName.CreateText(applicant.Email)
	applicantID := card.CreateElement("HAKIJAN_KORTTINUMERO")
	applicantID.CreateText(applicant.ID)
	email := card.CreateElement("EMAIL")
	email.CreateText(applicant.Email)
	lastName := card.CreateElement("SUKUNIMI")
	lastName.CreateText(applicant.LastName)
	nickName := card.CreateElement("KUTSUMANIMI")
	nickName.CreateText(applicant.NickName)
	firstName := card.CreateElement("ETUNIMET")
	firstName.CreateText(applicant.FirtsName)
	personalID := card.CreateElement("HETU")
	personalID.CreateText(applicant.PersonalID)
	school := card.CreateElement("KOULU")
	school.CreateAttr("CMD", "MODIFY")
	school.CreateAttr("LINE", "1")
	school.CreateText(applicant.School)
	class := card.CreateElement("KOULUTUS")
	class.CreateText(applicant.Class)
	exam := card.CreateElement("TUTKINTO")
	exam.CreateText(applicant.Exam)
	startDate := card.CreateElement("ALOITUSPÄIVÄ")
	startDate.CreateText(applicant.StartDate)
	memo := card.CreateElement("YHTEYSHENKILÖN_LISÄTIETOJA")
	memo.CreateText(applicant.Memo)
	campus := card.CreateElement("YKSIKKÖ")
	campus.CreateText(applicant.Campus)
	contactPersonAddress := card.CreateElement("YHTEYSHENKILÖN_LÄHIOSOITE")
	contactPersonAddress.CreateText(applicant.ContactPersonAddress)
	contactPersonPostalAddress := card.CreateElement("YHTEYSHENKILÖN_POSTIOSOITE")
	contactPersonPostalAddress.CreateText(applicant.ContactPersonPostalAddress)
	couch := card.CreateElement("OMAVALMENTAJA")
	couch.CreateText(applicant.Couch)
	education := card.CreateElement("KOULUTUSALA")
	education.CreateText(applicant.Education)
	costCentre := card.CreateElement("KUSTANNUSPAIKKA_ALUE")
	costCentre.CreateAttr("CMD", "MODIFY")
	costCentre.CreateAttr("LINE", "1")
	costCentre.CreateText(applicant.CostCentre)
	code5 := card.CreateElement("KOODI_5")
	code5.CreateText(applicant.Code5)
	gender := card.CreateElement("SUKUPUOLI")
	if applicant.Gender == "Mies" {
		gender.CreateText("1")
	} else {
		gender.CreateText("2")
	}
	nationally := card.CreateElement("KANSALAISUUS")
	nationally.CreateText(applicant.Nationally)
	nativeLanguage := card.CreateElement("ÄIDINKIELI")
	nativeLanguage.CreateText(applicant.NativelLanguage)
	homeCountry := card.CreateElement("KOTIMAA")
	homeCountry.CreateText(applicant.HomeCountry)
	streetAddress := card.CreateElement("KATUOSOITE")
	streetAddress.CreateText(applicant.StreetAddress)
	postaldAddress := card.CreateElement("POSTIOSOITE")
	postaldAddress.CreateText(applicant.PostalAddress)
	cellPhone := card.CreateElement("MATKAPUHELIN")
	cellPhone.CreateText(applicant.CellPhone)
	dataShieldStartDate := card.CreateElement("TURVAKIELTO_ALOITUS")
	dataShieldStartDate.CreateText(applicant.DataShieldStartDate)
	dataShieldEndDate := card.CreateElement("TURVAKIELTO_LOPETUS")
	dataShieldEndDate.CreateText(applicant.DataShieldEndDate)
	workExperience := card.CreateElement("TYÖKOKEMUS")
	workExperience.CreateText(applicant.WorkExperience)
	primaryEducation := card.CreateElement("POHJAKOULUTUS")
	primaryEducation.CreateText(applicant.PrimaryEducation)
	formerEducation := card.CreateElement("AIKAISEMPI_KOULUTUS")
	formerEducation.CreateAttr("CMD", "MODIFY")
	formerEducation.CreateAttr("LINE", "1")
	formerEducation.CreateText(applicant.FormerEducation)
	formerEducationYear := card.CreateElement("AIKAISEMPI_KOULUTUS_VUOSI")
	formerEducationYear.CreateAttr("CMD", "MODIFY")
	formerEducationYear.CreateAttr("LINE", "1")
	formerEducationYear.CreateText(applicant.FormerEducationYear)
	remark3 := card.CreateElement("HUOM_3")
	remark3.CreateText(applicant.Remark3)
	remark1 := card.CreateElement("HUOM_1")
	remark1.CreateText(applicant.Remark1)
	healthRequirements := card.CreateElement("TERVEYDENTILAVAATIMUS")
	healthRequirements.CreateText(applicant.HealthRequirements)
	ban := card.CreateElement("ESTE")
	ban.CreateText(applicant.Ban)
	banningSchool := card.CreateElement("OPPILAITOS_PERUUTTANUT")
	banningSchool.CreateText(applicant.BanningSchool)
	compulsoryEducationYear := card.CreateElement("OPPIVELVOLLISUUS_VUOSI")
	compulsoryEducationYear.CreateText(applicant.CompulsoryEducationYear)
	code10 := card.CreateElement("KOODI_10")
	code10.CreateText(applicant.Code10)
	applyDate := card.CreateElement("HAKUAJANKOHTA")
	applyDate.CreateText(applicant.ApplyDate)
	applyEducation := card.CreateElement("HAETTU_KOULUTUS")
	applyEducation.CreateText(applicant.ApplyEducation)
	homeCounty := card.CreateElement("KOTIKUNTA")
	homeCounty.CreateText(applicant.HomeCounty)
	studentType := card.CreateElement("OPISKELIJALAJI")
	studentType.CreateText("45")
	newTypeUserName := card.CreateElement("UUDENTYYPPINENTUNNUS")
	newTypeUserName.CreateText("Kyllä")
	teachingLanguage := card.CreateElement("OPETUSKIELI")
	teachingLanguage.CreateText(applicant.TeachingLanguage)

	applicantDoc.Indent(2)
	xmlAsString, _ := applicantDoc.WriteToString()
	filename, err := pq.CreateTMPFile(pq.StringWithCharset(128)+".xml", xmlAsString)
	if err != nil {
		return "", err
	}
	return filename, nil
}

func UserAccountXML(data []string) (string, error) {
	user := mapUserAccountData(data)
	emailLowerCase := strings.ToLower(user.Email)
	applicantDoc := etree.NewDocument()
	applicantDoc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)
	primusquery := applicantDoc.CreateElement("PRIMUSQUERY_IMPORT")
	primusquery.CreateAttr("ARCHIVEMODE", "1")
	primusquery.CreateAttr("CREATEIFNOTFOUND", "1")
	identity := primusquery.CreateElement("IDENTITY")
	identity.CreateText("service-create-wilma-accounts")
	card := primusquery.CreateElement("CARD")
	card.CreateAttr("FIND", user.Email)
	email := card.CreateElement("EMAIL")
	email.CreateText(emailLowerCase)
	userAccount := card.CreateElement("TUNNUS")
	userAccount.CreateText(user.Email)
	lastName := card.CreateElement("SUKUNIMI")
	lastName.CreateText(user.LastName)
	firstName := card.CreateElement("ETUNIMET")
	firstName.CreateText(user.FirstName)
	nickName := card.CreateElement("KUTSUMANIMI")
	nickName.CreateText(user.NickName)
	password := card.CreateElement("SALASANA")
	password.CreateText(user.Password)
	password2 := card.CreateElement("SALASANA2")
	password2.CreateText(user.Password)
	cellPhone := card.CreateElement("MATKAPUHELIN")
	cellPhone.CreateText(user.CellPhone)
	//personalID := card.CreateElement("HENKILÖTUNNUS")
	//personalID.CreateText(user.PersonalID)

	applicantDoc.Indent(2)
	xmlAsString, _ := applicantDoc.WriteToString()
	//log.Printf("XML: \n%s", xmlAsString)
	filename, err := pq.CreateTMPFile(pq.StringWithCharset(128)+".xml", xmlAsString)
	if err != nil {
		return "", err
	}
	return filename, nil
}

// Accounts query
func Accounts() pq.PrimusQuery {
	pq := pq.PrimusQuery{}
	pq.Charset = "UTF-8"
	pq.Database = "wpasswd"
	pq.Sort = ""
	pq.Search = "(K1=%B0% OR K1=%B1%)"
	pq.Header = "id;tunnus;arkistoitu"
	pq.Data = "#DATA{V1};#DATA{K1010};#DATA{K1}"
	pq.Footer = ""

	return pq
}

// Students query
func NewTypeUserAccount(username string) pq.PrimusQuery {
	pq := pq.PrimusQuery{}
	pq.Charset = "UTF-8"
	pq.Database = "opphenk"
	pq.Sort = ""
	pq.Search = "K1=%B0% AND (K1043[1]^K1010=\"" + username + "\"  AND K1040=%B1% AND K1048=%B1%)"
	pq.Data = "#DATA{K1043[1]^K1010}"
	pq.Footer = ""

	return pq
}

// Accounts query
func StudentRegistryApplicantID(id string) pq.PrimusQuery {
	pq := pq.PrimusQuery{}
	pq.Charset = "UTF-8"
	pq.Database = "opphenk"
	pq.Sort = ""
	pq.Search = "(K1705=" + id + ")"
	pq.Data = "#DATA{K1705}"
	pq.Footer = ""

	return pq
}

// Archived query
func Archived(id string) pq.PrimusQuery {
	pq := pq.PrimusQuery{}
	pq.Charset = "UTF-8"
	pq.Database = "hakijat"
	pq.Sort = ""
	pq.Search = "K1=%B0% AND (V1=\"" + id + "\")"
	pq.Data = "#DATA{V1}"
	pq.Footer = ""

	return pq
}

// Moved query
func Moved(id string) pq.PrimusQuery {
	pq := pq.PrimusQuery{}
	pq.Charset = "UTF-8"
	pq.Database = "opphenk"
	pq.Sort = ""
	pq.Search = "(K1705=\"" + id + "\")"
	pq.Data = "#DATA{V1}"
	pq.Footer = ""

	return pq
}
