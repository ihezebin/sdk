# logger
> This package is based on reading the logrus source code, 
encapsulating it, and then adding some commonly used custom functions, 
and adapting to logrus

For example:
- logger has built-in file segmentation hook;
- Built in quick creation of lohger, which contains the common fields of Web services;
- The default format is JSON;
- By default, the simplified file name and function name are carried;

Encapsulating logrus makes me better understand the principle of logrus. Even most of the code is similar to logrus. I must thank logrus !

I may add some other hooks later, such as outputting logs to Redis, ElasticSearch, etc.