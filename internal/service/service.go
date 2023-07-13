package service

import (
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/cloudwego/thriftgo/parser"
)

const (
	ApiGet        = "api.get"
	ApiPost       = "api.post"
	ApiPut        = "api.put"
	ApiPatch      = "api.patch"
	ApiDelete     = "api.delete"
	ApiOptions    = "api.options"
	ApiHEAD       = "api.head"
	ApiAny        = "api.any"
	ApiPath       = "api.path"
	ApiSerializer = "api.serializer"
	ApiGenPath    = "api.handler_path"
)

var (
	HttpMethodAnnotations = map[string]string{
		ApiGet:     "GET",
		ApiPost:    "POST",
		ApiPut:     "PUT",
		ApiPatch:   "PATCH",
		ApiDelete:  "DELETE",
		ApiOptions: "OPTIONS",
		ApiHEAD:    "HEAD",
		ApiAny:     "ANY",
	}
)

type Service struct {
	Name   string
	Routes map[string][]string
	Path   string
}

// Returns a map of service name to IDL files paths for the specified directory
func GetServicesFromIDLDir(idlDir string) ([]*Service, error) {
	idls := find(idlDir, ".thrift")

	ss := []*Service{}
	for _, idl := range idls {
		services, err := GetServicesFromIDL(idl)

		if err != nil {
			return nil, err
		}

		ss = append(ss, services...)
	}

	return ss, nil
}

func find(root, ext string) []string {
	var a []string
	filepath.WalkDir(root, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		if filepath.Ext(d.Name()) == ext {
			a = append(a, s)
		}
		return nil
	})
	return a
}

// Returns the services from the given IDL file
func GetServicesFromIDL(idlPath string) ([]*Service, error) {
	t, err := parser.ParseFile(idlPath, []string{""}, true)

	if err != nil {
		return nil, err
	}

	services := t.Services
	ss := []*Service{}

	for _, service := range services {
		s := NewService(service, idlPath)
		ss = append(ss, s)
	}

	return ss, nil
}

func NewService(ps *parser.Service, path string) *Service {
	s := &Service{
		Name:   ps.Name,
		Routes: make(map[string][]string),
		Path:   path,
	}

	fns := ps.GetFunctions()
	for _, fn := range fns {
		annotations := fn.GetAnnotations()
		a := getAnnotations(annotations, HttpMethodAnnotations)
		for k, v := range a {
			s.Routes[k] = append(s.Routes[k], v...)
		}
	}

	return s
}

func getAnnotations(input parser.Annotations, targets map[string]string) map[string][]string {
	if len(input) == 0 || len(targets) == 0 {
		return nil
	}
	out := map[string][]string{}
	for k, t := range targets {
		var ret *parser.Annotation
		for _, anno := range input {
			if strings.ToLower(anno.Key) == k {
				ret = anno
				break
			}
		}
		if ret == nil {
			continue
		}
		out[t] = ret.Values
	}
	return out
}
