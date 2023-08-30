// import original module declarations
import "styled-components";

// and extend them!
declare module "styled-components" {
  export interface DefaultTheme {
    colors: {
      primary: "#FF4F5A";
      secondary: "#F9F9F9";
      background: "#ffffff";
      text: "#000000";
      primaryTextColor: "#FF4F5A";
      secondaryTextColor: "#6B6B6B";
      inputPlaceholder: "#ABABAB";
      darkGrayText: "#787878";
      darkText: "#040404";
      black: "#000000";
      white: "#ffffff";
      dark: "";
      medium: "";
      light: "";
      danger: "";
      success: "#66A15A";
    };
    fonts: {
      rowdies: "Rowdies";
    };
    paddings: {};
    margins: {};
  }
}
