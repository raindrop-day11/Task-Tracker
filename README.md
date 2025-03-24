Task Tracker 是一个被用来追踪和管理你的任务的项目

要求：
在命令行中运行，接受用户的输入和行为作为参数，然后存储在JSON文件中。用户可以
1. Add, Update, Delete Tasks
2. Mark a task as in progress or done
3. List all tasks
4. List all tasks that are done
5. List all tasks that are not done
6. List all tasks that are in progress

限制：
1. 使用命令行的位置参数来接受用户输入
2. 不要使用外部库或者框架，使用编程语言的原生文件系统模块与JSON文件交互

任务属性：
1. id:
2. description:
3. status:
4. createdAt:
5. updatedAt:
在添加新任务时，务必将这些属性添加到JSON文件中，并且在更新任务时更新相应的属性