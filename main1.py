import pandas as pd
import numpy as np
import itertools


df = pd.read_csv('input_0.csv')
length = len(df)
distance = np.zeros((length, length))
for i in range(length):
    for j in range(length):
        distance[i, j] = (df.at[i, 'x']-df.at[j, 'x']) ** 2. + \
            (df.at[i, 'y']-df.at[j, 'y']) ** 2.
print(distance)
number = [int(i) for i in range(1, length)]
c = list(itertools.permutations(number))
df_c = pd.DataFrame(c)
df_c[length-1] = 0
df_d = df_c.copy()*0.
for index, row in df_c.iterrows():
    prev = 0
    for k in range(length):
        nex = df_c.at[index, k]
        df_d.at[index, row[k]] += distance[prev, nex]
        prev = df_c.at[index, k]

df_d[length] = df_d.sum(axis=1)
min_index = df_d.idxmin()[length]
df_e = df_c.iloc[min_index].values.tolist()
df_e = [str(int(o)) for o in df_e]
print(np.sqrt(df_d.iloc[min_index, length]))
with open('solution_yours_0.csv', 'w') as f:
    f.write('index'+'\n')
    for ans in df_e:
        f.write(ans+'\n')
