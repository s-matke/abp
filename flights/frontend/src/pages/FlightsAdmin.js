import React, {useState, useEffect} from 'react'
import axios from 'axios'
import { Button } from 'bootstrap';

export default function FlightsAdmin() {
    
    const [flights, setFlights] = useState([]);

  useEffect(() => {
    axios.get('http://localhost:8084/showFlights')
      .then(response => {
        setFlights(response.data);
      })
      .catch(error => {
        console.log(error);
      });
  }, []);
    
    /*const loadFlights=async() => {
        const result=await axios.get('http://localhost:8084/showFlights')
        setFlights(result.data);

    }*/
    
    return (
        <div className="container">
            <label>Letovi</label>
            <div className='py-4' style={{'width':'75%','margin-left':'280px'}}>
            <table className="table table-hover border rounded p-4 mt-2 shadow table-striped" style={{'cursor': 'pointer'}}>
      <thead>
        <tr>
          <th>Departure</th>
          <th>Origin</th>
          <th>Destination</th>
          <th>Price</th>
        </tr>
      </thead>
      <tbody>
        {flights.map(flight => (
          <tr key={flight.id}>
            <td>{flight.departure}</td>
            <td>{JSON.stringify(flight.origin)}</td>
            <td>{JSON.stringify(flight.destination)}</td>
            <td>{flight.price}</td>
          </tr>
        ))}
      </tbody>
    </table>
            </div>
        </div>
    );
}