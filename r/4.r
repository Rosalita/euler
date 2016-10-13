#A palindromic number reads the same both ways. 
#The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
#Find the largest palindrome made from the product of two 3-digit numbers.

#generate a vector of products of all two digit numbers
n <- 1:99
y <- as.numeric() 
for(i in n){
    x <- i * n
    y <- c(y, x)
}
unique(y)


# recursive function
result <- function(x,y){
  return(x*1:y)
}
y <- lapply(1:99, function(x) result(x,99))
y <- unique(unlist(test))

# matrix outer product operator %o%
x <- 1:99
y <- sort(unique(as.vector(x%o%x)))
y

#generate product of all three digit numbers
x<- 1:999
y <- sort(unique(as.vector(x%o%x)))
y


?nchar
nchar(y)
