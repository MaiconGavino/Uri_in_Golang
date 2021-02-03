line = input().split(" ")
entA, entB = line
entA = int(entA)
entB = int(entB)
if entA < entB:
    aux = entA
    entA = entB
    entB = aux
if entA % entB == 0:
    print("Sao Multiplos")
else:
    print("Nao sao Multiplos")