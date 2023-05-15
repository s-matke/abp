import axios from "axios";
import React, { useEffect,useState } from "react";
import { Link, useNavigate } from "react-router-dom";

export default function UpdateUser() {
  const [user, setUser] = useState({
    id : "",
    name: "",
    surname: "",
    email: "",
    password: "",
    username: "",
    phoneNumber: "",
    role: "",
    address: "",
    city: "",
    country: ""
  });

  const [first, setFirst] = useState(true);
  const [disabled, setDisabled] = useState(true);

  const navigate = useNavigate();

  const loadUser = async () => {
    const userId = JSON.parse(localStorage.getItem('userID'));
    if (!userId) {
        console.log("User ID is not set!");
        return;
    }
    

    const result = await axios.get(`http://localhost:8000/user/` + userId);

    setUser(result.data.user);
    
  };

  useEffect(() => {
    if (!localStorage.getItem('userID')){
      navigate("/signin")
    }
    
    loadUser();
    
   
  }, []);


  const onSubmit = async (e) => {
    e.preventDefault();
    try{
        console.log(user.id);
    await axios.put(`http://localhost:8000/user`, user);
    setDisabled(true);
    } catch (err) {
        console.log(err);
    }
  };

  function enableFields(){
    setDisabled(false)
  }

  const onInputChange = (e) => {
    setUser({ ...user, [e.target.name]: e.target.value });
  };
  
  
  
  return (
    <div className="container mt-5">
      <div className="row">
        <div className="col-md-6 offset-md-3 border rounded p-4 mt-2 shadow">
          <h2 className="text-center m-4">User Details</h2>

          <form onSubmit={(e) => onSubmit(e)}>
            <div className="mb-3">
              <label htmlFor="Name" className="form-label">
                First Name
              </label>
              <input
                type="text"
                className="form-control"
                placeholder="Enter your name"
                name="name"
                value={user.name || ""}
                disabled={disabled}
                onChange={(e) => onInputChange(e)}
              />
            </div>
            <div className="mb-3">
              <label htmlFor="Lastname" className="form-label">
                Last Name
              </label>
              <input
                type={"text"}
                className="form-control"
                placeholder="Enter your lastname"
                name="surname"
                value={user.surname || ""}
                disabled={disabled}
                onChange={(e) => onInputChange(e)}
              />
            </div>
            <div className="mb-3">
              <label htmlFor="Email" className="form-label">
                Email
              </label>
              <input
                type={"text"}
                className="form-control"
                placeholder="Enter your email"
                name="email"
                value={user.email || ""}
                disabled={disabled}
                onChange={(e) => onInputChange(e)}
              />
            </div>
            <div className="mb-3">
              <label htmlFor="Jmbg" className="form-label">
                Password
              </label>
              <input
                type={"password"}
                className="form-control"
                placeholder="Enter new password"
                name="password"
                value={user.password || ""}
                disabled={disabled}
                onChange={(e) => onInputChange(e)}
              />
            </div>
            <div className="mb-3">
              <label htmlFor="Phone" className="form-label">
                Username
              </label>
              <input
                type={"text"}
                className="form-control"
                placeholder="Enter new username"
                name="username"
                value={user.username || ""}
                disabled={disabled}
                onChange={(e) => onInputChange(e)}
              />
            </div>
            <div className="mb-3">
              <label htmlFor="Gender" className="form-label">
                Phone
              </label>
              <input
                type={"text"}
                className="form-control"
                placeholder="Enter your phone number"
                name="phoneNumber"
                value={user.phoneNumber || ""}
                disabled={disabled}
                onChange={(e) => onInputChange(e)}
              />
            </div>
            <div className="mb-3">
              <label htmlFor="Address" className="form-label">
                Address
              </label>
              <input
                type={"text"}
                className="form-control"
                placeholder="Enter your address"
                name="address"
                value={user.address || ""}
                disabled={disabled}
                onChange={(e) => onInputChange(e)}
              />
            </div>
            <div className="mb-3">
              <label htmlFor="City" className="form-label">
                City
              </label>
              <input
                type={"text"}
                className="form-control"
                placeholder="Enter your city"
                name="city"
                value={user.city || ""}
                disabled={disabled}
                onChange={(e) => onInputChange(e)}
              />
            </div>
            <div className="mb-3">
              <label htmlFor="Contry" className="form-label">
                Contry
              </label>
              <input
                type={"text"}
                className="form-control"
                placeholder="Enter your country"
                name="country"
                value={user.country || ""}
                disabled={disabled}
                onChange={(e) => onInputChange(e)}
              />
            </div>
            <div className="mb-3">
              <label htmlFor="Role" className="form-label">
                Role
              </label>
              <input
                type={"text"}
                className="form-control"
                placeholder="Enter your role"
                name="role"
                value={user.role || ""}
                disabled={true}
                onChange={(e) => onInputChange(e)}
              />
            </div>
            <input type="button" className="btn btn-outline-primary" value="Edit" onClick={enableFields}/>
            <input type="submit" className="btn btn-outline-primary" value="Submit"/>
          </form>
        </div>
      </div>
    </div>
  )
}
