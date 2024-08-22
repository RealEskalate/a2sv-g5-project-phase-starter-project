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
      style={{ margin: "10px", width: "auto", borderRadius: "20px" }}
    >
      <CardContent style={{ display: "flex", alignItems: "center" }}>
        <Box
          component="div"
          style={{
            width: "50px",
            height: "50px",
            backgroundImage: `url(${background})`,
            backgroundSize: "cover",
            display: "flex",
            justifyContent: "center",
            alignItems: "center",
            marginRight: "10px",
          }}
        >
          {icon}
        </Box>
        <div>
          <Typography
            variant="subtitle2"
            style={{ color: "gray", fontSize: "1.2rem" }}
          >
            {title}
          </Typography>
          <Typography
            variant="subtitle1"
            style={{ fontWeight: "bold", fontSize: "1.5rem" }}
          >
            {subtitle}
          </Typography>
        </div>
      </CardContent>
    </Card>
  );
};

export default LoanCard;
