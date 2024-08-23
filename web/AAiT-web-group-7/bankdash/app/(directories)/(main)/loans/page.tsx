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

const LoanDashboard = () => {
  return (
    <div className="w-fill h-fill flex flex-col gap-[26px] p-10">
      <div className="w-fill h-fill flex flex-row gap-[26px] ">
        <div className="w-[255px] h-[120px]  rounded-3xl flex justify-center items-center bg-white">
          <div className="flex flex-row w-[201px] h-[70px]  justify-between items-center ">
            <div className="rounded-full flex justify-center w-[70px] h-[70px] p-5 bg-[#E7EDFF]">
              <img src="pubimg/user.svg" />
            </div>
            <div className=" flex flex-col w-[116px] h-[51px] gap-1">
              <div className="font-[intel] text-[#718EBF]">Personal Loans</div>
              <div className="text-[#232323] font-[intel] font-bold">
                $50,000
              </div>
            </div>
          </div>
        </div>
        <div className="w-[255px] h-[120px]  rounded-3xl flex justify-center items-center bg-white">
          <div className="flex flex-row w-[201px] h-[70px]  justify-between items-center ">
            <div className=" rounded-full flex justify-center w-[70px] h-[70px] p-5 bg-[#FFF5D9]">
              <img src="pubimg/briefcase.svg" />
            </div>
            <div className=" flex flex-col w-[116px] h-[51px] gap-1">
              <div className="font-[intel] text-[#718EBF]">Corporate loans</div>
              <div className="text-[#232323] font-[intel] font-bold">
                $100,000
              </div>
            </div>
          </div>
        </div>
        <div className="w-[255px] h-[120px]  rounded-3xl flex justify-center items-center bg-white">
          <div className="flex flex-row w-[201px] h-[70px]  justify-between items-center ">
            <div className=" rounded-full flex justify-center w-[70px] h-[70px] p-5 bg-[#FFE0EB]">
              <img src="pubimg/group.svg" />
            </div>
            <div className=" flex flex-col w-[116px] h-[51px] gap-1">
              <div className="font-[intel] text-[#718EBF]">Business Loans</div>
              <div className="text-[#232323] font-[intel] font-bold">
                $500,000
              </div>
            </div>
          </div>
        </div>
        <div className="w-[255px] h-[120px] rounded-3xl flex justify-center items-center bg-white">
          <div className="flex flex-row w-[201px] h-[70px]  justify-between items-center ">
            <div className=" rounded-full flex justify-center w-[70px] h-[70px] p-5 bg-[#DCFAF8]">
              <img src="pubimg/support.svg" />
            </div>
            <div className=" flex flex-col w-[116px] h-[51px] gap-1">
              <div className="font-[intel] text-[#718EBF]">Custom Loans</div>
              <div className="text-[#232323] font-[intel] font-bold">
                Choose Money
              </div>
            </div>
          </div>
        </div>
      </div>
      <div className="text-[#333B69] w-fit font-bold">
        Active Loans Overview
      </div>
      <div className="bg-white px-[20px] py-[10px] w-[1110px] border rounded-3xl">
        <Table >
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

export default LoanDashboard;
