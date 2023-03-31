
import {BrowserRouter, Routes, Route} from "react-router-dom";
import Home from "./pages/Home";
import SearchFlights from "./pages/SearchFlights";
import ShowUserTickets from "./pages/ShowUserTickets";


function App() {
  return (
    <>
      <BrowserRouter>
    <Routes>
      <Route path="/" element={<Home />}></Route>
      <Route path="/search" element={<SearchFlights />}></Route>
      <Route path="/showUserTickets" element={<ShowUserTickets/>}></Route>
     
    </Routes>
    </BrowserRouter>
    </>
  );
}

export default App;
