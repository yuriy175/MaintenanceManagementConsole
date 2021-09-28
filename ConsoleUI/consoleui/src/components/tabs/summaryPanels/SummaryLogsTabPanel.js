import React, {useState, useContext} from 'react';
import { makeStyles } from '@material-ui/core/styles';

import Typography from '@material-ui/core/Typography';
import Checkbox from '@material-ui/core/Checkbox';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import TextareaAutosize from '@material-ui/core/TextareaAutosize';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import { UsersContext } from '../../../context/users-context';

import * as EquipWorker from '../../../workers/equipWorker'

const useStyles = makeStyles((theme) => ({
  root: {
    width:'100%',
    borderColor: 'darkgray'
  },
  commonSpacing:{
    marginLeft: theme.spacing(1),
    marginRight: theme.spacing(1),
  },
}));

export default function SummaryLogsTabPanel(props) {
  console.log("render SummaryLogsTabPanel");

  const classes = useStyles();const [usersState, usersDispatch] = useContext(UsersContext);

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

  const onDurationChange = async (event) =>{
    const val = event.target.value;
    setDuration(!isNaN(val) || val === '' ? val : defaultDuration);
  }

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
      <Button variant="contained" color="primary" className={classes.commonSpacing} onClick={onStart}>
          Запустить
      </Button>

      <TextField id="outlined-basic" className={classes.commonSpacing} value={duration} onChange={onDurationChange} label="Длительность, сек" variant="outlined" />
      <TextareaAutosize className={classes.root}
        rowsMax={40}
        aria-label="maximum height"
      />
    </div>
  );
}
  