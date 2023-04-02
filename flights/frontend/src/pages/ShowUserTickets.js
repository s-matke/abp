import React, {useState, useEffect} from 'react'
import axios from 'axios'
import 'bootstrap/dist/css/bootstrap.css'
import Logo from '../img/Logo.jpg'
import { useLocation } from 'react-router-dom'



function ShowUserTickets({params}) {

    const[tickets, setTickets] = useState([])
    const {state} = useLocation()
    useEffect(() => {
        axios.get('http://localhost:8084/showTickets/')
        .then(response => {
            console.log(response)
            setTickets(response.data)
        })
        .catch(error => {
            console.log(error)
        })
    },[])

    return (
        <div class="row mx-auto" >
           <div style={{textAlign : 'center'}}>
           <h1>Your tickets</h1> 
           </div>
            {
                tickets.map((item) => (
                <div class="card" style={{width: '18rem',
                                          margin : 10 }}>
                    <img src={Logo} class="card-img-top" alt="logo" style={{height:'200px',width:'285px'}}/>
                        <div class="card-body">
                            <h5 class="card-title">From: {item.Origin.Country}</h5>
                            <h5 class="card-title">To: {item.Destination.Country}</h5>
                            <h5 class="card-title">Date: {item.Departure}</h5>
                            <h5 class="card-title">Tickets: {item.NumberOfTickets}</h5>
                            <h5 class="card-title">Price: {item.NumberOfTickets * item.Price } </h5>
                            <a href="#" class="btn btn-primary" style={{margin:25}}>Update</a>
                            <a href="#" class="btn btn-primary">Delete</a>
                        </div>
                </div>
               
                ))
            
            }
        </div>
       
      )}
      
export default ShowUserTickets;