import React, {useState} from 'react';
import { makeStyles } from '@material-ui/core/styles';

import TextField from '@material-ui/core/TextField';
import NativeSelect from '@material-ui/core/NativeSelect';
import Button from '@material-ui/core/Button';

import {SearchPeriod} from '../model/constants'
import {getUSFullDate} from '../utilities/utils'

import SystemTable from './tables/historyTables/SystemTable'
import OrganAutoTable from './tables/historyTables/OrganAutoTable'
import GeneratorTable from './tables/historyTables/GeneratorTable'
import * as EquipWorker from '../workers/equipWorker'

const useStyles = makeStyles((theme) => ({
  root: {
    display:"flex"
  },
  commonSpacing:{
    marginLeft: theme.spacing(1),
    marginRight: theme.spacing(1),
  },
  textField: {    
    width: 200,
  },
  selectEmpty: {
    marginTop: theme.spacing(2),
  },
  optionStyle:{
    // backgroundColor: "#3f51b5",
    // color:"black",
  }
}));

export default function HistoryComponent(props) {
  console.log("render HistoryComponent");

  const classes = useStyles();

  const curDate = new Date();
  const [startDate, setStartDate] = useState(getUSFullDate(new Date(curDate.setDate(curDate.getDate() - SearchPeriod))));
  const [endDate, setEndDate] = useState(getUSFullDate(new Date()));
  const [currEquip, setCurrEquip] = useState("OrganAutos");

  const handleEquipsChange = async (event) => {
    const select = event.target;
    const val = select.options[select.selectedIndex].value;

    setCurrEquip(val);
  };

  const handleStartDateChange = (ev) => {
    setStartDate(ev.target.value);
  };

  const handleEndDateChange = (ev) => {
    setEndDate(ev.target.value);
  };

  const onSearch = async () => {
    const data = await EquipWorker.SearchEquip(currEquip, startDate, endDate);
    // const filter = {useLast: selectedValue, sincePeriod: sincePeriodValue, lastPeriod: lastPeriodValue};
    // deviceDispatch({
    //   type: 'SETDATEFILTER',
    //   payload:  filter,
    // });
    // await getFilteredMeasures(deviceState.currentdevice, filter, token, deviceDispatch, undefined, undefined, CurrentMeasureType);
  };


  return (
      <>
    <div className={classes.root}>
        <TextField
            id="startDate"
            label="Начальная дата"
            type="date"
            defaultValue={startDate}
            onChange={handleStartDateChange}
            className={classes.textField, classes.commonSpacing}
            InputLabelProps={{
            shrink: true,
            }}
        />
        <TextField
            id="endDate"
            label="Конечная дата"
            type="date"
            defaultValue={endDate}
            onChange={handleEndDateChange}
            className={classes.textField, classes.commonSpacing}
            InputLabelProps={{
            shrink: true,
            }}
        />
        <NativeSelect
                value={currEquip}
                onChange={handleEquipsChange}
                name="equips"
                className={classes.selectEmpty, classes.commonSpacing}
                variant="outlined"
              >
                <option value={"SystemInfo"} className={classes.optionStyle}>Система</option>
                <option value={"OrganAutos"} className={classes.optionStyle}>Орган авто</option>
                <option value={"Generators"} className={classes.optionStyle}>Генераторы</option>
        </NativeSelect>
        <Button variant="contained" color="primary" className={classes.commonSpacing} onClick={onSearch}>
            Искать
        </Button>
    </div>
    <div className={classes.root}>        
        {currEquip === "SystemInfo" ? <SystemTable></SystemTable> : <></>}
        {currEquip === "OrganAutos" ? <OrganAutoTable></OrganAutoTable> : <></>}     
        {currEquip === "Generators" ? <GeneratorTable></GeneratorTable> : <></>}    
           
    </div>
    </>
  );
}