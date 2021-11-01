package judge

import (
    "gamificacion/evaluate"
    "gamificacion/structures"
    "gamificacion/db"
)

var Body_values chan structures.Body
var State_values chan db.States

func init(){
    Body_values = make(chan structures.Body)
    State_values = make(chan db.States)
    go evaluate.All(Body_values, State_values)
}