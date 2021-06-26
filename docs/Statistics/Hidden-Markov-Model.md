# Hidden Markov Model

## Definition

A Hidden Markov Model (HMM) consisted by the following parts:

1. Ground truth states, represented by a Markov process $S_t=s_1,\dots, s_N$.
2. Observable events, represented by a Markov process $O_t=o_1,\dots,o_N$.
3. Emission probability or observation probability $\mathbb{P}(O_t=o_t|S_1=s_1, \dots, S_t=s_t) = \mathbb{P}(O_t=o_t|S_t=s_t)$. The emission probability can also be represented as a matrix $E$ such that the element $E_{i,s}=\mathbb{P}(O_t = i|S_t = s)$.
4. Tansition probability $T$ which the $ij$-th element $T_{ij}=\mathbb{P}(O_{t}=j|O_{t-1}=i)$ and then $\sum_iT_{ij}= 1,\forall j$.  
5. The initial probability distribution of truth states $\pi$ that $\pi_i = \mathbb{P}(S_1=i)$, $\sum_i\pi_i=1$. 

The HMM introduced here also assume that it is time-homogeneous which means that both transistion probability and emission probability is independent on the time subscripts. 

The inference problem of HMM can be classified into three types:

1. Given a observation sequence $O=\lbrace o_1,\dots,o_N\rbrace$ and a HMM assumption $\Lambda=(E,T)$, determine the likelihood $\mathbb{P}(O)$.
2. Given an observable sequence $O=\lbrace o_1,\dots,o_N\rbrace$ and a assumption $(E,T)$, discover the best hidden state sequence $S$.
3. Given a set of observables and sequence of stats, learn the observation probability and transition probability $(E,T)$.

## Likelihood Computation

For a HMM model with an assumption $\Lambda = (E,T)$ and a sequnece of observations $O=\lbrace o_1, \dots, o_{n}\rbrace$ is given, determine the likelihood $L_\Lambda(O)=\mathbb{P}(O|\Lambda)$ for obtaining the observation $O$ with respect to the assumption $\Lambda$. 

Based on the Markov assumption and the properties of conditional probability, it can be decomposed into

$$
\mathbb{P}(O) = \sum_k\mathbb{P}(O_n = o_n|S_n=k)\mathbb{P}(S_n=k|O_1= o_1,\dots, O_{n-1} = o_{n-1}) ,
$$
where the last term on the right side can be expressed recursively like:

$$
\mathbb{P}(S_t=s_t|O_1= o_1,\dots, O_{t-1} = o_{t-1}) = \sum_{s_{t-1}}\mathbb{P}(S_{t-1}=s_{t-1}|O_1= o_1,\dots, O_{t-2} = o_{n-2})\mathbb{P}(s_t|s_{t-1})\mathbb{P}(o_{t-1}|s_t)
$$ 

since by definition we have:

$$
\mathbb{P}(S_t=s_t|O_1= o_1,\dots, O_{t} = o_{t})=\mathbb{P}(o_t|s_t)\mathbb{P}(S_t=s_t|O_1= o_1,\dots, O_{t-1} = o_{t-1}).
$$

This all bring us to a algorithm for solving this problem. Let's first define a parameter $\alpha_{t}(s_t)$ as the probability of the $t$-th time step is $s_t$ with the observations are fixed up to time step $t$:

$$
\alpha_{t}(s_t) = \mathbb{P}(S_t=s_t|O_1= o_1,\dots, O_{t} = o_{t}).
$$

We can summary the algorithm as follow:

1. Initializing with $\alpha_{1}(s_1) = \mathbb{P}(s_1|o_1) = \pi_{s_1}E_{o_1, s_1}$.
2. Iteratively solving the $\alpha_{t}(s_t)$ by the formula:

$$
\alpha_{t}(s_t) = \sum_{s_{t-1}}\alpha_{t-1}(s_{t-1})T_{s_{t-1},s_t}E_{o_t, s_t}, \quad 1\le t\le N.
$$

3. The likelihood $\mathbb{P}(O|\Lambda)$ is then

$$
\mathbb{P}(O|\Lambda)=\sum_{s_n}\alpha_{n}(s_n).
$$

## Decoding: Viterbi Algorithm

Given a sequence of observation $O=\lbrace o_1, \dots, o_{n}\rbrace$ and a HMM assumption $(E,T)$. Determine the most probable hidden stats. 

The most probable hidden stats $S=\lbrace s_1, \dots, s_{n}\rbrace$ are those maximizing the probability 

$$
v_n(s_n) = \mathbb{P}(S_1=s_1,\dots, S_n=s_n,O_1=o_1,\dots O_n=o_n).
$$

Suppose $v_{n-1}(s)$ is known, then we have the relation

$$
v_n(s_n) = \max_{s_{n-1}}v_{n-1}(s_{n-1})T_{s_{n-1},s_n} E_{s_n,o_n}.
$$

A backtrace algorithm called Viterbi's Algorithm can be used to solve this problem by using the recurtion above. The pseudo code is

```py
def Viterbi(o : observation_array):
	T = transition_probability
	E = emission_probability
	p = initial_state_probability

	viterbi[M][N] # 2D array for total time step N and total candidate hidden states M 

	for s in hidden_states:
		viterbi[s,1] = p[s]E(o[1], s)
		backtrace[s,1] = 0
	for t in range(2, N):
		for s in hidden_states:
			viterbi[s,t] = max( viterbi[s2 , t-1]*T[s, s2]*E(o[t], s) for s2 in hidden_states)
			backtrace[s,t] = s2 # states that maximizing the element above
	bestpathprob = max( viterbi[s, N] for s in hidden_states)
	bestpathpointer = backtrace[s, N] # s is the one maximized the bestpathprob
	bestpath = []
	t = N
	while not bestpathpointer == 0:
		bestpath[t] = bestpathpointer
		t-=1
		bestpathpointer = backtrace[bestpathpointer, t] 

	return bestpath, bestpathprob
```

## Learning: The Forward-Backward Algorithm

Given sequences of observations and known the possible hidden state domain, learn the HMM parameters $(E,T)$. The Baum-Welch algorithm is desinged based on the idea of the large number theorem that the possible occurance frequency is an unbias estimator for these parameters. 

$$
\begin{aligned}
\hat T_{x,y}&=\frac{\text{number of states transfer } x\to y}{\text{number of states transfer from }x},\\
\hat E_{x,k}&=\frac{\text{number of observation of } k \text{, with hidden states } x}{\text{number of hidden states of }x}
\end{aligned}
$$

Let's first define a parameter $\beta_t(s_t)$ to be the probability of the observations after $t$ is given as $O_{t+1}=o_{t+1},\dots O_n=o_n$ provided the $t$-th hidden state is $s_t$:

$$
\beta_t(s_t) = \mathbb{P}(o_{t+1}, \dots, o_n|S_t=s_t).
$$

This value can be calculated in a **backward algorithm**:

1. Initialization with $\beta_n(x) = 1$ for all candidate state $x$.
2. Recusively calculate

$$
\beta_t(x) = \sum_{y}T_{x,y}E_{y,o_{t+1}}\beta_{t+1}(y), \quad 1\le t \le n.
$$ 

3. The likelihood function with assumption $\Lambda = (E,T)$ now can be expressed as 

$$
\mathbb{P}(O|\Lambda) = \sum{s_t}\alpha_t{s_t}\beta_t(s_t),\quad 1\le t \le n.
$$

Let the parameter $\xi_t(x,y)$ be the probability of the $S_t=x$ and $S_{t+1}=y$ provided the all observations are given under the assumption $\Lambda$:

$$
\xi_t(x,y) =\mathbb{P}(S_t=x, S_{t+1}=y|O,\Lambda).
$$

With the conditional probability:

$$
\mathbb{P}(X,Y|Z) = \mathbb{P}(X|Y,Z)\mathbb{P}(Y|Z),
$$

we have

$$
\mathbb{P}(S_t=x, S_{t+1}=y|O,\Lambda) = \frac{\mathbb{P}(S_t=x, S_{t+1}=y,O,|\Lambda)}{\mathbb{P}(O|\Lambda)},
$$
which is

$$
\xi_t(x,y)=\frac{\alpha_t(x)T_{x,y}E_{y,s_{t+1}}\beta_{t+1}(s_{t+1})}{\sum_{s_t}\alpha_t(s_t)\beta_t(s_t)}.
$$

The estimator of $\hat T_{x,y}$ is then

$$
\hat T_{x,y} = \frac{\sum_{t=1}^{n-1}\xi_t(x,y)}{\sum_{t=1}^{n-1}\sum_z\xi_t(x,z)}.
$$

Similarly, define a parameter $\gamma_t(s_t)$ as

$$
\gamma_t(x)=\mathbb{P}(S_t=x|O,\Lambda) = \frac{\alpha_t(x)\beta_t(x)}{\mathbb{P}(O|\Lambda)}
$$

then we have the estimator 

$$
\hat E_{x,k} = \frac{\sum_{\lbrace O_t=k \rbrace}\gamma_t(x)}{\sum_{t=1}^n\gamma_t(x)},
$$
where $\sum_{\lbrace O_t=k \rbrace}$ means sum over all the observation samples with $O_t=k$. After we get the two estimators, we can express the algorithm as the following steps:

1. Initializing the $(E,T)$ with a hyperthesis values. 
2. Calculate the parameter $\alpha$ from the forward algorithm and calculate $\beta$ from the backward algorithm. 
3. Calculate the $\gamma, \xi$ from the new $\alpha,\beta$.
4. Update the $(E,T)$.
5. Repeat the steps 2 to 4 until the parameters convergence.
