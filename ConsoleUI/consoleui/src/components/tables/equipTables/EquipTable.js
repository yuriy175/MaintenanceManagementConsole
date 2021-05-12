import React, {useContext} from 'react';
import CommonTable from '../CommonTable'
import { CurrentEquipContext } from '../../../context/currentEquip-context';
import { AllEquipsContext } from '../../../context/allEquips-context';
import * as EquipWorker from '../../../workers/equipWorker'
import {useSetCurrEquip} from '../../../hooks/useSetCurrEquip'

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
  //const [currEquipState, currEquipDispatch] = useContext(CurrentEquipContext);
  const [allEquipsState, allEquipsDispatch] = useContext(AllEquipsContext);
  const setCurrEquip = useSetCurrEquip();

  const rows = props.data;

  const handleSelect = async (event, row) => {

    const equipInfo = row.EquipName;
    setCurrEquip(equipInfo, 'SETEQUIPINFO');
    allEquipsDispatch({ type: 'ADDSELECTEDEQUIPS', payload: equipInfo }); 
    
    // currEquipDispatch({ type: 'RESET', payload: true });    
    // currEquipDispatch({ type: 'SETEQUIPINFO', payload: equipInfo }); 

    // // new software & system info come very slowly
    // const sysInfo = await EquipWorker.GetPermanentData("SystemInfo", equipInfo);
    // currEquipDispatch({ type: 'SETSYSTEM', payload: sysInfo }); 

    // const swInfo = await EquipWorker.GetPermanentData("Software", equipInfo);
    // currEquipDispatch({ type: 'SETSOFTWARE', payload: swInfo }); 

    // await EquipWorker.Activate(equipInfo, currEquipState.equipInfo);
  };

  return (
    <CommonTable columns={columns} rows={rows} onRowClick={handleSelect}></CommonTable>
  );
}