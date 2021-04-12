import React, {useContext} from 'react';
import CommonTable from '../CommonTable'
import * as AdminWorker from '../../../workers/adminWorker'
import { UsersContext } from '../../../context/users-context';

const columns = [
  { id: 'Login', label: 'Логин', minWidth: 170 },
  { id: 'Surname', label: 'ФИО', minWidth: 100 },
  { id: 'Email', label: 'Почта', minWidth: 100 },
  { id: 'Role', label: 'Роль', minWidth: 100 },
  { id: 'Disabled', label: 'Удален', checked: true, minWidth: 100 },
];

export default function UserTable(props) {
  console.log("render UserTable");
  const [usersState, usersDispatch] = useContext(UsersContext);

  const rows = props.data;

  const handleSelect = async (event, row) => {
    const Disabled = event.target.checked;//{id: "1", login, password, surname, email, role, disabled}
    const newRow = {...row, Disabled};
    const data = await AdminWorker.UpdateUser(newRow);
    const users = await AdminWorker.GetAllUsers();
    usersDispatch({ type: 'SETUSERS', payload: users }); 
  };

  return (
    <CommonTable columns={columns} rows={rows} onSelect={handleSelect}></CommonTable>
  );
}