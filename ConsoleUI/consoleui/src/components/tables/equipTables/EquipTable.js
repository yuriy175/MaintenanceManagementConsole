import React, {useContext} from 'react';
import CommonTable from '../CommonTable'

const columns = [
  { id: 'IsActive', label: 'Активен', checked: true, minWidth: 50 },
  { id: 'EquipName', label: 'Комплекс', minWidth: 170 },
  { id: 'RegisterDate', label: 'Дата регистрации', minWidth: 170 },
  { id: 'HospitalName', label: 'ЛПУ', minWidth: 100 },
  { id: 'HospitalAddress', label: 'Адрес', minWidth: 100 },
  { id: 'HospitalLatitude', label: 'Широта', minWidth: 100 },
  { id: 'HospitalLatitude', label: 'Долгота', minWidth: 100 },
  // { id: 'Disabled', label: 'Удален', checked: true, minWidth: 100 },
];

export default function EquipTable(props) {
  console.log("render EquipTable");

  const rows = props.data;

  const handleSelect = async (event, row) => {
    // const Disabled = event.target.checked;//{id: "1", login, password, surname, email, role, disabled}
    // const newRow = {...row, Disabled};
    // const data = await AdminWorker.UpdateUser(newRow);
    // const users = await AdminWorker.GetAllUsers();
    // usersDispatch({ type: 'SETUSERS', payload: users }); 
  };

  return (
    <CommonTable columns={columns} rows={rows} onSelect={handleSelect}></CommonTable>
  );
}