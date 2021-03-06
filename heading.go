package tocenize

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/writeas/go-strip-markdown"
)

type Heading struct {
	Title         string
	Depth         int
	Index         int
	UniqueCounter int
}

func NewHeadingATX(line string, index int) *Heading {
	return &Heading{
		Title: strings.Trim(line, "# "),
		Depth: len(line) - len(strings.TrimLeft(line, "#")),
		Index: index,
	}
}

func NewHeadingSE(title string, sep string, index int) *Heading {
	depth := 1
	if sep[0] == '-' {
		depth = 2
	}
	return &Heading{
		Title: strings.TrimSpace(title),
		Depth: depth,
		Index: index,
	}
}

func (h Heading) String() string {
	return h.string(1)
}

func (h Heading) string(minDepth int) string {
	return fmt.Sprintf(
		"%s- [%s](%s)",
		strings.Repeat("\t", h.Depth-minDepth),
		h.LinkTitle(),
		h.Anchor(),
	)
}

var (
	reLink  = regexp.MustCompile(`\[(.*?)\][\[\(].*?[\]\)]`)
	reImage = regexp.MustCompile(`\!\[.*?\]\s?[\[\(].*?[\]\)]`)
)

// LinkTitle returns a cleaned up title for use in links.
// It removes Markdown images and links, keeping only the link text.
func (h Heading) LinkTitle() string {
	t := h.Title
	t = reLink.ReplaceAllString(t, "$1") // remove links, keeping the text
	t = reImage.ReplaceAllString(t, "")  // remove images
	t = strings.TrimSpace(t)
	return t
}

var rePunct = regexp.MustCompile(`([^\w -]+)`)

func (h Heading) Anchor() string {
	// Strip Markdown
	a := stripmd.Strip(h.Title)
	a = strings.ToLower(a)
	a = rePunct.ReplaceAllString(a, "")
	a = strings.Replace(a, " ", "-", -1)
	if h.UniqueCounter > 0 {
		a = fmt.Sprintf("%s-%d", a, h.UniqueCounter)
	}
	return fmt.Sprintf("#%s", a)
}
