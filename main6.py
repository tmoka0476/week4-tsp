import pandas as pd

def　distance(input,output)
df = pd.read_csv('input_6.csv')
df_x = pd.DataFrame(
    {"x_int": df['x'], "x": df['x'], 'y_int': df['y'], "y": df['y']})
df_x['x_int'] = df_x['x_int']*0.01
df_x['y_int'] = df_x['y_int']*0.01
df_x['x_int'] = df_x['x_int'].astype(int)
df_x['y_int'] = df_x['y_int'].astype(int)
df_x = df_x.sort_values(['x_int', 'y_int'])
group_df_x = df_x.groupby(['x_int', 'y_int'])
index_list = []
for i in range(16):
    index_sub = []
    for j in range(1, 9):
        df_sub = group_df_x.get_group((i, j))
        df_sub = df_sub.sort_values('y')
        index_sub += list(df_sub.index)
    if i % 2 == 0:
        index_list += index_sub
    else:
        index_list += reversed(index_sub)
index_sub = []
for k in range(16):
    df_sub = group_df_x.get_group((k, 0))
    df_sub = df_sub.sort_values('x')
    index_sub += list(df_sub.index)
index_list += reversed(index_sub)
index_list = [str(int(o)) for o in index_list]
with open('solution_yours_6.csv', 'w') as f:
    f.write('index'+'\n')
    for ans in index_list:
        f.write(ans+'\n')
