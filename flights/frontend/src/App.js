
import {BrowserRouter, Routes, Route} from "react-router-dom";
import Home from "./pages/Home";
import SearchFlights from "./pages/SearchFlights";
import ShowUserTickets from "./pages/ShowUserTickets";
import 'react-toastify/dist/ReactToastify.css';
import { ToastContainer, toast } from 'react-toastify';

import Signup from "./pages/signup/Signup";
import Sidebar from "./components/Sidebar/Sidebar";
import CreateFlight from "./pages/CreateFlight";

function App() {
  return (
    <>
      <BrowserRouter>
        <Sidebar />
        <Routes>
          <Route path="/" element={<Home />}></Route>
          <Route path="/flight/search" element={<SearchFlights />}></Route>
          <Route path="/signup" element={<Signup />}></Route>
          <Route path="/tickets/owned" element={<ShowUserTickets/>}></Route>
          <Route path="/flight/create" element={<CreateFlight/>}></Route>
        </Routes>
      </BrowserRouter>
      <ToastContainer/>
    </>
  );
}

export default App;
