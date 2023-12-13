package pangea_proxy

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/pangeacyber/pangea-go/pangea-sdk/v3/service/audit"
)

type Proxy struct {
	client audit.Client
}

func New(client audit.Client) *Proxy {
	return &Proxy{
		client: client,
	}
}

type SafeSearchInput struct {
	Query string `json:"query"`

	// Specify the sort order of the response. "asc" or "desc"
	Order string `json:"order,omitempty"`

	// Name of column to sort the results by.
	OrderBy string `json:"order_by,omitempty"`

	// The start of the time range to perform the search on.
	Start string `json:"start,omitempty"`

	// The end of the time range to perform the search on. All records up to the latest if left out.
	End string `json:"end,omitempty"`

	// Number of audit records to include from the first page of the results.
	Limit int `json:"limit,omitempty"`

	// Maximum number of results to return.
	// min 1 max 10000
	MaxResults int `json:"max_results,omitempty"`

	// If true include root, membership and consistency proof
	Verbose *bool `json:"verbose,omitempty"`
}

func (p *Proxy) Search(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var input SafeSearchInput
	err := decodeJSONBody(w, r, &input)
	if err != nil {
		var mr *malformedRequest
		if errors.As(err, &mr) {
			http.Error(w, mr.msg, mr.status)
		} else {
			log.Print(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	search := &audit.SearchInput{
		Query:      input.Query,
		Order:      input.Order,
		OrderBy:    input.OrderBy,
		Limit:      input.Limit,
		MaxResults: input.MaxResults,
		Verbose:    input.Verbose,
	}
	resp, err := p.client.Search(ctx, search)
	if err != nil {
		log.Fatal(err)
	}

	b, err := json.Marshal(RawResponse{
		ResponseHeader: resp.ResponseHeader,
		RawResult:      resp.RawResult,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func (p *Proxy) Results(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var input audit.SearchResultsInput
	err := decodeJSONBody(w, r, &input)
	if err != nil {
		var mr *malformedRequest
		if errors.As(err, &mr) {
			http.Error(w, mr.msg, mr.status)
		} else {
			log.Print(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	resp, err := p.client.SearchResults(ctx, &input)
	if err != nil {
		log.Fatal(err)
	}

	b, err := json.Marshal(RawResponse{
		ResponseHeader: resp.ResponseHeader,
		RawResult:      resp.RawResult,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func (p *Proxy) Root(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var input audit.RootInput
	err := decodeJSONBody(w, r, &input)
	if err != nil {
		var mr *malformedRequest
		if errors.As(err, &mr) {
			http.Error(w, mr.msg, mr.status)
		} else {
			log.Print(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	resp, err := p.client.Root(ctx, &input)
	if err != nil {
		log.Fatal(err)
	}

	b, err := json.Marshal(RawResponse{
		ResponseHeader: resp.ResponseHeader,
		RawResult:      resp.RawResult,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}
