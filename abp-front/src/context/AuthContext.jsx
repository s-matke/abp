import { createContext, useState, useEffect } from "react"
import { useNavigate } from "react-router-dom";
import APIService from "../services/APIService";

const AuthContext = createContext()

export default AuthContext;



export const AuthProvider = ({children}) =>{
    const [loading, setLoading ] = useState(false);

    const handleLogin = async (e) => {
        e.preventDefault();
        console.log(e.target.username.value, e.target.password.value)
        try {
          const response = await APIService.Login(e.target.username.value, e.target.password.value);
          console.log(response);
          localStorage.setItem('userID', JSON.stringify(response.user.id))
          localStorage.setItem('userRole', JSON.stringify(response.user.role).toLowerCase())
        } catch (error) {
          console.error(error);
        }
      }
    let contextData = {
        handleLogin:handleLogin,
    }
    return(
        <AuthContext.Provider value={contextData}>
            {loading ?  null : children}
        </AuthContext.Provider>
    )
}