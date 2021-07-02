import React, {useState, useContext} from 'react';
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
import { EventsContext } from '../../../context/events-context';
import { UsersContext } from '../../../context/users-context';

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

export default function SummaryHistoryTabPanel(props) {
  console.log("render SummaryHistoryTabPanel");

  const classes = useStyles();
  // const [currEquipState, currEquipDispatch] = useContext(CurrentEquipContext);
  const [eventsState, eventsDispatch] = useContext(EventsContext);
  const [usersState, usersDispatch] = useContext(UsersContext);

  const curDate = new Date();
  const [startDate, setStartDate] = useState(getUSFullDate(new Date(curDate.setDate(curDate.getDate() - SearchPeriod))));
  const [endDate, setEndDate] = useState(getUSFullDate(new Date()));
  
  // const [events, setEvents] = useState([]);  
  const [importantOnly, setImportantOnly] = useState(false);  
  const [filters, setFilters] = useState({
    byTitle: '',
    byDescr: ''
  });


  const handleStartDateChange = (ev) => {
    setStartDate(ev.target.value);
  };

  const handleEndDateChange = (ev) => {
    setEndDate(ev.target.value);
  };

  // const handleNameChange = (ev) => {
  //   setEquipName(ev.target.value);
  // };  

  const equipName = props.equipName; // currEquipState?.equipInfo;
  let events = eventsState.currentEvents;
  const token = usersState.token;

  const onSearch = async () => {    
    const allEvents = await EquipWorker.SearchEquip(token, 'Events', equipName, startDate, endDate);
    // setEvents(data);
    eventsDispatch({ type: 'SETEVENTS', payload: allEvents }); 
  };

  const onSelect = async (event) => {
    setImportantOnly(!importantOnly);
  };

  
  if(filters.byTitle){
    events = events?.filter(p => p.Description.includes(filters.byTitle));
  }

  if(filters.byDescr){
    events = events?.filter(p => p.Details.includes(filters.byDescr));
  }

  const onTitleFilterChange = async (val) =>{
    setFilters({...filters, ...{byTitle: val.target?.value}});
  }

  const onDescrFilterChange = async (val) =>{
    setFilters({...filters, ...{byDescr: val.target?.value}});
  }


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
        <FormControlLabel
          control={
            <Checkbox
                checked={importantOnly}
                onChange={onSelect}
            />
          }
          label="Только важные"
        />
        
        <Button variant="contained" color="primary" className={classes.commonSpacing} onClick={onSearch}>
            Искать
        </Button>

        <TextField id="outlined-basic" className={classes.commonSpacing} onChange={onTitleFilterChange} label="По названию" variant="outlined" />
        <TextField id="outlined-basic" className={classes.commonSpacing} onChange={onDescrFilterChange} label="По описанию" variant="outlined" />
        
    </div>
    <div className={classes.root}>        
      <CommonTimeLine equipName={equipName} rows={events} showImportantOnly={importantOnly}></CommonTimeLine>
    </div>
    </>
  );
}