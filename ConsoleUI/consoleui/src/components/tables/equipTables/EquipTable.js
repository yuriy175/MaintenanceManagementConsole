import React, {useContext, useState} from 'react';
import { makeStyles } from '@material-ui/core/styles';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import Checkbox from '@material-ui/core/Checkbox';

import CommonTable from '../CommonTable'

import { SummaryTabIndex, MainTabPanelIndex } from '../../../model/constants';
import { CurrentEquipContext } from '../../../context/currentEquip-context';
import { AppContext } from '../../../context/app-context';
import { AllEquipsContext } from '../../../context/allEquips-context';
import * as EquipWorker from '../../../workers/equipWorker'
import {useSetCurrEquip} from '../../../hooks/useSetCurrEquip'

const useStyles = makeStyles((theme) => ({
  root: {
    display:"flex"
  }
}));

export default function EquipTable(props) {
  console.log("render EquipTable");
  const classes = useStyles();

  const [currEquipState, currEquipDispatch] = useContext(CurrentEquipContext);
  const [appState, appDispatch] = useContext(AppContext);
  const [allEquipsState, allEquipsDispatch] = useContext(AllEquipsContext);
  const setCurrEquip = useSetCurrEquip();
  const [visibleOnly, setVisibleOnly] = useState(true);  

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
    { id: 'Disabled', label: 'Скрыт', checked: true, minWidth: 100,
      format: (row) => row.Disabled
    },
    // { id: 'Disabled', label: 'Удален', checked: true, minWidth: 100 },
  ];

  const rows = props.data;
  const handleRowClick = async (event, row) => {

    var dataColumn = event.target.getAttribute("data-column");

    if(dataColumn === "Disabled")
    {
      return;
    }

    const equipInfo = row.EquipName;
    setCurrEquip(equipInfo, 'SETEQUIPINFO');
    allEquipsDispatch({ type: 'ADDSELECTEDEQUIPS', payload: equipInfo }); 
    appDispatch({ type: 'SETTAB', payload: {tab: SummaryTabIndex, panel: MainTabPanelIndex} }); 
  };

  const handleSelect = async (event, row) => {
    const equipInfo = row.EquipName;
    row.Disabled = !row.Disabled
    await EquipWorker.DisableEquipInfo(equipInfo, row.Disabled);
    const allEquips = await EquipWorker.GetAllEquips(!visibleOnly);
    allEquipsDispatch({ type: 'SETALLEQUIPS', payload: allEquips }); 
  };

  const onVisibleOnly = async (event) => {
    const value = !visibleOnly;
    setVisibleOnly(value);
    const allEquips = await EquipWorker.GetAllEquips(visibleOnly);
    allEquipsDispatch({ type: 'SETALLEQUIPS', payload: allEquips }); 
  };

  return (
    <>
      <div className={classes.root}>
        <FormControlLabel
              control={
                <Checkbox
                    checked={visibleOnly}
                    onChange={onVisibleOnly}
                />
              }
              label="Только нескрытые"
            />
      </div>
      <div className={classes.root}>
        <CommonTable readonly columns={columns} rows={rows} selectedRow={currEquipState.equipInfo} onRowClick={handleRowClick} onSelect={handleSelect}></CommonTable>
      </div>
    </>
  );
}