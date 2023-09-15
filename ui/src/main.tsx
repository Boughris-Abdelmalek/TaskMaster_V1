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
import { StartOptions } from "msw";
import { worker } from "./mocks/browser.ts";

if (import.meta.env.MODE === "development") {
  worker.start({
    onUnhandledRequest: "bypass", // Non mapped request, we'll simply ignore
  } as StartOptions);
}

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <Provider store={store}>
      <App />
    </Provider>
  </React.StrictMode>,
);
