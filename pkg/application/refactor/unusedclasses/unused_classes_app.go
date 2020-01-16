package unusedclasses

import (
	"github.com/phodal/coca/pkg/domain/jdomain"
	"sort"
	"strings"
)

var analysisPackage = ""

func Refactoring(parsedDeps []jdomain.JClassNode) []string {
	sourceClasses := make(map[string]string)
	targetClasses := make(map[string]string)

	for _, node := range parsedDeps {
		if strings.Contains(node.Package, analysisPackage) {
			className := node.Package + "." + node.Class
			sourceClasses[className] = className
		}

		for _, method := range node.Methods {
			for _, methodCall := range method.MethodCalls {
				if strings.Contains(methodCall.Package, analysisPackage) {
					className := methodCall.Package + "." + methodCall.Class
					targetClasses[className] = className
				}
			}
		}
	}

	var excludePackage []string = nil
	for _, clz := range sourceClasses {
		if targetClasses[clz] != clz {
			excludePackage = append(excludePackage, clz)
		}
	}

	sort.Strings(excludePackage)
	return excludePackage
}
