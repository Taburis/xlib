# Algorithm Cook Book

## Tomohiko Sakamoto’s Algorithm

This algorithm is used to find the date according to the Gregorian Calendar. The idea is that the date is determined by the number of days, $n$, away from the origin, the first day of Gregorian Calendar, which is "Monday". The exact date is the modulo $n\mid 7$. To simplify this procedure, we just need to calculate the date of the first day of each month in that year, which formed an array 

$$
t = \lbrace 0, 3, 2, 5, 0, 3, 5, 1, 4, 6, 2, 4\rbrace
$$

where the number means the days ahead of the Monday for each month. For instance, the 1st Jan is Monday so it is marked as 0 as the first element in the array and the 1st Feb is $31\mid 7 = 3$ and the 1st May is 2 comes from $29%4 = 1$ so that delayed 3 by 1 day and get the final result 2. Since $365\mid 7 = 1$, this array will needs to shifted by 1 every year. However, the leap year makes it complicate. A year $y$ is a leap year if and only if $y\mid 4 = 0 $ and $y\mid 100 \ne 0$ (except the special case that $y\mid 400 = 0$). Taking this into account, we got the pseudo-codes for the Sakamoto's algorithm as follow:

```py
def Sakamoto(int year, int month, int day):
	t = [0, 3, 2, 5, 0, 3, 5, 1, 4, 6, 2, 4]
	if month < 3 : year -= 1
	# y/4 is an estimate of how many leap years before `year`
	# y/100 the fake leap years
	# y/400 the exceptions need to be added back
	return (y + y/4 - y/100 + y/400 + t[month -1] + day)%7
```