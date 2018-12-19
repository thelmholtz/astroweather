package weather

import(
    "log"
    "strconv"
    "errors"
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
        case isDrought(key):
            log.Print("It's drougth")
            return "sequia", nil
        case isOptimal(key):
            log.Print("It's optimal")
            return "optimal", nil
        default:
            log.Print("It's default")
            return "normal", nil
        }
}

func isRaining(day int) bool { return false }

func isDrought(day int) bool { return true }

func isOptimal(day int) bool { return false }

