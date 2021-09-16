# go-perceptron

A simple implementation of the [perceptron algorithm](https://en.wikipedia.org/wiki/Perceptron)
in Go.

## Usage

```
$ git clone git@github.com:karantan/go-perceptron.git
$ cd go-perceptron
$ go run main.go
Starting weights 0.05628316484698421 0.46671059597618897 bias 0.3007628894709572
Iteration 0
0 or 0 = 0 (desired output: 0)
Adjusted weights 0.05628316484698421 0.46671059597618897 bias 0.3007628894709572
0 or 1 = 0 (desired output: 1)
Adjusted weights 0.05628316484698421 0.566710595976189 bias 0.4007628894709572
1 or 0 = 0 (desired output: 1)
Adjusted weights 0.1562831648469842 0.566710595976189 bias 0.5007628894709572
1 or 1 = 1 (desired output: 1)
Adjusted weights 0.1562831648469842 0.566710595976189 bias 0.5007628894709572
Iteration error 2

Iteration 1
0 or 0 = 0 (desired output: 0)
Adjusted weights 0.1562831648469842 0.566710595976189 bias 0.5007628894709572
0 or 1 = 1 (desired output: 1)
Adjusted weights 0.1562831648469842 0.566710595976189 bias 0.5007628894709572
1 or 0 = 0 (desired output: 1)
Adjusted weights 0.2562831648469842 0.566710595976189 bias 0.6007628894709571
1 or 1 = 1 (desired output: 1)
Adjusted weights 0.2562831648469842 0.566710595976189 bias 0.6007628894709571
Iteration error 1

Iteration 2
0 or 0 = 0 (desired output: 0)
Adjusted weights 0.2562831648469842 0.566710595976189 bias 0.6007628894709571
0 or 1 = 1 (desired output: 1)
Adjusted weights 0.2562831648469842 0.566710595976189 bias 0.6007628894709571
1 or 0 = 0 (desired output: 1)
Adjusted weights 0.3562831648469842 0.566710595976189 bias 0.7007628894709571
1 or 1 = 1 (desired output: 1)
Adjusted weights 0.3562831648469842 0.566710595976189 bias 0.7007628894709571
Iteration error 1

Iteration 3
0 or 0 = 0 (desired output: 0)
Adjusted weights 0.3562831648469842 0.566710595976189 bias 0.7007628894709571
0 or 1 = 1 (desired output: 1)
Adjusted weights 0.3562831648469842 0.566710595976189 bias 0.7007628894709571
1 or 0 = 1 (desired output: 1)
Adjusted weights 0.3562831648469842 0.566710595976189 bias 0.7007628894709571
1 or 1 = 1 (desired output: 1)
Adjusted weights 0.3562831648469842 0.566710595976189 bias 0.7007628894709571
Iteration error 0

Final weights 0.3562831648469842 0.566710595976189 bias 0.7007628894709571
```
