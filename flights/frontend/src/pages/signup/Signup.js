// import React from 'react'
// import { Col, Button, Row, Container, Card, Form } from "react-bootstrap";
// import 'bootstrap/dist/css/bootstrap.min.css';

// export default function Signup() {
//   return (
//     <div>
//       <Container>
//         <Row className="vh-100 d-flex justify-content-center align-items-center">
//           <Col md={8} lg={6} xs={12}>
//           <div className="border border-2 border-primary"></div>
//             <Card className="shadow px-4">
//               <Card.Body>
//                 <div className="mb-3 mt-md-4">
//                   <h2 className="fw-bold mb-2 text-center text-uppercase ">Logo</h2>
//                   <div className="mb-3">
//                     <Form>
//                       <Form.Group className="mb-3" controlId="Name">
//                         <Form.Label className="text-center">
//                           Name
//                         </Form.Label>
//                         <Form.Control type="text" placeholder="Enter Name" />
//                       </Form.Group>

//                       <Form.Group className="mb-3" controlId="formBasicEmail">
//                         <Form.Label className="text-center">
//                           Email address
//                         </Form.Label>
//                         <Form.Control type="email" placeholder="Enter email" />
//                       </Form.Group>

//                       <Form.Group
//                         className="mb-3"
//                         controlId="formBasicPassword"
//                       >
//                         <Form.Label>Password</Form.Label>
//                         <Form.Control type="password" placeholder="Password" />
//                       </Form.Group>
//                       <Form.Group
//                         className="mb-3"
//                         controlId="formBasicPassword"
//                       >
//                         <Form.Label>Confirm Password</Form.Label>
//                         <Form.Control type="password" placeholder="Password" />
//                       </Form.Group>
//                       <Form.Group
//                         className="mb-3"
//                         controlId="formBasicCheckbox"
//                       >
//                       </Form.Group>
//                       <div className="d-grid">
//                         <Button variant="primary" type="submit">
//                           Create Account
//                         </Button>
//                       </div>
//                     </Form>
//                     <div className="mt-3">
//                       <p className="mb-0  text-center">
//                       Already have an account??{" "}
//                         <a href="{''}" className="text-primary fw-bold">
//                           Sign In
//                         </a>
//                       </p>
//                     </div>
//                   </div>
//                 </div>
//               </Card.Body>
//             </Card>
//           </Col>
//         </Row>
//       </Container>
//     </div>
//   )
// }


// import React from 'react';
// // import 'mdb-react-ui-kit/dist/css/mdb.min.css';
// import { 
//   MDBBtn, 
//   MDBContainer, 
//   MDBRow, 
//   MDBCol, 
//   MDBCard, 
//   MDBCardBody, 
//   MDBCardImage, 
//   MDBInput, 
//   MDBIcon, 
//   MDBCheckbox 
// } 
// from 'mdb-react-ui-kit';
// import flight_img from "../../assets/img/21308.jpg"

// function Signup() {
//   return (
//     <MDBContainer fluid>

//       <MDBCard className='text-black m-5' style={{borderRadius: '25px'}}>
//         <MDBCardBody>
//           <MDBRow>
//             <MDBCol md='10' lg='6' className='order-2 order-lg-1 d-flex flex-column align-items-center'>

//               <p classNAme="text-center h2 fw-bold mb-5 mx-1 mx-md-4 mt-4">Sign up</p>

//               <div className="d-flex flex-row align-items-center mb-4 ">
//                	<MDBIcon fas icon="user me-3" size='lg'/>
//                 <MDBInput label='Your Name' id='form1' type='text' className='w-100'/>
//               </div>

//               <div className="d-flex flex-row align-items-center mb-4">
//                 <MDBIcon fas icon="envelope me-3" size='lg'/>
//                 <MDBInput label='Your Email' id='form2' type='email'/>
//               </div>

//               <div className="d-flex flex-row align-items-center mb-4">
//                 <MDBIcon fas icon="lock me-3" size='lg'/>
//                 <MDBInput label='Password' id='form3' type='password'/>
//               </div>

//               <div className="d-flex flex-row align-items-center mb-4">
//                 <MDBIcon fas icon="key me-3" size='lg'/>
//                 <MDBInput label='Repeat your password' id='form4' type='password'/>
//               </div>

//               <div className='mb-4'>
//                 <MDBCheckbox name='flexCheck' value='' id='flexCheckDefault' label='Subscribe to our newsletter' />
//               </div>
	
//               <MDBBtn className='mb-4' size='lg'>Register</MDBBtn>
              
//             </MDBCol>

//             <MDBCol md='10' lg='6' className='order-1 order-lg-2 d-flex align-items-center'>
//               <MDBCardImage src={flight_img} fluid/>
//             </MDBCol>
            
//           </MDBRow>
//         </MDBCardBody>
//       </MDBCard>



//       <link
//         href="https://fonts.googleapis.com/css?family=Roboto:300,400,500,700&display=swap"
//         rel="stylesheet"
//         />
//       <link
//         href="https://use.fontawesome.com/releases/v5.15.1/css/all.css"
//         rel="stylesheet"
//         />

//     </MDBContainer>
//   );
// }

// export default Signup;

import React, { useState } from "react";

import {
    Container,
    Button,
    Row,
    Col,
    Form,
    FormControl
  } 
from "react-bootstrap";

import "bootstrap/dist/css/bootstrap.min.css";

import { Link } from "react-router-dom";

const Signup = () => {

    const name = "";
    
    return (
    <Container>
      <Row>
        <Col md="4">
          <h1>Sign up</h1>
          <Form>
            <Form.Group controlId="firstNameId">
              <Form.Label>Name</Form.Label>
              <Form.Control
                type="text"
                name="firstName"
                placeholder="Enter name"
                value={name}
                // onChange={onChange}
                />
              <FormControl.Feedback type="invalid"></FormControl.Feedback>
            </Form.Group>
          </Form>
          <Button 
            color="primary"
            // onClick={onSignupClick}  
            >Sign up</Button>
          <p className="mt-2">
            Already have account? <Link to="/login">Login</Link>
          </p>
        </Col>
      </Row>
    </Container>
  );
}
export default Signup;