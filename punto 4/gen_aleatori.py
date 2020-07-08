import numpy as np

np.random.seed(0)
random_num = np.random.rand(500)
num1 = open("Number1.txt","w+")
for i in random_num:
    num1.write(str(i))
    num1.write(",")
num1.close()
np.random.seed(34)
random_num = np.random.rand(300)
num2 = open("Number2.txt","w+")
for i in random_num:
    num2.write(str(i))
    num2.write(",")
num2.close()
np.random.seed()
random_num = np.random.rand(10000)
num3 = open("Number3.txt","w+")
for i in random_num:
    num3.write(str(i))
    num3.write(",")
num3.close()