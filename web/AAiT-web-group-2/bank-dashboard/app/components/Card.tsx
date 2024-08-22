import React from "react";
import { Grid } from "@mui/material";
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
    icon: (
      <img
        src={icon1.src}
        alt="icon1"
        style={{ width: "30px", height: "30px" }}
      />
    ),
    title: "Personal Loans",
    subtitle: "$50,000",
    background: Ellipse1.src,
  },
  {
    icon: (
      <img
        src={icon2.src}
        alt="icon2"
        style={{ width: "30px", height: "30px" }}
      />
    ),
    title: "Corporation Loans",
    subtitle: "$100,000",
    background: Ellipse2.src,
  },
  {
    icon: (
      <img
        src={icon3.src}
        alt="icon3"
        style={{ width: "30px", height: "30px" }}
      />
    ),
    title: "Business Loans",
    subtitle: "$500,000",
    background: Ellipse3.src,
  },
  {
    icon: (
      <img
        src={icon4.src}
        alt="icon4"
        style={{ width: "30px", height: "30px" }}
      />
    ),
    title: "Custom Loans",
    subtitle: "Choose Money",
    background: Ellipse4.src,
  },
];

const Cards: React.FC = () => {
  return (
    <Grid container spacing={3}>
      {cards.map((card, index) => (
        <Grid item xs={12} sm={6} md={3} key={index}>
          <LoanCard
            icon={card.icon}
            title={card.title}
            subtitle={card.subtitle}
            background={card.background}
          />
        </Grid>
      ))}
    </Grid>
  );
};

export default Cards;
