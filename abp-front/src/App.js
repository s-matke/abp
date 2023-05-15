import logo from './logo.svg';
import 'react-toastify/dist/ReactToastify.css';
import {BrowserRouter, Routes, Route} from "react-router-dom";
import { ToastContainer, toast } from 'react-toastify';

import Sidebar from "./components/Sidebar/Sidebar";
import Register from './Pages/Register';
import CreateAccommodation from './Pages/CreateAccommodation';
import UpdateUser from './Pages/UpdateUser';
import ShowAccommodation from './Pages/ShowAccommodation';
import CancelReservation from './Pages/CancelReservation';
import Login from './Pages/Login';
import { AuthProvider } from './context/AuthContext';

// import Accommodation from './pages/Accommodation';
import ShowReservation from './Pages/ShowReservation';
import Accommodation from './Pages/Accommodation';

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
