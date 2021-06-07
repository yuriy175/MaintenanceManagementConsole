import react, { useContext } from 'react';
import { CurrentEquipContext } from '../context/currentEquip-context';
import { SystemVolatileContext } from '../context/systemVolatile-context';
import * as EquipWorker from '../workers/equipWorker'

export function useSetCurrEquip() {
  console.log("useSetCurrEquip");
  const [currEquipState, currEquipDispatch] = useContext(CurrentEquipContext);
  const [systemVolatileState, systemVolatileDispatch] = useContext(SystemVolatileContext);

  const setEquipInfo = async (equipInfo, type) => {

    //const equipInfo = row.EquipName;
    systemVolatileDispatch({ type: 'RESET', payload: true });    
    currEquipDispatch({ type: 'RESET', payload: true });    
    currEquipDispatch({ type: type, payload: equipInfo }); 

    // new software & system info come very slowly
    const sysInfo = await EquipWorker.GetPermanentData("SystemInfo", equipInfo);
    currEquipDispatch({ type: 'SETSYSTEM', payload: sysInfo[0] }); 

    const swInfo = await EquipWorker.GetPermanentData("Software", equipInfo);
    currEquipDispatch({ type: 'SETSOFTWARE', 
      payload: {
        Atlas: swInfo.Atlas ? swInfo.Atlas[0] : undefined, 
        Software: swInfo.Software ? swInfo.Software[0] : undefined} }); 

    await EquipWorker.Activate(equipInfo, currEquipState.equipInfo);
  };

  return setEquipInfo;
}
