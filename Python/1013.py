line = input().split(" ")

entA, entB, entC = line

maiorAB = (int(entA) + int(entB) + abs(int(entA)-int(entB)))/2
maiorABC = (int(maiorAB)+int(entC)+ abs(int(maiorAB)-int(entC)))/2
print("{} eh o maior".format(int(maiorABC)))