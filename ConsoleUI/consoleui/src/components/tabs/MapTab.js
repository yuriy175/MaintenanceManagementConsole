import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
// import { Map as LeafletMap, TileLayer, Marker, Popup } from 'react-leaflet';
// import { Map as LeafletMap, TileLayer, Marker, Popup } from 'react-leaflet';
import { MapContainer, TileLayer, Marker, Popup } from 'react-leaflet';

const useStyles = makeStyles((theme) => ({
  root: {
    // display:"flex",
    width: '1500px',
    height: '850px',
  },
}));

export default function MapTab(props) {
  console.log("render MapTab");

  const classes = useStyles();

  return (
    <div className={classes.root}>
      <MapContainer
        // className="markercluster-map"
        className={classes.root}
        center={[55.75222, 37.61556]}
        zoom={10}
        maxZoom={18}
      >
        <TileLayer
          url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
          attribution='&copy; <a href="http://osm.org/copyright">OpenStreetMap</a> contributors'
        />
        <Marker position={[55.61980, 37.65602]}>
          <Popup>
            ГБУЗ "ГКБ им. В.М. Буянова ДМЗ"
          </Popup>
        </Marker>
        <Marker position={[55.76273, 37.79743]}>
          <Popup>
            ГБУЗ МКНЦ им. А.С.Логинова ДЗМ
          </Popup>
        </Marker>
      </MapContainer>
      {/* <LeafletMap
        center={[50, 10]}
        zoom={6}
        maxZoom={10}
        attributionControl={true}
        zoomControl={true}
        doubleClickZoom={true}
        scrollWheelZoom={true}
        dragging={true}
        animate={true}
        easeLinearity={0.35}
      >
        <TileLayer
          url='http://{s}.tile.osm.org/{z}/{x}/{y}.png'
        />
        <Marker position={[50, 10]}>
          <Popup>
            Popup for any custom information.
          </Popup>
        </Marker>
      </LeafletMap> */}
    </div>
  );
}