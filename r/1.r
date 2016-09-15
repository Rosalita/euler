total <- 999
numbers <- 1:total
div3 <- (numbers %% 3)
div5 <- (numbers %% 5)
results <- which((div3 == "0")|(div5 == "0"))
answer <- sum(results)
answer
