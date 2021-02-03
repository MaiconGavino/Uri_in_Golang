line1 = input().split(" ")
line2 = input().split(" ")

codig1, qtd1, valor1 = line1
codig2, qtd2, valor2 = line2

total = (int(qtd1)*float(valor1)) + (int(qtd2)*float(valor2))

print("VALOR A PAGAR: R$ {:.2f}".format(total))