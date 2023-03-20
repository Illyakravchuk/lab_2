# Software Architecture. 2nd Lab
### *Created by: [Illya Kravchuk](https://github.com/Illyakravchuk) / [Dima Mudryk](https://github.com/TangiresH) / [Danylo Chohadze](https://github.com/GhostDolphin)*
---
## Uninterapted iteration & testing automation
> Team #12 - Transforming prefix expression into the infix one

The main goal is to implement the uninterapted iteration and enhance basic skills of using dependencies' injection pattern, with an eye to simplify the overall testing process.
---
Postfix expressions are to be calculated upon the use of this project. The list of supported mathematical operations is below:
- Addition (+)
- Subtraction (-)
- Multiplication (\*)
- Division (/)
- Exponentiation (^)
---
In order to start up the project, follow next steps:
1. Clone repo on your local device;
2. Choose project's directory in your terminal;
3. Command to launch project:
```
go run cmd/example/main.go
```
4. Command to read an input expression from the terminal:
```
go run main.go -e="+ * 2 3 4"
```
5. Command to read an input expression from an actual file:
```
go run main.go -f="test.txt"
```
6. Command to input the calculation result into specific file:
```
go run main.go -e="4 2 - " -o="result.txt
```
7. Command to conduct testing:
```
go test
```
---
> [Failed build](https://github.com/Illyakravchuk/lab_2/actions/runs/4463392235/jobs/7838565519)

> [Successful build](https://github.com/Illyakravchuk/lab_2/actions/runs/4463840600/jobs/7839437578)

> [Pull Request](https://github.com/Illyakravchuk/lab_2/actions/runs/4463848460/jobs/7839453486)
