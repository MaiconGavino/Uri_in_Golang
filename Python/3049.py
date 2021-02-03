baseA = int(input())
baseB = int(input())
base1 = 160 
base2 = 160 
altura = 70 
base1 = base1 - baseA
base2 = base2 - baseB
tp1 = ((baseA + baseB) * altura)/2
tp2 = ((base1 + base2) * altura)/2
if tp1 > tp2:
    print(1)
elif tp2 > tp1:
    print(2)
else:
    print(0)
