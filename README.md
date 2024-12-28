# TaskTrackerCLI

**TaskTrackerCLI** is a command-line interface (CLI) for managing tasks. You can use this tool to list, add, delete, update, and mark tasks.

## Features

- **List all tasks**: View all tasks in the system.
- **List tasks by status**: Filter tasks based on their status (to-do, in-progress, done).
- **Add new tasks**: Add a task with a description.
- **Delete tasks**: Remove tasks from the system by ID.
- **Update tasks**: Update tasks' descriptions.
- **Mark tasks**: Mark tasks as `done`, `in-progress`, or `to-do`.

## Usage

### 1. **List All Tasks**

To list all tasks, use the following command:
```bash
go run . list
```

### 2. **List Specific Tasks by Status**

You can filter tasks based on their status. Use the following commands to list tasks by status:

- **To-Do Tasks**:
  ```bash
  go run . list to-do
  ```

- **In-Progress Tasks**:
  ```bash
  go run . list in-progress
  ```

- **Completed Tasks (Done)**:
  ```bash
  go run . list done
  ```

### 3. **Add a New Task**

To add a new task, provide a description for the task:
```bash
go run . add "Task description"
```

Example:
```bash
go run . add "Finish writing README file"
```

### 4. **Delete a Task**

To delete a task, provide the ID of the task you want to remove:
```bash
go run . delete <task_id>
```

Example:
```bash
go run . delete 1
```

This will delete the task with the ID `1`.

### 5. **Update a Task Description**

To update the description of a task, use the following command with the task ID and the new description:
```bash
go run . update <task_id> "New task description"
```

Example:
```bash
go run . update 1 "Updated task description"
```

This will update the task with ID `1` to have the new description `"Updated task description"`.

### 6. **Mark a Task as In-Progress, Done, or To-Do**

You can mark a task with a specific status. The possible statuses are `done`, `in-progress`, or `to-do`. Use the following commands to mark tasks:

- **Mark Task as Done**:
  ```bash
  go run . mark-as-done <task_id>
  ```

Example:
```bash
go run . mark-as-done 2
```

This will mark the task with ID `2` as `done`.

- **Mark Task as In-Progress**:
  ```bash
  go run . mark-in-progress <task_id>
  ```

Example:
```bash
go run . mark-in-progress 1
```

This will mark the task with ID `1` as `in-progress`.

## File Structure

```
TaskTrackerCLI/
│
├── main.go          # Entry point for the CLI application
├── tasks.go         # Task data structure and logic
├── commands.go      # Command handling logic
└── README.md        # Project documentation
```

## Notes
- Task IDs are automatically assigned when tasks are created.
- Task statuses include `to-do`, `in-progress`, and `done`.

## Contribution

Feel free to contribute to this project by forking the repository and submitting a pull request. You can also report issues or request new features by opening an issue on GitHub.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

Happy task tracking! 🎉
