import React, {useContext, useState} from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import Button from '@material-ui/core/Button';
import Typography from '@material-ui/core/Typography';

import {useCardsStyles} from './CommonCard'
import { CurrentEquipContext } from '../../context/currentEquip-context';
import * as EquipWorker from '../../workers/equipWorker'
import {CardButtonedRow, CardSwitchedRow} from './CardRow'

const RemoteAccessCard = React.memo((props) => {  
  console.log(`! render RemoteAccessCard`);
  const [detailedXilib, setDetailedXilib] = useState(false);
  const [verboseXilib, setVerboseXilib] = useState(false);

  const classes = useCardsStyles();
  const bull = <span className={classes.bullet}>•</span>;

  const equipInfo = props.equipInfo;
  const token = props.token;
  const onRunTeamViewer = async () => {
    const res = await EquipWorker.RunTeamViewer(token, equipInfo);
  };

  const onRunTaskManager = async () => {
    const res = await EquipWorker.RunTaskManager(token, equipInfo);
  };

  const onAtlasLogs = async () => {
    const res = await EquipWorker.SendAtlasLogs(token, equipInfo);
  };

  const onXilibLogs = async () => {
    const res = await EquipWorker.XilibLogsOn(token, equipInfo, detailedXilib, verboseXilib);
  };

  const onDetailedXilib = async () => {
    setDetailedXilib(!detailedXilib);
  };

  const onVerboseXilib = async () => {
    setVerboseXilib(!verboseXilib);
  };

  const remoteaccess = props.remoteaccess;
  return (
    <Card className={classes.root}>
      <CardContent>
        <Typography variant="h5" component="h2">
          {bull}Удаленный доступ
        </Typography>
        <CardButtonedRow descr={'TeamViewer'} title={'Запустить'} onClick={onRunTeamViewer}></CardButtonedRow>
        <CardButtonedRow descr={'TaskManager'} title={'Запустить'} onClick={onRunTaskManager}></CardButtonedRow>
        <CardButtonedRow descr={'Логи Атлас'} title={'Прислать'} onClick={onAtlasLogs}></CardButtonedRow>
        <CardButtonedRow descr={'Логи Xilib'} 
          title={!remoteaccess.Xilogs? 'Не опред' : remoteaccess.Xilogs.on ? 'Выключить' : 'Включить'} 
          onClick={onXilibLogs}
          buttonColor={!remoteaccess.Xilogs? undefined : remoteaccess.Xilogs.on ? "secondary" : "primary"}
          disabled={!remoteaccess.Xilogs}
        ></CardButtonedRow>
        <CardSwitchedRow descr={'Детальный'} onChange={onDetailedXilib}></CardSwitchedRow>
        <CardSwitchedRow descr={'Подробный'} onChange={onVerboseXilib}></CardSwitchedRow>
      </CardContent>
    </Card>
  );
});

export default RemoteAccessCard;