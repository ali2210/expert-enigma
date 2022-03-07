package quanta

import (
	"reflect"

	"github.com/itsubaki/q"
	"github.com/itsubaki/q/pkg/quantum/qubit"
)

var bob_state []qubit.State
var alice_state []qubit.State

func QuantaCode() ([]qubit.State, []qubit.State) {

	qmac := q.New()
	bob := qmac.Zero()

	qmac.H(bob).CNOT(bob, qmac.One())
	qmac.Measure(bob, qmac.Zero())

	aqmac := q.New()
	alice := aqmac.One()

	aqmac.H(alice).CNOT(aqmac.Zero(), bob)
	aqmac.Measure(bob, aqmac.One())

	bob_state, alice_state = qmac.State(), aqmac.State()

	return bob_state, alice_state

}

func QuantaValid(p, q []string) bool {
	if reflect.DeepEqual(p, bob_state[0].BinaryString) && reflect.DeepEqual(q, alice_state[0].BinaryString) {
		return true
	}
	return false
}
