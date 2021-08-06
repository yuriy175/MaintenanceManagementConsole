import react, { useContext } from 'react';
import { CurrentEquipContext } from '../context/currentEquip-context';
import { AllEquipsContext } from '../context/allEquips-context';
//import { SystemVolatileContext } from '../context/systemVolatile-context';
import { UsersContext } from '../context/users-context';
import * as EquipWorker from '../workers/equipWorker'

export function useSetCurrEquip() {
  console.log("useSetCurrEquip");
  const [currEquipState, currEquipDispatch] = useContext(CurrentEquipContext);
  // const [systemVolatileState, systemVolatileDispatch] = useContext(SystemVolatileContext);
  const [allEquipsState, allEquipsDispatch] = useContext(AllEquipsContext);
  const [usersState, usersDispatch] = useContext(UsersContext);

  const setEquipInfo = async (equipInfo, type) => {
    const token = usersState.token;
    const connectedEquip = allEquipsState.connectedEquips?.includes(equipInfo)

    // systemVolatileDispatch({ type: 'RESET', payload: true });    
    currEquipDispatch({ type: 'RESET', payload: true });    
    currEquipDispatch({ type: type, payload: equipInfo }); 

    const fullInfo = await EquipWorker.GetPermanentData(token, "FullInfo", equipInfo);
    currEquipDispatch({ type: 'SETFULLINFO', payload: fullInfo }); 
    // if(connectedEquip){
      await EquipWorker.Activate(token, equipInfo, currEquipState.equipInfo);
    // }
  };

  return setEquipInfo;
}
