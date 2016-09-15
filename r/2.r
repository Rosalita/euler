options(scipen=999)
length <- 33 # use 33 as the 34th fib value is > 4,000,000
fib <- numeric(length) 
fib[1] <- 1
fib[2] <- 1
for (i in 3:length) { 
  fib[i] <- fib[i-1]+fib[i-2]
} 
even <- which (fib %% 2 == 0)
sum(fib[even])
