import { createContext, useContext, useEffect, useState } from "react";
import Cookies from 'js-cookie';
import { jwtDecode } from "jwt-decode";

const AuthContext = createContext(undefined);

export const AuthProvider = ({ children }) => {

  const [user, setUser] = useState(undefined);
  const [isLoading, setIsLoading] = useState(true);

  function getUserFromCookies() {
    const jwtToken = Cookies.get("token");
    if (jwtToken) {
      const decodedToken = jwtDecode(jwtToken);
      setUser(decodedToken);
    }
    setIsLoading(false);
  }

  useEffect(() => {
    const initializeAuth = () => {
      getUserFromCookies();
    };
    initializeAuth();
  }, []);

  const login = () => {
    getUserFromCookies();
  };

  const logout = () => {
    setUser(null);
    Cookies.remove("token");
  };

  if (isLoading) {
    return null;
  }

  return (
    <AuthContext.Provider value={{ user, isLoading, login, logout }}>
      {children}
    </AuthContext.Provider>
  );
};

export const useAuth = () => {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error("useAuth must be used within an AuthProvider");
  }
  return context;
};