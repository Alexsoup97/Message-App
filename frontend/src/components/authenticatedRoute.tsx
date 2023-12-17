import { Navigate, Outlet } from "react-router-dom";
import { useAuth } from "../utils/AuthProvider";

export function AuthenticatedRoutes() {
  const auth = useAuth();
  if (auth !== null && auth.user !== "") {
    return <Outlet />;
  } else {
    return <Navigate to="/SignIn" />;
  }
}
