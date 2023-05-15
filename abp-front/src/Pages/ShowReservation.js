import {React, useState, useEffect} from 'react'
import axios from 'axios'
import Button from 'react-bootstrap/Button';

function ShowReservation() {
    const [accommodation, setAccommodation] = useState([])
    const [showTable, setShowTable] = useState(false)
    const [reservation, setReservation] = useState([])
    const id = JSON.parse(localStorage.getItem('userID'))

    useEffect(() => {
        axios.get(`http://localhost:8000/accommodation/host/${id}`)
        .then(res => {
            
            setAccommodation(res.data.accommodations)
            console.log(res.data.accommodations)
            
        })
        .then(error => {
            console.log(error)
        })
    }, [])

    const handleReservation = (id) => {
        setShowTable(true)
        axios.get(`http://localhost:8000/reservation/pending/64621f6a8f37f1357a23b363`)
        .then(res => {
            
            setReservation(res.data.reservations)
            console.log(res.data.reservations)
            
        })
        .then(error => {
            console.log(error)
        })
       return id;
    }

    const cancelReservation = (id) => 
    {
        console.log(id)
        /*axios.post(`http://localhost:8000/reservation/confirm/${id}`)
        .then (res => {
            console.log(res.data)
        })
        .then(error => {
            console.log(error)
        })*/
    }
  return (
    <div>
        <div style={{marginLeft:'650px'}}>
       <table class="table">
    <thead>
      <tr>
        <th scope="col">Name</th>
      </tr>
    </thead>
    <tbody>
      {accommodation.map(a => (
        <tr key={a.id}>
          <th>{a.name}</th>
          <th><Button onClick={() => handleReservation(a.id)}>Reservation</Button></th>
        </tr>
      ))}
    </tbody>
  </table>
  { showTable && <table class="table">
    <thead>
      <tr>
        <th scope="col">Start Date</th>
        <th scope="col">End Date</th>
        <th scope="col">Number of guests</th>
        <th scope="col">Price</th>
      </tr>
    </thead>
    <tbody>
      {reservation.map(r => (
        <tr key={r.id}>
          <th>{r.startDate}</th>
          <th>{r.endDate}</th>
          <th>{r.numOfGuests}</th>
          <th>{r.price}</th>
          <th><Button onClick={()=> cancelReservation(r.id)}>Confirm</Button></th>
        </tr>
      ))}
    </tbody>
  </table>}
        </div>
    </div>
  )
}

export default ShowReservation