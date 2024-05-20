package inputs

import (
	"fmt"
	obs "github.com/openshift/cluster-logging-operator/api/observability/v1"
	internalobs "github.com/openshift/cluster-logging-operator/internal/api/observability"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ValidateAudit(spec obs.InputSpec) []metav1.Condition {
	if spec.Type != obs.InputTypeAudit {
		return nil
	}
	newCond := func(reason, message string, args ...any) metav1.Condition {
		if len(args) > 1 {
			message = fmt.Sprintf(message, args...)
		}
		return internalobs.NewCondition(obs.ValidationCondition,
			metav1.ConditionTrue,
			reason,
			message,
		)
	}

	if spec.Audit == nil {
		return []metav1.Condition{newCond(obs.ReasonMissingSpec, fmt.Sprintf("%s has nil audit spec", spec.Name))}
	}
	if len(spec.Audit.Sources) == 0 {
		return []metav1.Condition{newCond(obs.ReasonMissingSources, fmt.Sprintf("%s must define at least one valid source", spec.Name))}
	}
	return nil
}