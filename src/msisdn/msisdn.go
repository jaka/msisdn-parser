package msisdn

import "strings"
import "regexp"

type Provider struct {
  code string
  name string
  regex string
}

type Country struct {
  country map[uint8]Country
  iso string
  providers []Provider
}
type CountryMap map[uint8]Country

type Answer struct {
  countryDial string
  countryISO string
  providerName string
  subscriberNumber string
}

const ZERO = 48
const NINE = ZERO + 9

var GreeceProviders = []Provider{}
var CroatiaProviders = []Provider{}

var SloveniaProviders = []Provider{
  {"3", "Telekom", "^(\\d{7})$"},
  {"30", "Simobil", "^(\\d{6})$"},
  {"31", "Mobitel", "^(\\d{6})$"},
  {"4", "Telekom", "^(\\d{7})$"},
  {"40", "Simobil", "^(\\d{6})$"},
  {"41", "Mobitel", "^(\\d{6})$"},
  {"49", "Mobitel", "^(\\d{6})$"},
  {"51", "Mobitel", "^(\\d{6})$"},
  {"591", "Telemach", "^(\\d{5})$"},
  {"590", "T-2", "^(\\d{5})$"},
  {"599", "Telekom", "^(\\d{5})$"},
  {"64", "T-2", "^(\\d{6})$"},
  {"7", "Telekom", "^(\\d{7})$"},
  {"70", "Tušmobil", "^(\\d{6})$"},
  {"71", "Mobitel", "^(\\d{6})$"},
  {"7620", "Amis", "^(\\d{3})$"},
}

var countries = CountryMap{
  3: { country: CountryMap{
    0: { nil, "GR", GreeceProviders },
    8: { country: CountryMap{
      5: { nil, "HR", CroatiaProviders },
      6: { nil, "SI", SloveniaProviders },
    }},
  }},
  7: { nil, "RU", nil },
}

func getCountry(v string) (Country, int) {

  var max int
  countryList := countries
  if max = 3; len(v) < 3 {
    max = len(v)
  }

  for i := 0; i < max; i++ {
    digit := v[i] - ZERO
    country, ok := countryList[digit]
    if ok && (country.country == nil) {
      return country, (i + 1)
    }
    countryList = country.country
  }

  return Country{}, 0
}

func getProvider (providers []Provider, v string) (Provider, int) {

  var provider Provider
  var numberLength, providerCodeLength int

  numberLength = len(v)
  providerCodeLength = 0

  for _, curProvider := range providers {
    curProviderCodeLength := len(curProvider.code)
    if numberLength < curProviderCodeLength {
      continue
    }
    if (v[:curProviderCodeLength] == curProvider.code) && 
       (curProviderCodeLength > providerCodeLength) {
      provider = curProvider
      providerCodeLength = curProviderCodeLength
    }
  }

  return provider, providerCodeLength
}

func getSubscriberNumber(regex string, v string) (string, int) {

  matched, _ := regexp.MatchString(regex, v)
  if matched {
    return v, len(v)
  }
  return v, 0
}

func checkNumber(v string) (string, int) {

  o := strings.TrimRight(v, " ")
  o = strings.TrimLeft(o, " +")
  for i := 0; i < len(o); i++ {
    if o[i] < ZERO || o[i] > NINE {
      return "", 0
    }
  }
  return o, len(o)
}

func ParseMSISDN(msisdn string) (Answer, bool) {

  var answer Answer
 
  number, rv := checkNumber(msisdn)
  if rv == 0 {
    return answer, false
  }

  country, countryCodeLength := getCountry(number)
  if countryCodeLength == 0 {
    return answer, false
  }
  answer.countryDial = number[:countryCodeLength]
  answer.countryISO = country.iso

  provider, providerCodeLength := getProvider(country.providers, number[countryCodeLength:])
  if providerCodeLength == 0 {
    return answer, false
  }
  answer.providerName = provider.name

  subscriberNumber, subscriberNumberLength := getSubscriberNumber(provider.regex, number[countryCodeLength + providerCodeLength:])
  if subscriberNumberLength == 0 {
    return answer, false
  }
  answer.subscriberNumber = subscriberNumber

  return answer, true
}
