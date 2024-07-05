// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.731
package home

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"github.com/nayan9229/go-load-generator/model"
	"github.com/nayan9229/go-load-generator/views/components"
	"github.com/nayan9229/go-load-generator/views/layouts"
)

func Index(jobs []*model.Job) templ.Component {
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
		templ_7745c5c3_Var2 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
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
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"max-w-4xl mx-auto\"><h1 class=\"text-3xl font-bold mb-6\">Load generator</h1><!-- Add New Job Form --><div class=\"bg-white p-6 rounded-lg shadow mb-8\"><h2 class=\"text-2xl font-bold mb-4\">Add New Job</h2><form id=\"add-job-form\" class=\"space-y-4\" action=\"/\" method=\"POST\"><div><label for=\"url\" class=\"mt-1 block text-sm font-medium text-gray-700\">URL</label> <input type=\"text\" id=\"url\" name=\"url\" class=\"mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:ring-indigo-500 focus:border-indigo-500\"> <label for=\"url_timeout\" class=\"mt-1 block text-sm font-medium text-gray-700\">URL Timeout</label> <input type=\"text\" id=\"url_timeout\" name=\"url_timeout\" class=\"mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:ring-indigo-500 focus:border-indigo-500\"> <label for=\"runtime\" class=\"mt-1 block text-sm font-medium text-gray-700\">Runtime</label> <input type=\"text\" id=\"runtime\" name=\"runtime\" class=\"mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:ring-indigo-500 focus:border-indigo-500 \"> <label for=\"parallel_requests\" class=\"mt-1 block text-sm font-medium text-gray-700\">Parallel Requests</label> <input type=\"text\" id=\"parallel_requests\" name=\"parallel_requests\" class=\"mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:ring-indigo-500 focus:border-indigo-500\"></div><div><button type=\"submit\" class=\"inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500\">Add Job</button></div></form></div><!-- Job List --><div class=\"space-y-4\"><div class=\"p-8\"><ul class=\"bg-white shadow overflow-hidden mx-auto\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			for _, job := range jobs {
				templ_7745c5c3_Err = components.Job(job).Render(ctx, templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</ul></div></div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = layouts.Base().Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
