import {useState} from 'react'
import styled from "styled-components"
import { motion } from 'framer-motion';
import APIService from '../services/APIService'
import {useNavigate} from 'react-router-dom'
import { VisibilityOff, Visibility } from '@mui/icons-material'


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

const Register = () => {
  const navigate = useNavigate();
    const [firstName, setFirstName] = useState('');
  const [lastName, setLastName] = useState('');
  const [phone, setPhone] = useState('');
  const [email, setEmail] = useState('');
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [confirmPassword, setConfirmPassword] = useState('');
  const [isFocusedName, setIsFocusedName] = useState(false);
    const [isFocusedSurname, setIsFocusedSurname] = useState(false);
    const [isPasswordValid, setIsPasswordValid] = useState(true);
    const [isConfirmPasswordValid, setIsConfirmPasswordValid] = useState(true);
    const [passwordIconVisability, setPasswordIconVisability] = useState(false);
    const [confirmPasswordIconVisability, setConfirmPasswordIconVisability] = useState(false);

  function handlePasswordChange(e) {
    const password = e.target.value;
    setPassword(password);
    const valid = /^(?=.*[A-Z])(?=.*\d)[A-Za-z\d@$!%*#?&]{8,}$/.test(password);
    setIsPasswordValid(valid);
  }
  function handlePasswordConfirmChange(e) {
    const passwordConfirm = e.target.value;
    setConfirmPassword(passwordConfirm)
    if (password !== passwordConfirm){
      setIsConfirmPasswordValid(false)
    }
    else {
      setIsConfirmPasswordValid(true)
    }
  }
  const handleSubmit = async (e) => {
    e.preventDefault();
    if (!isConfirmPasswordValid || !isPasswordValid) {
      return;
    }
    try {
      const response = await APIService.Register(username, password, email, firstName, lastName, phone);
      console.log(response);
      navigate('/login')
      
    } catch (error) {
      console.error(error);
        if (error.response) {
          const { status, data } = error.response;
          if (status === 400 && data.username) {
            alert(`Korisnik sa korisničkim imenom "${username}" već postoji.`);
          } else {
            alert(`Došlo je do greške: ${data || status}`);
          }
        } else {
          alert('Došlo je do greške prilikom povezivanja sa serverom.');
        }
    }
  }

  const handlePasswordIconClick = () => {
    setPasswordIconVisability(!passwordIconVisability);
  };
  
  const handleConfirmPasswordIconClick = () => {
    setConfirmPasswordIconVisability(!confirmPasswordIconVisability);
  };
    
  return (
    <div>
        <FormContainer>
        <CheckoutForm onSubmit={handleSubmit} >
          <FormWrapper>
        <FormField>
          <Label>Ime*</Label>
          <Input
            type="text"
            value={firstName}
            onChange={e => setFirstName(e.target.value)}
            required
            min_width="230px"
            placeholder="Marko"
            onFocus={() => setIsFocusedName(true)}
            onBlur={() => setIsFocusedName(false)}
          />
          <motion.div
            initial={{ opacity: 0, y: 10 }}
            animate={{ opacity: isFocusedName ? 1 : 0, y: isFocusedName ? 0 : 10 }}
            transition={{ duration: 0.3 }}
            style={{ fontSize: '14px', marginTop: '5px' }}
      >
        Unesite Vaše ime
      </motion.div>
        </FormField>
        <FormField>
          <Label>Prezime*</Label>
          <Input
            type="text"
            value={lastName}
            onChange={e => setLastName(e.target.value)}
            required
            min_width="230px"
            placeholder="Markovic"
            onFocus={() => setIsFocusedSurname(true)}
            onBlur={() => setIsFocusedSurname(false)}
          />
          <motion.div
            initial={{ opacity: 0, y: 10 }}
            animate={{ opacity: isFocusedSurname ? 1 : 0, y: isFocusedSurname ? 0 : 10 }}
            transition={{ duration: 0.3 }}
            style={{ fontSize: '14px', marginTop: '5px' }}
      >
        Unesite Vaše Prezime
      </motion.div>
        </FormField>
        </FormWrapper>
        <FormField>
          <Label>Kontakt telefon *</Label>
          <Input
            type="tel"
            value={phone}
            onChange={e => setPhone(e.target.value)}
            required
            placeholder="060 123 1234"
          />
        </FormField>
        <FormField>
          <Label>Email adresa *</Label>
          <Input
            type="email"
            value={email}
            onChange={e => setEmail(e.target.value)}
            required
            placeholder="ime.prezima@gmail.com"
          />
        </FormField>
        <FormField>
          <Label>Korisnicko ime *</Label>
          <Input
            type="text"
            value={username}
            onChange={e => setUsername(e.target.value)}
            required
            placeholder="marko11"
          />
        </FormField>
        <FormField>
          <Label>Lozinka *</Label>
          <Input
             type={passwordIconVisability ? "text" : "password"}
            value={password}
            onChange={handlePasswordChange}
            style={{ borderColor: isPasswordValid ? 'initial' : 'red' }}
            required
          /> 
          {passwordIconVisability ? (
            <Visibility onClick={handlePasswordIconClick} />
          ) : (
            <VisibilityOff onClick={handlePasswordIconClick} />
          )}
        </FormField>
          {!isPasswordValid && (
        <div style={{ color: 'red' }}>
          Lozinka mora sadrzati najmanje jedno veliko slovo,jedan broj i da bude najmanje 8 karaktera dugacka.
        </div>
      )}
        <FormField>
          <Label>Potvrda lozinke *</Label>
          <Input
            type={confirmPasswordIconVisability ? "text" : "password"}
            value={confirmPassword}
            onChange={handlePasswordConfirmChange}
            style={{ borderColor: isConfirmPasswordValid ? 'initial' : 'red' }}
            required
          />
           {confirmPasswordIconVisability ? (
            <Visibility onClick={handleConfirmPasswordIconClick} />
          ) : (
            <VisibilityOff onClick={handleConfirmPasswordIconClick} />
          )}
        </FormField>
        {!isConfirmPasswordValid && (
        <div style={{ color: 'red' }}>
          Lozinke se ne poklapaju
        </div>
      )}
        <CheckoutButtonDiv>
        <CheckoutButton type="submit">Registracija</CheckoutButton>
        </CheckoutButtonDiv>
</CheckoutForm>
</FormContainer>
    </div>
  )
}

export default Register