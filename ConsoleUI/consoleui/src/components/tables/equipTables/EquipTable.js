import React, {useContext} from 'react';
import CommonTable from '../CommonTable'

import { SummaryTabIndex, MainTabPanelIndex } from '../../../model/constants';
import { CurrentEquipContext } from '../../../context/currentEquip-context';
import { AppContext } from '../../../context/app-context';
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
  const [appState, appDispatch] = useContext(AppContext);
  const [allEquipsState, allEquipsDispatch] = useContext(AllEquipsContext);
  const setCurrEquip = useSetCurrEquip();

  const rows = props.data;

  const handleSelect = async (event, row) => {

    const equipInfo = row.EquipName;
    setCurrEquip(equipInfo, 'SETEQUIPINFO');
    allEquipsDispatch({ type: 'ADDSELECTEDEQUIPS', payload: equipInfo }); 
    appDispatch({ type: 'SETTAB', payload: {tab: SummaryTabIndex, panel: MainTabPanelIndex} }); 
  };

  return (
    <CommonTable columns={columns} rows={rows} onRowClick={handleSelect}></CommonTable>
  );
}