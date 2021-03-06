/*
 * Copyright (C) 2019 Synopsys, Inc.
 *
 *  Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements. See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership. The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 *  with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations
 *  under the License.
 */

package utils

import (
	synopsysv1 "github.com/blackducksoftware/synopsys-operator/api/v1"
	//samplev1 "github.com/blackducksoftware/synopsys-operator/pkg/api/sample/v1"
)

const (
	// OPENSHIFT denotes to create an OpenShift routes
	OPENSHIFT = "OPENSHIFT"
	// NONE denotes no exposed service
	NONE = "NONE"
	// NODEPORT denotes to create a NodePort service
	NODEPORT = "NODEPORT"
	// LOADBALANCER denotes to create a LoadBalancer service
	LOADBALANCER = "LOADBALANCER"
)

const (
	// AlertCRDName is the name of an Alert CRD
	AlertCRDName = "alerts.synopsys.com"
	// BlackDuckCRDName is the name of the Black Duck CRD
	BlackDuckCRDName = "blackducks.synopsys.com"
	// OpsSightCRDName is the name of an OpsSight CRD
	OpsSightCRDName = "opssights.synopsys.com"
	// PrmCRDName is the name of the Polaris Reporting Module CRD
	PrmCRDName = "prms.synopsys.com"
	// SizeCRDName is the name of the Size CRD
	SizeCRDName = "sizes.synopsys.com"

	PolarisCRDName    = "polaris.synopsys.com"
	PolarisDBCRDName  = "polarisdbs.synopsys.com"
	AuthServerCRDName = "authservers.synopsys.com"
	ReportingCRDName  = "reportings.synopsys.com"

	// OperatorName is the name of an Operator
	OperatorName = "synopsys-operator"
	// AlertName is the name of an Alert app
	AlertName = "alert"
	// BlackDuckName is the name of the Black Duck app
	BlackDuckName = "blackduck"
	// OpsSightName is the name of an OpsSight app
	OpsSightName = "opssight"
	// PrmName is the name of the Prm app
	PrmName = "prm"

	PolarisName = "polaris"
)

// GetSampleDefaultValue creates a sample crd configuration object with defaults
//func GetSampleDefaultValue() *samplev1.SampleSpec {
//	return &samplev1.SampleSpec{
//		Namespace:   "namesapce",
//		SampleValue: "Value",
//	}
//}

// GetBlackDuckTemplate returns the required fields for Black Duck
func GetBlackDuckTemplate() *synopsysv1.BlackduckSpec {
	return &synopsysv1.BlackduckSpec{
		Size:              "Small",
		CertificateName:   "default",
		PersistentStorage: false,
		ExposeService:     NONE,
	}
}

// GetBlackDuckDefaultPersistentStorageLatest creates a Black Duck crd configuration object
// with defaults and persistent storage
func GetBlackDuckDefaultPersistentStorageLatest() *synopsysv1.BlackduckSpec {
	return &synopsysv1.BlackduckSpec{
		Namespace:         "blackduck-pvc",
		Size:              "Small",
		CertificateName:   "default",
		LivenessProbes:    false,
		PersistentStorage: true,
		ExposeService:     NONE,
		Environs:          []string{},
		ImageRegistries:   []string{},
		PVC:               []synopsysv1.PVC{},
	}
}

// GetBlackDuckDefaultExternalPersistentStorageLatest creates a BlackDuck crd configuration object
// with defaults and external persistent storage for latest BlackDuck
func GetBlackDuckDefaultExternalPersistentStorageLatest() *synopsysv1.BlackduckSpec {
	return &synopsysv1.BlackduckSpec{
		Namespace:         "synopsys-operator",
		Size:              "small",
		LivenessProbes:    false,
		PersistentStorage: true,
		PVC:               []synopsysv1.PVC{},
		CertificateName:   "default",
		Type:              "Artifacts",
		ExposeService:     NONE,
		Environs:          []string{},
		ImageRegistries:   []string{},
	}
}

// GetBlackDuckDefaultPersistentStorageV1 creates a BlackDuck crd configuration object
// with defaults and persistent storage for V1 BlackDuck
func GetBlackDuckDefaultPersistentStorageV1() *synopsysv1.BlackduckSpec {
	return &synopsysv1.BlackduckSpec{
		Namespace:         "synopsys-operator",
		Size:              "small",
		LivenessProbes:    false,
		PersistentStorage: true,
		PVC:               []synopsysv1.PVC{},
		CertificateName:   "default",
		Type:              "Artifacts",
		ExposeService:     NONE,
		Environs:          []string{},
		ImageRegistries:   []string{},
	}
}

// GetBlackDuckDefaultExternalPersistentStorageV1 creates a BlackDuck crd configuration object
// with defaults and external persistent storage for V1 BlackDuck
func GetBlackDuckDefaultExternalPersistentStorageV1() *synopsysv1.BlackduckSpec {
	return &synopsysv1.BlackduckSpec{
		Namespace:         "synopsys-operator",
		Size:              "small",
		LivenessProbes:    false,
		PersistentStorage: true,
		PVC:               []synopsysv1.PVC{},
		Type:              "Artifacts",
		ExposeService:     NONE,
	}
}

// GetBlackDuckDefaultBDBA returns a BlackDuck with BDBA
func GetBlackDuckDefaultBDBA() *synopsysv1.BlackduckSpec {
	return &synopsysv1.BlackduckSpec{
		Namespace:       "blackduck-bdba",
		CertificateName: "default",
		Environs: []string{
			"USE_BINARY_UPLOADS:1",
		},
		LivenessProbes:    false,
		PersistentStorage: false,
		Size:              "small",
		ExposeService:     NONE,
	}
}

// GetBlackDuckDefaultEphemeral returns a BlackDuck with ephemeral storage
func GetBlackDuckDefaultEphemeral() *synopsysv1.BlackduckSpec {
	return &synopsysv1.BlackduckSpec{
		Namespace:         "blackduck-ephemeral",
		CertificateName:   "default",
		LivenessProbes:    false,
		PersistentStorage: false,
		Size:              "small",
		Type:              "worker",
		ExposeService:     NONE,
	}
}

// GetBlackDuckDefaultEphemeralCustomAuthCA returns a BlackDuck with ephemeral storage
// using custom auth CA
func GetBlackDuckDefaultEphemeralCustomAuthCA() *synopsysv1.BlackduckSpec {
	return &synopsysv1.BlackduckSpec{
		Namespace:         "blackduck-auth-ca",
		CertificateName:   "default",
		LivenessProbes:    false,
		PersistentStorage: false,
		Size:              "Small",
		ExposeService:     NONE,
		AuthCustomCA:      "-----BEGIN CERTIFICATE-----\r\nMIIE1DCCArwCCQCuw9TgaoBKVDANBgkqhkiG9w0BAQsFADAsMQswCQYDVQQGEwJV\r\nUzELMAkGA1UECgwCYmQxEDAOBgNVBAMMB1JPT1QgQ0EwHhcNMTkwMjA2MDAzMjM3\r\nWhcNMjExMTI2MDAzMjM3WjAsMQswCQYDVQQGEwJVUzELMAkGA1UECgwCYmQxEDAO\r\nBgNVBAMMB1JPT1QgQ0EwggIiMA0GCSqGSIb3DQEBAQUAA4ICDwAwggIKAoICAQCr\r\nIctvPVoqRS3Ti38uFRVfJDovyi0p9PIaOmja3tMvkfecCsCVYHMo/vAy/fm9qiJI\r\nKutTwX9aLuiLO0tsDDUNwv0CrbXvuHpWvASOAdKyl6uxiYl0fq0cyBZSdKlsdDGk\r\nivENpN2gKHxDSUgAo74wUskfBrKvfKLhJhOmKCbN/NvxlsGMM5DgPgFGNegmw5r0\r\nZlDTXlWn3J/8C80dfGjT5hLr6Jtl0KTqxSREVTLT0fDk7bt9BHH/TCtNs9UwR1UI\r\nJVjjzW6pgS1DmGZ7Mfg2WBhhdDBuN0gxk/bcoiV2tfI0MLQyeVP+qWmdUXSNn9CT\r\nmpYdKezMfi5ieSy40fy23n+D1C+Xm5pnFErm3BwZYdN9gI633IBPQa0ELo28ZxhI\r\nIclGGyhUubZJ+ybNvGOIrgypTXYrZqvyWMV3qiMZb1EzpKdqAzGfsN1zmF+o4Rc3\r\ntBa2EF/lNSVCClUeFBA2UXvD/K9QA84cbLNJwpBZ9Bc6CZyvRTYGzXtAuZUVvNju\r\nMcWhsqXWzhVkChTyYicOdT8ZB+7/eC3tFyjAKSszIA5xuO8NtuIZBAc2AzRrkoE5\r\nCgHEUxNA3tbRUjYnH5HcgaQveFQtFwBWqIMxPeJixSLk2KYJSsWpTPC1x6s1IBLO\r\nITWhedDbtbs/FT9+cXd9K+/L+6UgR31oHaY/hYai1QIDAQABMA0GCSqGSIb3DQEB\r\nCwUAA4ICAQAz7aK5m9yPE/tTFQJfZRr35ug8ikBuGFvzb5s3fWYlQ1QbKUPBp9Q/\r\n1kUGJF2niOULUp5Gig6urz+E1m3wE5jgYRwZjgTmoEQEmN0/VQWTus72isWhTsZ5\r\nJKDSzcKGRJnHzO91gA3ZP1Cxoin5GX6w8eqEA2vh1hc7+GyKPTOsxu8hYMYI1yId\r\nfWAjqEUobLZZoijf+c3AqBVcf4tOpFMRTy4au3H+v7TNjc/fAeZUeAz7BswfqEV9\r\n0QNNTpezq5IS+pSPShRatL9k/BaE3MaF0Ossfnv3UPV80Yrup+9pRV8Lu6EXrdg5\r\n3L2+KK2Nz9A+iF2u9VqUw9lcJCIjgY+APf6Tf2AKQxNCA/pV1z0I8aQAlSLolgpx\r\nSMLwMecpjAcHPWF5ut3Re+8PfeyLGzeXCVyhZc9Aj9KaTNLRa/kb21KNVbcGGTu/\r\nuiGMEJXq1a1fKzMKTPnARz70XCS7nLJ7qEK3TuvrMhCqEEdFUf/S4yAmmWaEO9Fr\r\nUBk9ACW9UYBFtowqbJkbJm3KEXMMFP5cs33j/HEA1IkKDVT9Hi7NEK2/Y7e9afv7\r\no1UGNrGgU1rK8K+/2htOH9JhlPFWHQkk+wvGL6fFI7p+6TGes0KILN4WioOEKY0t\r\n0V1Zr8bejDW49cu1Awy443SrauhFLOInubZLA8S9ZvwTVIvpmTDjdQ==\r\n-----END CERTIFICATE-----",
	}
}

// GetBlackDuckDefaultExternalDB returns a BlackDuck with an external Data Base
func GetBlackDuckDefaultExternalDB() *synopsysv1.BlackduckSpec {
	return &synopsysv1.BlackduckSpec{
		Namespace:         "blackduck-externaldb",
		CertificateName:   "default",
		Size:              "small",
		PersistentStorage: false,
		ExternalPostgres: &synopsysv1.PostgresExternalDBConfig{
			PostgresHost:          "<<IP/FQDN>>",
			PostgresPort:          5432,
			PostgresAdmin:         "blackduck",
			PostgresUser:          "blackduck_user",
			PostgresSsl:           false,
			PostgresAdminPassword: "<<PASSWORD>>",
			PostgresUserPassword:  "<<PASSWORD>>",
		},
		Type:          "worker",
		ExposeService: NONE,
	}
}

// GetBlackDuckDefaultIPV6Disabled returns a BlackDuck with IPV6 Disabled
func GetBlackDuckDefaultIPV6Disabled() *synopsysv1.BlackduckSpec {
	return &synopsysv1.BlackduckSpec{
		Namespace:       "blackduck-ipv6disabled",
		CertificateName: "default",
		Environs: []string{
			"IPV4_ONLY:1",
			"BLACKDUCK_HUB_SERVER_ADDRESS:0.0.0.0",
		},
		Size:              "small",
		PersistentStorage: false,
		Type:              "worker",
		ExposeService:     NONE,
	}
}

// GetOpsSightUpstream returns the required fields for an upstream OpsSight
//func GetOpsSightUpstream() *synopsysv1.OpsSightSpec {
//	return &synopsysv1.OpsSightSpec{
//		Perceptor: &synopsysv1.Perceptor{
//			//Name:                           "perceptor",
//			//Port:                           3001,
//			//Image:                          "gcr.io/saas-hub-stg/blackducksoftware/perceptor:master",
//			CheckForStalledScansPauseHours: 999999,
//			StalledScanClientTimeoutHours:  999999,
//			ModelMetricsPauseSeconds:       15,
//			UnknownImagePauseMilliseconds:  15000,
//			ClientTimeoutMilliseconds:      100000,
//			Expose:                         NONE,
//		},
//		Perceiver: &synopsysv1.Perceiver{
//			EnableImagePerceiver: false,
//			EnablePodPerceiver:   true,
//			Port:                 3002,
//			ImagePerceiver: &opssightv1.ImagePerceiver{
//				Name:  "image-perceiver",
//				Image: "gcr.io/saas-hub-stg/blackducksoftware/image-perceiver:master",
//			},
//			PodPerceiver: &opssightv1.PodPerceiver{
//				Name:  "pod-perceiver",
//				Image: "gcr.io/saas-hub-stg/blackducksoftware/pod-perceiver:master",
//			},
//			ServiceAccount:            "perceiver",
//			AnnotationIntervalSeconds: 30,
//			DumpIntervalMinutes:       30,
//		},
//		ScannerPod: &opssightv1.ScannerPod{
//			Name: "perceptor-scanner",
//			ImageFacade: &opssightv1.ImageFacade{
//				Port:           3004,
//				Image:          "gcr.io/saas-hub-stg/blackducksoftware/perceptor-imagefacade:master",
//				ServiceAccount: "perceptor-scanner",
//				Name:           "perceptor-imagefacade",
//			},
//			Scanner: &opssightv1.Scanner{
//				Name:                 "perceptor-scanner",
//				Port:                 3003,
//				Image:                "gcr.io/saas-hub-stg/blackducksoftware/perceptor-scanner:master",
//				ClientTimeoutSeconds: 600,
//			},
//			ReplicaCount:   1,
//			ImageDirectory: "/var/images",
//		},
//		Prometheus: &opssightv1.Prometheus{
//			Name:   "prometheus",
//			Image:  "docker.io/prom/prometheus:v2.1.0",
//			Port:   9090,
//			Expose: NONE,
//		},
//		Skyfire: &opssightv1.Skyfire{
//			Image:                        "gcr.io/saas-hub-stg/blackducksoftware/pyfire:master",
//			Name:                         "skyfire",
//			Port:                         3005,
//			PrometheusPort:               3006,
//			ServiceAccount:               "skyfire",
//			HubClientTimeoutSeconds:      100,
//			HubDumpPauseSeconds:          240,
//			KubeDumpIntervalSeconds:      60,
//			PerceptorDumpIntervalSeconds: 60,
//		},
//		Blackduck: &opssightv1.Blackduck{
//			InitialCount:                       0,
//			MaxCount:                           0,
//			ConnectionsEnvironmentVariableName: "blackduck.json",
//			TLSVerification:                    false,
//			DeleteBlackduckThresholdPercentage: 50,
//			BlackduckSpec:                      GetBlackDuckTemplate(),
//		},
//		EnableMetrics: true,
//		EnableSkyfire: false,
//		DefaultCPU:    "300m",
//		DefaultMem:    "1300Mi",
//		ScannerCPU:    "300m",
//		ScannerMem:    "1300Mi",
//		LogLevel:      "debug",
//		SecretName:    "perceptor",
//		ConfigMapName: "opssight",
//	}
//}

// GetOpsSightDefault returns the required fields for OpsSight
//func GetOpsSightDefault() *opssightv1.OpsSightSpec {
//	return &opssightv1.OpsSightSpec{
//		Namespace: "opssight-test",
//		Perceptor: &opssightv1.Perceptor{
//			Name:                           "core",
//			Port:                           3001,
//			Image:                          "docker.io/blackducksoftware/opssight-core:2.2.4",
//			CheckForStalledScansPauseHours: 999999,
//			StalledScanClientTimeoutHours:  999999,
//			ModelMetricsPauseSeconds:       15,
//			UnknownImagePauseMilliseconds:  15000,
//			ClientTimeoutMilliseconds:      100000,
//			Expose:                         NONE,
//		},
//		ScannerPod: &opssightv1.ScannerPod{
//			Name: "scanner",
//			Scanner: &opssightv1.Scanner{
//				Name:                 "scanner",
//				Port:                 3003,
//				Image:                "docker.io/blackducksoftware/opssight-scanner:2.2.4",
//				ClientTimeoutSeconds: 600,
//			},
//			ImageFacade: &opssightv1.ImageFacade{
//				Name:               "image-getter",
//				Port:               3004,
//				InternalRegistries: []*opssightv1.RegistryAuth{},
//				Image:              "docker.io/blackducksoftware/opssight-image-getter:2.2.4",
//				ServiceAccount:     "scanner",
//				ImagePullerType:    "skopeo",
//			},
//			ReplicaCount: 1,
//		},
//		Perceiver: &opssightv1.Perceiver{
//			EnableImagePerceiver: false,
//			EnablePodPerceiver:   false,
//			Port:                 3002,
//			ImagePerceiver: &opssightv1.ImagePerceiver{
//				Name:  "image-processor",
//				Image: "docker.io/blackducksoftware/opssight-image-processor:2.2.4",
//			},
//			PodPerceiver: &opssightv1.PodPerceiver{
//				Name:  "pod-processor",
//				Image: "docker.io/blackducksoftware/opssight-pod-processor:2.2.4",
//			},
//			ServiceAccount:            "processor",
//			AnnotationIntervalSeconds: 30,
//			DumpIntervalMinutes:       30,
//		},
//		Prometheus: &opssightv1.Prometheus{
//			Name:   "prometheus",
//			Port:   9090,
//			Image:  "docker.io/prom/prometheus:v2.1.0",
//			Expose: NONE,
//		},
//		EnableSkyfire: false,
//		Skyfire: &opssightv1.Skyfire{
//			Image:                        "gcr.io/saas-hub-stg/blackducksoftware/pyfire:master",
//			Name:                         "skyfire",
//			Port:                         3005,
//			PrometheusPort:               3006,
//			ServiceAccount:               "skyfire",
//			HubClientTimeoutSeconds:      120,
//			HubDumpPauseSeconds:          240,
//			KubeDumpIntervalSeconds:      60,
//			PerceptorDumpIntervalSeconds: 60,
//		},
//		EnableMetrics: false,
//		DefaultCPU:    "300m",
//		DefaultMem:    "1300Mi",
//		ScannerCPU:    "300m",
//		ScannerMem:    "1300Mi",
//		LogLevel:      "debug",
//		SecretName:    "blackduck",
//		ConfigMapName: "opssight",
//		Blackduck: &opssightv1.Blackduck{
//			InitialCount:                       0,
//			MaxCount:                           0,
//			ConnectionsEnvironmentVariableName: "blackduck.json",
//			TLSVerification:                    false,
//			DeleteBlackduckThresholdPercentage: 50,
//			BlackduckSpec:                      GetBlackDuckTemplate(),
//		},
//	}
//}

// GetOpsSightDefaultWithIPV6DisabledBlackDuck retuns an OpsSight with a BlackDuck and
// IPV6 disabled
//func GetOpsSightDefaultWithIPV6DisabledBlackDuck() *opssightv1.OpsSightSpec {
//	return &opssightv1.OpsSightSpec{
//		Namespace: "opssight-test",
//		Perceptor: &opssightv1.Perceptor{
//			Name:                           "core",
//			Port:                           3001,
//			Image:                          "docker.io/blackducksoftware/opssight-core:2.2.4",
//			CheckForStalledScansPauseHours: 999999,
//			StalledScanClientTimeoutHours:  999999,
//			ModelMetricsPauseSeconds:       15,
//			UnknownImagePauseMilliseconds:  15000,
//			ClientTimeoutMilliseconds:      100000,
//			Expose:                         NONE,
//		},
//		ScannerPod: &opssightv1.ScannerPod{
//			Name: "scanner",
//			Scanner: &opssightv1.Scanner{
//				Name:                 "scanner",
//				Port:                 3003,
//				Image:                "docker.io/blackducksoftware/opssight-scanner:2.2.4",
//				ClientTimeoutSeconds: 600,
//			},
//			ImageFacade: &opssightv1.ImageFacade{
//				Name:               "image-getter",
//				Port:               3004,
//				InternalRegistries: []*opssightv1.RegistryAuth{},
//				Image:              "docker.io/blackducksoftware/opssight-image-getter:2.2.4",
//				ServiceAccount:     "scanner",
//				ImagePullerType:    "skopeo",
//			},
//			ReplicaCount: 1,
//		},
//		Perceiver: &opssightv1.Perceiver{
//			EnableImagePerceiver: false,
//			EnablePodPerceiver:   true,
//			ImagePerceiver: &opssightv1.ImagePerceiver{
//				Name:  "image-processor",
//				Image: "docker.io/blackducksoftware/opssight-image-processor:2.2.4",
//			},
//			PodPerceiver: &opssightv1.PodPerceiver{
//				Name:  "pod-processor",
//				Image: "docker.io/blackducksoftware/opssight-pod-processor:2.2.4",
//			},
//			ServiceAccount:            "opssight-processor",
//			AnnotationIntervalSeconds: 30,
//			DumpIntervalMinutes:       30,
//			Port:                      3002,
//		},
//		Prometheus: &opssightv1.Prometheus{
//			Name:   "prometheus",
//			Port:   9090,
//			Image:  "docker.io/prom/prometheus:v2.1.0",
//			Expose: NONE,
//		},
//		EnableSkyfire: false,
//		Skyfire: &opssightv1.Skyfire{
//			Image:                        "gcr.io/saas-hub-stg/blackducksoftware/pyfire:master",
//			Name:                         "skyfire",
//			Port:                         3005,
//			PrometheusPort:               3006,
//			ServiceAccount:               "skyfire",
//			HubClientTimeoutSeconds:      120,
//			HubDumpPauseSeconds:          240,
//			KubeDumpIntervalSeconds:      60,
//			PerceptorDumpIntervalSeconds: 60,
//		},
//		EnableMetrics: true,
//		DefaultCPU:    "300m",
//		DefaultMem:    "1300Mi",
//		ScannerCPU:    "300m",
//		ScannerMem:    "1300Mi",
//		LogLevel:      "debug",
//		SecretName:    "blackduck",
//		ConfigMapName: "opssight",
//		Blackduck: &opssightv1.Blackduck{
//			InitialCount:                       0,
//			MaxCount:                           0,
//			ConnectionsEnvironmentVariableName: "blackduck.json",
//			TLSVerification:                    false,
//			DeleteBlackduckThresholdPercentage: 50,
//			BlackduckSpec: &blackduckv1.BlackduckSpec{
//				PersistentStorage: false,
//				CertificateName:   "default",
//				Environs: []string{
//					"IPV4_ONLY:1",
//					"BLACKDUCK_HUB_SERVER_ADDRESS:0.0.0.0",
//				},
//				Size: "small",
//				Type: "worker",
//			},
//		},
//	}
//}

// GetAlertTemplate returns the required fields for Alert
func GetAlertTemplate() *synopsysv1.AlertSpec {
	return &synopsysv1.AlertSpec{}
}

// GetAlertDefault creates an Alert crd configuration object with defaults
func GetAlertDefault() *synopsysv1.AlertSpec {
	standAlone := false

	return &synopsysv1.AlertSpec{
		ExposeService:     NONE,
		Port:              IntToInt32(8443),
		PersistentStorage: false,
		PVCName:           "alert-pvc",
		StandAlone:        &standAlone,
		PVCSize:           "5G",
		AlertMemory:       "2560M",
		CfsslMemory:       "640M",
		Environs: []string{
			"ALERT_SERVER_PORT:8443",
			"PUBLIC_HUB_WEBSERVER_HOST:localhost",
			"PUBLIC_HUB_WEBSERVER_PORT:443",
		},
	}
}

// GetPolarisDBDefault returns PolarisDB default configuration
func GetPolarisDBDefault() *synopsysv1.PolarisDBSpec {
	return &synopsysv1.PolarisDBSpec{
		EnvironmentName:        "polaris",
		SMTPDetails:            synopsysv1.SMTPDetails{},
		PostgresInstanceType:   "internal",
		PostgresStorageDetails: synopsysv1.PostgresStorageDetails{},
		PostgresDetails: synopsysv1.PostgresDetails{
			Host: "postgresql",
			Port: 5432,
		},
		EventstoreDetails: synopsysv1.EventstoreDetails{
			Replicas: IntToInt32(3),
		},
		UploadServerDetails: synopsysv1.UploadServerDetails{
			Replicas: IntToInt32(1),
		},
	}
}
