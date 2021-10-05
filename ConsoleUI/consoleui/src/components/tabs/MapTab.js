import React, {useContext} from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';
import * as L from "leaflet";
import { MapContainer, TileLayer, Marker, Popup } from 'react-leaflet';

import { SummaryTabIndex, MainTabPanelIndex } from '../../model/constants';
import { AppContext } from '../../context/app-context';
import { AllEquipsContext } from '../../context/allEquips-context';
import {useSetCurrEquip} from '../../hooks/useSetCurrEquip'

const useStyles = makeStyles((theme) => ({
  root: {
    // display:"flex",
    // width: '1500px',
    height: '1000px',
  },
}));

const LeafIcon = L.Icon.extend({
  options: {}
});

const activeIcon = new LeafIcon({
  iconUrl:"./marker-icon.png"
}),
grayIcon = new LeafIcon({
  iconUrl:"./marker-icon-dark.png"
});

export default function MapTab(props) {
  console.log("render MapTab");

  const classes = useStyles();
  const [allEquipsState, allEquipsDispatch] = useContext(AllEquipsContext);
  const [appState, appDispatch] = useContext(AppContext);
  const setCurrEquip = useSetCurrEquip();

  const allEquips = allEquipsState.allEquips;
  const onSelect = async (ev, equip) => {
    const equipInfo = equip.EquipName;
    setCurrEquip(equipInfo, 'SETEQUIPINFO');
    allEquipsDispatch({ type: 'ADDSELECTEDEQUIPS', payload: equipInfo }); 
    appDispatch({ type: 'SETTAB', payload: {tab: SummaryTabIndex, panel: MainTabPanelIndex} }); 
  };

  return (
    <div className={classes.root}>
      <MapContainer
        // className="markercluster-map"
        className={classes.root}
        center={[ 59.8795351,30.3908424 ]}//55.75222, 37.61556]}
        zoom={10}
        maxZoom={18}
      >
        <TileLayer
          url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
          attribution='&copy; <a href="http://osm.org/copyright">OpenStreetMap</a> contributors'
        /> 
        {allEquips?.map((equip) => (
          <Marker position={[equip.HospitalLatitude, equip.HospitalLongitude]} icon={!equip.IsActive? activeIcon : grayIcon}>
            <Popup>
              {equip.HospitalName} ({equip.EquipName})
              <div>
                <Button variant="contained" color="primary" onClick={(ev) => onSelect(ev, equip)}>
                      Выбрать
                </Button>
              </div>
            </Popup>
          </Marker>
          ))} 
        {/* <Marker position={[55.61980, 37.65602]}>
          <Popup>
            ГБУЗ "ГКБ им. В.М. Буянова ДМЗ"
          </Popup>
        </Marker>
        <Marker position={[55.76273, 37.79743]} icon={greenIcon}>
          <Popup>
            ГБУЗ МКНЦ им. А.С.Логинова ДЗМ
          </Popup>
        </Marker> */}
      </MapContainer>
    </div>
  );
}