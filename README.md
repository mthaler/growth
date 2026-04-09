# growth
Growth code in Go

Exponential growth is quite simple:

$$
\frac{dN}{dt}=(b-d)N
$$

where $(b-d)$ is called intrinsic rate of natural increase. Solving this we get

$$
N(t)=c e^{t (b-d)}
$$