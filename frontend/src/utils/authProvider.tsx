import { createContext} from "react";

export const AuthContext = createContext<undefined>(undefined)

// function AuthProvider({children}){
//     const [user, setUser] = useState("")
//     return <AuthContext.Provider value={}> {children} </AuthContext.Provider>
// }

