import React, {useContext} from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import Button from '@material-ui/core/Button';
import Typography from '@material-ui/core/Typography';

import { CurrentEquipContext } from '../../context/currentEquip-context';
import {useCardsStyles} from './CommonCard'
import CardRow, {CardErrorRow} from './CardRow'

export default function SoftwareCard() {
  console.log(`! render SoftwareCard`);

  const [currEquipState, currEquipDispatch] = useContext(CurrentEquipContext);
  const classes = useCardsStyles();
  const bull = <span className={classes.bullet}>•</span>;

  return (
    <Card className={classes.root}>
      <CardContent>
        <Typography variant="h5" component="h2">
          {bull}Software
        </Typography>
        <CardRow descr={currEquipState.software?.Sysinfo?.OS} value={currEquipState.software?.Sysinfo?.Version}></CardRow>
        <CardRow descr={currEquipState.software?.MSSQL?.SQL} value={currEquipState.software?.MSSQL?.Version}></CardRow>
        <CardRow descr={'Пользователь'} value={currEquipState.software?.User?.Current_user}></CardRow>

        <Typography variant="h6" component="h2">
          {bull}Атлас
        </Typography>
        <CardRow descr="Версия" value={currEquipState.software?.Atlas?.Atlas_Version}></CardRow>
        <CardRow descr="Xilib" value={currEquipState.software?.Atlas?.XiLibs_Version}></CardRow>
        <CardRow descr="Конфигурация" value={currEquipState.software?.Atlas?.Complex_type}></CardRow>
        <CardRow descr="Язык" value={currEquipState.software?.Atlas?.Complex_type}></CardRow>
        <CardRow descr="Multimonitor" value={currEquipState.software?.Atlas?.Multimonitor}></CardRow>
        <CardRow descr="Ошибки" value={''}></CardRow>
        {currEquipState.software?.ErrorDescriptions?.length ? 
          currEquipState.software.ErrorDescriptions.map((i, ind) => (
            <CardErrorRow key={ind.toString()}  descr={i.Code} value={i.Description} ></CardErrorRow>
            ))
            :
            <></>          
        }
      </CardContent>
    </Card>
  );
}