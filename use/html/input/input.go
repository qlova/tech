/*
	Package input provides the HTML <input> element.

	The <input> HTML element is used to create interactive
	controls for web-based forms in order to accept data
	from the user; a wide variety of types of input data
	and control widgets are available, depending on the
	device and user agent. The <input> element is one of
	the most powerful and complex in all of HTML due to
	the sheer number of combinations of input types and
	attributes.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package input

import (
	"fmt"
	"strings"

	"qlova.tech/use/html"
	"qlova.tech/use/html/link"
	"qlova.tech/web/tree"
)

// Tag is the HTML tag of this element.
const Tag = html.Tag("input")

// New returns a <label> html tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag, html.Void)...)
}

const (
	// A push button with no default behavior displaying
	// the value of the value attribute, empty by default.
	TypeButton = html.Attribute("type=button")

	// A check box allowing single values to be
	// selected/deselected.
	TypeCheckbox = html.Attribute("type=checkbox")

	// A control for specifying a color; opening a
	// color picker when active in supporting browsers.
	TypeColor = html.Attribute("type=color")

	// A control for entering a date (year, month, and
	// day, with no time). Opens a date picker or numeric
	// wheels for year, month, day when active in
	// supporting browsers.
	TypeDate = html.Attribute("type=date")

	// A control for entering a date and time, with no
	// time zone. Opens a date picker or numeric wheels
	// for date- and time-components when active in supporting browsers.
	TypeDateTimeLocal = html.Attribute("type=datetime-local")

	// A field for editing an email address. Looks like
	// a text input, but has validation parameters and
	// relevant keyboard in supporting browsers and
	// devices with dynamic keyboards.
	TypeEmail = html.Attribute("type=email")

	// A control that lets the user select a file. Use
	// the accept attribute to define the types of files
	// that the control can select.
	TypeFile = html.Attribute("type=file")

	// A control that is not displayed but whose value
	// is submitted to the server.
	TypeHidden = html.Attribute("type=hidden")

	// A graphical submit button. Displays an image
	// defined by the src attribute. The alt attribute
	// displays if the image src is missing.
	TypeImage = html.Attribute("type=image")

	// A control for entering a month and year,
	// with no time zone.
	TypeMonth = html.Attribute("type=month")

	// A control for entering a number. Displays a
	// spinner and adds default validation when
	// supported. Displays a numeric keypad in
	// some devices with dynamic keypads.
	TypeNumber = html.Attribute("type=number")

	// A single-line text field whose value is
	// obscured. Will alert user if site is not secure.
	TypePassword = html.Attribute("type=password")

	// A radio button, allowing a single value to be
	// selected out of multiple choices with the same
	// name value.
	TypeRadio = html.Attribute("type=radio")

	// A control for entering a number whose exact
	// value is not important. Displays as a range
	// widget defaulting to the middle value. Used
	// in conjunction min and max to define the
	// range of acceptable values.
	TypeRange = html.Attribute("type=range")

	// A button that resets the contents of the
	// form to default values. Not recommended.
	TypeReset = html.Attribute("type=reset")

	// A single-line text field for entering search
	// strings. Line-breaks are automatically removed
	// from the input value. May include a delete icon
	// in supporting browsers that can be used to
	// clear the field. Displays a search icon instead
	// of enter key on some devices with dynamic keypads.
	TypeSearch = html.Attribute("type=search")

	// A button that submits the form.
	TypeSubmit = html.Attribute("type=submit")

	// A control for entering a telephone number.
	// Displays a telephone keypad in some devices
	// with dynamic keypads.
	TypeTelephoneNumber = html.Attribute("type=tel")

	// The default value. A single-line text field.
	// Line-breaks are automatically removed from
	// the input value.
	TypeText = html.Attribute("type=text")

	// A control for entering a time value with no time zone.
	TypeTime = html.Attribute("type=time")

	// A field for entering a URL. Looks like a text
	// input, but has validation parameters and relevant
	// keyboard in supporting browsers and devices with
	// dynamic keyboards.
	TypeURL = html.Attribute("type=url")

	// A control for entering a date consisting of
	// a week-year number and a week number with no time zone.
	TypeWeek = html.Attribute("type=week")
)

// Accept defines the file types that TypeFile should accept.
// This string is a comma-separated list of unique file type
// specifiers. Because a given file type may be identified
// in more than one manner, it's useful to provide a thorough
// set of type specifiers when you need files of a given format.
func Accept(extensions []string) html.Attribute {
	return html.Attr("accept", strings.Join(extensions, ","))
}

// Alt is valid for the TypeImage only, the alt attribute provides
// alternative text for the image, displaying the value of the
// attribute if the image src is missing or otherwise fails to load.
func Alt(description string) html.Attribute {
	return html.Attr("alt", description)
}

const (
	/*
		AutoFocus if present, indicates that the input should
		automatically have focus when the page has finished
		loading (or when the <dialog> containing the element
		has been displayed).

		No more than one element in the document may have the
		autofocus attribute. If put on more than one element,
		the first one with the attribute receives focus.

		The autofocus attribute cannot be used on inputs of
		type hidden, since hidden inputs cannot be focused.

		Warning: Automatically focusing a form control can
		confuse visually-impaired people using screen-reading
		technology and people with cognitive impairments. When
		autofocus is assigned, screen-readers "teleport" their
		user to the form control without warning them beforehand.

		Use careful consideration for accessibility when applying
		the autofocus attribute. Automatically focusing on a
		control can cause the page to scroll on load. The focus
		can also cause dynamic keyboards to display on some touch
		devices. While a screen reader will announce the label of
		the form control receiving focus, the screen reader will
		not announce anything before the label, and the sighted
		user on a small device will equally miss the context
		created by the preceding content.
	*/
	AutoFocus = html.Attribute("autofocus")
)

/*
	The capture attribute specifies that, optionally, a new file
	should be captured, and which device should be used to capture
	that new media of a type defined by the accept attribute.
*/
const (

	// The user-facing camera and/or microphone should be used.
	CaptureUser = html.Attribute("capture=user")

	// The outward-facing camera and/or microphone should be used
	CaptureEnvironment = html.Attribute("capture=environment")
)

// Valid for both radio and checkbox types, checked is a Boolean
// attribute. If present on a radio type, it indicates that the
// radio button is the currently selected one in the group of
// same-named radio buttons. If present on a checkbox type, it
// indicates that the checkbox is checked by default (when the
// page loads). It does not indicate whether this checkbox is
// currently checked: if the checkbox's state is changed, this
// content attribute does not reflect the change. (Only the
// HTMLInputElement's checked IDL attribute is updated.)
const Checked = html.Attribute("checked")

// Valid for text and search input types only, the dirname attribute
// enables the submission of the directionality of the element. When
// included, the form control will submit with two name/value pairs:
// the first being the name and value, the second being the value of
// the dirname as the name with the value of ltr or rtl being set by
// the browser.
func Dirname(name string) html.Attribute {
	return html.Attr("dirname", name)
}

// A Boolean attribute which, if present, indicates that the user
// should not be able to interact with the input. Disabled inputs
// are typically rendered with a dimmer color or using some other
// form of indication that the field is not available for use.
//
// Specifically, disabled inputs do not receive the click event,
// and disabled inputs are not submitted with the form.
const Disabled = html.Attribute("disabled")

// Form element with which the input is associated (that is, its
// form owner). This string's value, if present, must match the
// id of a <form> element in the same document. If this attribute
// isn't specified, the <input> element is associated with the
// nearest containing form, if any.
//
// The form attribute lets you place an input anywhere in the
// document but have it included with a form elsewhere in the document.
func Form(id html.ID) html.Attribute {
	return html.Attr("form", string(id))
}

// FormAction is a URL that processes the information submitted
// by the button. Overrides the action attribute of the button's
// form owner. Does nothing if there is no form owner.
type FormAction string

/*
	If the button is a submit button (it's inside/associated with a
	<form> and doesn't have type="button"), specifies how to encode
	the form data that is submitted.
*/
const (
	FormEncodingDefault   = html.Attribute("formenctype=application/x-www-form-urlencoded")
	FormEncodingMultipart = html.Attribute("formenctype=multipart/form-data")
	FormEncodingTextPlain = html.Attribute("formenctype=text/plain")
)

/*
	If the button is a submit button (it's inside/associated with a
	<form> and doesn't have type="button"), this attribute specifies
	the HTTP method used to submit the form.
*/
const (
	/*
		The data from the form are included in the body of the HTTP
		request when sent to the server. Use when the form contains
		information that shouldn't be public, like login credentials.
	*/
	FormMethodPost = html.Attribute("formmethod=post")

	/*
		The form data are appended to the form's action URL, with a ?
		as a separator, and the resulting URL is sent to the server.
		Use this method when the form has no side effects, like search forms.
	*/
	FormMethodGet = html.Attribute("formmethod=get")
)

/*
	If the button is a submit button, this Boolean attribute specifies
	that the form is not to be validated when it is submitted. If this
	attribute is specified, it overrides the novalidate attribute
	of the button's form owner.
*/
const FormNoValidate = html.Attribute("formnovalidate")

/*
	If the button is a submit button, this attribute is an author-defined
	name or standardized, underscore-prefixed keyword indicating where to
	display the response from submitting the form. This is the name of, or
	keyword for, a browsing context (a tab, window, or <iframe>). If this
	attribute is specified, it overrides the target attribute of the
	button's form owner.
*/
func FormTarget(target link.Target) html.Attribute {
	return html.Attr("formtarget", string(target))
}

// Width is the intrinsic width of the image in pixels.
type Width uint

// RenderAttr implements the attr.Renderer interface.
func (w Width) RenderAttr() []byte {
	return []byte(html.Attr("width", fmt.Sprint(w)))
}

// The height of the image in pixels.
type Height uint

// RenderAttr implements the attr.Renderer interface.
func (h Height) RenderAttr() []byte {
	return []byte(html.Attr("height", fmt.Sprint(h)))
}

/*
	The value given to the list attribute should be the id
	of a <datalist> element located in the same document.
	The <datalist> provides a list of predefined values to
	suggest to the user for this input. Any values in the
	list that are not compatible with the type are not
	included in the suggested options. The values provided
	are suggestions, not requirements: users can select
	from this predefined list or provide a different value.

	It is valid on text, search, url, tel, email, date, month,
	week, time, datetime-local, number, range, and color.

	Per the specifications, the list attribute is not supported
	by the hidden, password, checkbox, radio, file, or any of
	the button types.

	Depending on the browser, the user may see a custom color
	palette suggested, tic marks along a range, or even a input
	that opens like a <select> but allows for non-listed values.
	Check out the browser compatibility table for the other
	input types.
*/
func List(id html.ID) html.Attribute {
	return html.Attr("list", string(id))
}

/*
	Valid for date, month, week, time, datetime-local, number,
	and range, it defines the greatest value in the range of
	permitted values. If the value entered into the element
	exceeds this, the element fails constraint validation. If
	the value of the max attribute isn't a number, then the
	element has no maximum value.

	There is a special case: if the data type is periodic
	(such as for dates or times), the value of max may be
	lower than the value of min, which indicates that the
	range may wrap around; for example, this allows you to
	specify a time range from 10 PM to 4 AM.
*/
func Max[T any](v T) html.Attribute {
	return html.Attr("max", fmt.Sprint(v))
}

/*
	Valid for text, search, url, tel, email, and password, it
	defines the maximum number of characters (as UTF-16 code units)
	the user can enter into the field. This must be an integer value
	0 or higher. If no maxlength is specified, or an invalid value
	is specified, the field has no maximum length. This value must
	also be greater than or equal to the value of minlength.

	The input will fail constraint validation if the length of the
	text entered into the field is greater than maxlength UTF-16
	code units long. By default, browsers prevent users from
	entering more characters than allowed by the maxlength
	attribute. See Client-side validation for more information.
*/
type MaxLength uint

// RenderAttr implements the attr.Renderer interface.
func (l MaxLength) RenderAttr() []byte {
	return []byte(html.Attr("maxlength", fmt.Sprint(l)))
}

/*
	Valid for date, month, week, time, datetime-local, number, and
	range, it defines the most negative value in the range of
	permitted values. If the value entered into the element is
	less than this this, the element fails constraint validation.
	If the value of the min attribute isn't a number, then the
	element has no minimum value.

	This value must be less than or equal to the value of the max
	attribute. If the min attribute is present but is not
	specified or is invalid, no min value is applied. If the
	min attribute is valid and a non-empty value is less than
	the minimum allowed by the min attribute, constraint
	validation will prevent form submission. See Client-side
	validation for more information.

	There is a special case: if the data type is periodic (such
	as for dates or times), the value of max may be lower than
	the value of min, which indicates that the range may wrap
	around; for example, this allows you to specify a time
	range from 10 PM to 4 AM.
*/
func Min[T any](v T) html.Attribute {
	return html.Attr("min", fmt.Sprint(v))
}

/*
	Valid for text, search, url, tel, email, and password, it defines
	the minimum number of characters (as UTF-16 code units) the user
	can enter into the entry field. This must be an non-negative
	integer value smaller than or equal to the value specified by
	maxlength. If no minlength is specified, or an invalid value
	is specified, the input has no minimum length.

	The input will fail constraint validation if the length of the
	text entered into the field is fewer than minlength UTF-16 code
	units long, preventing form submission. See Client-side
	validation for more information.
*/
type MinLength uint

// RenderAttr implements the attr.Renderer interface.
func (l MinLength) RenderAttr() []byte {
	return []byte(html.Attr("minlength", fmt.Sprint(l)))
}

// Multiple if set, means the user can enter comma separated email
// addresses in the email widget or can choose more than one file
// with the file input. See the email and file input type.
const Multiple = html.Attribute("multiple")

/*
	Name for the input control. This name is submitted along with
	the control's value when the form data is submitted.

	Consider the name a required attribute (even though it's not).
	If an input has no name specified, or name is empty, the input's
	value is not submitted with the form! (Disabled controls,
	unchecked radio buttons, unchecked checkboxes, and reset
	buttons are also not sent.)

	There are two special cases:

		1. _charset_ : If used as the name of an <input> element of
		   type hidden, the input's value is automatically set by
		   the user agent to the character encoding being used to
		   submit the form.
		2. isindex: For historical reasons, the name isindex is
		   not allowed.

	The name attribute creates a unique behavior for radio buttons.

	Only one radio button in a same-named group of radio buttons can
	be checked at a time. Selecting any radio button in that group
	automatically deselects any currently-selected radio button in
	the same group. The value of that one checked radio button is
	sent along with the name if the form is submitted,

	When tabbing into a series of same-named group of radio buttons,
	if one is checked, that one will receive focus. If they aren't
	grouped together in source order, if one of the group is checked,
	tabbing into the group starts when the first one in the group is
	encountered, skipping all those that aren't checked. In other
	words, if one is checked, tabbing skips the unchecked radio
	buttons in the group. If none are checked, the radio button
	group receives focus when the first button in the same name
	group is reached.

	Once one of the radio buttons in a group has focus, using the
	arrow keys will navigate through all the radio buttons of the
	same name, even if the radio buttons are not grouped together
	in the source order.
*/
type Name string

// RenderAttr implements the attributes.Renderer interface.
func (n Name) RenderAttr() []byte {
	return []byte(html.Attr("name", string(n)))
}

/*
	The pattern attribute, when specified, is a regular expression
	that the input's value must match in order for the value to
	pass constraint validation. It must be a valid JavaScript
	regular expression, as used by the RegExp type, and as
	documented in our guide on regular expressions; the 'u'
	flag is specified when compiling the regular expression, so
	that the pattern is treated as a sequence of Unicode code
	points, instead of as ASCII. No forward slashes should be
	specified around the pattern text.

	If the pattern attribute is present but is not specified or is
	invalid, no regular expression is applied and this attribute
	is ignored completely. If the pattern attribute is valid and
	a non-empty value does not match the pattern, constraint
	validation will prevent form submission.
*/
type Pattern string

// RenderAttr implements the attributes.Renderer interface.
func (p Pattern) RenderAttr() []byte {
	return []byte(html.Attr("pattern", string(p)))
}

/*
	The placeholder attribute is a string that provides a brief
	hint to the user as to what kind of information is expected
	in the field. It should be a word or short phrase that provides
	a hint as to the expected type of data, rather than an
	explanation or prompt. The text must not include carriage
	returns or line feeds. So for example if a field is expected
	to capture a user's first name, and its label is "First Name",
	a suitable placeholder might be "e.g. Mustafa".
*/
type Placeholder string

// RenderAttr implements the attributes.Renderer interface.
func (p Placeholder) RenderAttr() []byte {
	return []byte(html.Attr("placeholder", string(p)))
}

/*
	ReadOnly, if present, indicates that the user
	should not be able to edit the value of the input. The readonly
	attribute is supported by the text, search, url, tel, email,
	date, month, week, time, datetime-local, number, and password
	input types.
*/
const ReadOnly = html.Attribute("readonly")

/*
	Required is a Boolean attribute which, if present, indicates
	that the user must specify a value for the input before the
	owning form can be submitted. The required attribute is
	supported by text, search, url, tel, email, date, month,
	week, time, datetime-local, number, password, checkbox,
	radio, and file inputs.
*/
const Required = html.Attribute("required")

/*
	Valid for email, password, tel, url and text input types only.
	Specifies how much of the input is shown. Basically creates
	same result as setting CSS width property with a few specialities.
	The actual unit of the value depends on the input type.
	For password and text, it is a number of characters (or em units)
	with a default value of 20, and for others, it is pixels.
	CSS width takes precedence over size attribute.
*/
type Size uint

// RenderAttr implements the attributes.Renderer interface.
func (s Size) RenderAttr() []byte {
	return []byte(html.Attr("size", fmt.Sprint(s)))
}

// Valid for the image input button only, Source specifies the URL of the
// image file to display to represent the graphical submit button.
type Source string

// RenderAttr implements the attr.Renderer interface.
func (s Source) RenderAttr() []byte {
	return []byte(html.Attr("src", string(s)))
}

/*
	Valid for the numeric input types, including number, date/time input
	types, and range, the step attribute is a number that specifies the
	granularity that the value must adhere to.

	If not explicitly included:

		- step defaults to 1 for number and range.
		- For the date/time input types, step is expressed in seconds,
		  with the default step being 60 seconds. The step scale factor
		  is 1000 (which converts the seconds to milliseconds, as used in
		  other algorithms).
*/
type Step float64

// RenderAttr implements the attr.Renderer interface.
func (s Step) RenderAttr() []byte {
	return []byte(html.Attr("step", fmt.Sprint(s)))
}

// Value is the input control's value. When specified in the HTML, this
// is the initial value, and from then on it can be altered or retrieved
// at any time using JavaScript to access the respective HTMLInputElement
// object's value property. The value attribute is always optional, though
// should be considered mandatory for checkbox, radio, and hidden.
type Value string

// RenderAttr implements the attr.Renderer interface.
func (v Value) RenderAttr() []byte {
	return []byte(html.Attr("value", string(v)))
}
