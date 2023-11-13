import React, {useState, useCallback} from 'react';
import DeckGL from '@deck.gl/react';

import {BitmapLayer} from '@deck.gl/layers';
import {TileLayer} from '@deck.gl/geo-layers';


// DeckGL react component
function Map(props) {
  const INITIAL_VIEW_STATE = props.initialViewState
  const layer = new TileLayer({
    // https://wiki.openstreetmap.org/wiki/Slippy_map_tilenames#Tile_servers
    data: 'https://c.tile.openstreetmap.org/{z}/{x}/{y}.png',

    minZoom: 0,
    maxZoom: 19,
    tileSize: 256,

    renderSubLayers: sprops => {
      const {
        bbox: {west, south, east, north}
      } = sprops.tile;

      return new BitmapLayer(sprops, {
        data: null,
        image: sprops.data,
        bounds: [west, south, east, north]
      });
  }})
    
  return <DeckGL
      initialViewState={INITIAL_VIEW_STATE}
      controller={true}
      layers={[layer]} />;
}

export default Map