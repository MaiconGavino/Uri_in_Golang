aux = 0
while True:
    entrada = int(input())
    if entrada == -1:
        break
    aux += 1
    cont = 0
    while entrada > 1:
        cont += 1
        entrada -= 2
    print("Experiment {}: {} full cycle(s)".format(aux, cont))