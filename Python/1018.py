valor = int(input())
notas = [100, 50, 20, 10, 5, 2, 1]
print(valor)
for nota in notas:
    qtdNotas = int(valor / nota)
    print("{} nota(s) de R$ {:.2f}".format(qtdNotas, nota))
    valor -= qtdNotas*nota
