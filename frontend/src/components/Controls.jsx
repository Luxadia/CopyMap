import * as React from "react";
import IconButton from "@mui/material/IconButton";
import ButtonGroup from "@mui/material/ButtonGroup";
import DeleteIcon from "@mui/icons-material/Delete";
import AlarmIcon from "@mui/icons-material/Alarm";
import AddShoppingCartIcon from "@mui/icons-material/AddShoppingCart";
import Divider from "@mui/material/Divider";

export default function IconButtons() {
  return (
    <ButtonGroup
      variant="contained"
      aria-label="outlined primary button group"
      orientation="vertical"
      sx={{
        position: "absolute",
        top: "0px",
        right: "0px",
        background: "black",
      }}
    >
      <IconButton aria-label="delete">
        <DeleteIcon />
      </IconButton>
      <IconButton aria-label="delete">
        <DeleteIcon />
      </IconButton>
      <Divider />
      <IconButton color="secondary">
        <AlarmIcon />
      </IconButton>
      <IconButton color="primary">
        <AddShoppingCartIcon />
      </IconButton>
    </ButtonGroup>
  );
}
