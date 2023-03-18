import React, {useState, useEffect, useCallback } from "react";

let logoutTimer;

const AuthContext = React.createContext({
    token: "",
    isLoggedIn: false,
    user: null,
    login: (token) => {},
    logout: () => {},
});

const calculateRemainingTime = (expiresIn) => {
    const currentTime = new Date().getTime();
    const expireTime = new Date().getTime() + +expiresIn;
  
    const remainingTime = expireTime - currentTime;
  
    return remainingTime;
};

const retrieveStoredToken = () => {
    const storedToken = localStorage.getItem("token");
    const expiresIn = localStorage.getItem("expires");
  
    const remainingTime = calculateRemainingTime(expiresIn);
    if (remainingTime <= 0) {
      localStorage.removeItem("token");
      localStorage.removeItem("expires");
      return null;
    }
  
    return { token: storedToken, duration: remainingTime };
};

const retriveUserFromToken = (token) => {
    const decoded = token;
  
    let user = {
      sub: decoded.sub,
      role: decoded.role[0].authority,
      id: decoded.id,
      firstTime: decoded.firstTime
    };
  
    return user;
};

 export const AuthContextProvider = (props) => {
    const tokenData = retrieveStoredToken();

    let initialToken;
    let initialUser = { role: "", id: 0 };
    if (tokenData !== null) {
        console.log(tokenData.token);
        initialToken = tokenData.token;
        initialUser = retriveUserFromToken(tokenData.token);
    }

    const [token, setToken] = useState(initialToken);
    const [user, setUser] = useState(initialUser);
    const userIsLoggedIn = !!token;

    const logoutHandler = useCallback(() => {
        setToken(null);
        setUser({ role: "", id: 0 })
        localStorage.removeItem("token")
        localStorage.removeItem("expires")

        if (logoutTimer) {
            clearTimeout(logoutTimer);
        }
    }, []);

    const loginHandler = (token) => {
        setToken(token);

        localStorage.setItem("token", token);

        setUser(retriveUserFromToken(token));

        const remainingTime = calculateRemainingTime(token.expiresIn);
        localStorage.setItem("expires", remainingTime);
        logoutTimer = setTimeout(logoutHandler, remainingTime);
    };

    useEffect(() => {
      if (tokenData) {
        logoutTimer = setTimeout(logoutHandler, tokenData.duration);
      }
    }, [tokenData, logoutHandler]);

    const contextValue = {
        token: token,
        isLoggedIn: userIsLoggedIn,
        user: user,
        login: loginHandler,
        logout: logoutHandler,
    };

    return (
        <AuthContext.Provider value={contextValue}>
            {props.children}
        </AuthContext.Provider>
    );
};


export default AuthContext;