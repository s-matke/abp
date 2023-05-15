import logo from './logo.svg';
import 'react-toastify/dist/ReactToastify.css';
import {BrowserRouter, Routes, Route} from "react-router-dom";
import { ToastContainer, toast } from 'react-toastify';

import Sidebar from "./components/Sidebar/Sidebar";
import Register from './pages/Register';
import CreateAccommodation from './Pages/CreateAccommodation';
import UpdateUser from './Pages/UpdateUser';
import ShowAccommodation from './Pages/ShowAccommodation';

import Login from './pages/Login';
import { AuthProvider } from './context/AuthContext';
import Accommodation from './pages/Accommodation';

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
          <Route path="/signin" element={<Login />}></Route> 
          <Route path="/accommodation/search" element={<Accommodation />}></Route> 
          {/* <Route path="/" element={<Signin />}></Route> */}
          <Route path="/createAccommodation" element={<CreateAccommodation/>}></Route>
          <Route path="/accommodation" element={<ShowAccommodation/>}></Route>
          <Route path="/accommodation/create" element={<CreateAccommodation/>}></Route>
        </Routes>
        </AuthProvider>
      </BrowserRouter>
      <ToastContainer/>
    </>
  );
}


export default App;
