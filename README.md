# sdk-go
The general web development library is encapsulated and managed in a unified way.
Integrating the unified log library, database tools, configuration analysis and HTTP middleware.
See readme, example directory and corresponding test case file of each package for specific use examples.

# Directory Structure
```
./
├─code
├─configure
├─email
├─example
│  ├─configure
│  └─httpserver
│      ├─config
│      ├─handlers
│      ├─middleware
│      ├─proto
│      ├─routes
│      └─server
├─httpserver
│  ├─middleware
│  └─result
├─jwt
├─logger
│  └─hooks
├─model
│  ├─mongo
│  │  └─test
│  ├─mysql
│  └─redis
│      └─test
├─oss
│  ├─qiniu
│  └─ucloud
└─utils
├─mapper
├─slice
└─timer
```