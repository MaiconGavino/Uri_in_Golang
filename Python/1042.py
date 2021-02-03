line = input().split(" ")
entA, entB, entC = line
entA = int(entA)
entB = int(entB)
entC = int(entC)

if entA > entB and entA > entC:
    z = entA
    if entB > entC:
        x = entC
        y = entB
    else:
        x = entB
        y = entC
if entB > entA and entB > entC:
    z = entB
    if entA > entC:
        x = entC
        y = entA
    else:
        x = entA
        y = entC
if entC > entA and entC > entB:
    z = entC
    if entA > entB:
        x = entB
        y = entA
    else:
        x = entA
        y = entB


print("{}".format(x))
print("{}".format(y))
print("{}".format(z))
print()
print("{}".format(int(entA)))
print("{}".format(int(entB)))
print("{}".format(int(entC)))