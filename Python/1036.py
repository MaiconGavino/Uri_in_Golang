line = input().split(" ")
entA, entB, entC = line
entA = float(entA)
entB = float(entB)
entC = float(entC)
delta = (entB**2 - 4 * entA * entC)
if delta < 0 or float(entA) == 0:
        print("Impossivel calcular")
else:
    x1 = (-float(entB) + (delta)**(1/2))/(2*entA)
    x2 =  (-float(entB) - (delta)**(1/2))/(2*entA)
    print("R1 = {:.5f}".format((x1)))
    print("R2 = {:.5f}".format((x2)))