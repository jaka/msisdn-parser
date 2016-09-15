package msisdn

import "testing"

func TestParseMSISDN_withLetters(t *testing.T) {

  _, ok := ParseMSISDN("01L465")
  if ok {
    t.Fail()
  }
}

func TestParseMSISDN_withLeadingPlusButWithoutLegalCountry(t *testing.T) {

  _, ok := ParseMSISDN("+0")
  if ok {
    t.Fail()
  }
}

func TestParseMSISDN_withLeadingPlusAndLegalCountry(t *testing.T) {

  answer, ok := ParseMSISDN("+30")
  if ok || (answer.CountryISO != "GR") || (answer.CountryDial != "30") {
    t.Fail()
  }
}

func TestParseMSISDN_withoutLeadingPlusAndLegalCountry(t *testing.T) {

  answer, ok := ParseMSISDN("7")
  if ok || (answer.CountryISO != "RU") || (answer.CountryDial != "7") {
    t.Fail()
  }
}

func TestParseMSISDN_withoutLeadingPlusAndLegalCountryAndIllegalProvider(t *testing.T) {

  answer, ok := ParseMSISDN("790")
  if ok || (answer.CountryISO != "RU") || (answer.CountryDial != "7") ||
     (answer.ProviderName != "") {
    t.Fail()
  }
  answer, ok = ParseMSISDN("38500000")
  if ok || (answer.CountryISO != "HR") || (answer.CountryDial != "385") ||
    (answer.ProviderName != "") {
    t.Fail()
  }
}

func TestParseMSISDN_LegalCountryAndProvider(t *testing.T) {

  answer, ok := ParseMSISDN("3863")
  if ok || (answer.CountryISO != "SI") || (answer.CountryDial != "386") ||
     (answer.ProviderName != "Telekom") {
    t.Fail()
  }
  answer, ok = ParseMSISDN("+38631")
  if ok || (answer.CountryISO != "SI") || (answer.CountryDial != "386") ||
     (answer.ProviderName != "Mobitel") {
    t.Fail()
  }
}

func TestParseMSISDN_LegalCountryAndProviderButIllegalNumber(t *testing.T) {

  answer, ok := ParseMSISDN("3867123456")
  if ok || (answer.CountryISO != "SI") || (answer.CountryDial != "386") ||
     (answer.ProviderName != "Mobitel") {
    t.Fail()
  }
  answer, ok = ParseMSISDN("38676201234567")
  if ok || (answer.CountryISO != "SI") || (answer.CountryDial != "386") ||
     (answer.ProviderName != "Amis") {
    t.Fail()
  }
}

func TestParseMSISDN_LegalCountryAndProviderAndLegalNumber(t *testing.T) {

  answer, ok := ParseMSISDN("38671234567")
  if !ok || (answer.CountryISO != "SI") || (answer.CountryDial != "386") ||
     (answer.ProviderName != "Mobitel") || (answer.SubscriberNumber != "234567") {
    t.Fail()
  }
  answer, ok = ParseMSISDN("3867620123")
  if !ok || (answer.CountryISO != "SI") || (answer.CountryDial != "386") ||
     (answer.ProviderName != "Amis") || (answer.SubscriberNumber != "123") {
    t.Fail()
  }
  answer, ok = ParseMSISDN("38659012345")
  if !ok || (answer.CountryISO != "SI") || (answer.CountryDial != "386") ||
     (answer.ProviderName != "T-2") || (answer.SubscriberNumber != "12345") {
    t.Fail()
  }
}

