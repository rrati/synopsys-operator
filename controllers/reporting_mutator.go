package controllers

import (
	"fmt"

	synopsysv1 "github.com/blackducksoftware/synopsys-operator/api/v1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	appsv1 "k8s.io/api/extensions/v1beta1"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func patchReporting(client client.Client, reportingCr *synopsysv1.Reporting, mapOfUniqueIdToBaseRuntimeObject map[string]runtime.Object) map[string]runtime.Object {
	patcher := ReportingPatcher{
		reportingCr:                      reportingCr,
		mapOfUniqueIdToBaseRuntimeObject: mapOfUniqueIdToBaseRuntimeObject,
		Client:                           client,
	}
	return patcher.patch()
}

type ReportingPatcher struct {
	reportingCr                      *synopsysv1.Reporting
	mapOfUniqueIdToBaseRuntimeObject map[string]runtime.Object
	client.Client
}

func (p *ReportingPatcher) patch() map[string]runtime.Object {
	patches := []func() error{
		p.patchNamespace,
		p.patchReportStorageSpec,
		p.patchReportingFrontendSpec,
		p.patchReportingIssueManagerSpec,
		p.patchReportingPortfolioServiceSpec,
		p.patchReportingReportServiceSpec,
		p.patchReportingToolsPortfolioServiceSpec,
		p.patchReportingSwaggerDoc,
	}
	for _, f := range patches {
		err := f()
		if err != nil {
			fmt.Printf("%s\n", err)
		}
	}
	return p.mapOfUniqueIdToBaseRuntimeObject
}

func (p *ReportingPatcher) patchNamespace() error {
	accessor := meta.NewAccessor()
	for _, runtimeObject := range p.mapOfUniqueIdToBaseRuntimeObject {
		accessor.SetNamespace(runtimeObject, p.reportingCr.Spec.Namespace)
	}
	return nil
}

func (p *ReportingPatcher) patchReportStorageSpec() error {
	ReportStorageUniqueID := "PersistentVolumeClaim." + "report-storage-pv-claim"
	reportStorageRuntimeObject, ok := p.mapOfUniqueIdToBaseRuntimeObject[ReportStorageUniqueID]
	if !ok {
		return nil
	}
	reportStorageInstance := reportStorageRuntimeObject.(*corev1.PersistentVolumeClaim)
	if size, err := resource.ParseQuantity(p.reportingCr.Spec.ReportStorageSpec.Volume.Size); err == nil {
		reportStorageInstance.Spec.Resources.Requests[v1.ResourceStorage] = size
	}
	return nil
}

func (p *ReportingPatcher) patchReportingFrontendSpec() error {
	if p.reportingCr.Spec.ReportingFrontendSpec.ImageDetails != nil {
		DeploymentUniqueID := "Deployment.rp-frontend"
		deploymentRuntimeObject, ok := p.mapOfUniqueIdToBaseRuntimeObject[DeploymentUniqueID]
		if !ok {
			return nil
		}
		PatchImageForService(
			p.reportingCr.Spec.ReportingFrontendSpec.ImageDetails,
			deploymentRuntimeObject.(*appsv1.Deployment),
		)
	}
	return nil
}

func (p *ReportingPatcher) patchReportingIssueManagerSpec() error {
	if p.reportingCr.Spec.ReportingIssueManagerSpec.ImageDetails != nil {
		DeploymentUniqueID := "Deployment.rp-issue-manager"
		deploymentRuntimeObject, ok := p.mapOfUniqueIdToBaseRuntimeObject[DeploymentUniqueID]
		if !ok {
			return nil
		}
		PatchImageForService(
			p.reportingCr.Spec.ReportingIssueManagerSpec.ImageDetails,
			deploymentRuntimeObject.(*appsv1.Deployment),
		)
	}
	return nil
}

func (p *ReportingPatcher) patchReportingPortfolioServiceSpec() error {
	if p.reportingCr.Spec.ReportingPortfolioServiceSpec.ImageDetails != nil {
		DeploymentUniqueID := "Deployment.rp-portfolio-service"
		deploymentRuntimeObject, ok := p.mapOfUniqueIdToBaseRuntimeObject[DeploymentUniqueID]
		if !ok {
			return nil
		}
		PatchImageForService(
			p.reportingCr.Spec.ReportingPortfolioServiceSpec.ImageDetails,
			deploymentRuntimeObject.(*appsv1.Deployment),
		)
	}
	return nil
}

func (p *ReportingPatcher) patchReportingReportServiceSpec() error {
	if p.reportingCr.Spec.ReportingReportServiceSpec.ImageDetails != nil {
		DeploymentUniqueID := "Deployment.rp-report-service"
		deploymentRuntimeObject, ok := p.mapOfUniqueIdToBaseRuntimeObject[DeploymentUniqueID]
		if !ok {
			return nil
		}
		PatchImageForService(
			p.reportingCr.Spec.ReportingReportServiceSpec.ImageDetails,
			deploymentRuntimeObject.(*appsv1.Deployment),
		)
	}
	return nil
}

func (p *ReportingPatcher) patchReportingToolsPortfolioServiceSpec() error {
	if p.reportingCr.Spec.ReportingToolsPortfolioServiceSpec.ImageDetails != nil {
		DeploymentUniqueID := "Deployment.rp-tools-portfolio-service"
		deploymentRuntimeObject, ok := p.mapOfUniqueIdToBaseRuntimeObject[DeploymentUniqueID]
		if !ok {
			return nil
		}
		PatchImageForService(
			p.reportingCr.Spec.ReportingToolsPortfolioServiceSpec.ImageDetails,
			deploymentRuntimeObject.(*appsv1.Deployment),
		)
	}
	return nil
}

func (p *ReportingPatcher) patchReportingSwaggerDoc() error {
	if p.reportingCr.Spec.ReportingSwaggerDoc.ImageDetails != nil {
		DeploymentUniqueID := "Deployment.rp-swagger-doc"
		deploymentRuntimeObject, ok := p.mapOfUniqueIdToBaseRuntimeObject[DeploymentUniqueID]
		if !ok {
			return nil
		}
		PatchImageForService(
			p.reportingCr.Spec.ReportingSwaggerDoc.ImageDetails,
			deploymentRuntimeObject.(*appsv1.Deployment),
		)
	}
	return nil
}
