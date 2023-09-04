import { BrowserRouter, Routes, Route } from "react-router-dom";
import LandingPage from "./pages/LandinPage.tsx";
import TodoApp from "./pages/TodoApp.tsx";
import SignIn from "./pages/SignIn.tsx";
import SignUp from "./pages/SignUp.tsx";
import { FC } from "react";
import { ThemeProvider } from "@mui/material";
import theme from "./material-theme.ts";

const App: FC = () => (
  <ThemeProvider theme={theme}>
    <BrowserRouter>
      <Routes>
        <Route index path={"/"} element={<LandingPage />} />
        <Route path={"/todo-app"} element={<TodoApp />} />
        <Route path={"/sign-in"} element={<SignIn />} />
        <Route path={"/sign-up"} element={<SignUp />} />
      </Routes>
    </BrowserRouter>
  </ThemeProvider>
);

export default App;
