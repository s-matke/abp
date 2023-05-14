import {useEffect, useState} from 'react'
import styled from "styled-components"
import axios from "axios";
import 'bootstrap/dist/css/bootstrap.css';
import { BrowserRouter, Route, Routes, useNavigate } from 'react-router-dom';
import { ToastContainer, toast } from 'react-toastify';
import APIService from '../services/APIService';

import DatePicker from 'react-datepicker';
import 'react-datepicker/dist/react-datepicker.css';

const Container = styled.div`
margin-left: 300px;
margin-top: 40px;
`;

const Title = styled.h1`
padding:0px 30px;
margin: 20px 0px 0px 0px;
letter-spacing: 1.4px;
color: #1a6666;

`;


const Form = styled.form`
  display: flex;
  gap: 30px;
  flex-wrap: wrap;
  justify-content: left;
  align-items: bottom;
  width: 100%;
  margin-top: 32px;
  padding:0px 30px;
`;

const Label = styled.label`
  font-size: 18px;
  font-weight: bold;
  color: teal;

`;

const Input = styled.input`
  font-size: 15px;
  padding: 8px;
  border-radius: 4px;
  border: 1px solid #ccc;
  background-color: #f1f1f1;
  width: 200px;
  font-weight: 500;
  &:active{
      border:1px solid teal;
  }
`;
const LabelInputDiv = styled.div`
    display:flex;
    flex-direction: column;
`;
const Button = styled.button`
    min-width: 200px;
    height: 40px;
  font-size: 14px;
  margin:0px;
  margin-top:14px;
  padding: 0px 16px;
  background-color: #007bff;
  color: #fff;
  border: none;
  border-radius: 30px;
  cursor: pointer;
  text-transform: uppercase;
  font-weight: 500;

  &:hover {
    background-color: #0062cc;
  }
`;

const SearchAccommodation = () => {
    const [startDate, setStartDate] = useState(null);
    const [endDate, setEndDate] = useState(new Date().toISOString().substr(0, 10));
  //const [departureEnd, setDepartureEnd] = useState(new Date().toISOString().substr(0, 10));
  const [destination, setDestination] = useState("");
  const [availableSeats, setAvailableSeats] = useState();
  const [accommodations, setAccommodations] = useState([])


  const navigate = useNavigate();

  useEffect(() => {
    async function fetchData() {
        try {
            const response = await APIService.GetAllAccommodations();
            console.log(response.accommodations);
            setAccommodations(response.accommodations);
        } catch (error) {
            console.error(error);
        }
    }
    fetchData();
}, []);

  

const handleSubmit = (e) => {
    e.preventDefault();
    let departureDate = new Date(startDate);
    departureDate.setHours(2, 0, 0, 0);
    let formattedDeparture = departureDate.toISOString().substring(0, 10);;

    let departureDateEnd = new Date(endDate);
    departureDateEnd.setHours(2, 0, 0, 0);
    let formattedDepartureEnd = departureDateEnd.toISOString().substring(0, 10);;

    console.log(formattedDeparture,formattedDepartureEnd);

    async function search(destination, availableSeats) {
        try {
            const response = await APIService.SearchAccommodations(destination, parseInt(availableSeats));
            console.log(response.accommodations);
            setAccommodations(response.accommodations);
        } catch (error) {
            console.error(error);
        }
    }
    search(destination, availableSeats);
}
 
const excludedDates = ['2023-05-15', '2023-05-20', '2023-05-25'];


const isDateExcluded = (date) => {
  const formattedDate = date.toISOString().slice(0, 10);
  return excludedDates.includes(formattedDate);
};

const filterExcludedDates = (date) => {
  return !isDateExcluded(date);
};

  
  return (
    <Container>
        <Title>Accommodations</Title>
        <Form onSubmit={e=>handleSubmit(e)}> {/* onSubmit={} */}
            <LabelInputDiv>
        <Label htmlFor="departure">Od Kog Datuma:</Label>
        <DatePicker
      selected={startDate}
      onChange={(date) => setStartDate(date)}
      filterDate={filterExcludedDates}
    />
</LabelInputDiv>
<LabelInputDiv>
<Label htmlFor="departure">Do Kog Datuma:</Label>
<DatePicker
      selected={startDate}
      onChange={(date) => setStartDate(date)}
      filterDate={filterExcludedDates}
    />
</LabelInputDiv>
<LabelInputDiv>
        <Label htmlFor="destination">Gde idete?</Label>
        <Input
          type="text"
          id="destination"
          value={destination}
          onChange={(e) => setDestination(e.target.value)}
          //required
        />
</LabelInputDiv>
<LabelInputDiv>
        <Label htmlFor="availableSeats">Koliko osoba?</Label>
        <Input
          type="number"
          id="availableSeats"
          value={availableSeats}
          onChange={(e) => setAvailableSeats(e.target.value)}
          //required
        />
</LabelInputDiv>
        <Button type="submit">Search</Button>
      </Form>
      {accommodations?.map((item,index)=>{
          
          return <div key={index}
                style={{
                    display:"flex",
                    justifyContent:"center",
                    alignItems:"center",
                    width:"800px",
                    height:"200px",
                    backgroundColor:"#ccc",
                    margin:"10px 0px",

                }}
          >
              <div style={{backgroundColor:"#ccc", flex:"1", width:"100%", height:"100%"}}><img alt="slika"/></div>
              <div style={{backgroundColor:"#c0bebe", flex:"1",width:"100%", height:"100%"}}>
                  <h3 style={{margin:"20px"}}>{item.location.city} ({item.location.address})</h3>
                  </div>
              <div style={{backgroundColor:"#e7e7e7", flex:"1",width:"100%", height:"100%"}}>sekcija 3</div>
              
          </div>
          ;
      })}
    </Container>
  )
}

export default SearchAccommodation