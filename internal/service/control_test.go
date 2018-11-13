package service

import (
	"testing"

	"github.com/moorara/flax/internal/model"
	"github.com/moorara/flax/pkg/log"
	"github.com/moorara/flax/pkg/metrics"
	"github.com/opentracing/opentracing-go/mocktracer"
	"github.com/stretchr/testify/assert"
)

func TestControlService(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			"Simple",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			logger := log.NewNopLogger()
			metrics := metrics.New("unit-test")
			tracer := mocktracer.New()
			service := NewControlService(logger, metrics, tracer)

			assert.NotNil(t, service)
		})
	}
}

func TestReadSpec(t *testing.T) {
	tests := []struct {
		name          string
		path          string
		expectedError string
		expectedSpec  *model.Spec
	}{
		{
			name:          "NoFile",
			path:          "./test/spec",
			expectedError: "no such file or directory",
			expectedSpec:  nil,
		},
		{
			name:          "UnknownFormat",
			path:          "./test/unknown.toml",
			expectedError: "unknown file format",
			expectedSpec:  nil,
		},
		{
			name:          "EmptyYAML",
			path:          "./test/empty.yaml",
			expectedError: "EOF",
			expectedSpec:  nil,
		},
		{
			name:          "EmptyJSON",
			path:          "./test/empty.json",
			expectedError: "EOF",
			expectedSpec:  nil,
		},
		{
			name:          "InvalidYAML",
			path:          "./test/error.yaml",
			expectedError: "cannot unmarshal",
			expectedSpec:  nil,
		},
		{
			name:          "InvalidJSON",
			path:          "./test/error.json",
			expectedError: "invalid character",
			expectedSpec:  nil,
		},
		{
			name:          "SimpleYAML",
			path:          "./test/simple.yaml",
			expectedError: "",
			expectedSpec: &model.Spec{
				HTTPMocks: []model.HTTPMock{
					model.HTTPMock{
						HTTPExpectation: model.HTTPExpectation{
							Path: "/health",
						},
						StatusCode: 200,
					},
				},
				RESTMock: []model.RESTMock{
					model.RESTMock{
						Name: "car",
						Store: []model.JSON{
							model.JSON{"id": "ad2bd67b-172e-4778-a8a3-7cfb626685b9", "make": "Mazda", "model": "CX-5"},
							model.JSON{"id": "26ee9c87-fdbb-48cf-be0a-add9a3d87189", "make": "Hyundai", "model": "Sonata"},
						},
					},
					model.RESTMock{
						Name: "team",
						Store: []model.JSON{
							model.JSON{"_id": "d93ce179-50f7-469e-bb36-1b3746145f00", "name": "Back-end", "tags": []interface{}{"go", "cloud"}},
							model.JSON{"_id": "8cd6ef6c-2095-4c75-bc66-6f38e785299d", "name": "Front-end", "tags": []interface{}{"react", "redux"}},
						},
					},
				}},
		},
		{
			name:          "SimpleJSON",
			path:          "./test/simple.json",
			expectedError: "",
			expectedSpec: &model.Spec{
				HTTPMocks: []model.HTTPMock{
					model.HTTPMock{
						HTTPExpectation: model.HTTPExpectation{
							Path: "/health",
						},
						StatusCode: 200,
					},
				},
				RESTMock: []model.RESTMock{
					model.RESTMock{
						Name: "car",
						Store: []model.JSON{
							model.JSON{"id": "ad2bd67b-172e-4778-a8a3-7cfb626685b9", "make": "Mazda", "model": "CX-5"},
							model.JSON{"id": "26ee9c87-fdbb-48cf-be0a-add9a3d87189", "make": "Hyundai", "model": "Sonata"},
						},
					},
					model.RESTMock{
						Name: "team",
						Store: []model.JSON{
							model.JSON{"_id": "d93ce179-50f7-469e-bb36-1b3746145f00", "name": "Back-end", "tags": []interface{}{"go", "cloud"}},
							model.JSON{"_id": "8cd6ef6c-2095-4c75-bc66-6f38e785299d", "name": "Front-end", "tags": []interface{}{"react", "redux"}},
						},
					},
				},
			},
		},
		{
			name:          "FullYAML",
			path:          "./test/full.yaml",
			expectedError: "",
			expectedSpec: &model.Spec{
				Config: model.Config{
					HTTPPort:  8080,
					HTTPSPort: 8443,
				},
				HTTPMocks: []model.HTTPMock{
					model.HTTPMock{
						HTTPExpectation: model.HTTPExpectation{
							Methods: []string{"GET"},
							Path:    "/health",
						},
						StatusCode: 200,
					},
					model.HTTPMock{
						HTTPExpectation: model.HTTPExpectation{
							Methods: []string{"GET"},
							Path:    "/current/user",
							Filters: model.HTTPFilters{
								Params: map[string]string{
									"tenantId": "\\w+",
								},
								Headers: map[string]string{
									"Authorization": "Bearer .*",
								},
							},
						},
						Delay: "100ms",
						StatusCode: 200,
						Headers: map[string]string{
							"Tenant-ID": "4c85a46f-e11b-4d5f-bca4-add78c1f6395",
						},
						Body: map[interface{}]interface{}{
							"id": "5da8349a-0707-4064-8fad-74cedb48a8fc",
							"name": "John Doe",
							"email": "john.doe@example.com",
						},
					},
				},
				RESTMock: []model.RESTMock{
					model.RESTMock{
						Name:     "car",
						BasePath: "/api/v1",
						Store: []model.JSON{
							model.JSON{"id": "ad2bd67b-172e-4778-a8a3-7cfb626685b9", "make": "Mazda", "model": "CX-5"},
							model.JSON{"id": "26ee9c87-fdbb-48cf-be0a-add9a3d87189", "make": "Hyundai", "model": "Sonata"},
						},
					},
					model.RESTMock{
						Name:       "team",
						BasePath:   "/api/v1",
						Delay:      "100ms",
						Identifier: "_id",
						ListHandle: "data",
						Filters: model.RESTFilters{
							Headers: map[string]string{
								"Authorization": "Bearer .*",
							},
						},
						Headers: map[string]string{
							"Tenant-ID": "4c85a46f-e11b-4d5f-bca4-add78c1f6395",
						},
						Store: []model.JSON{
							model.JSON{"_id": "d93ce179-50f7-469e-bb36-1b3746145f00", "name": "Back-end", "tags": []interface{}{"go", "cloud"}},
							model.JSON{"_id": "8cd6ef6c-2095-4c75-bc66-6f38e785299d", "name": "Front-end", "tags": []interface{}{"react", "redux"}},
						},
					},
				}},
		},
		{
			name:          "FullJSON",
			path:          "./test/full.json",
			expectedError: "",
			expectedSpec: &model.Spec{
				Config: model.Config{
					HTTPPort:  8080,
					HTTPSPort: 8443,
				},
				HTTPMocks: []model.HTTPMock{
					model.HTTPMock{
						HTTPExpectation: model.HTTPExpectation{
							Methods: []string{"GET"},
							Path:    "/health",
						},
						StatusCode: 200,
					},
					model.HTTPMock{
						HTTPExpectation: model.HTTPExpectation{
							Methods: []string{"GET"},
							Path:    "/current/user",
							Filters: model.HTTPFilters{
								Params: map[string]string{
									"tenantId": "\\w+",
								},
								Headers: map[string]string{
									"Authorization": "Bearer .*",
								},
							},
						},
						Delay: "100ms",
						StatusCode: 200,
						Headers: map[string]string{
							"Tenant-ID": "4c85a46f-e11b-4d5f-bca4-add78c1f6395",
						},
						Body: map[string]interface{}{
							"id": "5da8349a-0707-4064-8fad-74cedb48a8fc",
							"name": "John Doe",
							"email": "john.doe@example.com",
						},
					},
				},
				RESTMock: []model.RESTMock{
					model.RESTMock{
						Name:     "car",
						BasePath: "/api/v1",
						Store: []model.JSON{
							model.JSON{"id": "ad2bd67b-172e-4778-a8a3-7cfb626685b9", "make": "Mazda", "model": "CX-5"},
							model.JSON{"id": "26ee9c87-fdbb-48cf-be0a-add9a3d87189", "make": "Hyundai", "model": "Sonata"},
						},
					},
					model.RESTMock{
						Name:       "team",
						BasePath:   "/api/v1",
						Delay:      "100ms",
						Identifier: "_id",
						ListHandle: "data",
						Filters: model.RESTFilters{
							Headers: map[string]string{
								"Authorization": "Bearer .*",
							},
						},
						Headers: map[string]string{
							"Tenant-ID": "4c85a46f-e11b-4d5f-bca4-add78c1f6395",
						},
						Store: []model.JSON{
							model.JSON{"_id": "d93ce179-50f7-469e-bb36-1b3746145f00", "name": "Back-end", "tags": []interface{}{"go", "cloud"}},
							model.JSON{"_id": "8cd6ef6c-2095-4c75-bc66-6f38e785299d", "name": "Front-end", "tags": []interface{}{"react", "redux"}},
						},
					},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			logger := log.NewNopLogger()
			metrics := metrics.New("unit-test")
			tracer := mocktracer.New()
			service := &controlService{
				logger:  logger,
				metrics: metrics,
				tracer:  tracer,
			}

			err := service.ReadSpec(tc.path)

			if tc.expectedError == "" {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedSpec, service.spec)
			} else {
				assert.Contains(t, err.Error(), tc.expectedError)
				assert.Nil(t, service.spec)
			}
		})
	}
}
