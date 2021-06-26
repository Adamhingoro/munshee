# Munshee
A golang based microservice for **Chart-of-accounts** and **Ledger.** 

## Idea
Microservices are cool but developing them is hard. I had to deal with a lot of portals in my daily software-engineering life and creating same things again and again really grinds my gears. So... a **plugable, developer and business friendly** chart-of-accounts would be a awesome project to make. 

## Plan
I have experience in NodeJs, Python, Java, PHP and DotNet but Golang is a different beast. I have done quite a work in golang but creating a complete project is a different for me. 
I am going to create a **DDD (Domain Driven Design)** for this and I would use **hexagonal** architecture approach. 


![Image of hexagonal architecture](https://cdn-cgbdj.nitrocdn.com/RbczMDpxKIrQLdqnZdHDBvZTsISICJjh/assets/static/optimized/rev-11068f9/wp-content/uploads/2018/10/Screen-Shot-2018-10-26-at-09.36.50.png)
[Click here to read more about hexagonal architecture](https://apiumhub.com/tech-blog-barcelona/applying-hexagonal-architecture-symfony-project/)

And i am following golang project standards from [Golang-Standards](https://github.com/golang-standards/project-layout)

###### Http API's 
Will surely create RESTful API's along wit OAuth authentication. 
###### Event Driven 
Will create **ports** for Kafka, SQS and RabbitMQ. 
###### Fast
Will surely create a caching layer which can be configured for Redis or Memcache. 
###### Reports 
Will add standard Accounting reports and will create interfaces to implement custom reports. 

## More

this project is just started yet. And soon will be updated with more details.