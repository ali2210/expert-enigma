package quanta

import (
	"reflect"

	"github.com/itsubaki/q"
	"github.com/itsubaki/q/pkg/quantum/qubit"
)

var bob_state []qubit.State
var alice_state []qubit.State

func QuantaCode() ([]qubit.State, []qubit.State) {

	// get qubit systems
	qmac := q.New()

	// declare qubit as Zero
	bob := qmac.Zero()

	// Pauli Gate (cnot) is controlled qubit that changed when qubit other than Zero
	// convert Zero qubit into superposition
	qmac.H(bob).CNOT(bob, qmac.One())

	// then measure qubit with qubit as zero
	qmac.Measure(bob, qmac.Zero())

	// defined another qubit systems
	aqmac := q.New()

	// this qubit declare as One
	alice := aqmac.One()

	// create opposite effect over qubit One and then measure wuth qubit One
	aqmac.H(alice).CNOT(aqmac.Zero(), bob)
	aqmac.Measure(bob, aqmac.One())

	// get both qubits states
	bob_state, alice_state = qmac.State(), aqmac.State()

	return bob_state, alice_state

}

func QuantaValid(p, q []string) bool {

	// there may be more than one probable that both are different states
	if reflect.DeepEqual(p, bob_state[0].BinaryString) && reflect.DeepEqual(q, alice_state[0].BinaryString) {
		return true
	}
	return false
}
