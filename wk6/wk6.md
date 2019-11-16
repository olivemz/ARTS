## Algorithm

[wk6.go](wk6.go)

## Review
Domain driven design
https://en.wikipedia.org/wiki/Domain-driven_design
https://medium.com/the-coding-matrix/ddd-101-the-5-minute-tour-7a3037cf53b8
https://martinfowler.com/bliki/CQRS.html
https://martinfowler.com/tags/domain%20driven%20design.html


usually the best sample is database architecture. CRUD methods, with an abstract layer that exposed to users. So user does not really care about the underhood implementation. It provides huge benefit in following points
1. database migration:\
1.1. Dont really need to concern the underhood difference
2. service segamentation:\
2.1 since micro service is more and more popular, Reposotry driven is more popular as well. It provides convinient when design the service and how to combine them into one.


## Tips
use `nohup ... &` to make the process run in back end. 
use `sed -n '1,10p' x.file` to print file with specific lines. `sed` could accept regex

## Share
N/A
