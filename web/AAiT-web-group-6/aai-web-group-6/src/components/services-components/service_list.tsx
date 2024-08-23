import React from "react";
import Services from "./service";
import ServiceCard from "./card";

export default function ServicesPage() {
  return (
    <div className="px-5 md:px-10 py-7 bg-[rgba(245,247,250,1)]">
      <div className="flex mx-6 overflow-x-scroll space-x-4 md: overscroll-none justify-between  py-6 ">
        <ServiceCard
          text1="Life Insurance"
          text2="Unlimited Protection"
          imageSrc="/images/life-insurance_card.png"
          imageBackground="231,237,255,1"
        />
        <ServiceCard
          text1="Shopping"
          text2="Buy. Think. Grow."
          imageSrc="/images/shopping.png"
          imageBackground="255,245,217,1"
        />
        <ServiceCard
          text1="Safety"
          text2="We are your allies"
          imageSrc="/images/safety.png"
          imageBackground="220,250,248,1"
        />
      </div>

      <p className="font-semibold text-[rgba(52,60,106,1)] text-xl">
        Bank Services List
      </p>
      <Services />
    </div>
  );
}
