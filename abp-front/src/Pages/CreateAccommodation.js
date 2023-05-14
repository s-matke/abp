import {React, useState} from 'react'
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import axios from 'axios'

export default function CreateAccommodation() {
    
    const [accommodation, setAccommodation] = useState({})

    const handleFormInputChange = (name) => (event) => {
        const val = event.target.value;
        setAccommodation({...accommodation, [name] : val})
        
      };

      const handleUtilities = (name) => (event) => {
        const val = event.target.value;
        setAccommodation({...accommodation, [name] : val})
        
      };
      const CreateAccommodation = () => {
        const separatedVal = accommodation.utilities.split(',')
        
        var carArray = new Array();  
        for(var sv in separatedVal) {
            var jsonObj = new Object();
            jsonObj.name = separatedVal[sv];
            carArray.push(jsonObj);
          }
        axios.post('http://localhost:8000/accommodation',
        {
            hostId : JSON.parse(localStorage.getItem('userID')),
            name : accommodation.name,
            location : {
            city : accommodation.city,
            country : accommodation.country, 
            address : accommodation.address
        },
            utilities : carArray,
            minPeople : accommodation.minPeople,
            maxPeople : accommodation.maxPeople
          })
          .then(res => {
            console.log(res.data)
          }) 
          .then (error => {
            console.log(error)
          })
      };
    
  return (
    <div>
        <div style={{marginLeft:'650px',
                    width:'30%',
                    marginTop:'30px'}}>
        <h1>Create Accommodation</h1>
        <Form>
                <Form.Group className="mb-3" controlId="formBasicName">
                    <Form.Label>Name</Form.Label>
                    <Form.Control type="text" placeholder="Enter name" name="name" value={accommodation.name} onChange={handleFormInputChange("name")}/>
                </Form.Group>
                <Form.Group className="mb-3" controlId="formBasicCountry">
                    <Form.Label>Country</Form.Label>
                    <Form.Control type="text" placeholder="Enter country" name="country" value={accommodation.country} onChange={handleFormInputChange("country")}/>
                </Form.Group>
                <Form.Group className="mb-3" controlId="formBasicLastCity">
                    <Form.Label>City</Form.Label>
                    <Form.Control type="text" placeholder="Enter city" name="last_name" value={accommodation.city} onChange={handleFormInputChange("city")}/>
                </Form.Group>
                <Form.Group className="mb-3" controlId="formBasicLastAddress">
                    <Form.Label>Address</Form.Label>
                    <Form.Control type="text" placeholder="Enter address" name="address" value={accommodation.address} onChange={handleFormInputChange("address")}/>
                </Form.Group>
                <Form.Group className="mb-3" controlId="formBasicLastAddress">
                    <Form.Label>Min people</Form.Label>
                    <Form.Control type="text" placeholder="Enter min people" name="minPeople" value={accommodation.minPeople} onChange={handleFormInputChange("minPeople")}/>
                </Form.Group>
                <Form.Group className="mb-3" controlId="formBasicLastAddress">
                    <Form.Label>Max people</Form.Label>
                    <Form.Control type="text" placeholder="Enter max people" name="address" value={accommodation.maxPeople} onChange={handleFormInputChange("maxPeople")}/>
                </Form.Group>
                <Form.Group className="mb-3" controlId="formBasicName">
                    <Form.Label>Utlities</Form.Label>
                    <Form.Control type="text" placeholder="Enter utilities" name="name" value={accommodation.utilities} onChange={handleUtilities("utilities")}/>
                </Form.Group>
                <Button variant="primary" onClick={CreateAccommodation}>
                    Submit
                </Button>
        </Form>
        </div>
    </div>
  )
}
