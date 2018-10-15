# Installation and Development
0. git clone https://github.com/danackerson/outlyer outlyertest
0. cd outlyertest
0. go get -v ./...

# Test, Build and Run
0. go test
0. go build server.go
0. ./server
0. curl http://localhost:8080/metrics (addtl. http://localhost:8080/metrics/all to examine registry)

# Which parts did you spend the most time with? How long did it take you in total?
I approached it in 3 phases taking ~3.5hr:
0. HTTP server, metrics daemon (no real stats yet) (~1hr)
0. Find & utilize a stats library (github.com/shirou/gopsutil) (~1.5hr)
0. Code refactoring, clean-up & thread-safety (~1hr)

# What did you find to be the most difficult?
As I'm coding in a ChromeOS linux container (Crostini), the CPU load measurements were often 0 (I think the security sandboxing limits the ). After beating my head against the wall, I decided to build a binary and test it out on my Digital Ocean droplet running CoreOS. Worked perfectly :)

# What would you add to your solution if you had more time?
- The bonus section reading in some nginx/redis stats
- CirleCI build & deploy orchestration via docker

# Why did you choose this programming language?
As I'm applying for the Go Developer position, thought it would be helpful to have a working Golang implementation to examine.

# How did you find the test overall? If you have any suggestions on how we can improve the test, we'd love to hear them.
Interesting & fun: I actually looked forward to the time I could take to work on it and learnt about some cool stats libraries available to the community.
Always unsure about the bonus sections: Are they really just for bonus points or do you want to see if folks go the extra mile?
