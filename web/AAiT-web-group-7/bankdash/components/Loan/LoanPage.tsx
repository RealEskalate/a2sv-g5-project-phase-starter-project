import React from "react";
import {
  Table,
  TableBody,
  TableCaption,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import LoanCard from "@/components/Loan/LoanCard";

const LoanPage = () => {
  return (
    <div className="w-fill h-fill flex flex-col gap-[26px] p-10">
      <div className="w-fill h-fill flex flex-row gap-[26px] ">
        <LoanCard
          imageSrc="pubimg/user.svg"
          title="Personal Loans"
          amount="$100,000"
          bgColor="#E7EDFF"
        />
        <LoanCard
          imageSrc="pubimg/briefcase.svg"
          title="Corporate Loans"
          amount="$100,000"
          bgColor="#FFF5D9"
        />
        <LoanCard
          imageSrc="pubimg/Group.svg"
          title="Business Loans"
          amount="$500,000"
          bgColor="#FFE0EB"
        />
        <LoanCard
          imageSrc="pubimg/support.svg"
          title="Custom Loans"
          amount="Choose Money"
          bgColor="#DCFAF8"
        />
      </div>
      <div className="text-[#333B69] w-fit font-bold">
        Active Loans Overview
      </div>
      <div className="bg-white px-[20px] py-[10px] w-[1110px] border rounded-3xl">
        <Table>
          {/* <TableCaption>A list of your recent invoices.</TableCaption> */}
          <TableHeader>
            <TableRow>
              <TableHead className="text-[#718EBF]">SL No</TableHead>
              <TableHead className="text-[#718EBF]">Loan Money</TableHead>
              <TableHead className="text-[#718EBF]">Left to repay</TableHead>
              <TableHead className="text-[#718EBF]">Duration</TableHead>
              <TableHead className="text-[#718EBF]">Interset rate</TableHead>
              <TableHead className="text-[#718EBF]">Installment</TableHead>
              <TableHead className="text-[#718EBF]">Repay</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            <TableRow>
              <TableCell>01.</TableCell>
              <TableCell>$100,000</TableCell>
              <TableCell>$40,500</TableCell>
              <TableCell>8 Months</TableCell>
              <TableCell>12%</TableCell>
              <TableCell>$2,000 / month</TableCell>
              <TableCell>
                <button className="border-[1px] h-[35px] w-[100px] p-2 rounded-full text-[#232323]">
                  Repay
                </button>
              </TableCell>
            </TableRow>
            <TableRow>
              <TableCell>02.</TableCell>
              <TableCell>$100,000</TableCell>
              <TableCell>$40,500</TableCell>
              <TableCell>8 Months</TableCell>
              <TableCell>12%</TableCell>
              <TableCell>$2,000 / month</TableCell>
              <TableCell>
                <button className="border-[1px] h-[35px] w-[100px] p-2 rounded-full text-[#232323]">
                  Repay
                </button>
              </TableCell>
            </TableRow>
            <TableRow>
              <TableCell>02.</TableCell>
              <TableCell>$100,000</TableCell>
              <TableCell>$40,500</TableCell>
              <TableCell>8 Months</TableCell>
              <TableCell>12%</TableCell>
              <TableCell>$2,000 / month</TableCell>
              <TableCell>
                <button className="border-[1px] h-[35px] w-[100px] p-2 rounded-full text-[#232323]">
                  Repay
                </button>
              </TableCell>
            </TableRow>
            <TableRow>
              <TableCell>02.</TableCell>
              <TableCell>$100,000</TableCell>
              <TableCell>$40,500</TableCell>
              <TableCell>8 Months</TableCell>
              <TableCell>12%</TableCell>
              <TableCell>$2,000 / month</TableCell>
              <TableCell>
                <button className="border-[1px] h-[35px] w-[100px] p-2 rounded-full text-[#232323]">
                  Repay
                </button>
              </TableCell>
            </TableRow>
            <TableRow>
              <TableCell>02.</TableCell>
              <TableCell>$100,000</TableCell>
              <TableCell>$40,500</TableCell>
              <TableCell>8 Months</TableCell>
              <TableCell>12%</TableCell>
              <TableCell>$2,000 / month</TableCell>
              <TableCell>
                <button className="border-[1px] h-[35px] w-[100px] p-2 rounded-full text-[#232323]">
                  Repay
                </button>
              </TableCell>
            </TableRow>
            <TableRow>
              <TableCell>02.</TableCell>
              <TableCell>$100,000</TableCell>
              <TableCell>$40,500</TableCell>
              <TableCell>8 Months</TableCell>
              <TableCell>12%</TableCell>
              <TableCell>$2,000 / month</TableCell>
              <TableCell>
                <button className="border-[1px] h-[35px] w-[100px] p-2 rounded-full text-[#232323]">
                  Repay
                </button>
              </TableCell>
            </TableRow>
            <TableRow>
              <TableCell>02.</TableCell>
              <TableCell>$100,000</TableCell>
              <TableCell>$40,500</TableCell>
              <TableCell>8 Months</TableCell>
              <TableCell>12%</TableCell>
              <TableCell>$2,000 / month</TableCell>
              <TableCell>
                <button className="border-[1px] h-[35px] w-[100px] p-2 rounded-full text-[#232323]">
                  Repay
                </button>
              </TableCell>
            </TableRow>
            <TableRow>
              <TableCell>02.</TableCell>
              <TableCell>$100,000</TableCell>
              <TableCell>$40,500</TableCell>
              <TableCell>8 Months</TableCell>
              <TableCell>12%</TableCell>
              <TableCell>$2,000 / month</TableCell>
              <TableCell>
                <button className="border-[1px] h-[35px] w-[100px] p-2 rounded-full text-[#232323]">
                  Repay
                </button>
              </TableCell>
            </TableRow>

            <TableRow className="text-[#FE5C73] font-semibold">
              <TableCell>Total</TableCell>
              <TableCell>$125000</TableCell>
              <TableCell>$750,000</TableCell>
              <TableCell></TableCell>
              <TableCell></TableCell>
              <TableCell>$50,000 / month</TableCell>
            </TableRow>
          </TableBody>
        </Table>
        {/* <div className="flex justify-between w-[889px] h-[19px] text-[#FE5C73] ">
          <div className="gap-20 flex flex-row pl-4">
            <div>Total</div>
            <div>$125000</div>
            <div>$750,000</div>
          </div>
          <div className="">$50,000 / month</div>
        </div> */}
      </div>
    </div>
  );
};

export default LoanPage;
