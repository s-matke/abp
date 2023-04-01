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
// import 'mdb-react-ui-kit/dist/css/mdb.min.css'
import 'bootstrap/dist/css/bootstrap.min.css'

import axios from "axios";
import { toInteger } from "lodash";
import { useEffect, useState } from "react";
import { FaIcons } from "react-icons/fa";
import { useLocation, useNavigate } from "react-router";
import Select from "react-select";

// import { FcSurvey } from "react-icons/fc";
import { RiSurveyLine } from "react-icons/ri";
import { Link } from "react-router-dom";

// import 'mdbreact/dist/css/mdb.css'

function Signup() {


  // const btnClick = () => {
  //   console.log("Hello")
  // }

  // return (
  //   <MDBContainer fluid>

  //     <MDBCard className='text-black' style={{borderRadius: '25px'}}>
  //       <MDBCardBody>
  //         <MDBRow>
  //           <MDBCol md='10' lg='6' className='order-2 order-lg-2 d-flex flex-column align-items-center'>

  //             <p className="text-center h1 fw-bold mb-5 mx-1 mx-md-4 mt-4">Sign up</p>

  //             <div className="d-flex flex-row align-items-center mb-4 ">
  //              	<MDBIcon fas icon="user me-3" size='lg'/>
  //               <MDBInput label='Your Name' id='form1' type='text' className='w-100'/>
  //             </div>

  //             <div className="d-flex flex-row align-items-center mb-4">
  //               <MDBIcon fas icon="envelope me-3" size='lg'/>
  //               <MDBInput label='Your Email' id='form2' type='email'/>
  //             </div>

  //             <div className="d-flex flex-row align-items-center mb-4">
  //               <MDBIcon fas icon="lock me-3" size='lg'/>
  //               <MDBInput label='Password' id='form3' type='password'/>
  //             </div>

  //             <div className="d-flex flex-row align-items-center mb-4">
  //               <MDBIcon fas icon="key me-3" size='lg'/>
  //               <MDBInput label='Repeat your password' id='form4' type='password'/>
  //             </div>

  //             <div className='mb-4'>
  //               <MDBCheckbox name='flexCheck' value='' id='flexCheckDefault' label='Subscribe to our newsletter' />
  //             </div>
	
  //             <MDBBtn className='mb-4' size='lg' onClick={btnClick}>Register</MDBBtn>
              
  //           </MDBCol>

  //           <MDBCol md='10' lg='6' className='order-1 order-lg-2 d-flex align-items-center'>
  //             <MDBCardImage src={flight_img} fluid/>
  //           </MDBCol>
            
  //         </MDBRow>
  //       </MDBCardBody>
  //     </MDBCard>
  //   </MDBContainer>
  // );
    const [report, setReport] = useState(
        {
            user_id: 1,
            appointment_id: 2,
            attendanceStatus: "",
            bloodType: "",
            bloodQuantity: "",
            doctorsNote: "",
            equipmentQuantity: ""
        }
    )
  
    const [equipment, setEquipment] = useState({})

    const optionList = [
        { value: "0", label: "A+"},
        { value: "1", label: "A-"},
        { value: "2", label: "B+"},
        { value: "3", label: "B-"},
        { value: "4", label: "0+"},
        { value: "5", label: "0-"},
        { value: "6", label: "AB+"},
        { value: "7", label: "AB-"},
    ]

    const attendanceStatusList = [
        { value: "0", label: "Attended" },
        { value: "1", label: "Missed" },
        { value: "2", label: "Rejected" },
    ]

    const submitReport = () => {
        console.log("Hello")
    }

    const onInputChange = (e) => {
        console.log("helo")
    }

    const handleChange = (event, name) => {
      console.log("Hnadle")
    };

    console.log(report)

    const redirectSurvey = () => {
        console.log("hello")
    }

    return(
        <div className="container position-relative">
            <div className="row">
                <div className="col-md-6 offset-md-4 border rounded p-4 mt-2 shadow position-relative">
                    <h2 className="text-center m-5">Sign Up</h2>
                    {/* <div className="mb-4">
                            <h5 style={{'width':'50%%', 'text-align':'center', 'border-bottom':'1px solid #000', 'line-height':'0.1em', 'margin':'10px 0 20px'}}>
                                <span style={{'background':'#fff', 'padding':'0 5px'}}>Info about <b>Yes</b>'s donation</span></h5>
                    </div> */}
                    <form onSubmit={(e) => submitReport(e)}>
                        <div className="mb-3">
                            <label htmlFor="name" className="form-label">
                                Name:
                            </label>
                            <div className="input-group">
                                <input
                                    type={"text"}
                                    className="form-control"
                                    placeholder="Enter your name..."
                                    name="name"
                                    onChange={(e) => onInputChange(e)}
                                    required
                                    />
                            </div>
                        </div>
                        <div className="mb-3">
                            <label htmlFor="blood_type" className="form-label">
                                Surname:
                            </label>
                            <div className="input-group">
                                <input
                                    type={"text"}
                                    className="form-control"
                                    placeholder="Enter your surname..."
                                    name="surname"
                                    onChange={(e) => onInputChange(e)}
                                    required
                                    />
                            </div>
                        </div>
                        <div className="mb-4">
                            <label htmlFor="quantity" className="form-label">
                                Email:
                            </label>
                            <div className="input-group">
                                <input
                                    type={"email"}
                                    className="form-control"
                                    placeholder="Enter your email..."
                                    name="email"
                                    onChange={(e) => onInputChange(e)}
                                    required
                                    />
                            </div>
                        </div>
                        <div className="mb-4">
                            <label htmlFor="quantity" className="form-label">
                                Password:
                            </label>
                            <div className="input-group">
                                <input
                                    type={"password"}
                                    className="form-control"
                                    placeholder="Enter your password..."
                                    name="password"
                                    onChange={(e) => onInputChange(e)}
                                    required
                                    />
                                </div>
                        </div>
                        <div className="md-3">
                            <label id="description">Description:</label> <br/>
                            <textarea 
                                    className="form-control" 
                                    maxLength={255} 
                                    type={"text"} 
                                    name="doctorsNote"
                                    placeholder="Note..." 
                                    onChange={(e) => onInputChange(e)}
                                    required
                            />
                        </div>
                        {/* <div className="mb-4">
                            <h5 style={{'width':'50%%', 'text-align':'center', 'border-bottom':'1px solid #000', 'line-height':'0.1em', 'margin':'10px 0 20px'}}>
                                <span style={{'background':'#fff', 'padding':'0 5px'}}>Equipment</span></h5>
                        </div>
                        <div className="mb-5">
                            <label htmlFor="quantity" className="form-label">
                                Quantity:
                            </label>
                            <div className="input-group">
                                <input
                                    type={"number"}
                                    className="form-control"
                                    placeholder="Used equipment..."
                                    name="equipment_quantity"
                                    required
                                    />
                            </div>
                        </div> */}

                        <button type="submit" className="btn btn-outline-primary" 
                                style={{'width':'150px', 'height':'50px', 'margin-left':'39%', 'margin-top':'15px'}}>
                        Submit
                        </button>
                    </form>
                </div>
            </div>
        </div>
    );
}

export default Signup;

// import React, { useState } from "react";

// import {
//     Container,
//     Button,
//     Row,
//     Col,
//     Form,
//     FormControl
//   } 
// from "react-bootstrap";

// import "bootstrap/dist/css/bootstrap.min.css";

// import { Link } from "react-router-dom";

// const Signup = () => {

//     const name = "";
    
//     return (
//     <Container>
//       <Row>
//         <Col md="4">
//           <h1>Sign up</h1>
//           <Form>
//             <Form.Group controlId="firstNameId">
//               <Form.Label>Name</Form.Label>
//               <Form.Control
//                 type="text"
//                 name="firstName"
//                 placeholder="Enter name"
//                 value={name}
//                 // onChange={onChange}
//                 />
//               <FormControl.Feedback type="invalid"></FormControl.Feedback>
//             </Form.Group>
//           </Form>
//           <Button 
//             color="primary"
//             // onClick={onSignupClick}  
//             >Sign up</Button>
//           <p className="mt-2">
//             Already have account? <Link to="/login">Login</Link>
//           </p>
//         </Col>
//       </Row>
//     </Container>
//   );
// }
// export default Signup;