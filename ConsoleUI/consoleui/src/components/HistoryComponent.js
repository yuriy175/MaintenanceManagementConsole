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
import StudiesTable from './tables/historyTables/StudiesTable'
import SofwareTable from './tables/historyTables/SofwareTable'
import DetectorTable from './tables/historyTables/DetectorTable'
import StandTable from './tables/historyTables/StandTable'
import DicomTable from './tables/historyTables/DicomTable'

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
  
  const [systemInfos, setSystemInfos] = useState([]);  
  const [organAutos, setOrganAutos] = useState([]);  
  const [generators, setGenerators] = useState([]);
  const [studies, setStudies] = useState([]);
  const [stands, setStands] = useState([]);
  const [detectors, setDetectors] = useState([]);
  const [dosimeters, setDosimeters] = useState([]);
  const [software, setSoftware] = useState([]);
  const [dicom, setDicom] = useState([]);

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
    if(currEquip === "SystemInfo"){
        setSystemInfos(data);
    } else if (currEquip === "OrganAutos"){
        setOrganAutos(data);
    } else if (currEquip === "Generators"){
        setGenerators(data);
    } else if (currEquip === "Studies"){
        setStudies(data);
    } else if (currEquip === "Stands"){
      setStands(data);
    } else if (currEquip === "Dosimeters"){
      setDosimeters(data);
    } else if (currEquip === "Detectors"){
      setDetectors(data);
    } else if (currEquip === "Software"){
      setSoftware(data);
    } else if (currEquip === "Dicom"){
      setDicom(data);
    }
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
                <option value={"Stands"} className={classes.optionStyle}>Штативы</option>
                <option value={"Dosimeters"} className={classes.optionStyle}>Дозиметры</option>
                <option value={"Detectors"} className={classes.optionStyle}>Детекторы</option>
                <option value={"Studies"} className={classes.optionStyle}>Исследования</option>
                <option value={"Software"} className={classes.optionStyle}>Приложения</option>
                <option value={"Dicom"} className={classes.optionStyle}>Dicom</option>
        </NativeSelect>
        <Button variant="contained" color="primary" className={classes.commonSpacing} onClick={onSearch}>
            Искать
        </Button>
    </div>
    <div className={classes.root}>        
        {currEquip === "SystemInfo" ? <SystemTable data={systemInfos}></SystemTable> : <></>}
        {currEquip === "OrganAutos" ? <OrganAutoTable data={organAutos}></OrganAutoTable> : <></>}     
        {currEquip === "Generators" ? <GeneratorTable data={generators}></GeneratorTable> : <></>}    
        {currEquip === "Studies" ? <StudiesTable data={studies}></StudiesTable> : <></>}  
        {currEquip === "Software" ? <SofwareTable data={software}></SofwareTable> : <></>}  
        {currEquip === "Detectors" ? <DetectorTable data={software}></DetectorTable> : <></>}  
        {currEquip === "Stands" ? <StandTable data={software}></StandTable> : <></>}  
        {currEquip === "Dicom" ? <DicomTable data={software}></DicomTable> : <></>}  
    </div>
    </>
  );
}