#Entradas
line = input().split(" ")
entA, entB, entC = line
entA = float(entA)
entB = float(entB)
entC = float(entC)

#Operações
triangulo = (entA * entC)/2
areaCirculo = 3.14159 * (entC**2)
areaTrapezio = ((entA+entB)*entC)/2
areaQuadrado = entB**2
areaRetangulo = entA * entB

#saidas
print("TRIANGULO: {:.3f}".format(triangulo))
print("CIRCULO: {:.3f}".format(areaCirculo))
print("TRAPEZIO: {:.3f}".format(areaTrapezio))
print("QUADRADO: {:.3f}".format(areaQuadrado))
print("RETANGULO: {:.3f}".format(areaRetangulo))