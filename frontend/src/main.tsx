import React from "react";
import ReactDOM from "react-dom/client";
import "./index.css";
import { RouterProvider } from "react-router-dom";
import { CssBaseline, CssVarsProvider } from "@mui/joy";
import { theme } from "./utils/Theme";
import { AuthProvider, setUser } from "./utils/AuthProvider";
import { APIConstants } from "./utils/Constants";
import { QueryClientProvider } from "@tanstack/react-query";
import { ReactQueryDevtools } from "@tanstack/react-query-devtools";
import { router, queryClient } from "./utils/Router";
import axios from "axios";

const root = ReactDOM.createRoot(
  document.getElementById("root") as HTMLElement,
);

if (typeof window !== "undefined") {
  await axios
    .get(`${APIConstants.BackendUrl}/api/messages/heartbeat`, {
      withCredentials: true,
    })
    .then((resp: any) => {
      setUser(resp.user);
    })
    .catch(() => {
      setUser("");
    });
}

root.render(
  <React.StrictMode>
    <CssVarsProvider defaultMode="dark" theme={theme}>
      <QueryClientProvider client={queryClient}>
        <AuthProvider>
          <CssBaseline />
          <RouterProvider router={router} />
        </AuthProvider>
        <ReactQueryDevtools position="bottom" />
      </QueryClientProvider>
    </CssVarsProvider>
  </React.StrictMode>,
);
