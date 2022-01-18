package cell

var cell bool

func InitCell() bool {
    cell = false; // må sette eksplisitt
    return cell
}

/*
   INN-DATA --> FUNKSJON --> UT-DATA
      ()        InitCell      etter InitCell()
*/

func SetCellValue(value bool) bool {
    cell = value; // endret verdi
    return cell;  // returnerte den endrede verdien
}

func GetCellValue() bool {
    return cell  // sjekker verdien på variabelen i pakke-avgrensningen
}
