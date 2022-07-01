package main

import "fmt"

/* OPERATORS:
--------------------
+	ADDITION
-	SUBTRACTION
*	MULTIPLICATION
/	DIVISION
%	MODULE
++	ADD 1
--	SUBTRACT 1
+=	ADD VALUE
-=	SUBTRACT VALUE
*=	MULTIPLY BY VALUE
/=	DIVIDE BY VALUE
%=	MODULE BY VALUE
--------------------
*/

func main() {
	var a, b int

	a = 2
	fmt.Println("a =", a)

	// ADD
	b = a + 3
	fmt.Println("b = a + 2 =", b)

	// SUB
	b = b - a
	fmt.Println("b = b - a =", b)

	// MULT
	b = b * 3
	fmt.Println("b = b * 3 =", b)

	// DIV
	b = b / a
	fmt.Println("b = b / a =", b)

	//MOD
	fmt.Println("b % a =", b%a)

	// ADD 1
	a++
	fmt.Println("a++ =", a)

	// SUBTRACT 1
	a--
	fmt.Println("a-- =", a)

	// ADD value
	b += a
	fmt.Println("b += a =", b)

	// SUB value
	b -= a
	fmt.Println("b -= a =", b)

	// MULT by value
	b *= a
	fmt.Println("b *= a =", b)

	// DIV by value
	b /= a
	fmt.Println("b /= a =", b)

	// MOD by value
	b %= a
	fmt.Println("b %= a =", b)

	// SCIENTIFIC NOTATION

	const avogadro = 6.02214076e23
	const waterMolarMass = 18. // gr/mol

	waterMass := waterMolarMass / avogadro // (gr/mol) * ( 1 mol / avogadro)
	println("A water (H2O) molecule has a mass of", waterMass, "gr")

	const amc = 1.660538e-27 // kg = 1 amu

	oxygenMass := 16 * amc
	println("The atomic mass of the oxygen is", oxygenMass, "kg")

	const c = 299_792_458
	println("The speed of light is", c, "m/s")

	E := oxygenMass * c * c // E=mc2
	println("Kinetic energy (E) in the Oxygen is", E, "Joules")

}
