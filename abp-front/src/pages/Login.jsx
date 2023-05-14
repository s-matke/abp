import {useContext, useState} from 'react'
import styled from "styled-components"
import { motion } from 'framer-motion';
import APIService from '../services/APIService'
import {useNavigate} from 'react-router-dom'
import { VisibilityOff, Visibility } from '@mui/icons-material'
import AuthContext from '../context/AuthContext';


const FormContainer = styled.div`
    display: flex;
    justify-content: center;
    @media only screen and (max-width:768px){
       flex-direction: column;
       align-items: center;
    }
`;
const CheckoutForm = styled.form`

`;
const FormWrapper = styled.div`
    display:flex;
    justify-content: space-around;
    width: 100%;
    @media only screen and (max-width:768px){
       flex-direction: column;
    }
    
`;

const FormField = styled.div`
  display: flex;
  flex-direction: column;
  margin-bottom: 20px;
  width: 100%;
  
`;

const Label = styled.label`
  font-size: 14px;
  font-weight: 400;
  margin-bottom: 5px;
`;

const Input = styled.input`
    outline:none;
  padding: 10px;
  border: 0.5px solid #ccc;
  border-radius: 5px;
  box-shadow: 1px 1px 3px rgba(0, 0, 0, 0.3);
  width: ${(props) => props.min_width || '90%'};
  ::placeholder {
        font-size: 12px;
        font-weight: 400;
        color:#ccc;
        font-family: 'Urbanist';
  }
  font-weight: 600;
  font-family: 'Righteous';
  font-size: 12px;
    font-weight: 400;
    @media only screen and (max-width:768px){
       width:70vw;
    }
    
`;
const CheckoutButtonDiv = styled.div`
 @media only screen and (max-width:768px){
    display:flex;
    align-items: center;
    justify-content: center;
    position: fixed;
    bottom:0;
    left:0;
    right:0;
 }
`;
const CheckoutButton = styled.button`
    width: 50%;
    margin-top:20px;
    border-radius: 100px;
    font-size: 12px;
    font-weight: 700;
    letter-spacing: .04em;
    text-align: center;
    text-transform: uppercase;
    padding: 15px 20px;
    text-decoration: none;
    transition: .6s all cubic-bezier(.25,1,.5,1);
    background-color: ${props => props.disabled ? '#ccc' : '#5862bf'};
    border: ${props => props.disabled ? 'none' : '2px solid #5862bf'};
    color: #fff;
    display: block;
    margin-bottom: 1rem;
    cursor: ${props => props.disabled ? 'default' : 'pointer'};
    &:hover{
        background-color: ${props => props.disabled ? '#ccc' : '#6b76da'};
    }
    @media only screen and (max-width:768px){
        font-size: 12px;
        width: 70vw;
        
    }
`;
const Login = () => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const navigate = useNavigate();
    const [isFocusedUsername, setIsFocusedUsername] = useState(false);
    const [isFocusedPassword, setIsFocusedPassword] = useState(false);
    const { handleLogin }  = useContext(AuthContext)

  return (
    <div>
        <FormContainer>
        <CheckoutForm onSubmit={handleLogin} >
          <FormWrapper>
        <FormField>
          <Label>Username</Label>
          <Input
            type="text"
            name="username"
            value={username}
            onChange={e => setUsername(e.target.value)}
            required
            min_width="230px"
            placeholder="marko11"
            onFocus={() => setIsFocusedUsername(true)}
            onBlur={() => setIsFocusedUsername(false)}
          />
          <motion.div
            initial={{ opacity: 0, y: 10 }}
            animate={{ opacity: isFocusedUsername ? 1 : 0, y: isFocusedUsername ? 0 : 10 }}
            transition={{ duration: 0.3 }}
            style={{ fontSize: '14px', marginTop: '5px' }}
      >
        Unesite Vaš username
      </motion.div>
        </FormField>
        <FormField>
          <Label>Lozinka*</Label>
          <Input
            type="password"
            name="password"
            value={password}
            onChange={e => setPassword(e.target.value)}
            required
            min_width="230px"
            onFocus={() => setIsFocusedPassword(true)}
            onBlur={() => setIsFocusedPassword(false)}
          />
          <motion.div
            initial={{ opacity: 0, y: 10 }}
            animate={{ opacity: isFocusedPassword ? 1 : 0, y: isFocusedPassword ? 0 : 10 }}
            transition={{ duration: 0.3 }}
            style={{ fontSize: '14px', marginTop: '5px' }}
      >
        Unesite Vašu Lozinku
      </motion.div>
        </FormField>
        </FormWrapper>
        <CheckoutButtonDiv>
        <CheckoutButton type="submit">Prijava</CheckoutButton>
        </CheckoutButtonDiv>
</CheckoutForm>
</FormContainer>
    </div>
  )
}

export default Login