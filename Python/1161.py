while True:
    try:
        line = input().split(" ")
        entA, entB = line
        entA = int(entA)
        entB = int(entB)
        aux = entA - 1
        aux2 = entB - 1
        if entA == 0:
            entA = 1
        if entB == 0:
            entB = 1
        
        while aux > 0:
            entA = entA * aux
            aux -= 1
        while aux2 > 0:
            entB = entB * aux
            aux -= 1
        print(entA + entB)
    except:
        break