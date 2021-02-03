import math
def Prime(entrada):
    if entrada % 2 == 0 and entrada > 2 or entrada < 2:
        return False
    return all(entrada % i for i in range(3, int(math.sqrt(entrada))+1, 2))
casos = int(input())
while(casos > 0):
    casos -= 1
    entrada = int(input())
    if (Prime(entrada)):
        print("Prime")
    else:
        print("Not Prime")