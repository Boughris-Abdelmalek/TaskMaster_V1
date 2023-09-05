import Logo from "./Logo.tsx";
import { FC } from "react";
import { AppBar, Box, Button, Container, Toolbar } from "@mui/material";
import theme from "../material-theme.ts";

type NavBarProps = {
  variant?: "transparent" | "";
};

const pages = ["login", "signup"];

const Navbar: FC<NavBarProps> = ({ variant = "" }) => {
  return (
    <AppBar position="static" sx={{ p: 5 }}>
      <Container maxWidth="xl">
        <Toolbar disableGutters>
          <Logo variant={variant === "transparent" ? "" : "white"} />
          <Box sx={{ display: "flex", columnGap: 3 }}>
            {pages.map((page) => (
              <Button
                key={page}
                variant="text"
                sx={{
                  color: "black",
                  display: "block",
                  fontSize: "1.25rem",

                  "&:hover": {
                    backgroundColor: "transparent",
                    color: theme.palette.text.primary,
                  },
                }}
                href={`/${page}`}
              >
                {page}
              </Button>
            ))}
          </Box>
        </Toolbar>
      </Container>
    </AppBar>
  );
};

export default Navbar;
