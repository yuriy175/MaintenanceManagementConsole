import React, { useState, useEffect, useRef, useContext } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import FormHelperText from '@material-ui/core/FormHelperText';
import FormControl from '@material-ui/core/FormControl';
import NativeSelect from '@material-ui/core/NativeSelect';

import { AllEquipsContext } from '../../context/allEquips-context';
import { CurrentEquipContext } from '../../context/currentEquip-context';

import * as EquipWorker from '../../workers/equipWorker'

const drawerWidth = 240;

const useStyles = makeStyles((theme) => ({
  appBar: {
    width: `calc(100% - ${drawerWidth}px)`,
    marginLeft: drawerWidth,
  },
  formControl: {
    margin: theme.spacing(1),
    minWidth: 120,
  },
  selectEmpty: {
    marginTop: theme.spacing(2),
    color: "white",
  },
}));

export default function MainToolBar() {
  console.log(`! render MainToolBar`);

  const classes = useStyles();
  const [allEquipsState, allEquipsDispatch] = useContext(AllEquipsContext);
  const [currEquipState, currEquipDispatch] = useContext(CurrentEquipContext);
  const [currEquip, setCurrEquip] = React.useState('none');

  const handleEquipsChange = (event) => {
    const select = event.target;
    const val = select.options[select.selectedIndex].value;

    setCurrEquip(val);
  };

  useEffect(() => {
      (async () => {
          if(allEquipsState.equips !== null)
          {
            return;
          }

          const equips = await EquipWorker.GetAllEquips();
          allEquipsDispatch({ type: 'SETEQUIPS', payload: equips ? equips : [] });   
          if(equips?.length === 0)     
          {
            return;
          }

          const equipInfo = equips[0];
          await EquipWorker.Activate(equipInfo);
          currEquipDispatch({ type: 'SETEQUIPINFO', payload: equipInfo });   
      })();
  }, [allEquipsState.equips]);


  return (    
    <AppBar position="fixed" className={classes.appBar}>
        <Toolbar>
            <Typography variant="h6" noWrap>
            </Typography>
            <FormControl className={classes.formControl}>
              <NativeSelect
                value={currEquipState}
                onChange={handleEquipsChange}
                name="equips"
                className={classes.selectEmpty}
                variant="outlined"
              >
                {allEquipsState.equips?.map((i, ind) => (
                    <option key={ind.toString()} value={i}>{i}</option>
                    ))
                }

                {/* <option value="">None</option>
                <option value={10}>Ten</option>
                <option value={20}>Twenty</option>
                <option value={30}>Thirty</option> */}
              </NativeSelect>
            </FormControl>
        </Toolbar>
    </AppBar>
  );
}