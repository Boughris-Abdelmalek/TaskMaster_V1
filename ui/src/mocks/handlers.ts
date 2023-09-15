import { rest } from "msw";
import { BASE_URL, ENDPOINT_TODOS } from "../constants/endpoints.ts";
import { Todo } from "../types/todo.ts";

const mockTodos: Todo[] = [
  { id: 1, title: "Buy groceries", description: "Milk, eggs, bread", completed: false },
  { id: 2, title: "Go to the gym", description: "Workout for an hour", completed: true },
  { id: 3, title: "Read a book", description: "Finish 'The Great Gatsby'", completed: false },
  {
    id: 4,
    title: "Write a blog post",
    description: "About web development trends",
    completed: false,
  },
  { id: 5, title: "Clean the house", description: "Vacuum and dust", completed: true },
  { id: 6, title: "Take a walk", description: "Enjoy the park", completed: false },
  { id: 7, title: "Attend a meeting", description: "Team status update", completed: false },
  {
    id: 8,
    title: "Plan a vacation",
    description: "Choose destination and dates",
    completed: false,
  },
  { id: 9, title: "Cook dinner", description: "Make spaghetti and salad", completed: true },
  { id: 10, title: "Learn a new language", description: "Practice vocabulary", completed: false },
];

export const handlers = [
  // Mock GET request to get all todos
  rest.get(`${BASE_URL}${ENDPOINT_TODOS}`, (_req, res, ctx) => {
    return res(ctx.status(200), ctx.json(mockTodos));
  }),
];
