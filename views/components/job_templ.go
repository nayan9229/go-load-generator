// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.731
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"fmt"
	"github.com/nayan9229/go-load-generator/model"
)

func Job(job *model.Job) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li class=\"border-t border-gray-200\"><div class=\"px-4 py-5 sm:px-6\"><h3 class=\"text-lg leading-6 font-medium text-gray-900\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(job.ID.String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/components/job.templ`, Line: 12, Col: 79}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h3><p class=\"mt-1 max-w-2xl text-sm text-gray-500\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(job.Uri)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/components/job.templ`, Line: 13, Col: 63}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p><p class=\"mt-1 max-w-2xl text-sm text-gray-500\">Completed at: ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(job.CreatedAt.Format("2006-01-02 15:04:05"))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/components/job.templ`, Line: 14, Col: 113}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p><div class=\"mt-4 flex items-center justify-between\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if job.Status == model.Completed {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<p class=\"text-sm font-medium text-gray-500\">Status: <span class=\"text-green-600\">Completed</span></p><a href=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var5 templ.SafeURL = templ.URL(fmt.Sprintf("/result/%s", job.ID.String()))
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var5)))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"font-medium text-indigo-600 hover:text-indigo-500\">View result</a> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if job.Status == model.Running {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<p class=\"text-sm font-medium text-gray-500\">Status: <span class=\"text-green-600\">Running</span></p>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if job.Status == model.Pending {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<p class=\"text-sm font-medium text-gray-500\">Status: <span class=\"text-yellow-600\">Pending</span></p><a href=\"#\" class=\"font-medium text-indigo-600 hover:text-indigo-500\">Edit</a> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if job.Status == model.Failed {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<p class=\"text-sm font-medium text-gray-500\">Status: <span class=\"text-red-600\">Failed</span></p><a href=\"#\" class=\"font-medium text-indigo-600 hover:text-indigo-500\">Retry</a>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div></li>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
