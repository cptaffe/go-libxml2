package libxml2

import (
	"bytes"
	"io"

	"github.com/cptaffe/go-libxml2/clib"
	"github.com/cptaffe/go-libxml2/dom"
	"github.com/cptaffe/go-libxml2/parser"
	"github.com/cptaffe/go-libxml2/types"
	"github.com/pkg/errors"
)

func ParseHTMLString(content, url string, options ...parser.HTMLOption) (types.Document, error) {
	return ParseHTML([]byte(content), url, "UTF-8", options...)
}

// ParseHTML parses an HTML document. You can omit the options
// argument, or you can provide one bitwise-or'ed option
func ParseHTML(content []byte, url, encoding string, options ...parser.HTMLOption) (types.Document, error) {
	var option parser.HTMLOption
	if len(options) > 0 {
		option = options[0]
	} else {
		option = parser.DefaultHTMLOptions
	}
	docptr, err := clib.HTMLReadDoc(content, url, encoding, int(option))
	if err != nil {
		return nil, errors.Wrap(err, "failed to read document")
	}

	if docptr == 0 {
		return nil, errors.Wrap(clib.ErrInvalidDocument, "failed to get valid document pointer")
	}
	return dom.WrapDocument(docptr), nil
}

// ParseHTMLReader parses an HTML document. You can omit the options
// argument, or you can provide one bitwise-or'ed option
func ParseHTMLReader(in io.Reader, url, encoding string, options ...parser.HTMLOption) (types.Document, error) {
	buf := &bytes.Buffer{}
	if _, err := buf.ReadFrom(in); err != nil {
		return nil, errors.Wrap(err, "failed to rea from io.Reader")
	}

	return ParseHTML(buf.Bytes(), url, encoding, options...)
}
