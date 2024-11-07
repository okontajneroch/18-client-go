package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// príklad enumerácie
type Faction string

const (
	Rebellion Faction = "rebellion"
	Empire    Faction = "empire"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:generate=true
type Starfighter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StarfighterSpec   `json:"spec,omitempty"`
	Status            StarfighterStatus `json:"status,omitempty"`
}

// dátova štruktúra pre `spec`
type StarfighterSpec struct {
	Faction Faction `json:"faction"`
	Type    string  `json:"type"`
	Pilot   string  `json:"pilot"`
}

// dátova štruktúra pre `status`
type StarfighterStatus struct {
	Phases []string `json:"phases,omitempty"`
}
