import { createBrowserRouter } from "react-router-dom"
import { SignIn } from "../pages/SignIn"
import { Signup } from "../pages/Signup"
import { AuthenticatedRoutes } from "../components/authenticatedRoute"

export const router = createBrowserRouter([
    {
      path: "/Signin",
      element: <SignIn/>
    },{
      path: "/Signup",
      element: <Signup/>
    },{
        path: "/",
        element: <AuthenticatedRoutes/>,
        children: [

        ]
    }] )

