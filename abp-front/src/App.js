import logo from './logo.svg';
import 'react-toastify/dist/ReactToastify.css';
import {BrowserRouter, Routes, Route} from "react-router-dom";
import { ToastContainer, toast } from 'react-toastify';

import Sidebar from "./components/Sidebar/Sidebar";
import Register from './pages/Register';
import CreateAccommodation from './Pages/CreateAccommodation';

function App() {
  return (
    <>
      <BrowserRouter>
        <Sidebar />
        <Routes>
          <Route path="/signup" element={<Register />}></Route> 
          {/* <Route path="/" element={<Signin />}></Route> */}
          <Route path="/createAccommodation" element={<CreateAccommodation/>}></Route>
        </Routes>
      </BrowserRouter>
      <ToastContainer/>
    </>
  );
}


export default App;
