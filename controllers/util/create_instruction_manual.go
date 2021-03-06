package util

import (
	"io/ioutil"
	"strings"

	flying_dutchman "github.com/blackducksoftware/synopsys-operator/flying-dutchman"
	"gopkg.in/yaml.v2"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
)

// CreateInstructionManual creates the instruction manual on the fly based on the group annotation and labels
func CreateInstructionManual(mapOfUniqueIDToDesiredRuntimeObject map[string]runtime.Object) (*flying_dutchman.RuntimeObjectDependencyYaml, error) {

	dependencyYamlStruct := &flying_dutchman.RuntimeObjectDependencyYaml{}

	accessor := meta.NewAccessor()

	for uniqueID, desiredRuntimeObject := range mapOfUniqueIDToDesiredRuntimeObject {
		labels, err := accessor.Labels(desiredRuntimeObject)
		if err != nil {
			return nil, err
		}

		annotations, err := accessor.Annotations(desiredRuntimeObject)
		if err != nil {
			return nil, err
		}

		group, ok := labels["operator.synopsys.com/group-id"]
		if !ok {
			continue
		}

		if dependencyYamlStruct.Groups == nil {
			dependencyYamlStruct.Groups = make(map[string][]string)
		}
		dependencyYamlStruct.Groups[group] = append(dependencyYamlStruct.Groups[group], uniqueID)

		dependencies, ok := annotations["operator.synopsys.com/group-dependencies"]
		if !ok {
			continue
		}

		var depIndex *int
		// Verify if the dependency group already exists and create it if it doesn't
		for valueNb, value := range dependencyYamlStruct.Dependencies {
			if strings.Compare(value.Obj, group) == 0 {
				depIndex = &valueNb
				break
			}
		}
		if depIndex == nil {
			dependencyYamlStruct.Dependencies = append(dependencyYamlStruct.Dependencies, flying_dutchman.RuntimeObjectDependency{
				Obj: group,
			})
			depIndex = func(i int) *int { return &i }(len(dependencyYamlStruct.Dependencies) - 1)
		}

		if len(dependencies) > 0 {
			for _, dependency := range strings.Split(dependencies, ",") {
				// Add dependencies
				isDependencyAlreadyPresent := false
				for _, dep := range dependencyYamlStruct.Dependencies[*depIndex].IsDependentOn {
					if strings.Compare(dep, dependency) == 0 {
						isDependencyAlreadyPresent = true
						break
					}
				}
				if !isDependencyAlreadyPresent {
					dependencyYamlStruct.Dependencies[*depIndex].IsDependentOn = append(dependencyYamlStruct.Dependencies[*depIndex].IsDependentOn, dependency)
				}
			}
		}

	}

	return dependencyYamlStruct, nil
}

// CreateInstructionManualFromYaml creates the instruction manual from yaml
func CreateInstructionManualFromYaml(instructionManualLocation string) (*flying_dutchman.RuntimeObjectDependencyYaml, error) {
	// Read Dependency YAML File into Struct
	dependencyYamlBytes, err := ioutil.ReadFile(instructionManualLocation)
	if err != nil {
		return nil, err
	}

	dependencyYamlStruct := &flying_dutchman.RuntimeObjectDependencyYaml{}

	err = yaml.Unmarshal(dependencyYamlBytes, dependencyYamlStruct)
	if err != nil {
		return nil, err
	}

	return dependencyYamlStruct, nil
}
