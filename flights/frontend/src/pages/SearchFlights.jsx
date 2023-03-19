import {useEffect, useState} from 'react'
import styled from "styled-components"
import axios from "axios";


const Container = styled.div`
    background-color: red;
`;

const Form = styled.form`
  display: flex;
  //gap: 16px;
  flex-direction: column;
  justify-content: center;
  width: 300px;
  margin-top: 32px;
  padding:20px;
`;

const Label = styled.label`
  font-size: 18px;
  font-weight: bold;
`;

const Input = styled.input`
  font-size: 16px;
  padding: 8px;
  border-radius: 4px;
  border: 1px solid #ccc;
`;

const Button = styled.button`
  font-size: 16px;
  padding: 8px 16px;
  background-color: #007bff;
  color: #fff;
  border: none;
  border-radius: 4px;
  cursor: pointer;

  &:hover {
    background-color: #0062cc;
  }
`;
const SearchFlights = () => {
    const [departure, setDeparture] = useState("");
  const [origin, setOrigin] = useState("New York");
  const [destination, setDestination] = useState( "Los Angeles");
  const [availableSeats, setAvailableSeats] = useState(1);

  const handleSubmit = (e) => {
    e.preventDefault();
    console.log(departure,
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
        departure,
        origin,
        destination,
        availableSeats
      }
    };

    axios.request(options)
      .then(function (response) {
        console.log(response.data);
      })
      .catch(function (error) {
        console.error(error);
      });
  }

  return (
    <Container>
        <Form onSubmit={handleSubmit}>
        <Label htmlFor="departure">Departure:</Label>
        <Input
          type="datetime-local"
          id="departure"
          value={departure}
          onChange={(e) => setDeparture(e.target.value)}
          required
        />

        <Label htmlFor="origin">Origin:</Label>
        <Input
          type="text"
          id="origin"
          value={origin}
          onChange={(e) => setOrigin(e.target.value)}
          required
        />

        <Label htmlFor="destination">Destination:</Label>
        <Input
          type="text"
          id="destination"
          value={destination}
          onChange={(e) => setDestination(e.target.value)}
          required
        />

        <Label htmlFor="availableSeats">Available Seats:</Label>
        <Input
          type="number"
          id="availableSeats"
          value={availableSeats}
          onChange={(e) => setAvailableSeats(e.target.value)}
          required
        />

        <Button type="submit">Search</Button>
      </Form>
    </Container>
  )
}

export default SearchFlights