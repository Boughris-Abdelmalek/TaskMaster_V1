import { configureStore } from "@reduxjs/toolkit";
import todoReducer from "../store/reducers/todo-slice.ts";
import { todoApi } from "./reducers/todo-api-slice.ts";

export const store = configureStore({
  reducer: {
    todos: todoReducer,
    [todoApi.reducerPath]: todoApi.reducer,
  },

  middleware: (getDefaultMiddleware) => getDefaultMiddleware().concat(todoApi.middleware),

  devTools: true,
});

// Infer the `RootState` and `AppDispatch` types from the store itself
export type RootState = ReturnType<typeof store.getState>;
// Inferred type: {posts: PostsState, comments: CommentsState, users: UsersState}
export type AppDispatch = typeof store.dispatch;
