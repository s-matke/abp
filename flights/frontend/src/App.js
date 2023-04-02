
import {BrowserRouter, Routes, Route} from "react-router-dom";
import Home from "./pages/Home";
import SearchFlights from "./pages/SearchFlights";
import Signup from "./pages/signup/Signup";
import Sidebar from "./components/Sidebar/Sidebar";

function App() {
  return (
    <>
      <BrowserRouter>
        <Sidebar />
        <Routes>
          <Route path="/" element={<Home />}></Route>
          <Route path="/flight/search" element={<SearchFlights />}></Route>
          <Route path="/signup" element={<Signup />}></Route>
        </Routes>
      </BrowserRouter>
    </>
  );
}

export default App;
