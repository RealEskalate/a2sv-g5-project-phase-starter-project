import Card1 from "./components/Card1";
import Linechart from "./components/LineChart";
import Monthly from "./components/Monthly";
import {
  Table,
  TableBody,
  TableCaption,
  TableCell,
  TableHead,
  TableHeader,
  TableFooter,
  TableRow,
} from "@/app/loans/components/table";

const invoices = [
  {
    slno: "01.",
    name: "Trivago",
    price: "$520",
    return: "+5%",
  },
  {
    slno: "02.",
    name: "Canon",
    price: "$480",
    return: "+10%",
  },
  {
    slno: "03.",
    name: "Uber Food",
    price: "$350",
    return: "-3%",
  },
  {
    slno: "04.",
    name: "Nokia",
    price: "$40,500",
    return: "+2%",
  },
  {
    slno: "05.",
    name: "Tiktok",
    price: "$670",
    return: "-12%",
  },
];

export default function Home() {
  return (
    // <main className="mt-16 ml-72">
    <div className="bg-gray-100 p-6">
      <div className="flex justify-between flex-wrap lg:flex-nowrap">
        <Card1 text="Total Invested Amount" img="/total.png" num="$150,000" />
        <Card1 text="Number of Investments" img="/number.png" num="1,250" />
        <Card1 text="Rate of Return" img="/rate.png" num="+5.80%" />
      </div>
      <div className="grid grid-cols-2 gap-6">
        <div className="col-span-2 lg:col-span-1">
          <div className="my-4 text-2xl font-bold text-[#333B69]">
            Yearly Total Investment
          </div>
          <Linechart />
          {/* <div className="rounded-lg p-1"></div> */}
        </div>
        <div className="col-span-2 lg:col-span-1">
          <div className="my-4 text-2xl font-bold col-span-1 text-[#333B69]">
            Monthly Revenue
          </div>
          <div>
            <Monthly />
          </div>
        </div>
      </div>
      <div className="grid grid-cols-12 gap-8">
        <div className="col-span-12 lg:col-span-7">
          <div className="my-4 text-2xl font-bold text-[#333B69]">
            My Investment
          </div>
          <div>
            <div className="flex bg-white rounded-2xl p-5 justify-between my-3">
              <div>
                <svg
                  width="45"
                  height="45"
                  viewBox="0 0 45 45"
                  fill="none"
                  xmlns="http://www.w3.org/2000/svg"
                >
                  <rect width="45" height="45" rx="15" fill="#FFE0EB" />
                  <path
                    d="M30.4933 28.5861C30.1909 29.2848 29.8328 29.928 29.4181 30.5194C28.8526 31.3255 28.3897 31.8835 28.0329 32.1934C27.4798 32.702 26.8873 32.9625 26.2527 32.9773C25.7972 32.9773 25.2478 32.8477 24.6083 32.5847C23.9667 32.323 23.3771 32.1934 22.838 32.1934C22.2726 32.1934 21.6662 32.323 21.0176 32.5847C20.368 32.8477 19.8446 32.9847 19.4445 32.9983C18.836 33.0242 18.2295 32.7563 17.6241 32.1934C17.2377 31.8563 16.7544 31.2786 16.1754 30.4601C15.5541 29.586 15.0434 28.5725 14.6433 27.417C14.2148 26.1689 14 24.9603 14 23.7902C14 22.4498 14.2896 21.2938 14.8697 20.3251C15.3256 19.547 15.9322 18.9332 16.6913 18.4826C17.4504 18.032 18.2706 17.8023 19.1539 17.7877C19.6372 17.7877 20.271 17.9372 21.0587 18.231C21.8441 18.5258 22.3484 18.6753 22.5695 18.6753C22.7348 18.6753 23.295 18.5005 24.2447 18.152C25.1429 17.8288 25.9009 17.6949 26.5218 17.7477C28.2045 17.8834 29.4687 18.5468 30.3094 19.7418C28.8045 20.6536 28.06 21.9307 28.0749 23.5691C28.0885 24.8452 28.5514 25.9071 29.4612 26.7503C29.8736 27.1417 30.3341 27.4441 30.8464 27.659C30.7353 27.9812 30.618 28.2898 30.4933 28.5861V28.5861ZM26.6342 13.4001C26.6342 14.4003 26.2688 15.3343 25.5404 16.1987C24.6614 17.2263 23.5982 17.8201 22.4453 17.7264C22.4299 17.6007 22.4221 17.4741 22.4221 17.3474C22.4221 16.3872 22.8401 15.3596 23.5824 14.5193C23.953 14.0939 24.4244 13.7402 24.996 13.458C25.5663 13.1801 26.1058 13.0263 26.6132 13C26.628 13.1337 26.6342 13.2674 26.6342 13.4001V13.4001Z"
                    fill="#FF82AC"
                  />
                </svg>
              </div>
              <div className="pl-2">
                <div className="text-sm lg:text-base font-medium t-[#232323]">
                  Apple Store
                </div>
                <div className="text-xs lg:text-sm text-[#718EBF]">
                  E-commerce, Marketplace
                </div>
              </div>
              <div>
                <div className="text-sm lg:text-base font-medium t-[#232323]">
                  $54,000
                </div>
                <div className="text-xs lg:text-sm text-[#718EBF]">
                  Envestment Value
                </div>
              </div>
              <div>
                <div className="text-sm lg:text-base font-medium t-[#232323] text-[#16DBAA]">
                  +16%
                </div>
                <div className="text-xs lg:text-sm text-[#718EBF]">
                  Return Value
                </div>
              </div>
            </div>
            <div className="flex bg-white rounded-2xl p-5 justify-between my-3">
              <div>
                <svg
                  width="45"
                  height="45"
                  viewBox="0 0 45 45"
                  fill="none"
                  xmlns="http://www.w3.org/2000/svg"
                >
                  <rect width="45" height="45" rx="15" fill="#E7EDFF" />
                  <path
                    d="M32.6017 23.2281C32.6026 22.5467 32.545 21.8664 32.4295 21.1948H22.9993V25.0461H28.4005C28.29 25.6612 28.056 26.2475 27.7128 26.7698C27.3695 27.2921 26.9241 27.7394 26.4033 28.0849V30.5848H29.6268C31.5143 28.8446 32.6017 26.271 32.6017 23.2281Z"
                    fill="#4471FF"
                  />
                  <path
                    d="M22.9994 32.9999C25.6979 32.9999 27.97 32.1138 29.6269 30.5861L26.4034 28.0862C25.5062 28.6945 24.3507 29.0417 22.9994 29.0417C20.3911 29.0417 18.1773 27.2834 17.3857 24.9141H14.0649V27.4904C14.8972 29.1465 16.1735 30.5388 17.7511 31.5117C19.3288 32.4846 21.1458 32.9998 22.9994 32.9999Z"
                    fill="#4471FF"
                  />
                  <path
                    d="M17.3855 24.9142C16.967 23.6726 16.967 22.3281 17.3855 21.0866V18.5103H14.0648C13.3646 19.9035 13 21.4411 13 23.0004C13 24.5596 13.3646 26.0973 14.0648 27.4905L17.3855 24.9142Z"
                    fill="#4471FF"
                  />
                  <path
                    d="M22.9994 16.9589C24.4254 16.9356 25.8034 17.4744 26.8354 18.4588L29.6894 15.6047C27.8796 13.9049 25.4821 12.9717 22.9994 13.0007C21.1458 13.0007 19.3288 13.516 17.7511 14.4889C16.1735 15.4618 14.8972 16.854 14.0649 18.5102L17.3857 21.0865C18.1773 18.7171 20.3911 16.9589 22.9994 16.9589Z"
                    fill="#4471FF"
                  />
                </svg>
              </div>
              <div className="pl-2">
                <div className="text-sm lg:text-base font-medium t-[#232323]">
                  Samsung Mobile
                </div>
                <div className="text-xs lg:text-sm text-[#718EBF]">
                  E-commerce, Marketplace
                </div>
              </div>
              <div>
                <div className="text-sm lg:text-base font-medium t-[#232323]">
                  $25,300
                </div>
                <div className="text-xs lg:text-sm text-[#718EBF]">
                  Envestment Value
                </div>
              </div>
              <div>
                <div className="text-sm lg:text-base font-medium t-[#232323] text-[#FE5C73]">
                  -4%
                </div>
                <div className="text-xs lg:text-sm text-[#718EBF]">
                  Return Value
                </div>
              </div>
            </div>
            <div className="flex bg-white rounded-2xl p-5 justify-between my-3">
              <div>
                <svg
                  width="45"
                  height="45"
                  viewBox="0 0 45 45"
                  fill="none"
                  xmlns="http://www.w3.org/2000/svg"
                >
                  <rect width="45" height="45" rx="15" fill="#FFF5D9" />
                  <path
                    d="M22.044 33.0002L24.8539 17.1968C27.5323 17.1968 28.3772 17.4905 28.4991 18.6893C28.4991 18.6893 30.2959 18.0194 31.2021 16.6587C27.6658 15.0201 24.1126 14.9462 24.1126 14.9462L22.0393 17.4715L22.044 17.4712L19.9706 14.9458C19.9706 14.9458 16.4174 15.0198 12.8816 16.6584C13.787 18.019 15.5845 18.689 15.5845 18.689C15.7072 17.4901 16.5509 17.1964 19.2114 17.1945L22.044 33.0002Z"
                    fill="#FFBB38"
                  />
                  <path
                    d="M22.0429 14.2166C24.9014 14.1948 28.1734 14.6588 31.5229 16.1187C31.9706 15.3129 32.0857 14.9568 32.0857 14.9568C28.4242 13.5082 24.9952 13.0124 22.0425 13C19.0898 13.0124 15.661 13.5083 12 14.9568C12 14.9568 12.1633 15.3955 12.5624 16.1187C15.9113 14.6588 19.1838 14.1948 22.0425 14.2166H22.0429Z"
                    fill="#FFBB38"
                  />
                </svg>
              </div>
              <div className="pl-2">
                <div className="text-sm lg:text-base font-medium t-[#232323]">
                  Tesla Motors
                </div>
                <div className="text-xs lg:text-sm text-[#718EBF]">
                  E-commerce, Marketplace
                </div>
              </div>
              <div>
                <div className="text-sm lg:text-base font-medium t-[#232323]">
                  $8,200
                </div>
                <div className="text-xs lg:text-sm text-[#718EBF]">
                  Envestment Value
                </div>
              </div>
              <div>
                <div className="text-sm lg:text-base font-medium t-[#232323] text-[#16DBAA]">
                  +25%
                </div>
                <div className="text-xs lg:text-sm text-[#718EBF]">
                  Return Value
                </div>
              </div>
            </div>
          </div>
        </div>
        <div className="col-span-12 lg:col-span-5">
          <div className="my-4 text-2xl font-bold text-[#333B69]">
            Trending Stock
          </div>
          <Table className="bg-white shadow-1 rounded-3xl">
            <TableHeader>
              <TableRow className="text-[#718EBF] p1">
                <TableHead className="w-[100px] text-[#718EBF] h1 hidden md:table-cell">
                  SL No
                </TableHead>
                <TableHead className="text-[#718EBF]">Name</TableHead>
                <TableHead className="text-[#718EBF]">Price</TableHead>
                <TableHead className="text-[#718EBF]">Return</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {invoices.map((invoice) => (
                <TableRow key={invoice.slno}>
                  <TableCell className="font-medium text-[#232323] px-5  h-2 p-3">
                    {invoice.slno}
                  </TableCell>
                  <TableCell className="font-medium text-[#232323] h-2 p-3">
                    {invoice.name}
                  </TableCell>
                  <TableCell className="font-medium text-[#232323] h-2 p-3">
                    {invoice.price}
                  </TableCell>
                  <TableCell
                    className={
                      invoice.slno === "03." || invoice.slno === "05."
                        ? "font-medium text-[#16DBAA] h-2 p-3"
                        : "font-medium text-[#FE5C73] h-2 p-3"
                    }
                  >
                    {invoice.return}
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </div>
      </div>
    </div>
    // </main>
  );
}
