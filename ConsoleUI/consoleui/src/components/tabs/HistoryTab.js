import React, {useState, useContext} from 'react';
import { makeStyles } from '@material-ui/core/styles';

import TextField from '@material-ui/core/TextField';
import NativeSelect from '@material-ui/core/NativeSelect';
import Button from '@material-ui/core/Button';

import {SearchPeriod} from '../../model/constants'
import {getUSFullDate} from '../../utilities/utils'

import SystemTable from '../tables/historyTables/SystemTable'
import OrganAutoTable from '../tables/historyTables/OrganAutoTable'
import GeneratorTable from '../tables/historyTables/GeneratorTable'
import StudiesTable from '../tables/historyTables/StudiesTable'
import SofwareTable from '../tables/historyTables/SofwareTable'
import DetectorTable from '../tables/historyTables/DetectorTable'
import StandTable from '../tables/historyTables/StandTable'
import DicomTable from '../tables/historyTables/DicomTable'

import * as EquipWorker from '../../workers/equipWorker'
import { CurrentEquipContext } from '../../context/currentEquip-context';
import { UsersContext } from '../../context/users-context';

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

export default function HistoryTab(props) {
  console.log("render HistoryTab");

  const classes = useStyles();
  const [currEquipState, currEquipDispatch] = useContext(CurrentEquipContext);
  const [usersState, usersDispatch] = useContext(UsersContext);

  const curDate = new Date();
  const [startDate, setStartDate] = useState(getUSFullDate(new Date(curDate.setDate(curDate.getDate() - SearchPeriod))));
  const [endDate, setEndDate] = useState(getUSFullDate(new Date()));
  const [equipName, setEquipName] = useState(currEquipState?.equipInfo);
  const [currType, setCurrType] = useState("Software");//"SystemInfo");
  
  const [systemInfos, setSystemInfos] = useState([]);  
  const [organAutos, setOrganAutos] = useState([]);  
  const [generators, setGenerators] = useState([]);
  const [studies, setStudies] = useState([]);
  const [stands, setStands] = useState([]);
  const [detectors, setDetectors] = useState([]);
  const [dosimeters, setDosimeters] = useState([]);
  const [software, setSoftware] = useState([]);
  const [dicom, setDicom] = useState([]);

  const token = usersState.token;

  const handleTypeChange = async (event) => {
    const select = event.target;
    const val = select.options[select.selectedIndex].value;

    setCurrType(val);
  };

  const handleStartDateChange = (ev) => {
    setStartDate(ev.target.value);
  };

  const handleEndDateChange = (ev) => {
    setEndDate(ev.target.value);
  };

  const handleNameChange = (ev) => {
    setEquipName(ev.target.value);
  };  

  const onSearch = async () => {
    const data = await EquipWorker.SearchEquip(token, currType, equipName, startDate, endDate);
    switch (currType) {
      case "SystemInfo":
        setSystemInfos(data);
        break;
      case "OrganAutos":
        setOrganAutos(data);
        break;
      case "Generators":
        setGenerators(data);
        break;
      case "Studies":
        setStudies(data);
        break;
      case "Stands":
        setStands(data);
        break;
      case "Dosimeters":
        setDosimeters(data);
        break;
      case "Detectors":
        setDetectors(data);
        break;
      case "Software":
        setSoftware(data);
        break;
      case "Dicom":
        setDicom(data);
        break;
      default:
        alert( "Нет таких значений" );
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
                value={currType}
                onChange={handleTypeChange}
                name="equips"
                className={classes.selectEmpty, classes.commonSpacing}
                variant="outlined"
                label="Данные"
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
        <TextField id="standard-basic" label="Компекс" defaultValue={equipName} onChange={handleNameChange}/>
        <Button variant="contained" color="primary" className={classes.commonSpacing} onClick={onSearch}>
            Искать
        </Button>
    </div>
    <div className={classes.root}>        
        {currType === "SystemInfo" ? <SystemTable equipName={equipName} data={systemInfos}></SystemTable> : <></>}
        {currType === "OrganAutos" ? <OrganAutoTable data={organAutos}></OrganAutoTable> : <></>}     
        {currType === "Generators" ? <GeneratorTable data={generators}></GeneratorTable> : <></>}    
        {currType === "Studies" ? <StudiesTable data={studies}></StudiesTable> : <></>}  
        {currType === "Software" ? <SofwareTable equipName={equipName} data={software}></SofwareTable> : <></>}  
        {currType === "Detectors" ? <DetectorTable data={detectors}></DetectorTable> : <></>}  
        {currType === "Stands" ? <StandTable data={stands}></StandTable> : <></>}  
        {currType === "Dicom" ? <DicomTable data={dicom}></DicomTable> : <></>}  
    </div>
    </>
  );
}