"use client";
import Link from "next/link";
import { useState } from "react";
const Transactions = () => {
  const [selected, setSelected] = useState("All Transactions");
  const handleClick = (link: string) => {
    setSelected(link);
  };

  return (
    <div
      className="w-fit h-[923px] top-[101px] left-[250px] bg-[#F5F7FA] flex flex-col gap-6 p-6 "
      //   style={{ position: "relative" }}
    >
      <div className="flex justify-between">
        <div className="flex justify-between">
          <h5 className="w-[102px] h-[27px] top-[125px] left-[290px] font-inter text-[22px] font-semibold leading-[26.63px] text-left text-[#343C6A]">
            My Cards
          </h5>
          <button className="w-[93px] h-[21px] top-[130px] left-[927px] font-inter text-[17px] font-semibold leading-[20.57px] text-left text-[#343C6A]">
            + Add Card
          </button>
        </div>
        <div className="creditCards">
          <div className="creditcard"></div>
        </div>
        <div>
          <div className="w-[130px] h-[27px] top-[125px] left-[1050px] font-inter text-[22px] font-semibold leading-[26.63px] text-left text-[#343C6A]">
            My Expense
          </div>
          <div className="expenseCard"></div>
        </div>
      </div>

      <h5 className="recentTransactions  top-[421px] left-[290px] font-inter text-[22px] font-semibold leading-[26.63px] text-left text-[#343C6A]">
        Rececent Transactions
      </h5>
      <nav className="flex gap-20">
        <div className="flex flex-col gap-1 justify-center ">
          <Link
            href=""
            className={
              selected === "All Transactions"
                ? "w-[123px] h-[19px] top-[475px] left-[301px] font-inter text-[16px] font-medium leading-[19.36px] text-center text-[#1814F3]"
                : "text-w-[57px] h-[19px] top-[475px] left-[506px] font-inter text-[16px] font-medium leading-[19.36px] text-left text-[#718EBF]"
            }
            onClick={() => handleClick("All Transactions")}
          >
            All Transactions
          </Link>
          <div
            className={
              selected === "All Transactions"
                ? " h-[3px] top-[502px] left-[290px] gap-0 rounded-t-[10px] bg-[#1814F3]"
                : ""
            }
          ></div>
        </div>
        <div className="flex flex-col gap-1 justify-center">
          <Link
            href=""
            className={
              selected === "Incomes"
                ? "w-[123px] h-[19px] top-[475px] left-[301px] font-inter text-[16px] font-medium leading-[19.36px] text-center text-[#1814F3]"
                : "text-w-[57px] h-[19px] top-[475px] left-[506px] font-inter text-[16px] font-medium leading-[19.36px] text-left text-[#718EBF]"
            }
            onClick={() => handleClick("Incomes")}
          >
            Incomes
          </Link>
          <div
            className={
              selected === "Incomes"
                ? " h-[3px] top-[502px] left-[290px] gap-0 rounded-t-[10px] bg-[#1814F3]"
                : ""
            }
          ></div>
        </div>
        <div
          className="flex flex-col justify-center gap-1"
          onClick={() => handleClick("Expenses")}
        >
          <Link
            href=""
            className={
              selected === "Expenses"
                ? "w-[123px] h-[19px] top-[475px] left-[301px] font-inter text-[16px] font-medium leading-[19.36px] text-center text-[#1814F3]"
                : "text-w-[57px] h-[19px] top-[475px] left-[506px] font-inter text-[16px] font-medium leading-[19.36px] text-left text-[#718EBF]"
            }
          >
            Expenses
          </Link>
          <div
            className={
              selected === "Expenses"
                ? " h-[3px] top-[502px] left-[290px] gap-0 rounded-t-[10px] bg-[#1814F3]"
                : ""
            }
          ></div>
        </div>
      </nav>
      <div className="  text-left  p-8 rounded-[25px] bg-[#FFFF] ">
        <table className="w-[1110px] h-[397px]">
          <thead className="  text-[#718EBF]">
            <tr className="round-fullh-[35px]  ">
              <th className="font-inter text-[16px] font-medium leading-[19.36px] text-[#718EBF]">
                Description
              </th>
              <th className="font-inter text-[16px] font-medium leading-[19.36px] text-[#718EBF]">
                Transaction ID
              </th>
              <th className="font-inter text-[16px] font-medium leading-[19.36px] text-[#718EBF]">
                Type
              </th>
              <th className="font-inter text-[16px] font-medium leading-[19.36px] text-[#718EBF]">
                Card
              </th>
              <th className="font-inter text-[16px] font-medium leading-[19.36px] text-[#718EBF]">
                Date
              </th>
              <th className="font-inter text-[16px] font-medium leading-[19.36px] text-[#718EBF]">
                Amount
              </th>
              <th className="font-inter text-[16px] font-medium leading-[19.36px] text-[#718EBF]">
                Receipt
              </th>
            </tr>
          </thead>
          <tbody>
            <tr className="round-full h-[35px]  ">
              <td className="  ">
                <td className="flex gap-1 items-center">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke-width="1.5"
                    stroke="currentColor"
                    className="text-[#718EBF] w-[30px] h-[30px] "
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      d="m9 12.75 3 3m0 0 3-3m-3 3v-7.5M21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z"
                    />
                  </svg>

                  <p className="font-inter text-[16px]  font-normal leading-[19.36px]  text-[#232323]">
                    Spotify Subscription
                  </p>
                </td>
              </td>

              <td>
                <p className="font-inter text-[16px] font-normal leading-[19.36px] text-[#232323]">
                  #123546435
                </p>
              </td>
              <td>
                <p className="font-inter text-[16px] font-normal leading-[19.36px] text-[#232323]">
                  Shopping
                </p>
              </td>
              <td>
                <p className="font-inter text-[16px] font-normal leading-[19.36px] text-[#232323]">
                  123***
                </p>
              </td>
              <td>
                <p className="font-inter text-[16px] font-normal leading-[19.36px] text-[#232323]">
                  28 Jan, 12:30AM
                </p>
              </td>
              <td className="text-red-700">-$2,500</td>
              <td>
                <button className="w-[100px] h-[35px] top-[598px] left-[1270px] rounded-full border border-[#123288] text-[#1814F3]">
                  Download
                </button>
              </td>
            </tr>
            <tr className="round-full h-[35px]  ">
              <td className="  ">
                <td className="flex gap-1 items-center">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke-width="1.5"
                    stroke="currentColor"
                    className="text-[#718EBF] w-[30px] h-[30px] "
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      d="m9 12.75 3 3m0 0 3-3m-3 3v-7.5M21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z"
                    />
                  </svg>

                  <p className="font-inter text-[16px]  font-normal leading-[19.36px]  text-[#232323]">
                    Spotify Subscription
                  </p>
                </td>
              </td>

              <td>
                <p className="font-inter text-[16px] font-normal leading-[19.36px] text-[#232323]">
                  #123546435
                </p>
              </td>
              <td>
                <p className="font-inter text-[16px] font-normal leading-[19.36px] text-[#232323]">
                  Shopping
                </p>
              </td>
              <td>
                <p className="font-inter text-[16px] font-normal leading-[19.36px] text-[#232323]">
                  123***
                </p>
              </td>
              <td>
                <p className="font-inter text-[16px] font-normal leading-[19.36px] text-[#232323]">
                  28 Jan, 12:30AM
                </p>
              </td>
              <td className="text-red-700">-$2,500</td>
              <td>
                <button className="w-[100px] h-[35px] top-[598px] left-[1270px] rounded-full border border-[#123288] text-[#1814F3]">
                  Download
                </button>
              </td>
            </tr>
            <tr className="round-full h-[35px]  ">
              <td className="  ">
                <td className="flex gap-1 items-center">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke-width="1.5"
                    stroke="currentColor"
                    className="text-[#718EBF] w-[30px] h-[30px] "
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      d="m9 12.75 3 3m0 0 3-3m-3 3v-7.5M21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z"
                    />
                  </svg>

                  <p className="font-inter text-[16px]  font-normal leading-[19.36px]  text-[#232323]">
                    Spotify Subscription
                  </p>
                </td>
              </td>

              <td>
                <p className="font-inter text-[16px] font-normal leading-[19.36px] text-[#232323]">
                  #123546435
                </p>
              </td>
              <td>
                <p className="font-inter text-[16px] font-normal leading-[19.36px] text-[#232323]">
                  Shopping
                </p>
              </td>
              <td>
                <p className="font-inter text-[16px] font-normal leading-[19.36px] text-[#232323]">
                  123***
                </p>
              </td>
              <td>
                <p className="font-inter text-[16px] font-normal leading-[19.36px] text-[#232323]">
                  28 Jan, 12:30AM
                </p>
              </td>
              <td className="text-red-700">-$2,500</td>
              <td>
                <button className="w-[100px] h-[35px] top-[598px] left-[1270px] rounded-full border border-[#123288] text-[#1814F3]">
                  Download
                </button>
              </td>
            </tr>

            <tr className="round-full h-[35px]  ">
              <td className="  ">
                <td className="flex gap-1 items-center">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke-width="1.5"
                    stroke="currentColor"
                    className="text-[#718EBF] w-[30px] h-[30px] "
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      d="m9 12.75 3 3m0 0 3-3m-3 3v-7.5M21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z"
                    />
                  </svg>

                  <p className="font-inter text-[16px]  font-normal leading-[19.36px]  text-[#232323]">
                    Spotify Subscription
                  </p>
                </td>
              </td>

              <td>
                <p className="font-inter text-[16px] font-normal leading-[19.36px] text-[#232323]">
                  #123546435
                </p>
              </td>
              <td>
                <p className="font-inter text-[16px] font-normal leading-[19.36px] text-[#232323]">
                  Shopping
                </p>
              </td>
              <td>
                <p className="font-inter text-[16px] font-normal leading-[19.36px] text-[#232323]">
                  123***
                </p>
              </td>
              <td>
                <p className="font-inter text-[16px] font-normal leading-[19.36px] text-[#232323]">
                  28 Jan, 12:30AM
                </p>
              </td>
              <td className="text-red-700">-$2,500</td>
              <td>
                <button className="w-[100px] h-[35px] top-[598px] left-[1270px] rounded-full border border-[#123288] text-[#1814F3]">
                  Download
                </button>
              </td>
            </tr>
            <tr className="round-full h-[35px]  ">
              <td className="  ">
                <td className="flex gap-1 items-center">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke-width="1.5"
                    stroke="currentColor"
                    className="text-[#718EBF] w-[30px] h-[30px] "
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      d="m9 12.75 3 3m0 0 3-3m-3 3v-7.5M21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z"
                    />
                  </svg>

                  <p className="font-inter text-[16px]  font-normal leading-[19.36px]  text-[#232323]">
                    Spotify Subscription
                  </p>
                </td>
              </td>

              <td>
                <p className="font-inter text-[16px] font-normal leading-[19.36px] text-[#232323]">
                  #123546435
                </p>
              </td>
              <td>
                <p className="font-inter text-[16px] font-normal leading-[19.36px] text-[#232323]">
                  Shopping
                </p>
              </td>
              <td>
                <p className="font-inter text-[16px] font-normal leading-[19.36px] text-[#232323]">
                  123***
                </p>
              </td>
              <td>
                <p className="font-inter text-[16px] font-normal leading-[19.36px] text-[#232323]">
                  28 Jan, 12:30AM
                </p>
              </td>
              <td className="text-red-700">-$2,500</td>
              <td>
                <button className="w-[100px] h-[35px] top-[598px] left-[1270px] rounded-full border border-[#123288] text-[#1814F3]">
                  Download
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  );
};

export default Transactions;
