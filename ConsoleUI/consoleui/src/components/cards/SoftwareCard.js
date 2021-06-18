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

  const volatile = props.volatile;
  const dbStates = volatile?.DBStates;
  const software = props.software?.Software;
  const atlas = Array.isArray(software?.Atlas) ? software?.Atlas[0] : null;  // props.software?.Atlas;
  const atlasUser = volatile?.AtlasUser;
  const osInfo = Array.isArray(software?.OsInfos) ? software?.OsInfos[0] : null; 
  const sql = Array.isArray(software?.SqlServices) ? software?.SqlServices[0] : null; 
  const databases = dbStates ?? software?.SqlDatabases;
  
  return (
    <Card className={classes.root}>
      <CardContent>
        <Typography variant="h5" component="h2">
          {bull}Software
        </Typography>
        <CardRow descr={osInfo?.OsCaption} value={osInfo?.OsVersion}></CardRow>
        <CardRow descr={sql?.SqlName} value={sql?.SqlVersion}></CardRow>
        <CardRow descr={'Пользователь'} value={osInfo?.CurrentUser}></CardRow>
        <CardRow descr="Ошибки" value={''}></CardRow>
        {volatile?.ErrorDescriptions?.length ? 
          volatile.ErrorDescriptions.map((i, ind) => (
            <CardErrorRow key={ind.toString()}  descr={i.Code} value={i.Description} ></CardErrorRow>
            ))
            :
            <></>          
        }

        <Typography variant="h6" component="h2">
          {bull}Базы данных
        </Typography>
        {databases?.length ? 
          databases.map((i, ind) => (
            <CardRow key={ind.toString()} descr={i.Name} value={i.Status}></CardRow>
            ))
            :
            <></>          
        }

        <Typography variant="h6" component="h2">
          {bull}Атлас
        </Typography>
        <CardRow descr="Версия" value={atlas?.AtlasVersion}></CardRow>
        <CardRow descr="Xilib" value={atlas?.XilibsVersion}></CardRow>
        <CardRow descr="Конфигурация" value={atlas?.ComplexType}></CardRow>
        <CardRow descr="Язык" value={atlas?.Language}></CardRow>
        <CardRow descr="Multimonitor" value={atlas?.Multimonitor}></CardRow>
        <CardRow descr={atlasUser?.User ? atlasUser?.Role : 'Пользователь'} 
          value={atlasUser?.User ? atlasUser?.User : 'Не авторизован'}
          rightWidth={'100%'}
        ></CardRow>        
        <CardRow descr="Ошибки" value={''}></CardRow>
        {software?.AtlasErrorDescriptions?.length ? 
          software.AtlasErrorDescriptions.map((i, ind) => (
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