casos = int(input())
for i in range(0, casos):
    line = input().split(" ")
    entA, entB = line
    entA = int(entA)
    entB = int(entB)
    result = (entA * entB)/2
    print("{} cm2".format(int(result)))