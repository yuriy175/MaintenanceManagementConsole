import React, {useState, useContext} from 'react';
import { makeStyles } from '@material-ui/core/styles';

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
    width:'50%'
  }
}));


// flex-direction: row;
//   justify-content: center;
//   align-items: center;

export default function AdminTab(props) {
  console.log("render AdminTab");

  const classes = useStyles();
  const [usersState, usersDispatch] = useContext(UsersContext);
  const [login, setLogin] = useState('');
  const [password, setPassword] = useState('');
  const [surname, setSurname] = useState('');
  const [email, setEmail] = useState('');
  const [role, setRole] = useState('');

  const onLoginChange = (ev) => {
    setLogin(ev.target.value);
  }; 
  
  const onPasswordChange = (ev) => {
    setPassword(ev.target.value);
  }; 

  const onSurnameChange = (ev) => {
    setSurname(ev.target.value);
  }; 

  const onEmailChange = (ev) => {
    setEmail(ev.target.value);
  }; 

  const onRoleChange = (ev) => {
    setRole(ev.target.value);
  }; 

  const onAdd = async () => {
    const data = await AdminWorker.UpdateUser({login, password, surname, email, role});
  };

  return (
    <div className={classes.root}>
      <div className={classes.root}>
        <TextField className={classes.text} id="standard-basic" label="Логин" defaultValue={''} onChange={onLoginChange}/>
        <TextField className={classes.text} id="standard-basic" label="Пароль" defaultValue={''} onChange={onPasswordChange}/>
        <TextField className={classes.text} id="standard-basic" label="ФИО" defaultValue={''} onChange={onSurnameChange}/>
        <TextField className={classes.text} id="standard-basic" label="Почта" defaultValue={''} onChange={onEmailChange}/>
        <TextField className={classes.text} id="standard-basic" label="Роль" defaultValue={''} onChange={onRoleChange}/>
        <Button variant="contained" color="primary" className={classes.commonSpacing} onClick={onAdd}>
              Добавить
        </Button>
      </div>
      <UserTable></UserTable>
    </div>
  );
}