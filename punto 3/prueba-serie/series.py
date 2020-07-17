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
    file = open("Table.txt","w+")
    file.write("frecuencia obtenida:\n") 
    file.write("           |0.0-0.1 | 0.1-0.2 | 0.2-0.3 | 0.3-0.4 | 0.4-0.5 | 0.5-0.6 | 0.6-0.7 | 0.7-0.8 | 0.8-0.9 | 0.9-1.0|\n")
    int_inf = 0.0
    int_sup = 1/inc
    incr = 1 / inc
   
    for i in range(10):
        
        text_temp = f"{str(int_inf)[:3]} - {str(int_sup)[:3]}  |   {int(fo[i,0])}  |   {int(fo[i,1])}   |   {int(fo[i,2])}   |   {int(fo[i,3])}   |   {int(fo[i,4])}   |   {int(fo[i,5])}   |   {int(fo[i,6])}   |   {int(fo[i,7])}   |   {int(fo[i,8])}   |   {int(fo[i,9])}  |"
        file.write(text_temp + "\n")
        int_inf = int_sup
        int_sup += incr
        
    file.write("(fe-fo)²/fe:\n")
    file.write("           |  0.0-0.1 |   0.1-0.2 |   0.2-0.3 |   0.3-0.4 |   0.4-0.5 |   0.5-0.6 |   0.6-0.7 |   0.7-0.8 |   0.8-0.9 |   0.9-1.0|\n")
    int_inf = 0.0
    int_sup = 1/inc
    incr = 1 / inc
   
    for i in range(10):
        
        text_temp = f"{str(int_inf)[:3]} - {str(int_sup)[:3]}  |   {int(table_chi[i,0])}  |   {int(table_chi[i,1])}   |   {int(table_chi[i,2])}   |   {int(table_chi[i,3])}   |   {int(table_chi[i,4])}   |   {int(table_chi[i,5])}   |   {int(table_chi[i,6])}   |   {int(table_chi[i,7])}   |   {int(table_chi[i,8])}   |   {int(table_chi[i,9])}  |"
        file.write(text_temp + "\n")
        int_inf = int_sup
        int_sup += incr
    print("chi calculado:",chi_cal)
    chi_cri = 124.3421
    
    if chi_cal <= chi_cri:
        print("cumple la prueba de independencia")
    else:
        print("no cumple la prueba de independencia")



if  __name__ == "__main__":
    main()