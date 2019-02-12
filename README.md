# After Tax Income Service V2: This Time, It's Personal  

## Summary  
The OG [`after tax income service`](https://github.com/DylanLennard/after-tax-income-service) was a way to make a tool that I would use regularly as well as to learn about AWS deployment tools that I was exposed to at work. The goal was to learn about docker, EC2, ECS and Jenkins. A lot of the handler python code is from a command line app I made years ago (and am embarrased by), but besides that I felt the project was a success as I accomplished what I set out to do with it.   

I wanted to take a second stab at this service, but to learn some new tools in the process. I also hope to do this in such a way that the code flows a little better and makes more logical sense. For example, I feel that structs with methods are a much better way of representing the type of data structure I'd like for the Tax information for state and federal than a python class.  


## Tools of Interest   
* Golang  
* AWS Lambda  
* AWS API Gateway  


## What Does This Do  
Currently, the service will assume that you are in California and want your taxes for the 2018 year, and with that it will calculate your take home pay after taxes (assuming no deductions or pre-tax contributions). You pass the `Income` query parameter to give your pre-tax income, and you can optionally modify the `SelfEmploymentStatus` parameter if you'd like to see your rate assuming you were self-employed.  


## What This Does Not Do  
Replace Turbotax or HR Block.  


## TODO List (as of 02/11/2019) 
* Adjust the job to return an errorcode in event that bad parameter is passed or no parameters passed   
* Dockerize the job  
* Reconfigure the jenkins build server to run tests and such on this code and to update the lambda   
    * Once this done, can tear down the V1 version of the service  
* Nice to Have (but need a vacation to do): scrape tax data from appropriate sites and load it into a either a json file or a some sort of datastore. 
    * Can then have this vary by state  