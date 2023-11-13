import React, { useState, useEffect } from "react";
import CircularProgress from "@mui/material/CircularProgress";
import { styled, alpha } from "@mui/material/styles";
import { ThemeProvider, createTheme } from "@mui/material/styles";
import SearchIcon from "@mui/icons-material/Search";
import Popover from "@mui/material/Popover";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import Autocomplete from "@mui/material/Autocomplete";
import { Stack, TextField } from "@mui/material";
import InputBase from "@mui/material/InputBase";
import InputAdornment from "@mui/material/InputAdornment";

const PhotonSearchBox = ({goToLocation}) => {
  const [query, setQuery] = useState("");
  const [results, setResults] = useState([]);
  const [loading, setLoading] = useState(false);

  const handleSearch = async (value) => {
    try {
      setLoading(true);
      const response = await fetch(
        `https://photon.komoot.io/api/?q=${value}&limit=5&bbox=3.31497114423,50.803721015,7.09205325687,53.5104033474`
      );
      if (!response.ok) {
        throw new Error(
          `Komoot API request failed with status ${response.status}`
        );
      }
      const data = await response.json();
      setResults(data.features);
    } catch (error) {
      console.error("Error fetching Komoot search results", error);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    if (query.trim() !== "") {
      handleSearch(query);
    } else {
      setResults([]);
    }
  }, [query]);

   const handleChange = (event, value) => {
    // Update the selectedLocation state when an option is selected
    if (value && value.geometry && value.geometry.coordinates) {
      goToLocation(...value.geometry.coordinates)
    }}
  return (
    <Autocomplete
      sx={{
        ml: 5,
        width: "300px",
      }}
      options={results}
      size="small"
      getOptionLabel={(option) => option.properties.name}
      onChange={handleChange} 
      renderInput={(params) => (
        <TextField
          {...params}
          id="search-input"
          variant="filled"
          hiddenLabel
          fullWidth
          value={query}
          onChange={(e) => setQuery(e.target.value)}
          onFocus={() => handleSearch(query)}
          InputProps={{
            ...params.InputProps,
            startAdornment: <InputAdornment position="start"><SearchIcon/></InputAdornment>,
            endadornment: (
              <>
                {loading && <circularprogress color="inherit" size={20} />}
              </>
            ),
          }}
        />
      )}
      renderOption={(props, option) => (
        <li {...props}>
            <Stack variant="vertical">
                <Typography variant="subtitle1">{option.properties.name}</Typography>
                  <Typography variant="body2" color="textSecondary">
                    {option.properties.street}
                  </Typography>
            </Stack>
        </li>
      )}
    />
  );
};
export default PhotonSearchBox;