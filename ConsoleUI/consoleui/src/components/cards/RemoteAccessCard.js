import React, {useContext} from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import Button from '@material-ui/core/Button';
import Typography from '@material-ui/core/Typography';

import {useCardsStyles} from './CommonCard'
import { CurrentEquipContext } from '../../context/currentEquip-context';
import * as EquipWorker from '../../workers/equipWorker'

const RemoteAccessCard = React.memo((props) => {
// export default function RemoteAccessCard() {
  console.log(`! render RemoteAccessCard`);

  const classes = useCardsStyles();
  const bull = <span className={classes.bullet}>•</span>;

  const equipInfo = props.equipInfo;
  const onRunTV = async () => {
    const res = await EquipWorker.RunTeamViewer(equipInfo);
  };

  return (
    <Card className={classes.root}>
      <CardContent>
        <Typography variant="h5" component="h2">
          {bull}Удаленный доступ
        </Typography>
        <Typography className={classes.pos} color="textSecondary">
          TeamViewer
        </Typography>
        <Button variant="contained" color="primary" className={classes.commonSpacing} onClick={onRunTV}>
          Запустить
        </Button>
      </CardContent>
    </Card>
  );
});

export default RemoteAccessCard;