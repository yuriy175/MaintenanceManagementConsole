import React, {useContext} from 'react';
import { makeStyles } from '@material-ui/core/styles';

import { AllEquipsContext } from '../../context/allEquips-context';
import { CurrentEquipContext } from '../../context/currentEquip-context';
import EquipTable from '../tables/equipTables/EquipTable'

const useStyles = makeStyles((theme) => ({
  root: {
    // display:"flex"
  },
}));

export default function EquipsTab(props) {
  console.log("render EquipsTab");

  const classes = useStyles();  
  
  const [currEquipState, currEquipDispatch] = useContext(CurrentEquipContext);
  const [allEquipsState, allEquipsDispatch] = useContext(AllEquipsContext);

  return (
    <div className={classes.root}>
      <EquipTable data={allEquipsState.allEquips} equipInfo={currEquipState.equipInfo}></EquipTable>
    </div>
  );
}