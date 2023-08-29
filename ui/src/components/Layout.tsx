import Appbar from "./Appbar.tsx";
import Sidebar from "./Sidebar.tsx";
import { FCWithChildren } from "../../types/react.ts";

const BaseLayout: FCWithChildren = ({ children }) => {
  return (
    <>
      <Appbar />
      {children}
      <Sidebar />
    </>
  );
};

export default BaseLayout;
