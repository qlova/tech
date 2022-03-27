/*
	Package bidirectionalisolation provides the HTML <bdi> element.

	The <bdi> HTML element tells the browser's bidirectional algorithm
	to treat the text it contains in isolation from its surrounding text.
	It's particularly useful when a website dynamically inserts some text
	and doesn't know the directionality of the text being inserted.

	Bidirectional text is text that may contain both sequences of characters
	that are arranged left-to-right (LTR) and sequences of characters that
	are arranged right-to-left (RTL), such as an Arabic quotation embedded
	in an English string. Browsers implement the Unicode Bidirectional
	Algorithm to handle this. In this algorithm, characters are given
	an implicit directionality: for example, Latin characters are treated
	as LTR while Arabic characters are treated as RTL. Some other characters
	(such as spaces and some punctuation) are treated as neutral and are
	assigned directionality based on that of their surrounding characters.

	Usually, the bidirectional algorithm will do the right thing without
	the author having to provide any special markup but, occasionally,
	the algorithm needs help. That's where <bdi> comes in.

	The <bdi> element is used to wrap a span of text and instructs the
	bidirectional algorithm to treat this text in isolation from its
	surroundings. This works in two ways:

	The directionality of text embedded in <bdi> does not influence
	the directionality of the surrounding text.
	The directionality of text embedded in <bdi> is not influenced by the
	directionality of the surrounding text.

	For example, consider some text like:

		EMBEDDED-TEXT - 1st place

	If EMBEDDED-TEXT is LTR, this works fine. But if EMBEDDED-TEXT is
	RTL, then - 1 will be treated as RTL text (because it consists of
	neutral and weak characters). The result will be garbled:

		1 - EMBEDDED-TEXTst place

	If you know the directionality of EMBEDDED-TEXT in advance, you
	can fix this problem by wrapping EMBEDDED-TEXT in a <span> with
	the dir attribute set to the known directionality. But if you
	don't know the directionality - for example, because EMBEDDED-TEXT
	is being read from a database or entered by the user - you should
	use <bdi> to prevent the directionality of EMBEDDED-TEXT from
	affecting its surroundings.

	Though the same visual effect can be achieved using the CSS rule
	unicode-bidi: isolate on a <span> or another text-formatting
	element, HTML authors should not use this approach because it is
	not semantic and browsers are allowed to ignore CSS styling.

	Embedding the characters in <span dir="auto"> has the same effect
	as using <bdi>, but its semantics are less clear.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/bdi
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package bidirectionalisolation

import (
	"qlova.tech/use/html"
	"qlova.tech/web/tree"
)

// Tag is the html <bdi> tag.
const Tag = html.Tag("bdi")

// New returns a <bdi> html tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}
