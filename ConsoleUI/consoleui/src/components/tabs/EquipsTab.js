import React, {useContext} from 'react';
import { makeStyles } from '@material-ui/core/styles';

import { AllEquipsContext } from '../../context/allEquips-context';
import EquipTable from '../tables/equipTables/EquipTable'

const useStyles = makeStyles((theme) => ({
  root: {
    display:"flex"
  },
}));

export default function EquipsTab(props) {
  console.log("render EquipsTab");

  const classes = useStyles();  
  const [allEquipsState, allEquipsDispatch] = useContext(AllEquipsContext);

  return (
    <div className={classes.root}>
      <EquipTable data={allEquipsState.allEquips}></EquipTable>
    </div>
  );
}