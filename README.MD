<a name="readme-top"></a>

# Grinding - Never back down, never what?

## About this repository

Grinding is place that I created to store my material for Backend Engineering Interview Preparation.

## 1. System Design
### Fundamental Concepts
- CAP Theorem
- ACID vs BASE
- Consistency patterns
- Eventual vs Strong consistency
- Load balancing strategies
- DNS and CDN architecture
- Database sharding
- Replication strategies
- Failover and redundancy

### Scalability
- Horizontal vs Vertical scaling
- Database indexing strategies
- Caching layers (Client, CDN, Application, Database)
- Message queues and event-driven architectures
- Rate limiting
- Connection pooling
- Backpressure handling
- Data partitioning strategies

### High Availability
- Redundancy patterns
- Failover strategies
- Disaster recovery
- Health checks
- Circuit breakers
- Graceful degradation
- Blue-green deployment
- Canary releases
- Feature flags

### Common System Design Problems
- Design a CI/CD system
- Design Git hosting service
- Design notification system
- Design rate limiter
- Design distributed cache
- Design job scheduler
- Design monitoring system
- Design authentication system

## 2. Distributed Systems
### Core Concepts
- Distributed consensus (Raft, Paxos)
- Leader election
- Clock synchronization
- Vector clocks
- Quorum
- Gossip protocols
- Distributed transactions
- Split brain problem

### Patterns
- Saga pattern
- CQRS
- Event sourcing
- Bulkhead pattern
- Sidecar pattern
- Ambassador pattern
- Circuit breaker pattern
- Retry pattern

### Challenges
- Network partitions
- Distributed deadlocks
- Race conditions
- Thundering herd
- Hot spots
- Data consistency
- Service discovery
- Distributed tracing

## 3. Database Systems
### Concepts
- Transaction isolation levels
- Locking mechanisms
- Query optimization
- Execution plans
- Index types
- Materialized views
- Vacuum processing
- Write-ahead logging

### PostgreSQL Specific
- MVCC
- Table partitioning
- Inheritance
- Custom types
- Full-text search
- JSON operations
- Window functions
- Common table expressions

### Performance
- Query optimization
- Index strategy
- Connection pooling
- Parallel query execution
- Table partitioning
- Vacuum strategies
- Buffer management
- Statistics collection

## 4. Cloud Native & DevOps
### Container Orchestration
- Kubernetes architecture
- Service mesh
- Container security
- Resource management
- Auto-scaling
- Rolling updates
- ConfigMaps and Secrets
- Network policies

### CI/CD
- Pipeline design
- Automated testing
- Infrastructure as Code
- Artifact management
- Environment management
- Deployment strategies
- Rollback procedures
- Security scanning

### Monitoring & Observability
- Metrics collection
- Log aggregation
- Distributed tracing
- Alert management
- Performance monitoring
- Error tracking
- SLI/SLO/SLA
- Root cause analysis

## 5. Security
### Application Security
- Authentication methods
- Authorization patterns
- OAuth 2.0 and OpenID Connect
- JWT handling
- CSRF protection
- XSS prevention
- SQL injection prevention
- Input validation

### Infrastructure Security
- Network security
- Container security
- Secret management
- IAM best practices
- SSL/TLS configuration
- VPC design
- Security groups
- Penetration testing

## 6. Coding & Algorithms
### Data Structures
- Trees (Red-Black, B-Tree)
- Graphs
- Hash tables
- Priority queues
- LRU cache
- Bloom filters
- Skip lists
- Tries

### Algorithms
- Graph algorithms
- Dynamic programming
- Greedy algorithms
- Binary search variations
- String matching
- Sorting algorithms
- Tree traversal
- Path finding

### Go-Specific Topics
- Goroutines and channels
- Memory model
- Garbage collection
- Interface design
- Error handling
- Context package
- Reflection
- Testing patterns

## 7. API Design
### REST
- Resource modeling
- URL structure
- HTTP methods
- Status codes
- Error handling
- Versioning
- Documentation
- Rate limiting

### GraphQL
- Schema design
- Resolvers
- Type system
- Mutations
- Subscriptions
- Batching
- Caching
- N+1 problem

## 8. Software Architecture
### Patterns
- Microservices
- Event-driven
- Layered architecture
- Hexagonal architecture
- Domain-driven design
- Clean architecture
- SOLID principles
- Design patterns

## Study Resources
### Books
- Designing Data-Intensive Applications
- Clean Architecture
- Cloud Native Patterns
- Database Internals
- Building Microservices
- Go Programming Language
- Site Reliability Engineering
- System Design Interview

## Data structure and Algorithm

#### Prefix sum

Prefix sum helps solve problems that require calculating the sum or range of continuous elements in an array. It is a powerful technique often used in competitive programming and technical interviews to optimize the performance of sum queries.

Sample question:

[https://leetcode.com/problems/range-sum-query-immutable/](leetcode/1.prefix-sum/range-sum-query-immutable.go)


## [Golang](https://go.dev/)

Golang is my most beloved programming language, because it clean, simple syntax but have powerful feature like concurrency, blazing fast, good toolbox (go vet, linter, gofmt, go mod, ..ect). Here are some resource that can help you strengthen your Go skill.

- [Go 101 books](https://go101.org/) My first Golang book I read, that help me learn core fundamental of Go. Good for starter.
- [Concurrency in Go](https://www.oreilly.com/library/view/concurrency-in-go/9781491941294/) Famous book that help you deeply understanding concurrent concept and practical. High recommend for intermediate and upper Go learner.
- [Go Gotchas](https://yourbasic.org/golang/gotcha/) This blog help you point-out most mistake you can make when coding with very straightforward example and how to address the problem.

#### My public learning notes

[Today I learn](https://github.com/kurojs/til/issues) This is where I quickly note down things I found that usefull and interest.

<p align="right">(<a href="#readme-top">back to top</a>)</p>
