package main

import (
	"bytes"
	"errors"
	"fmt"
	"go/format"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"unicode"

	flags "github.com/jessevdk/go-flags"
	openapi "github.com/nasa9084/go-openapi"
)

const (
	arrayType       = "array"
	objectType      = "object"
	stringType      = "string"
	applicationjson = "application/json"
)

type options struct {
	SpecFile string `short:"f" long:"file"`
}

func main() { os.Exit(exec()) }

func exec() int {
	var opts options
	if _, err := flags.Parse(&opts); err != nil {
		log.Print(err)
		return 1
	}
	doc, err := openapi.Load(opts.SpecFile)
	if err != nil {
		log.Print(err)
		return 1
	}
	if err := generate(doc); err != nil {
		log.Print(err)
		return 1
	}
	return 0
}

func writeTo(src []byte, filename string) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(src)
	return err
}

func generate(doc *openapi.Document) error {
	if err := generateRoutes(doc.Paths); err != nil {
		return err
	}
	if err := generateRequestBodies(doc.Components.RequestBodies); err != nil {
		return err
	}
	if err := generateResponses(doc.Components.Responses); err != nil {
		return err
	}
	if err := generateSchemas(doc.Components.Schemas); err != nil {
		return err
	}
	if err := generateHandlers(doc.Paths); err != nil {
		return err
	}
	if _, err := os.Stat("logic.go"); os.IsNotExist(err) {
		if err := generateLogics(doc.Paths); err != nil {
			return err
		}
	}
	return nil
}

type buffer struct {
	bytes.Buffer
}

func (buf *buffer) WriteStringf(format string, args ...interface{}) (int, error) {
	return fmt.Fprintf(buf, format, args...)
}

func generateRoutes(paths openapi.Paths) error {
	var buf buffer
	buf.WriteString(`package faultinfo
// Code generated genserver.go. DO NOT EDIT.

import (
"net/http"

"github.com/gorilla/mux"
)`)
	buf.WriteString("\n\nfunc bindroutes(router *mux.Router) {")
	for path, pathitem := range paths {
		if pathitem.Get != nil {
			generateRoute(&buf, path, pathitem.Get.OperationID, http.MethodGet)
		}
		if pathitem.Post != nil {
			generateRoute(&buf, path, pathitem.Post.OperationID, http.MethodPost)
		}
		if pathitem.Put != nil {
			generateRoute(&buf, path, pathitem.Put.OperationID, http.MethodPut)
		}
		if pathitem.Delete != nil {
			generateRoute(&buf, path, pathitem.Delete.OperationID, http.MethodDelete)
		}
	}
	buf.WriteString("\n}")
	src, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}
	return writeTo(src, "routings_gen.go")
}

func generateRoute(buf *buffer, path, opid, method string) {
	buf.WriteStringf(
		"\nrouter.HandleFunc(%s, %sHandler).Methods(http.Method%s)",
		strconv.Quote(path), opid, string([]rune(method)[0])+string([]rune(strings.ToLower(method))[1:]),
	)
}

func generateRequestBodies(requestBodies map[string]*openapi.RequestBody) error {
	var buf buffer
	buf.WriteString(`package input
// Code generated by genserver.go. DO NOT EDIT.
`)
	for name, requestBody := range requestBodies {
		if err := generateRequestBody(&buf, name, requestBody); err != nil {
			return err
		}
	}

	src, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}
	return writeTo(src, "input/requestbody_gen.go")
}

func generateRequestBody(buf *buffer, name string, requestBody *openapi.RequestBody) error {
	buf.WriteStringf("\n\ntype %s ", name)
	switch requestBody.Content[applicationjson].Schema.Type {
	case objectType:
		buf.WriteString("struct {")
		for _, p := range requestBody.Content[applicationjson].Schema.Properties {
			if p.Type == objectType && p.EnableAdditionalProperties {
				buf.WriteStringf("\n%s map[string]string", p.Title)
				continue
			}
			buf.WriteStringf("\n%s %s", p.Title, p.Type)
		}
		buf.WriteString("\n}")
	default:
		buf.WriteString(requestBody.Content[applicationjson].Schema.Type)
	}
	return nil
}

func generateResponses(responses map[string]*openapi.Response) error {
	var buf buffer
	buf.WriteString(`package output
// Code generated by genserver.go. DO NOT EDIT.
`)
	for name, response := range responses {
		if err := generateResponse(&buf, name, response); err != nil {
			return err
		}
	}

	src, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}
	return writeTo(src, "output/response_gen.go")
}

func generateResponse(buf *buffer, name string, response *openapi.Response) error {
	buf.WriteStringf("\n\ntype %s struct {", name)
	for n, p := range response.Content[applicationjson].Schema.Properties {
		buf.WriteStringf("\n%s %s `json:%s`", p.Title, p.Type, strconv.Quote(n))
	}
	buf.WriteString("\n}")
	return nil
}

func generateSchemas(schemas map[string]*openapi.Schema) error {
	var buf buffer
	buf.WriteString(`package schemas
// Code generated by genserver.go. DO NOT EDIT.

import (
"time"
)
`)
	for name, schema := range schemas {
		if err := generateSchema(&buf, name, schema); err != nil {
			return err
		}
	}
	src, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}
	return writeTo(src, "schemas/schema_gen.go")
}

func generateSchema(buf *buffer, name string, schema *openapi.Schema) error {
	buf.WriteStringf("\n\ntype %s ", name)
	if schema.Type != objectType {
		buf.WriteString(schema.Type)
		return nil
	}
	buf.WriteString("struct {")
	for n, p := range schema.Properties {
		buf.WriteString("\n" + p.Title)
		switch p.Type {
		case stringType:
			var typ string
			switch p.Format {
			case "":
				typ = stringType
			case "date-time":
				typ = "time.Time"
			}
			buf.WriteStringf(" %s", typ)
		case arrayType:
			buf.WriteStringf(" []%s", p.Items.Type)
		case "boolean":
			buf.WriteString(" bool")
		default:
			buf.WriteStringf(" %s", p.Type)
		}
		buf.WriteStringf(" `json:%s`", strconv.Quote(n))
	}
	buf.WriteString("\n}")
	return nil
}

func generateHandlers(paths openapi.Paths) error {
	var buf buffer
	buf.WriteString(`package faultinfo
// Code generated by genserver.go. DO NOT EDIT.

import (
"bytes"
"encoding/json"
"net/http"
"strings"

"github.com/gorilla/mux"

"github.com/charakoba-com/everything/faultinfo/input"
)

const bearerType = "Bearer"

func Authenticate(token string) bool {
return true
}
`)
	for _, pathitem := range paths {
		if err := generateHandler(&buf, pathitem.Get, pathitem.Parameters); err != nil {
			return err
		}
		if err := generateHandler(&buf, pathitem.Post, pathitem.Parameters); err != nil {
			return err
		}
		if err := generateHandler(&buf, pathitem.Put, pathitem.Parameters); err != nil {
			return err
		}
		if err := generateHandler(&buf, pathitem.Delete, pathitem.Parameters); err != nil {
			return err
		}
	}
	src, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}
	return writeTo(src, "handlers_gen.go")
}

func generateHandler(buf *buffer, op *openapi.Operation, pathParam []*openapi.Parameter) error {
	if op == nil {
		// skip if undefined path
		return nil
	}

	buf.WriteStringf("\n\nfunc %sHandler(w http.ResponseWriter, r *http.Request) {", op.OperationID)

	// request header
	generateParseHeader(buf, op.Security)
	parameters := append(pathParam, op.Parameters...)
	if err := generateParseParams(buf, parameters); err != nil {
		return err
	}

	// request body
	if op.RequestBody != nil {
		buf.WriteString("\nvar req input." + strings.Split(op.RequestBody.Ref, "/")[3])
		buf.WriteString("\nif err := json.NewDecoder(r.Body).Decode(&req); err != nil {")
		buf.WriteString("\nw.WriteHeader(http.StatusBadRequest)")
		buf.WriteString("\nreturn")
		buf.WriteString("\n}")
	}

	hasResponseBody := false
	var respStatus string
	buf.WriteString("\n")
	if r, ok := op.Responses["200"]; ok && r.Content != nil {
		buf.WriteString("out, ")
		hasResponseBody = true
		respStatus = "http.StatusOK"
	}
	if r, ok := op.Responses["201"]; ok && r.Content != nil {
		buf.WriteString("out, ")
		hasResponseBody = true
		respStatus = "http.StatusCreated"
	}

	buf.WriteStringf("err := %s(", op.OperationID)
	var args []string
	args = append(args, "r.Context()")
	for _, p := range parameters {
		args = append(args, camelCase(p.Name))
	}
	if op.RequestBody != nil {
		args = append(args, "req")
	}
	buf.WriteString(strings.Join(args, ", "))
	buf.WriteString(")")
	buf.WriteString("\nif err != nil {")
	buf.WriteString("\nw.WriteHeader(http.StatusInternalServerError)")
	buf.WriteString("\nreturn")
	buf.WriteString("\n}")
	if hasResponseBody {
		buf.WriteString("\nvar buf bytes.Buffer")
		buf.WriteString("\nif err := json.NewEncoder(&buf).Encode(out); err != nil {")
		buf.WriteString("\nw.WriteHeader(http.StatusInternalServerError)")
		buf.WriteString("\nreturn")
		buf.WriteString("\n}")
		buf.WriteStringf("\nw.WriteHeader(%s)", respStatus)
		buf.WriteString("\nbuf.WriteTo(w)")
	}
	buf.WriteString("\n}")

	return nil
}

func generateParseHeader(buf *buffer, security *openapi.SecurityRequirement) {
	if security == nil || (*security)[0]["BearerAuth"] == nil {
		return
	}
	buf.WriteStringf("\nauthHeader := strings.Split(r.Header.Get(%s), %s)", strconv.Quote("Authorization"), strconv.Quote(" "))
	buf.WriteString("\nif len(authHeader) != 2 || authHeader[0] != bearerType {")
	buf.WriteString("\nw.WriteHeader(http.StatusUnauthorized)")
	buf.WriteString("\nreturn")
	buf.WriteString("\n}")
	buf.WriteString("\nbearerToken := authHeader[1]")
	buf.WriteString("\nif !Authenticate(bearerToken) {")
	buf.WriteStringf("\nw.Header().Set(%s, %s)", strconv.Quote("WWW-Authenticate"), strconv.Quote("Bearer realm=\"ident\""))
	buf.WriteString("\nw.WriteHeader(http.StatusUnauthorized)")
	buf.WriteString("\nreturn")
	buf.WriteString("\n}")
}

func generateParseParams(buf *buffer, params []*openapi.Parameter) error {
	for _, p := range params {
		if p.In != "path" {
			continue
		}
		if p.Name == "" {
			return errors.New("path parameter name is empty")
		}
		if p.Name == "type" {
			return errors.New("cannot use `type` as path parameter name")
		}
		buf.WriteStringf("\n%s := mux.Vars(r)[%s]", camelCase(p.Name), strconv.Quote(p.Name))
	}
	return nil
}

func camelCase(snakeCase string) string {
	var s []rune
	var afterUS bool
	for _, r := range snakeCase {
		if r == '_' {
			afterUS = true
			continue
		}
		if afterUS {
			r = unicode.ToUpper(r)
		}
		s = append(s, r)
		afterUS = false
	}
	return string(s)
}

func generateLogics(paths openapi.Paths) error {
	var buf buffer
	buf.WriteString(`package faultinfo

import (
"context"

"github.com/charakoba-com/everything/faultinfo/input"
"github.com/charakoba-com/everything/faultinfo/output"
"github.com/charakoba-com/everything/faultinfo/schemas"
)
`)
	for _, pathitem := range paths {
		if err := generateLogic(&buf, pathitem.Get, pathitem.Parameters); err != nil {
			return err
		}
		if err := generateLogic(&buf, pathitem.Post, pathitem.Parameters); err != nil {
			return err
		}
		if err := generateLogic(&buf, pathitem.Put, pathitem.Parameters); err != nil {
			return err
		}
		if err := generateLogic(&buf, pathitem.Delete, pathitem.Parameters); err != nil {
			return err
		}
	}

	src, err := format.Source(buf.Bytes())
	//src = buf.Bytes()
	//err = nil
	if err != nil {
		return err
	}

	return writeTo(src, "logic.go")
}

func generateLogic(buf *buffer, op *openapi.Operation, pathParam []*openapi.Parameter) error {
	if op == nil {
		return nil
	}
	params := append(op.Parameters, pathParam...)
	buf.WriteStringf("\n\nfunc %s(ctx context.Context, ", op.OperationID)
	var args []string
	for _, param := range params {
		args = append(args, fmt.Sprintf("%s %s", camelCase(param.Name), param.Schema.Type))
	}
	if op.RequestBody != nil {
		args = append(args, fmt.Sprintf("req input.%s", strings.Split(op.RequestBody.Ref, "/")[3]))
	}
	buf.WriteString(strings.Join(args, ", "))
	buf.WriteString(") (")

	resp, err := getSuccessResponse(op)
	if err != nil {
		return err
	}

	if resp.Content != nil {
		buf.WriteString("out ")
		schema := resp.Content[applicationjson].Schema
		typ := ref2type(schema.Ref)
		if typ == "" {
			if schema.Type == arrayType {
				typ = "[]" + schema.Items.Type
				if typ == "[]" {
					typ += ref2type(schema.Items.Ref)
				}
			} else {
				typ = schema.Type
			}
		}
		buf.WriteString(typ)
		buf.WriteString(", ")
	}
	buf.WriteString("err error) {")
	buf.WriteString("\nreturn")
	buf.WriteString("\n}")
	return nil
}

func getSuccessResponse(op *openapi.Operation) (*openapi.Response, error) {
	resp200, ok200 := op.Responses["200"]
	resp201, ok201 := op.Responses["201"]
	switch {
	case ok200 && ok201:
		return nil, errors.New("both ok and created response exists")
	case ok201:
		return resp201, nil
	case ok200:
		return resp200, nil
	default:
		return nil, errors.New("no ok or created response")
	}
}

func ref2type(refStr string) string {
	if refStr == "" {
		return ""
	}
	ref := strings.Split(strings.Replace(refStr, "responses", "output", -1), "/")
	return fmt.Sprintf("*%s.%s", ref[2], ref[3])
}
