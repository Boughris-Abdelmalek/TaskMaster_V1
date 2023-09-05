import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App.tsx";
import { store } from "./store/store.ts";
import { Provider } from "react-redux";
import "@fontsource/rowdies/300.css";
import "@fontsource/rowdies/400.css";
import "@fontsource/rowdies/700.css";
// Supports weights 300-700
import "@fontsource-variable/rosario";

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <Provider store={store}>
      <App />
    </Provider>
  </React.StrictMode>,
);
