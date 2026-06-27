# Entities

    - Category:
        - properties:
            title: string
            color: string
        
        - Behaviors:
            Create a new category
            Edit category details
            Delete a category
            View all categories
            Listing categories with their tasks count and status -> Single responsibility principle

    - Task:
        - properties:
            title: string
            description: string
            due_date: datetime
            is_completed: boolean
            category_id: foreign key to Category
        
        - Behaviors:
            Create a new task
            Edit task details
            Mark task as completed
            Delete a task
            View all tasks in a category
            Filter tasks by completion status
            Filter tasks by due date
            Filter tasks by date
            List tasks for today -> Single responsibility principle

    - User
        - properties:
            id: integer
            Email: string
            Password: string
        
        - Behaviors:
            Create a new user
            View user details
            Edit user details
            Delete a user
            User authentication and authorization

## UserStories

    - As a user, I want to create a new category so that I can organize my tasks.
    - As a user, I want to add a new task to a category so that I can keep track of my to-dos.
    - As a user, I want to mark a task as completed so that I can see which tasks I have finished.
    - As a user, I want to view all tasks in a category so that I can see what needs to be done.
    - As a user, I want to view all tasks count and status in a category so that I can see what needs to be done.
    - As a user, I want to edit the details of a task so that I can update its information.
    - As a user, I want to delete a task so that I can remove tasks that are no longer relevant.
    - As a user, I want to set a due date for a task so that I can manage my time effectively.
    - As a user, I want to filter tasks by their completion status so that I can focus on what still needs to be done.
    - As a user, I want to assign a color to a category so that I can visually distinguish between different categories.
    - As a user, I want to view all categories so that I can see how my tasks are organized.
    - As a user, I want to log-in to the application
