import { ReactNode, createContext, useState, useContext } from "react";

const authContext = createContext<AuthContextType | null>(null);

interface Props {
  children: ReactNode;
}

interface AuthContextType {
  Signin: any;
  Signout: any;
  user: string;
}

let username: string;

export function AuthProvider({ children }: Props) {
  const auth = useProvideAuth();

  return <authContext.Provider value={auth}> {children} </authContext.Provider>;
}

export function useAuth() {
  return useContext(authContext);
}

function useProvideAuth() {
  const [user, setUser] = useState(username);

  const Signin = (username: string) => {
    setUser(username);
  };

  const Signout = () => {
    setUser("");
  };

  return {
    Signin,
    Signout,
    user,
  };
}

export function setUser(user: string) {
  username = user;
}
