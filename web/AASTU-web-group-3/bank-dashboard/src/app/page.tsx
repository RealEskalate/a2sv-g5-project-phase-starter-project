"use client";
import React from "react";
import Image from "next/image";
import img from "./images/profile.png";
import emanuel from "../../public/images/emanuel-minca-jYv069cQuB8-unsplash 1.png";
import julia from "../../public/images/pexels-julia-volk-5273755 1.png";
import marcel from "../../public/images/marcel-strauss-Uc_tOqa_jDY-unsplash 1.png";
import paypal from "../../public/images/iconfinder_paypal_payment_pay_5340264 1.png";
import deposit from "../../public/images/iconfinder_business_finance_money-13_2784281 1.png";
import dollar from "../../public/images/iconfinder_6_4753731 1.png";
import Dashboard from "../app/dashboard/page";
import { AreaChart } from "lucide-react";
import { AreaChartComponent } from "./components/Chart/AreaChartComponent";
import { BarChartComponent } from "./components/Chart/Barchart";

const imageData = [
  { src: julia.src, alt: "julia", name: "Livia Bator", position: "CEO" },
  { src: marcel.src, alt: "marcel", name: "Randy Press", position: "Director" },
  { src: emanuel.src, alt: "emanuel", name: "Workman", position: "Designer" },
];

const transactionData = [
  {
    src: deposit.src,
    alt: "deposit",
    backgroundColor: "#FFF5D9",
    title: "Deposit from my",
    date: "28 January 2021",
    amount: "-$850",
    amountColor: "text-red-500",
  },
  {
    src: paypal.src,
    alt: "paypal",
    backgroundColor: "#E7EDFF",
    title: "Deposit Paypal",
    date: "25 January 2021",
    amount: "+$2,500",
    amountColor: "text-green-500",
  },
  {
    src: dollar.src,
    alt: "dollar",
    backgroundColor: "#DCFAF8",
    title: "Jemi Wilson",
    date: "21 January 2021",
    amount: "+$5,400",
    amountColor: "text-green-500",
  },
];

const HomePage: React.FC = () => {
  return <BarChartComponent />;
};

export default HomePage;
