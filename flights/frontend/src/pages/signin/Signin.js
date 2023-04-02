import React from 'react';
import 'mdb-react-ui-kit/dist/css/mdb.min.css';
import { 
  MDBBtn, 
  MDBContainer, 
  MDBRow, 
  MDBCol, 
  MDBCard, 
  MDBCardBody, 
  MDBCardImage, 
  MDBInput, 
} 
from 'mdb-react-ui-kit';
import flight_img from "../../assets/img/21308.jpg"

import PhoneInput from 'react-phone-input-2'
import 'react-phone-input-2/lib/high-res.css'

import 'bootstrap/dist/css/bootstrap.min.css'
import { useState } from 'react';

import axios from "axios";
import { useNavigate } from 'react-router-dom';
import { Form } from 'react-bootstrap';
import { ToastContainer, toast } from 'react-toastify';


function Signin() {

    const navigate = useNavigate();
    const [user, setUser] = useState(
        {
            email: "",
            password: ""
        }
    )

    const submitUser = (e) => {
        e.preventDefault();
        axios.post(`http://localhost:8084/signin`, user)
            .then(res => {
                if (res.status === 200) {
                    toast.success('Logged in!', {
                        position: toast.POSITION.TOP_CENTER
                    });
                    localStorage.setItem('user', JSON.stringify(res.data))
                    localStorage.setItem('role', res.data.Role)
                    console.log("YOOOOOOOOOOOOO")
                    let user = localStorage.getItem('user')
                    console.log("USER:")
                    let parsedUser = JSON.parse(user)
                    console.log(parsedUser["Role"])
                    navigate("/flight/search")
                    navigate(0)
                }
            })
            .catch(error => {
                console.log(error.response.status)
                if (error.response.status === 409) {
                    alert("Wrong email!")
                }
            })
    }

    const onInputChange = (e) => {
        setUser({...user, [e.target.name]: e.target.value});
    }

    return (
    <MDBContainer fluid>

      <MDBCard className='text-black' style={{borderRadius: '25px'}}>
        <MDBCardBody>
        <MDBRow>
            <MDBCol md='10' lg='6' className='order-2 order-lg-2 d-flex flex-column align-items-center'>
                <div className="container position-relative">
                    <div className="row">
                        <div className="col-md-5 offset-md-4 border rounded p-4 mt-2 shadow position-relative">
                            <Form onSubmit={(e) => submitUser(e)}>
                                <h2 className="text-center m-5">Sign Up</h2>
                                <div className="mb-4">
                                    <div className="input-group">
                                            <MDBCol md='6' lg='12'>
                                                <MDBInput 
                                                    label='Your email' 
                                                    name="email" 
                                                    type='email'
                                                    onChange={(e) => onInputChange(e)}
                                                    required/>
                                            </MDBCol>
                                    </div>
                                </div>
                                <div className="mb-4">
                                    <div className="input-group">
                                            <MDBCol md='6' lg='12'>
                                                <MDBInput 
                                                    label='Your password' 
                                                    name="password" 
                                                    type='password' 
                                                    onChange={(e) => onInputChange(e)}
                                                    required/>
                                                    
                                            </MDBCol>
                                    </div>
                                </div>
                                <div className='mb-4'>
                                    <MDBCol md='6' lg='12'>
                                        <MDBBtn 
                                            className='mb-1' 
                                            size='lg' 
                                            style={{width: "50%", marginLeft: "25%"}} 
                                            >Sign in</MDBBtn>
                                    </MDBCol>
                                </div>
                            </Form>
                        </div>
                    </div>
                </div>
            </MDBCol>
            <MDBCol md='10' lg='6' className='order-1 order-lg-2 d-flex align-items-center'>
              <MDBCardImage src={flight_img} fluid/>
            </MDBCol>
          </MDBRow>  
        </MDBCardBody>
      </MDBCard>
    </MDBContainer>
  );
}

export default Signin;