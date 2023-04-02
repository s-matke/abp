import React, {useState, useEffect} from 'react'
import axios from 'axios'
import 'bootstrap/dist/css/bootstrap.css'
import Logo from '../img/Logo.jpg'
import { useLocation, useNavigate } from 'react-router-dom'



function ShowUserTickets({params}) {

    const[tickets, setTickets] = useState([])
    const {state} = useLocation()
    const [userId, setUserId] = useState()

    console.log(state)
    const navigate = useNavigate()

    const fetchTickets = (id) => {
        axios.get('http://localhost:8084/showTickets/' + id)
        .then(response => {
            console.log(response.data)
            if (response.data !== null) {
                setTickets(response.data)
            }
        })
        .catch(error => {
            console.log(error)
        })
    }

    useEffect(() => {
        const role = localStorage.getItem('user') || 'guest';
        if (role !== "guest") {
            let parsedUser = JSON.parse(role)
            // setUserRole(parsedUser['Role'])
            setUserId(parsedUser["ID"])
            fetchTickets(parsedUser["ID"]);
        } else {
            // setUserRole(role)
            navigate("/flight/search")
            navigate(0)
        }

        
    },[])

    return (
        <div class="row mx" style={{marginLeft: '300px'}}>
           <div style={{textAlign : 'center'}}>
           <h1>Your tickets</h1> 
           </div>
            {
                tickets.map((item) => (
                <div class="card" style={{width: '18rem',
                                          margin : 10 }}>
                    <img src={Logo} class="card-img-top" alt="logo" style={{height:'200px',width:'285px'}}/>
                        <div class="card-body">
                            <h5 class="card-title">From: {item.Origin.country}</h5>
                            <h5 class="card-title">To: {item.Destination.country}</h5>
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