import {React, useEffect, useState} from 'react'
import axios from 'axios'
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import DatePicker from 'react-datepicker';
import 'react-datepicker/dist/react-datepicker.css';
import APIService from '../services/APIService';

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
    const [showDate, setShowDate] = useState(false);
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
        
      const [startDate, setStartDate] = useState(null);
      const [endDate, setEndDate] = useState(null);
      const [a_id, setAid] = useState(null);
      const handleSubmitClick= () => {

        async function availlability() {
          try {
              const response = await APIService.Availlabillity(a_id, startDate, endDate);
              alert("USPESNO")
              setShowDate(false)
              
              console.log(response);
          } catch (error) {
              console.error(error);
          }
      }
      availlability();
      }

    const handleShow= (aa_id) => {
        setAid(aa_id)
        setShowDate(true)
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
          <th><Button onClick={() => handlePrice(a.id)}>Click me</Button></th>
          {!showDate && <th><Button onClick={() => handleShow(a.id)}>Change Availabillity</Button></th>}
          <th><Button onClick={() => handlePrice(a.id)}>Create price</Button></th>
        </tr>
      ))}
    </tbody>
    {showDate && <div>
      <h3 style={{fontSize:"18px"}}>Od kada do kada zelite da zabranite rezervaciju?</h3>
      <h4 style={{fontSize:"14px"}}>Od</h4>
      <DatePicker
      selected={startDate}
      onChange={(date) => setStartDate(date)}
    />
    <h3 style={{fontSize:"14px"}}>Do</h3>
      <DatePicker
      selected={endDate}
      onChange={(date) => setEndDate(date)}
    />
    <button style={{backgroundColor:"teal", color:"white", borderRadius:"20px", minWidth:"200px", padding:"4px 10px", margin:"10px"}}
      onClick={()=>handleSubmitClick()}
    >Potvrdi</button>
    <button style={{backgroundColor:"white", color:"teal", borderRadius:"20px", minWidth:"200px", padding:"4px 10px", border:"none"}}
      onClick={()=>setShowDate(false)}
    >Ponisti</button>
      </div>}
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