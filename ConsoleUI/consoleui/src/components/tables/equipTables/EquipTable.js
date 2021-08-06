import React, {useContext, useState} from 'react';
import { makeStyles } from '@material-ui/core/styles';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import Checkbox from '@material-ui/core/Checkbox';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';

import {AdminRole} from '../../../model/constants'
import CommonTable from '../CommonTable'

import { SummaryTabIndex, MainTabPanelIndex } from '../../../model/constants';
import { AppContext } from '../../../context/app-context';
import { AllEquipsContext } from '../../../context/allEquips-context';
import { UsersContext } from '../../../context/users-context';
import * as EquipWorker from '../../../workers/equipWorker'
import {useSetCurrEquip} from '../../../hooks/useSetCurrEquip'
import {parseLocalString} from '../../../utilities/utils'
import ConfirmDlg from '../../dialogs/ConfirmDlg'


const useStyles = makeStyles((theme) => ({
  root: {
    display:"flex"
  },
  commonSpacing:{
    marginLeft: theme.spacing(1),
    marginRight: theme.spacing(1),
  },
}));

//export default function EquipTable(props) {
const EquipTable = React.memo((props) => {
  console.log("render EquipTable");
  const classes = useStyles();

  const [appState, appDispatch] = useContext(AppContext);
  const [usersState, usersDispatch] = useContext(UsersContext);
  const [allEquipsState, allEquipsDispatch] = useContext(AllEquipsContext);
  const setCurrEquip = useSetCurrEquip();
  const [visibleOnly, setVisibleOnly] = useState(true);  
  const [activeOnly, setActiveOnly] = useState(false);  
  const [filtrate, setFiltrate] = useState(false);  
  const [filters, setFilters] = React.useState({
    byEquip: '',
    byHospital: '',
    byAddress:''
  });
  const [openConfirm, setOpenConfirm] = React.useState({Result: false});

  const isAdmin = usersState.currentUser?.Role === AdminRole;
  const columns = [
    { id: 'IsActive', label: 'Активен', checked: true, minWidth: 50, sortable: true,
      // format: (row) => allEquipsState.connectedEquips?.includes(row.EquipName)
    },
    { id: 'EquipName', label: 'Комплекс', minWidth: 170, sortable: true },
    { id: 'RegisterDate', label: 'Дата регистрации', minWidth: 170,
      format: (value) => parseLocalString(value)
    },
    { id: 'HospitalName', label: 'ЛПУ', minWidth: 100, sortable: true },
    { id: 'HospitalAddress', label: 'Адрес', minWidth: 100, sortable: true },
    { id: 'HospitalLatitude', label: 'Широта', minWidth: 100 },
    { id: 'HospitalLongitude', label: 'Долгота', minWidth: 100 },
    { id: 'LastSeen', label: 'Посл. сообщение', minWidth: 100, sortable: true,
      //format: (value) => value ? parseLocalString(value) : ""
    },
    
    // { id: 'Disabled', label: 'Удален', checked: true, minWidth: 100 },
  ];
  if(isAdmin){
    columns.push({ id: 'Disabled', label: 'Скрыт', checked: true, minWidth: 100,
        format: (row) => row.Disabled
      });
  }

  const onFilter = async ()  => {
    setFiltrate(!filtrate);
  }

  let rows = props.data;
  if(activeOnly) {
    rows = rows?.filter(p => p.IsActive);
  }

  if(visibleOnly) {
    rows = rows?.filter(p => !p.Disabled);
  }
  
  if(filters.byEquip){
    rows = rows?.filter(p => p.EquipName.includes(filters.byEquip));
  }

  if(filters.byHospital){
    rows = rows?.filter(p => p.HospitalName.includes(filters.byHospital));
  }

  if(filters.byAddress){
    rows = rows?.filter(p => p.HospitalAddress.includes(filters.byAddress));
  }

  rows.forEach((row) => 
  {
    row.LastSeen = row.LastSeen ? parseLocalString(row.LastSeen) : "";
  });
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
    setOpenConfirm({Result: true, Context: row});
    /* */
  };

  const onVisibleOnly = async (event) => {
    const value = !visibleOnly;
    setVisibleOnly(value);
    const allEquips = await EquipWorker.GetAllEquips(usersState.token, visibleOnly);
    allEquipsDispatch({ type: 'SETALLEQUIPS', payload: allEquips }); 
  };

  const onActiveOnly = async (event) => {
    const value = !activeOnly;
    setActiveOnly(value);
  };

  const onEquipFilterChange = async (val) =>{
    setFilters({...filters, ...{byEquip: val.target?.value}});
  }

  const onHospFilterChange = async (val) =>{
    setFilters({...filters, ...{byHospital: val.target?.value}});
  }

  const onAddressFilterChange = async (val) =>{
    setFilters({...filters, ...{byAddress: val.target?.value}});
  }

  const onConfirmClose = async (result, context) => {
    if(result){
      const row = context;
      const equipInfo = row.EquipName;
      row.Disabled = !row.Disabled
      await EquipWorker.DisableEquipInfo(usersState.token, equipInfo, row.Disabled);
      allEquipsDispatch({ type: 'UPDATEALLEQUIPS', payload: row });
    }
    setOpenConfirm({Result: false});
  };

  const equipInfo = props.equipInfo;

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
        <FormControlLabel
              control={
                <Checkbox
                    checked={activeOnly}
                    onChange={onActiveOnly}
                />
              }
              label="Только активные"
            />
        <TextField id="outlined-basic" className={classes.commonSpacing} onChange={onEquipFilterChange} label="По комплексу" variant="outlined" />
        <TextField id="outlined-basic" className={classes.commonSpacing} onChange={onHospFilterChange} label="По ЛПУ" variant="outlined" />
        <TextField id="outlined-basic" className={classes.commonSpacing} onChange={onAddressFilterChange} label="По адресу" variant="outlined" />
        {/* <Button variant="contained" color="primary" className={classes.commonSpacing} onClick={onFilter}>
            Фильтровать
        </Button> */}
      </div>
      <div className={classes.root}>
        {/* <CommonTable readonly columns={columns} rows={rows} selectedRow={currEquipState.equipInfo} onRowClick={handleRowClick} onSelect={handleSelect}></CommonTable> */}
        <CommonTable readonly defaultSort={'IsActive'} defaultSortOrder={'desc'} columns={columns} rows={rows} selectedRow={equipInfo} onRowClick={handleRowClick} onSelect={handleSelect}></CommonTable>
      </div>
      <ConfirmDlg 
        open={openConfirm.Result} 
        сonfirmMessage={'Вы действительно хотите '+(openConfirm.Context?.Disabled ? 'открыть' : 'скрыть')+' комплекс?'}
        onClose={onConfirmClose}
        context={openConfirm.Context}
        >          
      </ConfirmDlg>
    </>
  );
});

export default EquipTable;