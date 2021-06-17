import React, {useContext} from 'react';
import CommonTable from '../CommonTable'

import { SummaryTabIndex, MainTabPanelIndex } from '../../../model/constants';
import { CurrentEquipContext } from '../../../context/currentEquip-context';
import { AppContext } from '../../../context/app-context';
import { AllEquipsContext } from '../../../context/allEquips-context';
import * as EquipWorker from '../../../workers/equipWorker'
import {useSetCurrEquip} from '../../../hooks/useSetCurrEquip'

export default function EquipTable(props) {
  console.log("render EquipTable");
  const [currEquipState, currEquipDispatch] = useContext(CurrentEquipContext);
  const [appState, appDispatch] = useContext(AppContext);
  const [allEquipsState, allEquipsDispatch] = useContext(AllEquipsContext);
  const setCurrEquip = useSetCurrEquip();

  const columns = [
    { id: 'IsActive', label: 'Активен', checked: true, minWidth: 50,
      format: (row) => allEquipsState.connectedEquips?.includes(row.EquipName)
    },
    { id: 'EquipName', label: 'Комплекс', minWidth: 170 },
    { id: 'RegisterDate', label: 'Дата регистрации', minWidth: 170 },
    { id: 'HospitalName', label: 'ЛПУ', minWidth: 100 },
    { id: 'HospitalAddress', label: 'Адрес', minWidth: 100 },
    { id: 'HospitalLatitude', label: 'Широта', minWidth: 100 },
    { id: 'HospitalLongitude', label: 'Долгота', minWidth: 100 },
    // { id: 'Disabled', label: 'Удален', checked: true, minWidth: 100 },
  ];

  const rows = props.data;
  const handleSelect = async (event, row) => {

    const equipInfo = row.EquipName;
    setCurrEquip(equipInfo, 'SETEQUIPINFO');
    allEquipsDispatch({ type: 'ADDSELECTEDEQUIPS', payload: equipInfo }); 
    appDispatch({ type: 'SETTAB', payload: {tab: SummaryTabIndex, panel: MainTabPanelIndex} }); 
  };

  return (
    <CommonTable readonly columns={columns} rows={rows} selectedRow={currEquipState.equipInfo} onRowClick={handleSelect}></CommonTable>
  );
}