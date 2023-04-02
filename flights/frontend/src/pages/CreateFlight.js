import React, {useState, useEffect} from 'react'
import axios from 'axios'
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

export default function CreateFlight() {

    



    const [flight, setFlight] = useState({
        departure: "",
        origin: {address: "", city:"", country:""},
        destination: {address: "", city:"", country:""},
        price: "",
        availableSeats: ""
        
    });
    
    const onInputChange = (e) => {
        if (e.target.name == "city" || e.target.name == "country" || e.target.name == "address"){
            
            setFlight({...flight, destination: {...flight.destination, [e.target.name]: e.target.value}})}
        else if (e.target.name == "origin.city" || e.target.name == "origin.country" || e.target.name == "origin.address"){
            
            let text = e.target.name;
            const name = text.split(".")[1]
            setFlight({...flight, origin: {...flight.origin, [name]: e.target.value}})
        }
        else {setFlight({...flight, [e.target.name]: e.target.value});}
        
    }

    const onSubmit = async (e) => {
        e.preventDefault();
       
        // flight.departure = new Date(flight.departure);
        let departureDate = new Date(flight.departure);
        departureDate.setHours(2, 0, 0, 0);
        flight.departure = departureDate.toISOString();
        flight.availableSeats = parseInt(flight.availableSeats);
        flight.price = parseFloat(flight.price);
        
        await axios.post("http://localhost:8084/createFlight", flight);
        toast.success('Successfully created flight !', {
            position: toast.POSITION.TOP_RIGHT
        });
        
    };
    

  return (
    
    <div className="container">
        
      <div className="row">
        <div className="col-md-6 offset-md-3 border rounded p-4 mt-2 shadow">
          <h2 className="text-center m-4">Create a Flight</h2>
          <form  onSubmit={(e) => onSubmit(e)} >
            <div className="mb-3">
              <label htmlFor="Departure" className="form-label">
                Departure:
              </label>
              <input
                type={"datetime-local"}
                className="form-control"
                name="departure"  
                required
                onChange={(e) => onInputChange(e)} 
              />
            </div>
            <div className="mb-3">
              <label htmlFor="Origin" className="form-label">
                Origin:
              </label>
              <input
                type={"text"}
                className="form-control"
                placeholder="Enter address"
                name="origin.address" 
                required
                onChange={(e) => onInputChange(e)}    
              />
              <input
                type={"text"}
                className="form-control"
                placeholder="Enter city"
                name="origin.city" 
                required
                onChange={(e) => onInputChange(e)}    
              />
              <input
                type={"text"}
                className="form-control"
                placeholder="Enter country"
                name="origin.country" 
                required
                onChange={(e) => onInputChange(e)}    
              />
            </div>
            <div className="mb-3">
              <label htmlFor="Destination" className="form-label">
                Destination:
              </label>
              <input
                type={"text"}
                className="form-control"
                placeholder="Enter address"
                name="address"               
                required
                onChange={(e) => onInputChange(e)}            
              />
              <input
                type={"text"}
                className="form-control"
                placeholder="Enter city"
                name="city"               
                required
                onChange={(e) => onInputChange(e)}            
              />
              <input
                type={"text"}
                className="form-control"
                placeholder="Enter country"
                name="country"               
                required
                onChange={(e) => onInputChange(e)}            
              />
            </div>
            <div className="mb-3">
              <label htmlFor="Price" className="form-label">
                Price:
              </label>
              <input
                type="number" 
                min = "0"
                className="form-control"
                placeholder="Enter price"
                name="price"               
                required
                onChange={(e) => onInputChange(e)}            
              />
            </div>
            <div className="mb-3">
              <label htmlFor="Seats" className="form-label">
                Available seats:
              </label>
              <input
                type="number"
                min="0"
                className="form-control"
                placeholder="Enter available seats"
                name="availableSeats"               
                required
                onChange={(e) => onInputChange(e)}            
              />
            </div>
            
            
           
            <button type="submit" className="btn btn-outline-primary">
              Submit
            </button>
          </form>
        </div>
      </div>
    </div>
  );
}