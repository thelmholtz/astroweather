package weather

import(
    "log"
    "strconv"
    "errors"
    "github.com/thelmholtz/astroweather/planets"
)

func Day(day string) (string, error) {

    log.Print("Queried weather for day: " + day)

    key, err := strconv.Atoi(day)

    if err != nil {
        log.Print("ERROR: El dia no es numerico")
        return "undefined", err
    }

    switch {
        case key < 0:
            log.Print("ERROR: El dia tiene que ser positivo")
            return "undefined", errors.New("El dia no es positivo")
        case isRaining(key):
            log.Print("It's raining")
            return "lluvia", nil
        case isOptimal(key):
            log.Print("It's optimal")
            return "optimal", nil
        case isDry(key):
            log.Print("It's dry")
            return "sequia", nil
        default:
            log.Print("It's normal")
            return "normal", nil
        }
}

//TODO Not sure this funcs are necesary...

func isRaining(day int) bool {
    return false
}

func isDry(day int) bool {
    return planets.RadiallyAligned(day)
}


func isOptimal(day int) bool {
    return planets.Aligned(day)
}

