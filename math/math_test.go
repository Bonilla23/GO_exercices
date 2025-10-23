// math_test.go
package math
// Paquete para hacer el test
import "testing"

// testsuma( parametro que el paquete testing proporciona )
func TestSuma(t *testing.T) {
    // resultado = 6
    resultado := Suma(2, 3)
    // esperado = 5
    esperado := 5
    // 6 no es 5 = Error ( se esperaba 5 )
    if resultado != esperado {
        t.Errorf("Suma(2, 3) = %d; se esperaba %d", resultado, esperado)
    }
}

func TestSumaIncorrecta(t *testing.T) {
    // resultado = 6
    resultado := Suma(3, 3)
    // esperado = 5
    esperado := 6
    // 6 no es 5 = Error ( se esperaba 5 )
    if resultado != esperado {
        t.Errorf("Suma(2, 3) = %d; se esperaba %d", resultado, esperado)
    }
}
