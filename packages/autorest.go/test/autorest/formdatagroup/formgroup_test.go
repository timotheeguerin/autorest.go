// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package formdatagroup

import (
	"context"
	"generatortests"
	"io"
	"strings"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming"
	"github.com/stretchr/testify/require"
)

func newFormdataClient(t *testing.T) *FormdataClient {
	client, err := NewFormdataClient(&azcore.ClientOptions{
		TracingProvider: generatortests.NewTracingProvider(t),
	})
	require.NoError(t, err)
	return client
}

func NewFormdataClient(options *azcore.ClientOptions) (*FormdataClient, error) {
	client, err := azcore.NewClient("formdatagroup.FormdataClient", generatortests.ModuleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &FormdataClient{internal: client}, nil
}

func TestUploadFile(t *testing.T) {
	client := newFormdataClient(t)
	s := strings.NewReader("the data")
	resp, err := client.UploadFile(context.Background(), streaming.NopCloser(s), "sample", nil)
	require.NoError(t, err)
	b, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	if string(b) != "the data" {
		t.Fatalf("unexpected result %s", string(b))
	}
}

func TestUploadFileViaBody(t *testing.T) {
	client := newFormdataClient(t)
	s := strings.NewReader("the data")
	resp, err := client.UploadFileViaBody(context.Background(), streaming.NopCloser(s), nil)
	require.NoError(t, err)
	b, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	if string(b) != "the data" {
		t.Fatalf("unexpected result %s", string(b))
	}
}

func TestUploadFiles(t *testing.T) {
	t.Skip("missing route in test server")
	client := newFormdataClient(t)
	s1 := strings.NewReader("the data")
	s2 := strings.NewReader(" to be uploaded")
	resp, err := client.UploadFiles(context.Background(), []io.ReadSeekCloser{
		streaming.NopCloser(s1),
		streaming.NopCloser(s2),
	}, nil)
	require.NoError(t, err)
	b, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	if string(b) != "the data" {
		t.Fatalf("unexpected result %s", string(b))
	}
}
