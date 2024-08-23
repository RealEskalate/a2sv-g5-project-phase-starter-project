import React from "react";
import "@/styles/globals.css";
import Nav from "@/components/Nav";

const layout = ({ children }) => {
  return (
    <html>
      <body>
        <div>
          <Nav />
          {children}
        </div>
      </body>
    </html>
  );
};

export default layout;
