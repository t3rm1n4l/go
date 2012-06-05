package argo

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type Namespace string

func NewNamespace(base string) (ns Namespace) {
	return Namespace(base)
}

func (ns Namespace) Get(name string) (term Term) {
	return NewResource(string(ns) + name)
}

func Lookup(prefix string) (uri string, err error) {
	reqURL := fmt.Sprintf("http://prefix.cc/%s.file.txt", prefix)

	resp, err := http.Get(reqURL)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		dataBuffer := make([]byte, 1024)
		_, err := resp.Body.Read(dataBuffer)
		if err != nil {
			return "", err
		}

		data := strings.Trim(string(dataBuffer), " \r\n\x00")
		parts := strings.Split(data, "\t")
		return parts[1], nil
	}

	return "", errors.New(fmt.Sprintf("HTTP request returned status %d", resp.StatusCode))
}

var RDF = NewNamespace("http://www.w3.org/1999/02/22-rdf-syntax-ns#")
var RDFS = NewNamespace("http://www.w3.org/2000/01/rdf-schema#")
var OWL = NewNamespace("http://www.w3.org/2002/07/owl#")
var FOAF = NewNamespace("http://xmlns.com/foaf/0.1/")
var DC = NewNamespace("http://purl.org/dc/elements/1.1/")
var DCT = NewNamespace("http://purl.org/dc/terms/")
var A = RDF.Get("type")
