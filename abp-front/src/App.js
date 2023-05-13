import logo from './logo.svg';
import 'react-toastify/dist/ReactToastify.css';
import {BrowserRouter, Routes, Route} from "react-router-dom";
import { ToastContainer, toast } from 'react-toastify';

import Sidebar from "./components/Sidebar/Sidebar";

function App() {
  return (
    <>
      <BrowserRouter>
        <Sidebar />
        <Routes>
          {/* <Route path="/" element={<Signin />}></Route> */}
        </Routes>
      </BrowserRouter>
      <ToastContainer/>
    </>
  );
}


export default App;
