'use client';
import * as React from 'react';
import { styled } from '@mui/material/styles';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Paper from '@mui/material/Paper';
import Button from '@mui/material/Button';

const StyledTableCell = styled(TableCell)(({ theme }) => ({
  // Customize the styles if needed
}));

const StyledTableRow = styled(TableRow)(({ theme }) => ({
  // Customize the styles if needed
}));

function createData(
  SL_NO: string,
  Loan_Money: string,
  Left_To_repay: string,
  Duration: string,
  Interest_Rate: string,
  Installment: string,
  Repay: string,
) {
  return { SL_NO, Loan_Money, Left_To_repay, Duration, Interest_Rate, Installment, Repay };
}

const rows = [
  createData('01.', '$10000', '$6000', '24 months', '4.0%', '$500 /month', 'Repay'),
  createData('02.', '$23700', '$9000', '37 months', '4.3%', '$700 /month', 'Repay'),
  createData('03.', '$26200', '$16000', '24 months', '6.0%', '$800 /month', 'Repay'),
  createData('04.', '$30500', '$3700', '67 months', '4.3%', '$900 /month', 'Repay'),
  createData('05.', '$35600', '$16000', '49 months', '3.9%', '$1000 /month', 'Repay'),
  createData('06.', '$40000', '$20000', '36 months', '5.0%', '$1100 /month', 'Repay'),
];

const totalLoanMoney = rows.reduce((sum, row) => sum + parseFloat(row.Loan_Money.replace('$', '')), 0);
const totalLeftToRepay = rows.reduce((sum, row) => sum + parseFloat(row.Left_To_repay.replace('$', '')), 0);
const totalInstallment = rows.reduce((sum, row) => sum + parseFloat(row.Installment.replace('$', '').replace(' /month', '')), 0);

export default function CustomizedTables() {
  return (
    <div className='rounded-[25px]w-full px-2.5 mt-12 ml-2 md:ml-40 mr-2 md:mr-40'>
      <TableContainer component={Paper} className="p-4">
        <Table sx={{ minWidth: 700 }} aria-label="customized table">
          <TableHead>
            <TableRow>
              <StyledTableCell className="text-[#718EBF] hidden md:table-cell">SL NO</StyledTableCell>
              <StyledTableCell className="text-[#718EBF] p-1 md:p-4" align="center">Loan Money</StyledTableCell>
              <StyledTableCell className="text-[#718EBF] p-1 md:p-4" align="center">Left To Repay</StyledTableCell>
              <StyledTableCell className="text-[#718EBF] hidden md:table-cell p-1 md:p-4" align="center">Duration</StyledTableCell>
              <StyledTableCell className="text-[#718EBF] hidden md:table-cell" align="center">Interest Rate</StyledTableCell>
              <StyledTableCell className="text-[#718EBF] hidden md:table-cell" align="center">Installment</StyledTableCell>
              <StyledTableCell className="text-[#718EBF] p-1 md:p-4" align="center">Repay</StyledTableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {rows.map((row) => (
              <StyledTableRow key={row.SL_NO}>
                <StyledTableCell component="th" scope="row" className="hidden md:table-cell">
                  {row.SL_NO}
                </StyledTableCell>
                <StyledTableCell align="center" className='p-1 md:p-4'>{row.Loan_Money}</StyledTableCell>
                <StyledTableCell align="center" className='p-1 md:p-4'>{row.Left_To_repay}</StyledTableCell>
                <StyledTableCell align="center" className="hidden md:table-cell">{row.Duration}</StyledTableCell>
                <StyledTableCell align="center" className="hidden md:table-cell">{row.Interest_Rate}</StyledTableCell>
                <StyledTableCell align="center" className="hidden md:table-cell">{row.Installment}</StyledTableCell>
                <StyledTableCell align="center">
                  <Button
                    className="text-gray-800 font-semibold w-[100px] h-[35px] py-2 px-4 border rounded-full"
                    sx={{
                      borderColor: '#1814F3',
                      borderWidth: '1px',
                      borderStyle: 'solid',
                    }}
                  >
                    {row.Repay}
                  </Button>
                </StyledTableCell>
              </StyledTableRow>
            ))}
            <StyledTableRow>
              <StyledTableCell component="th" scope="row" className="text-[#FE5C73] font-bold hidden md:table-cell" sx={{ border: 'none' }}>
                Total
              </StyledTableCell>
              <StyledTableCell align="center" className="text-[#FE5C73] font-bold" sx={{ border: 'none' }}>
                ${totalLoanMoney}
              </StyledTableCell>
              <StyledTableCell align="center" className="text-[#FE5C73] font-bold" sx={{ border: 'none' }}>
                ${totalLeftToRepay}
              </StyledTableCell>
              <StyledTableCell colSpan={2} sx={{ border: 'none' }} className="hidden md:table-cell" />
              <StyledTableCell align="center" className="text-[#FE5C73] font-bold hidden md:table-cell" sx={{ border: 'none' }}>
                ${totalInstallment} /month
              </StyledTableCell>
              <StyledTableCell sx={{ border: 'none' }} />
            </StyledTableRow>
          </TableBody>
        </Table>
      </TableContainer>
    </div>
  );
}
