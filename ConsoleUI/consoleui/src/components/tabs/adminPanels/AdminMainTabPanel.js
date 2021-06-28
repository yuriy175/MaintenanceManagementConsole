import React, {useState, useContext} from 'react';
import { makeStyles } from '@material-ui/core/styles';

import TextField from '@material-ui/core/TextField';
import NativeSelect from '@material-ui/core/NativeSelect';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import InputLabel from '@material-ui/core/InputLabel';

import UserTable from '../../tables/adminTables/UserTable'
import * as AdminWorker from '../../../workers/adminWorker'
import { UsersContext } from '../../../context/users-context';
import { Roles, UserRole } from '../../../model/constants';

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


export default function AdminMainTabPanel(props) {
  console.log("render AdminMainTabPanel");

  const classes = useStyles();
  const [usersState, usersDispatch] = useContext(UsersContext);
  const [login, setLogin] = useState('');
  const [password, setPassword] = useState('');
  const [surname, setSurname] = useState('');
  const [email, setEmail] = useState('');
  const [role, setRole] = useState(UserRole);

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

  const handleRoleChange = async (event) => {
    const select = event.target;
    const val = select.options[select.selectedIndex].value;

    setRole(val);
  };

  const onAdd = async () => {
    const token = usersState.token
    const data = await AdminWorker.UpdateUser({id: '', login, password, surname, email, role, disabled: false}, token);
    const users = await AdminWorker.GetAllUsers(token);
    usersDispatch({ type: 'SETUSERS', payload: users }); 
  };

  const onLogin = async () => {
    const data = await AdminWorker.Login({login, password});
  };

  return (
    <div className={classes.root}>
      <div className={classes.root}>
        <TextField className={classes.text} id="standard-basic" required={true} label="Логин" defaultValue={''} onChange={onLoginChange}/>
        <TextField className={classes.text} id="standard-basic" required={true} label="Пароль" defaultValue={''} onChange={onPasswordChange}/>
        <TextField className={classes.text} id="standard-basic" required={true} label="ФИО" defaultValue={''} onChange={onSurnameChange}/>
        <TextField className={classes.text} id="standard-basic" label="Почта" defaultValue={''} onChange={onEmailChange}/>
        {/* <TextField className={classes.text} id="standard-basic" label="Роль" defaultValue={''} onChange={onRoleChange}/> */}
        <FormControl className={classes.text}>
              <InputLabel htmlFor="grouped-native-select">Роль</InputLabel>
              <NativeSelect
                value={role}
                onChange={handleRoleChange}
                name="roles"
                className={classes.selectEmpty}
                variant="outlined"
              >
                {Roles.map((i, ind) => (
                    <option key={ind.toString()} value={i} className={classes.optionStyle}>{i}</option>
                    ))
                }
              </NativeSelect>
            </FormControl>
        <Button variant="contained" color="primary" className={classes.commonSpacing} onClick={onAdd}>
              Добавить
        </Button>
      </div>
      <UserTable data={usersState.users}></UserTable>
    </div>
  );
}