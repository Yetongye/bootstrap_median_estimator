# Load required library
if (!require("boot")) {
  install.packages("boot", repos = "http://cran.us.r-project.org")
  library(boot)
}

# Set seed for reproducibility
set.seed(123)

# Generate sample data from a normal distribution
sample_data <- rnorm(100, mean = 0, sd = 1)

# Define statistic function: compute median of resampled data
median_func <- function(data, indices) {
  return(median(data[indices]))
}

# Record start time
start_time <- Sys.time()

# Perform bootstrap: 1000 resamples
bootstrap_result <- boot(data = sample_data, statistic = median_func, R = 1000)

# Record end time
end_time <- Sys.time()

# Calculate standard error of the median
se_median <- sd(bootstrap_result$t)

# Print and save results
sink("bootstrap_r/bootstrap_r_output.txt")
cat("This is a bootstrap median estimator using R:\n")
print(bootstrap_result)
cat("\nStandard Error of the Median:", se_median, "\n")
cat("\nExecution time:", round(difftime(end_time, start_time, units = "secs"), 3), "seconds\n")

# Record memory usage
cat("\nMemory usage snapshot from gc():\n")
print(gc())
sink()

# Plot the bootstrap distribution
png("bootstrap_r/bootstrap_r_plot.png")
plot(bootstrap_result)
dev.off()
