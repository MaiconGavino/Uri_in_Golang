entrada = int(input())
for i in range(0, entrada):
    line = input().split(" ")
    entA, entB = line
    entA = int(entA)
    entB = int(entB)
    aux = 0
    if entA > entB:
        aux = entA
        entA  = entB
        entB = aux
    while(entB % entA != 0):
        aux = entB % entA
        entB = entA
        entA = aux
    print(entA)