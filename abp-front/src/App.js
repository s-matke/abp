import logo from './logo.svg';
import 'react-toastify/dist/ReactToastify.css';
import {BrowserRouter, Routes, Route} from "react-router-dom";
import { ToastContainer, toast } from 'react-toastify';

import Sidebar from "./components/Sidebar/Sidebar";
import Register from './Pages/Register';
import CreateAccommodation from './Pages/CreateAccommodation';
import Login from './Pages/Login';
import { AuthProvider } from './context/AuthContext';
import UpdateUser from './Pages/UpdateUser';

function App() {
  return (
    <>
      <BrowserRouter>
      <AuthProvider>
        <Sidebar />
        <Routes>
          {/*<Route path="/signup" element={<Register />}></Route> */}
          <Route path="/signup" element={<Register />}></Route> 
          <Route path="/signin" element={<Login />}></Route>
          <Route path="/userupdate" element={<UpdateUser/>}></Route>
         
          {/* <Route path="/" element={<Signin />}></Route> */}
          <Route path="/accommodation/create" element={<CreateAccommodation/>}></Route>
        </Routes>
        </AuthProvider>
      </BrowserRouter>
      <ToastContainer/>
    </>
  );
}


export default App;
