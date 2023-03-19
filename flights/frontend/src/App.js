
import {BrowserRouter, Routes, Route} from "react-router-dom";
import Home from "./pages/Home";
import SearchFlights from "./pages/SearchFlights";

function App() {
  return (
    <>
      <BrowserRouter>
    <Routes>
      <Route path="/" element={<Home />}></Route>
      <Route path="/search" element={<SearchFlights />}></Route>
    </Routes>
    </BrowserRouter>
    </>
  );
}

export default App;
