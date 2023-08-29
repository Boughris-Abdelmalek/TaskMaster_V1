import { useGetTodosQuery } from "./store/reducers/todo-api-slice.ts";
import { useEffect } from "react";
import { useAppDispatch, useAppSelector } from "./store/hooks.ts";
import { selectAllTodos, setAllTodos } from "./store/reducers/todo-slice.ts";
import BaseLayout from "./components/Layout.tsx";

const App = () => {
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
      <ul>
        {todos &&
          Object.values(todos).map((todo) => {
            return <li key={todo.id}>{todo.title}</li>;
          })}
      </ul>
    </BaseLayout>
  );
};

export default App;
