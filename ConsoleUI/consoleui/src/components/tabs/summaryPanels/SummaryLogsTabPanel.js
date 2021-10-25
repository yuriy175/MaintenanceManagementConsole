import React, {useState, useContext} from 'react';
import { makeStyles } from '@material-ui/core/styles';

import Typography from '@material-ui/core/Typography';
import Checkbox from '@material-ui/core/Checkbox';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import TextareaAutosize from '@material-ui/core/TextareaAutosize';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';
import { UsersContext } from '../../../context/users-context';

import * as EquipWorker from '../../../workers/equipWorker'
import { EquipLogContext } from '../../../context/equipLog-context';
import {parseLocalString} from '../../../utilities/utils'


const useStyles = makeStyles((theme) => ({
  root: {
    width:'100%',
    borderColor: 'darkgray'
  },
  commonSpacing:{
    marginLeft: theme.spacing(1),
    marginRight: theme.spacing(1),
  },
  logsArea: {
    width:'100%',
    borderColor: 'darkgray',
    height: '70em',
    overflowY: 'auto',
  },
  logClass: {
    textAlign: 'left',
  }
}));

const SummaryLogsTabPanel = React.memo((props) => {
// export default function StandCard() {
// export default function SummaryLogsTabPanel(props) {
  console.log("render SummaryLogsTabPanel");

  const classes = useStyles();
  const [usersState, usersDispatch] = useContext(UsersContext);
  const [equipLogs, equipLogsDispatch] = useContext(EquipLogContext);

  const equipName = props.equipName;
  const token = usersState.token;
  const defaultDuration = 30;
  const [generatorLogs, setGeneratorLogs] = useState(false);  
  const [standLogs, setStandLogs] = useState(false);  
  const [detectorLogs, setDetectorLogs] = useState(false);  
  const [duration, setDuration] = useState(defaultDuration);

  const onGeneratorLogs = async (event) => {
    const value = !generatorLogs;
    setGeneratorLogs(value);
  };

  const onStandLogs = async (event) => {
    const value = !standLogs;
    setStandLogs(value);
  };

  const onDetectorLogs = async (event) => {
    const value = !detectorLogs;
    setDetectorLogs(value);
  };

  const onStart = async () => { 
    const equipTypes = [];   
    if(generatorLogs){
      equipTypes.push("generator");
    }

    if(standLogs){
      equipTypes.push("stand");
    }

    if(detectorLogs){
      equipTypes.push("detector");
    }

    if(!duration){
      setDuration(defaultDuration);
    }

    const result = await EquipWorker.SetEquipLogsOn(token, equipName, equipTypes, duration ?? defaultDuration);
  };

  const onReset = async () => { 
    equipLogsDispatch({ type: 'RESETLOG', payload: true}); 
  }

  const onDurationChange = async (event) =>{
    const val = event.target.value;
    setDuration(!isNaN(val) || val === '' ? val : defaultDuration);
  }

  // const newLog = equipLogs.currentLog?.Type + ' ' + equipLogs.currentLog?.State?.Timestamp + ' ' + equipLogs.currentLog?.State?.Data;
  const logs = equipLogs?.currentLogs;
  return (
    <div>      
      <FormControlLabel
              control={
                <Checkbox
                    checked={generatorLogs}
                    onChange={onGeneratorLogs}
                />
              }
              label="Логи генератора"
            />
      <FormControlLabel
              control={
                <Checkbox
                    checked={standLogs}
                    onChange={onStandLogs}
                />
              }
              label="Логи штатива"
            />
      <FormControlLabel
              control={
                <Checkbox
                    checked={detectorLogs}
                    onChange={onDetectorLogs}
                />
              }
              label="Логи детектора"
            />
      
      <TextField id="outlined-basic" className={classes.commonSpacing} value={duration} onChange={onDurationChange} label="Длительность, сек" variant="outlined" />

      <Button variant="contained" color="primary" className={classes.commonSpacing} onClick={onStart}>
          Запустить
      </Button>

      <Button variant="contained" color="primary" className={classes.commonSpacing} onClick={onReset}>
          Очистить
      </Button>
      
      <List className={classes.logsArea}>
        {logs?.map((log, ind) => (
          <ListItemText 
              key={log?.State?.Timestamp.toString()} 
              className={classes.logClass}
              primary={parseLocalString(log?.State?.Timestamp) + ' ' + log?.Type  + ' ' + log?.State?.Data} />
            ))}          
      </List>
      {/* {logs?.map((log, ind) => (
              <Typography>{log?.Type + ' ' + log?.State?.Timestamp + ' ' + log?.State?.Data}</Typography>
            ))} */}
    </div>
  );
});

export default SummaryLogsTabPanel;
  