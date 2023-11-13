import React, {useState, useCallback} from 'react';
import CssBaseline from "@mui/material/CssBaseline";
import Map from "./components/Map";
import Nav from "./components/Nav";
import Controls from "./components/Controls";
import { ThemeProvider, createTheme } from "@mui/material/styles";
import {FlyToInterpolator} from '@deck.gl/core';


const darkTheme = createTheme({
  palette: {
    mode: "dark",
  },
});

function App() {
    const [initialViewState, setInitialViewState] = useState({
    latitude: 37.7751,
    longitude: -122.4193,
    zoom: 11,
    bearing: 0,
    pitch: 0,
  });

  const goToLocation = useCallback((longitude, latitude) => {
    setInitialViewState({
      longitude,
      latitude,
      zoom: 14,
      pitch: 0,
      bearing: 0,
      transitionDuration: 8000,
      transitionInterpolator: new FlyToInterpolator()
    });
  }, [setInitialViewState]);

  return (
    <>
      <ThemeProvider theme={darkTheme}>
        <CssBaseline />
        <Map initialViewState={initialViewState}/>
        <Nav goToLocation={goToLocation}>
          <Controls />
        </Nav>
      </ThemeProvider>
    </>
  );
}

export default App;
