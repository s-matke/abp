import {React, useEffect, useState} from 'react'
import axios from 'axios'
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';

function ShowAccommodation() {
   
    const [accommodation, setAccommodation] = useState([])
    const [showInputs, setShowInputs] = useState(false)
    const [price, setPrice] = useState({})
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

    const handleChange = (name) => (event) => {
        const val = event.target.value;
        setPrice({...price, [name] : val})
        
      };

      var handleChangeSelected = (selected) => {
        setPrice({...price, ["type"] : selected.value})
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
        alert("Successfully create price!")
        window.location.assign("/accommodation")
    })
}
        
    
   
  return (
    <div>
        <div style={{marginLeft:'500px', width:'50%'}}>
        <table class="table">
    <thead>
      <tr>
        <th scope="col">Name</th>
        <th scope="col">Price</th>
      </tr>
    </thead>
    <tbody>
      {accommodation.map(a => (
        <tr key={a.id}>
          <th>{a.name}</th>
          <th><Button onClick={() => handlePrice(a.id)}>Create price</Button></th>
        </tr>
      ))}
    </tbody>
  </table>
        {showInputs && <div style={{width:'50%'}}>
            <Form>
              <div className='row'>
                <div className='col'>
            <Form.Group>
               <Form.Label>Price</Form.Label> 
            <Form.Control  type="text" value={price.prc} onChange={handleChange("prc")}/>
            </Form.Group>
            </div>
            <div className='col'>
            <Form.Group>
               <Form.Label>Holiday</Form.Label> 
            <Form.Control  type="text" value={price.holiday} onChange={handleChange("holiday")}/>
           </Form.Group> 
           </div>
           
          <div className='col'>
           <Form.Group>
               <Form.Label>Summer</Form.Label> 
            <Form.Control type="text" value={price.summer} onChange={handleChange("summer")}/>
          </Form.Group> 
          </div>
          </div>
          <div className=' row'>
            <div className='col'>
          <Form.Group>
               <Form.Label>Winter</Form.Label> 
            <Form.Control type="text" value={price.winter} onChange={handleChange("winter")}/>
         </Form.Group>
         </div>
         <div className='col'>
         <Form.Group>
               <Form.Label>Spring</Form.Label>   
            <Form.Control type="text" value={price.spring} onChange={handleChange("spring")}/>
          </Form.Group>
          </div>
          <div className='col'>
          <Form.Group>
               <Form.Label>Fall</Form.Label> 
            <Form.Control type="text" value={price.fall} onChange={handleChange("fall")}/>
          </Form.Group> 
          </div>
         </div>
         <div className='row'> 
          <div className='col'>
          <Form.Group>
               <Form.Label>Weekend</Form.Label> 
            <Form.Control type="text" value={price.weekend} onChange={handleChange("weekend")}/>
          </Form.Group>
          </div>
          <div className='col'>
          <Form.Group>
               <Form.Label>Workday</Form.Label>  
            <Form.Control type="text" value={price.workday} onChange={handleChange("workday")}/>
            </Form.Group>
            </div>
            <div className='col' style={{paddingTop:'30px'}}>
           <Form.Select className="form-select" aria-label="Default select example" value={price.type} onChange={handleChange("type")} >
              <option>Type</option>
              <option value="PER_PERSON">PER_PERSON</option>
              <option value="PER_HOUSEHOLD">PER_HOUSEHOLD</option> 
            </Form.Select>
          </div> 
            </div>
            </Form>
            <div style={{paddingTop:'30px'}}>
            <Button variant="primary" onClick={createPrice} >Submit</Button>
            </div>
            </div>
            }

        </div>
    </div>
  )
}

export default ShowAccommodation