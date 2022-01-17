# Cellemodell (fra konseptuell modell på norsk til konseptuell modell på Golang)
Simulering av en digital celle.

Funksjoner: 
- lage en celle (og eventuelt initialisere den)
- sette en verdi i celle (tillatte verdier 0 eller 1)
- avlese den verdien som cellen har (tilstand)

# Lage en celle
Hvordan allokerer vi minne for nye typer i vårt programmeringsspråk?
Kan vurdere å bruke bool for representere cellen.
https://go.dev/ref/spec#Boolean_types 
TestInitCell

# Sette en verdi 
TestSetCellValue

# Avlese verdien i cellen
TestViewCellStatus
TestGetCellValue