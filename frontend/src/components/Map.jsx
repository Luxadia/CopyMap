import React, { useState, useCallback, useEffect } from "react";
import { DeckGL, GeoJsonLayer, ScatterplotLayer } from "deck.gl";

import { BitmapLayer } from "@deck.gl/layers";
import { TileLayer } from "@deck.gl/geo-layers";

// DeckGL react component
function Map(props) {

  const url = 'output.json'
  //const url = 'https://raw.githubusercontent.com/visgl/deck.gl-data/master/examples/collision-filter/ne_10_roads_mexico.json';
  const [data, setdata] = useState([])
  useEffect(() => {
    fetch(url)
      .then(r => r.json())
      .then(data3 => setdata(data3.features));
  }, [url]);


  const tile1 = new TileLayer({
    data: "http://www.basemaps.cartocdn.com/light_all/{z}/{x}/{y}.png",

    minZoom: 0,
    maxZoom: 19,
    tileSize: 256,

    renderSubLayers: (sprops) => {
      const {
        bbox: { west, south, east, north },
      } = sprops.tile;

      return new BitmapLayer(sprops, {
        data: null,
        image: sprops.data,
        bounds: [west, south, east, north],
      });
    }})

  const tile2 = new GeoJsonLayer({
    data: data,
    getLineColor: [255, 0, 0],
    lineWidthMinPixels: 4,
  })

  console.log(tile2)

  const layers = [tile1, tile2]




  const INITIAL_VIEW_STATE = props.initialViewState;

  return (
    <DeckGL
      initialViewState={INITIAL_VIEW_STATE}
      controller={true}
      layers={layers} />);

  }
export default Map;


// //exmaple tooltip
//   new TileLayer({
//     // https://wiki.openstreetmap.org/wiki/Slippy_map_tilenames#Tile_servers
//     data: "https://c.tile.openstreetmap.org/{z}/{x}/{y}.png",

//     minZoom: 0,
//     maxZoom: 19,
//     tileSize: 256,

//     renderSubLayers: (sprops) => {
//       const {
//         bbox: { west, south, east, north },
//       } = sprops.tile;

//       return new BitmapLayer(sprops, {
//         data: null,
//         image: sprops.data,
//         bounds: [west, south, east, north],
//       });
//     }}),
//     new ScatterplotLayer({
//       data,
//       getPosition: (d) => d.position,
//       getRadius: 10000,
//       getFillColor: [255, 0, 0],
//       // Enable picking
//       pickable: true,
//       // Update app state
//       onHover: (info) => setHoverInfo(info),
//     }),
//   ];

//   const INITIAL_VIEW_STATE = props.initialViewState;

//   return (
//     <DeckGL
//       initialViewState={{ longitude: -122.45, latitude: 27.78, zoom: 12 }}
//       controller={true}
//       layers={layers}
//       getTooltip={({ object }) => object && object.message}
//     />
//   );