
import {BrowserRouter, Routes, Route} from "react-router-dom";
import Home from "./pages/Home";
import SearchFlights from "./pages/SearchFlights";
import ShowUserTickets from "./pages/ShowUserTickets";
import 'react-toastify/dist/ReactToastify.css';
import { ToastContainer, toast } from 'react-toastify';

import Signup from "./pages/signup/Signup";
import Signin from "./pages/signin/Signin";
import Sidebar from "./components/Sidebar/Sidebar";
import CreateFlight from "./pages/CreateFlight";
import Signout from "./utils/Signout";

function App() {
  return (
    <>
      <BrowserRouter>
        <Sidebar />
        <Routes>
          <Route path="/" element={<SearchFlights />}></Route>
          <Route path="/flight/search" element={<SearchFlights />}></Route>
          <Route path="/signup" element={<Signup />}></Route>
          <Route path="/signin" element={<Signin />}></Route>
          <Route path="/tickets/owned" element={<ShowUserTickets/>}></Route>
          <Route path="/flight/create" element={<CreateFlight/>}></Route>
          <Route path="/signout" element={<Signout/>}></Route>
        </Routes>
      </BrowserRouter>
      <ToastContainer/>
    </>
  );
}

export default App;
