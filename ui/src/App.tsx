import { useGetTodosQuery } from "./store/reducers/todo-api-slice.ts";
import { useEffect } from "react";
import { useAppDispatch, useAppSelector } from "./store/hooks.ts";
import { selectAllTodos, setAllTodos } from "./store/reducers/todo-slice.ts";

const App = () => {
  const { data } = useGetTodosQuery("");
  const dispatch = useAppDispatch();
  const todos = useAppSelector(selectAllTodos);

  useEffect(() => {
    if (data) {
      dispatch(setAllTodos(data));
    }
  }, [data]);

  return (
    <ul>
      {todos &&
        Object.values(todos).map((todo) => {
          return <li key={todo.id}>{todo.title}</li>;
        })}
    </ul>
  );
};

export default App;
