## 巡回セールスマン問題を解こう！
### まず問題を正確に定式化をしましょう
二次元ユークリッド空間上にN個の点<a href="https://www.codecogs.com/eqnedit.php?latex=P_i(x_i,y_i)_i\in&space;N" target="_blank"><img src="https://latex.codecogs.com/gif.latex?P_i(x_i,y_i)_i\in&space;N" title="P_i(x_i,y_i)_i\in N" /></a>
が与えられます。このとき、任意の点からスタートして、全ての点を通り再びはじめの点に戻るような経路のうち、なるべく短いのものをなるべく高速に求めてください。一度通った点は二度通ることができません。

以下の制約条件を用いることができます。
- 任意の3点は三角不等式を満たす。すなわち<a href="https://www.codecogs.com/eqnedit.php?latex=d(P_i,P_j)&space;\leq&space;d(P_i,P_k)&space;&plus;&space;d(P_k,P_j)" target="_blank"><img src="https://latex.codecogs.com/gif.latex?d(P_i,P_j)&space;\leq&space;d(P_i,P_k)&space;&plus;&space;d(P_k,P_j)" title="d(P_i,P_j) \leq d(P_i,P_k) + d(P_k,P_j)" /></a>が成立します。
- <a href="https://www.codecogs.com/eqnedit.php?latex=P_i" target="_blank"><img src="https://latex.codecogs.com/gif.latex?P_i" title="P_i" /></a>から<a href="https://www.codecogs.com/eqnedit.php?latex=P_j" target="_blank"><img src="https://latex.codecogs.com/gif.latex?P_j" title="P_j" /></a>の距離と<a href="https://www.codecogs.com/eqnedit.php?latex=P_j" target="_blank"><img src="https://latex.codecogs.com/gif.latex?P_j" title="P_j" /></a>から<a href="https://www.codecogs.com/eqnedit.php?latex=P_i" target="_blank"><img src="https://latex.codecogs.com/gif.latex?P_i" title="P_i" /></a>の距離が異なることはありません。<a href="https://www.codecogs.com/eqnedit.php?latex=d(P_i,P_j)&space;=&space;d(P_j,P_i)" target="_blank"><img src="https://latex.codecogs.com/gif.latex?d(P_i,P_j)&space;=&space;d(P_j,P_i)" title="d(P_i,P_j) = d(P_j,P_i)" /></a>が常に成立します。

巡回セールスマン問題は一般には、「無向完全グラフ<a href="https://www.codecogs.com/eqnedit.php?latex=G" target="_blank"><img src="https://latex.codecogs.com/gif.latex?G" title="G" /></a>と重み関数<a href="https://www.codecogs.com/eqnedit.php?latex=c:E(G)\rightarrow&space;\mathbb{R}_&plus;" target="_blank"><img src="https://latex.codecogs.com/gif.latex?c:E(G)\rightarrow&space;\mathbb{R}_&plus;" title="c:E(G)\rightarrow \mathbb{R}_+" /></a>が与えられて、<a href="https://www.codecogs.com/eqnedit.php?latex=G" target="_blank"><img src="https://latex.codecogs.com/gif.latex?G" title="G" /></a>の最小重みのハミルトン閉路を求める問題」（コルテ・フィーゲン, 2009, p. 350）として定式化されます。無向完全グラフとはどの2つの点の間にも向きのない辺があるグラフのことで、ハミルトン閉路とは同じ点を2度以上通ることなく全ての点を通り、始点と終点が一致しているような閉路を指します。

### とりあえずBrute Forceでやってみる


### 参考文献


## 以下は元のREADMEの内容です
Build@Mercari 2020 Week4 - Travelling Salesman PRoblem Challenges.

This is forked from [https://github.com/hayatoito/google-step-tsp-2016](https://github.com/hayatoito/google-step-tsp-2016).

1. 問題
[巡回セールスマン問題](https://ja.wikipedia.org/wiki/%E5%B7%A1%E5%9B%9E%E3%82%BB%E3%83%BC%E3%83%AB%E3%82%B9%E3%83%9E%E3%83%B3%E5%95%8F%E9%A1%8C) を解くアルゴリズムを考えてください。

2. 課題
----
このrepositoryを自分のgithubにforkして使ってください。
N = 5 から N = 2048までの７つの課題があります。

| Challenge    | N (= the number of cities) | Input file  | Output (Solution) file |
| ------------ | -------------------------: | ----------- | ---------------------- |
| Challenge 0  |                          5 | input_0.csv | solution_yours_0.csv   |
| Challenge 1  |                          8 | input_1.csv | solution_yours_1.csv   |
| Challenge 2  |                         16 | input_2.csv | solution_yours_2.csv   |
| Challenge 3  |                         64 | input_3.csv | solution_yours_3.csv   |
| Challenge 4  |                        128 | input_4.csv | solution_yours_4.csv   |
| Challenge 5  |                        512 | input_5.csv | solution_yours_5.csv   |
| Challenge 6  |                       2048 | input_6.csv | solution_yours_6.csv   |

inputとoutputの形式については *3. Data Format Specification* を見てください。
### Your tasks

* 巡回セールスマン問題をとくアルゴリズムを考えて実装してください。
* `solution_yours_{0-6}.csv` をあなたのアルゴリズムでといた結果で上書きしてください。
* あなたの解法の*path length*を[scoreboard]に書いてください

[scoreboard]: https://docs.google.com/spreadsheets/d/1t4ScULZ7aZpDJL8i9AVFQfqL7sErjT5i3cmC1G5ecR8/edit?usp=sharing
### An optional task (Speed challenge)

What matters in this optional task is your program's *speed* (execution time). The path length does not matter as long as it meets the condition.
Your task is:

* Given `input_6.csv`, write a program which outputs a path shorter than `47,000`

Input your program's execution time in the [scoreboard]. Faster (smaller) is better.

You can measure the execution time by `time` command.

```shellsession
$ time yourprogram input_6.csv solution_yours_6.csv
2.96s user 0.07s system 97% cpu 3.116 total
```

In this case, your score is `3.116` (s).

### Visualizer

The demo page of the visualizer is:
https://mercari-build.github.io/week4-tsp/visualizer/.

The assignment includes a helper Web page, `visualizer/index.html`, which
visualizes your solutions. You need to run a HTTP server on your local machine
to access the visualizer. Any HTTP server is okay. If you are not sure how to
run a web server, use the following command to run the HTTP server included in
the assignment. Make sure that you are in the top directory of the assignment
before running the command.

``` shellsession
./nocache_server.py # For Python 3
./nocache_server.py2.py # If you don’t want to install Python3
```

Then, open a browser and navigate to the
[http://localhost:8000/visualizer/](http://localhost:8000/visualizer/). Note
that visualizer was only tested by Google Chrome.  Using the visualizer is
up-to you. You don’t have to use the visualizer to finish the assignment. The
visualizer is provided for the purpose of helping you understand the problem.

3. Data Format Specification
----

### Input Format

The input consists of `N + 1` lines. The first line is always `x,y`. It is followed by `N` lines, each line represents an i-th city’s location, point `xi,yi` where `xi`, `yi` is a floating point number.

```
x,y
x_0,y_0
x_1,y_1
…
x_N-1,y_N-1
```

### Output Format

Output has `N + 1` lines. The first line should be “index”. It is followed by `N` lines, each line is the index of city, which represents the visitation order.

```
index
v_0
v_1
v_2
…
v_N-1
```

### Example (Challenge 0, N = 5)

Input Example:

```
x,y
214.98279057984195,762.6903632435094
1222.0393903625825,229.56212316547953
792.6961393471055,404.5419583098643
1042.5487563564207,709.8510160219619
150.17533883877582,25.512728869805677
```

Output (Solution) Example:

```
index
0
2
3
1
4
```

These formats are requirements for the visualizer, which can take only properly formatted CSV files as input.

5. What’s included in the assignment
----

To help you understand the problem, there are some sample scripts / resources
in the assignment, including, but not limited to:

- `solver_random.py` - Sample stupid solver. You never lose to this stupid one.
- `solution_random_{0-6}.csv` - Sample solutions by solver_random.py.
- `solver_greedy.py` - Sample solver using the greedy algorithm. You should beat this definitely.
- `solution_greedy_{0-6}.csv` - Sample solutions by solver_greedy.py.
- `solution_sa_{0-6}.csv` - Yet another sample solutions. I expect all of you will beat this one too. The solver itself is not included intentionally.
- `solution_wakanapo_{0-6}.csv` - Sample solutions I solved when I was joined Google STEP 2016.
- `solution_yours_{0-6}.csv` - You should overwrite these files with your solution.
- `solution_verifier.py` - Try to validate your solution and print the path length.
- `input_generator.py` - Python script which was used to create input files, `input_{0-6}.csv`
- `visualizer/` - The directory for visualizer.

6. Acknowledgments
----
この課題は[私](https://github.com/wakanapo)がgoogle step 2016に参加したときにやったものです。問題のわかりやすさ、visualizerによるアルゴリズムのみやすさ、楽しさなどにおいてこれを上回る課題はないと思ったので、Build@Mercariでも採用することにしました。
