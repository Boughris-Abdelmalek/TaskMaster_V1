import { AppBar, Avatar, Box, Button, Toolbar } from "@mui/material";
import Logo from "./Logo.tsx";
import { FC } from "react";

type NavBarProps = {
  variant?: "transparent" | "";
};

const Appbar: FC<NavBarProps> = ({ variant = "" }) => {
  return (
    <AppBar position="static" sx={{ p: 2, backgroundColor: "#FF4F5A" }}>
      <Box>
        <Toolbar disableGutters>
          <Logo variant={variant === "transparent" ? "" : "white"} />
          <Box sx={{ display: "flex" }}>
            <Button color="primary">Malik</Button>
            <Avatar />
          </Box>
        </Toolbar>
      </Box>
    </AppBar>
  );
};

export default Appbar;
