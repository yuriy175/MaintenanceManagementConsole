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

const SoftwareCard = React.memo((props) => {
//export default function SoftwareCard() {
  console.log(`! render SoftwareCard`);

  const classes = useCardsStyles();
  const bull = <span className={classes.bullet}>•</span>;

  const software = props.software;
  return (
    <Card className={classes.root}>
      <CardContent>
        <Typography variant="h5" component="h2">
          {bull}Software
        </Typography>
        <CardRow descr={software?.Sysinfo?.OS} value={software?.Sysinfo?.Version}></CardRow>
        <CardRow descr={software?.MSSQL?.SQL} value={software?.MSSQL?.Version}></CardRow>
        <CardRow descr={'Пользователь'} value={software?.User?.Current_user}></CardRow>

        <Typography variant="h6" component="h2">
          {bull}Атлас
        </Typography>
        <CardRow descr="Версия" value={software?.Atlas?.Atlas_Version}></CardRow>
        <CardRow descr="Xilib" value={software?.Atlas?.XiLibs_Version}></CardRow>
        <CardRow descr="Конфигурация" value={software?.Atlas?.Complex_type}></CardRow>
        <CardRow descr="Язык" value={software?.Atlas?.Complex_type}></CardRow>
        <CardRow descr="Multimonitor" value={software?.Atlas?.Multimonitor}></CardRow>
        <CardRow descr={software?.Atlas?.Atlas_User?.Role} value={software?.Atlas?.Atlas_User?.User}></CardRow>        
        <CardRow descr="Ошибки" value={''}></CardRow>
        {software?.ErrorDescriptions?.length ? 
          software.ErrorDescriptions.map((i, ind) => (
            <CardErrorRow key={ind.toString()}  descr={i.Code} value={i.Description} ></CardErrorRow>
            ))
            :
            <></>          
        }
      </CardContent>
    </Card>
  );
});

export default SoftwareCard;