package quill

import (
	"io"
	"strconv"
)

type boldFormat struct{}

func (*boldFormat) Fmt() *Format {
	return &Format{
		Val:   "strong",
		Place: Tag,
	}
}

func (*boldFormat) HasFormat(o *Op) bool {
	return o.HasAttr("bold")
}

type imageFormat struct {
	src, alt string
}

func (*imageFormat) Fmt() *Format { return new(Format) } // The body contains the entire element.

func (imf *imageFormat) HasFormat(o *Op) bool {
	return o.Type == "image" && o.Data == imf.src
}

// imageFormat implements the FormatWriter interface.
func (imf *imageFormat) Write(buf io.Writer) {

	buf.Write([]byte("<img src="))
	buf.Write([]byte(strconv.Quote(imf.src)))
	if imf.alt != "" {
		buf.Write([]byte(" alt="))
		buf.Write([]byte(strconv.Quote(imf.alt)))
	}
	buf.Write([]byte{'>'})

}

type italicFormat struct{}

func (*italicFormat) Fmt() *Format {
	return &Format{
		Val:   "em",
		Place: Tag,
	}
}

func (*italicFormat) HasFormat(o *Op) bool {
	return o.HasAttr("italic")
}

type colorFormat struct {
	c string
}

func (cf *colorFormat) Fmt() *Format {
	return &Format{
		Val:   "color:" + cf.c + ";",
		Place: Style,
	}
}

func (cf *colorFormat) HasFormat(o *Op) bool {
	return o.Attrs["color"] == cf.c
}

type linkFormat struct {
	href string
}

func (*linkFormat) Fmt() *Format { return new(Format) } // a wrapper only

func (lf *linkFormat) HasFormat(o *Op) bool {
	return o.Attrs["link"] == lf.href
}

func (lf *linkFormat) PreWrap(_ []*Format) string {
	return `<a href=` + strconv.Quote(lf.href) + ` target="_blank">`
}

func (lf *linkFormat) PostWrap(_ []*Format, o *Op) string {
	if lf.HasFormat(o) {
		return ""
	}
	return "</a>"
}
