import React, {useState, useContext, useEffect} from 'react';
import { makeStyles } from '@material-ui/core/styles';

import TextField from '@material-ui/core/TextField';
import NativeSelect from '@material-ui/core/NativeSelect';
import Button from '@material-ui/core/Button';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import Checkbox from '@material-ui/core/Checkbox';

import {SearchPeriod} from '../../../model/constants'
import {getUSFullDate} from '../../../utilities/utils'

import SystemTable from '../../tables/historyTables/SystemTable'
import OrganAutoTable from '../../tables/historyTables/OrganAutoTable'
import GeneratorTable from '../../tables/historyTables/GeneratorTable'
import StudiesTable from '../../tables/historyTables/StudiesTable'
import SofwareTable from '../../tables/historyTables/SofwareTable'
import DetectorTable from '../../tables/historyTables/DetectorTable'
import StandTable from '../../tables/historyTables/StandTable'
import DicomTable from '../../tables/historyTables/DicomTable'
import CommonTimeLine from '../../timelines/CommonTimeLine'

import * as EquipWorker from '../../../workers/equipWorker'
import { CurrentEquipContext } from '../../../context/currentEquip-context';
import { UsersContext } from '../../../context/users-context';
import {parseLocalString} from '../../../utilities/utils'

const useStyles = makeStyles((theme) => ({
  root: {
    // display:"flex"
  },
  commonSpacing:{
    marginLeft: theme.spacing(1),
    marginRight: theme.spacing(1),
  },
  textField: {    
    marginTop: theme.spacing(2),
  },
  textFieldInline: {    
    marginLeft: theme.spacing(2),
  },
  selectEmpty: {
    marginTop: theme.spacing(2),
  },
  optionStyle:{
    // backgroundColor: "#3f51b5",
    // color:"black",
  },
  legend:{
    display: "flex",
    justifyContent: "flex-end",
    marginLeft: "0px",
  }
}));

export default function SummaryInfoTabPanel(props) {
  console.log("render SummaryInfoTabPanel");

  const classes = useStyles();
  const [usersState] = useContext(UsersContext);
  const [currEquipState, currEquipDispatch] = useContext(CurrentEquipContext);

  const [info, setInfo] = React.useState(currEquipState.info);

  const equipName = props.equipName; 
  const token = usersState.token;

  useEffect(() => {
    const info = currEquipState.info;
    if(!info.EquipName){
      info.EquipName = equipName;
    }

    info.ManufacturingDate = getUSFullDate(new Date(info.ManufacturingDate));
    info.MontageDate = getUSFullDate(new Date(info.MontageDate));
    info.WarrantyStartDate = getUSFullDate(new Date(info.WarrantyStartDate));
    info.WarrantyEndDate = getUSFullDate(new Date(info.WarrantyEndDate));

    setInfo(info);
  }, [currEquipState.info]);

  const onUpdate = async () => {    
    const allEvents = await EquipWorker.UpdateEquipInfo(token, equipName, info);
    currEquipDispatch({ type: 'SETINFO', payload: info }); 
  };

  const onSerialNumChange = async (val) =>{
    setInfo({...info, SerialNum: val.target?.value});
  }  

  const onModelChange = async (val) =>{
    setInfo({...info, Model: val.target?.value});
  }  

  const onAgreementChange = async (val) =>{
    setInfo({...info, Agreement: val.target?.value});
  }  

  const onManufDateChange = (val) => {
    setInfo({...info, ManufacturingDate: val.target?.value});
  };

  const onMontageDateChange = (val) => {
    setInfo({...info, MontageDate: val.target?.value});
  };

  const onWarrantyStartDateChange = (val) => {
    setInfo({...info, WarrantyStartDate: val.target?.value});
  };

  const onWarrantyEndDateChange = (val) => {
    setInfo({...info, WarrantyEndDate: val.target?.value});
  };  

  const onContactInfoChange = async (val) =>{
    setInfo({...info, ContactInfo: val.target?.value});
  }  

  const onReparInfoChange = async (val) =>{
    setInfo({...info, ReparInfo: val.target?.value});
  } 

  return (
    <>
      <div className={classes.root}>
          <TextField className={classes.textField} margin="dense" id="serNum" label="Серийный номер" fullWidth variant="standard"
              value={info.SerialNum} onChange={onSerialNumChange}/>
          <TextField className={classes.textField} margin="dense" id="model" label="Модель" fullWidth variant="standard"
              value={info.Model} onChange={onModelChange}/>
          <TextField className={classes.textField} margin="dense" id="agree" label="Договор" fullWidth variant="standard"
              value={info.Agreement} onChange={onAgreementChange}/>

          <TextField className={classes.textField} id="manufDate" label="Дата производства" fullWidth type="date"
              value={info.ManufacturingDate} onChange={onManufDateChange} 
              InputLabelProps={{ shrink: true}}/>
          
          <TextField className={classes.textField} id="montageDate" label="Дата монтажа" fullWidth type="date"
              value={info.MontageDate} onChange={onMontageDateChange} 
              InputLabelProps={{ shrink: true}}/>

          <FormControlLabel className={classes.legend} fullWidth
                    control={
                      <div>
                        <TextField className={classes.textFieldInline} id="warrantyStartDate" label="Дата начала" type="date"
                            value={info.WarrantyStartDate} onChange={onWarrantyStartDateChange} 
                            InputLabelProps={{ shrink: true}}/>          
                        <TextField className={classes.textFieldInline} id="warrantyEndDate" label="Дата конца" type="date"
                            value={info.WarrantyEndDate} onChange={onWarrantyEndDateChange} 
                            InputLabelProps={{ shrink: true}}/>
                      </div>
                    }
                    label="Гарантийное обслуживание"
                    labelPlacement="start"
                  />

          <TextField
              id="outlined-multiline-static" className={classes.textField} label="Контакная информация"
              multiline rows={5} fullWidth defaultValue="" variant="outlined"
              value={info.ContactInfo} onChange={onContactInfoChange}
            />
          <TextField
              id="outlined-multiline-static" className={classes.textField} label="Информация о ремонтах"
              multiline rows={10} fullWidth defaultValue="" variant="outlined"
              value={info.ReparInfo} onChange={onReparInfoChange}
            />
      </div>
      <Button variant="contained" color="primary" className={classes.commonSpacing} onClick={onUpdate}>
            Обновить
      </Button>
    </>
  );
}