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

package soperator

import (
	"crypto/x509/pkix"
	"encoding/json"
	"fmt"
	"github.com/blackducksoftware/synopsys-operator/utils"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	synopsysv1 "github.com/blackducksoftware/synopsys-operator/api/v1"
	log "github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

//GetBlackduckVersionsToRemove finds all Blackducks with a different version, returns their specs with the new version
func GetBlackduckVersionsToRemove(restClient *rest.RESTClient, namespace string, newVersion string) ([]synopsysv1.Blackduck, error) {
	log.Debugf("Collecting all Blackducks that are not version: %s", newVersion)
	currBlackDucks, err := utils.ListBlackduck(restClient, namespace, v1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get BlackDucks: %s", err)
	}
	newBlackDucks := []synopsysv1.Blackduck{}
	for _, blackDuck := range currBlackDucks.Items {
		log.Debugf("Found Blackduck version '%s': %s", blackDuck.TypeMeta.APIVersion, blackDuck.Name)
		if blackDuck.TypeMeta.APIVersion != newVersion {
			blackDuck.TypeMeta.APIVersion = newVersion
			newBlackDucks = append(newBlackDucks, blackDuck)
		}
	}
	return newBlackDucks, nil
}

// GetOpsSightVersionsToRemove finds all OpsSights with a different version, returns their specs with the new version
func GetOpsSightVersionsToRemove(restClient *rest.RESTClient, namespace string, newVersion string) ([]synopsysv1.OpsSight, error) {
	log.Debugf("Collecting all OpsSights that are not version: %s", newVersion)
	currOpsSights, err := utils.ListOpsSight(restClient, namespace, v1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get OpsSights: %s", err)
	}
	newOpsSights := []synopsysv1.OpsSight{}
	for _, opsSight := range currOpsSights.Items {
		log.Debugf("Found OpsSight version '%s': %s", opsSight.TypeMeta.APIVersion, opsSight.Name)
		if opsSight.TypeMeta.APIVersion != newVersion {
			opsSight.TypeMeta.APIVersion = newVersion
			newOpsSights = append(newOpsSights, opsSight)
		}
	}
	return newOpsSights, nil
}

// GetAlertVersionsToRemove finds all Alerts with a different version, returns their specs with the new version
func GetAlertVersionsToRemove(restClient *rest.RESTClient, namespace string, newVersion string) ([]synopsysv1.Alert, error) {
	log.Debugf("Collecting all Alerts that are not version: %s", newVersion)
	currAlerts, err := utils.ListAlerts(restClient, namespace, v1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get Alerts: %s", err)
	}
	newAlerts := []synopsysv1.Alert{}
	for _, alert := range currAlerts.Items {
		log.Debugf("found Alert version '%s': %s", alert.TypeMeta.APIVersion, alert.Name)
		if alert.TypeMeta.APIVersion != newVersion {
			alert.TypeMeta.APIVersion = newVersion
			newAlerts = append(newAlerts, alert)
		}
	}
	return newAlerts, nil
}

// GetOperatorImage returns the image for Synopsys Operator from
// the cluster
func GetOperatorImage(kubeClient *kubernetes.Clientset, namespace string) (string, error) {
	currCM, err := utils.GetConfigMap(kubeClient, namespace, "synopsys-operator")
	if err != nil {
		return "", fmt.Errorf("unable to get synopsys operator image due to %s", err)
	}
	return currCM.Data["image"], nil
}

// GetOldOperatorSpec returns a spec that respesents the current Synopsys Operator in the cluster
func GetOldOperatorSpec(restConfig *rest.Config, kubeClient *kubernetes.Clientset, namespace string) (*SpecConfig, error) {
	log.Debugf("creating new Synopsys Operator spec")
	currCM, err := utils.GetConfigMap(kubeClient, namespace, "synopsys-operator")
	if err != nil {
		return nil, fmt.Errorf("unable to get synopsys operator config map in namespace %s due to %+v", namespace, err)
	}

	sOperatorSpec := SpecConfig{}
	sOperatorSpec.Namespace = namespace
	sOperatorSpec.RestConfig = restConfig
	sOperatorSpec.KubeClient = kubeClient

	data := currCM.Data["config.json"]
	err = json.Unmarshal([]byte(data), &sOperatorSpec)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal Synopsys Operator configMap data due to %+v", err)
	}

	var cmData map[string]interface{}
	err = json.Unmarshal([]byte(data), &cmData)
	if err != nil {
		log.Errorf("unable to unmarshal config map data due to %+v", err)
	}
	if crdNames, ok := cmData["CrdNames"]; ok {
		sOperatorSpec.Crds = utils.StringToStringSlice(crdNames.(string), ",")
	}

	// Set the secretType and secret data
	currSecret, err := utils.GetSecret(kubeClient, namespace, "blackduck-secret")
	if err != nil {
		return nil, fmt.Errorf("unable to get Synopsys Operator secret due to %+v", err)
	}
	currKubeSecretData := currSecret.Data
	sealKey := string(currKubeSecretData["SEAL_KEY"])
	if len(sealKey) == 0 {
		sealKey, err = utils.GetRandomString(32)
		if err != nil {
			log.Panicf("unable to generate the random string for SEAL_KEY due to %+v", err)
		}
	}
	sOperatorSpec.SealKey = sealKey

	// Set the secretType and secret data
	tlsSecret, err := utils.GetSecret(kubeClient, namespace, "synopsys-operator-tls")
	if err != nil {
		return nil, fmt.Errorf("unable to get Synopsys Operator tls secret due to %+v", err)
	}
	certificate := string(tlsSecret.Data["cert.crt"])
	certificateKey := string(tlsSecret.Data["cert.key"])
	if len(certificate) == 0 || len(certificateKey) == 0 {
		certificate, certificateKey, err = utils.GeneratePemSelfSignedCertificateAndKey(
			pkix.Name{CommonName: fmt.Sprintf("synopsys-operator.%s.svc", namespace)},
		)
		if err != nil {
			log.Panicf("unable to generate tls certificate and key due to %+v", err)
		}
	}
	sOperatorSpec.Certificate = certificate
	sOperatorSpec.CertificateKey = certificateKey
	log.Debugf("got current Synopsys Operator secret data from Cluster")

	return &sOperatorSpec, nil
}

// GetOldPrometheusSpec returns a spec that respesents the current prometheus in the cluster
func GetOldPrometheusSpec(restConfig *rest.Config, kubeClient *kubernetes.Clientset, namespace string) (*PrometheusSpecConfig, error) {
	log.Debugf("creating New Prometheus SpecConfig")
	prometheusSpec := PrometheusSpecConfig{}
	// Set Namespace
	prometheusSpec.Namespace = namespace
	// Set Image
	currCM, err := utils.GetConfigMap(kubeClient, namespace, "prometheus")
	if err != nil {
		return nil, fmt.Errorf("Failed to get Prometheus ConfigMap: %s", err)
	}
	prometheusSpec.Image = currCM.Data["Image"]
	prometheusSpec.Expose = currCM.Data["Expose"]
	prometheusSpec.RestConfig = restConfig
	prometheusSpec.KubeClient = kubeClient
	log.Debugf("added image %s to Prometheus SpecConfig", prometheusSpec.Image)

	return &prometheusSpec, nil

}

// GetClusterType returns the Cluster type. It defaults to Kubernetes
func GetClusterType(kubeClient *kubernetes.Clientset) ClusterType {
	if utils.IsOpenshift(kubeClient) {
		return OpenshiftClusterType
	}
	return KubernetesClusterType
}
