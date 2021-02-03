caso = int(input())

for teste in range(caso):
    entrada = int(input())

    #verificar se eh primo
    primo = True
    for i in range(2, entrada):
        if entrada % i == 0:
            primo = False
            break

    #imprimir o resultado

    if primo:
        print('{} eh primo'.format(entrada))
    else:
        print('{} nao eh primo'.format(entrada))
