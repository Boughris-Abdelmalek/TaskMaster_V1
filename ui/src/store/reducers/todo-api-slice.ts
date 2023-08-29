import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";
import { Todo } from "../../types/todo.ts";
import { BASE_URL, ENDPOINT_TODOS } from "../../constants/endpoints.ts";

export const todoApi = createApi({
  reducerPath: "todoApi",
  baseQuery: fetchBaseQuery({ baseUrl: BASE_URL }),
  endpoints: (builder) => ({
    getTodos: builder.query<Todo[], void>({
      query: () => ENDPOINT_TODOS,
    }),
    getTodosById: builder.query<Todo, number>({
      query: (id) => `${ENDPOINT_TODOS}/${id}`,
    }),
  }),
});
export const { useGetTodosQuery } = todoApi;
