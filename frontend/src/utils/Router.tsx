import { createBrowserRouter } from "react-router-dom";
import { SignIn } from "../pages/registration/Signin";
import { Signup } from "../pages/registration/Signup";
import { AuthenticatedRoutes } from "../components/authenticatedRoute";
import { Dashboard } from "../pages/Dashboard";
import { loadChatData } from "./services/ChatService";
import { QueryClient } from "@tanstack/react-query";
import { ErrorPage } from "../pages/NotFound";

export const queryClient = new QueryClient();
export const router = createBrowserRouter([
  {
    path: "dashboard",
    element: <AuthenticatedRoutes />,
    children: [
      {
        path: "",
        element: <Dashboard />,
        loader: loadChatData(queryClient),
      },
    ],
  },
  {
    path: "Signin",
    element: <SignIn />,
  },
  {
    path: "Signup",
    element: <Signup />,
  },
  {
    path: "*",
    element: <ErrorPage />,
    children: [],
  },
]);
