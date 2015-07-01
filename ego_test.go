package ego

import (
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func trimSpaceLine(s string) string {
	buf := []string{}
	for _, line := range strings.Split(s, "\n") {
		line = strings.TrimSpace(line)
		if len(line) != 0 {
			buf = append(buf, line)
		}
	}
	return strings.TrimSpace(strings.Join(buf, "\n"))
}

func errorIfNotEqual(t *testing.T, v1, v2 interface{}) {
	if v1 != v2 {
		_, file, line, _ := runtime.Caller(1)
		t.Errorf("%v line %v: '%v' expected, but got '%v'", filepath.Base(file), line, v1, v2)
	}
}

type User struct {
	FirstName      string
	FavoriteColors []string
}

func TestRendererHtml(t *testing.T) {
	renderer := NewEgoRenderer()
	user := &User{
		FirstName:      "Alice",
		FavoriteColors: []string{"red", "orange"},
	}

	writer := httptest.NewRecorder()
	renderer.Html(writer, UserView, user)
	errorIfNotEqual(t, trimSpaceLine(`
<html>
  <body>
    <h1>Hello Alice!</h1>

    <p>Here's a list of your favorite colors:</p>
    <ul>
     
        <li>red</li>
      
        <li>orange</li>
      
    </ul>
  </body>
</html>
`), trimSpaceLine(writer.Body.String()))
	errorIfNotEqual(t, "text/html; charset=UTF-8", writer.Header().Get("Content-Type"))
}

func TestRendererHtmlLayout(t *testing.T) {
	renderer := NewEgoRenderer()
	user := &User{
		FirstName:      "Alice",
		FavoriteColors: []string{"red", "orange"},
	}

	writer := httptest.NewRecorder()
	renderer.Html(writer, MyContents, user)
	errorIfNotEqual(t, trimSpaceLine(`
<html>
  <body>
    <h1>Hello Alice!</h1>

    <p>Here's a list of your favorite colors:</p>
    <ul>
     
        <li>red</li>
      
        <li>orange</li>
      
    </ul>
  </body>
</html>
`), trimSpaceLine(writer.Body.String()))
	errorIfNotEqual(t, "text/html; charset=UTF-8", writer.Header().Get("Content-Type"))
}
