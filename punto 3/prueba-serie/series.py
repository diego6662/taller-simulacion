import numpy as np
def main():
    file = open("../../punto 1/Number.txt").read()
    file = file.split(",")
    nums = []
    #se lee el array y se pasa su informacion de string a float
    for i in range(len(file) - 1):
        nums.append(float(file[i]))

    group_number = len(nums) / 2 ##numero de parejas
    
    class_number = 10.0 ## numero de clases
    
    FE = len(nums) / class_number
    inc = 10.0
    """
    incremento es el valor de como se debe dividir el intervalo
    que se da por la formula de techo(((n/2)^1/2)^1/2)
    """
    int_inf = 0
    int_sup = 1/inc
    incr = 1 / inc
    # si la cantidad de parejas es impar agrego el valor promedio de la secuencia
    # para que el ultimo numero de la misma tenga una pareja
    
    pairs = []
    fo = np.zeros((int(inc) ,int(inc) ))
    row = 0
    for i in range(int(group_number)):
        pairs.append((nums[i],nums[i + 1]))
    while int_sup <= 1:
        col = 0
        row_inf = 0
        row_sup = incr
        nums_temp = []
        for i in range(len(pairs)):
            if pairs[i][0] >= int_inf and pairs[i][0] < int_sup:
                nums_temp.append(pairs[i])
        
        while row_sup <= 1:
            for i in range(len(nums_temp)):
                if nums_temp[i][1] >= row_inf and nums_temp[i][1] < row_sup:
                    
                    fo[row,col] += 1
            row_inf = row_sup
            row_sup += incr
            col += 1
        row += 1
        int_inf = int_sup
        int_sup += incr
    table_chi = ((FE - fo) ** 2) / FE
    chi_cal = np.sum(table_chi)
    print("frecuencia obtenida:\n",fo)
    print("(fe-fo)²/fe\n",table_chi)
    print("chi calculado:",chi_cal)
    chi_cri = 36.42
    
    if chi_cal <= chi_cri:
        print("cumple la prueba de independencia")
    else:
        print("no cumple la prueba de independencia")



if  __name__ == "__main__":
    main()