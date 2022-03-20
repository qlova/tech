/*
	Package time provides the HTML <time> element.

	The <time> HTML element represents a specific period in time.
	It may include the datetime attribute to translate dates into
	machine-readable format, allowing for better search engine
	results or custom features such as reminders.

	It may represent one of the following:

		- A time on a 24-hour clock.
		- A precise date in the Gregorian calendar (with optional time and timezone information).
		- A valid time duration.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/time
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package time

import (
	"qlova.tech/new/tree"
	"qlova.tech/use/html"
)

// Tag is the html <time> tag.
const Tag = html.Tag("time")

// New returns a <time> html tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}

/*
	This attribute indicates the time and/or date of the element and must be in one of the formats described below.

	Valid Datetime Values

	a valid year string
		2011

	a valid month string
		2011-11

	a valid date string
		2011-11-18

	a valid yearless date string
		11-18

	a valid week string
		2011-W47

	a valid time string
		14:54
		14:54:39
		14:54:39.929

	a valid local date and time string
		2011-11-18T14:54:39.929
		2011-11-18 14:54:39.929

	a valid global date and time string
		2011-11-18T14:54:39.929Z
		2011-11-18T14:54:39.929-0400
		2011-11-18T14:54:39.929-04:00
		2011-11-18 14:54:39.929Z
		2011-11-18 14:54:39.929-0400
		2011-11-18 14:54:39.929-04:00

	a valid duration string
		PT4H18M3S
*/
func DateTime(value string) html.Attribute {
	return html.Attr("datetime", value)
}
