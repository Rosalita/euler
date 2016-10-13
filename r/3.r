#The prime factors of 13195 are 5, 7, 13 and 29.
#What is the largest prime factor of the number 600851475143 ?

#install.packages("gmp")
library(gmp) 
factors <- factorize(600851475143)
max(factors)
