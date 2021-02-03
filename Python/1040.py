line = input().split(" ")
entA, entB, entC, entD = line
entA = float(entA)
entB = float(entB)
entC = float(entC)
entD = float(entD)
media = (entA * 2 + entB * 3 + entC * 4 + entD)/10
if media >= 7:
    print("Media: {:.1f}".format(float(media)))
    print("Aluno aprovado.")
elif media >= 5 and media < 7:
    exame = float(input())
    resultExame = (exame + media)/2
    if resultExame >= 5:
        print("Media: {:.1f}".format(float(media)))
        print("Aluno em exame.")
        print("Nota do exame: {:.1f}".format(exame))
        print("Aluno aprovado.")
        print("Media final: {:.1f}".format(resultExame))
    else:
        print("Media: {:.1f}".format(float(media)))
        print("Aluno em exame.")
        print("Nota do exame: {:.1f}".format(exame))
        print("Aluno reprovado.")
        print("Media final: {:.1f}".format(resultExame))
else:
    print("Media: {:.1f}".format(float(media)))
    print("Aluno reprovado.")
    
