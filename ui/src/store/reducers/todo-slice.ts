import { createSlice } from "@reduxjs/toolkit";
import { RootState } from "../store.ts";
import { TodoState } from "../../types/todo.ts";

const initialState: TodoState = {
  todos: [],
};

export const todoSlice = createSlice({
  name: "todo",
  initialState,
  reducers: {
    setAllTodos: (state, action) => {
      state.todos = action.payload;
    },
  },
});

// Action creators are generated for each case reducer function
export const { setAllTodos } = todoSlice.actions;

export default todoSlice.reducer;

// Selector functions
export const selectAllTodos = (state: RootState) => state.todos.todos;
