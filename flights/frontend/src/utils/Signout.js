import { useEffect } from "react";
import { useNavigate } from "react-router";

function Signout() {
    // console.log(localStorage.getItem('user'))
    localStorage.clear();
    // window.location.href = '/';

    const navigate = useNavigate();
    useEffect(() => {
        localStorage.clear();
        // window.location.href = '/';
        navigate("/signin")
        navigate(0)
    },[])

    return(
        <div></div>
    )
    
}

export default Signout;
