import React from "react";
import { Box, useMediaQuery } from "@mui/material";
import Ellipse1 from "../../public/Ellipse1.png";
import Ellipse2 from "../../public/Ellipse2.png";
import Ellipse3 from "../../public/Ellipse3.png";
import Ellipse4 from "../../public/Ellipse4.png";
import icon1 from "../../public/icon1.png";
import icon2 from "../../public/icon2.png";
import icon3 from "../../public/icon3.png";
import icon4 from "../../public/icon4.png";
import LoanCard from "./LoanCard";

const cards = [
  {
    icon: <img src={icon1.src} alt="icon1" />,
    title: "Personal Loans",
    subtitle: "$50,000",
    background: Ellipse1.src,
  },
  {
    icon: <img src={icon2.src} alt="icon2" />,
    title: "Corporation Loans",
    subtitle: "$100,000",
    background: Ellipse2.src,
  },
  {
    icon: <img src={icon3.src} alt="icon3" />,
    title: "Business Loans",
    subtitle: "$500,000",
    background: Ellipse3.src,
  },
  {
    icon: <img src={icon4.src} alt="icon4" />,
    title: "Custom Loans",
    subtitle: "Choose Money",
    background: Ellipse4.src,
  },
];

const Cards: React.FC = () => {
  const isMobile = useMediaQuery("(max-width: 600px)");

  return (
    <Box
      sx={{
        display: "flex",
        overflowX: isMobile ? "scroll" : "visible",
        scrollSnapType: isMobile ? "x mandatory" : "none",
        paddingLeft: isMobile ? "10px" : "0",
        "&::-webkit-scrollbar": {
          display: "none",
        },
        scrollbarWidth: "none",
      }}
    >
      {cards.map((card, index) => (
        <Box
          key={index}
          sx={{
            flexShrink: 0,
            scrollSnapAlign: "center",
            width: "240px",
            marginRight: isMobile ? "-5px" : "0",
            display: "flex",
            justifyContent: "center",
            alignItems: "center",
          }}
        >
          <LoanCard
            icon={card.icon}
            title={card.title}
            subtitle={card.subtitle}
            background={card.background}
          />
        </Box>
      ))}
    </Box>
  );
};

export default Cards;
