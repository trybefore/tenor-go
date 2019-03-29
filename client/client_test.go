package main

import (
	"fmt"
	"testing"
)

func TestTenorClient_Search(t *testing.T) {
	type args struct {
		endpoint Endpoint
		query    map[string]interface{}
	}
	tests := []struct {
		name string
		t    *TenorClient
		args args
		want *TenorResponse
	}{
		{
			name: "test 1",
			t:    NewTenorClient("ETFH67ZEE2WH"),
			args: args{endpoint: SearchEndpoint, query: map[string]interface{}{"q": "headpat anime"}},
			want: &TenorResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Printf("%v", tt.t.search(tt.args.endpoint, tt.args.query))
		})
	}
}
