// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package service

import (
	"context"
	"errors"
	"ml_metadata/metadata_store/mlmetadata"
	mlpb "ml_metadata/proto/metadata_store_go_proto"
	"reflect"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	pb "github.com/kubeflow/metadata/api"
)

func TestValidTypeNames(t *testing.T) {
	names := []string{
		"Model",
		"MyModel123",
		"MyModel123_-123lkj",
		"myModel12-lkj*0972)-",
	}

	for _, n := range names {
		if err := validTypeName(n); err != nil {
			t.Errorf("validTypeName(%q) = %v\nWant nil error", n, err)
		}
	}
}

func TestInvalidTypeNames(t *testing.T) {
	names := []string{
		"1Model",
		"1 Model",
		" 1Model",
		"/1Model",
		"Model/",
		"Model//",
		"A1 Model",
		"A 1Model",
		"A/1Model",
		"A ",
	}

	for _, n := range names {
		if err := validTypeName(n); err == nil {
			t.Errorf("validTypeName(%q) = nil\nWant non-nil error", n)
		}
	}
}

func TestValidNamespaces(t *testing.T) {
	names := []string{
		"kubeflow.org",
		"kubeflow.org/v1",
		"Kubeflow.org/v1",
		"www.kubeflow.org/v1",
		"aa",
	}

	for _, n := range names {
		if err := validNamespace(n); err != nil {
			t.Errorf("validNamespace(%q) = %v\nWant nil error", n, err)
		}
	}
}

func TestInvalidNamespaces(t *testing.T) {
	names := []string{
		"kube flow.org",
		"123kubeflow.org",
		"kubeflow.org/v1/",
		"Kubeflow.org/v1//",
		"/kubeflow.org",
		"a",
		"/",
	}

	for _, n := range names {
		if err := validNamespace(n); err == nil {
			t.Errorf("validNamespace(%q) = nil\nWant non-nil error", n)
		}
	}
}

func TestGetNamespacedName(t *testing.T) {
	tests := []struct {
		in            string
		wantNamespace string
		wantName      string
		wantErr       error
	}{
		{"kubeflow.org/v1", "kubeflow.org", "v1", nil},
		{"kubeflow.org/v1/Model", "kubeflow.org/v1", "Model", nil},
		{"kubeflow.org/v1/Model1/Model", "kubeflow.org/v1/Model1", "Model", nil},
		{"kubeflow.org/v1/a/b/c/d", "kubeflow.org/v1/a/b/c", "d", nil},
		{"kubeflow.org/v1/", "", "", errors.New("malformed type name: \"kubeflow.org/v1/\"")},
		{"kubeflow.org", "", "", errors.New("malformed type name: \"kubeflow.org\"")},
	}

	for _, test := range tests {
		namespace, name, err := getNamespacedName(test.in)

		if !reflect.DeepEqual(err, test.wantErr) || namespace != test.wantNamespace || name != test.wantName {
			t.Errorf("getNamespacedName(%q) = %q, %q, %v\nWant %q, %q, %v",
				test.in, namespace, name, err, test.wantNamespace, test.wantName, test.wantErr)
		}
	}
}

func testMLMDStore(t *testing.T) *mlmetadata.Store {
	cfg := &mlpb.ConnectionConfig{
		Config: &mlpb.ConnectionConfig_FakeDatabase{
			&mlpb.FakeDatabaseConfig{},
		},
	}

	store, err := mlmetadata.NewStore(cfg)
	if err != nil {
		t.Fatalf("Failed to create ML Metadata Store: %v", err)
	}

	return store
}

func TestArtifactTypeCreation(t *testing.T) {
	store := testMLMDStore(t)
	svc := New(store)

	tests := []struct {
		request    string
		wantStored string
	}{
		{
			request: ` artifact_type {
				name:      "Model"
				namespace: { name: "my_namespace"}
				type_properties { key: "string_field" value { string_type {} } }
				type_properties { key: "int_field" value { int_type {} } }
				type_properties { key: "double_field" value { double_type {} } }
			}`,
			wantStored: `
				name:      "my_namespace/Model"
				properties { key: "string_field" value: STRING }
				properties { key: "int_field" value: INT }
				properties { key: "double_field" value: DOUBLE }
			`,
		},
		// TODO(neuromage): Add more test cases.
	}

	for i, test := range tests {
		req := &pb.CreateArtifactTypeRequest{}
		if err := proto.UnmarshalText(test.request, req); err != nil {
			t.Errorf("Test case %d\nproto.UnmarshalText failure: %v ", i, err)
			continue
		}
		want := req.GetArtifactType()

		ctx := context.Background()
		response, err := svc.CreateArtifactType(ctx, req)
		if err != nil {
			t.Errorf("Test case %d\nCreateArtifactType\nRequest:\n%v\nGot error:\n%v\nWant nil error\n", i, req, err)
			continue
		}

		got := response.GetArtifactType()

		if !cmp.Equal(got, want, cmpopts.IgnoreFields(pb.ArtifactType{}, "Id")) {
			t.Errorf("Test case %d\nCreateArtifactType\nRequest:\n%v\nGot:\n%v\nError:\n%v\nWant:\n%v\nDiff\n%v\n",
				i, req, got, err, want, cmp.Diff(want, got))
		}

		wantStored := &mlpb.ArtifactType{}
		if err := proto.UnmarshalText(test.wantStored, wantStored); err != nil {
			t.Errorf("Test case %d\nproto.UnmarshalText failure: %v ", i, err)
			continue
		}

		gotStored, err := store.GetArtifactTypesByID([]mlmetadata.ArtifactTypeID{mlmetadata.ArtifactTypeID(got.GetId())})
		if err != nil || len(gotStored) != 1 {
			t.Errorf("Test case %d\nstore.GetArtifactType:\n%v, %v\nWant single artifact type and nil error\n", i, gotStored, err)
			continue
		}

		if !cmp.Equal(gotStored[0], wantStored, cmpopts.IgnoreFields(mlpb.ArtifactType{}, "Id")) {
			t.Errorf("Test case %d\nStored ArtifactType:\n%v\nWant:\n%v\nDiff:\n%s",
				i, gotStored, wantStored, cmp.Diff(wantStored, gotStored[0]))
		}
	}
}
