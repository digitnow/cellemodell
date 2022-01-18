package cell

import "testing"

func TestInitCell(t *testing.T) {
    wanted := false; // Bruker semantikken i Golangs konseptuelle modell
    state := InitCell()
    if state != wanted {
        t.Errorf("Feil, fikk %t, ønsket %t.", state, wanted)
    }
}

func TestSetCellValue(t *testing.T) {
    wanted := true
    state := SetCellValue(true)
    if state != wanted {
        t.Errorf("Feil, fikk %t, ønsket %t.", state, wanted)
    }
}

// Ikke en enhetstest strengt tatt?
func TestGetCellValue(t *testing.T) {
    wanted := true
    //InitCell()
    SetCellValue(true)
    state := GetCellValue()
    if state != wanted {
        t.Errorf("Feil, fikk %t, ønsket %t.", state, wanted)
    }
}
