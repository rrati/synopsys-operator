/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"fmt"
	"strings"

	synopsysv1 "github.com/blackducksoftware/synopsys-operator/meta-builder/api/v1"
	controllers_utils "github.com/blackducksoftware/synopsys-operator/meta-builder/controllers/util"
	flying_dutchman "github.com/blackducksoftware/synopsys-operator/meta-builder/flying-dutchman"
	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// AlertReconciler reconciles a Alert object
type AlertReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	Log    logr.Logger
}

var (
	apiGVStr = synopsysv1.GroupVersion.String()
)

func (r *AlertReconciler) GetClient() client.Client {
	return r.Client
}

func (r *AlertReconciler) GetCustomResource(req ctrl.Request) (metav1.Object, error) {
	var alert synopsysv1.Alert
	if err := r.Get(context.Background(), req.NamespacedName, &alert); err != nil {
		//log.Error(err, "unable to fetch Alert")
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		// TODO:
		// We generally want to ignore (not requeue) NotFound errors, since we’ll get a reconciliation request once the object exists, and requeuing in the meantime won’t help.
		if !apierrs.IsNotFound(err) {
			return nil, err
		}
		if apierrs.IsNotFound(err) {
			return nil, nil
		}
	}
	return &alert, nil
}

func (r *AlertReconciler) GetRuntimeObjects(cr interface{}) (map[string]runtime.Object, error) {
	alertCr := cr.(*synopsysv1.Alert)
	//TODO: either read contents of yaml from locally mounted file

	// only fetch the location of the latest if the version in the spec is not given
	var version string
	if 0 == len(alertCr.Spec.Version) {
		latestUrl := "https://raw.githubusercontent.com/blackducksoftware/releases/master/alert/latest"
		if latestArrayOfByte, err := controllers_utils.HTTPGet(latestUrl); err != nil {
			// TODO: error getting latest
			return nil, err
		} else {
			version = string(latestArrayOfByte)
		}
	} else {
		version = alertCr.Spec.Version
	}

	latestBaseYamlUrl := fmt.Sprintf("https://raw.githubusercontent.com/blackducksoftware/releases/master/alert/%s/alert_base.yaml", version)
	latestBaseYamlAsByteArray, err := controllers_utils.HTTPGet(latestBaseYamlUrl)
	if err != nil {
		return nil, err
	}

	//FinalYamlPath := "config/samples/alert_runtime_objects.yaml"
	//byteArrayContentFromFile, err := ioutil.ReadFile(FinalYamlPath)
	//if err != nil {
	//	return nil, err
	//}

	latestBaseYamlAsString := string(latestBaseYamlAsByteArray)
	latestBaseYamlAsString = strings.ReplaceAll(latestBaseYamlAsString, "${NAME}", alertCr.Name)
	latestBaseYamlAsString = strings.ReplaceAll(latestBaseYamlAsString, "${NAMESPACE}", alertCr.Spec.Namespace)
	if len(alertCr.Spec.ExposeService) > 0 {
		latestBaseYamlAsString = strings.ReplaceAll(latestBaseYamlAsString, "${SERVICE_TYPE}", alertCr.Spec.ExposeService)
	} else {
		latestBaseYamlAsString = strings.ReplaceAll(latestBaseYamlAsString, "${SERVICE_TYPE}", string(corev1.ServiceTypeClusterIP))
	}

	if len(alertCr.Spec.AlertMemory) > 0 {
		latestBaseYamlAsString = strings.ReplaceAll(latestBaseYamlAsString, "${ALERT_MEM}", alertCr.Spec.AlertMemory)
	} else {
		latestBaseYamlAsString = strings.ReplaceAll(latestBaseYamlAsString, "${ALERT_MEM}", "2560M")
	}

	if len(alertCr.Spec.CfsslMemory) > 0 {
		latestBaseYamlAsString = strings.ReplaceAll(latestBaseYamlAsString, "${CFSSL_MEM}", alertCr.Spec.CfsslMemory)
	} else {
		latestBaseYamlAsString = strings.ReplaceAll(latestBaseYamlAsString, "${CFSSL_MEM}", "640M")
	}

	if 0 != *alertCr.Spec.Port {
		latestBaseYamlAsString = strings.ReplaceAll(latestBaseYamlAsString, "${ALERT_PORT}", string(*alertCr.Spec.Port))
	} else {
		latestBaseYamlAsString = strings.ReplaceAll(latestBaseYamlAsString, "${ALERT_PORT}", "8443")
	}

	mapOfUniqueIdToBaseRuntimeObject := controllers_utils.ConvertYamlFileToRuntimeObjects(latestBaseYamlAsString)
	for _, desiredRuntimeObject := range mapOfUniqueIdToBaseRuntimeObject {
		// set an owner reference
		if err := ctrl.SetControllerReference(alertCr, desiredRuntimeObject.(metav1.Object), r.Scheme); err != nil {
			// requeue if we cannot set owner on the object
			// TODO: change this to requeue, and only not requeue when we get "newAlreadyOwnedError", i.e: if it's already owned by our CR
			//return ctrl.Result{}, err
			return mapOfUniqueIdToBaseRuntimeObject, nil
		}
	}
	mapOfUniqueIdToDesiredRuntimeObject := patchAlert(alertCr, mapOfUniqueIdToBaseRuntimeObject)

	return mapOfUniqueIdToDesiredRuntimeObject, nil
}

func (r *AlertReconciler) GetInstructionManual(mapOfUniqueIdToDesiredRuntimeObject map[string]runtime.Object) (*flying_dutchman.RuntimeObjectDependencyYaml, error) {
	//instructionManualFile := "https://raw.githubusercontent.com/yashbhutwala/kb-synopsys-operator/master/config/samples/dependency_sample_alert.yaml"
	instructionManual, err := controllers_utils.CreateInstructionManual(mapOfUniqueIdToDesiredRuntimeObject)
	if err != nil {
		return nil, err
	}
	return instructionManual, nil
}

func (r *AlertReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	//ctx := context.Background()
	//log := r.Log.WithValues("name and namespace of the alertCr to reconcile", req.NamespacedName)

	// TODO: By adding sha, we no longer need to requeue after (awesome!!), but it's here just in case you need to re-enable it
	//return ctrl.Result{RequeueAfter: 10 * time.Minute}, nil
	return flying_dutchman.MetaReconcile(req, r)
}

// +kubebuilder:rbac:groups=alerts.synopsys.com,resources=alerts,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=alerts.synopsys.com,resources=alerts/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=alerts,resources=pods,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=alerts,resources=pods/status,verbs=get
// +kubebuilder:rbac:groups=alerts,resources=namespaces,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=alerts,resources=services,verbs=get;list;watch;create;update;patch;delete

func (r *AlertReconciler) SetIndexingForChildrenObjects(mgr ctrl.Manager, ro runtime.Object) error {
	if err := mgr.GetFieldIndexer().IndexField(ro, flying_dutchman.JobOwnerKey, func(rawObj runtime.Object) []string {
		// grab the job object, extract the owner...
		owner := metav1.GetControllerOf(ro.(metav1.Object))
		if owner == nil {
			return nil
		}
		// ...make sure it's a Alert...
		if owner.APIVersion != apiGVStr || owner.Kind != "Alert" {
			return nil
		}

		// ...and if so, return it
		return []string{owner.Name}
	}); err != nil {
		return err
	}
	return nil
}

func (r *AlertReconciler) SetupWithManager(mgr ctrl.Manager) error {
	// Code here allows to kick off a reconciliation when runtime objects our controller manages are changed somehow
	r.SetIndexingForChildrenObjects(mgr, &corev1.ConfigMap{})
	r.SetIndexingForChildrenObjects(mgr, &corev1.Service{})
	r.SetIndexingForChildrenObjects(mgr, &corev1.ReplicationController{})
	r.SetIndexingForChildrenObjects(mgr, &corev1.Secret{})

	alertBuilder := ctrl.NewControllerManagedBy(mgr).For(&synopsysv1.Alert{}).Named("alert")
	alertBuilder = alertBuilder.Owns(&corev1.ConfigMap{})
	alertBuilder = alertBuilder.Owns(&corev1.Service{})
	alertBuilder = alertBuilder.Owns(&corev1.ReplicationController{})
	alertBuilder = alertBuilder.Owns(&corev1.Secret{})

	return alertBuilder.Complete(r)
}
