package template

import (
	_ "embed"
	"fmt"
	"regexp"
	"strings"

	"text/template"

	"github.com/openshift/cluster-logging-operator/internal/generator/framework"
	"github.com/openshift/cluster-logging-operator/internal/generator/vector/elements"
	vectorhelpers "github.com/openshift/cluster-logging-operator/internal/generator/vector/helpers"
)

var (
	//go:embed template.vrl.tmpl
	templateVRLTmplStr   string
	pathRegex            = regexp.MustCompile(`\{([^{}]+)\}`)
	splitRegex           = regexp.MustCompile(`to_string!\(([^)]+)\)`)
	GroupNameVRLTemplate = template.Must(template.New("template VRL").Parse(templateVRLTmplStr))
)

type Template struct {
	Field     string
	VRLString string
}

func TemplateRemap(componentID string, inputs []string, userTemplate, field, description string) framework.Element {
	// Generate template
	w := &strings.Builder{}

	_ = GroupNameVRLTemplate.Execute(w,
		Template{
			Field:     field,
			VRLString: TransformUserTemplateToVRL(userTemplate),
		},
	)

	return elements.Remap{
		Desc:        description,
		ComponentID: componentID,
		Inputs:      vectorhelpers.MakeInputs(inputs...),
		VRL:         w.String(),
	}
}

// TransformUserTemplateToVRL converts the user entered template to VRL compatible syntax
// Example: foo-{.log_type||"none"} -> "foo-" + to_string!(.log_type||"none")
func TransformUserTemplateToVRL(groupName string) string {
	// Finds and replaces expressions defined in `{}` with to_string!()
	replacedGroupName := pathRegex.ReplaceAllStringFunc(groupName, func(match string) string {
		matches := pathRegex.FindStringSubmatch(match)
		replaced := fmt.Sprintf("to_string!(%s)", matches[1])
		return replaced
	})

	// Finding all matches of to_string!() returning their start + end indices
	matchedIndices := splitRegex.FindAllStringSubmatchIndex(replacedGroupName, -1)
	if len(matchedIndices) == 0 {
		return fmt.Sprintf("%q", groupName)
	}

	var result []string
	lastIndex := 0
	// Make the final resulting array with the appropriate pieces so that it can be concatenated together with `+`
	for _, match := range matchedIndices {
		// Append the part before the match. Check if empty string so we don't concat it
		partBeforeMatch := replacedGroupName[lastIndex:match[0]]
		if partBeforeMatch != "" {
			result = append(result, fmt.Sprintf("%q", partBeforeMatch))
		}

		// Append the to_string!() group
		result = append(result, replacedGroupName[match[0]:match[1]])
		lastIndex = match[1]
	}
	// Append the remaining part of the string after the last match making sure it isn't the empty string
	endOfString := replacedGroupName[lastIndex:]
	if endOfString != "" {
		result = append(result, fmt.Sprintf("%q", endOfString))
	}

	// Join array with `+`
	return strings.Join(result, " + ")
}
