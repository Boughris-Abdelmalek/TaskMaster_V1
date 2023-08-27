import {useEffect, useState} from "react";

const App = () => {
    const [todos, setTodos] = useState(null)
    const fetchTodos = async () => {
        const response = await fetch("http://localhost:8080/todos");
        const data = await response.json();

        setTodos(data);
    }

    useEffect(() => {
        fetchTodos();
    }, []);

    if (todos) {
        console.log(todos)
    }

    return (
        <h1>Hello World</h1>
    )
}

export default App