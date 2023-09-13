import Appbar from "./Appbar.tsx";
import Sidebar from "./Sidebar.tsx";
import { FCWithChildren } from "../../types/react.ts";
import { Box } from "@mui/material";

const BaseLayout: FCWithChildren = ({ children }) => {
  return (
    <Box sx={{ m: 0 }}>
      <Appbar variant="" />
      {children}
      <Sidebar />
    </Box>
  );
};

export default BaseLayout;
