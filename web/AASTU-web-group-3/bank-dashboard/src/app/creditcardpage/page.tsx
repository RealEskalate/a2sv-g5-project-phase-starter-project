"use client";
import Image from "next/image";
import DoughnutChart from "../components/charts/Doughnut/page";
import React from "react";
import CardListPage from "../components/cardList/CardList";

const creditCardPage = () => {
  return (
    <div className="body bg-[#F5F7FA] w-full h-[100vh] overflow-y-auto pb-5">
      {/* <div className="nav w-[375px] h-[140px]  bg-white rounded-[15px]  ">
        Navigation bar
      </div> */}
      <div className="p-5 font-semibold text-[16px] leading-[19.36px] text-[#343C6A]">
        My Cards
      </div>
      <div className=" h-44 border-2 p-5  bg-white rounded-[15px] mx-5">
        My card import here
      </div>
      <div className="statandlist md:flex">
        <div className="md:w-[35%]">
          <p className=" p-4 font-semibold text-[16px] leading-[19.36px] text-[#343C6A] mx-2">
            Card Expense Statistics
          </p>
          <div className="piechart md:w-auto md:h-auto flex justify-center  md:ml-5 md:px-0 border-[1px] w-auto h-[300px] bg-white rounded-[15px] mx-5  px-8">
            <DoughnutChart />
          </div>
        </div>

        <div className="md:w-[65%] md:mr-3  ">
          <p className="p-4 md:pb-2 md:mx-0 font-semibold text-[16px] leading-[19.36px] text-[#343C6A] mx-2 ">
            Card List
          </p>
          <div className="cardList md:mx-0 w-auto h-auto  mx-4 ">
            <CardListPage />
          </div>
        </div>
      </div>

      <div className="addnewandcardsetting md:flex md:flex-row md:w-full ">
        <div className="addnewcard flex flex-col md:w-[65%]">
          <div className="p-5 font-semibold text-[16px] leading-[19.36px] text-[#343C6A] mx-2">
            Add New Card
          </div>
          <form className="newcard flex flex-col md:flex md:flex-row md:flex-wrap md:w-auto  md:h-[321px] justify-between  w-auto h-[527px]  p-4 border-2 bg-white rounded-[15px] mx-6">
            <p className="description text-[#718EBF] text-[12px] leading-[22px]">
              Credit Card generally means a plastic card issued by Scheduled
              Commercial Banks assigned to a Cardholder, with a credit limit,
              that can be used to purchase goods and services on credit or
              obtain cash advances.
            </p>
            <div className="flex flex-col md:w-[50%] ">
              <label htmlFor="CardType" className="md:text-xs md:font-normal">
                Card Type
              </label>
              <input
                type="text"
                name="CardType"
                id="CardType"
                placeholder="Classic"
                className="border-[1px] md:w-[90%] w-auto  h-[40px] rounded-[10px] p-3 md:text-xs"
              />
            </div>
            <div className="flex flex-col md:w-[50%] ">
              <label htmlFor="nameoncard" className="md:text-xs md:font-normal">
                Name On Card
              </label>
              <input
                type="text"
                name="nameoncard"
                id="nameoncard"
                placeholder="My Cards"
                className="border-[1px] md:w-[90%] w-auto h-[40px] rounded-[10px] p-3 md:text-xs"
              />
            </div>
            <div className="flex flex-col md:w-[50%] ">
              <label htmlFor="cardnumber" className="md:text-xs md:font-normal">
                Card Number
              </label>
              <input
                type="text"
                name="cardnumber"
                id="cardnumber"
                placeholder="**** **** **** ****"
                className="border-[1px] md:w-[90%] w-auto h-[40px] rounded-[10px] p-3 md:text-xs"
              />
            </div>
            <div className="flex flex-col md:w-[50%] ">
              <label
                htmlFor="expirationdate"
                className="md:text-xs md:font-normal"
              >
                Expiration Date
              </label>
              <input
                type=" date"
                name="expirationdate"
                id="expirationdate"
                placeholder="25 January 2025"
                className="border-[1px] md:w-[90%] w-auto h-[40px] rounded-[10px] p-3 md:text-xs"
              />
            </div>

            <button
              type="submit"
              className="border-2 md:w-[30%] bg-[#1814F3] text-white w-auto  h-[40px] rounded-[9px]"
            >
              Add To Cart
            </button>
          </form>
        </div>
        <div className="card setting md:w-[35%]  md:mr-8">
          <div className="p-5 font-semibold text-[16px] leading-[19.36px] text-[#343C6A] mx-2 ">
            Card Setting
          </div>
          <div className="flex flex-col justify-around self-center md:w-full md:mx-2 w-auto p-4 h-[325px] border-[1px] rounded-[15px] bg-white mx-6 ">
            <div className="blockcard md:pl-0 flex w-auto h-[45px] pl-5  ">
              <div className="left">
                <svg
                  width="45"
                  height="45"
                  viewBox="0 0 45 45"
                  fill="none"
                  xmlns="http://www.w3.org/2000/svg"
                >
                  <rect width="45" height="45" rx="12" fill="#FFF5D9" />
                  <path
                    d="M31.9036 22.2365V16.3049C31.9036 15.087 30.9127 14.0962 29.6949 14.0962H15.2086C13.9908 14.0962 13 15.087 13 16.3049V25.3089C13 26.5268 13.9908 27.5176 15.2086 27.5176H20.9365C21.7011 30.0522 24.057 31.9034 26.8375 31.9034C30.2355 31.9034 32.9999 29.139 32.9999 25.7409C33 24.4399 32.5944 23.2321 31.9036 22.2365ZM23.3331 20.6748H14.3605V18.7461H30.543V20.8202C29.511 20.0411 28.2273 19.5785 26.8375 19.5785C25.5366 19.5784 24.3288 19.984 23.3331 20.6748ZM15.2086 15.4567H29.6949C30.1626 15.4567 30.543 15.8372 30.543 16.3049V17.3856H14.3605V16.3049C14.3605 15.8372 14.741 15.4567 15.2086 15.4567ZM15.2086 26.1571C14.741 26.1571 14.3605 25.7766 14.3605 25.3089V22.0354H21.9168C21.1377 23.0674 20.6751 24.3511 20.6751 25.7409C20.6751 25.8808 20.6802 26.0195 20.6894 26.1571H15.2086V26.1571ZM26.8376 30.5428C24.1898 30.5428 22.0356 28.3887 22.0356 25.7408C22.0356 23.0931 24.1898 20.939 26.8376 20.939C29.4853 20.939 31.6395 23.0931 31.6395 25.7408C31.6395 28.3887 29.4853 30.5428 26.8376 30.5428Z"
                    fill="#FFBB38"
                  />
                  <path
                    d="M23.5482 23.9644C23.1725 23.9644 22.8679 24.2689 22.8679 24.6446V26.8375C22.8679 27.2132 23.1725 27.5178 23.5482 27.5178H30.1268C30.5025 27.5178 30.8071 27.2132 30.8071 26.8375V24.6446C30.8071 24.2689 30.5025 23.9644 30.1268 23.9644H23.5482ZM29.4466 26.1573H24.2285V25.325H29.4466V26.1573Z"
                    fill="#FFBB38"
                  />{" "}
                </svg>
              </div>
              <div className="right flex-row  w-auto h-[36px] p-2 pl-3">
                <div className="w-[73px] font-medium text-sm">Block Card</div>
                <div className="w-[140px] font-normal text-xs text-[#718EBF]">
                  Instantly block your card
                </div>
              </div>
            </div>

            <div className="blockcard md:pl-0 flex w-auto h-[45px] pl-5 ">
              <div className="left">
                <svg
                  width="45"
                  height="45"
                  viewBox="0 0 45 45"
                  fill="none"
                  xmlns="http://www.w3.org/2000/svg"
                >
                  <rect width="45" height="45" rx="12" fill="#E7EDFF" />
                  <path
                    d="M28.625 20.3438H27.2173V17.1316C27.2173 14.8534 25.3248 13 22.9986 13C20.6723 13 18.7798 14.8534 18.7798 17.1316V20.3438H17.375C16.0827 20.3438 15.0312 21.3952 15.0312 22.6875V30.6562C15.0312 31.9486 16.0827 33 17.375 33H28.625C29.9173 33 30.9688 31.9486 30.9688 30.6562V22.6875C30.9688 21.3952 29.9173 20.3438 28.625 20.3438ZM20.3423 17.1316C20.3423 15.715 21.5339 14.5625 22.9986 14.5625C24.4632 14.5625 25.6548 15.715 25.6548 17.1316V20.3438H20.3423V17.1316ZM29.4062 30.6562C29.4062 31.087 29.0558 31.4375 28.625 31.4375H17.375C16.9442 31.4375 16.5938 31.087 16.5938 30.6562V22.6875C16.5938 22.2567 16.9442 21.9062 17.375 21.9062H28.625C29.0558 21.9062 29.4062 22.2567 29.4062 22.6875V30.6562Z"
                    fill="#396AFF"
                  />
                  <path
                    d="M23 24.1719C22.2018 24.1719 21.5547 24.8189 21.5547 25.6172C21.5547 26.1269 21.8187 26.5746 22.2173 26.832V28.5469C22.2173 28.9783 22.5671 29.3281 22.9986 29.3281C23.43 29.3281 23.7798 28.9783 23.7798 28.5469V26.8338C24.18 26.5768 24.4453 26.1281 24.4453 25.6172C24.4453 24.8189 23.7982 24.1719 23 24.1719Z"
                    fill="#396AFF"
                  />
                </svg>
              </div>
              <div className="right flex-row  w-auto h-[36px] p-2 pl-3">
                <div className="w-[116px] font-medium text-sm">
                  Change Pin Code
                </div>
                <div className="w-[140px] font-normal text-[11px] text-[#718EBF]">
                  Withdraw without any card
                </div>
              </div>
            </div>

            <div className="blockcard md:pl-0 flex w-auto h-[45px] pl-5 ">
              <div className="left">
                <svg
                  width="45"
                  height="45"
                  viewBox="0 0 45 45"
                  fill="none"
                  xmlns="http://www.w3.org/2000/svg"
                >
                  <rect width="45" height="45" rx="12" fill="#FFE0EB" />
                  <g clip-path="url(#clip0_196_642)">
                    <path
                      d="M24.25 22V24.5H30.15C29.55 28 26.75 30.5 23.25 30.5C19.15 30.5 15.75 27.1 15.75 23C15.75 18.9 19.15 15.5 23.25 15.5C25.35 15.5 27.15 16.4 28.45 17.8L30.25 16C28.45 14.2 26.05 13 23.25 13C17.75 13 13.25 17.5 13.25 23C13.25 28.5 17.75 33 23.25 33C28.75 33 32.75 28.5 32.75 23V22H24.25Z"
                      fill="#FF82AC"
                    />
                  </g>
                  <defs>
                    <clipPath id="clip0_196_642">
                      <rect
                        width="20"
                        height="20"
                        fill="white"
                        transform="translate(13 13)"
                      />
                    </clipPath>
                  </defs>
                </svg>
              </div>
              <div className="right flex-row  w-auto h-[36px] p-2 pl-3">
                <div className="w-[125px] font-medium text-sm">
                  Add to Google Pay
                </div>
                <div className="w-[140px] font-normal text-[11px] text-[#718EBF]">
                  Withdraw without any card
                </div>
              </div>
            </div>

            <div className="blockcard md:pl-0 flex w-auto h-[45px] pl-5 ">
              <div className="left">
                <Image
                  src="/images/apple.png"
                  alt="applepay"
                  width={45}
                  height={45}
                />
              </div>
              <div className="right flex-row  w-auto h-[36px] p-2 pl-3">
                <div className="w-[125px] font-medium text-sm">
                  Add to Apple Pay
                </div>
                <div className="w-[140px] font-normal text-[11px] text-[#718EBF]">
                  Withdraw without any card
                </div>
              </div>
            </div>

            <div className="blockcard md:pl-0 flex w-auto h-[45px] pl-5 ">
              <div className="left">
                <Image
                  src="/images/apple.png"
                  alt="applepay"
                  width={45}
                  height={45}
                />
              </div>
              <div className="right flex-row  w-auto h-[36px] p-2 pl-3">
                <div className="w-[127px] font-medium text-sm">
                  Add to Apple Store
                </div>
                <div className="w-[140px] font-normal text-[11px] text-[#718EBF]">
                  Withdraw without any card
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default creditCardPage;
