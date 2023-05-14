import {React, useEffect, useState} from 'react'
import axios from 'axios'
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';

function ShowAccommodation() {
   
    const [accommodation, setAccommodation] = useState([])
    const [showInputs, setShowInputs] = useState(false)
    const [price, setPrice] = useState({})
    useEffect(() => {
        axios.get('http://localhost:8000/accommodation')
        .then(res => {
            
            setAccommodation(res.data.accommodations)
            console.log(res.data.accommodations)
            
        })
        .then(error => {
            console.log(error)
        })
    }, [])

    const handleChange = (name) => (event) => {
        const val = event.target.value;
        setPrice({...price, [name] : val})
        
      };

    const handlePrice = (id) => {
        setShowInputs(true)
        setPrice({...price, ["accommodationId"] : id})
       
    }
    
    const createPrice  = () => {
        axios.post('http://localhost:8000/pricing',
        {
            "accommodationId": price.accommodationId,
        "price": price.prc,
        "season": {
            "springMultiplier": price.spring,
            "summerMultiplier": price.summer,
            "fallMultiplier": price.fall,
            "winterMultiplier":price.winter
        },
        "week": {
            "workdayMultiplier": price.workday,
            "weekendMultiplier": price.weekend
        },
        "holiday": price.holiday,
        "pricingType": price.type
    })
    .then(res => {
        console.log(res.data)
    })
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
          <th><Button onClick={() => handlePrice(a.id)}>Click me</Button></th>
        </tr>
      ))}
    </tbody>
  </table>
        {showInputs && <div style={{width:'30%'}}>
            <Form>
            <Form.Group>
               <Form.Label>Price</Form.Label> 
            <Form.Control  type="text" value={price.prc} onChange={handleChange("prc")}/>
            </Form.Group>
            <Form.Group>
               <Form.Label>Holiday</Form.Label> 
            <Form.Control  type="text" value={price.holiday} onChange={handleChange("holiday")}/>
           </Form.Group> 
           <Form.Group>
               <Form.Label>Type</Form.Label> 
            <Form.Control type="text" value={price.type} onChange={handleChange("type")} placeholder='PER_PERSON/PER_HOUSEHOLD'/>
           </Form.Group> 
           <Form.Group>
               <Form.Label>Summer</Form.Label> 
            <Form.Control type="text" value={price.summer} onChange={handleChange("summer")}/>
          </Form.Group> 
          <Form.Group>
               <Form.Label>Winter</Form.Label> 
            <Form.Control type="text" value={price.winter} onChange={handleChange("winter")}/>
         </Form.Group>
         <Form.Group>
               <Form.Label>Spring</Form.Label>   
            <Form.Control type="text" value={price.spring} onChange={handleChange("spring")}/>
          </Form.Group>
          <Form.Group>
               <Form.Label>Fall</Form.Label> 
            <Form.Control type="text" value={price.fall} onChange={handleChange("fall")}/>
          </Form.Group> 
          <Form.Group>
               <Form.Label>Weekend</Form.Label> 
            <Form.Control type="text" value={price.weekend} onChange={handleChange("weekend")}/>
          </Form.Group>
          <Form.Group>
               <Form.Label>Workday</Form.Label>  
            <Form.Control type="text" value={price.workday} onChange={handleChange("workday")}/>
            </Form.Group>
            </Form>
            <Button variant="primary" onClick={createPrice}>Click</Button>
            </div>}

        </div>
    </div>
  )
}

export default ShowAccommodation