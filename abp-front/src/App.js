import logo from './logo.svg';
import 'react-toastify/dist/ReactToastify.css';
import {BrowserRouter, Routes, Route} from "react-router-dom";
import { ToastContainer, toast } from 'react-toastify';

import Sidebar from "./components/Sidebar/Sidebar";
import Register from './pages/Register';
import CreateAccommodation from './pages/CreateAccommodation';
import UpdateUser from './pages/UpdateUser';
import ShowAccommodation from './pages/ShowAccommodation';

import Login from './pages/Login';
import CancelReservation from './pages/CancelReservation';
import { AuthProvider } from './context/AuthContext';

// import Accommodation from './pages/Accommodation';
import ShowReservation from './pages/ShowReservation';
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
          <Route path="/userupdate" element={<UpdateUser/>}></Route>        
          <Route path="/signin" element={<Login />}></Route> 
          {/*<Route path="/accommodation" element={<Accommodation />}></Route> */}
          <Route path="/accommodation/search" element={<Accommodation />}></Route> 
          {/* <Route path="/" element={<Signin />}></Route> */}
          <Route path="/accommodation" element={<ShowAccommodation/>}></Route>
          <Route path="/accommodation/create" element={<CreateAccommodation/>}></Route>
          <Route path="/reservation" element={<ShowReservation/>}></Route>
          <Route path="/reservation/cancel" element={<CancelReservation/>}></Route>
        </Routes>
        </AuthProvider>
      </BrowserRouter>
      <ToastContainer/>
    </>
  );
}


export default App;
