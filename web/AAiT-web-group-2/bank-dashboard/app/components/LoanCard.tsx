import React from "react";
import { Card, CardContent, Typography, Box } from "@mui/material";

interface LoanCardProps {
  icon: React.ReactNode;
  title: string;
  subtitle: string;
  background: string;
}

const LoanCard: React.FC<LoanCardProps> = ({
  icon,
  title,
  subtitle,
  background,
}) => {
  return (
    <Card
      className="shadow-lg p-4"
      style={{
        margin: "5px",
        width: "220px",
        height: "115px",
        borderRadius: "30px",
        overflow: "hidden",
        display: "flex",
        alignItems: "center",
      }}
    >
      <CardContent
        style={{
          display: "flex",
          alignItems: "center",
          padding: "0 8px",
          height: "100%",
          width: "100%",
        }}
      >
        <Box
          component="div"
          style={{
            width: "55px",
            height: "55px",
            backgroundImage: `url(${background})`,
            backgroundSize: "contain",
            display: "flex",
            justifyContent: "center",
            alignItems: "center",
            marginRight: "8px",
            borderRadius: "50%",
          }}
        >
          <div
            style={{
              width: "45px",
              height: "45px",
              display: "flex",
              justifyContent: "center",
              alignItems: "center",
            }}
          >
            {icon}
          </div>
        </Box>
        <div style={{ overflow: "hidden", width: "calc(100% - 60px)" }}>
          <Typography
            variant="subtitle2"
            style={{
              color: "gray",
              fontSize: "0.8rem",
              whiteSpace: "nowrap",
              overflow: "hidden",
              textOverflow: "ellipsis",
            }}
          >
            {title}
          </Typography>
          <Typography
            variant="subtitle1"
            style={{
              fontWeight: "bold",
              fontSize: "0.9rem",
              whiteSpace: "nowrap",
              overflow: "hidden",
              textOverflow: "ellipsis",
            }}
          >
            {subtitle}
          </Typography>
        </div>
      </CardContent>
    </Card>
  );
};

export default LoanCard;
