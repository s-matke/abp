import {useEffect, useState} from 'react'
import styled from "styled-components"
import axios from "axios";
import 'bootstrap/dist/css/bootstrap.css';
import ShowUserTickets from './ShowUserTickets';
import { BrowserRouter, Route, Routes, useNavigate } from 'react-router-dom';

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

const FlightContainer = styled.div`
    display:flex;
    align-items: center;
    justify-content: space-between;
    margin:20px;
    height: 200px;
    padding:0px;
    
`;
const FlightWrapper = styled.div`
    border:1px solid #ccc;
    display:flex;
    width: ${(props) => props.width}   ;;
    height:100%;
    padding: 0px 20px 20px 20px;
    flex-direction:  ${(props) => props.flex_direction};
    justify-content:  ${(props) => props.justify_content};
    align-items: center;

`
const FlightLeftWrapper = styled.div`
    display:flex;
    flex-direction: column;
`;
const FlightTitle = styled.h1`
    font-size: 30px;
    letter-spacing: 1.5px;
    color:teal;
    padding:0px;
    margin:0px;
`;

const FlightDeparture = styled.h3`
    font-size: 16px;
    font-weight: 700;
    padding:5px;
    margin:0px;
`;

const FlightRightWrapper = styled.div`
    display:flex;
    gap:20px;
    align-items: center;
`;
const TicketPrice = styled.h1`
    font-size: 25px;
    color:#ccc;
    font-weight: 300;
`;
const TotalPrice = styled.h1`
    font-size: 25px;
    color:black;
    font-weight: 600;
`;

const SearchFlights = () => {
  const [departure, setDeparture] = useState(new Date().toISOString().substr(0, 10));
  const [origin, setOrigin] = useState("");
  const [destination, setDestination] = useState("");
  const [availableSeats, setAvailableSeats] = useState();
  const [flights, setFlights] = useState([])
  const [NumberOfTicket,setNumberOfTicket] = useState(0);
  const navigate = useNavigate()
 
  const handleSubmit = (e) => {
    e.preventDefault();
    let departureDate = new Date(departure);
    departureDate.setHours(0, 0, 0, 0);
    let formattedDeparture = departureDate.toISOString();

    console.log(formattedDeparture,
        origin,
        destination,
        availableSeats
      )
    const options = {
      method: 'POST',
      url: 'http://localhost:8084/searchFlights',
      headers: {
        Authorization: 'Token 9b85c3784a75fa85e8ae3fc58ca69ff9136af186',
        'Content-Type': 'application/json'
      },
      data: {
        "departure":formattedDeparture,
        "origin":origin,
        "destination":destination,
        "availableSeats":parseInt(availableSeats)
      }
    };

    axios.request(options)
      .then(function (response) {
        console.log(response.data);
        setFlights(response.data);
      })
      .catch(function (error) {
        console.error(error);
      });
  }


const handleClick = index =>
{
  if (NumberOfTicket <= 0)
    {
      alert("Incorrect input!")
      return
    }
  else
    {
      axios.post('http://localhost:8084/buyTicket',
        {
          "IdFlight":flights[index]._id,
          "flight":
          {
          "Origin":{"Country":flights[index].origin.country,"City":flights[index].origin.city,"Address":flights[index].origin.address},
          "Destination":{"Country":flights[index].destination.country,"City":flights[index].destination.city,"Address":flights[index].destination.address},
          "Price":flights[index].price,
          "AvailableSeats":flights[index].availableseats,
          "Departure" : flights[index].departure
          },
          "IdUser":"64270b757820ffdc26d3506b",
          "NumberOfTickets":parseInt(NumberOfTicket)
        })

  navigate('/ShowUserTickets', {state : {id:"64270b757820ffdc26d3506b"}})

    }
}
  

  
  return (
    <Container>
        <Title>Pretrazi letove po sledecim parametrima:</Title>
        <Form onSubmit={handleSubmit}>
            <LabelInputDiv>
        <Label htmlFor="departure">Datum:</Label>
        <Input
          type="date"
          id="departure"
          value={departure}
          onChange={(e) => setDeparture(e.target.value)}
          required
        />
</LabelInputDiv>
<LabelInputDiv>
        <Label htmlFor="origin">Polazno Mesto:</Label>
        <Input
          type="text"
          id="origin"
          value={origin}
          onChange={(e) => setOrigin(e.target.value)}
          //required
        />
</LabelInputDiv>
<LabelInputDiv>
        <Label htmlFor="destination">Destinacija:</Label>
        <Input
          type="text"
          id="destination"
          value={destination}
          onChange={(e) => setDestination(e.target.value)}
          //required
        />
</LabelInputDiv>
<LabelInputDiv>
        <Label htmlFor="availableSeats">Slobodnih Mesta:</Label>
        <Input
          type="number"
          id="availableSeats"
          value={availableSeats}
          onChange={(e) => setAvailableSeats(e.target.value)}
          //required
        />
</LabelInputDiv>
        <Button type="submit">Pretrazi</Button>
      </Form>
      {flights?.map((item,index)=>{
        
          const departureDate = item.departure.substring(0, 10);
          const departureTime = item.departure.substring(11, 19);
          
          return <FlightContainer key={index}>
          <FlightWrapper flex_direction="normal" justify_content="space-around" width="70%">
              <FlightLeftWrapper>
              <FlightTitle>{item.origin.city}</FlightTitle>
              <FlightDeparture>{departureDate}</FlightDeparture>
              <FlightDeparture>{departureTime}</FlightDeparture>
              </FlightLeftWrapper>
              <FlightLeftWrapper>
                  <FlightTitle>{item.destination.city}</FlightTitle>
              </FlightLeftWrapper>
          </FlightWrapper>
          <FlightWrapper  flex_direction="column" justify_content="center" width="30%">
              <FlightRightWrapper>
                  <FlightTitle>Cena:</FlightTitle>
              <TicketPrice>{item.price}.O EUR</TicketPrice>
              </FlightRightWrapper>
              <FlightRightWrapper>
                  <FlightTitle> Ukupna cena: </FlightTitle>
                {item.totalPrice===0 ? <TotalPrice>-</TotalPrice> : <TotalPrice>{item.totalPrice}.O EUR</TotalPrice>}
              </FlightRightWrapper>
             <FlightRightWrapper>
              <LabelInputDiv>
                <Label htmlFor='origin'>Enter the number of ticket:</Label>
                  <Input type = 'number' min = '1' onChange = {e => setNumberOfTicket(e.target.value)}/>
              </LabelInputDiv>
                <button type = 'button' class = 'btn btn-danger' onClick = {() => handleClick(index)} value = {NumberOfTicket} 
                 >Buy ticket</button>
              </FlightRightWrapper> 
          </FlightWrapper>
          </FlightContainer>
          ;
      })}
    </Container>
  )
}

export default SearchFlights