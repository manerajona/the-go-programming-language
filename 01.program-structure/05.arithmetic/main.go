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
--	SUBSTRACT 1
+=	ADD VALUE
-=	SUBSTRACT VALUE
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

	// ++
	a++
	fmt.Println("a++ =", a)

	// --
	a--
	fmt.Println("a-- =", a)

	// ADD
	b += a
	fmt.Println("b += a =", b)

	// SUB
	b -= a
	fmt.Println("b -= a =", b)

	// MULT
	b *= a
	fmt.Println("b *= a =", b)

	// DIV
	b /= a
	fmt.Println("b /= a =", b)

	// MOD
	b %= a
	fmt.Println("b %= a =", b)

	// SCIENTIFIC NOTATION

	const avogadro = 6.02214076e23

	twoMole := 2 * avogadro
	println("The number of particles in two Mole are", twoMole)

	const amc = 1.660538e-27 // kg = 1 amu

	oxigenMass := 16 * amc
	println("The atomic mass of the oxigen is", oxigenMass, "kg")

	const c = 299_792_458
	println("The speed of light is", c, "m/s")

	E := oxigenMass * c * c // E=mc2
	println("Kinetic energy (E) in the Oxigen is", E, "Joules")

}
