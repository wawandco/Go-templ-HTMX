// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package tasks

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func New(err ...string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex gap-3 text items-center p-5 rounded-lg w-full bg-slate-800\" id=\"input\"><div><button hx-post=\"/tasks\" hx-include=\"[name=&#39;title&#39;]\" hx-target=\"#input\" hx-swap=\"outerHTML\" type=\"button\" class=\"rounded-full flex justify-center items-center p-1 border border-gray-600\" _=\"on htmx:afterRequest(issueRequest) trigger updateList on #list end\n            on htmx:afterRequest(issueRequest) trigger updateCount on #count end\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"20\" height=\"20\" viewBox=\"0 0 20 20\" fill=\"none\"><path d=\"M10 4.16667V15.8333M4.16667 10H15.8333\" stroke=\"white\" stroke-width=\"1.66667\" stroke-linecap=\"round\" stroke-linejoin=\"round\"></path></svg></button></div><div class=\"flex flex-col w-full\"><input type=\"text\" name=\"title\" autofocus hx-post=\"/tasks\" hx-include=\"this\" hx-target=\"#input\" hx-trigger=\"keyup[keyCode==13]\" hx-swap=\"outerHTML\" placeholder=\"Write a title\" class=\"text-white w-full border-none bg-transparent outline-none\" _=\"on htmx:afterRequest(issueRequest) trigger updateList on #list end\n            on htmx:afterRequest(issueRequest) trigger updateCount on #count end\"><div><p>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if len(err) > 0 {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<span class=\"text-red-500 text-sm\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var2 string = err[0]
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
