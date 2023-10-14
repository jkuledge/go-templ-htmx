// Code generated by templ@v0.2.364 DO NOT EDIT.

package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "strconv"

var CountsForm = struct {
	Global  string
	Session string
}{
	Global:  "global",
	Session: "session",
}

func Counts(global, user int) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_1 := templ.GetChildren(ctx)
		if var_1 == nil {
			var_1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<form id=\"countsForm\" action=\"/\" method=\"POST\" hx-post=\"/\" hx-select=\"#countsForm\" hx-swap=\"outerHTML\"><div class=\"columns\">")
		if err != nil {
			return err
		}
		var var_2 = []any{"column", "has-text-centered", "is-primary", "border"}
		err = templ.RenderCSSItems(ctx, templBuffer, var_2...)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("<div class=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(templ.CSSClasses(var_2).String()))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\"><h1 class=\"title is-size-1 has-text-centered\">")
		if err != nil {
			return err
		}
		var var_3 string = strconv.Itoa(global)
		_, err = templBuffer.WriteString(templ.EscapeString(var_3))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</h1><p class=\"subtitle has-text-centerd\">")
		if err != nil {
			return err
		}
		var_4 := `Global`
		_, err = templBuffer.WriteString(var_4)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</p><div><button class=\"button is-primary\" type=\"submit\" name=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(CountsForm.Global))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\" value=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(CountsForm.Global))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\">")
		if err != nil {
			return err
		}
		var_5 := `+1`
		_, err = templBuffer.WriteString(var_5)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</button></div></div>")
		if err != nil {
			return err
		}
		var var_6 = []any{"column", "has-text-centered", "border"}
		err = templ.RenderCSSItems(ctx, templBuffer, var_6...)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("<div class=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(templ.CSSClasses(var_6).String()))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\"><h1 class=\"title is-size-1 has-text-centered\">")
		if err != nil {
			return err
		}
		var var_7 string = strconv.Itoa(user)
		_, err = templBuffer.WriteString(templ.EscapeString(var_7))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</h1><p class=\"subtitle has-text-centerd\">")
		if err != nil {
			return err
		}
		var_8 := `Session`
		_, err = templBuffer.WriteString(var_8)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</p><div><button class=\"button is-secondary\" type=\"submit\" name=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(CountsForm.Session))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\" value=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(CountsForm.Session))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\">")
		if err != nil {
			return err
		}
		var_9 := `+1`
		_, err = templBuffer.WriteString(var_9)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</button></div></div></div></form>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}

func Form() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_10 := templ.GetChildren(ctx)
		if var_10 == nil {
			var_10 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<form action=\"/\" method=\"POST\"><div><button type=\"submit\" name=\"global\" value=\"global\">")
		if err != nil {
			return err
		}
		var_11 := `Global`
		_, err = templBuffer.WriteString(var_11)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</button></div><div><button type=\"submit\" name=\"user\" value=\"user\">")
		if err != nil {
			return err
		}
		var_12 := `User`
		_, err = templBuffer.WriteString(var_12)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</button></div></form>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}

func Header() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_13 := templ.GetChildren(ctx)
		if var_13 == nil {
			var_13 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<head><script src=\"/htmx.min.js\">")
		if err != nil {
			return err
		}
		var_14 := ``
		_, err = templBuffer.WriteString(var_14)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</script></head>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}
