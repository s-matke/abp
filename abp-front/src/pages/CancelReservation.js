import React, { useState, useEffect } from 'react'
import axios from 'axios'
import Button from 'react-bootstrap/Button';

function CancelReservation() {
    const [reservation, setReservation] = useState([])
    const id = JSON.parse(localStorage.getItem('userID'))

    useEffect(() => {
        axios.get(`http://localhost:8000/reservation/guest/${id}`)
        .then(res => {
            console.log(res.data)
            setReservation(res.data.reservations)
        })
        .catch(error => {
            console.log(error)
        })
    }, [])

    const cancelReservation = (id) => 
    {
        console.log(id)
        axios.get(`http://localhost:8000/reservation/cancel/${id}`)
        .then (res => {
            console.log(res.data)
        })
        .then(error => {
            console.log(error)
        })
    }
  return (
    <div>
        <div style={{marginLeft:'650px'}}>
        <table class="table">
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
          <th><Button onClick={()=> cancelReservation(r.id)}>Cancel</Button></th>
        </tr>
      ))}
    </tbody>
  </table>
        </div>
    </div>
  )
}

export default CancelReservation