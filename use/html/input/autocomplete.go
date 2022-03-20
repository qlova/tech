package input

/*
	Autocomplete attribute lets web developers specify what if any
	permission the user agent has to provide automated assistance
	in filling out form field values, as well as guidance to the
	browser as to the type of information expected in the field.

	It is available on <input> elements that take a text or
	numeric value as input, <textarea> elements, <select>
	elements, and <form> elements.

	The source of the suggested values is generally up to the
	browser; typically values come from past values entered by
	the user, but they may also come from pre-configured values.
	For instance, a browser might let the user save their name,
	address, phone number, and email addresses for autocomplete
	purposes. Perhaps the browser offers the ability to save
	encrypted credit card information, for autocompletion
	following an authentication procedure.

	If an <input>, <select> or <textarea> element has no
	autocomplete attribute, then browsers use the autocomplete
	attribute of the element's form owner, which is either the
	<form> element that the element is a descendant of, or the
	<form> whose id is specified by the form attribute of the
	element.
*/
type AutoComplete []AutoCompletion

// AutoCompletion values for AutoComplete.
type AutoCompletion string

const (
	/*
		The browser is not permitted to automatically enter or select
		a value for this field. It is possible that the document or
		application provides its own autocomplete feature, or that
		security concerns require that the field's value not be
		automatically entered.
	*/
	AutoCompleteOff AutoCompletion = "off"

	/*
		The browser is allowed to automatically complete the input.
		No guidance is provided as to the type of data expected
		in the field, so the browser may use its own judgement.
	*/
	AutoCompleteOn AutoCompletion = "on"

	/*
		The field expects the value to be a person's full name.
		Using "name" rather than breaking the name down into its
		components is generally preferred because it avoids
		dealing with the wide diversity of human names and how
		they are structured;
	*/
	AutoCompleteName AutoCompletion = "name"

	// The prefix or title, such as "Mrs.", "Mr.", "Miss",
	// "Ms.", "Dr.", or "Mlle.".
	AutoCompleteHonorificPrefix AutoCompletion = "honorific-prefix"

	// The given (or "first") name.
	AutoCompleteGivenName AutoCompletion = "given-name"

	// The middle name.
	AutoCompleteAdditionalName AutoCompletion = "additional-name"

	// The family (or "last") name.
	AutoCompleteFamilyName AutoCompletion = "family-name"

	// The suffix, such as "Jr.", "B.Sc.", "PhD.", "MBASW", or "IV".
	AutoCompleteHonorificSuffix AutoCompletion = "honorific-suffix"

	// A nickname or handle.
	AutoCompleteNickname AutoCompletion = "nickname"

	// An email address
	AutoCompleteEmail AutoCompletion = "email"

	// A username or account name.
	AutoCompleteUsername AutoCompletion = "username"

	/*
		A new password. When creating a new account or changing
		passwords, this should be used for an "Enter your new password"
		or "Confirm new password" field, as opposed to a general
		"Enter your current password" field that might be present.
		This may be used by the browser both to avoid accidentally
		filling in an existing password and to offer assistance in
		creating a secure password.
	*/
	AutoCompleteNewPassword AutoCompletion = "new-password"

	// The user's current password.
	AutoCompleteCurrentPassword AutoCompletion = "current-password"

	// A one-time code used for verifying user identity.
	AutoCompleteOneTimeCode AutoCompletion = "one-time-code"

	// A job title, or the title a person has within an
	// organization, such as "Senior Technical Writer",
	// "President", or "Assistant Troop Leader".
	AutoCompleteOrganisationTitle AutoCompletion = "organization-title"

	// A company or organization name, such as "Acme Widget Company"
	// or "Girl Scouts of America".
	AutoCompleteOrganisation AutoCompletion = "organization"

	// A street address. This can be multiple lines of text, and
	// should fully identify the location of the address within
	// its second administrative level (typically a city or town),
	// but should not include the city name, ZIP or postal code, or country name.
	AutoCompleteStreetAddress AutoCompletion = "street-address"

	// Each individual line of the street address. These should
	// only be present if the "street-address" is not present.
	AutoCompleteAddressLine1 AutoCompletion = "address-line1"
	AutoCompleteAddressLine2 AutoCompletion = "address-line2"
	AutoCompleteAddressLine3 AutoCompletion = "address-line3"

	// The finest-grained administrative level, in addresses
	// which have four levels.
	AutoCompleteAddressLevel4 AutoCompletion = "address-level4"

	// The third administrative level, in addresses with at
	// least three administrative levels.
	AutoCompleteAddressLevel3 AutoCompletion = "address-level3"

	// The second administrative level, in addresses with at
	// least two of them. In countries with two administrative
	// levels, this would typically be the city, town, village,
	// or other locality in which the address is located.
	AutoCompleteAddressLevel2 AutoCompletion = "address-level2"

	// The first administrative level in the address. This is
	// typically the province in which the address is located.
	// In the United States, this would be the state. In
	// Switzerland, the canton. In the United Kingdom, the
	// post town.
	AutoCompleteAddressLevel1 AutoCompletion = "address-level1"

	// A country or territory code.
	AutoCompleteCountry AutoCompletion = "country"

	// A postal code (in the United States, this is the ZIP code).
	AutoCompletePostalCode AutoCompletion = "postal-code"

	// The full name as printed on or associated with a payment
	// instrument such as a credit card. Using a full name field
	// is preferred, typically, over breaking the name into pieces.
	AutoCompleteCreditCardName AutoCompletion = "cc-name"

	// A given (first) name as given on a payment instrument
	// like a credit card.
	AutoCompleteCreditCardGivenName AutoCompletion = "cc-given-name"

	// A middle name as given on a payment instrument or credit card.
	AutoCompleteCreditCardAdditionalName AutoCompletion = "cc-additional-name"

	// A family name, as given on a credit card.
	AutoCompleteCreditCardFamilyName AutoCompletion = "cc-family-name"

	// A credit card number or other number identifying a
	// payment method, such as an account number.
	AutoCompleteCreditCardNumber AutoCompletion = "cc-number"

	// A payment method expiration date, typically in the form "MM/YY" or "MM/YYYY".
	AutoCompleteCreditCardExpirationDate AutoCompletion = "cc-exp"

	// The month in which the payment method expires.
	AutoCompleteCreditCardExpirationMonth AutoCompletion = "cc-exp-month"

	// The year in which the payment method expires.
	AutoCompleteCreditCardExpirationYear AutoCompletion = "cc-exp-year"

	// The security code for the payment instrument; on credit cards,
	// this is the 3-digit verification number on the back of the card
	AutoCompleteCreditCardSecurityCode AutoCompletion = "cc-csc"

	// The type of payment instrument (such as "Visa" or "Master Card").
	AutoCompleteCreditCardType AutoCompletion = "cc-type"

	// The currency in which the transaction is to take place.
	AutoCompleteTransactionCurrency AutoCompletion = "transaction-currency"

	// The amount, given in the currency specified by
	// "transaction-currency", of the transaction, for a payment form.
	AutoCompleteTransactionAmount AutoCompletion = "transaction-amount"

	// A preferred language, given as a valid BCP 47 language tag.
	AutoCompleteLanguage AutoCompletion = "language"

	// A birth date, as a full date.
	AutoCompleteBirthDate AutoCompletion = "bday"

	// The day of the month of a birth date.
	AutoCompleteBirthDateDay AutoCompletion = "bday-day"

	// The month of a birth date.
	AutoCompleteBirthDateMonth AutoCompletion = "bday-month"

	// The year of a birth date.
	AutoCompleteBirthDateYear AutoCompletion = "bday-year"

	// A gender identity (such as "Female", "Fa'afafine",
	// "Male"), as freeform text without newlines.
	AutoCompleteSex AutoCompletion = "sex"

	// A full telephone number, including the country code.
	AutoCompleteTelephoneNumber AutoCompletion = "tel"

	// The country code, such as "1" for the United States,
	// Canada, and other areas in North America and parts of the Caribbean.
	AutoCompleteTelephoneNumberCountryCode AutoCompletion = "tel-country-code"

	// The entire phone number without the country code component,
	// including a country-internal prefix. For the phone number
	// "1-855-555-6502", this field's value would be "855-555-6502".
	AutoCompleteTelephoneNumberNational AutoCompletion = "tel-national"

	// The area code of a phone number, without the country code
	// or the country-internal prefix if appropriate.
	AutoCompleteTelephoneNumberAreaCode AutoCompletion = "tel-area-code"

	// The local number of a phone number, without the country code
	// or the country-internal prefix if appropriate.
	AutoCompleteTelephoneNumberLocal AutoCompletion = "tel-local"

	// A telephone extension code within the phone number, such as a
	// room or suite number in a hotel or an office extension in a company.
	AutoCompleteTelephoneNumberExtension AutoCompletion = "tel-extension"

	// A URL for an instant messaging protocol endpoint, such as
	// "xmpp:username@example.net".
	AutoCompleteInstantMessagingProtocolEndpoint AutoCompletion = "impp"

	// A URL, such as a home page or company web site address as
	// appropriate given the context of the other fields in the form.
	AutoCompleteURL AutoCompletion = "url"

	// The URL of an image representing the person, company, or
	// contact information given in the other fields in the form.
	AutoCompletePhoto AutoCompletion = "photo"
)
