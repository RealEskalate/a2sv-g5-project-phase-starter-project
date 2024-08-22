import React from "react";
import {
  Button,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Typography,
  Card,
  CardContent,
} from "@mui/material";

const loans = [
  {
    slNo: "1",
    loanMoney: "$10,000",
    leftToRepay: "$5,000",
    duration: "12 months",
    interestRate: "5%",
    installment: "$500 / month",
    repay: "Repay",
  },
  {
    slNo: "2",
    loanMoney: "$15,000",
    leftToRepay: "$7,500",
    duration: "24 months",
    interestRate: "4%",
    installment: "$600 / month",
    repay: "Repay",
  },
  {
    slNo: "3",
    loanMoney: "$20,000",
    leftToRepay: "$10,000",
    duration: "36 months",
    interestRate: "3.5%",
    installment: "$700 / month",
    repay: "Repay",
  },
  {
    slNo: "4",
    loanMoney: "$25,000",
    leftToRepay: "$12,500",
    duration: "48 months",
    interestRate: "3%",
    installment: "$800 / month",
    repay: "Repay",
  },
  {
    slNo: "5",
    loanMoney: "$30,000",
    leftToRepay: "$15,000",
    duration: "60 months",
    interestRate: "2.5%",
    installment: "$900 / month",
    repay: "Repay",
  },
  {
    slNo: "6",
    loanMoney: "$35,000",
    leftToRepay: "$17,500",
    duration: "72 months",
    interestRate: "2%",
    installment: "$1,000 / month",
    repay: "Repay",
  },
  {
    slNo: "7",
    loanMoney: "$40,000",
    leftToRepay: "$20,000",
    duration: "84 months",
    interestRate: "1.5%",
    installment: "$1,100 / month",
    repay: "Repay",
  },
  {
    slNo: "8",
    loanMoney: "$45,000",
    leftToRepay: "$22,500",
    duration: "54 months",
    interestRate: "2.5%",
    installment: "$1,200 / month",
    repay: "Repay",
  },
];

// Calculate totals
const totalLoanMoney = loans.reduce(
  (sum, loan) => sum + parseFloat(loan.loanMoney.replace(/[^0-9.-]+/g, "")),
  0
);
const totalLeftToRepay = loans.reduce(
  (sum, loan) => sum + parseFloat(loan.leftToRepay.replace(/[^0-9.-]+/g, "")),
  0
);
const totalInstallment = loans.reduce(
  (sum, loan) => sum + parseFloat(loan.installment.replace(/[^0-9.-]+/g, "")),
  0
);

// Add the total row
loans.push({
  slNo: "Total",
  loanMoney: `$${totalLoanMoney.toLocaleString()}`,
  leftToRepay: `$${totalLeftToRepay.toLocaleString()}`,
  duration: "",
  interestRate: "",
  installment: `$${totalInstallment.toLocaleString()} / month`,
  repay: "",
});

const ActiveLoans: React.FC = () => {
  return (
    <Card style={{ margin: "20px", borderRadius: "20px" }}>
      <CardContent>
        <Typography variant="h5" style={{ margin: "20px 0" }}>
          Active Loans Overview
        </Typography>
        <TableContainer>
          <Table>
            <TableHead>
              <TableRow>
                <TableCell style={{ color: "#718EBF" }}>SL No</TableCell>
                <TableCell style={{ color: "#718EBF" }}>Loan Money</TableCell>
                <TableCell style={{ color: "#718EBF" }}>
                  Left to Repay
                </TableCell>
                <TableCell style={{ color: "#718EBF" }}>Duration</TableCell>
                <TableCell style={{ color: "#718EBF" }}>
                  Interest Rate
                </TableCell>
                <TableCell style={{ color: "#718EBF" }}>Installment</TableCell>
                <TableCell style={{ color: "#718EBF" }}>Repay</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {loans.map((loan, index) => (
                <TableRow key={index}>
                  <TableCell
                    style={loan.slNo === "Total" ? { color: "red" } : {}}
                  >
                    {loan.slNo}
                  </TableCell>
                  <TableCell
                    style={loan.slNo === "Total" ? { color: "red" } : {}}
                  >
                    {loan.loanMoney}
                  </TableCell>
                  <TableCell
                    style={loan.slNo === "Total" ? { color: "red" } : {}}
                  >
                    {loan.leftToRepay}
                  </TableCell>
                  <TableCell
                    style={loan.slNo === "Total" ? { color: "red" } : {}}
                  >
                    {loan.duration}
                  </TableCell>
                  <TableCell
                    style={loan.slNo === "Total" ? { color: "red" } : {}}
                  >
                    {loan.interestRate}
                  </TableCell>
                  <TableCell
                    style={loan.slNo === "Total" ? { color: "red" } : {}}
                  >
                    {loan.installment}
                  </TableCell>
                  <TableCell
                    style={loan.slNo === "Total" ? { color: "red" } : {}}
                  >
                    {loan.repay && (
                      <Button
                        variant="outlined"
                        style={{ borderRadius: "20px" }}
                      >
                        {loan.repay}
                      </Button>
                    )}
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
      </CardContent>
    </Card>
  );
};

export default ActiveLoans;
