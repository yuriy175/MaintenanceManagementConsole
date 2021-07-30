import React, {useState, useContext} from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Redirect } from 'react-router-dom';
import TextField from '@material-ui/core/TextField';
import NativeSelect from '@material-ui/core/NativeSelect';
import Button from '@material-ui/core/Button';
import PasswordComponent from '../commonComponents/PasswordComponent'

import UserTable from '../tables/adminTables/UserTable'

import * as AdminWorker from '../../workers/adminWorker'
import { UsersContext } from '../../context/users-context';

const useStyles = makeStyles((theme) => ({
  root: {
    display:"flex",
    flexDirection: "column",
    justifyContent: "center",
    alignItems: "center",
    width:'100%',
    marginBottom:'1em',
    backgroundColor: theme.palette.background.default,  
  },
  commonSpacing:{
    marginLeft: theme.spacing(1),
    marginRight: theme.spacing(1),
  },
  text:{
    width:'50%',
    marginBottom:'1em'
  },
  alert:{
    backgroundColor: '#f44336',
    width: '50%',
    color: 'white',
    height: '3em',
    borderRadius: '0.3em',
    marginBottom:'1em'
  }
}));


export default function AuthComponent(props) {
  console.log("render AuthComponent");

  const classes = useStyles();
  const [usersState, usersDispatch] = useContext(UsersContext);
  const [login, setLogin] = useState('');
  const [password, setPassword] = useState('');
  const [email, setEmail] = useState('');
  const [redirect, setRedirect] = useState(false);
  const [showError, setShowError] = useState(false);

  if (redirect) {
    return <Redirect to="/info" />;  
  }

  const onLoginChange = (ev) => {
    setLogin(ev.target.value);
    if(showError)setShowError(false);
  }; 
  
  const onPasswordChange = (ev) => {
    setPassword(ev.target.value);
    if(showError)setShowError(false);
  }; 

  const onEmailChange = (ev) => {
    setEmail(ev.target.value);
    if(showError)setShowError(false);
  }; 

  const onLogin = async () => {
    const parseJwt = (token) => {
      try {
        return JSON.parse(atob(token.split('.')[1]));
      } catch (e) {
        return null;
      }
    };
    const data = await AdminWorker.Login({login: 'se', password: '1', email});
    // const data = await AdminWorker.Login({login: 'sa', password: 'medtex', email});
    // const data = await AdminWorker.Login({login, password, email});
    
    if(data && data.Token){
      const claims = parseJwt(data.Token);      
      usersDispatch({ type: 'SETUSER', payload: {Token: data.Token, Claims: claims, Surname: data.Surname} }); 
      setRedirect(true);
    }
    else{
      setShowError(true);
    }
  };

  return (
    <div className={classes.root}>
      <div className={classes.root}>
        <TextField className={classes.text} id="standard-basic" label="Логин" defaultValue={''} onChange={onLoginChange}/>
        {/* <TextField className={classes.text} id="standard-basic" label="Пароль" defaultValue={''} onChange={onPasswordChange}/> */}
        <PasswordComponent password={password} onChange={onPasswordChange}></PasswordComponent>
        <TextField className={classes.text} id="standard-basic" label="Почта" defaultValue={''} onChange={onEmailChange}/>        
        {showError ? <div className={classes.alert}>
          Логин или пароль неверны
        </div> : <></>}
        <Button variant="contained" color="primary" className={classes.commonSpacing} onClick={onLogin}>
              Login
        </Button>
      </div>
    </div>
  );
}