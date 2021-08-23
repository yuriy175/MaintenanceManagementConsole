import React, {useState, useContext} from 'react';
import { makeStyles } from '@material-ui/core/styles';

import Typography from '@material-ui/core/Typography';
import Checkbox from '@material-ui/core/Checkbox';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import TextareaAutosize from '@material-ui/core/TextareaAutosize';
import { UsersContext } from '../../../context/users-context';

import * as EquipWorker from '../../../workers/equipWorker'

const useStyles = makeStyles((theme) => ({
  root: {
    width:'100%',
    borderColor: 'darkgray'
  },
}));

export default function SummaryLogsTabPanel(props) {
  console.log("render SummaryLogsTabPanel");

  const classes = useStyles();const [usersState, usersDispatch] = useContext(UsersContext);

  const equipName = props.equipName;
  const token = usersState.token;
  const [generatorLogs, setGeneratorLogs] = useState(false);  
  const [standLogs, setStandLogs] = useState(false);  
  const [detectorLogs, setDetectorLogs] = useState(false);  

  const onGeneratorLogs = async (event) => {
    const value = !generatorLogs;
    setGeneratorLogs(value);
    const result = await EquipWorker.SetEquipLogsOn(token, equipName, "generator", value);
  };

  const onStandLogs = async (event) => {
    const value = !standLogs;
    setStandLogs(value);
    const result = await EquipWorker.SetEquipLogsOn(token, equipName, "stand", value);
  };

  const onDetectorLogs = async (event) => {
    const value = !detectorLogs;
    setDetectorLogs(value);
    const result = await EquipWorker.SetEquipLogsOn(token, equipName, "detector", value);
  };

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
      <TextareaAutosize className={classes.root}
        rowsMax={40}
        aria-label="maximum height"
      />
    </div>
  );
}
  