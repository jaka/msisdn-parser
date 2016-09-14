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
  if ok || (answer.countryISO != "GR") || (answer.countryDial != "30") {
    t.Fail()
  }
}

func TestParseMSISDN_withoutLeadingPlusAndLegalCountry(t *testing.T) {

  answer, ok := ParseMSISDN("7")
  if ok || (answer.countryISO != "RU") || (answer.countryDial != "7") {
    t.Fail()
  }
}

func TestParseMSISDN_withoutLeadingPlusAndLegalCountryAndIllegalProvider(t *testing.T) {

  answer, ok := ParseMSISDN("790")
  if ok || (answer.countryISO != "RU") || (answer.countryDial != "7") {
    t.Fail()
  }
  answer, ok = ParseMSISDN("38500000")
  if ok || (answer.countryISO != "HR") || (answer.countryDial != "385") {
    t.Fail()
  }
}

func TestParseMSISDN_LegalCountryAndProvider(t *testing.T) {

  answer, ok := ParseMSISDN("3863")
  if ok || (answer.countryISO != "SI") || (answer.countryDial != "386") ||
     (answer.providerName != "Telekom") {
    t.Fail()
  }
  answer, ok = ParseMSISDN("+38631")
  if ok || (answer.countryISO != "SI") || (answer.countryDial != "386") ||
     (answer.providerName != "Mobitel") {
    t.Fail()
  }
}

func TestParseMSISDN_LegalCountryAndProviderButIllegalNumber(t *testing.T) {

  answer, ok := ParseMSISDN("3867123456")
  if ok || (answer.countryISO != "SI") || (answer.countryDial != "386") || 
     (answer.providerName != "Mobitel") {
    t.Fail()
  }
  answer, ok = ParseMSISDN("38676201234567")
  if ok || (answer.countryISO != "SI") || (answer.countryDial != "386") ||
     (answer.providerName != "Amis") {
    t.Fail()
  }
}

func TestParseMSISDN_LegalCountryAndProviderAndLegalNumber(t *testing.T) {

  answer, ok := ParseMSISDN("38671234567")
  if !ok || (answer.countryISO != "SI") || (answer.countryDial != "386") ||
     (answer.providerName != "Mobitel") || (answer.subscriberNumber != "234567") {
    t.Fail()
  }
  answer, ok = ParseMSISDN("3867620123")
  if !ok || (answer.countryISO != "SI") || (answer.countryDial != "386") || 
     (answer.providerName != "Amis") || (answer.subscriberNumber != "123") {
    t.Fail()
  }
  answer, ok = ParseMSISDN("38659012345")
  if !ok || (answer.countryISO != "SI") || (answer.countryDial != "386") || 
     (answer.providerName != "T-2") || (answer.subscriberNumber != "12345") {
    t.Fail()
  }
}



