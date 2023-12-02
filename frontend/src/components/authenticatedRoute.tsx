import {Navigate, Outlet} from 'react-router-dom'
import { useContext } from 'react'
import { AuthContext } from '../utils/authProvider'
export function AuthenticatedRoutes(){
   const isLoggedIn = useContext(AuthContext) 

    return(
       isLoggedIn ? <Outlet/> : <Navigate to ="/SignIn"/> 
    )
}

