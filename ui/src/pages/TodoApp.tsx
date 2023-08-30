import { FC, useEffect } from "react";
import BaseLayout from "../components/Layout.tsx";
import { useGetTodosQuery } from "../store/reducers/todo-api-slice.ts";
import { useAppDispatch, useAppSelector } from "../store/hooks.ts";
import { selectAllTodos, setAllTodos } from "../store/reducers/todo-slice.ts";

const TodoApp: FC = () => {
  const { data } = useGetTodosQuery();
  const dispatch = useAppDispatch();
  const todos = useAppSelector(selectAllTodos);

  useEffect(() => {
    if (data) {
      dispatch(setAllTodos(data));
    }
  }, [data]);

  return (
    <BaseLayout>
      <h3>Todos for today</h3>
      <ul>
        {todos &&
          Object.values(todos).map((todo) => {
            return <li key={todo.id}>{todo.title}</li>;
          })}
      </ul>
    </BaseLayout>
  );
};

export default TodoApp;
