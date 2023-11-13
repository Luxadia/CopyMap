import * as React from "react";
import Box from "@mui/material/Box";
import Drawer from "@mui/material/Drawer";
import AppBar from "@mui/material/AppBar";
import CssBaseline from "@mui/material/CssBaseline";
import Toolbar from "@mui/material/Toolbar";
import Typography from "@mui/material/Typography";
import FormatAlignCenterIcon from "@mui/icons-material/FormatAlignCenter";
import ToggleButton from "@mui/material/ToggleButton";
import ToggleButtonGroup from "@mui/material/ToggleButtonGroup";
import SearchIcon from "@mui/icons-material/Search";
import { styled, alpha } from "@mui/material/styles";
import InputBase from "@mui/material/InputBase";
import FilterIcon from "@mui/icons-material/Filter";
import FilterVintageTwoToneIcon from "@mui/icons-material/FilterVintageTwoTone";
import SettingsInputComponentIcon from "@mui/icons-material/SettingsInputComponent";

import { useState } from "react";
import List from "@mui/material/List";
import ListItem from "@mui/material/ListItem";
import ListItemText from "@mui/material/ListItemText";
import Checkbox from "@mui/material/Checkbox";
import Collapse from "@mui/material/Collapse";
import ListItemIcon from "@mui/material/ListItemIcon";
import ExpandLess from "@mui/icons-material/ExpandLess";
import ExpandMore from "@mui/icons-material/ExpandMore";

import PhotonSearchBox from "./PhotonSearch";
const drawerWidth = 240;



const LayerSection = ({ title, items }) => {
    const [open, setOpen] = useState(true);
    const [selectAll, setSelectAll] = useState(false);
    const [selectedItems, setSelectedItems] = useState([]);
  
    const handleToggleAll = () => {
      setSelectAll(!selectAll);
      setSelectedItems(selectAll ? [] : items.map((item) => item.id));
    };
  
    const handleClick = () => {
      setOpen(!open);
    };
  
    const handleToggleItem = (itemId) => {
      const index = selectedItems.indexOf(itemId);
      if (index === -1) {
        setSelectedItems([...selectedItems, itemId]);
      } else {
        setSelectedItems([...selectedItems.slice(0, index), ...selectedItems.slice(index + 1)]);
      }
    };
  
    return (
      <>
        <ListItem dense onClick={handleClick} sx={{ paddingLeft: 0 }}>
          <ListItemIcon>
            <Checkbox
              edge="start"
              checked={selectAll}
              onChange={handleToggleAll}
            />
          </ListItemIcon>
          <ListItemText primary={title} />
          {open ? <ExpandLess /> : <ExpandMore />}
        </ListItem>
        <Collapse in={open} timeout="auto" unmountOnExit>
          <List sx={{p:"0px"}} dense>
            {items.map((item) => (
              <ListItem key={item.id} sx={{ paddingLeft: 4 }}>
                <ListItemIcon>
                  <Checkbox
                    size="small" 
                    edge="start"
                    checked={selectedItems.includes(item.id)}
                    onChange={() => handleToggleItem(item.id)}
                    sx={{padding:"0px"}}
                  />
                </ListItemIcon>
                <ListItemText sx={{p:"0px"}} primary={item.label} />
              </ListItem>
            ))}
          </List>
        </Collapse>
      </>
    );
  };
  
  const FilterSidebar = () => {
    const layerSections = [
      {
        title: 'Section 1',
        items: [
          { id: 1, label: 'Item 1.1' },
          { id: 2, label: 'Item 1.2' },
        ],
      },
      {
        title: 'Section 2',
        items: [
          { id: 3, label: 'Item 2.1' },
          { id: 4, label: 'Item 2.2' },
        ],
      },
    ];
  
    return (
      <List component="nav" aria-labelledby="nested-list-subheader" dense>
        {layerSections.map((section) => (
          <LayerSection key={section.title} {...section} />
        ))}
      </List>
    );
  };

// function FilterSidebar() {
//   return <p>filter content here</p>;
// }

function SettingsSidebar() {
  return <p>setting content here</p>;
}

function SidebarContent(props) {
  const input = props.input;
  if (input == "Filter") {
    return FilterSidebar();
  } else {
    return SettingsSidebar();
  }
}

export default function Nav({ goToLocation, children }) {
  const [open, setOpen] = React.useState(false);
  const [alignment, setAlignment] = React.useState(null);
  const handleAlignment = (event, newAlignment) => {
    setAlignment(newAlignment);
    if (newAlignment == null) {
      setOpen(false);
    } else {
      setOpen(true);
    }
  };
  return (
    <Box sx={{ display: "flex", position: "relative" }}>
      <CssBaseline />
      <AppBar
        position="fixed"
        sx={{ zIndex: (theme) => theme.zIndex.drawer + 1 }}
      >
        <Toolbar>
          <FilterVintageTwoToneIcon />
          <Typography variant="h6" noWrap component="div">
            MapApp
          </Typography>
          <PhotonSearchBox goToLocation={goToLocation}/>
          <ToggleButtonGroup
            size="small"
            value={alignment}
            exclusive
            onChange={handleAlignment}
            aria-label="text alignment"
            sx={{ml:"6px"}}
          >
            <ToggleButton
              value="Filter"
              aria-label="filter"
              sx={{ border: "None", borderRadius: 0 }}
            >
              <FilterIcon />
            </ToggleButton>
            <ToggleButton
              value="Settings"
              aria-label="settings"
              sx={{ border: "None", borderRadius: 0 }}
            >
              <SettingsInputComponentIcon />
            </ToggleButton>
          </ToggleButtonGroup>
        </Toolbar>
      </AppBar>
      <Drawer
        variant="persistent"
        open={open}
        sx={{
          width: drawerWidth,
          flexShrink: 0,
          [`& .MuiDrawer-paper`]: {
            width: drawerWidth,
            boxSizing: "border-box",
          },
        }}
      >
        <Toolbar />
        <Box
          sx={{
            overflow: "auto",
            pl: 3,
          }}
        >
          <SidebarContent input={alignment} />
        </Box>
      </Drawer>
      <Box component="main" sx={{ flexGrow: 1, p: 3 }}>
        <Toolbar />
        <div style={{ position: "relative" }}>{children}</div>
      </Box>
    </Box>
  );
}
