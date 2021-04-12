import React, {useState, useContext} from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Redirect } from 'react-router-dom';

import TextField from '@material-ui/core/TextField';
import NativeSelect from '@material-ui/core/NativeSelect';
import Button from '@material-ui/core/Button';

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
    marginBottom:'1em'
  },
  commonSpacing:{
    marginLeft: theme.spacing(1),
    marginRight: theme.spacing(1),
  },
  text:{
    width:'50%',
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

  if (redirect) {
    return <Redirect to="/info" />;  
  }

  const onLoginChange = (ev) => {
    setLogin(ev.target.value);
  }; 
  
  const onPasswordChange = (ev) => {
    setPassword(ev.target.value);
  }; 

  const onEmailChange = (ev) => {
    setEmail(ev.target.value);
  }; 

  const onLogin = async () => {
    const data = await AdminWorker.Login({login, password});
    if(data){
      usersDispatch({ type: 'SETUSER', payload: data }); 
      setRedirect(true);
    }
  };

  return (
    <div className={classes.root}>
      <div className={classes.root}>
        <TextField className={classes.text} id="standard-basic" label="Логин" defaultValue={''} onChange={onLoginChange}/>
        <TextField className={classes.text} id="standard-basic" label="Пароль" defaultValue={''} onChange={onPasswordChange}/>
        <TextField className={classes.text} id="standard-basic" label="Почта" defaultValue={''} onChange={onEmailChange}/>
        <Button variant="contained" color="primary" className={classes.commonSpacing} onClick={onLogin}>
              Login
        </Button>
      </div>
    </div>
  );
}